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

type EmrCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrClusterSpec   `json:"spec,omitempty"`
	Status            EmrClusterStatus `json:"status,omitempty"`
}

type EmrClusterSpecBootstrapAction struct {
	// +optional
	Args []string `json:"args,omitempty" tf:"args,omitempty"`
	Name string   `json:"name" tf:"name"`
	Path string   `json:"path" tf:"path"`
}

type EmrClusterSpecEc2Attributes struct {
	// +optional
	AdditionalMasterSecurityGroups string `json:"additionalMasterSecurityGroups,omitempty" tf:"additional_master_security_groups,omitempty"`
	// +optional
	AdditionalSlaveSecurityGroups string `json:"additionalSlaveSecurityGroups,omitempty" tf:"additional_slave_security_groups,omitempty"`
	// +optional
	EmrManagedMasterSecurityGroup string `json:"emrManagedMasterSecurityGroup,omitempty" tf:"emr_managed_master_security_group,omitempty"`
	// +optional
	EmrManagedSlaveSecurityGroup string `json:"emrManagedSlaveSecurityGroup,omitempty" tf:"emr_managed_slave_security_group,omitempty"`
	InstanceProfile              string `json:"instanceProfile" tf:"instance_profile"`
	// +optional
	KeyName string `json:"keyName,omitempty" tf:"key_name,omitempty"`
	// +optional
	ServiceAccessSecurityGroup string `json:"serviceAccessSecurityGroup,omitempty" tf:"service_access_security_group,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
}

type EmrClusterSpecKerberosAttributes struct {
	// +optional
	AdDomainJoinPassword string `json:"adDomainJoinPassword,omitempty" tf:"ad_domain_join_password,omitempty"`
	// +optional
	AdDomainJoinUser string `json:"adDomainJoinUser,omitempty" tf:"ad_domain_join_user,omitempty"`
	// +optional
	CrossRealmTrustPrincipalPassword string `json:"crossRealmTrustPrincipalPassword,omitempty" tf:"cross_realm_trust_principal_password,omitempty"`
	KdcAdminPassword                 string `json:"kdcAdminPassword" tf:"kdc_admin_password"`
	Realm                            string `json:"realm" tf:"realm"`
}

type EmrClusterSpec struct {
	// +optional
	AdditionalInfo string `json:"additionalInfo,omitempty" tf:"additional_info,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Applications []string `json:"applications,omitempty" tf:"applications,omitempty"`
	// +optional
	AutoscalingRole string `json:"autoscalingRole,omitempty" tf:"autoscaling_role,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	BootstrapAction []EmrClusterSpecBootstrapAction `json:"bootstrapAction,omitempty" tf:"bootstrap_action,omitempty"`
	// +optional
	Configurations string `json:"configurations,omitempty" tf:"configurations,omitempty"`
	// +optional
	ConfigurationsJSON string `json:"configurationsJSON,omitempty" tf:"configurations_json,omitempty"`
	// +optional
	CustomAmiID string `json:"customAmiID,omitempty" tf:"custom_ami_id,omitempty"`
	// +optional
	EbsRootVolumeSize int `json:"ebsRootVolumeSize,omitempty" tf:"ebs_root_volume_size,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Ec2Attributes []EmrClusterSpecEc2Attributes `json:"ec2Attributes,omitempty" tf:"ec2_attributes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KerberosAttributes []EmrClusterSpecKerberosAttributes `json:"kerberosAttributes,omitempty" tf:"kerberos_attributes,omitempty"`
	// +optional
	LogURI       string `json:"logURI,omitempty" tf:"log_uri,omitempty"`
	Name         string `json:"name" tf:"name"`
	ReleaseLabel string `json:"releaseLabel" tf:"release_label"`
	// +optional
	SecurityConfiguration string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`
	ServiceRole           string `json:"serviceRole" tf:"service_role"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VisibleToAllUsers bool                      `json:"visibleToAllUsers,omitempty" tf:"visible_to_all_users,omitempty"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type EmrClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EmrClusterList is a list of EmrClusters
type EmrClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EmrCluster CRD objects
	Items []EmrCluster `json:"items,omitempty"`
}
