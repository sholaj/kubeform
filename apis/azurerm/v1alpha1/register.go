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

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&SqlServer{},
		&SqlServerList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&Iothub{},
		&IothubList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&Eventhub{},
		&EventhubList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&StorageShare{},
		&StorageShareList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&ManagementLock{},
		&ManagementLockList{},

		&MysqlServer{},
		&MysqlServerList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&LbRule{},
		&LbRuleList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&RouteTable{},
		&RouteTableList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&MariadbServer{},
		&MariadbServerList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&AppService{},
		&AppServiceList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&Route{},
		&RouteList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ContainerService{},
		&ContainerServiceList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&CdnProfile{},
		&CdnProfileList{},

		&DataFactory{},
		&DataFactoryList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&Firewall{},
		&FirewallList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&IotDps{},
		&IotDpsList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&KeyVault{},
		&KeyVaultList{},

		&LbProbe{},
		&LbProbeList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&StorageQueue{},
		&StorageQueueList{},

		&NotificationHub{},
		&NotificationHubList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&SearchService{},
		&SearchServiceList{},

		&SignalrService{},
		&SignalrServiceList{},

		&StorageTable{},
		&StorageTableList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DevTestLab{},
		&DevTestLabList{},

		&DnsZone{},
		&DnsZoneList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&Subnet{},
		&SubnetList{},

		&ApiManagement{},
		&ApiManagementList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&SharedImage{},
		&SharedImageList{},

		&BatchPool{},
		&BatchPoolList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&Image{},
		&ImageList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&StorageAccount{},
		&StorageAccountList{},

		&Snapshot{},
		&SnapshotList{},

		&StorageContainer{},
		&StorageContainerList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&Lb{},
		&LbList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&BatchAccount{},
		&BatchAccountList{},

		&DnsARecord{},
		&DnsARecordList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&FunctionApp{},
		&FunctionAppList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&RedisCache{},
		&RedisCacheList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&PublicIP{},
		&PublicIPList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
