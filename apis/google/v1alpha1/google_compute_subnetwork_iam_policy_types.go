package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeSubnetworkIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSubnetworkIamPolicySpec   `json:"spec,omitempty"`
	Status            GoogleComputeSubnetworkIamPolicyStatus `json:"status,omitempty"`
}

type GoogleComputeSubnetworkIamPolicySpec struct {
	PolicyData string `json:"policy_data"`
	Etag       string `json:"etag"`
	Region     string `json:"region"`
	Subnetwork string `json:"subnetwork"`
	Project    string `json:"project"`
}

type GoogleComputeSubnetworkIamPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeSubnetworkIamPolicyList is a list of GoogleComputeSubnetworkIamPolicys
type GoogleComputeSubnetworkIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSubnetworkIamPolicy CRD objects
	Items []GoogleComputeSubnetworkIamPolicy `json:"items,omitempty"`
}