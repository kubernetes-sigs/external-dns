// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20201028

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// token校验失败。
	AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 私有域关联VPC失败。
	FAILEDOPERATION_BINDZONEVPCFAILED = "FailedOperation.BindZoneVpcFailed"

	// 记录创建失败。
	FAILEDOPERATION_CREATERECORDFAILED = "FailedOperation.CreateRecordFailed"

	// 私有域创建失败。
	FAILEDOPERATION_CREATEZONEFAILED = "FailedOperation.CreateZoneFailed"

	// 当前私有域已关联 VPC，如需清空解析记录请先解除 VPC 关联。
	FAILEDOPERATION_DELETELASTBINDVPCRECORDFAILED = "FailedOperation.DeleteLastBindVpcRecordFailed"

	// 解析域删除失败。
	FAILEDOPERATION_DELETEZONEFAILED = "FailedOperation.DeleteZoneFailed"

	// 记录修改失败。
	FAILEDOPERATION_MODIFYRECORDFAILED = "FailedOperation.ModifyRecordFailed"

	// 私有域修改失败。
	FAILEDOPERATION_MODIFYZONEFAILED = "FailedOperation.ModifyZoneFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 错误未定义。
	INTERNALERROR_UNDEFIENDERROR = "InternalError.UndefiendError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 已经存在绑定的账号。
	INVALIDPARAMETER_ACCOUNTEXIST = "InvalidParameter.AccountExist"

	// 非法CIDR。
	INVALIDPARAMETER_ILLEGALCIDR = "InvalidParameter.IllegalCidr"

	// 域名不正确。
	INVALIDPARAMETER_ILLEGALDOMAIN = "InvalidParameter.IllegalDomain"

	// 顶级域名不正确。
	INVALIDPARAMETER_ILLEGALDOMAINTLD = "InvalidParameter.IllegalDomainTld"

	// PTR记录非法。
	INVALIDPARAMETER_ILLEGALPTRRECORD = "InvalidParameter.IllegalPTRRecord"

	// 记录不合法。
	INVALIDPARAMETER_ILLEGALRECORD = "InvalidParameter.IllegalRecord"

	// 无效的记录值。
	INVALIDPARAMETER_ILLEGALRECORDVALUE = "InvalidParameter.IllegalRecordValue"

	// VPC非法。
	INVALIDPARAMETER_ILLEGALVPCINFO = "InvalidParameter.IllegalVpcInfo"

	// MX 必须为5-50之间且为5的倍数。
	INVALIDPARAMETER_INVALIDMX = "InvalidParameter.InvalidMX"

	// AAAA记录负载均衡数量超过50。
	INVALIDPARAMETER_RECORDAAAACOUNTEXCEED = "InvalidParameter.RecordAAAACountExceed"

	// A记录负载均衡数量超过50。
	INVALIDPARAMETER_RECORDACOUNTEXCEED = "InvalidParameter.RecordACountExceed"

	// CNAME记录负载均衡数量超过50。
	INVALIDPARAMETER_RECORDCNAMECOUNTEXCEED = "InvalidParameter.RecordCNAMECountExceed"

	// 记录冲突。
	INVALIDPARAMETER_RECORDCONFLICT = "InvalidParameter.RecordConflict"

	// 记录数量超过限制。
	INVALIDPARAMETER_RECORDCOUNTEXCEED = "InvalidParameter.RecordCountExceed"

	// 记录已经存在。
	INVALIDPARAMETER_RECORDEXIST = "InvalidParameter.RecordExist"

	// 记录层级超过限制。
	INVALIDPARAMETER_RECORDLEVELEXCEED = "InvalidParameter.RecordLevelExceed"

	// MX记录负载均衡数量超过50。
	INVALIDPARAMETER_RECORDMXCOUNTEXCEED = "InvalidParameter.RecordMXCountExceed"

	// 记录不存在。
	INVALIDPARAMETER_RECORDNOTEXIST = "InvalidParameter.RecordNotExist"

	// 记录负载均衡数量超过限制。
	INVALIDPARAMETER_RECORDROLLLIMITCOUNTEXCEED = "InvalidParameter.RecordRolllimitCountExceed"

	// TXT记录负载均衡数量超过10。
	INVALIDPARAMETER_RECORDTXTCOUNTEXCEED = "InvalidParameter.RecordTXTCountExceed"

	// 当前记录类型不支持权重。
	INVALIDPARAMETER_RECORDUNSUPPORTWEIGHT = "InvalidParameter.RecordUnsupportWeight"

	// VPC已绑定其它解析域。
	INVALIDPARAMETER_VPCBINDED = "InvalidParameter.VpcBinded"

	// 当前VPC已关联相同主域名。
	INVALIDPARAMETER_VPCBINDEDMAINDOMAIN = "InvalidParameter.VpcBindedMainDomain"

	// VPC关联反解析域超过限制。
	INVALIDPARAMETER_VPCPTRZONEBINDEXCEED = "InvalidParameter.VpcPtrZoneBindExceed"

	// 解析域不存在。
	INVALIDPARAMETER_ZONENOTEXISTS = "InvalidParameter.ZoneNotExists"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 私有域解析服务未开通。
	RESOURCENOTFOUND_SERVICENOTSUBSCRIBED = "ResourceNotFound.ServiceNotSubscribed"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 角色未授权。
	UNAUTHORIZEDOPERATION_ROLEUNAUTHORIZED = "UnauthorizedOperation.RoleUnAuthorized"

	// 未实名账号。
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDACCOUNT = "UnauthorizedOperation.UnauthorizedAccount"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 账号未绑定。
	UNSUPPORTEDOPERATION_ACCOUNTNOTBOUND = "UnsupportedOperation.AccountNotBound"

	// 存在绑定的VPC资源。
	UNSUPPORTEDOPERATION_EXISTBOUNDVPC = "UnsupportedOperation.ExistBoundVpc"
)
