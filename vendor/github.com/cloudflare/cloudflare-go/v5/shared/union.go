// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsAuditLogListParamsBeforeUnion() {}
func (UnionTime) ImplementsAuditLogListParamsSinceUnion()  {}

type UnionString string

func (UnionString) ImplementsReceivedGetParamsEndUnion()   {}
func (UnionString) ImplementsReceivedGetParamsStartUnion() {}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersFiltersValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion() {}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion() {}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapUnionItem() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapUnionItem() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion() {}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapUnionItem() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapUnionItem() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArrayUnionItem() {
}
func (UnionString) ImplementsObservabilityTelemetryQueryResponsePatternsSeriesDataGroupsValueUnion() {
}
func (UnionString) ImplementsObservabilityTelemetryValuesResponseValueUnion()                      {}
func (UnionString) ImplementsObservabilityTelemetryKeysParamsFiltersValueUnion()                   {}
func (UnionString) ImplementsObservabilityTelemetryKeysParamsKeyNeedleValueUnion()                 {}
func (UnionString) ImplementsObservabilityTelemetryKeysParamsNeedleValueUnion()                    {}
func (UnionString) ImplementsObservabilityTelemetryQueryParamsParametersFiltersValueUnion()        {}
func (UnionString) ImplementsObservabilityTelemetryQueryParamsParametersNeedleValueUnion()         {}
func (UnionString) ImplementsObservabilityTelemetryValuesParamsFiltersValueUnion()                 {}
func (UnionString) ImplementsObservabilityTelemetryValuesParamsNeedleValueUnion()                  {}
func (UnionString) ImplementsNamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion()           {}
func (UnionString) ImplementsNamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion()        {}
func (UnionString) ImplementsOriginPortUnionParam()                                                {}
func (UnionString) ImplementsOriginPortUnion()                                                     {}
func (UnionString) ImplementsHealthCheckTargetUnionParam()                                         {}
func (UnionString) ImplementsHealthCheckTargetUnion()                                              {}
func (UnionString) ImplementsGRETunnelNewResponseHealthCheckTargetUnion()                          {}
func (UnionString) ImplementsGRETunnelUpdateResponseModifiedGRETunnelHealthCheckTargetUnion()      {}
func (UnionString) ImplementsGRETunnelListResponseGRETunnelsHealthCheckTargetUnion()               {}
func (UnionString) ImplementsGRETunnelDeleteResponseDeletedGRETunnelHealthCheckTargetUnion()       {}
func (UnionString) ImplementsGRETunnelBulkUpdateResponseModifiedGRETunnelsHealthCheckTargetUnion() {}
func (UnionString) ImplementsGRETunnelGetResponseGRETunnelHealthCheckTargetUnion()                 {}
func (UnionString) ImplementsGRETunnelNewParamsHealthCheckTargetUnion()                            {}
func (UnionString) ImplementsGRETunnelUpdateParamsHealthCheckTargetUnion()                         {}
func (UnionString) ImplementsIPSECTunnelNewResponseHealthCheckTargetUnion()                        {}
func (UnionString) ImplementsIPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetUnion()  {}
func (UnionString) ImplementsIPSECTunnelListResponseIPSECTunnelsHealthCheckTargetUnion()           {}
func (UnionString) ImplementsIPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetUnion()   {}
func (UnionString) ImplementsIPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetUnion() {
}
func (UnionString) ImplementsIPSECTunnelGetResponseIPSECTunnelHealthCheckTargetUnion()         {}
func (UnionString) ImplementsIPSECTunnelNewParamsHealthCheckTargetUnion()                      {}
func (UnionString) ImplementsIPSECTunnelUpdateParamsHealthCheckTargetUnion()                   {}
func (UnionString) ImplementsAccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationNewParamsBodySaaSApplicationPolicyUnion()        {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyBrowserSSHApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyBrowserVNCApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyAppLauncherApplicationPolicyUnion() {}
func (UnionString) ImplementsAccessApplicationNewParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationNewParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationNewParamsBodyBrowserRdpApplicationPolicyUnion()     {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodySaaSApplicationPolicyUnion()        {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyBrowserSSHApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyBrowserVNCApplicationPolicyUnion()  {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyAppLauncherApplicationPolicyUnion() {}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyDeviceEnrollmentPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyBrowserIsolationPermissionsApplicationPolicyUnion() {
}
func (UnionString) ImplementsAccessApplicationUpdateParamsBodyBrowserRdpApplicationPolicyUnion() {}
func (UnionString) ImplementsAccessApplicationPolicyTestNewParamsPolicyUnion()                   {}
func (UnionString) ImplementsDLPEmailRuleNewResponseConditionsValueUnion()                       {}
func (UnionString) ImplementsDLPEmailRuleUpdateResponseConditionsValueUnion()                    {}
func (UnionString) ImplementsDLPEmailRuleListResponseConditionsValueUnion()                      {}
func (UnionString) ImplementsDLPEmailRuleDeleteResponseConditionsValueUnion()                    {}
func (UnionString) ImplementsDLPEmailRuleBulkEditResponseConditionsValueUnion()                  {}
func (UnionString) ImplementsDLPEmailRuleGetResponseConditionsValueUnion()                       {}
func (UnionString) ImplementsDLPEmailRuleNewParamsConditionsValueUnion()                         {}
func (UnionString) ImplementsDLPEmailRuleUpdateParamsConditionsValueUnion()                      {}
func (UnionString) ImplementsRankingTimeseriesGroupsResponseSerie0Union()                        {}
func (UnionString) ImplementsRankingInternetServiceTimeseriesGroupsResponseSerie0Union()         {}
func (UnionString) ImplementsConfigurationToolsZarazManagedComponentDefaultFieldsUnion()         {}
func (UnionString) ImplementsConfigurationToolsZarazManagedComponentSettingsUnion()              {}
func (UnionString) ImplementsConfigurationToolsWorkerDefaultFieldsUnion()                        {}
func (UnionString) ImplementsConfigurationToolsWorkerSettingsUnion()                             {}
func (UnionString) ImplementsConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion()    {}
func (UnionString) ImplementsConfigUpdateParamsToolsZarazManagedComponentSettingsUnion()         {}
func (UnionString) ImplementsConfigUpdateParamsToolsWorkerDefaultFieldsUnion()                   {}
func (UnionString) ImplementsConfigUpdateParamsToolsWorkerSettingsUnion()                        {}
func (UnionString) ImplementsSettingValueUnionParam()                                            {}
func (UnionString) ImplementsSettingValueUnion()                                                 {}
func (UnionString) ImplementsThreatEventListParamsSearchValueUnion()                             {}
func (UnionString) ImplementsThreatEventListParamsSearchValueArrayItemUnion()                    {}
func (UnionString) ImplementsLogListParamsFiltersValueUnion()                                    {}
func (UnionString) ImplementsLogDeleteParamsFiltersValueUnion()                                  {}
func (UnionString) ImplementsLogEditParamsMetadataUnion()                                        {}
func (UnionString) ImplementsDatasetNewResponseFiltersValueUnion()                               {}
func (UnionString) ImplementsDatasetUpdateResponseFiltersValueUnion()                            {}
func (UnionString) ImplementsDatasetListResponseFiltersValueUnion()                              {}
func (UnionString) ImplementsDatasetDeleteResponseFiltersValueUnion()                            {}
func (UnionString) ImplementsDatasetGetResponseFiltersValueUnion()                               {}
func (UnionString) ImplementsDatasetNewParamsFiltersValueUnion()                                 {}
func (UnionString) ImplementsDatasetUpdateParamsFiltersValueUnion()                              {}
func (UnionString) ImplementsEvaluationNewResponseDatasetsFiltersValueUnion()                    {}
func (UnionString) ImplementsEvaluationListResponseDatasetsFiltersValueUnion()                   {}
func (UnionString) ImplementsEvaluationDeleteResponseDatasetsFiltersValueUnion()                 {}
func (UnionString) ImplementsEvaluationGetResponseDatasetsFiltersValueUnion()                    {}
func (UnionString) ImplementsInstanceGetResponseOutputUnion()                                    {}
func (UnionString) ImplementsAIRunResponseUnion()                                                {}
func (UnionString) ImplementsAIRunParamsBodyTextEmbeddingsTextUnion()                            {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsHeightUnion()                                 {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsMarginBottomUnion()                           {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsMarginLeftUnion()                             {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsMarginRightUnion()                            {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsMarginTopUnion()                              {}
func (UnionString) ImplementsPDFNewParamsPDFOptionsWidthUnion()                                  {}

type UnionBool bool

func (UnionBool) ImplementsScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion()              {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersFiltersValueUnion() {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion()  {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion()   {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion()   {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventUnion() {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapUnionItem() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapUnionItem() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion()        {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventUnion() {}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapUnionItem() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapUnionItem() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArrayUnionItem() {
}
func (UnionBool) ImplementsObservabilityTelemetryQueryResponsePatternsSeriesDataGroupsValueUnion() {}
func (UnionBool) ImplementsObservabilityTelemetryValuesResponseValueUnion()                        {}
func (UnionBool) ImplementsObservabilityTelemetryKeysParamsFiltersValueUnion()                     {}
func (UnionBool) ImplementsObservabilityTelemetryKeysParamsKeyNeedleValueUnion()                   {}
func (UnionBool) ImplementsObservabilityTelemetryKeysParamsNeedleValueUnion()                      {}
func (UnionBool) ImplementsObservabilityTelemetryQueryParamsParametersFiltersValueUnion()          {}
func (UnionBool) ImplementsObservabilityTelemetryQueryParamsParametersNeedleValueUnion()           {}
func (UnionBool) ImplementsObservabilityTelemetryValuesParamsFiltersValueUnion()                   {}
func (UnionBool) ImplementsObservabilityTelemetryValuesParamsNeedleValueUnion()                    {}
func (UnionBool) ImplementsNamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion()             {}
func (UnionBool) ImplementsNamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion()          {}
func (UnionBool) ImplementsDispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion() {
}
func (UnionBool) ImplementsConfigurationToolsZarazManagedComponentDefaultFieldsUnion()      {}
func (UnionBool) ImplementsConfigurationToolsZarazManagedComponentSettingsUnion()           {}
func (UnionBool) ImplementsConfigurationToolsWorkerDefaultFieldsUnion()                     {}
func (UnionBool) ImplementsConfigurationToolsWorkerSettingsUnion()                          {}
func (UnionBool) ImplementsConfigUpdateParamsToolsZarazManagedComponentDefaultFieldsUnion() {}
func (UnionBool) ImplementsConfigUpdateParamsToolsZarazManagedComponentSettingsUnion()      {}
func (UnionBool) ImplementsConfigUpdateParamsToolsWorkerDefaultFieldsUnion()                {}
func (UnionBool) ImplementsConfigUpdateParamsToolsWorkerSettingsUnion()                     {}
func (UnionBool) ImplementsLogListParamsFiltersValueUnion()                                 {}
func (UnionBool) ImplementsLogDeleteParamsFiltersValueUnion()                               {}
func (UnionBool) ImplementsLogEditParamsMetadataUnion()                                     {}
func (UnionBool) ImplementsDatasetNewResponseFiltersValueUnion()                            {}
func (UnionBool) ImplementsDatasetUpdateResponseFiltersValueUnion()                         {}
func (UnionBool) ImplementsDatasetListResponseFiltersValueUnion()                           {}
func (UnionBool) ImplementsDatasetDeleteResponseFiltersValueUnion()                         {}
func (UnionBool) ImplementsDatasetGetResponseFiltersValueUnion()                            {}
func (UnionBool) ImplementsDatasetNewParamsFiltersValueUnion()                              {}
func (UnionBool) ImplementsDatasetUpdateParamsFiltersValueUnion()                           {}
func (UnionBool) ImplementsEvaluationNewResponseDatasetsFiltersValueUnion()                 {}
func (UnionBool) ImplementsEvaluationListResponseDatasetsFiltersValueUnion()                {}
func (UnionBool) ImplementsEvaluationDeleteResponseDatasetsFiltersValueUnion()              {}
func (UnionBool) ImplementsEvaluationGetResponseDatasetsFiltersValueUnion()                 {}

type UnionInt int64

func (UnionInt) ImplementsReceivedGetParamsEndUnion()                             {}
func (UnionInt) ImplementsReceivedGetParamsStartUnion()                           {}
func (UnionInt) ImplementsPageRuleActionsCacheTTLByStatusValueUnion()             {}
func (UnionInt) ImplementsPageRuleNewParamsActionsCacheTTLByStatusValueUnion()    {}
func (UnionInt) ImplementsPageRuleUpdateParamsActionsCacheTTLByStatusValueUnion() {}
func (UnionInt) ImplementsPageRuleEditParamsActionsCacheTTLByStatusValueUnion()   {}
func (UnionInt) ImplementsOriginPortUnionParam()                                  {}
func (UnionInt) ImplementsOriginPortUnion()                                       {}

type UnionFloat float64

func (UnionFloat) ImplementsSettingEditParamsBodyValueValueUnion() {}
func (UnionFloat) ImplementsTTLParam()                             {}
func (UnionFloat) ImplementsTTL()                                  {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersFiltersValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion() {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion() {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion() {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventUnion() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapUnionItem() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapUnionItem() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion()        {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventUnion() {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapUnionItem() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapUnionItem() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArrayUnionItem() {
}
func (UnionFloat) ImplementsObservabilityTelemetryQueryResponsePatternsSeriesDataGroupsValueUnion() {}
func (UnionFloat) ImplementsObservabilityTelemetryValuesResponseValueUnion()                        {}
func (UnionFloat) ImplementsObservabilityTelemetryKeysParamsFiltersValueUnion()                     {}
func (UnionFloat) ImplementsObservabilityTelemetryKeysParamsKeyNeedleValueUnion()                   {}
func (UnionFloat) ImplementsObservabilityTelemetryKeysParamsNeedleValueUnion()                      {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryParamsParametersFiltersValueUnion()          {}
func (UnionFloat) ImplementsObservabilityTelemetryQueryParamsParametersNeedleValueUnion()           {}
func (UnionFloat) ImplementsObservabilityTelemetryValuesParamsFiltersValueUnion()                   {}
func (UnionFloat) ImplementsObservabilityTelemetryValuesParamsNeedleValueUnion()                    {}
func (UnionFloat) ImplementsNamespaceBulkGetResponseWorkersKVBulkGetResultValuesUnion()             {}
func (UnionFloat) ImplementsNamespaceKeyBulkGetResponseWorkersKVBulkGetResultValuesUnion()          {}
func (UnionFloat) ImplementsRankingTimeseriesGroupsResponseSerie0Union()                            {}
func (UnionFloat) ImplementsRankingInternetServiceTimeseriesGroupsResponseSerie0Union()             {}
func (UnionFloat) ImplementsSettingValueUnionParam()                                                {}
func (UnionFloat) ImplementsSettingValueUnion()                                                     {}
func (UnionFloat) ImplementsThreatEventListParamsSearchValueUnion()                                 {}
func (UnionFloat) ImplementsThreatEventListParamsSearchValueArrayItemUnion()                        {}
func (UnionFloat) ImplementsLogListParamsFiltersValueUnion()                                        {}
func (UnionFloat) ImplementsLogDeleteParamsFiltersValueUnion()                                      {}
func (UnionFloat) ImplementsLogEditParamsMetadataUnion()                                            {}
func (UnionFloat) ImplementsDatasetNewResponseFiltersValueUnion()                                   {}
func (UnionFloat) ImplementsDatasetUpdateResponseFiltersValueUnion()                                {}
func (UnionFloat) ImplementsDatasetListResponseFiltersValueUnion()                                  {}
func (UnionFloat) ImplementsDatasetDeleteResponseFiltersValueUnion()                                {}
func (UnionFloat) ImplementsDatasetGetResponseFiltersValueUnion()                                   {}
func (UnionFloat) ImplementsDatasetNewParamsFiltersValueUnion()                                     {}
func (UnionFloat) ImplementsDatasetUpdateParamsFiltersValueUnion()                                  {}
func (UnionFloat) ImplementsEvaluationNewResponseDatasetsFiltersValueUnion()                        {}
func (UnionFloat) ImplementsEvaluationListResponseDatasetsFiltersValueUnion()                       {}
func (UnionFloat) ImplementsEvaluationDeleteResponseDatasetsFiltersValueUnion()                     {}
func (UnionFloat) ImplementsEvaluationGetResponseDatasetsFiltersValueUnion()                        {}
func (UnionFloat) ImplementsInstanceGetResponseOutputUnion()                                        {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsHeightUnion()                                     {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsMarginBottomUnion()                               {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsMarginLeftUnion()                                 {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsMarginRightUnion()                                {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsMarginTopUnion()                                  {}
func (UnionFloat) ImplementsPDFNewParamsPDFOptionsWidthUnion()                                      {}
