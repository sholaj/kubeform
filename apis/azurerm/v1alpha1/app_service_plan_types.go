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

type AppServicePlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServicePlanSpec   `json:"spec,omitempty"`
	Status            AppServicePlanStatus `json:"status,omitempty"`
}

type AppServicePlanSpecSku struct {
	Size string `json:"size" tf:"size"`
	Tier string `json:"tier" tf:"tier"`
}

type AppServicePlanSpec struct {
	// +optional
	IsXenon bool `json:"isXenon,omitempty" tf:"is_xenon,omitempty"`
	// +optional
	Kind              string `json:"kind,omitempty" tf:"kind,omitempty"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku         []AppServicePlanSpecSku   `json:"sku" tf:"sku"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AppServicePlanStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServicePlanList is a list of AppServicePlans
type AppServicePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppServicePlan CRD objects
	Items []AppServicePlan `json:"items,omitempty"`
}
