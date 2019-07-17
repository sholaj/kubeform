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

type DbInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbInstanceSpec   `json:"spec,omitempty"`
	Status            DbInstanceStatus `json:"status,omitempty"`
}

type DbInstanceSpecS3Import struct {
	BucketName string `json:"bucketName" tf:"bucket_name"`
	// +optional
	BucketPrefix        string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`
	IngestionRole       string `json:"ingestionRole" tf:"ingestion_role"`
	SourceEngine        string `json:"sourceEngine" tf:"source_engine"`
	SourceEngineVersion string `json:"sourceEngineVersion" tf:"source_engine_version"`
}

type DbInstanceSpec struct {
	// +optional
	AllowMajorVersionUpgrade bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	// +optional
	CopyTagsToSnapshot bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`
	// +optional
	DeletionProtection bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`
	// +optional
	Domain string `json:"domain,omitempty" tf:"domain,omitempty"`
	// +optional
	DomainIamRoleName string `json:"domainIamRoleName,omitempty" tf:"domain_iam_role_name,omitempty"`
	// +optional
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports,omitempty"`
	// +optional
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`
	// +optional
	IamDatabaseAuthenticationEnabled bool   `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`
	InstanceClass                    string `json:"instanceClass" tf:"instance_class"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	MaxAllocatedStorage int `json:"maxAllocatedStorage,omitempty" tf:"max_allocated_storage,omitempty"`
	// +optional
	MonitoringInterval int `json:"monitoringInterval,omitempty" tf:"monitoring_interval,omitempty"`
	// +optional
	Password string `json:"password,omitempty" tf:"password,omitempty"`
	// +optional
	PerformanceInsightsEnabled bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`
	// +optional
	ReplicateSourceDb string `json:"replicateSourceDb,omitempty" tf:"replicate_source_db,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Import []DbInstanceSpecS3Import `json:"s3Import,omitempty" tf:"s3_import,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`
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

type DbInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbInstanceList is a list of DbInstances
type DbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbInstance CRD objects
	Items []DbInstance `json:"items,omitempty"`
}
