package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcrRepository struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEcrRepositorySpec   `json:"spec,omitempty"`
	Status            AwsEcrRepositoryStatus `json:"status,omitempty"`
}

type AwsEcrRepositorySpec struct {
	Arn           string            `json:"arn"`
	RegistryId    string            `json:"registry_id"`
	RepositoryUrl string            `json:"repository_url"`
	Name          string            `json:"name"`
	Tags          map[string]string `json:"tags"`
}

type AwsEcrRepositoryStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcrRepositoryList is a list of AwsEcrRepositorys
type AwsEcrRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEcrRepository CRD objects
	Items []AwsEcrRepository `json:"items,omitempty"`
}