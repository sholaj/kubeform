package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermAutomationVariableBool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationVariableBoolSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationVariableBoolStatus `json:"status,omitempty"`
}

type AzurermAutomationVariableBoolSpec struct {
	ResourceGroupName     string `json:"resource_group_name"`
	Name                  string `json:"name"`
	AutomationAccountName string `json:"automation_account_name"`
	Description           string `json:"description"`
	Encrypted             bool   `json:"encrypted"`
	Value                 bool   `json:"value"`
}

type AzurermAutomationVariableBoolStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermAutomationVariableBoolList is a list of AzurermAutomationVariableBools
type AzurermAutomationVariableBoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationVariableBool CRD objects
	Items []AzurermAutomationVariableBool `json:"items,omitempty"`
}