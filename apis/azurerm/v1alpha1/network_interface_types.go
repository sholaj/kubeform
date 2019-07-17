package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceStatus `json:"status,omitempty"`
}

type NetworkInterfaceSpecIpConfiguration struct {
	Name string `json:"name" tf:"name"`
	// +optional
	PrivateIPAddress           string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	PrivateIPAddressAllocation string `json:"privateIPAddressAllocation" tf:"private_ip_address_allocation"`
	// +optional
	PrivateIPAddressVersion string `json:"privateIPAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`
	// +optional
	PublicIPAddressID string `json:"publicIPAddressID,omitempty" tf:"public_ip_address_id,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
}

type NetworkInterfaceSpec struct {
	// +optional
	EnableAcceleratedNetworking bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking,omitempty"`
	// +optional
	EnableIPForwarding bool                                  `json:"enableIPForwarding,omitempty" tf:"enable_ip_forwarding,omitempty"`
	IpConfiguration    []NetworkInterfaceSpecIpConfiguration `json:"ipConfiguration" tf:"ip_configuration"`
	Location           string                                `json:"location" tf:"location"`
	Name               string                                `json:"name" tf:"name"`
	// +optional
	NetworkSecurityGroupID string                    `json:"networkSecurityGroupID,omitempty" tf:"network_security_group_id,omitempty"`
	ResourceGroupName      string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ProviderRef            core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type NetworkInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkInterfaceList is a list of NetworkInterfaces
type NetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkInterface CRD objects
	Items []NetworkInterface `json:"items,omitempty"`
}
