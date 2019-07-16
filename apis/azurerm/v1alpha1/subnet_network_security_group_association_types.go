package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SubnetNetworkSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetNetworkSecurityGroupAssociationSpec   `json:"spec,omitempty"`
	Status            SubnetNetworkSecurityGroupAssociationStatus `json:"status,omitempty"`
}

type SubnetNetworkSecurityGroupAssociationSpec struct {
	NetworkSecurityGroupId string `json:"network_security_group_id"`
	SubnetId               string `json:"subnet_id"`
}

type SubnetNetworkSecurityGroupAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SubnetNetworkSecurityGroupAssociationList is a list of SubnetNetworkSecurityGroupAssociations
type SubnetNetworkSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SubnetNetworkSecurityGroupAssociation CRD objects
	Items []SubnetNetworkSecurityGroupAssociation `json:"items,omitempty"`
}