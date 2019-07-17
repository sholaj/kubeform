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

type ComputeForwardingRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeForwardingRuleSpec   `json:"spec,omitempty"`
	Status            ComputeForwardingRuleStatus `json:"status,omitempty"`
}

type ComputeForwardingRuleSpec struct {
	// +optional
	BackendService string `json:"backendService,omitempty" tf:"backend_service,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	IpVersion string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`
	// +optional
	// Deprecated
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	LoadBalancingScheme string `json:"loadBalancingScheme,omitempty" tf:"load_balancing_scheme,omitempty"`
	Name                string `json:"name" tf:"name"`
	// +optional
	PortRange string `json:"portRange,omitempty" tf:"port_range,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:UniqueItems=true
	Ports []string `json:"ports,omitempty" tf:"ports,omitempty"`
	// +optional
	// Deprecated
	ServiceLabel string `json:"serviceLabel,omitempty" tf:"service_label,omitempty"`
	// +optional
	Target      string                    `json:"target,omitempty" tf:"target,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeForwardingRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeForwardingRuleList is a list of ComputeForwardingRules
type ComputeForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeForwardingRule CRD objects
	Items []ComputeForwardingRule `json:"items,omitempty"`
}
