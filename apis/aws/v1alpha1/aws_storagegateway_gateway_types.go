package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsStoragegatewayGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewayGatewaySpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewayGatewayStatus `json:"status,omitempty"`
}

type AwsStoragegatewayGatewaySpecSmbActiveDirectorySettings struct {
	DomainName string `json:"domain_name"`
	Password   string `json:"password"`
	Username   string `json:"username"`
}

type AwsStoragegatewayGatewaySpec struct {
	GatewayTimezone            string                         `json:"gateway_timezone"`
	SmbActiveDirectorySettings []AwsStoragegatewayGatewaySpec `json:"smb_active_directory_settings"`
	Arn                        string                         `json:"arn"`
	GatewayId                  string                         `json:"gateway_id"`
	GatewayIpAddress           string                         `json:"gateway_ip_address"`
	GatewayName                string                         `json:"gateway_name"`
	GatewayType                string                         `json:"gateway_type"`
	MediumChangerType          string                         `json:"medium_changer_type"`
	SmbGuestPassword           string                         `json:"smb_guest_password"`
	TapeDriveType              string                         `json:"tape_drive_type"`
	ActivationKey              string                         `json:"activation_key"`
}

type AwsStoragegatewayGatewayStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsStoragegatewayGatewayList is a list of AwsStoragegatewayGateways
type AwsStoragegatewayGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewayGateway CRD objects
	Items []AwsStoragegatewayGateway `json:"items,omitempty"`
}