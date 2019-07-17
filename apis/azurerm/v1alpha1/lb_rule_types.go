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

type LbRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbRuleSpec   `json:"spec,omitempty"`
	Status            LbRuleStatus `json:"status,omitempty"`
}

type LbRuleSpec struct {
	BackendPort int `json:"backendPort" tf:"backend_port"`
	// +optional
	DisableOutboundSnat bool `json:"disableOutboundSnat,omitempty" tf:"disable_outbound_snat,omitempty"`
	// +optional
	EnableFloatingIP            bool   `json:"enableFloatingIP,omitempty" tf:"enable_floating_ip,omitempty"`
	FrontendIPConfigurationName string `json:"frontendIPConfigurationName" tf:"frontend_ip_configuration_name"`
	FrontendPort                int    `json:"frontendPort" tf:"frontend_port"`
	LoadbalancerID              string `json:"loadbalancerID" tf:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location          string                    `json:"location,omitempty" tf:"location,omitempty"`
	Name              string                    `json:"name" tf:"name"`
	Protocol          string                    `json:"protocol" tf:"protocol"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LbRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbRuleList is a list of LbRules
type LbRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbRule CRD objects
	Items []LbRule `json:"items,omitempty"`
}
