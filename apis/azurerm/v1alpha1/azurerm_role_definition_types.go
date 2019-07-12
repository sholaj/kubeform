package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermRoleDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRoleDefinitionSpec   `json:"spec,omitempty"`
	Status            AzurermRoleDefinitionStatus `json:"status,omitempty"`
}

type AzurermRoleDefinitionSpecPermissions struct {
	NotActions     []string `json:"not_actions"`
	DataActions    []string `json:"data_actions"`
	NotDataActions []string `json:"not_data_actions"`
	Actions        []string `json:"actions"`
}

type AzurermRoleDefinitionSpec struct {
	RoleDefinitionId string                      `json:"role_definition_id"`
	Name             string                      `json:"name"`
	Scope            string                      `json:"scope"`
	Description      string                      `json:"description"`
	Permissions      []AzurermRoleDefinitionSpec `json:"permissions"`
	AssignableScopes []string                    `json:"assignable_scopes"`
}

type AzurermRoleDefinitionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermRoleDefinitionList is a list of AzurermRoleDefinitions
type AzurermRoleDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRoleDefinition CRD objects
	Items []AzurermRoleDefinition `json:"items,omitempty"`
}