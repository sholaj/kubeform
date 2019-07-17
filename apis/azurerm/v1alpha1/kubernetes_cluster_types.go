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

type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterSpec   `json:"spec,omitempty"`
	Status            KubernetesClusterStatus `json:"status,omitempty"`
}

type KubernetesClusterSpecAgentPoolProfile struct {
	// +optional
	Count int    `json:"count,omitempty" tf:"count,omitempty"`
	Name  string `json:"name" tf:"name"`
	// +optional
	OsType string `json:"osType,omitempty" tf:"os_type,omitempty"`
	// +optional
	Type   string `json:"type,omitempty" tf:"type,omitempty"`
	VmSize string `json:"vmSize" tf:"vm_size"`
	// +optional
	VnetSubnetID string `json:"vnetSubnetID,omitempty" tf:"vnet_subnet_id,omitempty"`
}

type KubernetesClusterSpecLinuxProfileSshKey struct {
	KeyData string `json:"keyData" tf:"key_data"`
}

type KubernetesClusterSpecLinuxProfile struct {
	AdminUsername string `json:"adminUsername" tf:"admin_username"`
	// +kubebuilder:validation:MaxItems=1
	SshKey []KubernetesClusterSpecLinuxProfileSshKey `json:"sshKey" tf:"ssh_key"`
}

type KubernetesClusterSpecServicePrincipal struct {
	ClientID     string `json:"clientID" tf:"client_id"`
	ClientSecret string `json:"clientSecret" tf:"client_secret"`
}

type KubernetesClusterSpec struct {
	AgentPoolProfile []KubernetesClusterSpecAgentPoolProfile `json:"agentPoolProfile" tf:"agent_pool_profile"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ApiServerAuthorizedIPRanges []string `json:"apiServerAuthorizedIPRanges,omitempty" tf:"api_server_authorized_ip_ranges,omitempty"`
	DnsPrefix                   string   `json:"dnsPrefix" tf:"dns_prefix"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LinuxProfile      []KubernetesClusterSpecLinuxProfile `json:"linuxProfile,omitempty" tf:"linux_profile,omitempty"`
	Location          string                              `json:"location" tf:"location"`
	Name              string                              `json:"name" tf:"name"`
	ResourceGroupName string                              `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	ServicePrincipal []KubernetesClusterSpecServicePrincipal `json:"servicePrincipal" tf:"service_principal"`
	ProviderRef      core.LocalObjectReference               `json:"providerRef" tf:"-"`
}

type KubernetesClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KubernetesClusterList is a list of KubernetesClusters
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesCluster CRD objects
	Items []KubernetesCluster `json:"items,omitempty"`
}
