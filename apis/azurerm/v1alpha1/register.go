package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubeform.dev/kubeform/apis/azurerm"
)

var SchemeGroupVersion = schema.GroupVersion{Group: azurerm.GroupName, Version: "v1alpha1"}

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

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&RouteTable{},
		&RouteTableList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&SqlServer{},
		&SqlServerList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&StorageTable{},
		&StorageTableList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&KeyVault{},
		&KeyVaultList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&AppService{},
		&AppServiceList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&PublicIP{},
		&PublicIPList{},

		&Route{},
		&RouteList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&RedisCache{},
		&RedisCacheList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&Image{},
		&ImageList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&Iothub{},
		&IothubList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&IotDps{},
		&IotDpsList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&NotificationHub{},
		&NotificationHubList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&MysqlServer{},
		&MysqlServerList{},

		&StorageAccount{},
		&StorageAccountList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&Eventhub{},
		&EventhubList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DnsZone{},
		&DnsZoneList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&SharedImage{},
		&SharedImageList{},

		&CdnProfile{},
		&CdnProfileList{},

		&ContainerService{},
		&ContainerServiceList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&BatchAccount{},
		&BatchAccountList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&Snapshot{},
		&SnapshotList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&DevTestLab{},
		&DevTestLabList{},

		&Lb{},
		&LbList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&StorageQueue{},
		&StorageQueueList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&LbProbe{},
		&LbProbeList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&DataFactory{},
		&DataFactoryList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&ApiManagement{},
		&ApiManagementList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&FunctionApp{},
		&FunctionAppList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ManagementLock{},
		&ManagementLockList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&StorageBlob{},
		&StorageBlobList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&Firewall{},
		&FirewallList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&LbRule{},
		&LbRuleList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&SearchService{},
		&SearchServiceList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&Subnet{},
		&SubnetList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&BatchPool{},
		&BatchPoolList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&SignalrService{},
		&SignalrServiceList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&MariadbServer{},
		&MariadbServerList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&StorageShare{},
		&StorageShareList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&DnsARecord{},
		&DnsARecordList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&StorageContainer{},
		&StorageContainerList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
