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

type Ami struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmiSpec   `json:"spec,omitempty"`
	Status            AmiStatus `json:"status,omitempty"`
}

type AmiSpec struct {
	// +optional
	Architecture string `json:"architecture,omitempty" tf:"architecture,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	EnaSupport bool `json:"enaSupport,omitempty" tf:"ena_support,omitempty"`
	// +optional
	KernelID string `json:"kernelID,omitempty" tf:"kernel_id,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	RamdiskID string `json:"ramdiskID,omitempty" tf:"ramdisk_id,omitempty"`
	// +optional
	RootDeviceName string `json:"rootDeviceName,omitempty" tf:"root_device_name,omitempty"`
	// +optional
	SriovNetSupport string `json:"sriovNetSupport,omitempty" tf:"sriov_net_support,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VirtualizationType string                    `json:"virtualizationType,omitempty" tf:"virtualization_type,omitempty"`
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AmiStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AmiList is a list of Amis
type AmiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ami CRD objects
	Items []Ami `json:"items,omitempty"`
}
