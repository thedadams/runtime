//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppImage) DeepCopyInto(out *AppImage) {
	*out = *in
	in.ImageData.DeepCopyInto(&out.ImageData)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppImage.
func (in *AppImage) DeepCopy() *AppImage {
	if in == nil {
		return nil
	}
	out := new(AppImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppInstance) DeepCopyInto(out *AppInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppInstance.
func (in *AppInstance) DeepCopy() *AppInstance {
	if in == nil {
		return nil
	}
	out := new(AppInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppInstanceList) DeepCopyInto(out *AppInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppInstanceList.
func (in *AppInstanceList) DeepCopy() *AppInstanceList {
	if in == nil {
		return nil
	}
	out := new(AppInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppInstanceSpec) DeepCopyInto(out *AppInstanceSpec) {
	*out = *in
	if in.Stop != nil {
		in, out := &in.Stop, &out.Stop
		*out = new(bool)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]VolumeBinding, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppInstanceSpec.
func (in *AppInstanceSpec) DeepCopy() *AppInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(AppInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppInstanceStatus) DeepCopyInto(out *AppInstanceStatus) {
	*out = *in
	in.AppImage.DeepCopyInto(&out.AppImage)
	in.AppSpec.DeepCopyInto(&out.AppSpec)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[string]Condition, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppInstanceStatus.
func (in *AppInstanceStatus) DeepCopy() *AppInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(AppInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpec) DeepCopyInto(out *AppSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make(map[string]Container, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(map[string]Image, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make(map[string]VolumeRequest, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]Secret, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpec.
func (in *AppSpec) DeepCopy() *AppSpec {
	if in == nil {
		return nil
	}
	out := new(AppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Build) DeepCopyInto(out *Build) {
	*out = *in
	if in.ContextDirs != nil {
		in, out := &in.ContextDirs, &out.ContextDirs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Build.
func (in *Build) DeepCopy() *Build {
	if in == nil {
		return nil
	}
	out := new(Build)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderSpec) DeepCopyInto(out *BuilderSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make(map[string]ContainerImageBuilderSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(map[string]ImageBuilderSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderSpec.
func (in *BuilderSpec) DeepCopy() *BuilderSpec {
	if in == nil {
		return nil
	}
	out := new(BuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	if in.Dirs != nil {
		in, out := &in.Dirs, &out.Dirs
		*out = make(map[string]VolumeMount, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make(map[string]File, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(Build)
		(*in).DeepCopyInto(*out)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Entrypoint != nil {
		in, out := &in.Entrypoint, &out.Entrypoint
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make(map[string]Container, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerData) DeepCopyInto(out *ContainerData) {
	*out = *in
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make(map[string]ImageData, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerData.
func (in *ContainerData) DeepCopy() *ContainerData {
	if in == nil {
		return nil
	}
	out := new(ContainerData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerImageBuilderSpec) DeepCopyInto(out *ContainerImageBuilderSpec) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(Build)
		(*in).DeepCopyInto(*out)
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make(map[string]ContainerImageBuilderSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerImageBuilderSpec.
func (in *ContainerImageBuilderSpec) DeepCopy() *ContainerImageBuilderSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerImageBuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvSecretVal) DeepCopyInto(out *EnvSecretVal) {
	*out = *in
	if in.Optional != nil {
		in, out := &in.Optional, &out.Optional
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvSecretVal.
func (in *EnvSecretVal) DeepCopy() *EnvSecretVal {
	if in == nil {
		return nil
	}
	out := new(EnvSecretVal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
	in.Secret.DeepCopyInto(&out.Secret)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSecret) DeepCopyInto(out *FileSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSecret.
func (in *FileSecret) DeepCopy() *FileSecret {
	if in == nil {
		return nil
	}
	out := new(FileSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericMap.
func (in GenericMap) DeepCopy() GenericMap {
	if in == nil {
		return nil
	}
	out := new(GenericMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(Build)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilderSpec) DeepCopyInto(out *ImageBuilderSpec) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(Build)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilderSpec.
func (in *ImageBuilderSpec) DeepCopy() *ImageBuilderSpec {
	if in == nil {
		return nil
	}
	out := new(ImageBuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageData) DeepCopyInto(out *ImageData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageData.
func (in *ImageData) DeepCopy() *ImageData {
	if in == nil {
		return nil
	}
	out := new(ImageData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagesData) DeepCopyInto(out *ImagesData) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make(map[string]ContainerData, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(map[string]ImageData, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagesData.
func (in *ImagesData) DeepCopy() *ImagesData {
	if in == nil {
		return nil
	}
	out := new(ImagesData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	out.Params = in.Params.DeepCopy()
	if in.Optional != nil {
		in, out := &in.Optional, &out.Optional
		*out = new(bool)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBinding) DeepCopyInto(out *VolumeBinding) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBinding.
func (in *VolumeBinding) DeepCopy() *VolumeBinding {
	if in == nil {
		return nil
	}
	out := new(VolumeBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeMount) DeepCopyInto(out *VolumeMount) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeMount.
func (in *VolumeMount) DeepCopy() *VolumeMount {
	if in == nil {
		return nil
	}
	out := new(VolumeMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeRequest) DeepCopyInto(out *VolumeRequest) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]AccessMode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeRequest.
func (in *VolumeRequest) DeepCopy() *VolumeRequest {
	if in == nil {
		return nil
	}
	out := new(VolumeRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSecretMount) DeepCopyInto(out *VolumeSecretMount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSecretMount.
func (in *VolumeSecretMount) DeepCopy() *VolumeSecretMount {
	if in == nil {
		return nil
	}
	out := new(VolumeSecretMount)
	in.DeepCopyInto(out)
	return out
}
