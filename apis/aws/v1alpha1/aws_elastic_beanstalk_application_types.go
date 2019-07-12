package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticBeanstalkApplicationSpec   `json:"spec,omitempty"`
	Status            AwsElasticBeanstalkApplicationStatus `json:"status,omitempty"`
}

type AwsElasticBeanstalkApplicationSpecAppversionLifecycle struct {
	ServiceRole        string `json:"service_role"`
	MaxAgeInDays       int    `json:"max_age_in_days"`
	MaxCount           int    `json:"max_count"`
	DeleteSourceFromS3 bool   `json:"delete_source_from_s3"`
}

type AwsElasticBeanstalkApplicationSpec struct {
	Name                string                               `json:"name"`
	Description         string                               `json:"description"`
	AppversionLifecycle []AwsElasticBeanstalkApplicationSpec `json:"appversion_lifecycle"`
}

type AwsElasticBeanstalkApplicationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkApplicationList is a list of AwsElasticBeanstalkApplications
type AwsElasticBeanstalkApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticBeanstalkApplication CRD objects
	Items []AwsElasticBeanstalkApplication `json:"items,omitempty"`
}