package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermSearchService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSearchServiceSpec   `json:"spec,omitempty"`
	Status            AzurermSearchServiceStatus `json:"status,omitempty"`
}

type AzurermSearchServiceSpec struct {
	Tags              map[string]string `json:"tags"`
	Location          string            `json:"location"`
	Sku               string            `json:"sku"`
	PartitionCount    int               `json:"partition_count"`
	PrimaryKey        string            `json:"primary_key"`
	SecondaryKey      string            `json:"secondary_key"`
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	ReplicaCount      int               `json:"replica_count"`
}

type AzurermSearchServiceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermSearchServiceList is a list of AzurermSearchServices
type AzurermSearchServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSearchService CRD objects
	Items []AzurermSearchService `json:"items,omitempty"`
}