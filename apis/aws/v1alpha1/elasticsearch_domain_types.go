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

type ElasticsearchDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticsearchDomainSpec   `json:"spec,omitempty"`
	Status            ElasticsearchDomainStatus `json:"status,omitempty"`
}

type ElasticsearchDomainSpecCognitoOptions struct {
	// +optional
	Enabled        bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	IdentityPoolID string `json:"identityPoolID" tf:"identity_pool_id"`
	RoleArn        string `json:"roleArn" tf:"role_arn"`
	UserPoolID     string `json:"userPoolID" tf:"user_pool_id"`
}

type ElasticsearchDomainSpecLogPublishingOptions struct {
	CloudwatchLogGroupArn string `json:"cloudwatchLogGroupArn" tf:"cloudwatch_log_group_arn"`
	// +optional
	Enabled bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	LogType string `json:"logType" tf:"log_type"`
}

type ElasticsearchDomainSpecSnapshotOptions struct {
	AutomatedSnapshotStartHour int `json:"automatedSnapshotStartHour" tf:"automated_snapshot_start_hour"`
}

type ElasticsearchDomainSpecVpcOptions struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids,omitempty"`
}

type ElasticsearchDomainSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CognitoOptions []ElasticsearchDomainSpecCognitoOptions `json:"cognitoOptions,omitempty" tf:"cognito_options,omitempty"`
	DomainName     string                                  `json:"domainName" tf:"domain_name"`
	// +optional
	ElasticsearchVersion string `json:"elasticsearchVersion,omitempty" tf:"elasticsearch_version,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	LogPublishingOptions []ElasticsearchDomainSpecLogPublishingOptions `json:"logPublishingOptions,omitempty" tf:"log_publishing_options,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnapshotOptions []ElasticsearchDomainSpecSnapshotOptions `json:"snapshotOptions,omitempty" tf:"snapshot_options,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcOptions  []ElasticsearchDomainSpecVpcOptions `json:"vpcOptions,omitempty" tf:"vpc_options,omitempty"`
	ProviderRef core.LocalObjectReference           `json:"providerRef" tf:"-"`
}

type ElasticsearchDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticsearchDomainList is a list of ElasticsearchDomains
type ElasticsearchDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticsearchDomain CRD objects
	Items []ElasticsearchDomain `json:"items,omitempty"`
}
