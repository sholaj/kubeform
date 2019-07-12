package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorAssessmentTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsInspectorAssessmentTemplateSpec   `json:"spec,omitempty"`
	Status            AwsInspectorAssessmentTemplateStatus `json:"status,omitempty"`
}

type AwsInspectorAssessmentTemplateSpec struct {
	TargetArn        string   `json:"target_arn"`
	Arn              string   `json:"arn"`
	Duration         int      `json:"duration"`
	RulesPackageArns []string `json:"rules_package_arns"`
	Name             string   `json:"name"`
}

type AwsInspectorAssessmentTemplateStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInspectorAssessmentTemplateList is a list of AwsInspectorAssessmentTemplates
type AwsInspectorAssessmentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsInspectorAssessmentTemplate CRD objects
	Items []AwsInspectorAssessmentTemplate `json:"items,omitempty"`
}