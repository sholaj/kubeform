package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type BudgetsBudget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BudgetsBudgetSpec   `json:"spec,omitempty"`
	Status            BudgetsBudgetStatus `json:"status,omitempty"`
}

type BudgetsBudgetSpecNotification struct {
	ComparisonOperator string `json:"comparisonOperator" tf:"comparison_operator"`
	NotificationType   string `json:"notificationType" tf:"notification_type"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubscriberEmailAddresses []string `json:"subscriberEmailAddresses,omitempty" tf:"subscriber_email_addresses,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubscriberSnsTopicArns []string    `json:"subscriberSnsTopicArns,omitempty" tf:"subscriber_sns_topic_arns,omitempty"`
	Threshold              json.Number `json:"threshold" tf:"threshold"`
	ThresholdType          string      `json:"thresholdType" tf:"threshold_type"`
}

type BudgetsBudgetSpec struct {
	BudgetType  string `json:"budgetType" tf:"budget_type"`
	LimitAmount string `json:"limitAmount" tf:"limit_amount"`
	LimitUnit   string `json:"limitUnit" tf:"limit_unit"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Notification []BudgetsBudgetSpecNotification `json:"notification,omitempty" tf:"notification,omitempty"`
	// +optional
	TimePeriodEnd   string                    `json:"timePeriodEnd,omitempty" tf:"time_period_end,omitempty"`
	TimePeriodStart string                    `json:"timePeriodStart" tf:"time_period_start"`
	TimeUnit        string                    `json:"timeUnit" tf:"time_unit"`
	ProviderRef     core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BudgetsBudgetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BudgetsBudgetList is a list of BudgetsBudgets
type BudgetsBudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BudgetsBudget CRD objects
	Items []BudgetsBudget `json:"items,omitempty"`
}
