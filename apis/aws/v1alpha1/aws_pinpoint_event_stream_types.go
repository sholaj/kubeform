package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsPinpointEventStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsPinpointEventStreamSpec   `json:"spec,omitempty"`
	Status            AwsPinpointEventStreamStatus `json:"status,omitempty"`
}

type AwsPinpointEventStreamSpec struct {
	DestinationStreamArn string `json:"destination_stream_arn"`
	RoleArn              string `json:"role_arn"`
	ApplicationId        string `json:"application_id"`
}

type AwsPinpointEventStreamStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsPinpointEventStreamList is a list of AwsPinpointEventStreams
type AwsPinpointEventStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsPinpointEventStream CRD objects
	Items []AwsPinpointEventStream `json:"items,omitempty"`
}