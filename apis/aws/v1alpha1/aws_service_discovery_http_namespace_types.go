package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryHttpNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsServiceDiscoveryHttpNamespaceSpec   `json:"spec,omitempty"`
	Status            AwsServiceDiscoveryHttpNamespaceStatus `json:"status,omitempty"`
}

type AwsServiceDiscoveryHttpNamespaceSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Arn         string `json:"arn"`
}

type AwsServiceDiscoveryHttpNamespaceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryHttpNamespaceList is a list of AwsServiceDiscoveryHttpNamespaces
type AwsServiceDiscoveryHttpNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsServiceDiscoveryHttpNamespace CRD objects
	Items []AwsServiceDiscoveryHttpNamespace `json:"items,omitempty"`
}