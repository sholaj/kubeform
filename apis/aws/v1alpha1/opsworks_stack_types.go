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

type OpsworksStack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksStackSpec   `json:"spec,omitempty"`
	Status            OpsworksStackStatus `json:"status,omitempty"`
}

type OpsworksStackSpec struct {
	// +optional
	BerkshelfVersion string `json:"berkshelfVersion,omitempty" tf:"berkshelf_version,omitempty"`
	// +optional
	Color string `json:"color,omitempty" tf:"color,omitempty"`
	// +optional
	ConfigurationManagerName string `json:"configurationManagerName,omitempty" tf:"configuration_manager_name,omitempty"`
	// +optional
	ConfigurationManagerVersion string `json:"configurationManagerVersion,omitempty" tf:"configuration_manager_version,omitempty"`
	// +optional
	CustomJSON                string `json:"customJSON,omitempty" tf:"custom_json,omitempty"`
	DefaultInstanceProfileArn string `json:"defaultInstanceProfileArn" tf:"default_instance_profile_arn"`
	// +optional
	DefaultOs string `json:"defaultOs,omitempty" tf:"default_os,omitempty"`
	// +optional
	DefaultRootDeviceType string `json:"defaultRootDeviceType,omitempty" tf:"default_root_device_type,omitempty"`
	// +optional
	DefaultSSHKeyName string `json:"defaultSSHKeyName,omitempty" tf:"default_ssh_key_name,omitempty"`
	// +optional
	HostnameTheme string `json:"hostnameTheme,omitempty" tf:"hostname_theme,omitempty"`
	// +optional
	ManageBerkshelf bool   `json:"manageBerkshelf,omitempty" tf:"manage_berkshelf,omitempty"`
	Name            string `json:"name" tf:"name"`
	Region          string `json:"region" tf:"region"`
	ServiceRoleArn  string `json:"serviceRoleArn" tf:"service_role_arn"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UseCustomCookbooks bool `json:"useCustomCookbooks,omitempty" tf:"use_custom_cookbooks,omitempty"`
	// +optional
	UseOpsworksSecurityGroups bool                      `json:"useOpsworksSecurityGroups,omitempty" tf:"use_opsworks_security_groups,omitempty"`
	ProviderRef               core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type OpsworksStackStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksStackList is a list of OpsworksStacks
type OpsworksStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksStack CRD objects
	Items []OpsworksStack `json:"items,omitempty"`
}
