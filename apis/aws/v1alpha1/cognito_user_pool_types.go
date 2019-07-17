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

type CognitoUserPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolSpec   `json:"spec,omitempty"`
	Status            CognitoUserPoolStatus `json:"status,omitempty"`
}

type CognitoUserPoolSpecDeviceConfiguration struct {
	// +optional
	ChallengeRequiredOnNewDevice bool `json:"challengeRequiredOnNewDevice,omitempty" tf:"challenge_required_on_new_device,omitempty"`
	// +optional
	DeviceOnlyRememberedOnUserPrompt bool `json:"deviceOnlyRememberedOnUserPrompt,omitempty" tf:"device_only_remembered_on_user_prompt,omitempty"`
}

type CognitoUserPoolSpecEmailConfiguration struct {
	// +optional
	EmailSendingAccount string `json:"emailSendingAccount,omitempty" tf:"email_sending_account,omitempty"`
	// +optional
	ReplyToEmailAddress string `json:"replyToEmailAddress,omitempty" tf:"reply_to_email_address,omitempty"`
	// +optional
	SourceArn string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`
}

type CognitoUserPoolSpecSchemaNumberAttributeConstraints struct {
	// +optional
	MaxValue string `json:"maxValue,omitempty" tf:"max_value,omitempty"`
	// +optional
	MinValue string `json:"minValue,omitempty" tf:"min_value,omitempty"`
}

type CognitoUserPoolSpecSchemaStringAttributeConstraints struct {
	// +optional
	MaxLength string `json:"maxLength,omitempty" tf:"max_length,omitempty"`
	// +optional
	MinLength string `json:"minLength,omitempty" tf:"min_length,omitempty"`
}

type CognitoUserPoolSpecSchema struct {
	AttributeDataType string `json:"attributeDataType" tf:"attribute_data_type"`
	// +optional
	DeveloperOnlyAttribute bool `json:"developerOnlyAttribute,omitempty" tf:"developer_only_attribute,omitempty"`
	// +optional
	Mutable bool   `json:"mutable,omitempty" tf:"mutable,omitempty"`
	Name    string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NumberAttributeConstraints []CognitoUserPoolSpecSchemaNumberAttributeConstraints `json:"numberAttributeConstraints,omitempty" tf:"number_attribute_constraints,omitempty"`
	// +optional
	Required bool `json:"required,omitempty" tf:"required,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StringAttributeConstraints []CognitoUserPoolSpecSchemaStringAttributeConstraints `json:"stringAttributeConstraints,omitempty" tf:"string_attribute_constraints,omitempty"`
}

type CognitoUserPoolSpecSmsConfiguration struct {
	ExternalID   string `json:"externalID" tf:"external_id"`
	SnsCallerArn string `json:"snsCallerArn" tf:"sns_caller_arn"`
}

type CognitoUserPoolSpecUserPoolAddOns struct {
	AdvancedSecurityMode string `json:"advancedSecurityMode" tf:"advanced_security_mode"`
}

type CognitoUserPoolSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AliasAttributes []string `json:"aliasAttributes,omitempty" tf:"alias_attributes,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AutoVerifiedAttributes []string `json:"autoVerifiedAttributes,omitempty" tf:"auto_verified_attributes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DeviceConfiguration []CognitoUserPoolSpecDeviceConfiguration `json:"deviceConfiguration,omitempty" tf:"device_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EmailConfiguration []CognitoUserPoolSpecEmailConfiguration `json:"emailConfiguration,omitempty" tf:"email_configuration,omitempty"`
	// +optional
	MfaConfiguration string `json:"mfaConfiguration,omitempty" tf:"mfa_configuration,omitempty"`
	Name             string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Schema []CognitoUserPoolSpecSchema `json:"schema,omitempty" tf:"schema,omitempty"`
	// +optional
	SmsAuthenticationMessage string `json:"smsAuthenticationMessage,omitempty" tf:"sms_authentication_message,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SmsConfiguration []CognitoUserPoolSpecSmsConfiguration `json:"smsConfiguration,omitempty" tf:"sms_configuration,omitempty"`
	// +optional
	SmsVerificationMessage string `json:"smsVerificationMessage,omitempty" tf:"sms_verification_message,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	UserPoolAddOns []CognitoUserPoolSpecUserPoolAddOns `json:"userPoolAddOns,omitempty" tf:"user_pool_add_ons,omitempty"`
	// +optional
	UsernameAttributes []string                  `json:"usernameAttributes,omitempty" tf:"username_attributes,omitempty"`
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CognitoUserPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitoUserPoolList is a list of CognitoUserPools
type CognitoUserPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitoUserPool CRD objects
	Items []CognitoUserPool `json:"items,omitempty"`
}
