package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueCrawler struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlueCrawlerSpec   `json:"spec,omitempty"`
	Status            AwsGlueCrawlerStatus `json:"status,omitempty"`
}

type AwsGlueCrawlerSpecSchemaChangePolicy struct {
	DeleteBehavior string `json:"delete_behavior"`
	UpdateBehavior string `json:"update_behavior"`
}

type AwsGlueCrawlerSpecS3Target struct {
	Path       string   `json:"path"`
	Exclusions []string `json:"exclusions"`
}

type AwsGlueCrawlerSpecJdbcTarget struct {
	ConnectionName string   `json:"connection_name"`
	Path           string   `json:"path"`
	Exclusions     []string `json:"exclusions"`
}

type AwsGlueCrawlerSpecDynamodbTarget struct {
	Path string `json:"path"`
}

type AwsGlueCrawlerSpec struct {
	Role                  string               `json:"role"`
	SchemaChangePolicy    []AwsGlueCrawlerSpec `json:"schema_change_policy"`
	Arn                   string               `json:"arn"`
	S3Target              []AwsGlueCrawlerSpec `json:"s3_target"`
	JdbcTarget            []AwsGlueCrawlerSpec `json:"jdbc_target"`
	SecurityConfiguration string               `json:"security_configuration"`
	Configuration         string               `json:"configuration"`
	DatabaseName          string               `json:"database_name"`
	Description           string               `json:"description"`
	Schedule              string               `json:"schedule"`
	Classifiers           []string             `json:"classifiers"`
	DynamodbTarget        []AwsGlueCrawlerSpec `json:"dynamodb_target"`
	Name                  string               `json:"name"`
	TablePrefix           string               `json:"table_prefix"`
}

type AwsGlueCrawlerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueCrawlerList is a list of AwsGlueCrawlers
type AwsGlueCrawlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlueCrawler CRD objects
	Items []AwsGlueCrawler `json:"items,omitempty"`
}