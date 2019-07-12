package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamAccessKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamAccessKeySpec   `json:"spec,omitempty"`
	Status            AwsIamAccessKeyStatus `json:"status,omitempty"`
}

type AwsIamAccessKeySpec struct {
	SesSmtpPassword string `json:"ses_smtp_password"`
	PgpKey          string `json:"pgp_key"`
	KeyFingerprint  string `json:"key_fingerprint"`
	EncryptedSecret string `json:"encrypted_secret"`
	User            string `json:"user"`
	Status          string `json:"status"`
	Secret          string `json:"secret"`
}

type AwsIamAccessKeyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamAccessKeyList is a list of AwsIamAccessKeys
type AwsIamAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamAccessKey CRD objects
	Items []AwsIamAccessKey `json:"items,omitempty"`
}