package config

import (
	"path"
	"strings"

	"github.com/layer5io/meshery-adapter-library/common"
	"github.com/layer5io/meshery-adapter-library/config"
	"github.com/layer5io/meshery-adapter-library/status"
	configprovider "github.com/layer5io/meshkit/config/provider"
	"github.com/layer5io/meshkit/utils"
	smp "github.com/layer5io/service-mesh-performance/spec"
)

const (
	// Constants to use in log statements
	LabelNamespace = "label-namespace"

	ServicePatchFile = "service-patch-file"
	CPPatchFile      = "cp-patch-file"
	ControlPatchFile = "control-patch-file"
	FilterPatchFile  = "filter-patch-file"

	// Istio vet operation
	IstioVetOperation = "istio-vet"

	// Configure Envoy filter operation
	EnvoyFilterOperation = "envoy-filter-operation"

	// Addons that the adapter supports
	PrometheusAddon = "prometheus-addon"
	GrafanaAddon    = "grafana-addon"
	KialiAddon      = "kiali-addon"
	JaegerAddon     = "jaeger-addon"
	ZipkinAddon     = "zipkin-addon"

	// Policies
	DenyAllPolicyOperation     = "deny-all-policy-operation"
	StrictMTLSPolicyOperation  = "strict-mtls-policy-operation"
	MutualMTLSPolicyOperation  = "mutual-mtls-policy-operation"
	DisableMTLSPolicyOperation = "disable-mtls-policy-operation"

	// OAM Metadata constants
	OAMAdapterNameMetadataKey       = "adapter.meshery.io/name"
	OAMComponentCategoryMetadataKey = "ui.meshery.io/category"
)

var (
	// IstioMeshOperation is the default name for the install
	// and uninstall commands on the istio mesh
	IstioOperation = strings.ToLower(smp.ServiceMesh_ISTIO.Enum().String())

	ServerVersion  = status.None
	ServerGitSHA   = status.None
	configRootPath = path.Join(utils.GetHome(), ".meshery")

	Config = configprovider.Options{
		FilePath: configRootPath,
		FileName: "istio",
		FileType: "yaml",
	}

	ServerConfig = map[string]string{
		"name":     smp.ServiceMesh_ISTIO.Enum().String(),
		"type":     "adapter",
		"port":     "10000",
		"traceurl": status.None,
	}

	MeshSpec = map[string]string{
		"name":    smp.ServiceMesh_ISTIO.Enum().String(),
		"status":  status.NotInstalled,
		"version": status.None,
	}

	ProviderConfig = map[string]string{
		configprovider.FilePath: configRootPath,
		configprovider.FileType: "yaml",
		configprovider.FileName: "istio",
	}

	// KubeConfig - Controlling the kubeconfig lifecycle with viper
	KubeConfig = map[string]string{
		configprovider.FilePath: configRootPath,
		configprovider.FileType: "yaml",
		configprovider.FileName: "kubeconfig",
	}

	Operations = getOperations(common.Operations)
)

// New creates a new config instance
func New(provider string) (config.Handler, error) {
	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(Config)
	case configprovider.InMemKey:
		return configprovider.NewInMem(Config)
	}

	return nil, ErrEmptyConfig
}

func NewKubeconfigBuilder(provider string) (config.Handler, error) {
	opts := configprovider.Options{
		FilePath: configRootPath,
		FileType: "yaml",
		FileName: "kubeconfig",
	}

	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(opts)
	case configprovider.InMemKey:
		return configprovider.NewInMem(opts)
	}
	return nil, ErrEmptyConfig
}

// RootPath returns the config root path for the adapter
func RootPath() string {
	return configRootPath
}
