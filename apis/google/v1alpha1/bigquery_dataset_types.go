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

type BigqueryDataset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BigqueryDatasetSpec   `json:"spec,omitempty"`
	Status            BigqueryDatasetStatus `json:"status,omitempty"`
}

type BigqueryDatasetSpec struct {
	DatasetID string `json:"datasetID" tf:"dataset_id"`
	// +optional
	DefaultTableExpirationMs int `json:"defaultTableExpirationMs,omitempty" tf:"default_table_expiration_ms,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	FriendlyName string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	Location    string                    `json:"location,omitempty" tf:"location,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BigqueryDatasetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BigqueryDatasetList is a list of BigqueryDatasets
type BigqueryDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BigqueryDataset CRD objects
	Items []BigqueryDataset `json:"items,omitempty"`
}
