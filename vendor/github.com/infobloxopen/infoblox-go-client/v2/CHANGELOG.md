# Infoblox Go Client Release Notes

## v2.6.0

### Release Summary

- Added a generic function to fetch objects by internal id.
- Added support for the http calls with headers to WAPI.

### Major Changes

- Added a generic function `SearchObjectByAltId` to fetch objects by internal id.
- Added validations for objects fetched by internal id.

### Minor Changes

- Added two functions `GetDnsMember` and `GetDhcpMember` to fetch DNS and DHCP members respectively.
- Added two functions `UpdateDnsStatus` and `UpdateDhcpStatus` to update DNS and DHCP status respectively.

## v2.4.0

### Release Summary

- Added 'object_generated' file contains auto generated WAPI object structures and associated functions from WAPI instance.
- Added E2E tests for functionality validation on WAPI's instance.
- Updated CHANGELOG file with structs and fields of corresponding objects generated.
- Some fields of structs are updated with pointers corresponding to WAPI instance.

### Major Changes

- Updated vendor directory and some dependencies like `ginkgo` with version update for test files.
- Added `object_generated` file contains supported objects from go-client and other WAPI objects with structs.
- Added Multi Value Extensible Attribute search validation for terraform plugin.
- Removed `record_ns` file as `RecordNS` object struct already exists in `object_generated` file.

## v2.1.1

This is just a bugfix release.

## v2.1.0

### Release Summary

- Enhancements in Host Record functionality
- Code refinements

### Minor Changes

- `SearchHostRecordByAltId` function to search Host Record By alternative ID from terraform.
- The code for every record has been seperated and added under a new file.


## v2.0.0

### Release Summary

Create, Update, Delete and Get operation on below records are being added or enhanced.

- Network View with comment and EAs
- IPv4 and IPv6 network containers with comment and EAs
- IPv4 and IPv6 network with comment and EAs
- Host Record with comment, EAs, EnableDns, EnableDhcp, UseTtl, Ttl, Alias attributes
- Fixed Address record with comment and EAs
- A record with comment, EAs, UseTtl, Ttl
- AAAA record with comment, EAs, UseTtl, Ttl
- PTR record with comment, EAs, UseTtl, Ttl
- Added IPv6 support for PTR record
- CNAME record with comment, EAs, UseTtl, Ttl
- Adds a compile-time check to the interface to make sure it stays in sync with the actual implementation.
- Added apt UTs and updated respective UTs

### Minor Changes

- Added default value of network view in AllocateIP, CreateHostRecord and CreatePTRRecord Function

### Bugfixes

- IPv6 Support `#86 <https://github.com/infobloxopen/infoblox-go-client/issues/86>`_
- Possibility to UPDATE a CNAME entry `#110 <https://github.com/infobloxopen/infoblox-go-client/issues/110>`_
- Feature Request: Ability to add comments `#116 <https://github.com/infobloxopen/infoblox-go-client/issues/116>`_
- Feature: Add ability to set extensible attributes when creating a network `#119 <https://github.com/infobloxopen/infoblox-go-client/issues/119>`_
- Feature request: add host aliases get/create/update `#126 <https://github.com/infobloxopen/infoblox-go-client/issues/126>`_

## v2.2.0

WAPI version used to generate the client is "2.12.1"

### Created structs:
- `DtcAllrecords`
- `PxgridEndpoint`
- `ThreatinsightCloudclient`
- `LicenseGridwide`
- `ParentalcontrolSubscribersite`
- `RecordDtclbdn`
- `AdmingroupDnsshowcommands`
- `DiscoveryDevicesupportbundle`
- `Ipv6dhcpoptiondefinition`
- `IPv6Range`
- `AdAuthServer`
- `GridRestartbannersetting`
- `GridServicerestartGroup`
- `Lease`
- `MemberDHCPProperties`
- `DiscoverySdnnetwork`
- `DtcMonitorHttp`
- `RecordNaptr`
- `Discovery`
- `Dhcpddns`
- `ParentalcontrolSitemember`
- `ThreatprotectionNatrule`
- `DdnsPrincipalclusterGroup`
- `ParentalcontrolSubscriberrecord`
- `RecordTlsa`
- `AdmingroupGridshowcommands`
- `Vlan`
- `ZoneRp`
- `AdmingroupLicensingshowcommands`
- `Allrecords`
- `DiscoveryStatus`
- `Filterrelayagent`
- `Filetransfersetting`
- `GridLoggingcategories`
- `Monitoreddomains`
- `Dtc`
- `Addressac`
- `DiscoveryAutoconversionsetting`
- `Ruleset`
- `Cacertificate`
- `Ipv6filteroption`
- `RoamingHost`
- `NsgroupForwardingmember`
- `Clientsubnetdomain`
- `Discoverydata`
- `GridAttackmitigation`
- `GridCspgridsetting`
- `Authpolicy`
- `Dhcpoptionspace`
- `DiscoveryDevicecomponent`
- `Msdnsserver`
- `SettingAtpoutbound`
- `NotificationRule`
- `DtcMonitorSip`
- `ParentalcontrolSubscriber`
- `ThreatprotectionNatport`
- `RirOrganization`
- `RecordMX`
- `RecordRpzCnameIpaddress`
- `Awsrte53zoneinfo`
- `ExtensibleattributedefDescendants`
- `DiscoveryMemberproperties`
- `GridServicerestartStatus`
- `Permission`
- `AdmingroupDatabaseshowcommands`
- `SettingViewaddress`
- `Mastergrid`
- `Trapreceiver`
- `View`
- `GridDhcpproperties`
- `Ipv6Network`
- `DiscoveryGridproperties`
- `DiscoveryClicredential`
- `Superhostchild`
- `MemberDnsgluerecordaddr`
- `GridLicensePool`
- `Ipv6fixedaddresstemplate`
- `Tftpfiledir`
- `DiscoveryVlaninfo`
- `MemberNtp`
- `Approvalworkflow`
- `DtcMonitorTcp`
- `Fixedaddresstemplate`
- `Dnsserver`
- `Mgmtportsetting`
- `SettingIpamThreshold`
- `SettingSchedule`
- `Ipv6rangetemplate`
- `SmartfolderGlobal`
- `Vlanrange`
- `SharedNetwork`
- `GridConsentbannersetting`
- `GridServicerestart`
- `GridServicerestartRequestChangedobject`
- `IPv4Address`
- `Ipv6dhcpoptionspace`
- `ZoneStub`
- `Lomuser`
- `GridDns`
- `RecordRpzSrv`
- `Upgradeschedule`
- `DtcMonitorPdp`
- `DtcHealth`
- `Ipv4Network`
- `GridCloudapiVmaddress`
- `HsmThalesgroup`
- `IPv6SharedNetwork`
- `Awsrte53task`
- `DiscoverySdnconfig`
- `DtcServerLink`
- `GridmemberSoaserial`
- `Sortlist`
- `MsserverAdsitesSite`
- `RecordDname`
- `AdmingroupSamlsetting`
- `Filteroption`
- `ParentalcontrolBlockingpolicy`
- `NsgroupStubmember`
- `RecordDhcid`
- `DtcServer`
- `GridMemberCloudapi`
- `AdmingroupDhcpshowcommands`
- `ThreatanalyticsAnalyticsWhitelist`
- `AdAuthService`
- `DbObjects`
- `GridThreatanalytics`
- `Exclusionrange`
- `HsmThales`
- `Adminrole`
- `Discoverytaskport`
- `DtcMonitorSnmp`
- `AdmingroupAdminsetcommands`
- `SmartfolderQueryitemvalue`
- `DtcRecordNaptr`
- `GridMaxminddbinfo`
- `OutboundCloudclient`
- `SettingSecuritybanner`
- `Discoverytaskvserver`
- `GridCloudapiInfo`
- `RadiusServer`
- `DiscoveryBasicsdnpollsettings`
- `GridCloudapiUser`
- `Superhost`
- `ThreatprotectionRuleset`
- `ThreatprotectionRuletemplate`
- `GridThreatprotection`
- `Networkuser`
- `RecordAlias`
- `DiscoveryPort`
- `DiscoveryCredentialgroup`
- `DiscoveryDevice`
- `GridCloudapiVm`
- `DiscoveryStatusinfo`
- `DdnsPrincipalcluster`
- `AdmingroupDockersetcommands`
- `AdmingroupMachinecontroltoplevelcommands`
- `Logicfilterrule`
- `Discoverytask`
- `Nsgroup`
- `Bulkhostnametemplate`
- `DxlEndpoint`
- `SharedrecordSrv`
- `AdmingroupAdminshowcommands`
- `SettingNetwork`
- `DiscoveryBasicpollsettings`
- `GridLicensesubpool`
- `AdmingroupDockershowcommands`
- `AdmingroupLockoutsetting`
- `DtcPoolLink`
- `ThreatanalyticsWhitelist`
- `Changedobject`
- `OutboundCloudclientEvent`
- `SettingPassword`
- `SettingSnmp`
- `RecordNsec3`
- `TacacsplusAuthservice`
- `DiscoveryPortConfigAdminstatus`
- `Scheduledbackup`
- `IPv6NetworkTemplate`
- `Lomnetworkconfig`
- `Memberserver`
- `RecordRpzCnameClientipaddressdn`
- `Lanhaportsetting`
- `Fingerprint`
- `Ipv4FixedAddress`
- `RecordRpzCname`
- `Ntpkey`
- `Dbsnapshot`
- `HsmSafenetgroup`
- `AdmingroupNetworkingshowcommands`
- `Ntpaccess`
- `Physicalportsetting`
- `SettingDynamicratio`
- `DtcTopology`
- `RecordRpzAaaaIpaddress`
- `SmartfolderPersonal`
- `DiscoveryAdvancedpollsetting`
- `MemberCspmembersetting`
- `Allnsgroup`
- `DiscoveryIfaddrinfo`
- `LdapEamapping`
- `DtcPool`
- `Trapnotification`
- `DtcCertificate`
- `NetworkviewAssocmember`
- `ZoneNameServer`
- `Adminuser`
- `AdmingroupNetworkingsetcommands`
- `Dhcpmember`
- `SettingAutomatedtrafficcapture`
- `DiscoverySnmpcredential`
- `Dnstapsetting`
- `Ipv6networksetting`
- `TaxiiRpzconfig`
- `MemberDnsip`
- `ParentalcontrolMsp`
- `CertificateAuthservice`
- `Extsyslogbackupserver`
- `GridServicerestartGroupSchedule`
- `Vtftpdirmember`
- `DiscoverySeedrouter`
- `NotificationRestTemplateinstance`
- `OcspResponder`
- `Exclusionrangetemplate`
- `SettingTrafficcapturechr`
- `LdapAuthService`
- `NotificationRestTemplate`
- `DtcPoolConsolidatedMonitorHealth`
- `Memberservicestatus`
- `SmartfolderQueryitem`
- `CiscoiseEndpoint`
- `ThreatanalyticsModuleset`
- `AdmingroupDnssetcommands`
- `RecordRpzAaaa`
- `DiscoveryDevicePortstatistics`
- `DiscoveryPortControlTaskdetails`
- `DtcRecordA`
- `DtcRecordSrv`
- `Orderedresponsepolicyzones`
- `Admingroup`
- `ZoneAuthDiscrepancy`
- `SettingScavenging`
- `ThreatprotectionGridRule`
- `EADefListValue`
- `GridAttackdetect`
- `GridCspapiconfig`
- `SettingTriggerrecdnslatency`
- `MACFilterAddress`
- `Mssuperscope`
- `RadiusAuthservice`
- `LdapServer`
- `ThreatprotectionRuleparam`
- `Bfdtemplate`
- `GridCloudapiCloudstatistics`
- `MemberFiledistribution`
- `Sharedrecordgroup`
- `DhcpStatistics`
- `Filterfingerprint`
- `GridServicerestartRequest`
- `Nxdomainrule`
- `Csvimporttask`
- `NotificationRestEndpoint`
- `GridAutoblackhole`
- `Preprovision`
- `MsserverAdsitesDomain`
- `ThreatprotectionStatistics`
- `DiscoveryPortConfigVlaninfo`
- `RecordUnknown`
- `AdmingroupDhcpsetcommands`
- `AdmingroupLicensingsetcommands`
- `DiscoveryVrfmappingrule`
- `NotificationRuleexpressionop`
- `Allrpzrecords`
- `Ipv6NetworkContainer`
- `Msserver`
- `Dhcpoption`
- `Syslogserver`
- `ScheduledTask`
- `Vdiscoverytask`
- `CapacityreportObjectcount`
- `SamlAuthservice`
- `Ipv6setting`
- `Natsetting`
- `Updatesdownloadmemberconfig`
- `Range`
- `RecordNsec`
- `Filtermac`
- `RecordRrsig`
- `Awsrte53recordinfo`
- `GridCloudapiTenant`
- `DiscoveryAdvancedsdnpollsettings`
- `Vlanlink`
- `DiscoveryNetworkinfo`
- `Dnssectrustedkey`
- `Expressionop`
- `Rdatasubfield`
- `SamlIdp`
- `DtcMonitor`
- `RecordRpzPtr`
- `SNMPUser`
- `SmartfolderGroupby`
- `Ipv4NetworkContainer`
- `ParentalcontrolNasgateway`
- `DtcRecordCname`
- `MsserverDcnsrecordcreation`
- `SettingIpamTrap`
- `ThreatprotectionStatinfo`
- `RecordNS`
- `Upgradegroup`
- `Memberservicecommunication`
- `SharedrecordCname`
- `ThreatprotectionRule`
- `ZoneForward`
- `NotificationRestTemplateparameter`
- `MemberThreatprotection`
- `RecordRpzCnameIpaddressdn`
- `Recordnamepolicy`
- `GridCloudapiGatewayEndpointmapping`
- `SettingDnsresolver`
- `TacacsplusServer`
- `Search`
- `RecordRpzMx`
- `Upgradestep`
- `Zoneassociation`
- `DtcLbdn`
- `Filternac`
- `Namedacl`
- `SettingInactivelockout`
- `ThreatprotectionRulecategory`
- `AdmingroupDnstoplevelcommands`
- `Filterrule`
- `PropertiesBlackoutsetting`
- `Captiveportal`
- `LocaluserAuthservice`
- `Memberdfp`
- `AdmingroupTroubleshootingtoplevelcommands`
- `DtcTopologyRuleSource`
- `GridResponseratelimiting`
- `DtcMonitorIcmp`
- `GridInformationalbannersetting`
- `Awsrte53taskgroup`
- `Bulkhost`
- `GridFiledistribution`
- `AdmingroupCloudsetcommands`
- `Hotfix`
- `Tsigac`
- `Ftpuser`
- `RecordRpzNaptr`
- `Restartservicestatus`
- `DeletedObjects`
- `DiscoverySnmp3credential`
- `Queriesuser`
- `SettingMsserver`
- `SettingTrafficcaptureqps`
- `ParentalcontrolAbs`
- `IPv6Address`
- `NetworkTemplate`
- `SharedRecordTXT`
- `RecordRpzTxt`
- `Taxii`
- `IpamStatistics`
- `NsgroupDelegation`
- `HsmAllgroups`
- `Dnsseckey`
- `DtcServerMonitor`
- `DiscoveryDeviceneighbor`
- `RecordDnskey`
- `GridmemberSoamname`
- `CaptiveportalFile`
- `CiscoiseEaassociation`
- `MemberLicense`
- `Natgroup`
- `Rangetemplate`
- `SharedRecordAAAA`
- `DtcObject`
- `Orderedranges`
- `RecordRpzAIpaddress`
- `SettingSecurity`
- `MemberDns`
- `Msdhcpoption`
- `SettingHttpproxyserver`
- `MsserverAduserData`
- `UpgradegroupMember`
- `Scavengingtask`
- `Dhcpserver`
- `FireeyeAlertmap`
- `SshKey`
- `NetworkDiscovery`
- `AdmingroupCloudshowcommands`
- `Option60matchrule`
- `Adsites`
- `DiscoveryDeviceinterface`
- `ThreatprotectionProfile`
- `AdmingroupGridsetcommands`
- `Dnsseckeyalgorithm`
- `Interface`
- `UpgradegroupSchedule`
- `Ipv6FixedAddress`
- `MemberParentalcontrol`
- `AdmingroupPasswordsetting`
- `CiscoisePublishsetting`
- `MsserverAduser`
- `SettingSyslogproxy`
- `Distributionschedule`
- `GridDashboard`
- `NsgroupForwardstubserver`
- `Bgpas`
- `SyslogEndpointServers`
- `DtcRecordAaaa`
- `Kerberoskey`
- `ParentalcontrolAvp`
- `Servicestatus`
- `ThreatprotectionRuleconfig`
- `DiscoveryVrf`
- `GridLicensePoolContainer`
- `RecordRpzA`
- `SettingTriggerrecqueries`
- `Dhcpoptiondefinition`
- `DxlEndpointBroker`
- `Lan2portsetting`
- `DiscoveryDiagnostictask`
- `Ospf`
- `RecordCaa`
- `AdmingroupSecuritysetcommands`
- `Bgpneighbor`
- `Msdhcpserver`
- `GridCloudapi`
- `RecordNsec3param`
- `DiscoveryAdvisorsetting`
- `Dnsseckeyparams`
- `Hostnamerewritepolicy`
- `GridDnsFixedrrsetorderfqdn`
- `SharedRecordMX`
- `ThreatprotectionProfileRule`
- `AdmingroupDatabasesetcommands`
- `Remoteddnszone`
- `Allendpoints`
- `Fileop`
- `Forwardingmemberserver`
- `HsmSafenet`
- `Thresholdtrap`
- `MemberThreatanalytics`
- `DtcMonitorSnmpOid`
- `SettingTriggeruthdnslatency`
- `Dhcpfailover`
- `SharedRecordA`
- `Vlanview`
- `SyslogEndpoint`
- `Nodeinfo`
- `Awsuser`
- `DtcTopologyLabel`
- `SmartfolderChildren`
- `GridCloudapiGatewayConfig`
- `ParentalcontrolSpm`
- `GridServicerestartGroupOrder`
- `AdmingroupSecurityshowcommands`
- `Eaexpressionop`
- `Dns64group`
- `RecordSRV`
- `DiscoveryPortConfigDescription`
- `GridLockoutsetting`
- `Preprovisionhardware`
- `GridX509certificate`
- `MsserverDns`
- `RecordRpzCnameClientipaddress`
- `DtcTopologyRule`
- `SettingEmail`
- `Ntpac`
- `MsserverDhcp`
- `AdmingroupAdmintoplevelcommands`
- `FireeyeRulemapping`
- `DiscoveryScaninterface`
- `Objectschangestrackingsetting`
- `RecordDs`
- `Rir`
- `CiscoiseSubscribesetting`

### Deleted structs:
- `PhysicalPortSetting`
- `License`
- `ServiceStatus`
- `SubElementsStatus`
- `Network`
- `LanHaPortSetting`
- `NetworkContainer`
- `FixedAddress`
- `IBBase`
- `NetworkSetting`
- `NodeInfo`
- `NetworkContainerNextAvailableInfo`
- `NetworkContainerNextAvailable`
- `Ipv6Setting`
- `QueryParams`
- `RequestBody`
- `SingleRequest`
- `MultiRequest`

### Updated structs:
##### `CapacityReport`

Created Fields:
- `ObjectCounts`

Deleted Fields:
- `ObjectCount`

##### `EADefinition`

Created Fields:
- `DefaultValue`
- `DescendantsAction`
- `Max`
- `Min`
- `Namespace`

##### `Grid`

Created Fields:
- `AllowRecursiveDeletion`
- `AuditLogFormat`
- `AuditToSyslogEnable`
- `AutomatedTrafficCaptureSetting`
- `ConsentBannerSetting`
- `CspApiConfig`
- `CspGridSetting`
- `DenyMgmSnapshots`
- `DescendantsAction`
- `DnsResolverSetting`
- `Dscp`
- `EmailSetting`
- `EnableGuiApiForLanVip`
- `EnableLom`
- `EnableMemberRedirect`
- `EnableRecycleBin`
- `EnableRirSwip`
- `ExternalSyslogBackupServers`
- `ExternalSyslogServerEnable`
- `HttpProxyServerSetting`
- `InformationalBannerSetting`
- `IsGridVisualizationVisible`
- `LockoutSetting`
- `LomUsers`
- `MgmStrictDelegateMode`
- `MsSetting`
- `NatGroups`
- `ObjectsChangesTrackingSetting`
- `PasswordSetting`
- `RestartBannerSetting`
- `RestartStatus`
- `RpzHitRateInterval`
- `RpzHitRateMaxQuery`
- `RpzHitRateMinQuery`
- `ScheduledBackup`
- `Secret`
- `SecurityBannerSetting`
- `SecuritySetting`
- `ServiceStatus`
- `SnmpSetting`
- `SupportBundleDownloadTimeout`
- `SyslogFacility`
- `SyslogServers`
- `SyslogSize`
- `ThresholdTraps`
- `TimeZone`
- `TokenUsageDelay`
- `TrafficCaptureAuthDnsSetting`
- `TrafficCaptureChrSetting`
- `TrafficCaptureQpsSetting`
- `TrafficCaptureRecDnsSetting`
- `TrafficCaptureRecQueriesSetting`
- `TrapNotifications`
- `UpdatesDownloadMemberConfig`
- `VpnPort`

##### `HostRecord`

Created Fields:
- `AllowTelnet`
- `CliCredentials`
- `CloudInfo`
- `CreationTime`
- `DdnsProtected`
- `DeviceDescription`
- `DeviceLocation`
- `DeviceType`
- `DeviceVendor`
- `Disable`
- `DisableDiscovery`
- `DnsAliases`
- `DnsName`
- `EnableImmediateDiscovery`
- `LastQueried`
- `MsAdUserData`
- `RestartIfNeeded`
- `RrsetOrder`
- `Snmp3Credential`
- `SnmpCredential`
- `UseCliCredentials`
- `UseSnmp3Credential`
- `UseSnmpCredential`

Deleted Fields:
- `Ipv4Addr`
- `Ipv6Addr`

##### `HostRecordIpv4Addr`

Created Fields:
- `Bootfile`
- `Bootserver`
- `DenyBootp`
- `DiscoverNowStatus`
- `DiscoveredData`
- `EnablePxeLeaseTime`
- `Host`
- `IgnoreClientRequestedOptions`
- `IsInvalidMac`
- `LastQueried`
- `LogicFilterRules`
- `MatchClient`
- `MsAdUserData`
- `Network`
- `NetworkView`
- `Nextserver`
- `Options`
- `PxeLeaseTime`
- `ReservedInterface`
- `UseBootfile`
- `UseBootserver`
- `UseDenyBootp`
- `UseForEaInheritance`
- `UseIgnoreClientRequestedOptions`
- `UseLogicFilterRules`
- `UseNextserver`
- `UseOptions`
- `UsePxeLeaseTime`

Deleted Fields:
- `Cidr`
- `View`

##### `HostRecordIpv6Addr`

Created Fields:
- `AddressType`
- `DiscoverNowStatus`
- `DiscoveredData`
- `DomainName`
- `DomainNameServers`
- `Host`
- `Ipv6prefix`
- `Ipv6prefixBits`
- `LastQueried`
- `LogicFilterRules`
- `MatchClient`
- `MsAdUserData`
- `Network`
- `NetworkView`
- `Options`
- `PreferredLifetime`
- `ReservedInterface`
- `UseDomainName`
- `UseDomainNameServers`
- `UseForEaInheritance`
- `UseLogicFilterRules`
- `UseOptions`
- `UsePreferredLifetime`
- `UseValidLifetime`
- `ValidLifetime`

Deleted Fields:
- `Cidr`
- `View`

##### `Member`

Created Fields:
- `ActivePosition`
- `AdditionalIpList`
- `AutomatedTrafficCaptureSetting`
- `BgpAs`
- `Comment`
- `CspAccessKey`
- `CspMemberSetting`
- `DnsResolverSetting`
- `Dscp`
- `Ea`
- `EmailSetting`
- `EnableHa`
- `EnableLom`
- `EnableMemberRedirect`
- `EnableRoApiAccess`
- `ExternalSyslogBackupServers`
- `ExternalSyslogServerEnable`
- `Ipv6Setting`
- `Ipv6StaticRoutes`
- `IsDscpCapable`
- `Lan2Enabled`
- `Lan2PortSetting`
- `LcdInput`
- `LomNetworkConfig`
- `LomUsers`
- `MasterCandidate`
- `MemberServiceCommunication`
- `MgmtPortSetting`
- `MmdbEaBuildTime`
- `MmdbGeoipBuildTime`
- `NTPSetting`
- `NatSetting`
- `NodeInfo`
- `OspfList`
- `PassiveHaArpEnabled`
- `Platform`
- `PreProvisioning`
- `PreserveIfOwnsDelegation`
- `RemoteConsoleAccessEnable`
- `RouterId`
- `ServiceStatus`
- `SnmpSetting`
- `StaticRoutes`
- `SupportAccessEnable`
- `SupportAccessInfo`
- `SyslogProxySetting`
- `SyslogServers`
- `SyslogSize`
- `ThresholdTraps`
- `TrafficCaptureAuthDnsSetting`
- `TrafficCaptureChrSetting`
- `TrafficCaptureQpsSetting`
- `TrafficCaptureRecDnsSetting`
- `TrafficCaptureRecQueriesSetting`
- `TrapNotifications`
- `UpgradeGroup`
- `UseAutomatedTrafficCapture`
- `UseDnsResolverSetting`
- `UseDscp`
- `UseEmailSetting`
- `UseEnableLom`
- `UseEnableMemberRedirect`
- `UseExternalSyslogBackupServers`
- `UseLcdInput`
- `UseRemoteConsoleAccessEnable`
- `UseSnmpSetting`
- `UseSupportAccessEnable`
- `UseSyslogProxySetting`
- `UseThresholdTraps`
- `UseTimeZone`
- `UseTrafficCaptureAuthDns`
- `UseTrafficCaptureChr`
- `UseTrafficCaptureQps`
- `UseTrafficCaptureRecDns`
- `UseTrafficCaptureRecQueries`
- `UseTrapNotifications`
- `UseV4Vrrp`
- `VipSetting`
- `VpnMtu`

Deleted Fields:
- `Nodeinfo`
- `PLATFORM`

##### `NTPSetting`

Created Fields:
- `GmLocalNtpStratum`
- `LocalNtpStratum`
- `UseDefaultStratum`

##### `NTPserver`

Created Fields:
- `NtpKeyNumber`
- `Preferred`

Deleted Fields:
- `NTPKeyNumber`
- `Preffered`

##### `NameServer`

Created Fields:
- `SharedWithMsParentDelegation`
- `Stealth`
- `TsigKey`
- `TsigKeyAlg`
- `TsigKeyName`
- `UseTsigKeyName`

##### `NetworkView`

Created Fields:
- `AssociatedDnsViews`
- `AssociatedMembers`
- `CloudInfo`
- `DdnsDnsView`
- `DdnsZonePrimaries`
- `InternalForwardZones`
- `IsDefault`
- `MgmPrivate`
- `MsAdUserData`
- `RemoteForwardZones`
- `RemoteReverseZones`

##### `RecordA`

Created Fields:
- `AwsRte53RecordInfo`
- `CloudInfo`
- `CreationTime`
- `Creator`
- `DdnsPrincipal`
- `DdnsProtected`
- `Disable`
- `DiscoveredData`
- `DnsName`
- `ForbidReclamation`
- `LastQueried`
- `MsAdUserData`
- `Reclaimable`
- `RemoveAssociatedPtr`
- `SharedRecordGroup`

##### `RecordAAAA`

Created Fields:
- `AwsRte53RecordInfo`
- `CloudInfo`
- `CreationTime`
- `Creator`
- `DdnsPrincipal`
- `DdnsProtected`
- `Disable`
- `DiscoveredData`
- `DnsName`
- `ForbidReclamation`
- `LastQueried`
- `MsAdUserData`
- `Reclaimable`
- `RemoveAssociatedPtr`
- `SharedRecordGroup`

##### `RecordCNAME`

Created Fields:
- `AwsRte53RecordInfo`
- `CloudInfo`
- `CreationTime`
- `Creator`
- `DdnsPrincipal`
- `DdnsProtected`
- `Disable`
- `DnsCanonical`
- `DnsName`
- `ForbidReclamation`
- `LastQueried`
- `Reclaimable`
- `SharedRecordGroup`

##### `RecordPTR`

Created Fields:
- `AwsRte53RecordInfo`
- `CloudInfo`
- `CreationTime`
- `Creator`
- `DdnsPrincipal`
- `DdnsProtected`
- `Disable`
- `DiscoveredData`
- `DnsName`
- `DnsPtrdname`
- `ForbidReclamation`
- `LastQueried`
- `MsAdUserData`
- `Reclaimable`
- `SharedRecordGroup`

##### `RecordTXT`

Created Fields:
- `AwsRte53RecordInfo`
- `CloudInfo`
- `CreationTime`
- `Creator`
- `DdnsPrincipal`
- `DdnsProtected`
- `Disable`
- `DnsName`
- `ForbidReclamation`
- `LastQueried`
- `Reclaimable`
- `SharedRecordGroup`

##### `UpgradeStatus`

Created Fields:
- `AllowDistribution`
- `AllowDistributionScheduling`
- `AllowUpgrade`
- `AllowUpgradeCancel`
- `AllowUpgradePause`
- `AllowUpgradeResume`
- `AllowUpgradeScheduling`
- `AllowUpgradeTest`
- `AllowUpload`
- `AlternateVersion`
- `Comment`
- `CurrentVersion`
- `CurrentVersionSummary`
- `DistributionScheduleActive`
- `DistributionScheduleTime`
- `DistributionState`
- `DistributionVersion`
- `DistributionVersionSummary`
- `ElementStatus`
- `GridState`
- `GroupState`
- `HaStatus`
- `Hotfixes`
- `Ipv4Address`
- `Ipv6Address`
- `Member`
- `Message`
- `PnodeRole`
- `Reverted`
- `StatusTime`
- `StatusValue`
- `StatusValueUpdateTime`
- `Steps`
- `StepsCompleted`
- `StepsTotal`
- `SubelementType`
- `SubelementsCompleted`
- `SubelementsStatus`
- `SubelementsTotal`
- `UpgradeScheduleActive`
- `UpgradeState`
- `UpgradeTestStatus`
- `UploadVersion`
- `UploadVersionSummary`

Deleted Fields:
- `SubElementStatus`

##### `UserProfile`

Created Fields:
- `ActiveDashboardType`
- `AdminGroup`
- `DaysToExpire`
- `Email`
- `GlobalSearchOnEa`
- `GlobalSearchOnNiData`
- `GridAdminGroups`
- `LastLogin`
- `LbTreeNodesAtGenLevel`
- `LbTreeNodesAtLastLevel`
- `MaxCountWidgets`
- `OldPassword`
- `Password`
- `TableSize`
- `TimeZone`
- `UseTimeZone`
- `UserType`

##### `ZoneAuth`

Created Fields:
- `Address`
- `AllowActiveDir`
- `AllowFixedRrsetOrder`
- `AllowGssTsigForUnderscoreZone`
- `AllowGssTsigZoneUpdates`
- `AllowQuery`
- `AllowTransfer`
- `AllowUpdate`
- `AllowUpdateForwarding`
- `AwsRte53ZoneInfo`
- `CloudInfo`
- `Comment`
- `CopyXferToNotify`
- `CreatePtrForBulkHosts`
- `CreatePtrForHosts`
- `CreateUnderscoreZones`
- `DdnsForceCreationTimestampUpdate`
- `DdnsPrincipalGroup`
- `DdnsPrincipalTracking`
- `DdnsRestrictPatterns`
- `DdnsRestrictPatternsList`
- `DdnsRestrictProtected`
- `DdnsRestrictSecure`
- `DdnsRestrictStatic`
- `Disable`
- `DisableForwarding`
- `DisplayDomain`
- `DnsFqdn`
- `DnsIntegrityEnable`
- `DnsIntegrityFrequency`
- `DnsIntegrityMember`
- `DnsIntegrityVerboseLogging`
- `DnsSoaEmail`
- `DnssecKeyParams`
- `DnssecKeys`
- `DnssecKskRolloverDate`
- `DnssecZskRolloverDate`
- `DoHostAbstraction`
- `EffectiveCheckNamesPolicy`
- `EffectiveRecordNamePolicy`
- `ExternalPrimaries`
- `ExternalSecondaries`
- `GridPrimary`
- `GridPrimarySharedWithMsParentDelegation`
- `GridSecondaries`
- `ImportFrom`
- `IsDnssecEnabled`
- `IsDnssecSigned`
- `IsMultimaster`
- `LastQueried`
- `LastQueriedAcl`
- `Locked`
- `LockedBy`
- `MaskPrefix`
- `MemberSoaMnames`
- `MemberSoaSerials`
- `MsAdIntegrated`
- `MsAllowTransfer`
- `MsAllowTransferMode`
- `MsDcNsRecordCreation`
- `MsDdnsMode`
- `MsManaged`
- `MsPrimaries`
- `MsReadOnly`
- `MsSecondaries`
- `MsSyncDisabled`
- `MsSyncMasterName`
- `NetworkAssociations`
- `NetworkView`
- `NotifyDelay`
- `NsGroup`
- `Parent`
- `Prefix`
- `PrimaryType`
- `RecordNamePolicy`
- `RecordsMonitored`
- `RestartIfNeeded`
- `RrNotQueriedEnabledTime`
- `ScavengingSettings`
- `SetSoaSerialNumber`
- `SoaDefaultTtl`
- `SoaEmail`
- `SoaExpire`
- `SoaNegativeTtl`
- `SoaRefresh`
- `SoaRetry`
- `SoaSerialNumber`
- `Srgs`
- `UpdateForwarding`
- `UseAllowActiveDir`
- `UseAllowQuery`
- `UseAllowTransfer`
- `UseAllowUpdate`
- `UseAllowUpdateForwarding`
- `UseCheckNamesPolicy`
- `UseCopyXferToNotify`
- `UseDdnsForceCreationTimestampUpdate`
- `UseDdnsPatternsRestriction`
- `UseDdnsPrincipalSecurity`
- `UseDdnsRestrictProtected`
- `UseDdnsRestrictStatic`
- `UseDnssecKeyParams`
- `UseExternalPrimary`
- `UseGridZoneTimer`
- `UseImportFrom`
- `UseNotifyDelay`
- `UseRecordNamePolicy`
- `UseScavengingSettings`
- `UseSoaEmail`
- `UsingSrgAssociations`
- `ZoneFormat`
- `ZoneNotQueriedEnabledTime`

##### `ZoneDelegated`

Created Fields:
- `Address`
- `Comment`
- `DelegatedTtl`
- `Disable`
- `DisplayDomain`
- `DnsFqdn`
- `EnableRfc2317Exclusion`
- `Locked`
- `LockedBy`
- `MaskPrefix`
- `MsAdIntegrated`
- `MsDdnsMode`
- `MsManaged`
- `MsReadOnly`
- `MsSyncMasterName`
- `NsGroup`
- `Parent`
- `Prefix`
- `UseDelegatedTtl`
- `UsingSrgAssociations`
- `ZoneFormat`

