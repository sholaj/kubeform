package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamUserLoginProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamUserLoginProfileSpec   `json:"spec,omitempty"`
	Status            AwsIamUserLoginProfileStatus `json:"status,omitempty"`
}

type AwsIamUserLoginProfileSpec struct {
	KeyFingerprint        string `json:"key_fingerprint"`
	EncryptedPassword     string `json:"encrypted_password"`
	User                  string `json:"user"`
	PgpKey                string `json:"pgp_key"`
	PasswordResetRequired bool   `json:"password_reset_required"`
	PasswordLength        int    `json:"password_length"`
}

type AwsIamUserLoginProfileStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamUserLoginProfileList is a list of AwsIamUserLoginProfiles
type AwsIamUserLoginProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamUserLoginProfile CRD objects
	Items []AwsIamUserLoginProfile `json:"items,omitempty"`
}