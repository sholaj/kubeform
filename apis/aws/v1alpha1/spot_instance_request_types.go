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

type SpotInstanceRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotInstanceRequestSpec   `json:"spec,omitempty"`
	Status            SpotInstanceRequestStatus `json:"status,omitempty"`
}

type SpotInstanceRequestSpecCreditSpecification struct {
	// +optional
	CpuCredits string `json:"cpuCredits,omitempty" tf:"cpu_credits,omitempty"`
}

type SpotInstanceRequestSpec struct {
	Ami string `json:"ami" tf:"ami"`
	// +optional
	BlockDurationMinutes int `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CreditSpecification []SpotInstanceRequestSpecCreditSpecification `json:"creditSpecification,omitempty" tf:"credit_specification,omitempty"`
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
	// +optional
	InstanceInterruptionBehaviour string `json:"instanceInterruptionBehaviour,omitempty" tf:"instance_interruption_behaviour,omitempty"`
	InstanceType                  string `json:"instanceType" tf:"instance_type"`
	// +optional
	LaunchGroup string `json:"launchGroup,omitempty" tf:"launch_group,omitempty"`
	// +optional
	Monitoring bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`
	// +optional
	SourceDestCheck bool `json:"sourceDestCheck,omitempty" tf:"source_dest_check,omitempty"`
	// +optional
	SpotPrice string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`
	// +optional
	SpotType string `json:"spotType,omitempty" tf:"spot_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	// +optional
	UserDataBase64 string `json:"userDataBase64,omitempty" tf:"user_data_base64,omitempty"`
	// +optional
	VolumeTags map[string]string `json:"volumeTags,omitempty" tf:"volume_tags,omitempty"`
	// +optional
	WaitForFulfillment bool                      `json:"waitForFulfillment,omitempty" tf:"wait_for_fulfillment,omitempty"`
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SpotInstanceRequestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpotInstanceRequestList is a list of SpotInstanceRequests
type SpotInstanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpotInstanceRequest CRD objects
	Items []SpotInstanceRequest `json:"items,omitempty"`
}
