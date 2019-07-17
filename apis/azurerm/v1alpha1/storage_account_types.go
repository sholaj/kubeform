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

type StorageAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountSpec   `json:"spec,omitempty"`
	Status            StorageAccountStatus `json:"status,omitempty"`
}

type StorageAccountSpecCustomDomain struct {
	Name string `json:"name" tf:"name"`
	// +optional
	UseSubdomain bool `json:"useSubdomain,omitempty" tf:"use_subdomain,omitempty"`
}

type StorageAccountSpecNetworkRules struct {
	// +optional
	DefaultAction string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`
}

type StorageAccountSpec struct {
	// +optional
	AccountEncryptionSource string `json:"accountEncryptionSource,omitempty" tf:"account_encryption_source,omitempty"`
	// +optional
	AccountKind            string `json:"accountKind,omitempty" tf:"account_kind,omitempty"`
	AccountReplicationType string `json:"accountReplicationType" tf:"account_replication_type"`
	AccountTier            string `json:"accountTier" tf:"account_tier"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomDomain []StorageAccountSpecCustomDomain `json:"customDomain,omitempty" tf:"custom_domain,omitempty"`
	// +optional
	EnableBlobEncryption bool `json:"enableBlobEncryption,omitempty" tf:"enable_blob_encryption,omitempty"`
	// +optional
	EnableFileEncryption bool `json:"enableFileEncryption,omitempty" tf:"enable_file_encryption,omitempty"`
	// +optional
	EnableHTTPSTrafficOnly bool `json:"enableHTTPSTrafficOnly,omitempty" tf:"enable_https_traffic_only,omitempty"`
	// +optional
	IsHnsEnabled bool   `json:"isHnsEnabled,omitempty" tf:"is_hns_enabled,omitempty"`
	Location     string `json:"location" tf:"location"`
	Name         string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkRules      []StorageAccountSpecNetworkRules `json:"networkRules,omitempty" tf:"network_rules,omitempty"`
	ResourceGroupName string                           `json:"resourceGroupName" tf:"resource_group_name"`
	ProviderRef       core.LocalObjectReference        `json:"providerRef" tf:"-"`
}

type StorageAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageAccountList is a list of StorageAccounts
type StorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageAccount CRD objects
	Items []StorageAccount `json:"items,omitempty"`
}
