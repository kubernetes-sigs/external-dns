package v2

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/exoscale/egoscale/v2/oapi"
)

type oapiClientMock struct {
	oapi.ClientWithResponsesInterface
	mock.Mock
}

func (m *oapiClientMock) AddExternalSourceToSecurityGroupWithResponse(
	ctx context.Context,
	id string,
	body oapi.AddExternalSourceToSecurityGroupJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.AddExternalSourceToSecurityGroupResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.AddExternalSourceToSecurityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) AddRuleToSecurityGroupWithResponse(
	ctx context.Context,
	id string,
	body oapi.AddRuleToSecurityGroupJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.AddRuleToSecurityGroupResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.AddRuleToSecurityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) AddServiceToLoadBalancerWithResponse(
	ctx context.Context,
	id string,
	body oapi.AddServiceToLoadBalancerJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.AddServiceToLoadBalancerResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.AddServiceToLoadBalancerResponse), args.Error(1)
}

func (m *oapiClientMock) AttachInstanceToElasticIpWithResponse( // nolint:revive
	ctx context.Context,
	id string,
	body oapi.AttachInstanceToElasticIpJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.AttachInstanceToElasticIpResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.AttachInstanceToElasticIpResponse), args.Error(1)
}

func (m *oapiClientMock) AttachInstanceToPrivateNetworkWithResponse(
	ctx context.Context,
	id string,
	body oapi.AttachInstanceToPrivateNetworkJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.AttachInstanceToPrivateNetworkResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.AttachInstanceToPrivateNetworkResponse), args.Error(1)
}

func (m *oapiClientMock) AttachInstanceToSecurityGroupWithResponse(
	ctx context.Context,
	id string,
	body oapi.AttachInstanceToSecurityGroupJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.AttachInstanceToSecurityGroupResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.AttachInstanceToSecurityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) CreateAntiAffinityGroupWithResponse(
	ctx context.Context,
	body oapi.CreateAntiAffinityGroupJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateAntiAffinityGroupResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateAntiAffinityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) CreateElasticIpWithResponse( // nolint:revive
	ctx context.Context,
	body oapi.CreateElasticIpJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateElasticIpResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateElasticIpResponse), args.Error(1)
}

func (m *oapiClientMock) CreateAccessKeyWithResponse(
	ctx context.Context,
	body oapi.CreateAccessKeyJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateAccessKeyResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateAccessKeyResponse), args.Error(1)
}

func (m *oapiClientMock) CreateInstanceWithResponse(
	ctx context.Context,
	body oapi.CreateInstanceJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateInstanceResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) CreateInstancePoolWithResponse(
	ctx context.Context,
	body oapi.CreateInstancePoolJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateInstancePoolResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateInstancePoolResponse), args.Error(1)
}

func (m *oapiClientMock) CreateLoadBalancerWithResponse(
	ctx context.Context,
	body oapi.CreateLoadBalancerJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateLoadBalancerResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateLoadBalancerResponse), args.Error(1)
}

func (m *oapiClientMock) CreatePrivateNetworkWithResponse(
	ctx context.Context,
	body oapi.CreatePrivateNetworkJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreatePrivateNetworkResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreatePrivateNetworkResponse), args.Error(1)
}

func (m *oapiClientMock) CreateSecurityGroupWithResponse(
	ctx context.Context,
	body oapi.CreateSecurityGroupJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateSecurityGroupResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateSecurityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) CreateSksClusterWithResponse(
	ctx context.Context,
	body oapi.CreateSksClusterJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateSksClusterResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateSksClusterResponse), args.Error(1)
}

func (m *oapiClientMock) CreateSksNodepoolWithResponse(
	ctx context.Context,
	id string,
	body oapi.CreateSksNodepoolJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateSksNodepoolResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.CreateSksNodepoolResponse), args.Error(1)
}

func (m *oapiClientMock) CreateSnapshotWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn) (*oapi.CreateSnapshotResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.CreateSnapshotResponse), args.Error(1)
}

func (m *oapiClientMock) CopyTemplateWithResponse(
	ctx context.Context,
	id string,
	body oapi.CopyTemplateJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CopyTemplateResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.CopyTemplateResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteAntiAffinityGroupWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteAntiAffinityGroupResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteAntiAffinityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteDbaasServiceWithResponse(
	ctx context.Context,
	name string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteDbaasServiceResponse, error) {
	args := m.Called(ctx, name, reqEditors)
	return args.Get(0).(*oapi.DeleteDbaasServiceResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteElasticIpWithResponse( // nolint:revive
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteElasticIpResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteElasticIpResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteInstanceWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteInstanceResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteInstancePoolWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteInstancePoolResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteInstancePoolResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteLoadBalancerWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteLoadBalancerResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteLoadBalancerResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteLoadBalancerServiceWithResponse(
	ctx context.Context,
	id string,
	serviceId string, // nolint:revive
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteLoadBalancerServiceResponse, error) {
	args := m.Called(ctx, id, serviceId, reqEditors)
	return args.Get(0).(*oapi.DeleteLoadBalancerServiceResponse), args.Error(1)
}

func (m *oapiClientMock) DeletePrivateNetworkWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeletePrivateNetworkResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeletePrivateNetworkResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteRuleFromSecurityGroupWithResponse(
	ctx context.Context,
	id string,
	ruleId string, // nolint:revive
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteRuleFromSecurityGroupResponse, error) {
	args := m.Called(ctx, id, ruleId, reqEditors)
	return args.Get(0).(*oapi.DeleteRuleFromSecurityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteSecurityGroupWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteSecurityGroupResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteSecurityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteSksClusterWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteSksClusterResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteSksClusterResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteSksNodepoolWithResponse(
	ctx context.Context,
	id string,
	sksNodepoolId string, // nolint:revive
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteSksNodepoolResponse, error) {
	args := m.Called(ctx, id, sksNodepoolId, reqEditors)
	return args.Get(0).(*oapi.DeleteSksNodepoolResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteSnapshotWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteSnapshotResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteSnapshotResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteSshKeyWithResponse( // nolint:revive
	ctx context.Context,
	name string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteSshKeyResponse, error) {
	args := m.Called(ctx, name, reqEditors)
	return args.Get(0).(*oapi.DeleteSshKeyResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteTemplateWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteTemplateResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteTemplateResponse), args.Error(1)
}

func (m *oapiClientMock) DetachInstanceFromElasticIpWithResponse( // nolint:revive
	ctx context.Context,
	id string,
	body oapi.DetachInstanceFromElasticIpJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DetachInstanceFromElasticIpResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.DetachInstanceFromElasticIpResponse), args.Error(1)
}

func (m *oapiClientMock) DetachInstanceFromPrivateNetworkWithResponse(
	ctx context.Context,
	id string,
	body oapi.DetachInstanceFromPrivateNetworkJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DetachInstanceFromPrivateNetworkResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.DetachInstanceFromPrivateNetworkResponse), args.Error(1)
}

func (m *oapiClientMock) DetachInstanceFromSecurityGroupWithResponse(
	ctx context.Context,
	id string,
	body oapi.DetachInstanceFromSecurityGroupJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DetachInstanceFromSecurityGroupResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.DetachInstanceFromSecurityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) EvictInstancePoolMembersWithResponse(
	ctx context.Context,
	id string,
	body oapi.EvictInstancePoolMembersJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.EvictInstancePoolMembersResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.EvictInstancePoolMembersResponse), args.Error(1)
}

func (m *oapiClientMock) EvictSksNodepoolMembersWithResponse(
	ctx context.Context,
	id string,
	sksNodepoolId string, // nolint:revive
	body oapi.EvictSksNodepoolMembersJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.EvictSksNodepoolMembersResponse, error) {
	args := m.Called(ctx, id, sksNodepoolId, body, reqEditors)
	return args.Get(0).(*oapi.EvictSksNodepoolMembersResponse), args.Error(1)
}

func (m *oapiClientMock) ExportSnapshotWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ExportSnapshotResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.ExportSnapshotResponse), args.Error(1)
}

func (m *oapiClientMock) GenerateSksClusterKubeconfigWithResponse(
	ctx context.Context,
	id string,
	body oapi.GenerateSksClusterKubeconfigJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GenerateSksClusterKubeconfigResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.GenerateSksClusterKubeconfigResponse), args.Error(1)
}

func (m *oapiClientMock) GetAntiAffinityGroupWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetAntiAffinityGroupResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetAntiAffinityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) GetDbaasCaCertificateWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetDbaasCaCertificateResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.GetDbaasCaCertificateResponse), args.Error(1)
}

func (m *oapiClientMock) GetDbaasServiceTypeWithResponse(
	ctx context.Context,
	name string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetDbaasServiceTypeResponse, error) {
	args := m.Called(ctx, name, reqEditors)
	return args.Get(0).(*oapi.GetDbaasServiceTypeResponse), args.Error(1)
}

func (m *oapiClientMock) GetDeployTargetWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetDeployTargetResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetDeployTargetResponse), args.Error(1)
}

func (m *oapiClientMock) GetElasticIpWithResponse( // nolint:revive
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetElasticIpResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetElasticIpResponse), args.Error(1)
}

func (m *oapiClientMock) GetAccessKeyWithResponse(
	ctx context.Context,
	key string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetAccessKeyResponse, error) {
	args := m.Called(ctx, key, reqEditors)
	return args.Get(0).(*oapi.GetAccessKeyResponse), args.Error(1)
}

func (m *oapiClientMock) GetInstanceWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetInstanceResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) GetInstancePoolWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetInstancePoolResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetInstancePoolResponse), args.Error(1)
}

func (m *oapiClientMock) GetInstanceTypeWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetInstanceTypeResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetInstanceTypeResponse), args.Error(1)
}

func (m *oapiClientMock) GetLoadBalancerWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetLoadBalancerResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetLoadBalancerResponse), args.Error(1)
}

func (m *oapiClientMock) GetPrivateNetworkWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetPrivateNetworkResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetPrivateNetworkResponse), args.Error(1)
}

func (m *oapiClientMock) GetQuotaWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetQuotaResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetQuotaResponse), args.Error(1)
}

func (m *oapiClientMock) GetSecurityGroupWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetSecurityGroupResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetSecurityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) GetSksClusterWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetSksClusterResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetSksClusterResponse), args.Error(1)
}

func (m *oapiClientMock) GetSksClusterAuthorityCertWithResponse(
	ctx context.Context,
	id string,
	authority oapi.GetSksClusterAuthorityCertParamsAuthority,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetSksClusterAuthorityCertResponse, error) {
	args := m.Called(ctx, id, authority, reqEditors)
	return args.Get(0).(*oapi.GetSksClusterAuthorityCertResponse), args.Error(1)
}

func (m *oapiClientMock) GetSksNodepoolWithResponse(
	ctx context.Context,
	id string,
	sksNodepoolId string, // nolint:revive
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetSksNodepoolResponse, error) {
	args := m.Called(ctx, id, sksNodepoolId, reqEditors)
	return args.Get(0).(*oapi.GetSksNodepoolResponse), args.Error(1)
}

func (m *oapiClientMock) GetSnapshotWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetSnapshotResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetSnapshotResponse), args.Error(1)
}

func (m *oapiClientMock) GetSshKeyWithResponse( // nolint:revive
	ctx context.Context,
	name string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetSshKeyResponse, error) {
	args := m.Called(ctx, name, reqEditors)
	return args.Get(0).(*oapi.GetSshKeyResponse), args.Error(1)
}

func (m *oapiClientMock) GetTemplateWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetTemplateResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetTemplateResponse), args.Error(1)
}

func (m *oapiClientMock) GetOperationWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetOperationResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetOperationResponse), args.Error(1)
}

func (m *oapiClientMock) ListAccessKeyKnownOperationsWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListAccessKeyKnownOperationsResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListAccessKeyKnownOperationsResponse), args.Error(1)
}

func (m *oapiClientMock) ListAccessKeyOperationsWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListAccessKeyOperationsResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListAccessKeyOperationsResponse), args.Error(1)
}

func (m *oapiClientMock) ListAccessKeysWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListAccessKeysResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListAccessKeysResponse), args.Error(1)
}

func (m *oapiClientMock) ListAntiAffinityGroupsWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListAntiAffinityGroupsResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListAntiAffinityGroupsResponse), args.Error(1)
}

func (m *oapiClientMock) ListDbaasServiceTypesWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListDbaasServiceTypesResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListDbaasServiceTypesResponse), args.Error(1)
}

func (m *oapiClientMock) ListDbaasServicesWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListDbaasServicesResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListDbaasServicesResponse), args.Error(1)
}

func (m *oapiClientMock) ListDeployTargetsWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListDeployTargetsResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListDeployTargetsResponse), args.Error(1)
}

func (m *oapiClientMock) ListElasticIpsWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListElasticIpsResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListElasticIpsResponse), args.Error(1)
}

func (m *oapiClientMock) ListInstancesWithResponse(
	ctx context.Context,
	params *oapi.ListInstancesParams,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListInstancesResponse, error) {
	args := m.Called(ctx, params, reqEditors)
	return args.Get(0).(*oapi.ListInstancesResponse), args.Error(1)
}

func (m *oapiClientMock) ListInstancePoolsWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListInstancePoolsResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListInstancePoolsResponse), args.Error(1)
}

func (m *oapiClientMock) ListInstanceTypesWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListInstanceTypesResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListInstanceTypesResponse), args.Error(1)
}

func (m *oapiClientMock) ListLoadBalancersWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListLoadBalancersResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListLoadBalancersResponse), args.Error(1)
}

func (m *oapiClientMock) ListPrivateNetworksWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListPrivateNetworksResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListPrivateNetworksResponse), args.Error(1)
}

func (m *oapiClientMock) ListQuotasWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListQuotasResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListQuotasResponse), args.Error(1)
}

func (m *oapiClientMock) ListSecurityGroupsWithResponse(
	ctx context.Context,
	params *oapi.ListSecurityGroupsParams,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListSecurityGroupsResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListSecurityGroupsResponse), args.Error(1)
}

func (m *oapiClientMock) ListSksClustersWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListSksClustersResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListSksClustersResponse), args.Error(1)
}

func (m *oapiClientMock) ListSksClusterVersionsWithResponse(
	ctx context.Context,
	params *oapi.ListSksClusterVersionsParams,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListSksClusterVersionsResponse, error) {
	args := m.Called(ctx, params, reqEditors)
	return args.Get(0).(*oapi.ListSksClusterVersionsResponse), args.Error(1)
}

func (m *oapiClientMock) ListSnapshotsWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListSnapshotsResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListSnapshotsResponse), args.Error(1)
}

func (m *oapiClientMock) ListSshKeysWithResponse( // nolint:revive
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListSshKeysResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListSshKeysResponse), args.Error(1)
}

func (m *oapiClientMock) ListTemplatesWithResponse(
	ctx context.Context,
	params *oapi.ListTemplatesParams,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListTemplatesResponse, error) {
	args := m.Called(ctx, params, reqEditors)
	return args.Get(0).(*oapi.ListTemplatesResponse), args.Error(1)
}

func (m *oapiClientMock) ListZonesWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListZonesResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListZonesResponse), args.Error(1)
}

func (m *oapiClientMock) RebootInstanceWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.RebootInstanceResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.RebootInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) RegisterSshKeyWithResponse( // nolint:revive
	ctx context.Context,
	body oapi.RegisterSshKeyJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.RegisterSshKeyResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.RegisterSshKeyResponse), args.Error(1)
}

func (m *oapiClientMock) RegisterTemplateWithResponse(
	ctx context.Context,
	body oapi.RegisterTemplateJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.RegisterTemplateResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.RegisterTemplateResponse), args.Error(1)
}

func (m *oapiClientMock) RemoveExternalSourceFromSecurityGroupWithResponse(
	ctx context.Context,
	id string,
	body oapi.RemoveExternalSourceFromSecurityGroupJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.RemoveExternalSourceFromSecurityGroupResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.RemoveExternalSourceFromSecurityGroupResponse), args.Error(1)
}

func (m *oapiClientMock) ResetInstanceWithResponse(
	ctx context.Context,
	id string,
	body oapi.ResetInstanceJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ResetInstanceResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.ResetInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) ResizeInstanceDiskWithResponse(
	ctx context.Context,
	id string,
	body oapi.ResizeInstanceDiskJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ResizeInstanceDiskResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.ResizeInstanceDiskResponse), args.Error(1)
}

func (m *oapiClientMock) RevertInstanceToSnapshotWithResponse(
	ctx context.Context,
	instanceId string, // nolint:revive
	body oapi.RevertInstanceToSnapshotJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.RevertInstanceToSnapshotResponse, error) {
	args := m.Called(ctx, instanceId, body, reqEditors)
	return args.Get(0).(*oapi.RevertInstanceToSnapshotResponse), args.Error(1)
}

func (m *oapiClientMock) RevokeAccessKeyWithResponse(
	ctx context.Context,
	key string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.RevokeAccessKeyResponse, error) {
	args := m.Called(ctx, key, reqEditors)
	return args.Get(0).(*oapi.RevokeAccessKeyResponse), args.Error(1)
}

func (m *oapiClientMock) RotateSksCcmCredentialsWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.RotateSksCcmCredentialsResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.RotateSksCcmCredentialsResponse), args.Error(1)
}

func (m *oapiClientMock) ScaleInstanceWithResponse(
	ctx context.Context,
	id string,
	body oapi.ScaleInstanceJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ScaleInstanceResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.ScaleInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) ScaleInstancePoolWithResponse(
	ctx context.Context,
	id string,
	body oapi.ScaleInstancePoolJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ScaleInstancePoolResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.ScaleInstancePoolResponse), args.Error(1)
}

func (m *oapiClientMock) ScaleSksNodepoolWithResponse(
	ctx context.Context,
	id string,
	sksNodepoolId string, // nolint:revive
	body oapi.ScaleSksNodepoolJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ScaleSksNodepoolResponse, error) {
	args := m.Called(ctx, id, sksNodepoolId, body, reqEditors)
	return args.Get(0).(*oapi.ScaleSksNodepoolResponse), args.Error(1)
}

func (m *oapiClientMock) StartInstanceWithResponse(
	ctx context.Context,
	id string,
	body oapi.StartInstanceJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.StartInstanceResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.StartInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) StopInstanceWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.StopInstanceResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.StopInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateElasticIpWithResponse( // nolint:revive
	ctx context.Context,
	id string,
	body oapi.UpdateElasticIpJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateElasticIpResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdateElasticIpResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateInstanceWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpdateInstanceJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateInstanceResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdateInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateInstancePoolWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpdateInstancePoolJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateInstancePoolResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdateInstancePoolResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateLoadBalancerWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpdateLoadBalancerJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateLoadBalancerResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdateLoadBalancerResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateLoadBalancerServiceWithResponse(
	ctx context.Context,
	id string,
	serviceId string, // nolint:revive
	body oapi.UpdateLoadBalancerServiceJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateLoadBalancerServiceResponse, error) {
	args := m.Called(ctx, id, serviceId, body, reqEditors)
	return args.Get(0).(*oapi.UpdateLoadBalancerServiceResponse), args.Error(1)
}

func (m *oapiClientMock) UpdatePrivateNetworkWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpdatePrivateNetworkJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdatePrivateNetworkResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdatePrivateNetworkResponse), args.Error(1)
}

func (m *oapiClientMock) UpdatePrivateNetworkInstanceIpWithResponse( // nolint:revive
	ctx context.Context,
	id string,
	body oapi.UpdatePrivateNetworkInstanceIpJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdatePrivateNetworkInstanceIpResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdatePrivateNetworkInstanceIpResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateSksClusterWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpdateSksClusterJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateSksClusterResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdateSksClusterResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateSksNodepoolWithResponse(
	ctx context.Context,
	id string,
	sksNodepoolId string, // nolint:revive
	body oapi.UpdateSksNodepoolJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateSksNodepoolResponse, error) {
	args := m.Called(ctx, id, sksNodepoolId, body, reqEditors)
	return args.Get(0).(*oapi.UpdateSksNodepoolResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateTemplateWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpdateTemplateJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateTemplateResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdateTemplateResponse), args.Error(1)
}

func (m *oapiClientMock) UpgradeSksClusterWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpgradeSksClusterJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpgradeSksClusterResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpgradeSksClusterResponse), args.Error(1)
}

func (m *oapiClientMock) UpgradeSksClusterServiceLevelWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpgradeSksClusterServiceLevelResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.UpgradeSksClusterServiceLevelResponse), args.Error(1)
}

func (m *oapiClientMock) GetDbaasMigrationStatusWithResponse(
	ctx context.Context,
	name oapi.DbaasServiceName,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetDbaasMigrationStatusResponse, error) {
	args := m.Called(ctx, name, reqEditors)
	return args.Get(0).(*oapi.GetDbaasMigrationStatusResponse), args.Error(1)
}

func (m *oapiClientMock) ListDnsDomainsWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListDnsDomainsResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListDnsDomainsResponse), args.Error(1)
}

func (m *oapiClientMock) GetDnsDomainWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetDnsDomainResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetDnsDomainResponse), args.Error(1)
}

func (m *oapiClientMock) CreateDnsDomainWithResponse(
	ctx context.Context,
	body oapi.CreateDnsDomainJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateDnsDomainResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateDnsDomainResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteDnsDomainWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteDnsDomainResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteDnsDomainResponse), args.Error(1)
}

func (m *oapiClientMock) GetDnsDomainZoneFileWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetDnsDomainZoneFileResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetDnsDomainZoneFileResponse), args.Error(1)
}

func (m *oapiClientMock) ListDnsDomainRecordsWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListDnsDomainRecordsResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.ListDnsDomainRecordsResponse), args.Error(1)
}

func (m *oapiClientMock) GetDnsDomainRecordWithResponse(
	ctx context.Context,
	domainId string,
	recordId string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetDnsDomainRecordResponse, error) {
	args := m.Called(ctx, domainId, recordId, reqEditors)
	return args.Get(0).(*oapi.GetDnsDomainRecordResponse), args.Error(1)
}

func (m *oapiClientMock) CreateDnsDomainRecordWithResponse(
	ctx context.Context,
	domainId string,
	body oapi.CreateDnsDomainRecordJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateDnsDomainRecordResponse, error) {
	args := m.Called(ctx, domainId, body, reqEditors)
	return args.Get(0).(*oapi.CreateDnsDomainRecordResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteDnsDomainRecordWithResponse(
	ctx context.Context,
	domainId string,
	recordId string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteDnsDomainRecordResponse, error) {
	args := m.Called(ctx, domainId, recordId, reqEditors)
	return args.Get(0).(*oapi.DeleteDnsDomainRecordResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateDnsDomainRecordWithResponse(
	ctx context.Context,
	domainId string,
	recordId string,
	body oapi.UpdateDnsDomainRecordJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateDnsDomainRecordResponse, error) {
	args := m.Called(ctx, domainId, recordId, body, reqEditors)
	return args.Get(0).(*oapi.UpdateDnsDomainRecordResponse), args.Error(1)
}

func (m *oapiClientMock) GetReverseDnsInstanceWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetReverseDnsInstanceResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetReverseDnsInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteReverseDnsInstanceWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteReverseDnsInstanceResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteReverseDnsInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateReverseDnsInstanceWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpdateReverseDnsInstanceJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateReverseDnsInstanceResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdateReverseDnsInstanceResponse), args.Error(1)
}

func (m *oapiClientMock) GetReverseDnsElasticIpWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetReverseDnsElasticIpResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetReverseDnsElasticIpResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteReverseDnsElasticIpWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteReverseDnsElasticIpResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteReverseDnsElasticIpResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateReverseDnsElasticIpWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpdateReverseDnsElasticIpJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateReverseDnsElasticIpResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdateReverseDnsElasticIpResponse), args.Error(1)
}

func (m *oapiClientMock) StopDbaasRedisMigrationWithResponse(
	ctx context.Context,
	name oapi.DbaasServiceName,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.StopDbaasRedisMigrationResponse, error) {
	args := m.Called(ctx, name, reqEditors)
	return args.Get(0).(*oapi.StopDbaasRedisMigrationResponse), args.Error(1)
}

func (m *oapiClientMock) StopDbaasPgMigrationWithResponse(
	ctx context.Context,
	name oapi.DbaasServiceName,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.StopDbaasPgMigrationResponse, error) {
	args := m.Called(ctx, name, reqEditors)
	return args.Get(0).(*oapi.StopDbaasPgMigrationResponse), args.Error(1)
}

func (m *oapiClientMock) StopDbaasMysqlMigrationWithResponse(
	ctx context.Context,
	name oapi.DbaasServiceName,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.StopDbaasMysqlMigrationResponse, error) {
	args := m.Called(ctx, name, reqEditors)
	return args.Get(0).(*oapi.StopDbaasMysqlMigrationResponse), args.Error(1)
}

func (m *oapiClientMock) GetIamOrganizationPolicyWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetIamOrganizationPolicyResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.GetIamOrganizationPolicyResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateIamOrganizationPolicyWithResponse(
	ctx context.Context,
	body oapi.UpdateIamOrganizationPolicyJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateIamOrganizationPolicyResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.UpdateIamOrganizationPolicyResponse), args.Error(1)
}

func (m *oapiClientMock) GetIamRoleWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetIamRoleResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetIamRoleResponse), args.Error(1)
}

func (m *oapiClientMock) ListIamRolesWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListIamRolesResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListIamRolesResponse), args.Error(1)
}

func (m *oapiClientMock) CreateIamRoleWithResponse(
	ctx context.Context,
	body oapi.CreateIamRoleJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateIamRoleResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateIamRoleResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteIamRoleWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteIamRoleResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteIamRoleResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateIamRoleWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpdateIamRoleJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateIamRoleResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdateIamRoleResponse), args.Error(1)
}

func (m *oapiClientMock) UpdateIamRolePolicyWithResponse(
	ctx context.Context,
	id string,
	body oapi.UpdateIamRolePolicyJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.UpdateIamRolePolicyResponse, error) {
	args := m.Called(ctx, id, body, reqEditors)
	return args.Get(0).(*oapi.UpdateIamRolePolicyResponse), args.Error(1)
}

func (m *oapiClientMock) GetApiKeyWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.GetApiKeyResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.GetApiKeyResponse), args.Error(1)
}

func (m *oapiClientMock) ListApiKeysWithResponse(
	ctx context.Context,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.ListApiKeysResponse, error) {
	args := m.Called(ctx, reqEditors)
	return args.Get(0).(*oapi.ListApiKeysResponse), args.Error(1)
}

func (m *oapiClientMock) CreateApiKeyWithResponse(
	ctx context.Context,
	body oapi.CreateApiKeyJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CreateApiKeyResponse, error) {
	args := m.Called(ctx, body, reqEditors)
	return args.Get(0).(*oapi.CreateApiKeyResponse), args.Error(1)
}

func (m *oapiClientMock) DeleteApiKeyWithResponse(
	ctx context.Context,
	id string,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.DeleteApiKeyResponse, error) {
	args := m.Called(ctx, id, reqEditors)
	return args.Get(0).(*oapi.DeleteApiKeyResponse), args.Error(1)
}
