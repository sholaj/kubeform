package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserGroupMembership struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamUserGroupMembershipSpec   `json:"spec,omitempty"`
	Status            AwsIamUserGroupMembershipStatus `json:"status,omitempty"`
}

type AwsIamUserGroupMembershipSpec struct {
	User   string   `json:"user"`
	Groups []string `json:"groups"`
}

type AwsIamUserGroupMembershipStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserGroupMembershipList is a list of AwsIamUserGroupMemberships
type AwsIamUserGroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamUserGroupMembership CRD objects
	Items []AwsIamUserGroupMembership `json:"items,omitempty"`
}