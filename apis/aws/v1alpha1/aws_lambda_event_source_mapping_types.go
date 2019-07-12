package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaEventSourceMapping struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLambdaEventSourceMappingSpec   `json:"spec,omitempty"`
	Status            AwsLambdaEventSourceMappingStatus `json:"status,omitempty"`
}

type AwsLambdaEventSourceMappingSpec struct {
	LastProcessingResult      string `json:"last_processing_result"`
	StateTransitionReason     string `json:"state_transition_reason"`
	EventSourceArn            string `json:"event_source_arn"`
	BatchSize                 int    `json:"batch_size"`
	Enabled                   bool   `json:"enabled"`
	FunctionArn               string `json:"function_arn"`
	State                     string `json:"state"`
	Uuid                      string `json:"uuid"`
	FunctionName              string `json:"function_name"`
	StartingPosition          string `json:"starting_position"`
	StartingPositionTimestamp string `json:"starting_position_timestamp"`
	LastModified              string `json:"last_modified"`
}

type AwsLambdaEventSourceMappingStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaEventSourceMappingList is a list of AwsLambdaEventSourceMappings
type AwsLambdaEventSourceMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLambdaEventSourceMapping CRD objects
	Items []AwsLambdaEventSourceMapping `json:"items,omitempty"`
}