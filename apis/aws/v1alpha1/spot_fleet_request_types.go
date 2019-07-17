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

type SpotFleetRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotFleetRequestSpec   `json:"spec,omitempty"`
	Status            SpotFleetRequestStatus `json:"status,omitempty"`
}

type SpotFleetRequestSpecLaunchSpecification struct {
	Ami string `json:"ami" tf:"ami"`
	// +optional
	AssociatePublicIPAddress bool `json:"associatePublicIPAddress,omitempty" tf:"associate_public_ip_address,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile,omitempty"`
	// +optional
	IamInstanceProfileArn string `json:"iamInstanceProfileArn,omitempty" tf:"iam_instance_profile_arn,omitempty"`
	InstanceType          string `json:"instanceType" tf:"instance_type"`
	// +optional
	Monitoring bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`
	// +optional
	PlacementTenancy string `json:"placementTenancy,omitempty" tf:"placement_tenancy,omitempty"`
	// +optional
	SpotPrice string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	// +optional
	WeightedCapacity string `json:"weightedCapacity,omitempty" tf:"weighted_capacity,omitempty"`
}

type SpotFleetRequestSpec struct {
	// +optional
	AllocationStrategy string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`
	// +optional
	ExcessCapacityTerminationPolicy string `json:"excessCapacityTerminationPolicy,omitempty" tf:"excess_capacity_termination_policy,omitempty"`
	// +optional
	FleetType    string `json:"fleetType,omitempty" tf:"fleet_type,omitempty"`
	IamFleetRole string `json:"iamFleetRole" tf:"iam_fleet_role"`
	// +optional
	InstanceInterruptionBehaviour string `json:"instanceInterruptionBehaviour,omitempty" tf:"instance_interruption_behaviour,omitempty"`
	// +optional
	InstancePoolsToUseCount int `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	LaunchSpecification []SpotFleetRequestSpecLaunchSpecification `json:"launchSpecification" tf:"launch_specification"`
	// +optional
	ReplaceUnhealthyInstances bool `json:"replaceUnhealthyInstances,omitempty" tf:"replace_unhealthy_instances,omitempty"`
	// +optional
	SpotPrice      string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`
	TargetCapacity int    `json:"targetCapacity" tf:"target_capacity"`
	// +optional
	TerminateInstancesWithExpiration bool `json:"terminateInstancesWithExpiration,omitempty" tf:"terminate_instances_with_expiration,omitempty"`
	// +optional
	ValidFrom string `json:"validFrom,omitempty" tf:"valid_from,omitempty"`
	// +optional
	ValidUntil string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
	// +optional
	WaitForFulfillment bool                      `json:"waitForFulfillment,omitempty" tf:"wait_for_fulfillment,omitempty"`
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SpotFleetRequestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpotFleetRequestList is a list of SpotFleetRequests
type SpotFleetRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpotFleetRequest CRD objects
	Items []SpotFleetRequest `json:"items,omitempty"`
}
