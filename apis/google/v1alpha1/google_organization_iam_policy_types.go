package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleOrganizationIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleOrganizationIamPolicySpec   `json:"spec,omitempty"`
	Status            GoogleOrganizationIamPolicyStatus `json:"status,omitempty"`
}

type GoogleOrganizationIamPolicySpec struct {
	PolicyData string `json:"policy_data"`
	Etag       string `json:"etag"`
	OrgId      string `json:"org_id"`
}

type GoogleOrganizationIamPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleOrganizationIamPolicyList is a list of GoogleOrganizationIamPolicys
type GoogleOrganizationIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleOrganizationIamPolicy CRD objects
	Items []GoogleOrganizationIamPolicy `json:"items,omitempty"`
}