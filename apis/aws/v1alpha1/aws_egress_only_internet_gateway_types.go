package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEgressOnlyInternetGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEgressOnlyInternetGatewaySpec   `json:"spec,omitempty"`
	Status            AwsEgressOnlyInternetGatewayStatus `json:"status,omitempty"`
}

type AwsEgressOnlyInternetGatewaySpec struct {
	VpcId string `json:"vpc_id"`
}

type AwsEgressOnlyInternetGatewayStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEgressOnlyInternetGatewayList is a list of AwsEgressOnlyInternetGateways
type AwsEgressOnlyInternetGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEgressOnlyInternetGateway CRD objects
	Items []AwsEgressOnlyInternetGateway `json:"items,omitempty"`
}