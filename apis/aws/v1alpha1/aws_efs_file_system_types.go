package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsFileSystem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEfsFileSystemSpec   `json:"spec,omitempty"`
	Status            AwsEfsFileSystemStatus `json:"status,omitempty"`
}

type AwsEfsFileSystemSpec struct {
	CreationToken                string            `json:"creation_token"`
	ReferenceName                string            `json:"reference_name"`
	DnsName                      string            `json:"dns_name"`
	ProvisionedThroughputInMibps float64           `json:"provisioned_throughput_in_mibps"`
	ThroughputMode               string            `json:"throughput_mode"`
	Arn                          string            `json:"arn"`
	PerformanceMode              string            `json:"performance_mode"`
	Encrypted                    bool              `json:"encrypted"`
	KmsKeyId                     string            `json:"kms_key_id"`
	Tags                         map[string]string `json:"tags"`
}

type AwsEfsFileSystemStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEfsFileSystemList is a list of AwsEfsFileSystems
type AwsEfsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEfsFileSystem CRD objects
	Items []AwsEfsFileSystem `json:"items,omitempty"`
}