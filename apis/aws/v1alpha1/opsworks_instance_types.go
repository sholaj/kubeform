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

type OpsworksInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksInstanceSpec   `json:"spec,omitempty"`
	Status            OpsworksInstanceStatus `json:"status,omitempty"`
}

type OpsworksInstanceSpec struct {
	// +optional
	AgentVersion string `json:"agentVersion,omitempty" tf:"agent_version,omitempty"`
	// +optional
	Architecture string `json:"architecture,omitempty" tf:"architecture,omitempty"`
	// +optional
	AutoScalingType string `json:"autoScalingType,omitempty" tf:"auto_scaling_type,omitempty"`
	// +optional
	DeleteEbs bool `json:"deleteEbs,omitempty" tf:"delete_ebs,omitempty"`
	// +optional
	DeleteEip bool `json:"deleteEip,omitempty" tf:"delete_eip,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	InstallUpdatesOnBoot bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`
	// +optional
	InstanceType string   `json:"instanceType,omitempty" tf:"instance_type,omitempty"`
	LayerIDS     []string `json:"layerIDS" tf:"layer_ids"`
	StackID      string   `json:"stackID" tf:"stack_id"`
	// +optional
	State       string                    `json:"state,omitempty" tf:"state,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type OpsworksInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksInstanceList is a list of OpsworksInstances
type OpsworksInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksInstance CRD objects
	Items []OpsworksInstance `json:"items,omitempty"`
}
