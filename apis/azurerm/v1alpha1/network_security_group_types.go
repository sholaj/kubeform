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

type NetworkSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityGroupSpec   `json:"spec,omitempty"`
	Status            NetworkSecurityGroupStatus `json:"status,omitempty"`
}

type NetworkSecurityGroupSpec struct {
	Location          string                    `json:"location" tf:"location"`
	Name              string                    `json:"name" tf:"name"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type NetworkSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkSecurityGroupList is a list of NetworkSecurityGroups
type NetworkSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkSecurityGroup CRD objects
	Items []NetworkSecurityGroup `json:"items,omitempty"`
}
