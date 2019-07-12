package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesReceiptFilterSpec   `json:"spec,omitempty"`
	Status            AwsSesReceiptFilterStatus `json:"status,omitempty"`
}

type AwsSesReceiptFilterSpec struct {
	Name   string `json:"name"`
	Cidr   string `json:"cidr"`
	Policy string `json:"policy"`
}

type AwsSesReceiptFilterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesReceiptFilterList is a list of AwsSesReceiptFilters
type AwsSesReceiptFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesReceiptFilter CRD objects
	Items []AwsSesReceiptFilter `json:"items,omitempty"`
}