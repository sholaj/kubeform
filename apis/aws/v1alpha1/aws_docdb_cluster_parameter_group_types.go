package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDocdbClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDocdbClusterParameterGroupSpec   `json:"spec,omitempty"`
	Status            AwsDocdbClusterParameterGroupStatus `json:"status,omitempty"`
}

type AwsDocdbClusterParameterGroupSpecParameter struct {
	ApplyMethod string `json:"apply_method"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

type AwsDocdbClusterParameterGroupSpec struct {
	Parameter   []AwsDocdbClusterParameterGroupSpec `json:"parameter"`
	Tags        map[string]string                   `json:"tags"`
	Arn         string                              `json:"arn"`
	Name        string                              `json:"name"`
	NamePrefix  string                              `json:"name_prefix"`
	Family      string                              `json:"family"`
	Description string                              `json:"description"`
}

type AwsDocdbClusterParameterGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDocdbClusterParameterGroupList is a list of AwsDocdbClusterParameterGroups
type AwsDocdbClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDocdbClusterParameterGroup CRD objects
	Items []AwsDocdbClusterParameterGroup `json:"items,omitempty"`
}