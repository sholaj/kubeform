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

type ComputeInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeInstanceSpec   `json:"spec,omitempty"`
	Status            ComputeInstanceStatus `json:"status,omitempty"`
}

type ComputeInstanceSpecAttachedDisk struct {
	// +optional
	DiskEncryptionKeyRaw string `json:"diskEncryptionKeyRaw,omitempty" tf:"disk_encryption_key_raw,omitempty"`
	// +optional
	Mode   string `json:"mode,omitempty" tf:"mode,omitempty"`
	Source string `json:"source" tf:"source"`
}

type ComputeInstanceSpecBootDisk struct {
	// +optional
	AutoDelete bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`
	// +optional
	DiskEncryptionKeyRaw string `json:"diskEncryptionKeyRaw,omitempty" tf:"disk_encryption_key_raw,omitempty"`
}

type ComputeInstanceSpecNetworkInterfaceAccessConfig struct {
	// +optional
	PublicPtrDomainName string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty"`
}

type ComputeInstanceSpecNetworkInterfaceAliasIPRange struct {
	IpCIDRRange string `json:"ipCIDRRange" tf:"ip_cidr_range"`
	// +optional
	SubnetworkRangeName string `json:"subnetworkRangeName,omitempty" tf:"subnetwork_range_name,omitempty"`
}

type ComputeInstanceSpecNetworkInterface struct {
	// +optional
	AccessConfig []ComputeInstanceSpecNetworkInterfaceAccessConfig `json:"accessConfig,omitempty" tf:"access_config,omitempty"`
	// +optional
	AliasIPRange []ComputeInstanceSpecNetworkInterfaceAliasIPRange `json:"aliasIPRange,omitempty" tf:"alias_ip_range,omitempty"`
}

type ComputeInstanceSpecScratchDisk struct {
	// +optional
	Interface string `json:"interface,omitempty" tf:"interface,omitempty"`
}

type ComputeInstanceSpecServiceAccount struct {
	// +kubebuilder:validation:UniqueItems=true
	Scopes []string `json:"scopes" tf:"scopes"`
}

type ComputeInstanceSpec struct {
	// +optional
	AllowStoppingForUpdate bool `json:"allowStoppingForUpdate,omitempty" tf:"allow_stopping_for_update,omitempty"`
	// +optional
	AttachedDisk []ComputeInstanceSpecAttachedDisk `json:"attachedDisk,omitempty" tf:"attached_disk,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	BootDisk []ComputeInstanceSpecBootDisk `json:"bootDisk" tf:"boot_disk"`
	// +optional
	CanIPForward bool `json:"canIPForward,omitempty" tf:"can_ip_forward,omitempty"`
	// +optional
	// Deprecated
	CreateTimeout int `json:"createTimeout,omitempty" tf:"create_timeout,omitempty"`
	// +optional
	DeletionProtection bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Labels      map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	MachineType string            `json:"machineType" tf:"machine_type"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	// +optional
	MetadataStartupScript string `json:"metadataStartupScript,omitempty" tf:"metadata_startup_script,omitempty"`
	// +optional
	MinCPUPlatform   string                                `json:"minCPUPlatform,omitempty" tf:"min_cpu_platform,omitempty"`
	Name             string                                `json:"name" tf:"name"`
	NetworkInterface []ComputeInstanceSpecNetworkInterface `json:"networkInterface" tf:"network_interface"`
	// +optional
	ScratchDisk []ComputeInstanceSpecScratchDisk `json:"scratchDisk,omitempty" tf:"scratch_disk,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceAccount []ComputeInstanceSpecServiceAccount `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags        []string                  `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeInstanceList is a list of ComputeInstances
type ComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeInstance CRD objects
	Items []ComputeInstance `json:"items,omitempty"`
}
