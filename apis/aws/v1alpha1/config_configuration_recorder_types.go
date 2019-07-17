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

type ConfigConfigurationRecorder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigurationRecorderSpec   `json:"spec,omitempty"`
	Status            ConfigConfigurationRecorderStatus `json:"status,omitempty"`
}

type ConfigConfigurationRecorderSpec struct {
	// +optional
	Name        string                    `json:"name,omitempty" tf:"name,omitempty"`
	RoleArn     string                    `json:"roleArn" tf:"role_arn"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ConfigConfigurationRecorderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigConfigurationRecorderList is a list of ConfigConfigurationRecorders
type ConfigConfigurationRecorderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigConfigurationRecorder CRD objects
	Items []ConfigConfigurationRecorder `json:"items,omitempty"`
}
