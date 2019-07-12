package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermLb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLbSpec   `json:"spec,omitempty"`
	Status            AzurermLbStatus `json:"status,omitempty"`
}

type AzurermLbSpecFrontendIpConfiguration struct {
	Name                       string   `json:"name"`
	PrivateIpAddress           string   `json:"private_ip_address"`
	PublicIpAddressId          string   `json:"public_ip_address_id"`
	PrivateIpAddressAllocation string   `json:"private_ip_address_allocation"`
	LoadBalancerRules          []string `json:"load_balancer_rules"`
	InboundNatRules            []string `json:"inbound_nat_rules"`
	OutboundRules              []string `json:"outbound_rules"`
	Zones                      []string `json:"zones"`
	SubnetId                   string   `json:"subnet_id"`
	PublicIpPrefixId           string   `json:"public_ip_prefix_id"`
}

type AzurermLbSpec struct {
	FrontendIpConfiguration []AzurermLbSpec   `json:"frontend_ip_configuration"`
	PrivateIpAddress        string            `json:"private_ip_address"`
	PrivateIpAddresses      []string          `json:"private_ip_addresses"`
	Tags                    map[string]string `json:"tags"`
	Name                    string            `json:"name"`
	Location                string            `json:"location"`
	ResourceGroupName       string            `json:"resource_group_name"`
	Sku                     string            `json:"sku"`
}

type AzurermLbStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermLbList is a list of AzurermLbs
type AzurermLbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLb CRD objects
	Items []AzurermLb `json:"items,omitempty"`
}