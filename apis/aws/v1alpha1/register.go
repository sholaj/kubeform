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

		&GameliftAlias{},
		&GameliftAliasList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&Alb{},
		&AlbList{},

		&SwfDomain{},
		&SwfDomainList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&Elb{},
		&ElbList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&Route{},
		&RouteList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&EmrCluster{},
		&EmrClusterList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&BackupVault{},
		&BackupVaultList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&IotThing{},
		&IotThingList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&LbListener{},
		&LbListenerList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&TransferUser{},
		&TransferUserList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&SsmParameter{},
		&SsmParameterList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&NetworkACL{},
		&NetworkACLList{},

		&Lb{},
		&LbList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&IotCertificate{},
		&IotCertificateList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&EipAssociation{},
		&EipAssociationList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&RouteTable{},
		&RouteTableList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&MqBroker{},
		&MqBrokerList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&Subnet{},
		&SubnetList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&DbInstance{},
		&DbInstanceList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&DxLag{},
		&DxLagList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&IamUser{},
		&IamUserList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&IamPolicy{},
		&IamPolicyList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&TransferServer{},
		&TransferServerList{},

		&DxGateway{},
		&DxGatewayList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&Eip{},
		&EipList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&FlowLog{},
		&FlowLogList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&AmiCopy{},
		&AmiCopyList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SsmActivation{},
		&SsmActivationList{},

		&SfnActivity{},
		&SfnActivityList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&IotPolicy{},
		&IotPolicyList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&KmsAlias{},
		&KmsAliasList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&SesTemplate{},
		&SesTemplateList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&Ami{},
		&AmiList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&Vpc{},
		&VpcList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&Route53Record{},
		&Route53RecordList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&EksCluster{},
		&EksClusterList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&PinpointApp{},
		&PinpointAppList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&IamRole{},
		&IamRoleList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&WafIpset{},
		&WafIpsetList{},

		&BackupPlan{},
		&BackupPlanList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&WafRule{},
		&WafRuleList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&DaxCluster{},
		&DaxClusterList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&Instance{},
		&InstanceList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&RdsCluster{},
		&RdsClusterList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&EcsService{},
		&EcsServiceList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&KmsKey{},
		&KmsKeyList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&SnsTopic{},
		&SnsTopicList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&IotThingType{},
		&IotThingTypeList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&Codepipeline{},
		&CodepipelineList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&DxConnection{},
		&DxConnectionList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&IamGroup{},
		&IamGroupList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&S3Bucket{},
		&S3BucketList{},

		&SqsQueue{},
		&SqsQueueList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&KeyPair{},
		&KeyPairList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&AlbListener{},
		&AlbListenerList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&KmsGrant{},
		&KmsGrantList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&NatGateway{},
		&NatGatewayList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&WafWebACL{},
		&WafWebACLList{},

		&EcsCluster{},
		&EcsClusterList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&GlueJob{},
		&GlueJobList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&MskCluster{},
		&MskClusterList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
