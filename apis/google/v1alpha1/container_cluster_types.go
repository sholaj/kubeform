package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ContainerCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerClusterSpec   `json:"spec,omitempty"`
	Status            ContainerClusterStatus `json:"status,omitempty"`
}

type ContainerClusterSpecIpAllocationPolicy struct {
	// +optional
	CreateSubnetwork bool `json:"createSubnetwork,omitempty" tf:"create_subnetwork,omitempty"`
	// +optional
	SubnetworkName string `json:"subnetworkName,omitempty" tf:"subnetwork_name,omitempty"`
}

type ContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow struct {
	StartTime string `json:"startTime" tf:"start_time"`
}

type ContainerClusterSpecMaintenancePolicy struct {
	// +kubebuilder:validation:MaxItems=1
	DailyMaintenanceWindow []ContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow `json:"dailyMaintenanceWindow" tf:"daily_maintenance_window"`
}

type ContainerClusterSpecMasterAuthorizedNetworksConfig struct{}

type ContainerClusterSpecPodSecurityPolicyConfig struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type ContainerClusterSpec struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// Deprecated
	EnableBinaryAuthorization bool `json:"enableBinaryAuthorization,omitempty" tf:"enable_binary_authorization,omitempty"`
	// +optional
	EnableKubernetesAlpha bool `json:"enableKubernetesAlpha,omitempty" tf:"enable_kubernetes_alpha,omitempty"`
	// +optional
	EnableLegacyAbac bool `json:"enableLegacyAbac,omitempty" tf:"enable_legacy_abac,omitempty"`
	// +optional
	// Deprecated
	EnableTpu bool `json:"enableTpu,omitempty" tf:"enable_tpu,omitempty"`
	// +optional
	InitialNodeCount int `json:"initialNodeCount,omitempty" tf:"initial_node_count,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IpAllocationPolicy []ContainerClusterSpecIpAllocationPolicy `json:"ipAllocationPolicy,omitempty" tf:"ip_allocation_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MaintenancePolicy []ContainerClusterSpecMaintenancePolicy `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MasterAuthorizedNetworksConfig []ContainerClusterSpecMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty" tf:"master_authorized_networks_config,omitempty"`
	// +optional
	MinMasterVersion string `json:"minMasterVersion,omitempty" tf:"min_master_version,omitempty"`
	Name             string `json:"name" tf:"name"`
	// +optional
	Network string `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	PodSecurityPolicyConfig []ContainerClusterSpecPodSecurityPolicyConfig `json:"podSecurityPolicyConfig,omitempty" tf:"pod_security_policy_config,omitempty"`
	// +optional
	RemoveDefaultNodePool bool `json:"removeDefaultNodePool,omitempty" tf:"remove_default_node_pool,omitempty"`
	// +optional
	ResourceLabels map[string]string         `json:"resourceLabels,omitempty" tf:"resource_labels,omitempty"`
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ContainerClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerClusterList is a list of ContainerClusters
type ContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerCluster CRD objects
	Items []ContainerCluster `json:"items,omitempty"`
}
