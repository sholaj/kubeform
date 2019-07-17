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

type Loadbalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadbalancerSpec   `json:"spec,omitempty"`
	Status            LoadbalancerStatus `json:"status,omitempty"`
}

type LoadbalancerSpecForwardingRule struct {
	// +optional
	CertificateID  string `json:"certificateID,omitempty" tf:"certificate_id,omitempty"`
	EntryPort      int    `json:"entryPort" tf:"entry_port"`
	EntryProtocol  string `json:"entryProtocol" tf:"entry_protocol"`
	TargetPort     int    `json:"targetPort" tf:"target_port"`
	TargetProtocol string `json:"targetProtocol" tf:"target_protocol"`
	// +optional
	TlsPassthrough bool `json:"tlsPassthrough,omitempty" tf:"tls_passthrough,omitempty"`
}

type LoadbalancerSpec struct {
	// +optional
	Algorithm string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`
	// +optional
	DropletTag string `json:"dropletTag,omitempty" tf:"droplet_tag,omitempty"`
	// +optional
	EnableProxyProtocol bool `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol,omitempty"`
	// +kubebuilder:validation:MinItems=1
	ForwardingRule []LoadbalancerSpecForwardingRule `json:"forwardingRule" tf:"forwarding_rule"`
	Name           string                           `json:"name" tf:"name"`
	// +optional
	RedirectHTTPToHTTPS bool                      `json:"redirectHTTPToHTTPS,omitempty" tf:"redirect_http_to_https,omitempty"`
	Region              string                    `json:"region" tf:"region"`
	ProviderRef         core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LoadbalancerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoadbalancerList is a list of Loadbalancers
type LoadbalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Loadbalancer CRD objects
	Items []Loadbalancer `json:"items,omitempty"`
}
