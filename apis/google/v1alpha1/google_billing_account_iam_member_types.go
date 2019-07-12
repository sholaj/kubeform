package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleBillingAccountIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBillingAccountIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleBillingAccountIamMemberStatus `json:"status,omitempty"`
}

type GoogleBillingAccountIamMemberSpec struct {
	Role             string `json:"role"`
	Member           string `json:"member"`
	Etag             string `json:"etag"`
	BillingAccountId string `json:"billing_account_id"`
}

type GoogleBillingAccountIamMemberStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleBillingAccountIamMemberList is a list of GoogleBillingAccountIamMembers
type GoogleBillingAccountIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBillingAccountIamMember CRD objects
	Items []GoogleBillingAccountIamMember `json:"items,omitempty"`
}