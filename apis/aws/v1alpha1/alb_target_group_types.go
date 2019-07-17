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

type AlbTargetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbTargetGroupSpec   `json:"spec,omitempty"`
	Status            AlbTargetGroupStatus `json:"status,omitempty"`
}

type AlbTargetGroupSpec struct {
	// +optional
	DeregistrationDelay int `json:"deregistrationDelay,omitempty" tf:"deregistration_delay,omitempty"`
	// +optional
	LambdaMultiValueHeadersEnabled bool `json:"lambdaMultiValueHeadersEnabled,omitempty" tf:"lambda_multi_value_headers_enabled,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	ProxyProtocolV2 bool `json:"proxyProtocolV2,omitempty" tf:"proxy_protocol_v2,omitempty"`
	// +optional
	SlowStart int `json:"slowStart,omitempty" tf:"slow_start,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TargetType string `json:"targetType,omitempty" tf:"target_type,omitempty"`
	// +optional
	VpcID       string                    `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AlbTargetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AlbTargetGroupList is a list of AlbTargetGroups
type AlbTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AlbTargetGroup CRD objects
	Items []AlbTargetGroup `json:"items,omitempty"`
}
