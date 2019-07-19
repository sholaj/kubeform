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

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&DataflowJob{},
		&DataflowJobList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&Project{},
		&ProjectList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ProjectService{},
		&ProjectServiceList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&Folder{},
		&FolderList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
