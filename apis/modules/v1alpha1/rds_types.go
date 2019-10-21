package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type RDS struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RDSSpec   `json:"spec,omitempty"`
	Status            RDSStatus `json:"status,omitempty"`
}

type RDSSpec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// The allocated storage in gigabytes
	AllocatedStorage string `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`
	// +optional
	// Indicates that major version upgrades are allowed. Changing this parameter does not result in an outage and the change is asynchronously applied as soon as possible
	AllowMajorVersionUpgrade bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`
	// +optional
	// Specifies whether any database modifications are applied immediately, or during the next maintenance window
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	// +optional
	// The Availability Zone of the RDS instance
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	// The days to retain backups for
	BackupRetentionPeriod json.Number `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`
	// +optional
	// The daily time range (in UTC) during which automated backups are created if they are enabled. Example: '09:46-10:16'. Must not overlap with maintenance_window
	BackupWindow string `json:"backupWindow,omitempty" tf:"backup_window,omitempty"`
	// +optional
	// (Optional) The character set name to use for DB encoding in Oracle instances. This can't be changed. See Oracle Character Sets Supported in Amazon RDS for more information
	CharacterSetName string `json:"characterSetName,omitempty" tf:"character_set_name,omitempty"`
	// +optional
	// On delete, copy all Instance tags to the final snapshot (if final_snapshot_identifier is specified)
	CopyTagsToSnapshot bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`
	// +optional
	// Whether to create a database instance
	CreateDbInstance bool `json:"createDbInstance,omitempty" tf:"create_db_instance,omitempty"`
	// +optional
	// Whether to create a database option group
	CreateDbOptionGroup bool `json:"createDbOptionGroup,omitempty" tf:"create_db_option_group,omitempty"`
	// +optional
	// Whether to create a database parameter group
	CreateDbParameterGroup bool `json:"createDbParameterGroup,omitempty" tf:"create_db_parameter_group,omitempty"`
	// +optional
	// Whether to create a database subnet group
	CreateDbSubnetGroup bool `json:"createDbSubnetGroup,omitempty" tf:"create_db_subnet_group,omitempty"`
	// +optional
	// Create IAM role with a defined name that permits RDS to send enhanced monitoring metrics to CloudWatch Logs.
	CreateMonitoringRole bool `json:"createMonitoringRole,omitempty" tf:"create_monitoring_role,omitempty"`
	// +optional
	// Name of DB subnet group. DB instance will be created in the VPC associated with the DB subnet group. If unspecified, will be created in the default VPC
	DbSubnetGroupName string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`
	// +optional
	// The database can't be deleted when this value is set to true.
	DeletionProtection bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`
	// +optional
	// List of log types to enable for exporting to CloudWatch logs. If omitted, no logs will be exported. Valid values (depending on engine): alert, audit, error, general, listener, slowquery, trace, postgresql (PostgreSQL), upgrade (PostgreSQL).
	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports,omitempty"`
	// +optional
	// The database engine to use
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	// The engine version to use
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	// The family of the DB parameter group
	Family string `json:"family,omitempty" tf:"family,omitempty"`
	// +optional
	// The name of your final DB snapshot when this DB instance is deleted.
	FinalSnapshotIdentifier string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`
	// +optional
	// Specifies whether or mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled
	IamDatabaseAuthenticationEnabled bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`
	// +optional
	// The name of the RDS instance, if omitted, Terraform will assign a random, unique identifier
	Identifier string `json:"identifier,omitempty" tf:"identifier,omitempty"`
	// +optional
	// The instance type of the RDS instance
	InstanceClass string `json:"instanceClass,omitempty" tf:"instance_class,omitempty"`
	// +optional
	// The amount of provisioned IOPS. Setting this implies a storage_type of 'io1'
	Iops json.Number `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	// The ARN for the KMS encryption key. If creating an encrypted replica, set this to the destination KMS ARN. If storage_encrypted is set to true and kms_key_id is not specified the default KMS key created in your account will be used
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	// License model information for this DB instance. Optional, but required for some DB engines, i.e. Oracle SE1
	LicenseModel string `json:"licenseModel,omitempty" tf:"license_model,omitempty"`
	// +optional
	// The window to perform maintenance in. Syntax: 'ddd:hh24:mi-ddd:hh24:mi'. Eg: 'Mon:00:00-Mon:03:00'
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
	// +optional
	// Specifies the major version of the engine that this option group should be associated with
	MajorEngineVersion string `json:"majorEngineVersion,omitempty" tf:"major_engine_version,omitempty"`
	// +optional
	// Specifies the value for Storage Autoscaling
	MaxAllocatedStorage json.Number `json:"maxAllocatedStorage,omitempty" tf:"max_allocated_storage,omitempty"`
	// +optional
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	MonitoringInterval json.Number `json:"monitoringInterval,omitempty" tf:"monitoring_interval,omitempty"`
	// +optional
	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to CloudWatch Logs. Must be specified if monitoring_interval is non-zero.
	MonitoringRoleArn string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn,omitempty"`
	// +optional
	// Name of the IAM role which will be created when create_monitoring_role is enabled.
	MonitoringRoleName string `json:"monitoringRoleName,omitempty" tf:"monitoring_role_name,omitempty"`
	// +optional
	// Specifies if the RDS instance is multi-AZ
	MultiAz bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`
	// +optional
	// The DB name to create. If omitted, no database is created initially
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	// The description of the option group
	OptionGroupDescription string `json:"optionGroupDescription,omitempty" tf:"option_group_description,omitempty"`
	// +optional
	// Name of the DB option group to associate. Setting this automatically disables option_group creation
	OptionGroupName string `json:"optionGroupName,omitempty" tf:"option_group_name,omitempty"`
	// +optional
	// A list of Options to apply.
	Options json.RawMessage `json:"options,omitempty" tf:"options,omitempty"`
	// +optional
	// Description of the DB parameter group to create
	ParameterGroupDescription string `json:"parameterGroupDescription,omitempty" tf:"parameter_group_description,omitempty"`
	// +optional
	// Name of the DB parameter group to associate or create
	ParameterGroupName string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`
	// +optional
	// A list of DB parameters (map) to apply
	Parameters json.RawMessage `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	// Password for the master DB user. Note that this may show up in logs, and it will be stored in the state file
	Password string `json:"password,omitempty" tf:"password,omitempty"`
	// +optional
	// Specifies whether Performance Insights are enabled
	PerformanceInsightsEnabled bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled,omitempty"`
	// +optional
	// The amount of time in days to retain Performance Insights data. Either 7 (7 days) or 731 (2 years).
	PerformanceInsightsRetentionPeriod json.Number `json:"performanceInsightsRetentionPeriod,omitempty" tf:"performance_insights_retention_period,omitempty"`
	// +optional
	// The port on which the DB accepts connections
	Port string `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	// Bool to control if instance is publicly accessible
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`
	// +optional
	// Specifies that this resource is a Replicate database, and to use this value as the source database. This correlates to the identifier of another Amazon RDS Database to replicate.
	ReplicateSourceDb string `json:"replicateSourceDb,omitempty" tf:"replicate_source_db,omitempty"`
	// +optional
	// Determines whether a final DB snapshot is created before the DB instance is deleted. If true is specified, no DBSnapshot is created. If false is specified, a DB snapshot is created before the DB instance is deleted, using the value from final_snapshot_identifier
	SkipFinalSnapshot bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`
	// +optional
	// Specifies whether or not to create this database from a snapshot. This correlates to the snapshot ID you'd find in the RDS console, e.g: rds:production-2015-06-26-06-05.
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`
	// +optional
	// Specifies whether the DB instance is encrypted
	StorageEncrypted bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
	// +optional
	// One of 'standard' (magnetic), 'gp2' (general purpose SSD), or 'io1' (provisioned IOPS SSD). The default is 'io1' if iops is specified, 'standard' if not. Note that this behaviour is different from the AWS web console, where the default is 'gp2'.
	StorageType string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
	// +optional
	// A list of VPC subnet IDs
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids,omitempty"`
	// +optional
	// A mapping of tags to assign to all resources
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// (Optional) Updated Terraform resource management timeouts. Applies to `aws_db_instance` in particular to permit resource management times
	Timeouts map[string]string `json:"timeouts,omitempty" tf:"timeouts,omitempty"`
	// +optional
	// (Optional) Time zone of the DB instance. timezone is currently only supported by Microsoft SQL Server. The timezone can only be set on creation. See MSSQL User Guide for more information.
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
	// +optional
	// Whether to use the parameter group name prefix or not
	UseParameterGroupNamePrefix bool `json:"useParameterGroupNamePrefix,omitempty" tf:"use_parameter_group_name_prefix,omitempty"`
	// +optional
	// Username for the master DB user
	Username string `json:"username,omitempty" tf:"username,omitempty"`
	// +optional
	// List of VPC security groups to associate
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type RDSOutput struct {
	// The address of the RDS instance
	ThisDbInstanceAddress string `json:"thisDbInstanceAddress" tf:"this_db_instance_address"`
	// The ARN of the RDS instance
	ThisDbInstanceArn string `json:"thisDbInstanceArn" tf:"this_db_instance_arn"`
	// The availability zone of the RDS instance
	ThisDbInstanceAvailabilityZone string `json:"thisDbInstanceAvailabilityZone" tf:"this_db_instance_availability_zone"`
	// The connection endpoint
	ThisDbInstanceEndpoint string `json:"thisDbInstanceEndpoint" tf:"this_db_instance_endpoint"`
	// The canonical hosted zone ID of the DB instance (to be used in a Route 53 Alias record)
	ThisDbInstanceHostedZoneID string `json:"thisDbInstanceHostedZoneID" tf:"this_db_instance_hosted_zone_id"`
	// The RDS instance ID
	ThisDbInstanceID string `json:"thisDbInstanceID" tf:"this_db_instance_id"`
	// The database name
	ThisDbInstanceName string `json:"thisDbInstanceName" tf:"this_db_instance_name"`
	// The database password (this password may be old, because Terraform doesn't track it after initial creation)
	ThisDbInstancePassword string `json:"thisDbInstancePassword" tf:"this_db_instance_password"`
	// The database port
	ThisDbInstancePort string `json:"thisDbInstancePort" tf:"this_db_instance_port"`
	// The RDS Resource ID of this instance
	ThisDbInstanceResourceID string `json:"thisDbInstanceResourceID" tf:"this_db_instance_resource_id"`
	// The RDS instance status
	ThisDbInstanceStatus string `json:"thisDbInstanceStatus" tf:"this_db_instance_status"`
	// The master username for the database
	ThisDbInstanceUsername string `json:"thisDbInstanceUsername" tf:"this_db_instance_username"`
	// The ARN of the db option group
	ThisDbOptionGroupArn string `json:"thisDbOptionGroupArn" tf:"this_db_option_group_arn"`
	// The db option group id
	ThisDbOptionGroupID string `json:"thisDbOptionGroupID" tf:"this_db_option_group_id"`
	// The ARN of the db parameter group
	ThisDbParameterGroupArn string `json:"thisDbParameterGroupArn" tf:"this_db_parameter_group_arn"`
	// The db parameter group id
	ThisDbParameterGroupID string `json:"thisDbParameterGroupID" tf:"this_db_parameter_group_id"`
	// The ARN of the db subnet group
	ThisDbSubnetGroupArn string `json:"thisDbSubnetGroupArn" tf:"this_db_subnet_group_arn"`
	// The db subnet group name
	ThisDbSubnetGroupID string `json:"thisDbSubnetGroupID" tf:"this_db_subnet_group_id"`
}

type RDSStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RDSOutput `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RDSList is a list of RDSs
type RDSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RDS CRD objects
	Items []RDS `json:"items,omitempty"`
}