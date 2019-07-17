package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubeform.dev/kubeform/apis/aws"
)

var SchemeGroupVersion = schema.GroupVersion{Group: aws.GroupName, Version: "v1alpha1"}

var (
	// TODO: move SchemeBuilder with zz_generated.deepcopy.go to k8s.io/api.
	// localSchemeBuilder and AddToScheme will stay in k8s.io/kubernetes.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
}

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&MskCluster{},
		&MskClusterList{},

		&RdsCluster{},
		&RdsClusterList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&Instance{},
		&InstanceList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&IotCertificate{},
		&IotCertificateList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&Ami{},
		&AmiList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&Route{},
		&RouteList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&SqsQueue{},
		&SqsQueueList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&MqBroker{},
		&MqBrokerList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&Lb{},
		&LbList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&Codepipeline{},
		&CodepipelineList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&SfnActivity{},
		&SfnActivityList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&EcsService{},
		&EcsServiceList{},

		&IamPolicy{},
		&IamPolicyList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&BackupPlan{},
		&BackupPlanList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&TransferUser{},
		&TransferUserList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&SsmActivation{},
		&SsmActivationList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&AmiCopy{},
		&AmiCopyList{},

		&SwfDomain{},
		&SwfDomainList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&Alb{},
		&AlbList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&NetworkACL{},
		&NetworkACLList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&KmsGrant{},
		&KmsGrantList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&KmsKey{},
		&KmsKeyList{},

		&Subnet{},
		&SubnetList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&SsmParameter{},
		&SsmParameterList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&EmrCluster{},
		&EmrClusterList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&NatGateway{},
		&NatGatewayList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&Vpc{},
		&VpcList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&Elb{},
		&ElbList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&IamRole{},
		&IamRoleList{},

		&GlueJob{},
		&GlueJobList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&IotPolicy{},
		&IotPolicyList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&FlowLog{},
		&FlowLogList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&S3Bucket{},
		&S3BucketList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&DbInstance{},
		&DbInstanceList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&AlbListener{},
		&AlbListenerList{},

		&Eip{},
		&EipList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&IotThingType{},
		&IotThingTypeList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&DaxCluster{},
		&DaxClusterList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&EcsCluster{},
		&EcsClusterList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&BackupVault{},
		&BackupVaultList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&IotThing{},
		&IotThingList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&LbListener{},
		&LbListenerList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&DxGateway{},
		&DxGatewayList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&Route53Record{},
		&Route53RecordList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&SnsTopic{},
		&SnsTopicList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&RouteTable{},
		&RouteTableList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&KmsAlias{},
		&KmsAliasList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&TransferServer{},
		&TransferServerList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&IamGroup{},
		&IamGroupList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&EipAssociation{},
		&EipAssociationList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&DxConnection{},
		&DxConnectionList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&WafRule{},
		&WafRuleList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&SesTemplate{},
		&SesTemplateList{},

		&PinpointApp{},
		&PinpointAppList{},

		&EksCluster{},
		&EksClusterList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&WafIpset{},
		&WafIpsetList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&DxLag{},
		&DxLagList{},

		&IamUser{},
		&IamUserList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&WafWebACL{},
		&WafWebACLList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&KeyPair{},
		&KeyPairList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
