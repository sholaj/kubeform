package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleKmsCryptoKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleKmsCryptoKeySpec   `json:"spec,omitempty"`
	Status            GoogleKmsCryptoKeyStatus `json:"status,omitempty"`
}

type GoogleKmsCryptoKeySpec struct {
	Name           string `json:"name"`
	KeyRing        string `json:"key_ring"`
	RotationPeriod string `json:"rotation_period"`
	SelfLink       string `json:"self_link"`
}

type GoogleKmsCryptoKeyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleKmsCryptoKeyList is a list of GoogleKmsCryptoKeys
type GoogleKmsCryptoKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleKmsCryptoKey CRD objects
	Items []GoogleKmsCryptoKey `json:"items,omitempty"`
}