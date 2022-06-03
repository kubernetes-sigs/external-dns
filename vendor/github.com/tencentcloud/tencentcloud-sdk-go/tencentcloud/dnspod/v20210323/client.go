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

package v20210323

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-03-23"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateDomainRequest() (request *CreateDomainRequest) {
    request = &CreateDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDomain")
    
    
    return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
    response = &CreateDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDomain
// 添加域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINEXISTS = "FailedOperation.DomainExists"
//  FAILEDOPERATION_DOMAINOWNEDBYOTHERUSER = "FailedOperation.DomainOwnedByOtherUser"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    
    response = NewCreateDomainResponse()
    err = c.Send(request, response)
    return
}

// CreateDomain
// 添加域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINEXISTS = "FailedOperation.DomainExists"
//  FAILEDOPERATION_DOMAINOWNEDBYOTHERUSER = "FailedOperation.DomainOwnedByOtherUser"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
func (c *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainAliasRequest() (request *CreateDomainAliasRequest) {
    request = &CreateDomainAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDomainAlias")
    
    
    return
}

func NewCreateDomainAliasResponse() (response *CreateDomainAliasResponse) {
    response = &CreateDomainAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDomainAlias
// 创建域名别名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINEXISTS = "FailedOperation.DomainExists"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINOWNEDBYOTHERUSER = "FailedOperation.DomainOwnedByOtherUser"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  LIMITEXCEEDED_DOMAINALIASCOUNTEXCEEDED = "LimitExceeded.DomainAliasCountExceeded"
//  LIMITEXCEEDED_DOMAINALIASNUMBERLIMIT = "LimitExceeded.DomainAliasNumberLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDomainAlias(request *CreateDomainAliasRequest) (response *CreateDomainAliasResponse, err error) {
    if request == nil {
        request = NewCreateDomainAliasRequest()
    }
    
    response = NewCreateDomainAliasResponse()
    err = c.Send(request, response)
    return
}

// CreateDomainAlias
// 创建域名别名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINEXISTS = "FailedOperation.DomainExists"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINOWNEDBYOTHERUSER = "FailedOperation.DomainOwnedByOtherUser"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  LIMITEXCEEDED_DOMAINALIASCOUNTEXCEEDED = "LimitExceeded.DomainAliasCountExceeded"
//  LIMITEXCEEDED_DOMAINALIASNUMBERLIMIT = "LimitExceeded.DomainAliasNumberLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDomainAliasWithContext(ctx context.Context, request *CreateDomainAliasRequest) (response *CreateDomainAliasResponse, err error) {
    if request == nil {
        request = NewCreateDomainAliasRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDomainAliasResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainBatchRequest() (request *CreateDomainBatchRequest) {
    request = &CreateDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDomainBatch")
    
    
    return
}

func NewCreateDomainBatchResponse() (response *CreateDomainBatchResponse) {
    response = &CreateDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDomainBatch
// 批量添加域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_BATCHDOMAINCREATEACTIONERROR = "InvalidParameter.BatchDomainCreateActionError"
//  INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_JOBGREATERTHANLIMIT = "InvalidParameter.JobGreaterThanLimit"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_TOOMANYINVALIDDOMAINS = "InvalidParameter.TooManyInvalidDomains"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) CreateDomainBatch(request *CreateDomainBatchRequest) (response *CreateDomainBatchResponse, err error) {
    if request == nil {
        request = NewCreateDomainBatchRequest()
    }
    
    response = NewCreateDomainBatchResponse()
    err = c.Send(request, response)
    return
}

// CreateDomainBatch
// 批量添加域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_BATCHDOMAINCREATEACTIONERROR = "InvalidParameter.BatchDomainCreateActionError"
//  INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_JOBGREATERTHANLIMIT = "InvalidParameter.JobGreaterThanLimit"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_TOOMANYINVALIDDOMAINS = "InvalidParameter.TooManyInvalidDomains"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) CreateDomainBatchWithContext(ctx context.Context, request *CreateDomainBatchRequest) (response *CreateDomainBatchResponse, err error) {
    if request == nil {
        request = NewCreateDomainBatchRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainGroupRequest() (request *CreateDomainGroupRequest) {
    request = &CreateDomainGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDomainGroup")
    
    
    return
}

func NewCreateDomainGroupResponse() (response *CreateDomainGroupResponse) {
    response = &CreateDomainGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDomainGroup
// 创建域名分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPNAMEEXISTS = "InvalidParameter.GroupNameExists"
//  INVALIDPARAMETER_GROUPNAMEINVALID = "InvalidParameter.GroupNameInvalid"
//  LIMITEXCEEDED_GROUPNUMBERLIMIT = "LimitExceeded.GroupNumberLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDomainGroup(request *CreateDomainGroupRequest) (response *CreateDomainGroupResponse, err error) {
    if request == nil {
        request = NewCreateDomainGroupRequest()
    }
    
    response = NewCreateDomainGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateDomainGroup
// 创建域名分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPNAMEEXISTS = "InvalidParameter.GroupNameExists"
//  INVALIDPARAMETER_GROUPNAMEINVALID = "InvalidParameter.GroupNameInvalid"
//  LIMITEXCEEDED_GROUPNUMBERLIMIT = "LimitExceeded.GroupNumberLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDomainGroupWithContext(ctx context.Context, request *CreateDomainGroupRequest) (response *CreateDomainGroupResponse, err error) {
    if request == nil {
        request = NewCreateDomainGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDomainGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordRequest() (request *CreateRecordRequest) {
    request = &CreateRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateRecord")
    
    
    return
}

func NewCreateRecordResponse() (response *CreateRecordResponse) {
    response = &CreateRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRecord
// 添加记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DNSSECADDCNAMEERROR = "InvalidParameter.DnssecAddCnameError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_INVALIDWEIGHT = "InvalidParameter.InvalidWeight"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_RECORDTTLLIMIT = "LimitExceeded.RecordTtlLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) CreateRecord(request *CreateRecordRequest) (response *CreateRecordResponse, err error) {
    if request == nil {
        request = NewCreateRecordRequest()
    }
    
    response = NewCreateRecordResponse()
    err = c.Send(request, response)
    return
}

// CreateRecord
// 添加记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DNSSECADDCNAMEERROR = "InvalidParameter.DnssecAddCnameError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_INVALIDWEIGHT = "InvalidParameter.InvalidWeight"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_RECORDTTLLIMIT = "LimitExceeded.RecordTtlLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) CreateRecordWithContext(ctx context.Context, request *CreateRecordRequest) (response *CreateRecordResponse, err error) {
    if request == nil {
        request = NewCreateRecordRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordBatchRequest() (request *CreateRecordBatchRequest) {
    request = &CreateRecordBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateRecordBatch")
    
    
    return
}

func NewCreateRecordBatchResponse() (response *CreateRecordBatchResponse) {
    response = &CreateRecordBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRecordBatch
// 批量添加记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_BATCHRECORDCREATEACTIONERROR = "InvalidParameter.BatchRecordCreateActionError"
//  INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_JOBGREATERTHANLIMIT = "InvalidParameter.JobGreaterThanLimit"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_RECORDSEMPTY = "InvalidParameter.RecordsEmpty"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) CreateRecordBatch(request *CreateRecordBatchRequest) (response *CreateRecordBatchResponse, err error) {
    if request == nil {
        request = NewCreateRecordBatchRequest()
    }
    
    response = NewCreateRecordBatchResponse()
    err = c.Send(request, response)
    return
}

// CreateRecordBatch
// 批量添加记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_BATCHRECORDCREATEACTIONERROR = "InvalidParameter.BatchRecordCreateActionError"
//  INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_JOBGREATERTHANLIMIT = "InvalidParameter.JobGreaterThanLimit"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_RECORDSEMPTY = "InvalidParameter.RecordsEmpty"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) CreateRecordBatchWithContext(ctx context.Context, request *CreateRecordBatchRequest) (response *CreateRecordBatchResponse, err error) {
    if request == nil {
        request = NewCreateRecordBatchRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateRecordBatchResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainRequest() (request *DeleteDomainRequest) {
    request = &DeleteDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteDomain")
    
    
    return
}

func NewDeleteDomainResponse() (response *DeleteDomainResponse) {
    response = &DeleteDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDomain
// 删除域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISKEYDOMAIN = "FailedOperation.DomainIsKeyDomain"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    if request == nil {
        request = NewDeleteDomainRequest()
    }
    
    response = NewDeleteDomainResponse()
    err = c.Send(request, response)
    return
}

// DeleteDomain
// 删除域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISKEYDOMAIN = "FailedOperation.DomainIsKeyDomain"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DeleteDomainWithContext(ctx context.Context, request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    if request == nil {
        request = NewDeleteDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainAliasRequest() (request *DeleteDomainAliasRequest) {
    request = &DeleteDomainAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteDomainAlias")
    
    
    return
}

func NewDeleteDomainAliasResponse() (response *DeleteDomainAliasResponse) {
    response = &DeleteDomainAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDomainAlias
// 删除域名别名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINALIASIDINVALID = "InvalidParameter.DomainAliasIdInvalid"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDomainAlias(request *DeleteDomainAliasRequest) (response *DeleteDomainAliasResponse, err error) {
    if request == nil {
        request = NewDeleteDomainAliasRequest()
    }
    
    response = NewDeleteDomainAliasResponse()
    err = c.Send(request, response)
    return
}

// DeleteDomainAlias
// 删除域名别名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINALIASIDINVALID = "InvalidParameter.DomainAliasIdInvalid"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDomainAliasWithContext(ctx context.Context, request *DeleteDomainAliasRequest) (response *DeleteDomainAliasResponse, err error) {
    if request == nil {
        request = NewDeleteDomainAliasRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteDomainAliasResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordRequest() (request *DeleteRecordRequest) {
    request = &DeleteRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteRecord")
    
    
    return
}

func NewDeleteRecordResponse() (response *DeleteRecordResponse) {
    response = &DeleteRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRecord
// 删除记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DeleteRecord(request *DeleteRecordRequest) (response *DeleteRecordResponse, err error) {
    if request == nil {
        request = NewDeleteRecordRequest()
    }
    
    response = NewDeleteRecordResponse()
    err = c.Send(request, response)
    return
}

// DeleteRecord
// 删除记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DeleteRecordWithContext(ctx context.Context, request *DeleteRecordRequest) (response *DeleteRecordResponse, err error) {
    if request == nil {
        request = NewDeleteRecordRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShareDomainRequest() (request *DeleteShareDomainRequest) {
    request = &DeleteShareDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteShareDomain")
    
    
    return
}

func NewDeleteShareDomainResponse() (response *DeleteShareDomainResponse) {
    response = &DeleteShareDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteShareDomain
// 删除域名共享
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_EMAILINVALID = "InvalidParameter.EmailInvalid"
//  INVALIDPARAMETER_EMAILORQQINVALID = "InvalidParameter.EmailOrQqInvalid"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DeleteShareDomain(request *DeleteShareDomainRequest) (response *DeleteShareDomainResponse, err error) {
    if request == nil {
        request = NewDeleteShareDomainRequest()
    }
    
    response = NewDeleteShareDomainResponse()
    err = c.Send(request, response)
    return
}

// DeleteShareDomain
// 删除域名共享
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_EMAILINVALID = "InvalidParameter.EmailInvalid"
//  INVALIDPARAMETER_EMAILORQQINVALID = "InvalidParameter.EmailOrQqInvalid"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DeleteShareDomainWithContext(ctx context.Context, request *DeleteShareDomainRequest) (response *DeleteShareDomainResponse, err error) {
    if request == nil {
        request = NewDeleteShareDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteShareDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchTaskRequest() (request *DescribeBatchTaskRequest) {
    request = &DescribeBatchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeBatchTask")
    
    
    return
}

func NewDescribeBatchTaskResponse() (response *DescribeBatchTaskResponse) {
    response = &DescribeBatchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBatchTask
// 获取任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTBATCHTASKOWNER = "FailedOperation.NotBatchTaskOwner"
//  INVALIDPARAMETER_BATCHTASKNOTEXIST = "InvalidParameter.BatchTaskNotExist"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
func (c *Client) DescribeBatchTask(request *DescribeBatchTaskRequest) (response *DescribeBatchTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBatchTaskRequest()
    }
    
    response = NewDescribeBatchTaskResponse()
    err = c.Send(request, response)
    return
}

// DescribeBatchTask
// 获取任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTBATCHTASKOWNER = "FailedOperation.NotBatchTaskOwner"
//  INVALIDPARAMETER_BATCHTASKNOTEXIST = "InvalidParameter.BatchTaskNotExist"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
func (c *Client) DescribeBatchTaskWithContext(ctx context.Context, request *DescribeBatchTaskRequest) (response *DescribeBatchTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBatchTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBatchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainRequest() (request *DescribeDomainRequest) {
    request = &DescribeDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomain")
    
    
    return
}

func NewDescribeDomainResponse() (response *DescribeDomainResponse) {
    response = &DescribeDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomain
// 获取域名信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomain(request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    if request == nil {
        request = NewDescribeDomainRequest()
    }
    
    response = NewDescribeDomainResponse()
    err = c.Send(request, response)
    return
}

// DescribeDomain
// 获取域名信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainWithContext(ctx context.Context, request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    if request == nil {
        request = NewDescribeDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainAliasListRequest() (request *DescribeDomainAliasListRequest) {
    request = &DescribeDomainAliasListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainAliasList")
    
    
    return
}

func NewDescribeDomainAliasListResponse() (response *DescribeDomainAliasListResponse) {
    response = &DescribeDomainAliasListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainAliasList
// 获取域名别名列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NODATAOFDOMAINALIAS = "ResourceNotFound.NoDataOfDomainAlias"
func (c *Client) DescribeDomainAliasList(request *DescribeDomainAliasListRequest) (response *DescribeDomainAliasListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainAliasListRequest()
    }
    
    response = NewDescribeDomainAliasListResponse()
    err = c.Send(request, response)
    return
}

// DescribeDomainAliasList
// 获取域名别名列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NODATAOFDOMAINALIAS = "ResourceNotFound.NoDataOfDomainAlias"
func (c *Client) DescribeDomainAliasListWithContext(ctx context.Context, request *DescribeDomainAliasListRequest) (response *DescribeDomainAliasListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainAliasListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDomainAliasListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainListRequest() (request *DescribeDomainListRequest) {
    request = &DescribeDomainListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainList")
    
    
    return
}

func NewDescribeDomainListResponse() (response *DescribeDomainListResponse) {
    response = &DescribeDomainListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainList
// 获取域名列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPIDINVALID = "InvalidParameter.GroupIdInvalid"
//  INVALIDPARAMETER_OFFSETINVALID = "InvalidParameter.OffsetInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_RESULTMORETHAN500 = "InvalidParameter.ResultMoreThan500"
//  INVALIDPARAMETERVALUE_LIMITINVALID = "InvalidParameterValue.LimitInvalid"
//  OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDomainList(request *DescribeDomainListRequest) (response *DescribeDomainListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainListRequest()
    }
    
    response = NewDescribeDomainListResponse()
    err = c.Send(request, response)
    return
}

// DescribeDomainList
// 获取域名列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPIDINVALID = "InvalidParameter.GroupIdInvalid"
//  INVALIDPARAMETER_OFFSETINVALID = "InvalidParameter.OffsetInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_RESULTMORETHAN500 = "InvalidParameter.ResultMoreThan500"
//  INVALIDPARAMETERVALUE_LIMITINVALID = "InvalidParameterValue.LimitInvalid"
//  OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDomainListWithContext(ctx context.Context, request *DescribeDomainListRequest) (response *DescribeDomainListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDomainListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainLogListRequest() (request *DescribeDomainLogListRequest) {
    request = &DescribeDomainLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainLogList")
    
    
    return
}

func NewDescribeDomainLogListResponse() (response *DescribeDomainLogListResponse) {
    response = &DescribeDomainLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainLogList
// 获取域名日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainLogList(request *DescribeDomainLogListRequest) (response *DescribeDomainLogListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainLogListRequest()
    }
    
    response = NewDescribeDomainLogListResponse()
    err = c.Send(request, response)
    return
}

// DescribeDomainLogList
// 获取域名日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainLogListWithContext(ctx context.Context, request *DescribeDomainLogListRequest) (response *DescribeDomainLogListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainLogListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDomainLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainPurviewRequest() (request *DescribeDomainPurviewRequest) {
    request = &DescribeDomainPurviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainPurview")
    
    
    return
}

func NewDescribeDomainPurviewResponse() (response *DescribeDomainPurviewResponse) {
    response = &DescribeDomainPurviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainPurview
// 获取域名权限
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DescribeDomainPurview(request *DescribeDomainPurviewRequest) (response *DescribeDomainPurviewResponse, err error) {
    if request == nil {
        request = NewDescribeDomainPurviewRequest()
    }
    
    response = NewDescribeDomainPurviewResponse()
    err = c.Send(request, response)
    return
}

// DescribeDomainPurview
// 获取域名权限
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DescribeDomainPurviewWithContext(ctx context.Context, request *DescribeDomainPurviewRequest) (response *DescribeDomainPurviewResponse, err error) {
    if request == nil {
        request = NewDescribeDomainPurviewRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDomainPurviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainShareInfoRequest() (request *DescribeDomainShareInfoRequest) {
    request = &DescribeDomainShareInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainShareInfo")
    
    
    return
}

func NewDescribeDomainShareInfoResponse() (response *DescribeDomainShareInfoResponse) {
    response = &DescribeDomainShareInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainShareInfo
// 获取域名共享信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DescribeDomainShareInfo(request *DescribeDomainShareInfoRequest) (response *DescribeDomainShareInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainShareInfoRequest()
    }
    
    response = NewDescribeDomainShareInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeDomainShareInfo
// 获取域名共享信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DescribeDomainShareInfoWithContext(ctx context.Context, request *DescribeDomainShareInfoRequest) (response *DescribeDomainShareInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainShareInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDomainShareInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordRequest() (request *DescribeRecordRequest) {
    request = &DescribeRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecord")
    
    
    return
}

func NewDescribeRecordResponse() (response *DescribeRecordResponse) {
    response = &DescribeRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecord
// 获取记录信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DescribeRecord(request *DescribeRecordRequest) (response *DescribeRecordResponse, err error) {
    if request == nil {
        request = NewDescribeRecordRequest()
    }
    
    response = NewDescribeRecordResponse()
    err = c.Send(request, response)
    return
}

// DescribeRecord
// 获取记录信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DescribeRecordWithContext(ctx context.Context, request *DescribeRecordRequest) (response *DescribeRecordResponse, err error) {
    if request == nil {
        request = NewDescribeRecordRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordLineListRequest() (request *DescribeRecordLineListRequest) {
    request = &DescribeRecordLineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordLineList")
    
    
    return
}

func NewDescribeRecordLineListResponse() (response *DescribeRecordLineListResponse) {
    response = &DescribeRecordLineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordLineList
// 获取等级允许的线路
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINGRADEINVALID = "InvalidParameterValue.DomainGradeInvalid"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DescribeRecordLineList(request *DescribeRecordLineListRequest) (response *DescribeRecordLineListResponse, err error) {
    if request == nil {
        request = NewDescribeRecordLineListRequest()
    }
    
    response = NewDescribeRecordLineListResponse()
    err = c.Send(request, response)
    return
}

// DescribeRecordLineList
// 获取等级允许的线路
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINGRADEINVALID = "InvalidParameterValue.DomainGradeInvalid"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) DescribeRecordLineListWithContext(ctx context.Context, request *DescribeRecordLineListRequest) (response *DescribeRecordLineListResponse, err error) {
    if request == nil {
        request = NewDescribeRecordLineListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRecordLineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordListRequest() (request *DescribeRecordListRequest) {
    request = &DescribeRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordList")
    
    
    return
}

func NewDescribeRecordListResponse() (response *DescribeRecordListResponse) {
    response = &DescribeRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordList
// 获取某个域名下的解析记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RESULTMORETHAN500 = "InvalidParameter.ResultMoreThan500"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_LIMITINVALID = "InvalidParameterValue.LimitInvalid"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_NODATAOFRECORD = "ResourceNotFound.NoDataOfRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRecordList(request *DescribeRecordListRequest) (response *DescribeRecordListResponse, err error) {
    if request == nil {
        request = NewDescribeRecordListRequest()
    }
    
    response = NewDescribeRecordListResponse()
    err = c.Send(request, response)
    return
}

// DescribeRecordList
// 获取某个域名下的解析记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RESULTMORETHAN500 = "InvalidParameter.ResultMoreThan500"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_LIMITINVALID = "InvalidParameterValue.LimitInvalid"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_NODATAOFRECORD = "ResourceNotFound.NoDataOfRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRecordListWithContext(ctx context.Context, request *DescribeRecordListRequest) (response *DescribeRecordListResponse, err error) {
    if request == nil {
        request = NewDescribeRecordListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordTypeRequest() (request *DescribeRecordTypeRequest) {
    request = &DescribeRecordTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordType")
    
    
    return
}

func NewDescribeRecordTypeResponse() (response *DescribeRecordTypeResponse) {
    response = &DescribeRecordTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRecordType
// 获取等级允许的记录类型
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINGRADEINVALID = "InvalidParameterValue.DomainGradeInvalid"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecordType(request *DescribeRecordTypeRequest) (response *DescribeRecordTypeResponse, err error) {
    if request == nil {
        request = NewDescribeRecordTypeRequest()
    }
    
    response = NewDescribeRecordTypeResponse()
    err = c.Send(request, response)
    return
}

// DescribeRecordType
// 获取等级允许的记录类型
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINGRADEINVALID = "InvalidParameterValue.DomainGradeInvalid"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRecordTypeWithContext(ctx context.Context, request *DescribeRecordTypeRequest) (response *DescribeRecordTypeResponse, err error) {
    if request == nil {
        request = NewDescribeRecordTypeRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRecordTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDetailRequest() (request *DescribeUserDetailRequest) {
    request = &DescribeUserDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeUserDetail")
    
    
    return
}

func NewDescribeUserDetailResponse() (response *DescribeUserDetailResponse) {
    response = &DescribeUserDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserDetail
// 获取帐户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserDetail(request *DescribeUserDetailRequest) (response *DescribeUserDetailResponse, err error) {
    if request == nil {
        request = NewDescribeUserDetailRequest()
    }
    
    response = NewDescribeUserDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeUserDetail
// 获取帐户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUserDetailWithContext(ctx context.Context, request *DescribeUserDetailRequest) (response *DescribeUserDetailResponse, err error) {
    if request == nil {
        request = NewDescribeUserDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUserDetailResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainLockRequest() (request *ModifyDomainLockRequest) {
    request = &ModifyDomainLockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainLock")
    
    
    return
}

func NewModifyDomainLockResponse() (response *ModifyDomainLockResponse) {
    response = &ModifyDomainLockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainLock
// 锁定域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDLOCK = "InvalidParameter.DomainNotAllowedLock"
//  INVALIDPARAMETER_LOCKDAYSINVALID = "InvalidParameter.LockDaysInvalid"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyDomainLock(request *ModifyDomainLockRequest) (response *ModifyDomainLockResponse, err error) {
    if request == nil {
        request = NewModifyDomainLockRequest()
    }
    
    response = NewModifyDomainLockResponse()
    err = c.Send(request, response)
    return
}

// ModifyDomainLock
// 锁定域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDLOCK = "InvalidParameter.DomainNotAllowedLock"
//  INVALIDPARAMETER_LOCKDAYSINVALID = "InvalidParameter.LockDaysInvalid"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyDomainLockWithContext(ctx context.Context, request *ModifyDomainLockRequest) (response *ModifyDomainLockResponse, err error) {
    if request == nil {
        request = NewModifyDomainLockRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDomainLockResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainOwnerRequest() (request *ModifyDomainOwnerRequest) {
    request = &ModifyDomainOwnerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainOwner")
    
    
    return
}

func NewModifyDomainOwnerResponse() (response *ModifyDomainOwnerResponse) {
    response = &ModifyDomainOwnerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainOwner
// 域名过户
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_TRANSFERTOENTERPRISEDENIED = "FailedOperation.TransferToEnterpriseDenied"
//  FAILEDOPERATION_TRANSFERTOPERSONDENIED = "FailedOperation.TransferToPersonDenied"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_EMAILINVALID = "InvalidParameter.EmailInvalid"
//  INVALIDPARAMETER_EMAILORQQINVALID = "InvalidParameter.EmailOrQqInvalid"
//  INVALIDPARAMETER_EMAILSAME = "InvalidParameter.EmailSame"
//  INVALIDPARAMETER_OTHERACCOUNTUNREALNAME = "InvalidParameter.OtherAccountUnrealName"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) ModifyDomainOwner(request *ModifyDomainOwnerRequest) (response *ModifyDomainOwnerResponse, err error) {
    if request == nil {
        request = NewModifyDomainOwnerRequest()
    }
    
    response = NewModifyDomainOwnerResponse()
    err = c.Send(request, response)
    return
}

// ModifyDomainOwner
// 域名过户
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_TRANSFERTOENTERPRISEDENIED = "FailedOperation.TransferToEnterpriseDenied"
//  FAILEDOPERATION_TRANSFERTOPERSONDENIED = "FailedOperation.TransferToPersonDenied"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_EMAILINVALID = "InvalidParameter.EmailInvalid"
//  INVALIDPARAMETER_EMAILORQQINVALID = "InvalidParameter.EmailOrQqInvalid"
//  INVALIDPARAMETER_EMAILSAME = "InvalidParameter.EmailSame"
//  INVALIDPARAMETER_OTHERACCOUNTUNREALNAME = "InvalidParameter.OtherAccountUnrealName"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) ModifyDomainOwnerWithContext(ctx context.Context, request *ModifyDomainOwnerRequest) (response *ModifyDomainOwnerResponse, err error) {
    if request == nil {
        request = NewModifyDomainOwnerRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDomainOwnerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainRemarkRequest() (request *ModifyDomainRemarkRequest) {
    request = &ModifyDomainRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainRemark")
    
    
    return
}

func NewModifyDomainRemarkResponse() (response *ModifyDomainRemarkResponse) {
    response = &ModifyDomainRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainRemark
// 设置域名备注
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REMARKTOOLONG = "InvalidParameter.RemarkTooLong"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyDomainRemark(request *ModifyDomainRemarkRequest) (response *ModifyDomainRemarkResponse, err error) {
    if request == nil {
        request = NewModifyDomainRemarkRequest()
    }
    
    response = NewModifyDomainRemarkResponse()
    err = c.Send(request, response)
    return
}

// ModifyDomainRemark
// 设置域名备注
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REMARKTOOLONG = "InvalidParameter.RemarkTooLong"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyDomainRemarkWithContext(ctx context.Context, request *ModifyDomainRemarkRequest) (response *ModifyDomainRemarkResponse, err error) {
    if request == nil {
        request = NewModifyDomainRemarkRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDomainRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainStatusRequest() (request *ModifyDomainStatusRequest) {
    request = &ModifyDomainStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainStatus")
    
    
    return
}

func NewModifyDomainStatusResponse() (response *ModifyDomainStatusResponse) {
    response = &ModifyDomainStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainStatus
// 修改域名状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) ModifyDomainStatus(request *ModifyDomainStatusRequest) (response *ModifyDomainStatusResponse, err error) {
    if request == nil {
        request = NewModifyDomainStatusRequest()
    }
    
    response = NewModifyDomainStatusResponse()
    err = c.Send(request, response)
    return
}

// ModifyDomainStatus
// 修改域名状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) ModifyDomainStatusWithContext(ctx context.Context, request *ModifyDomainStatusRequest) (response *ModifyDomainStatusResponse, err error) {
    if request == nil {
        request = NewModifyDomainStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDomainStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainUnlockRequest() (request *ModifyDomainUnlockRequest) {
    request = &ModifyDomainUnlockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainUnlock")
    
    
    return
}

func NewModifyDomainUnlockResponse() (response *ModifyDomainUnlockResponse) {
    response = &ModifyDomainUnlockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainUnlock
// 域名锁定解锁
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINISNOTLOCKED = "InvalidParameter.DomainIsNotlocked"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNLOCKCODEEXPIRED = "InvalidParameter.UnLockCodeExpired"
//  INVALIDPARAMETER_UNLOCKCODEINVALID = "InvalidParameter.UnLockCodeInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyDomainUnlock(request *ModifyDomainUnlockRequest) (response *ModifyDomainUnlockResponse, err error) {
    if request == nil {
        request = NewModifyDomainUnlockRequest()
    }
    
    response = NewModifyDomainUnlockResponse()
    err = c.Send(request, response)
    return
}

// ModifyDomainUnlock
// 域名锁定解锁
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINISNOTLOCKED = "InvalidParameter.DomainIsNotlocked"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNLOCKCODEEXPIRED = "InvalidParameter.UnLockCodeExpired"
//  INVALIDPARAMETER_UNLOCKCODEINVALID = "InvalidParameter.UnLockCodeInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyDomainUnlockWithContext(ctx context.Context, request *ModifyDomainUnlockRequest) (response *ModifyDomainUnlockResponse, err error) {
    if request == nil {
        request = NewModifyDomainUnlockRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDomainUnlockResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDynamicDNSRequest() (request *ModifyDynamicDNSRequest) {
    request = &ModifyDynamicDNSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDynamicDNS")
    
    
    return
}

func NewModifyDynamicDNSResponse() (response *ModifyDynamicDNSResponse) {
    response = &ModifyDynamicDNSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDynamicDNS
// 更新动态 DNS 记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyDynamicDNS(request *ModifyDynamicDNSRequest) (response *ModifyDynamicDNSResponse, err error) {
    if request == nil {
        request = NewModifyDynamicDNSRequest()
    }
    
    response = NewModifyDynamicDNSResponse()
    err = c.Send(request, response)
    return
}

// ModifyDynamicDNS
// 更新动态 DNS 记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyDynamicDNSWithContext(ctx context.Context, request *ModifyDynamicDNSRequest) (response *ModifyDynamicDNSResponse, err error) {
    if request == nil {
        request = NewModifyDynamicDNSRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDynamicDNSResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordRequest() (request *ModifyRecordRequest) {
    request = &ModifyRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecord")
    
    
    return
}

func NewModifyRecordResponse() (response *ModifyRecordResponse) {
    response = &ModifyRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRecord
// 修改记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_INVALIDWEIGHT = "InvalidParameter.InvalidWeight"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_RECORDTTLLIMIT = "LimitExceeded.RecordTtlLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyRecord(request *ModifyRecordRequest) (response *ModifyRecordResponse, err error) {
    if request == nil {
        request = NewModifyRecordRequest()
    }
    
    response = NewModifyRecordResponse()
    err = c.Send(request, response)
    return
}

// ModifyRecord
// 修改记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_INVALIDWEIGHT = "InvalidParameter.InvalidWeight"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_RECORDTTLLIMIT = "LimitExceeded.RecordTtlLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyRecordWithContext(ctx context.Context, request *ModifyRecordRequest) (response *ModifyRecordResponse, err error) {
    if request == nil {
        request = NewModifyRecordRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyRecordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordBatchRequest() (request *ModifyRecordBatchRequest) {
    request = &ModifyRecordBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecordBatch")
    
    
    return
}

func NewModifyRecordBatchResponse() (response *ModifyRecordBatchResponse) {
    response = &ModifyRecordBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRecordBatch
// 批量修改记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_BATCHRECORDMODIFYACTIONERROR = "InvalidParameter.BatchRecordModifyActionError"
//  INVALIDPARAMETER_BATCHRECORDMODIFYACTIONINVALIDVALUE = "InvalidParameter.BatchRecordModifyActionInvalidValue"
//  INVALIDPARAMETER_BATCHRECORDREPLACEACTIONERROR = "InvalidParameter.BatchRecordReplaceActionError"
//  INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_JOBGREATERTHANLIMIT = "InvalidParameter.JobGreaterThanLimit"
//  INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_RECORDSEMPTY = "InvalidParameter.RecordsEmpty"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyRecordBatch(request *ModifyRecordBatchRequest) (response *ModifyRecordBatchResponse, err error) {
    if request == nil {
        request = NewModifyRecordBatchRequest()
    }
    
    response = NewModifyRecordBatchResponse()
    err = c.Send(request, response)
    return
}

// ModifyRecordBatch
// 批量修改记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_BATCHRECORDMODIFYACTIONERROR = "InvalidParameter.BatchRecordModifyActionError"
//  INVALIDPARAMETER_BATCHRECORDMODIFYACTIONINVALIDVALUE = "InvalidParameter.BatchRecordModifyActionInvalidValue"
//  INVALIDPARAMETER_BATCHRECORDREPLACEACTIONERROR = "InvalidParameter.BatchRecordReplaceActionError"
//  INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_JOBGREATERTHANLIMIT = "InvalidParameter.JobGreaterThanLimit"
//  INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"
//  INVALIDPARAMETER_RECORDSEMPTY = "InvalidParameter.RecordsEmpty"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyRecordBatchWithContext(ctx context.Context, request *ModifyRecordBatchRequest) (response *ModifyRecordBatchResponse, err error) {
    if request == nil {
        request = NewModifyRecordBatchRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyRecordBatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordRemarkRequest() (request *ModifyRecordRemarkRequest) {
    request = &ModifyRecordRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecordRemark")
    
    
    return
}

func NewModifyRecordRemarkResponse() (response *ModifyRecordRemarkResponse) {
    response = &ModifyRecordRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRecordRemark
// 设置记录备注
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_REMARKLENGTHEXCEEDED = "InvalidParameter.RemarkLengthExceeded"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyRecordRemark(request *ModifyRecordRemarkRequest) (response *ModifyRecordRemarkResponse, err error) {
    if request == nil {
        request = NewModifyRecordRemarkRequest()
    }
    
    response = NewModifyRecordRemarkResponse()
    err = c.Send(request, response)
    return
}

// ModifyRecordRemark
// 设置记录备注
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_REMARKLENGTHEXCEEDED = "InvalidParameter.RemarkLengthExceeded"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyRecordRemarkWithContext(ctx context.Context, request *ModifyRecordRemarkRequest) (response *ModifyRecordRemarkResponse, err error) {
    if request == nil {
        request = NewModifyRecordRemarkRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyRecordRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordStatusRequest() (request *ModifyRecordStatusRequest) {
    request = &ModifyRecordStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecordStatus")
    
    
    return
}

func NewModifyRecordStatusResponse() (response *ModifyRecordStatusResponse) {
    response = &ModifyRecordStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRecordStatus
// 修改解析记录的状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DNSSECADDCNAMEERROR = "InvalidParameter.DnssecAddCnameError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyRecordStatus(request *ModifyRecordStatusRequest) (response *ModifyRecordStatusResponse, err error) {
    if request == nil {
        request = NewModifyRecordStatusRequest()
    }
    
    response = NewModifyRecordStatusResponse()
    err = c.Send(request, response)
    return
}

// ModifyRecordStatus
// 修改解析记录的状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DNSSECADDCNAMEERROR = "InvalidParameter.DnssecAddCnameError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifyRecordStatusWithContext(ctx context.Context, request *ModifyRecordStatusRequest) (response *ModifyRecordStatusResponse, err error) {
    if request == nil {
        request = NewModifyRecordStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyRecordStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubdomainStatusRequest() (request *ModifySubdomainStatusRequest) {
    request = &ModifySubdomainStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifySubdomainStatus")
    
    
    return
}

func NewModifySubdomainStatusResponse() (response *ModifySubdomainStatusResponse) {
    response = &ModifySubdomainStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubdomainStatus
// 暂停子域名的解析记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DNSSECADDCNAMEERROR = "InvalidParameter.DnssecAddCnameError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINSELFNOCOPY = "InvalidParameter.DomainSelfNoCopy"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_STATUSCODEINVALID = "InvalidParameter.StatusCodeInvalid"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifySubdomainStatus(request *ModifySubdomainStatusRequest) (response *ModifySubdomainStatusResponse, err error) {
    if request == nil {
        request = NewModifySubdomainStatusRequest()
    }
    
    response = NewModifySubdomainStatusResponse()
    err = c.Send(request, response)
    return
}

// ModifySubdomainStatus
// 暂停子域名的解析记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DNSSECADDCNAMEERROR = "InvalidParameter.DnssecAddCnameError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"
//  INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"
//  INVALIDPARAMETER_DOMAINSELFNOCOPY = "InvalidParameter.DomainSelfNoCopy"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"
//  INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_STATUSCODEINVALID = "InvalidParameter.StatusCodeInvalid"
//  INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"
//  LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"
//  LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"
//  LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"
//  LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"
//  LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"
//  LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"
//  LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
func (c *Client) ModifySubdomainStatusWithContext(ctx context.Context, request *ModifySubdomainStatusRequest) (response *ModifySubdomainStatusResponse, err error) {
    if request == nil {
        request = NewModifySubdomainStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifySubdomainStatusResponse()
    err = c.Send(request, response)
    return
}
