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

type VpcPeeringConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcPeeringConnectionSpec   `json:"spec,omitempty"`
	Status            VpcPeeringConnectionStatus `json:"status,omitempty"`
}

type VpcPeeringConnectionSpec struct {
	// +optional
	AutoAccept bool   `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`
	PeerVpcID  string `json:"peerVpcID" tf:"peer_vpc_id"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	VpcID       string                    `json:"vpcID" tf:"vpc_id"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VpcPeeringConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcPeeringConnectionList is a list of VpcPeeringConnections
type VpcPeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcPeeringConnection CRD objects
	Items []VpcPeeringConnection `json:"items,omitempty"`
}
