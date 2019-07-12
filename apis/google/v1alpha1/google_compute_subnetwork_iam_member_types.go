package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeSubnetworkIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSubnetworkIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleComputeSubnetworkIamMemberStatus `json:"status,omitempty"`
}

type GoogleComputeSubnetworkIamMemberSpec struct {
	Role       string `json:"role"`
	Member     string `json:"member"`
	Etag       string `json:"etag"`
	Project    string `json:"project"`
	Region     string `json:"region"`
	Subnetwork string `json:"subnetwork"`
}

type GoogleComputeSubnetworkIamMemberStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeSubnetworkIamMemberList is a list of GoogleComputeSubnetworkIamMembers
type GoogleComputeSubnetworkIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSubnetworkIamMember CRD objects
	Items []GoogleComputeSubnetworkIamMember `json:"items,omitempty"`
}