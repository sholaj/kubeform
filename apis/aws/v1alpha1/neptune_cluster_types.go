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

type NeptuneCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneClusterSpec   `json:"spec,omitempty"`
	Status            NeptuneClusterStatus `json:"status,omitempty"`
}

type NeptuneClusterSpec struct {
	// +optional
	BackupRetentionPeriod int `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`
	// +optional
	IamDatabaseAuthenticationEnabled bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IamRoles []string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`
	// +optional
	NeptuneClusterParameterGroupName string `json:"neptuneClusterParameterGroupName,omitempty" tf:"neptune_cluster_parameter_group_name,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	ReplicationSourceIdentifier string `json:"replicationSourceIdentifier,omitempty" tf:"replication_source_identifier,omitempty"`
	// +optional
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`
	// +optional
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`
	// +optional
	StorageEncrypted bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type NeptuneClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NeptuneClusterList is a list of NeptuneClusters
type NeptuneClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NeptuneCluster CRD objects
	Items []NeptuneCluster `json:"items,omitempty"`
}
