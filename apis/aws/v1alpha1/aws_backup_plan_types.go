package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBackupPlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsBackupPlanSpec   `json:"spec,omitempty"`
	Status            AwsBackupPlanStatus `json:"status,omitempty"`
}

type AwsBackupPlanSpecRuleLifecycle struct {
	ColdStorageAfter int `json:"cold_storage_after"`
	DeleteAfter      int `json:"delete_after"`
}

type AwsBackupPlanSpecRule struct {
	RuleName          string                  `json:"rule_name"`
	TargetVaultName   string                  `json:"target_vault_name"`
	Schedule          string                  `json:"schedule"`
	StartWindow       int                     `json:"start_window"`
	CompletionWindow  int                     `json:"completion_window"`
	Lifecycle         []AwsBackupPlanSpecRule `json:"lifecycle"`
	RecoveryPointTags map[string]string       `json:"recovery_point_tags"`
}

type AwsBackupPlanSpec struct {
	Name    string              `json:"name"`
	Rule    []AwsBackupPlanSpec `json:"rule"`
	Arn     string              `json:"arn"`
	Version string              `json:"version"`
	Tags    map[string]string   `json:"tags"`
}

type AwsBackupPlanStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBackupPlanList is a list of AwsBackupPlans
type AwsBackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsBackupPlan CRD objects
	Items []AwsBackupPlan `json:"items,omitempty"`
}