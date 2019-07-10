package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	ResourceCodePostgresVersion     = "pgversion"
	ResourceKindPostgresVersion     = "PostgresVersion"
	ResourceSingularPostgresVersion = "postgresversion"
	ResourcePluralPostgresVersion   = "postgresversions"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PostgresVersion defines a Postgres database version.
type PostgresVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresVersionSpec   `json:"spec,omitempty"`
	Status            PostgresVersionStatus `json:"status,omitempty"`
}

// PostgresVersionSpec is the spec for postgres version
type PostgresVersionSpec struct {
}

// PostgresVersionStatus is the spec for postgres version
type PostgresVersionStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PostgresVersionList is a list of PostgresVersions
type PostgresVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PostgresVersion CRD objects
	Items []PostgresVersion `json:"items,omitempty"`
}
