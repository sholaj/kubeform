package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRedshiftParameterGroupSpec   `json:"spec,omitempty"`
	Status            AwsRedshiftParameterGroupStatus `json:"status,omitempty"`
}

type AwsRedshiftParameterGroupSpecParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsRedshiftParameterGroupSpec struct {
	Name        string                          `json:"name"`
	Family      string                          `json:"family"`
	Description string                          `json:"description"`
	Parameter   []AwsRedshiftParameterGroupSpec `json:"parameter"`
}

type AwsRedshiftParameterGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftParameterGroupList is a list of AwsRedshiftParameterGroups
type AwsRedshiftParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRedshiftParameterGroup CRD objects
	Items []AwsRedshiftParameterGroup `json:"items,omitempty"`
}