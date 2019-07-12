package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleBillingAccountIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBillingAccountIamPolicySpec   `json:"spec,omitempty"`
	Status            GoogleBillingAccountIamPolicyStatus `json:"status,omitempty"`
}

type GoogleBillingAccountIamPolicySpec struct {
	PolicyData       string `json:"policy_data"`
	Etag             string `json:"etag"`
	BillingAccountId string `json:"billing_account_id"`
}

type GoogleBillingAccountIamPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleBillingAccountIamPolicyList is a list of GoogleBillingAccountIamPolicys
type GoogleBillingAccountIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBillingAccountIamPolicy CRD objects
	Items []GoogleBillingAccountIamPolicy `json:"items,omitempty"`
}