package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecCreditSpecification struct {
	// +optional
	CpuCredits string `json:"cpuCredits,omitempty" tf:"cpu_credits,omitempty"`
}

type InstanceSpec struct {
	Ami string `json:"ami" tf:"ami"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CreditSpecification []InstanceSpecCreditSpecification `json:"creditSpecification,omitempty" tf:"credit_specification,omitempty"`
	// +optional
	DisableAPITermination bool `json:"disableAPITermination,omitempty" tf:"disable_api_termination,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	GetPasswordData bool `json:"getPasswordData,omitempty" tf:"get_password_data,omitempty"`
	// +optional
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`
	// +optional
	InstanceInitiatedShutdownBehavior string `json:"instanceInitiatedShutdownBehavior,omitempty" tf:"instance_initiated_shutdown_behavior,omitempty"`
	InstanceType                      string `json:"instanceType" tf:"instance_type"`
	// +optional
	Monitoring bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`
	// +optional
	SourceDestCheck bool `json:"sourceDestCheck,omitempty" tf:"source_dest_check,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	// +optional
	UserDataBase64 string                    `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
