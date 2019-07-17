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

type FunctionApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionAppSpec   `json:"spec,omitempty"`
	Status            FunctionAppStatus `json:"status,omitempty"`
}

type FunctionAppSpec struct {
	AppServicePlanID string `json:"appServicePlanID" tf:"app_service_plan_id"`
	// +optional
	AppSettings map[string]string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`
	// +optional
	EnableBuiltinLogging bool `json:"enableBuiltinLogging,omitempty" tf:"enable_builtin_logging,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	HttpsOnly               bool   `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`
	Location                string `json:"location" tf:"location"`
	Name                    string `json:"name" tf:"name"`
	ResourceGroupName       string `json:"resourceGroupName" tf:"resource_group_name"`
	StorageConnectionString string `json:"storageConnectionString" tf:"storage_connection_string"`
	// +optional
	Version     string                    `json:"version,omitempty" tf:"version,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type FunctionAppStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FunctionAppList is a list of FunctionApps
type FunctionAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FunctionApp CRD objects
	Items []FunctionApp `json:"items,omitempty"`
}
