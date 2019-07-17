package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MonitorMetricAlertrule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorMetricAlertruleSpec   `json:"spec,omitempty"`
	Status            MonitorMetricAlertruleStatus `json:"status,omitempty"`
}

type MonitorMetricAlertruleSpec struct {
	Aggregation string `json:"aggregation" tf:"aggregation"`
	// +optional
	Enabled           bool                      `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Location          string                    `json:"location" tf:"location"`
	MetricName        string                    `json:"metricName" tf:"metric_name"`
	Name              string                    `json:"name" tf:"name"`
	Operator          string                    `json:"operator" tf:"operator"`
	Period            string                    `json:"period" tf:"period"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ResourceID        string                    `json:"resourceID" tf:"resource_id"`
	Threshold         json.Number               `json:"threshold" tf:"threshold"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type MonitorMetricAlertruleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorMetricAlertruleList is a list of MonitorMetricAlertrules
type MonitorMetricAlertruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorMetricAlertrule CRD objects
	Items []MonitorMetricAlertrule `json:"items,omitempty"`
}
