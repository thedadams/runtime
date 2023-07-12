package profiles

import (
	apiv1 "github.com/acorn-io/runtime/pkg/apis/api.acorn.io/v1"
	"github.com/acorn-io/z"
)

// Default values
var (
	ClusterDomainDefault         = ".local.oss-acorn.io"
	InternalClusterDomainDefault = "svc.cluster.local"

	AcornDNSEndpointDefault = "https://oss-dns.acrn.io/v1"
	AcornDNSStateDefault    = "auto"

	// LetsEncryptOptionDefault is the default state for the Let's Encrypt integration
	LetsEncryptOptionDefault = "disabled"

	// AutoUpgradeIntervalDefault is the default value for the DefaultImageCheckInterval field
	AutoUpgradeIntervalDefault = "5m"

	// HttpEndpointPatternDefault is a pattern that works with Let's Encrypt
	HttpEndpointPatternDefault = "{{hashConcat 8 .Container .App .Namespace | truncate}}.{{.ClusterDomain}}"

	// Features
	FeatureImageAllowRules = "image-allow-rules"
	FeatureDefaults        = map[string]bool{
		FeatureImageAllowRules: false,
	}
)

func defaultProfile() apiv1.Config {
	return apiv1.Config{
		AcornDNS:                       z.P(AcornDNSStateDefault),
		AcornDNSEndpoint:               z.P(AcornDNSEndpointDefault),
		AutoUpgradeInterval:            z.P(AutoUpgradeIntervalDefault),
		AWSIdentityProviderARN:         new(string),
		BuilderPerProject:              new(bool),
		CertManagerIssuer:              new(string),
		EventTTL:                       new(string),
		Features:                       FeatureDefaults,
		HttpEndpointPattern:            z.P(HttpEndpointPatternDefault),
		IgnoreUserLabelsAndAnnotations: new(bool),
		IngressClassName:               new(string),
		IngressControllerNamespace:     new(string),
		InternalClusterDomain:          InternalClusterDomainDefault,
		InternalRegistryPrefix:         new(string),
		LetsEncrypt:                    z.P(LetsEncryptOptionDefault),
		LetsEncryptEmail:               "",
		LetsEncryptTOSAgree:            new(bool),
		ManageVolumeClasses:            new(bool),
		NetworkPolicies:                new(bool),
		PodSecurityEnforceProfile:      "baseline",
		Profile:                        new(string),
		PublishBuilders:                new(bool),
		RecordBuilds:                   new(bool),
		SetPodSecurityEnforceProfile:   z.P(true),
		UseCustomCABundle:              new(bool),
		WorkloadMemoryDefault:          new(int64),
		WorkloadMemoryMaximum:          new(int64),
		RegistryMemory:                 new(string),
		RegistryCPU:                    new(string),
		BuildkitdMemory:                new(string),
		BuildkitdCPU:                   new(string),
		BuildkitdServiceMemory:         new(string),
		BuildkitdServiceCPU:            new(string),
		ControllerMemory:               new(string),
		ControllerCPU:                  new(string),
		APIServerMemory:                new(string),
		APIServerCPU:                   new(string),
	}
}