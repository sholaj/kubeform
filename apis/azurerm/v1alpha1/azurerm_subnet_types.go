package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermSubnet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSubnetSpec   `json:"spec,omitempty"`
	Status            AzurermSubnetStatus `json:"status,omitempty"`
}

type AzurermSubnetSpecDelegationServiceDelegation struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

type AzurermSubnetSpecDelegation struct {
	Name              string                        `json:"name"`
	ServiceDelegation []AzurermSubnetSpecDelegation `json:"service_delegation"`
}

type AzurermSubnetSpec struct {
	Name                   string              `json:"name"`
	VirtualNetworkName     string              `json:"virtual_network_name"`
	IpConfigurations       []string            `json:"ip_configurations"`
	ServiceEndpoints       []string            `json:"service_endpoints"`
	ResourceGroupName      string              `json:"resource_group_name"`
	AddressPrefix          string              `json:"address_prefix"`
	NetworkSecurityGroupId string              `json:"network_security_group_id"`
	RouteTableId           string              `json:"route_table_id"`
	Delegation             []AzurermSubnetSpec `json:"delegation"`
}

type AzurermSubnetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermSubnetList is a list of AzurermSubnets
type AzurermSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSubnet CRD objects
	Items []AzurermSubnet `json:"items,omitempty"`
}