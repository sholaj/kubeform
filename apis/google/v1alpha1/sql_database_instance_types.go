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

type SqlDatabaseInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseInstanceSpec   `json:"spec,omitempty"`
	Status            SqlDatabaseInstanceStatus `json:"status,omitempty"`
}

type SqlDatabaseInstanceSpecSettingsDatabaseFlags struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type SqlDatabaseInstanceSpecSettingsMaintenanceWindow struct {
	// +optional
	Day int `json:"day,omitempty" tf:"day,omitempty"`
	// +optional
	Hour int `json:"hour,omitempty" tf:"hour,omitempty"`
	// +optional
	UpdateTrack string `json:"updateTrack,omitempty" tf:"update_track,omitempty"`
}

type SqlDatabaseInstanceSpecSettings struct {
	// +optional
	AuthorizedGaeApplications []string `json:"authorizedGaeApplications,omitempty" tf:"authorized_gae_applications,omitempty"`
	// +optional
	DatabaseFlags []SqlDatabaseInstanceSpecSettingsDatabaseFlags `json:"databaseFlags,omitempty" tf:"database_flags,omitempty"`
	// +optional
	DiskAutoresize bool `json:"diskAutoresize,omitempty" tf:"disk_autoresize,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MaintenanceWindow []SqlDatabaseInstanceSpecSettingsMaintenanceWindow `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
	// +optional
	PricingPlan string `json:"pricingPlan,omitempty" tf:"pricing_plan,omitempty"`
	// +optional
	ReplicationType string `json:"replicationType,omitempty" tf:"replication_type,omitempty"`
	Tier            string `json:"tier" tf:"tier"`
	// +optional
	UserLabels map[string]string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

type SqlDatabaseInstanceSpec struct {
	// +optional
	DatabaseVersion string `json:"databaseVersion,omitempty" tf:"database_version,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Settings    []SqlDatabaseInstanceSpecSettings `json:"settings" tf:"settings"`
	ProviderRef core.LocalObjectReference         `json:"providerRef" tf:"-"`
}

type SqlDatabaseInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlDatabaseInstanceList is a list of SqlDatabaseInstances
type SqlDatabaseInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlDatabaseInstance CRD objects
	Items []SqlDatabaseInstance `json:"items,omitempty"`
}
