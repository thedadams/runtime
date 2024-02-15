package appdefinition

import (
	"strconv"
	"strings"

	"github.com/acorn-io/baaah/pkg/apply"
	"github.com/acorn-io/baaah/pkg/router"
	"github.com/acorn-io/baaah/pkg/typed"
	v1 "github.com/acorn-io/runtime/pkg/apis/internal.acorn.io/v1"
	"github.com/acorn-io/runtime/pkg/jobs"
	"github.com/acorn-io/runtime/pkg/labels"
	"github.com/acorn-io/runtime/pkg/publicname"
	"github.com/acorn-io/runtime/pkg/secrets"
	"github.com/acorn-io/runtime/pkg/system"
	"github.com/acorn-io/z"
	"github.com/google/go-containerregistry/pkg/name"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func stripPruneAndUpdate(annotations map[string]string) map[string]string {
	result := map[string]string{}
	for k, v := range annotations {
		if k == apply.AnnotationPrune || k == apply.AnnotationUpdate {
			continue
		}
		result[k] = v
	}
	return result
}

func toJobs(req router.Request, appInstance *v1.AppInstance, pullSecrets *PullSecrets, tag name.Reference, interpolator *secrets.Interpolator) (result []kclient.Object, _ error) {
	for _, jobName := range typed.SortedKeys(appInstance.Status.AppSpec.Jobs) {
		jobDef := appInstance.Status.AppSpec.Jobs[jobName]
		jobDef, err := augmentContainerWithConsumerInfo(req.Ctx, req.Client, appInstance.Status.Namespace, jobDef)
		if err != nil {
			return nil, err
		}
		addBusybox := false
		for _, v := range jobDef.Dirs {
			if v.Preload {
				addBusybox = true
			}
		}
		jbs, err := toJobAndCronJob(req, appInstance, pullSecrets, tag, jobName, jobDef, interpolator, addBusybox)
		if err != nil {
			return nil, err
		}
		if len(jbs) == 0 {
			continue
		}

		job := jbs[0]
		perms := v1.FindPermission(jobName, appInstance.Status.Permissions)
		sa, err := toServiceAccount(req, job.GetName(), job.GetLabels(), stripPruneAndUpdate(job.GetAnnotations()), appInstance, perms)
		if err != nil {
			return nil, err
		}
		if perms.HasRules() {
			perms, err := toPermissions(req.Ctx, req.Client, perms, job.GetLabels(), stripPruneAndUpdate(job.GetAnnotations()), appInstance)
			if err != nil {
				return nil, err
			}
			result = append(result, perms...)
		}
		result = append(result, sa)
		result = append(result, jbs...)
	}

	return result, nil
}

func setJobEventName(containers []corev1.Container, eventName string) (result []corev1.Container) {
	for _, c := range containers {
		c.Env = append(c.Env, corev1.EnvVar{
			Name:  "ACORN_EVENT",
			Value: eventName,
		})
		result = append(result, c)
	}
	return
}

func setSecretOutputVolume(containers []corev1.Container) (result []corev1.Container) {
	for _, c := range containers {
		c.VolumeMounts = append([]corev1.VolumeMount{
			{Name: jobs.Helper, MountPath: "/run/secrets"},
		}, c.VolumeMounts...)
		result = append(result, c)
	}
	return
}

func toJobAndCronJob(req router.Request, appInstance *v1.AppInstance, pullSecrets *PullSecrets, tag name.Reference, name string, container v1.Container, interpolator *secrets.Interpolator, addBusybox bool) ([]kclient.Object, error) {
	var result []kclient.Object
	interpolator = interpolator.ForJob(name)
	jobEventName := jobs.GetEvent(name, appInstance)

	jobStatus := appInstance.Status.AppStatus.Jobs[name]
	jobStatus.Skipped = !jobs.ShouldRunForEvent(jobEventName, container)
	appInstance.Status.AppStatus.Jobs = z.AddToMap(appInstance.Status.AppStatus.Jobs, name, jobStatus)

	if jobStatus.Skipped && container.Schedule == "" {
		return nil, nil
	}

	containers, initContainers := toContainers(appInstance, tag, name, container, interpolator, false, addBusybox)

	containers = append(containers, corev1.Container{
		Name:            jobs.Helper,
		Image:           system.DefaultImage(),
		Command:         []string{"/usr/local/bin/acorn-job-helper-init"},
		ImagePullPolicy: corev1.PullIfNotPresent,
	})

	secretAnnotations, err := getSecretAnnotations(req, appInstance, container, interpolator)
	if err != nil {
		return nil, err
	}

	volumes, err := toVolumes(appInstance, container, interpolator, false, addBusybox)
	if err != nil {
		return nil, err
	}

	baseAnnotations := labels.Merge(
		secretAnnotations,
		labels.GatherScoped(name, v1.LabelTypeJob, appInstance.Status.AppSpec.Annotations, container.Annotations, appInstance.Spec.Annotations),
	)

	baseAnnotations[labels.AcornConfigHashAnnotation] = appInstance.Status.AppStatus.Jobs[name].ConfigHash
	baseAnnotations[labels.AcornAppGeneration] = strconv.FormatInt(appInstance.Generation, 10)

	podLabels, err := jobLabels(appInstance, container, name, interpolator,
		labels.AcornManaged, "true",
		labels.AcornAppPublicName, publicname.Get(appInstance),
		labels.AcornJobName, name,
		labels.AcornContainerName, "",
	)
	if err != nil {
		return nil, err
	}

	jobSpec := batchv1.JobSpec{
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels:      podLabels,
				Annotations: labels.Merge(podAnnotations(appInstance, name, container), baseAnnotations),
			},
			Spec: corev1.PodSpec{
				Affinity:                      appInstance.Status.Scheduling[name].Affinity,
				Tolerations:                   appInstance.Status.Scheduling[name].Tolerations,
				RuntimeClassName:              stringOrNilPtr(appInstance.Status.Scheduling[name].RuntimeClassName),
				TerminationGracePeriodSeconds: z.Pointer[int64](5),
				ImagePullSecrets:              pullSecrets.ForContainer(name, append(containers, initContainers...)),
				EnableServiceLinks:            new(bool),
				RestartPolicy:                 corev1.RestartPolicyNever,
				Containers:                    setSecretOutputVolume(containers),
				InitContainers:                setSecretOutputVolume(initContainers),
				Volumes: append(volumes, corev1.Volume{
					Name: jobs.Helper,
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{
							Medium:    corev1.StorageMediumMemory,
							SizeLimit: resource.NewScaledQuantity(1, resource.Mega),
						},
					},
				}),
				ServiceAccountName: name,
			},
		},
	}

	objectMeta := metav1.ObjectMeta{
		Name:        name,
		Namespace:   appInstance.Status.Namespace,
		Labels:      jobSpec.Template.Labels,
		Annotations: labels.Merge(getDependencyAnnotations(appInstance, name, container.Dependencies), baseAnnotations),
	}

	interpolator.AddMissingAnnotations(appInstance.GetStopped(), baseAnnotations)

	if container.Schedule == "" || !jobStatus.Skipped {
		job := &batchv1.Job{
			ObjectMeta: *objectMeta.DeepCopy(),
			Spec:       *jobSpec.DeepCopy(),
		}

		job.Spec.BackoffLimit = z.Pointer[int32](1000)
		job.Spec.Template.Spec.Containers = setJobEventName(setSecretOutputVolume(containers), jobEventName)
		job.Spec.Template.Spec.InitContainers = setJobEventName(setSecretOutputVolume(initContainers), jobEventName)
		job.Annotations[apply.AnnotationPrune] = "false"
		if job.Annotations[apply.AnnotationUpdate] == "" {
			// getDependencyAnnotations may set this annotation, so don't override here
			job.Annotations[apply.AnnotationUpdate] = "true"
		}

		result = append(result, job)
	}

	if container.Schedule != "" {
		result = append(result, &batchv1.CronJob{
			ObjectMeta: objectMeta,
			Spec: batchv1.CronJobSpec{
				FailedJobsHistoryLimit:     z.Pointer[int32](3),
				SuccessfulJobsHistoryLimit: z.Pointer[int32](1),
				ConcurrencyPolicy:          batchv1.ReplaceConcurrent,
				Schedule:                   toCronJobSchedule(container.Schedule),
				JobTemplate: batchv1.JobTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: jobSpec.Template.Labels,
					},
					Spec: jobSpec,
				},
			},
		})
	}

	return result, nil
}

func toCronJobSchedule(schedule string) string {
	switch strings.TrimSpace(schedule) {
	case "year":
	case "annually":
	case "monthly":
	case "weekly":
	case "daily":
	case "midnight":
	case "hourly":
	default:
		return schedule
	}
	return "@" + strings.TrimSpace(schedule)
}
