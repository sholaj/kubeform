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

type SqlDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseSpec   `json:"spec,omitempty"`
	Status            SqlDatabaseStatus `json:"status,omitempty"`
}

type SqlDatabaseSpecImport struct {
	AdministratorLogin         string `json:"administratorLogin" tf:"administrator_login"`
	AdministratorLoginPassword string `json:"administratorLoginPassword" tf:"administrator_login_password"`
	AuthenticationType         string `json:"authenticationType" tf:"authentication_type"`
	// +optional
	OperationMode  string `json:"operationMode,omitempty" tf:"operation_mode,omitempty"`
	StorageKey     string `json:"storageKey" tf:"storage_key"`
	StorageKeyType string `json:"storageKeyType" tf:"storage_key_type"`
	StorageURI     string `json:"storageURI" tf:"storage_uri"`
}

type SqlDatabaseSpec struct {
	// +optional
	CreateMode string `json:"createMode,omitempty" tf:"create_mode,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Import   []SqlDatabaseSpecImport `json:"import,omitempty" tf:"import,omitempty"`
	Location string                  `json:"location" tf:"location"`
	Name     string                  `json:"name" tf:"name"`
	// +optional
	ReadScale         bool                      `json:"readScale,omitempty" tf:"read_scale,omitempty"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ServerName        string                    `json:"serverName" tf:"server_name"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SqlDatabaseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlDatabaseList is a list of SqlDatabases
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlDatabase CRD objects
	Items []SqlDatabase `json:"items,omitempty"`
}
