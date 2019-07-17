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

type AppService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceSpec   `json:"spec,omitempty"`
	Status            AppServiceStatus `json:"status,omitempty"`
}

type AppServiceSpec struct {
	AppServicePlanID string `json:"appServicePlanID" tf:"app_service_plan_id"`
	// +optional
	ClientCertEnabled bool `json:"clientCertEnabled,omitempty" tf:"client_cert_enabled,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	HttpsOnly         bool                      `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`
	Location          string                    `json:"location" tf:"location"`
	Name              string                    `json:"name" tf:"name"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AppServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServiceList is a list of AppServices
type AppServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppService CRD objects
	Items []AppService `json:"items,omitempty"`
}
