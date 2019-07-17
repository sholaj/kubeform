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

type KeyVaultSecret struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyVaultSecretSpec   `json:"spec,omitempty"`
	Status            KeyVaultSecretStatus `json:"status,omitempty"`
}

type KeyVaultSecretSpec struct {
	// +optional
	ContentType string                    `json:"contentType,omitempty" tf:"content_type,omitempty"`
	Name        string                    `json:"name" tf:"name"`
	Value       string                    `json:"value" tf:"value"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type KeyVaultSecretStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KeyVaultSecretList is a list of KeyVaultSecrets
type KeyVaultSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KeyVaultSecret CRD objects
	Items []KeyVaultSecret `json:"items,omitempty"`
}
