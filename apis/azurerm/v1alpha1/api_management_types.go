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

type ApiManagement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementSpec   `json:"spec,omitempty"`
	Status            ApiManagementStatus `json:"status,omitempty"`
}

type ApiManagementSpecAdditionalLocation struct {
	Location string `json:"location" tf:"location"`
}

type ApiManagementSpecCertificate struct {
	CertificatePassword string `json:"certificatePassword" tf:"certificate_password"`
	EncodedCertificate  string `json:"encodedCertificate" tf:"encoded_certificate"`
	StoreName           string `json:"storeName" tf:"store_name"`
}

type ApiManagementSpecIdentity struct {
	Type string `json:"type" tf:"type"`
}

type ApiManagementSpecSku struct {
	Capacity int    `json:"capacity" tf:"capacity"`
	Name     string `json:"name" tf:"name"`
}

type ApiManagementSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AdditionalLocation []ApiManagementSpecAdditionalLocation `json:"additionalLocation,omitempty" tf:"additional_location,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Certificate []ApiManagementSpecCertificate `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity          []ApiManagementSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Location          string                      `json:"location" tf:"location"`
	Name              string                      `json:"name" tf:"name"`
	PublisherEmail    string                      `json:"publisherEmail" tf:"publisher_email"`
	PublisherName     string                      `json:"publisherName" tf:"publisher_name"`
	ResourceGroupName string                      `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku         []ApiManagementSpecSku    `json:"sku" tf:"sku"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApiManagementStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementList is a list of ApiManagements
type ApiManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagement CRD objects
	Items []ApiManagement `json:"items,omitempty"`
}
