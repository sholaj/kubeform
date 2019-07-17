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

type CloudfunctionsFunction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfunctionsFunctionSpec   `json:"spec,omitempty"`
	Status            CloudfunctionsFunctionStatus `json:"status,omitempty"`
}

type CloudfunctionsFunctionSpec struct {
	// +optional
	AvailableMemoryMb int `json:"availableMemoryMb,omitempty" tf:"available_memory_mb,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	EntryPoint string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`
	// +optional
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`
	// +optional
	Labels              map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name                string            `json:"name" tf:"name"`
	SourceArchiveBucket string            `json:"sourceArchiveBucket" tf:"source_archive_bucket"`
	SourceArchiveObject string            `json:"sourceArchiveObject" tf:"source_archive_object"`
	// +optional
	Timeout int `json:"timeout,omitempty" tf:"timeout,omitempty"`
	// +optional
	TriggerHTTP bool                      `json:"triggerHTTP,omitempty" tf:"trigger_http,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CloudfunctionsFunctionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudfunctionsFunctionList is a list of CloudfunctionsFunctions
type CloudfunctionsFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudfunctionsFunction CRD objects
	Items []CloudfunctionsFunction `json:"items,omitempty"`
}
