package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubeform.dev/kubeform/apis/google"
)

var SchemeGroupVersion = schema.GroupVersion{Group: google.GroupName, Version: "v1alpha1"}

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

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&Project{},
		&ProjectList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&DataprocJob{},
		&DataprocJobList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&StorageBucket{},
		&StorageBucketList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&BigtableTable{},
		&BigtableTableList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&Folder{},
		&FolderList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ProjectService{},
		&ProjectServiceList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&DataflowJob{},
		&DataflowJobList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
