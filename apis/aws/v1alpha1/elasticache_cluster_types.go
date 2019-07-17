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

type ElasticacheCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheClusterSpec   `json:"spec,omitempty"`
	Status            ElasticacheClusterStatus `json:"status,omitempty"`
}

type ElasticacheClusterSpec struct {
	ClusterID string `json:"clusterID" tf:"cluster_id"`
	// +optional
	NotificationTopicArn string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	PreferredAvailabilityZones []string `json:"preferredAvailabilityZones,omitempty" tf:"preferred_availability_zones,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SnapshotArns []string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`
	// +optional
	SnapshotName string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`
	// +optional
	SnapshotRetentionLimit int `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ElasticacheClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticacheClusterList is a list of ElasticacheClusters
type ElasticacheClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticacheCluster CRD objects
	Items []ElasticacheCluster `json:"items,omitempty"`
}
