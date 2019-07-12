package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermStorageQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStorageQueueSpec   `json:"spec,omitempty"`
	Status            AzurermStorageQueueStatus `json:"status,omitempty"`
}

type AzurermStorageQueueSpec struct {
	StorageAccountName string `json:"storage_account_name"`
	Name               string `json:"name"`
	ResourceGroupName  string `json:"resource_group_name"`
}

type AzurermStorageQueueStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermStorageQueueList is a list of AzurermStorageQueues
type AzurermStorageQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStorageQueue CRD objects
	Items []AzurermStorageQueue `json:"items,omitempty"`
}