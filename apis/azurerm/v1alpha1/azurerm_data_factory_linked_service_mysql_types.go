package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDataFactoryLinkedServiceMysql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataFactoryLinkedServiceMysqlSpec   `json:"spec,omitempty"`
	Status            AzurermDataFactoryLinkedServiceMysqlStatus `json:"status,omitempty"`
}

type AzurermDataFactoryLinkedServiceMysqlSpec struct {
	Parameters             map[string]string `json:"parameters"`
	Annotations            []string          `json:"annotations"`
	Name                   string            `json:"name"`
	DataFactoryName        string            `json:"data_factory_name"`
	ResourceGroupName      string            `json:"resource_group_name"`
	Description            string            `json:"description"`
	ConnectionString       string            `json:"connection_string"`
	IntegrationRuntimeName string            `json:"integration_runtime_name"`
	AdditionalProperties   map[string]string `json:"additional_properties"`
}

type AzurermDataFactoryLinkedServiceMysqlStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDataFactoryLinkedServiceMysqlList is a list of AzurermDataFactoryLinkedServiceMysqls
type AzurermDataFactoryLinkedServiceMysqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactoryLinkedServiceMysql CRD objects
	Items []AzurermDataFactoryLinkedServiceMysql `json:"items,omitempty"`
}