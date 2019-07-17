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

type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineSpec   `json:"spec,omitempty"`
	Status            VirtualMachineStatus `json:"status,omitempty"`
}

type VirtualMachineSpecBootDiagnostics struct {
	Enabled    bool   `json:"enabled" tf:"enabled"`
	StorageURI string `json:"storageURI" tf:"storage_uri"`
}

type VirtualMachineSpecOsProfile struct {
	// +optional
	AdminPassword string `json:"adminPassword,omitempty" tf:"admin_password,omitempty"`
	AdminUsername string `json:"adminUsername" tf:"admin_username"`
	ComputerName  string `json:"computerName" tf:"computer_name"`
}

type VirtualMachineSpecOsProfileLinuxConfigSshKeys struct {
	KeyData string `json:"keyData" tf:"key_data"`
	Path    string `json:"path" tf:"path"`
}

type VirtualMachineSpecOsProfileLinuxConfig struct {
	DisablePasswordAuthentication bool `json:"disablePasswordAuthentication" tf:"disable_password_authentication"`
	// +optional
	SshKeys []VirtualMachineSpecOsProfileLinuxConfigSshKeys `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
}

type VirtualMachineSpecOsProfileSecretsVaultCertificates struct {
	// +optional
	CertificateStore string `json:"certificateStore,omitempty" tf:"certificate_store,omitempty"`
	CertificateURL   string `json:"certificateURL" tf:"certificate_url"`
}

type VirtualMachineSpecOsProfileSecrets struct {
	SourceVaultID string `json:"sourceVaultID" tf:"source_vault_id"`
	// +optional
	VaultCertificates []VirtualMachineSpecOsProfileSecretsVaultCertificates `json:"vaultCertificates,omitempty" tf:"vault_certificates,omitempty"`
}

type VirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig struct {
	Component   string `json:"component" tf:"component"`
	Content     string `json:"content" tf:"content"`
	Pass        string `json:"pass" tf:"pass"`
	SettingName string `json:"settingName" tf:"setting_name"`
}

type VirtualMachineSpecOsProfileWindowsConfigWinrm struct {
	// +optional
	CertificateURL string `json:"certificateURL,omitempty" tf:"certificate_url,omitempty"`
	Protocol       string `json:"protocol" tf:"protocol"`
}

type VirtualMachineSpecOsProfileWindowsConfig struct {
	// +optional
	AdditionalUnattendConfig []VirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig `json:"additionalUnattendConfig,omitempty" tf:"additional_unattend_config,omitempty"`
	// +optional
	EnableAutomaticUpgrades bool `json:"enableAutomaticUpgrades,omitempty" tf:"enable_automatic_upgrades,omitempty"`
	// +optional
	ProvisionVmAgent bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent,omitempty"`
	// +optional
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
	// +optional
	Winrm []VirtualMachineSpecOsProfileWindowsConfigWinrm `json:"winrm,omitempty" tf:"winrm,omitempty"`
}

type VirtualMachineSpecPlan struct {
	Name      string `json:"name" tf:"name"`
	Product   string `json:"product" tf:"product"`
	Publisher string `json:"publisher" tf:"publisher"`
}

type VirtualMachineSpecStorageOsDisk struct {
	CreateOption string `json:"createOption" tf:"create_option"`
	// +optional
	ImageURI string `json:"imageURI,omitempty" tf:"image_uri,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	VhdURI string `json:"vhdURI,omitempty" tf:"vhd_uri,omitempty"`
	// +optional
	WriteAcceleratorEnabled bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type VirtualMachineSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BootDiagnostics []VirtualMachineSpecBootDiagnostics `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics,omitempty"`
	// +optional
	DeleteDataDisksOnTermination bool `json:"deleteDataDisksOnTermination,omitempty" tf:"delete_data_disks_on_termination,omitempty"`
	// +optional
	DeleteOsDiskOnTermination bool     `json:"deleteOsDiskOnTermination,omitempty" tf:"delete_os_disk_on_termination,omitempty"`
	Location                  string   `json:"location" tf:"location"`
	Name                      string   `json:"name" tf:"name"`
	NetworkInterfaceIDS       []string `json:"networkInterfaceIDS" tf:"network_interface_ids"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	OsProfile []VirtualMachineSpecOsProfile `json:"osProfile,omitempty" tf:"os_profile,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	OsProfileLinuxConfig []VirtualMachineSpecOsProfileLinuxConfig `json:"osProfileLinuxConfig,omitempty" tf:"os_profile_linux_config,omitempty"`
	// +optional
	OsProfileSecrets []VirtualMachineSpecOsProfileSecrets `json:"osProfileSecrets,omitempty" tf:"os_profile_secrets,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	OsProfileWindowsConfig []VirtualMachineSpecOsProfileWindowsConfig `json:"osProfileWindowsConfig,omitempty" tf:"os_profile_windows_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Plan []VirtualMachineSpecPlan `json:"plan,omitempty" tf:"plan,omitempty"`
	// +optional
	PrimaryNetworkInterfaceID string `json:"primaryNetworkInterfaceID,omitempty" tf:"primary_network_interface_id,omitempty"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	StorageOsDisk []VirtualMachineSpecStorageOsDisk `json:"storageOsDisk" tf:"storage_os_disk"`
	VmSize        string                            `json:"vmSize" tf:"vm_size"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones       []string                  `json:"zones,omitempty" tf:"zones,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VirtualMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualMachineList is a list of VirtualMachines
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualMachine CRD objects
	Items []VirtualMachine `json:"items,omitempty"`
}
