package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchEventRuleSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchEventRuleStatus `json:"status,omitempty"`
}

type AwsCloudwatchEventRuleSpec struct {
	RoleArn            string            `json:"role_arn"`
	Arn                string            `json:"arn"`
	Name               string            `json:"name"`
	NamePrefix         string            `json:"name_prefix"`
	ScheduleExpression string            `json:"schedule_expression"`
	EventPattern       string            `json:"event_pattern"`
	Description        string            `json:"description"`
	IsEnabled          bool              `json:"is_enabled"`
	Tags               map[string]string `json:"tags"`
}

type AwsCloudwatchEventRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventRuleList is a list of AwsCloudwatchEventRules
type AwsCloudwatchEventRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchEventRule CRD objects
	Items []AwsCloudwatchEventRule `json:"items,omitempty"`
}