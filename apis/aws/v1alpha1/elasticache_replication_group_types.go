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

type ElasticacheReplicationGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheReplicationGroupSpec   `json:"spec,omitempty"`
	Status            ElasticacheReplicationGroupStatus `json:"status,omitempty"`
}

type ElasticacheReplicationGroupSpec struct {
	// +optional
	AtRestEncryptionEnabled bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled,omitempty"`
	// +optional
	AuthToken string `json:"authToken,omitempty" tf:"auth_token,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	// +optional
	AutomaticFailoverEnabled bool `json:"automaticFailoverEnabled,omitempty" tf:"automatic_failover_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	NotificationTopicArn string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
	// +optional
	Port                        int    `json:"port,omitempty" tf:"port,omitempty"`
	ReplicationGroupDescription string `json:"replicationGroupDescription" tf:"replication_group_description"`
	ReplicationGroupID          string `json:"replicationGroupID" tf:"replication_group_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SnapshotArns []string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`
	// +optional
	SnapshotName string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`
	// +optional
	SnapshotRetentionLimit int `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TransitEncryptionEnabled bool                      `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`
	ProviderRef              core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ElasticacheReplicationGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticacheReplicationGroupList is a list of ElasticacheReplicationGroups
type ElasticacheReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticacheReplicationGroup CRD objects
	Items []ElasticacheReplicationGroup `json:"items,omitempty"`
}
