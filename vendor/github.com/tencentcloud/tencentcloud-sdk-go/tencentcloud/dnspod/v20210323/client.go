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
<<<<<<< HEAD
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
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
    "errors"
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


func NewCheckRecordSnapshotRollbackRequest() (request *CheckRecordSnapshotRollbackRequest) {
    request = &CheckRecordSnapshotRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "CheckRecordSnapshotRollback")
    
    
    return
}

func NewCheckRecordSnapshotRollbackResponse() (response *CheckRecordSnapshotRollbackResponse) {
    response = &CheckRecordSnapshotRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckRecordSnapshotRollback
// 回滚前检查单条记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) CheckRecordSnapshotRollback(request *CheckRecordSnapshotRollbackRequest) (response *CheckRecordSnapshotRollbackResponse, err error) {
    return c.CheckRecordSnapshotRollbackWithContext(context.Background(), request)
}

// CheckRecordSnapshotRollback
// 回滚前检查单条记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) CheckRecordSnapshotRollbackWithContext(ctx context.Context, request *CheckRecordSnapshotRollbackRequest) (response *CheckRecordSnapshotRollbackResponse, err error) {
    if request == nil {
        request = NewCheckRecordSnapshotRollbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckRecordSnapshotRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckRecordSnapshotRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewCheckSnapshotRollbackRequest() (request *CheckSnapshotRollbackRequest) {
    request = &CheckSnapshotRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "CheckSnapshotRollback")
    
    
    return
}

func NewCheckSnapshotRollbackResponse() (response *CheckSnapshotRollbackResponse) {
    response = &CheckSnapshotRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckSnapshotRollback
// 快照回滚前检查
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) CheckSnapshotRollback(request *CheckSnapshotRollbackRequest) (response *CheckSnapshotRollbackResponse, err error) {
    return c.CheckSnapshotRollbackWithContext(context.Background(), request)
}

// CheckSnapshotRollback
// 快照回滚前检查
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) CheckSnapshotRollbackWithContext(ctx context.Context, request *CheckSnapshotRollbackRequest) (response *CheckSnapshotRollbackResponse, err error) {
    if request == nil {
        request = NewCheckSnapshotRollbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckSnapshotRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckSnapshotRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDealRequest() (request *CreateDealRequest) {
    request = &CreateDealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDeal")
    
    
    return
}

func NewCreateDealResponse() (response *CreateDealResponse) {
    response = &CreateDealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeal
// DNSPod商品下单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTISLOCKED = "FailedOperation.AccountIsLocked"
//  FAILEDOPERATION_CONTAINSPERSONALVIP = "FailedOperation.ContainsPersonalVip"
//  FAILEDOPERATION_DOMAINISPERSONALTYPE = "FailedOperation.DomainIsPersonalType"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_EIPCHECKFAILED = "FailedOperation.EipCheckFailed"
//  FAILEDOPERATION_FUNCTIONNOTALLOWEDAPPLY = "FailedOperation.FunctionNotAllowedApply"
//  FAILEDOPERATION_GETWHOISFAILED = "FailedOperation.GetWhoisFailed"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_LOGINTIMEOUT = "FailedOperation.LoginTimeout"
//  FAILEDOPERATION_MOBILENOTVERIFIED = "FailedOperation.MobileNotVerified"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_ORDERCANNOTPAY = "FailedOperation.OrderCanNotPay"
//  FAILEDOPERATION_ORDERHASPAID = "FailedOperation.OrderHasPaid"
//  FAILEDOPERATION_RESOURCENOTBIND = "FailedOperation.ResourceNotBind"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  FAILEDOPERATION_VERIFYINGBILLEXISTS = "FailedOperation.VerifyingBillExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"
//  INVALIDPARAMETER_ACTIONSUCCESS = "InvalidParameter.ActionSuccess"
//  INVALIDPARAMETER_ACTIVITY = "InvalidParameter.Activity"
//  INVALIDPARAMETER_BILLNUMBERINVALID = "InvalidParameter.BillNumberInvalid"
//  INVALIDPARAMETER_COMMON = "InvalidParameter.Common"
//  INVALIDPARAMETER_DATAEXPIRED = "InvalidParameter.DataExpired"
//  INVALIDPARAMETER_DATAINVALID = "InvalidParameter.DataInvalid"
//  INVALIDPARAMETER_DEALTYPEINVALID = "InvalidParameter.DealTypeInvalid"
//  INVALIDPARAMETER_DNSDEALDOMAINUPGRADED = "InvalidParameter.DnsDealDomainUpgraded"
//  INVALIDPARAMETER_DNSDEALLOCKED = "InvalidParameter.DnsDealLocked"
//  INVALIDPARAMETER_DNSINVALIDDEAL = "InvalidParameter.DnsInvalidDeal"
//  INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"
//  INVALIDPARAMETER_DOMAINALIASIDINVALID = "InvalidParameter.DomainAliasIdInvalid"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINISMODIFYINGDNS = "InvalidParameter.DomainIsModifyingDns"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_DOMAINNOTVIP = "InvalidParameter.DomainNotVip"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_DOMAINTYPEINVALID = "InvalidParameter.DomainTypeInvalid"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_GOODSCHILDTYPEINVALID = "InvalidParameter.GoodsChildTypeInvalid"
//  INVALIDPARAMETER_GOODSNUMINVALID = "InvalidParameter.GoodsNumInvalid"
//  INVALIDPARAMETER_GOODSTYPEINVALID = "InvalidParameter.GoodsTypeInvalid"
//  INVALIDPARAMETER_GROUPNAMEINVALID = "InvalidParameter.GroupNameInvalid"
//  INVALIDPARAMETER_ILLEGALNEWDEAL = "InvalidParameter.IllegalNewDeal"
//  INVALIDPARAMETER_INVALIDDEALNAME = "InvalidParameter.InvalidDealName"
//  INVALIDPARAMETER_INVALIDSECRETID = "InvalidParameter.InvalidSecretId"
//  INVALIDPARAMETER_INVALIDSIGNATURE = "InvalidParameter.InvalidSignature"
//  INVALIDPARAMETER_INVALIDTIME = "InvalidParameter.InvalidTime"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SHAREDUSERSUNREALNAME = "InvalidParameter.SharedUsersUnrealName"
//  INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"
//  INVALIDPARAMETER_TIMESTAMPEXPIRED = "InvalidParameter.TimestampExpired"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"
//  OPERATIONDENIED_CANCELBILLNOTALLOWED = "OperationDenied.CancelBillNotAllowed"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_MONITORCALLBACKNOTENABLED = "OperationDenied.MonitorCallbackNotEnabled"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTGRANTEDBYOWNER = "OperationDenied.NotGrantedByOwner"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  OPERATIONDENIED_NOTORDEROWNER = "OperationDenied.NotOrderOwner"
//  OPERATIONDENIED_POSTREQUESTACCEPTONLY = "OperationDenied.PostRequestAcceptOnly"
//  OPERATIONDENIED_VIPDOMAINALLOWED = "OperationDenied.VipDomainAllowed"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NODATAOFDOMAIN = "ResourceNotFound.NoDataOfDomain"
//  RESOURCENOTFOUND_NODATAOFDOMAINALIAS = "ResourceNotFound.NoDataOfDomainAlias"
//  RESOURCENOTFOUND_NODATAOFGIFT = "ResourceNotFound.NoDataOfGift"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDeal(request *CreateDealRequest) (response *CreateDealResponse, err error) {
    return c.CreateDealWithContext(context.Background(), request)
}

// CreateDeal
// DNSPod商品下单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTISLOCKED = "FailedOperation.AccountIsLocked"
//  FAILEDOPERATION_CONTAINSPERSONALVIP = "FailedOperation.ContainsPersonalVip"
//  FAILEDOPERATION_DOMAINISPERSONALTYPE = "FailedOperation.DomainIsPersonalType"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_EIPCHECKFAILED = "FailedOperation.EipCheckFailed"
//  FAILEDOPERATION_FUNCTIONNOTALLOWEDAPPLY = "FailedOperation.FunctionNotAllowedApply"
//  FAILEDOPERATION_GETWHOISFAILED = "FailedOperation.GetWhoisFailed"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_LOGINTIMEOUT = "FailedOperation.LoginTimeout"
//  FAILEDOPERATION_MOBILENOTVERIFIED = "FailedOperation.MobileNotVerified"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_ORDERCANNOTPAY = "FailedOperation.OrderCanNotPay"
//  FAILEDOPERATION_ORDERHASPAID = "FailedOperation.OrderHasPaid"
//  FAILEDOPERATION_RESOURCENOTBIND = "FailedOperation.ResourceNotBind"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  FAILEDOPERATION_VERIFYINGBILLEXISTS = "FailedOperation.VerifyingBillExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"
//  INVALIDPARAMETER_ACTIONSUCCESS = "InvalidParameter.ActionSuccess"
//  INVALIDPARAMETER_ACTIVITY = "InvalidParameter.Activity"
//  INVALIDPARAMETER_BILLNUMBERINVALID = "InvalidParameter.BillNumberInvalid"
//  INVALIDPARAMETER_COMMON = "InvalidParameter.Common"
//  INVALIDPARAMETER_DATAEXPIRED = "InvalidParameter.DataExpired"
//  INVALIDPARAMETER_DATAINVALID = "InvalidParameter.DataInvalid"
//  INVALIDPARAMETER_DEALTYPEINVALID = "InvalidParameter.DealTypeInvalid"
//  INVALIDPARAMETER_DNSDEALDOMAINUPGRADED = "InvalidParameter.DnsDealDomainUpgraded"
//  INVALIDPARAMETER_DNSDEALLOCKED = "InvalidParameter.DnsDealLocked"
//  INVALIDPARAMETER_DNSINVALIDDEAL = "InvalidParameter.DnsInvalidDeal"
//  INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"
//  INVALIDPARAMETER_DOMAINALIASIDINVALID = "InvalidParameter.DomainAliasIdInvalid"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINISMODIFYINGDNS = "InvalidParameter.DomainIsModifyingDns"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_DOMAINNOTVIP = "InvalidParameter.DomainNotVip"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_DOMAINTYPEINVALID = "InvalidParameter.DomainTypeInvalid"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_GOODSCHILDTYPEINVALID = "InvalidParameter.GoodsChildTypeInvalid"
//  INVALIDPARAMETER_GOODSNUMINVALID = "InvalidParameter.GoodsNumInvalid"
//  INVALIDPARAMETER_GOODSTYPEINVALID = "InvalidParameter.GoodsTypeInvalid"
//  INVALIDPARAMETER_GROUPNAMEINVALID = "InvalidParameter.GroupNameInvalid"
//  INVALIDPARAMETER_ILLEGALNEWDEAL = "InvalidParameter.IllegalNewDeal"
//  INVALIDPARAMETER_INVALIDDEALNAME = "InvalidParameter.InvalidDealName"
//  INVALIDPARAMETER_INVALIDSECRETID = "InvalidParameter.InvalidSecretId"
//  INVALIDPARAMETER_INVALIDSIGNATURE = "InvalidParameter.InvalidSignature"
//  INVALIDPARAMETER_INVALIDTIME = "InvalidParameter.InvalidTime"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SHAREDUSERSUNREALNAME = "InvalidParameter.SharedUsersUnrealName"
//  INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"
//  INVALIDPARAMETER_TIMESTAMPEXPIRED = "InvalidParameter.TimestampExpired"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"
//  OPERATIONDENIED_CANCELBILLNOTALLOWED = "OperationDenied.CancelBillNotAllowed"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_MONITORCALLBACKNOTENABLED = "OperationDenied.MonitorCallbackNotEnabled"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTGRANTEDBYOWNER = "OperationDenied.NotGrantedByOwner"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  OPERATIONDENIED_NOTORDEROWNER = "OperationDenied.NotOrderOwner"
//  OPERATIONDENIED_POSTREQUESTACCEPTONLY = "OperationDenied.PostRequestAcceptOnly"
//  OPERATIONDENIED_VIPDOMAINALLOWED = "OperationDenied.VipDomainAllowed"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NODATAOFDOMAIN = "ResourceNotFound.NoDataOfDomain"
//  RESOURCENOTFOUND_NODATAOFDOMAINALIAS = "ResourceNotFound.NoDataOfDomainAlias"
//  RESOURCENOTFOUND_NODATAOFGIFT = "ResourceNotFound.NoDataOfGift"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDealWithContext(ctx context.Context, request *CreateDealRequest) (response *CreateDealResponse, err error) {
    if request == nil {
        request = NewCreateDealRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeal require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDealResponse()
    err = c.Send(request, response)
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINEXISTS = "FailedOperation.DomainExists"
//  FAILEDOPERATION_DOMAININENTERPRISEMAILACCOUNT = "FailedOperation.DomainInEnterpriseMailAccount"
//  FAILEDOPERATION_DOMAINOWNEDBYOTHERUSER = "FailedOperation.DomainOwnedByOtherUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_DOMAININBLACKLIST = "InvalidParameter.DomainInBlackList"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINISMYALIAS = "InvalidParameter.DomainIsMyAlias"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_QUHUITXTNOTMATCH = "InvalidParameter.QuhuiTxtNotMatch"
//  INVALIDPARAMETER_QUHUITXTRECORDWAIT = "InvalidParameter.QuhuiTxtRecordWait"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    return c.CreateDomainWithContext(context.Background(), request)
}

// CreateDomain
// 添加域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINEXISTS = "FailedOperation.DomainExists"
//  FAILEDOPERATION_DOMAININENTERPRISEMAILACCOUNT = "FailedOperation.DomainInEnterpriseMailAccount"
//  FAILEDOPERATION_DOMAINOWNEDBYOTHERUSER = "FailedOperation.DomainOwnedByOtherUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_DOMAININBLACKLIST = "InvalidParameter.DomainInBlackList"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINISMYALIAS = "InvalidParameter.DomainIsMyAlias"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_QUHUITXTNOTMATCH = "InvalidParameter.QuhuiTxtNotMatch"
//  INVALIDPARAMETER_QUHUITXTRECORDWAIT = "InvalidParameter.QuhuiTxtRecordWait"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
func (c *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomain require credential")
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
//  INVALIDPARAMETER_ALIASISMYDOMAIN = "InvalidParameter.AliasIsMyDomain"
//  INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  LIMITEXCEEDED_DOMAINALIASCOUNTEXCEEDED = "LimitExceeded.DomainAliasCountExceeded"
//  LIMITEXCEEDED_DOMAINALIASNUMBERLIMIT = "LimitExceeded.DomainAliasNumberLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDomainAlias(request *CreateDomainAliasRequest) (response *CreateDomainAliasResponse, err error) {
    return c.CreateDomainAliasWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_ALIASISMYDOMAIN = "InvalidParameter.AliasIsMyDomain"
//  INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainAlias require credential")
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
//  FAILEDOPERATION = "FailedOperation"
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
    return c.CreateDomainBatchWithContext(context.Background(), request)
}

// CreateDomainBatch
// 批量添加域名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainCustomLineRequest() (request *CreateDomainCustomLineRequest) {
    request = &CreateDomainCustomLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDomainCustomLine")
    
    
    return
}

func NewCreateDomainCustomLineResponse() (response *CreateDomainCustomLineResponse) {
    response = &CreateDomainCustomLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDomainCustomLine
// 创建域名的自定义线路
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"
//  INVALIDPARAMETER_IPALREADYEXIST = "InvalidParameter.IpAlreadyExist"
//  INVALIDPARAMETER_LINEINUSE = "InvalidParameter.LineInUse"
//  INVALIDPARAMETER_LINENAMEINVALID = "InvalidParameter.LineNameInvalid"
//  INVALIDPARAMETER_LINENAMEINVALIDCHARACTER = "InvalidParameter.LineNameInvalidCharacter"
//  INVALIDPARAMETER_LINENAMEOCCUPIED = "InvalidParameter.LineNameOccupied"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYIP = "InvalidParameterValue.IpAreaEmptyIp"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYLINENAME = "InvalidParameterValue.IpAreaEmptyLineName"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_CUSTOMLINELIMITED = "LimitExceeded.CustomLineLimited"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_EDITUSINGRECORDLINENOTALLOWED = "OperationDenied.EditUsingRecordLineNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDomainCustomLine(request *CreateDomainCustomLineRequest) (response *CreateDomainCustomLineResponse, err error) {
    return c.CreateDomainCustomLineWithContext(context.Background(), request)
}

// CreateDomainCustomLine
// 创建域名的自定义线路
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"
//  INVALIDPARAMETER_IPALREADYEXIST = "InvalidParameter.IpAlreadyExist"
//  INVALIDPARAMETER_LINEINUSE = "InvalidParameter.LineInUse"
//  INVALIDPARAMETER_LINENAMEINVALID = "InvalidParameter.LineNameInvalid"
//  INVALIDPARAMETER_LINENAMEINVALIDCHARACTER = "InvalidParameter.LineNameInvalidCharacter"
//  INVALIDPARAMETER_LINENAMEOCCUPIED = "InvalidParameter.LineNameOccupied"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYIP = "InvalidParameterValue.IpAreaEmptyIp"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYLINENAME = "InvalidParameterValue.IpAreaEmptyLineName"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_CUSTOMLINELIMITED = "LimitExceeded.CustomLineLimited"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_EDITUSINGRECORDLINENOTALLOWED = "OperationDenied.EditUsingRecordLineNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateDomainCustomLineWithContext(ctx context.Context, request *CreateDomainCustomLineRequest) (response *CreateDomainCustomLineResponse, err error) {
    if request == nil {
        request = NewCreateDomainCustomLineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainCustomLine require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainCustomLineResponse()
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
    return c.CreateDomainGroupWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainGroup require credential")
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
// 备注：新添加的解析记录存在短暂的索引延迟，如果查询不到新增记录，请在 30 秒后重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DNSSECINCOMPLETECLOSED = "FailedOperation.DNSSECIncompleteClosed"
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
    return c.CreateRecordWithContext(context.Background(), request)
}

// CreateRecord
// 添加记录
//
// 备注：新添加的解析记录存在短暂的索引延迟，如果查询不到新增记录，请在 30 秒后重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DNSSECINCOMPLETECLOSED = "FailedOperation.DNSSECIncompleteClosed"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecord require credential")
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
    return c.CreateRecordBatchWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecordBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecordBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordGroupRequest() (request *CreateRecordGroupRequest) {
    request = &CreateRecordGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateRecordGroup")
    
    
    return
}

func NewCreateRecordGroupResponse() (response *CreateRecordGroupResponse) {
    response = &CreateRecordGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRecordGroup
// 添加记录分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) CreateRecordGroup(request *CreateRecordGroupRequest) (response *CreateRecordGroupResponse, err error) {
    return c.CreateRecordGroupWithContext(context.Background(), request)
}

// CreateRecordGroup
// 添加记录分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) CreateRecordGroupWithContext(ctx context.Context, request *CreateRecordGroupRequest) (response *CreateRecordGroupResponse, err error) {
    if request == nil {
        request = NewCreateRecordGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecordGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecordGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotRequest() (request *CreateSnapshotRequest) {
    request = &CreateSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateSnapshot")
    
    
    return
}

func NewCreateSnapshotResponse() (response *CreateSnapshotResponse) {
    response = &CreateSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSnapshot
// 创建快照
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
    return c.CreateSnapshotWithContext(context.Background(), request)
}

// CreateSnapshot
// 创建快照
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) CreateSnapshotWithContext(ctx context.Context, request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSnapshotResponse()
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
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    return c.DeleteDomainWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomain require credential")
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
    return c.DeleteDomainAliasWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomainAlias require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainAliasResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainBatchRequest() (request *DeleteDomainBatchRequest) {
    request = &DeleteDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteDomainBatch")
    
    
    return
}

func NewDeleteDomainBatchResponse() (response *DeleteDomainBatchResponse) {
    response = &DeleteDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDomainBatch
// 批量删除域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_BATCHRECORDREMOVEACTIONERROR = "InvalidParameter.BatchRecordRemoveActionError"
//  INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
func (c *Client) DeleteDomainBatch(request *DeleteDomainBatchRequest) (response *DeleteDomainBatchResponse, err error) {
    return c.DeleteDomainBatchWithContext(context.Background(), request)
}

// DeleteDomainBatch
// 批量删除域名
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_BATCHRECORDREMOVEACTIONERROR = "InvalidParameter.BatchRecordRemoveActionError"
//  INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
func (c *Client) DeleteDomainBatchWithContext(ctx context.Context, request *DeleteDomainBatchRequest) (response *DeleteDomainBatchResponse, err error) {
    if request == nil {
        request = NewDeleteDomainBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomainBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainCustomLineRequest() (request *DeleteDomainCustomLineRequest) {
    request = &DeleteDomainCustomLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteDomainCustomLine")
    
    
    return
}

func NewDeleteDomainCustomLineResponse() (response *DeleteDomainCustomLineResponse) {
    response = &DeleteDomainCustomLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDomainCustomLine
// 删除域名的自定义线路
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"
//  INVALIDPARAMETER_IPALREADYEXIST = "InvalidParameter.IpAlreadyExist"
//  INVALIDPARAMETER_IPAREA = "InvalidParameter.IpArea"
//  INVALIDPARAMETER_LINEINUSE = "InvalidParameter.LineInUse"
//  INVALIDPARAMETER_LINENAMEINVALID = "InvalidParameter.LineNameInvalid"
//  INVALIDPARAMETER_LINENAMEINVALIDCHARACTER = "InvalidParameter.LineNameInvalidCharacter"
//  INVALIDPARAMETER_LINENAMEOCCUPIED = "InvalidParameter.LineNameOccupied"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYIP = "InvalidParameterValue.IpAreaEmptyIp"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYLINENAME = "InvalidParameterValue.IpAreaEmptyLineName"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_CUSTOMLINELIMITED = "LimitExceeded.CustomLineLimited"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DELETEUSINGRECORDLINENOTALLOWED = "OperationDenied.DeleteUsingRecordLineNotAllowed"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_EDITUSINGRECORDLINENOTALLOWED = "OperationDenied.EditUsingRecordLineNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteDomainCustomLine(request *DeleteDomainCustomLineRequest) (response *DeleteDomainCustomLineResponse, err error) {
    return c.DeleteDomainCustomLineWithContext(context.Background(), request)
}

// DeleteDomainCustomLine
// 删除域名的自定义线路
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"
//  INVALIDPARAMETER_IPALREADYEXIST = "InvalidParameter.IpAlreadyExist"
//  INVALIDPARAMETER_IPAREA = "InvalidParameter.IpArea"
//  INVALIDPARAMETER_LINEINUSE = "InvalidParameter.LineInUse"
//  INVALIDPARAMETER_LINENAMEINVALID = "InvalidParameter.LineNameInvalid"
//  INVALIDPARAMETER_LINENAMEINVALIDCHARACTER = "InvalidParameter.LineNameInvalidCharacter"
//  INVALIDPARAMETER_LINENAMEOCCUPIED = "InvalidParameter.LineNameOccupied"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYIP = "InvalidParameterValue.IpAreaEmptyIp"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYLINENAME = "InvalidParameterValue.IpAreaEmptyLineName"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_CUSTOMLINELIMITED = "LimitExceeded.CustomLineLimited"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DELETEUSINGRECORDLINENOTALLOWED = "OperationDenied.DeleteUsingRecordLineNotAllowed"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_EDITUSINGRECORDLINENOTALLOWED = "OperationDenied.EditUsingRecordLineNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteDomainCustomLineWithContext(ctx context.Context, request *DeleteDomainCustomLineRequest) (response *DeleteDomainCustomLineResponse, err error) {
    if request == nil {
        request = NewDeleteDomainCustomLineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDomainCustomLine require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDomainCustomLineResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
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
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
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
    return c.DeleteRecordWithContext(context.Background(), request)
}

// DeleteRecord
// 删除记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
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
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordBatchRequest() (request *DeleteRecordBatchRequest) {
    request = &DeleteRecordBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteRecordBatch")
    
    
    return
}

func NewDeleteRecordBatchResponse() (response *DeleteRecordBatchResponse) {
    response = &DeleteRecordBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordBatch
// 批量删除解析记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_BATCHRECORDREMOVEACTIONERROR = "InvalidParameter.BatchRecordRemoveActionError"
//  INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
//  INVALIDPARAMETER_RECORDSEMPTY = "InvalidParameter.RecordsEmpty"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
func (c *Client) DeleteRecordBatch(request *DeleteRecordBatchRequest) (response *DeleteRecordBatchResponse, err error) {
    return c.DeleteRecordBatchWithContext(context.Background(), request)
}

// DeleteRecordBatch
// 批量删除解析记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_BATCHRECORDREMOVEACTIONERROR = "InvalidParameter.BatchRecordRemoveActionError"
//  INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
//  INVALIDPARAMETER_RECORDSEMPTY = "InvalidParameter.RecordsEmpty"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
func (c *Client) DeleteRecordBatchWithContext(ctx context.Context, request *DeleteRecordBatchRequest) (response *DeleteRecordBatchResponse, err error) {
    if request == nil {
        request = NewDeleteRecordBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordBatchResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordGroupRequest() (request *DeleteRecordGroupRequest) {
    request = &DeleteRecordGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteRecordGroup")
    
    
    return
}

func NewDeleteRecordGroupResponse() (response *DeleteRecordGroupResponse) {
    response = &DeleteRecordGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordGroup
// 删除记录分组
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DeleteRecordGroup(request *DeleteRecordGroupRequest) (response *DeleteRecordGroupResponse, err error) {
    return c.DeleteRecordGroupWithContext(context.Background(), request)
}

// DeleteRecordGroup
// 删除记录分组
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DeleteRecordGroupWithContext(ctx context.Context, request *DeleteRecordGroupRequest) (response *DeleteRecordGroupResponse, err error) {
    if request == nil {
        request = NewDeleteRecordGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordGroupResponse()
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
// 按账号删除域名共享
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
    return c.DeleteShareDomainWithContext(context.Background(), request)
}

// DeleteShareDomain
// 按账号删除域名共享
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShareDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShareDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotRequest() (request *DeleteSnapshotRequest) {
    request = &DeleteSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteSnapshot")
    
    
    return
}

func NewDeleteSnapshotResponse() (response *DeleteSnapshotResponse) {
    response = &DeleteSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSnapshot
// 删除快照
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (response *DeleteSnapshotResponse, err error) {
    return c.DeleteSnapshotWithContext(context.Background(), request)
}

// DeleteSnapshot
// 删除快照
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DeleteSnapshotWithContext(ctx context.Context, request *DeleteSnapshotRequest) (response *DeleteSnapshotResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSnapshotResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTBATCHTASKOWNER = "FailedOperation.NotBatchTaskOwner"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_BATCHTASKNOTEXIST = "InvalidParameter.BatchTaskNotExist"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
func (c *Client) DescribeBatchTask(request *DescribeBatchTaskRequest) (response *DescribeBatchTaskResponse, err error) {
    return c.DescribeBatchTaskWithContext(context.Background(), request)
}

// DescribeBatchTask
// 获取任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTBATCHTASKOWNER = "FailedOperation.NotBatchTaskOwner"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_BATCHTASKNOTEXIST = "InvalidParameter.BatchTaskNotExist"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
func (c *Client) DescribeBatchTaskWithContext(ctx context.Context, request *DescribeBatchTaskRequest) (response *DescribeBatchTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBatchTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchTask require credential")
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomain(request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    return c.DescribeDomainWithContext(context.Background(), request)
}

// DescribeDomain
// 获取域名信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomain require credential")
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
    return c.DescribeDomainAliasListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainAliasList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainAliasListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainAnalyticsRequest() (request *DescribeDomainAnalyticsRequest) {
    request = &DescribeDomainAnalyticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainAnalytics")
    
    
    return
}

func NewDescribeDomainAnalyticsResponse() (response *DescribeDomainAnalyticsResponse) {
    response = &DescribeDomainAnalyticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainAnalytics
// 统计各个域名的解析量，帮助您了解流量情况、时间段分布。支持查看近 3 个月内的统计情况
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINNOTINSERVICE = "FailedOperation.DomainNotInService"
//  FAILEDOPERATION_TEMPORARYERROR = "FailedOperation.TemporaryError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DescribeDomainAnalytics(request *DescribeDomainAnalyticsRequest) (response *DescribeDomainAnalyticsResponse, err error) {
    return c.DescribeDomainAnalyticsWithContext(context.Background(), request)
}

// DescribeDomainAnalytics
// 统计各个域名的解析量，帮助您了解流量情况、时间段分布。支持查看近 3 个月内的统计情况
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINNOTINSERVICE = "FailedOperation.DomainNotInService"
//  FAILEDOPERATION_TEMPORARYERROR = "FailedOperation.TemporaryError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DescribeDomainAnalyticsWithContext(ctx context.Context, request *DescribeDomainAnalyticsRequest) (response *DescribeDomainAnalyticsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainAnalyticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainAnalytics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainAnalyticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainCustomLineListRequest() (request *DescribeDomainCustomLineListRequest) {
    request = &DescribeDomainCustomLineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainCustomLineList")
    
    
    return
}

func NewDescribeDomainCustomLineListResponse() (response *DescribeDomainCustomLineListResponse) {
    response = &DescribeDomainCustomLineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainCustomLineList
// 获取域名的自定义线路列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDomainCustomLineList(request *DescribeDomainCustomLineListRequest) (response *DescribeDomainCustomLineListResponse, err error) {
    return c.DescribeDomainCustomLineListWithContext(context.Background(), request)
}

// DescribeDomainCustomLineList
// 获取域名的自定义线路列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDomainCustomLineListWithContext(ctx context.Context, request *DescribeDomainCustomLineListRequest) (response *DescribeDomainCustomLineListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainCustomLineListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainCustomLineList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainCustomLineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainFilterListRequest() (request *DescribeDomainFilterListRequest) {
    request = &DescribeDomainFilterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainFilterList")
    
    
    return
}

func NewDescribeDomainFilterListResponse() (response *DescribeDomainFilterListResponse) {
    response = &DescribeDomainFilterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainFilterList
// 获取域名筛选列表
//
// 备注：新添加的解析记录存在短暂的索引延迟，如果查询不到新增记录，请在 30 秒后重试
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
func (c *Client) DescribeDomainFilterList(request *DescribeDomainFilterListRequest) (response *DescribeDomainFilterListResponse, err error) {
    return c.DescribeDomainFilterListWithContext(context.Background(), request)
}

// DescribeDomainFilterList
// 获取域名筛选列表
//
// 备注：新添加的解析记录存在短暂的索引延迟，如果查询不到新增记录，请在 30 秒后重试
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
func (c *Client) DescribeDomainFilterListWithContext(ctx context.Context, request *DescribeDomainFilterListRequest) (response *DescribeDomainFilterListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainFilterListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainFilterList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainFilterListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainGroupListRequest() (request *DescribeDomainGroupListRequest) {
    request = &DescribeDomainGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainGroupList")
    
    
    return
}

func NewDescribeDomainGroupListResponse() (response *DescribeDomainGroupListResponse) {
    response = &DescribeDomainGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainGroupList
// 获取域名分组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDomainGroupList(request *DescribeDomainGroupListRequest) (response *DescribeDomainGroupListResponse, err error) {
    return c.DescribeDomainGroupListWithContext(context.Background(), request)
}

// DescribeDomainGroupList
// 获取域名分组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDomainGroupListWithContext(ctx context.Context, request *DescribeDomainGroupListRequest) (response *DescribeDomainGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainGroupListResponse()
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
    return c.DescribeDomainListWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainList require credential")
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
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainLogList(request *DescribeDomainLogListRequest) (response *DescribeDomainLogListResponse, err error) {
    return c.DescribeDomainLogListWithContext(context.Background(), request)
}

// DescribeDomainLogList
// 获取域名日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainPreviewRequest() (request *DescribeDomainPreviewRequest) {
    request = &DescribeDomainPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainPreview")
    
    
    return
}

func NewDescribeDomainPreviewResponse() (response *DescribeDomainPreviewResponse) {
    response = &DescribeDomainPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainPreview
// 获取域名概览信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainPreview(request *DescribeDomainPreviewRequest) (response *DescribeDomainPreviewResponse, err error) {
    return c.DescribeDomainPreviewWithContext(context.Background(), request)
}

// DescribeDomainPreview
// 获取域名概览信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainPreviewWithContext(ctx context.Context, request *DescribeDomainPreviewRequest) (response *DescribeDomainPreviewResponse, err error) {
    if request == nil {
        request = NewDescribeDomainPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainPreviewResponse()
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
    return c.DescribeDomainPurviewWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainPurview require credential")
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
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainShareInfo(request *DescribeDomainShareInfoRequest) (response *DescribeDomainShareInfoResponse, err error) {
    return c.DescribeDomainShareInfoWithContext(context.Background(), request)
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
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainShareInfoWithContext(ctx context.Context, request *DescribeDomainShareInfoRequest) (response *DescribeDomainShareInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainShareInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainShareInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainShareInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainWhoisRequest() (request *DescribeDomainWhoisRequest) {
    request = &DescribeDomainWhoisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainWhois")
    
    
    return
}

func NewDescribeDomainWhoisResponse() (response *DescribeDomainWhoisResponse) {
    response = &DescribeDomainWhoisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainWhois
// 获取域名Whois信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainWhois(request *DescribeDomainWhoisRequest) (response *DescribeDomainWhoisResponse, err error) {
    return c.DescribeDomainWhoisWithContext(context.Background(), request)
}

// DescribeDomainWhois
// 获取域名Whois信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeDomainWhoisWithContext(ctx context.Context, request *DescribeDomainWhoisRequest) (response *DescribeDomainWhoisResponse, err error) {
    if request == nil {
        request = NewDescribeDomainWhoisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainWhois require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainWhoisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackageDetailRequest() (request *DescribePackageDetailRequest) {
    request = &DescribePackageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribePackageDetail")
    
    
    return
}

func NewDescribePackageDetailResponse() (response *DescribePackageDetailResponse) {
    response = &DescribePackageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePackageDetail
// 获取各套餐配置详情
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTRESOURCEOWNER = "OperationDenied.NotResourceOwner"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribePackageDetail(request *DescribePackageDetailRequest) (response *DescribePackageDetailResponse, err error) {
    return c.DescribePackageDetailWithContext(context.Background(), request)
}

// DescribePackageDetail
// 获取各套餐配置详情
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTRESOURCEOWNER = "OperationDenied.NotResourceOwner"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribePackageDetailWithContext(ctx context.Context, request *DescribePackageDetailRequest) (response *DescribePackageDetailResponse, err error) {
    if request == nil {
        request = NewDescribePackageDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePackageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePackageDetailResponse()
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
//  FAILEDOPERATION = "FailedOperation"
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
    return c.DescribeRecordWithContext(context.Background(), request)
}

// DescribeRecord
// 获取记录信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordExistExceptDefaultNSRequest() (request *DescribeRecordExistExceptDefaultNSRequest) {
    request = &DescribeRecordExistExceptDefaultNSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordExistExceptDefaultNS")
    
    
    return
}

func NewDescribeRecordExistExceptDefaultNSResponse() (response *DescribeRecordExistExceptDefaultNSResponse) {
    response = &DescribeRecordExistExceptDefaultNSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordExistExceptDefaultNS
// 判断是否有除系统默认的@-NS记录之外的记录存在
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
func (c *Client) DescribeRecordExistExceptDefaultNS(request *DescribeRecordExistExceptDefaultNSRequest) (response *DescribeRecordExistExceptDefaultNSResponse, err error) {
    return c.DescribeRecordExistExceptDefaultNSWithContext(context.Background(), request)
}

// DescribeRecordExistExceptDefaultNS
// 判断是否有除系统默认的@-NS记录之外的记录存在
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
func (c *Client) DescribeRecordExistExceptDefaultNSWithContext(ctx context.Context, request *DescribeRecordExistExceptDefaultNSRequest) (response *DescribeRecordExistExceptDefaultNSResponse, err error) {
    if request == nil {
        request = NewDescribeRecordExistExceptDefaultNSRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordExistExceptDefaultNS require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordExistExceptDefaultNSResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordFilterListRequest() (request *DescribeRecordFilterListRequest) {
    request = &DescribeRecordFilterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordFilterList")
    
    
    return
}

func NewDescribeRecordFilterListResponse() (response *DescribeRecordFilterListResponse) {
    response = &DescribeRecordFilterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordFilterList
// 获取某个域名下的解析记录列表
//
// 备注：
//
// 1. 新添加的解析记录存在短暂的索引延迟，如果查询不到新增记录，请在 30 秒后重试
//
// 2.  API获取的记录总条数会比控制台多2条，原因是： 为了防止用户误操作导致解析服务不可用，对2021-10-29 14:24:26之后添加的域名，在控制台都不显示这2条NS记录。
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
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OFFSETINVALID = "InvalidParameter.OffsetInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RESULTMORETHAN500 = "InvalidParameter.ResultMoreThan500"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_LIMITINVALID = "InvalidParameterValue.LimitInvalid"
//  LIMITEXCEEDED_OFFSETEXCEEDED = "LimitExceeded.OffsetExceeded"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_NODATAOFRECORD = "ResourceNotFound.NoDataOfRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRecordFilterList(request *DescribeRecordFilterListRequest) (response *DescribeRecordFilterListResponse, err error) {
    return c.DescribeRecordFilterListWithContext(context.Background(), request)
}

// DescribeRecordFilterList
// 获取某个域名下的解析记录列表
//
// 备注：
//
// 1. 新添加的解析记录存在短暂的索引延迟，如果查询不到新增记录，请在 30 秒后重试
//
// 2.  API获取的记录总条数会比控制台多2条，原因是： 为了防止用户误操作导致解析服务不可用，对2021-10-29 14:24:26之后添加的域名，在控制台都不显示这2条NS记录。
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
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OFFSETINVALID = "InvalidParameter.OffsetInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"
//  INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"
//  INVALIDPARAMETER_RESULTMORETHAN500 = "InvalidParameter.ResultMoreThan500"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_LIMITINVALID = "InvalidParameterValue.LimitInvalid"
//  LIMITEXCEEDED_OFFSETEXCEEDED = "LimitExceeded.OffsetExceeded"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_NODATAOFRECORD = "ResourceNotFound.NoDataOfRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRecordFilterListWithContext(ctx context.Context, request *DescribeRecordFilterListRequest) (response *DescribeRecordFilterListResponse, err error) {
    if request == nil {
        request = NewDescribeRecordFilterListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordFilterList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordFilterListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordGroupListRequest() (request *DescribeRecordGroupListRequest) {
    request = &DescribeRecordGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordGroupList")
    
    
    return
}

func NewDescribeRecordGroupListResponse() (response *DescribeRecordGroupListResponse) {
    response = &DescribeRecordGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordGroupList
// 查询解析记录分组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeRecordGroupList(request *DescribeRecordGroupListRequest) (response *DescribeRecordGroupListResponse, err error) {
    return c.DescribeRecordGroupListWithContext(context.Background(), request)
}

// DescribeRecordGroupList
// 查询解析记录分组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeRecordGroupListWithContext(ctx context.Context, request *DescribeRecordGroupListRequest) (response *DescribeRecordGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeRecordGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordLineCategoryListRequest() (request *DescribeRecordLineCategoryListRequest) {
    request = &DescribeRecordLineCategoryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordLineCategoryList")
    
    
    return
}

func NewDescribeRecordLineCategoryListResponse() (response *DescribeRecordLineCategoryListResponse) {
    response = &DescribeRecordLineCategoryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordLineCategoryList
// 按分类返回线路列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeRecordLineCategoryList(request *DescribeRecordLineCategoryListRequest) (response *DescribeRecordLineCategoryListResponse, err error) {
    return c.DescribeRecordLineCategoryListWithContext(context.Background(), request)
}

// DescribeRecordLineCategoryList
// 按分类返回线路列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeRecordLineCategoryListWithContext(ctx context.Context, request *DescribeRecordLineCategoryListRequest) (response *DescribeRecordLineCategoryListResponse, err error) {
    if request == nil {
        request = NewDescribeRecordLineCategoryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordLineCategoryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordLineCategoryListResponse()
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
//  FAILEDOPERATION = "FailedOperation"
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
    return c.DescribeRecordLineListWithContext(context.Background(), request)
}

// DescribeRecordLineList
// 获取等级允许的线路
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordLineList require credential")
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
// 获取某个域名下的解析记录列表
//
// 备注：
//
// 1. 新添加的解析记录存在短暂的索引延迟，如果查询不到新增记录，请在 30 秒后重试
//
// 2.  API获取的记录总条数会比控制台多2条，原因是： 为了防止用户误操作导致解析服务不可用，对2021-10-29 14:24:26之后添加的域名，在控制台都不显示这2条NS记录。
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
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
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
    return c.DescribeRecordListWithContext(context.Background(), request)
}

// DescribeRecordList
// 获取某个域名下的解析记录列表
//
// 备注：
//
// 1. 新添加的解析记录存在短暂的索引延迟，如果查询不到新增记录，请在 30 秒后重试
//
// 2.  API获取的记录总条数会比控制台多2条，原因是： 为了防止用户误操作导致解析服务不可用，对2021-10-29 14:24:26之后添加的域名，在控制台都不显示这2条NS记录。
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
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordSnapshotRollbackResultRequest() (request *DescribeRecordSnapshotRollbackResultRequest) {
    request = &DescribeRecordSnapshotRollbackResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordSnapshotRollbackResult")
    
    
    return
}

func NewDescribeRecordSnapshotRollbackResultResponse() (response *DescribeRecordSnapshotRollbackResultResponse) {
    response = &DescribeRecordSnapshotRollbackResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordSnapshotRollbackResult
// 查询解析记录重新回滚的结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_BATCHTASKNOTEXIST = "InvalidParameter.BatchTaskNotExist"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeRecordSnapshotRollbackResult(request *DescribeRecordSnapshotRollbackResultRequest) (response *DescribeRecordSnapshotRollbackResultResponse, err error) {
    return c.DescribeRecordSnapshotRollbackResultWithContext(context.Background(), request)
}

// DescribeRecordSnapshotRollbackResult
// 查询解析记录重新回滚的结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_BATCHTASKNOTEXIST = "InvalidParameter.BatchTaskNotExist"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeRecordSnapshotRollbackResultWithContext(ctx context.Context, request *DescribeRecordSnapshotRollbackResultRequest) (response *DescribeRecordSnapshotRollbackResultResponse, err error) {
    if request == nil {
        request = NewDescribeRecordSnapshotRollbackResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordSnapshotRollbackResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordSnapshotRollbackResultResponse()
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
    return c.DescribeRecordTypeWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotConfigRequest() (request *DescribeSnapshotConfigRequest) {
    request = &DescribeSnapshotConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeSnapshotConfig")
    
    
    return
}

func NewDescribeSnapshotConfigResponse() (response *DescribeSnapshotConfigResponse) {
    response = &DescribeSnapshotConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshotConfig
// 查询解析快照配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeSnapshotConfig(request *DescribeSnapshotConfigRequest) (response *DescribeSnapshotConfigResponse, err error) {
    return c.DescribeSnapshotConfigWithContext(context.Background(), request)
}

// DescribeSnapshotConfig
// 查询解析快照配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeSnapshotConfigWithContext(ctx context.Context, request *DescribeSnapshotConfigRequest) (response *DescribeSnapshotConfigResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotListRequest() (request *DescribeSnapshotListRequest) {
    request = &DescribeSnapshotListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeSnapshotList")
    
    
    return
}

func NewDescribeSnapshotListResponse() (response *DescribeSnapshotListResponse) {
    response = &DescribeSnapshotListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshotList
// 查询快照列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeSnapshotList(request *DescribeSnapshotListRequest) (response *DescribeSnapshotListResponse, err error) {
    return c.DescribeSnapshotListWithContext(context.Background(), request)
}

// DescribeSnapshotList
// 查询快照列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeSnapshotListWithContext(ctx context.Context, request *DescribeSnapshotListRequest) (response *DescribeSnapshotListResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotRollbackResultRequest() (request *DescribeSnapshotRollbackResultRequest) {
    request = &DescribeSnapshotRollbackResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeSnapshotRollbackResult")
    
    
    return
}

func NewDescribeSnapshotRollbackResultResponse() (response *DescribeSnapshotRollbackResultResponse) {
    response = &DescribeSnapshotRollbackResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshotRollbackResult
// 查询快照回滚结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeSnapshotRollbackResult(request *DescribeSnapshotRollbackResultRequest) (response *DescribeSnapshotRollbackResultResponse, err error) {
    return c.DescribeSnapshotRollbackResultWithContext(context.Background(), request)
}

// DescribeSnapshotRollbackResult
// 查询快照回滚结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeSnapshotRollbackResultWithContext(ctx context.Context, request *DescribeSnapshotRollbackResultRequest) (response *DescribeSnapshotRollbackResultResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotRollbackResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotRollbackResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotRollbackResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotRollbackTaskRequest() (request *DescribeSnapshotRollbackTaskRequest) {
    request = &DescribeSnapshotRollbackTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeSnapshotRollbackTask")
    
    
    return
}

func NewDescribeSnapshotRollbackTaskResponse() (response *DescribeSnapshotRollbackTaskResponse) {
    response = &DescribeSnapshotRollbackTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshotRollbackTask
// 查询最近一次回滚
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_INNERTASKNOTEXIST = "InvalidParameter.InnerTaskNotExist"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeSnapshotRollbackTask(request *DescribeSnapshotRollbackTaskRequest) (response *DescribeSnapshotRollbackTaskResponse, err error) {
    return c.DescribeSnapshotRollbackTaskWithContext(context.Background(), request)
}

// DescribeSnapshotRollbackTask
// 查询最近一次回滚
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_INNERTASKNOTEXIST = "InvalidParameter.InnerTaskNotExist"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DescribeSnapshotRollbackTaskWithContext(ctx context.Context, request *DescribeSnapshotRollbackTaskRequest) (response *DescribeSnapshotRollbackTaskResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotRollbackTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotRollbackTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotRollbackTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubdomainAnalyticsRequest() (request *DescribeSubdomainAnalyticsRequest) {
    request = &DescribeSubdomainAnalyticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeSubdomainAnalytics")
    
    
    return
}

func NewDescribeSubdomainAnalyticsResponse() (response *DescribeSubdomainAnalyticsResponse) {
    response = &DescribeSubdomainAnalyticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubdomainAnalytics
// 统计子域名的解析量，帮助您了解流量情况、时间段分布。支持查看近 3 个月内的统计情况。仅付费套餐域名可用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINNOTINSERVICE = "FailedOperation.DomainNotInService"
//  FAILEDOPERATION_TEMPORARYERROR = "FailedOperation.TemporaryError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINGRADEINVALID = "InvalidParameterValue.DomainGradeInvalid"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DescribeSubdomainAnalytics(request *DescribeSubdomainAnalyticsRequest) (response *DescribeSubdomainAnalyticsResponse, err error) {
    return c.DescribeSubdomainAnalyticsWithContext(context.Background(), request)
}

// DescribeSubdomainAnalytics
// 统计子域名的解析量，帮助您了解流量情况、时间段分布。支持查看近 3 个月内的统计情况。仅付费套餐域名可用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINNOTINSERVICE = "FailedOperation.DomainNotInService"
//  FAILEDOPERATION_TEMPORARYERROR = "FailedOperation.TemporaryError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINGRADEINVALID = "InvalidParameterValue.DomainGradeInvalid"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) DescribeSubdomainAnalyticsWithContext(ctx context.Context, request *DescribeSubdomainAnalyticsRequest) (response *DescribeSubdomainAnalyticsResponse, err error) {
    if request == nil {
        request = NewDescribeSubdomainAnalyticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubdomainAnalytics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubdomainAnalyticsResponse()
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
// 获取账户信息
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
    return c.DescribeUserDetailWithContext(context.Background(), request)
}

// DescribeUserDetail
// 获取账户信息
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVASStatisticRequest() (request *DescribeVASStatisticRequest) {
    request = &DescribeVASStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeVASStatistic")
    
    
    return
}

func NewDescribeVASStatisticResponse() (response *DescribeVASStatisticResponse) {
    response = &DescribeVASStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVASStatistic
// 获取域名增值服务用量
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTRESOURCEOWNER = "OperationDenied.NotResourceOwner"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVASStatistic(request *DescribeVASStatisticRequest) (response *DescribeVASStatisticResponse, err error) {
    return c.DescribeVASStatisticWithContext(context.Background(), request)
}

// DescribeVASStatistic
// 获取域名增值服务用量
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTRESOURCEOWNER = "OperationDenied.NotResourceOwner"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeVASStatisticWithContext(ctx context.Context, request *DescribeVASStatisticRequest) (response *DescribeVASStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeVASStatisticRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVASStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVASStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadSnapshotRequest() (request *DownloadSnapshotRequest) {
    request = &DownloadSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "DownloadSnapshot")
    
    
    return
}

func NewDownloadSnapshotResponse() (response *DownloadSnapshotResponse) {
    response = &DownloadSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadSnapshot
// 下载快照
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DownloadSnapshot(request *DownloadSnapshotRequest) (response *DownloadSnapshotResponse, err error) {
    return c.DownloadSnapshotWithContext(context.Background(), request)
}

// DownloadSnapshot
// 下载快照
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) DownloadSnapshotWithContext(ctx context.Context, request *DownloadSnapshotRequest) (response *DownloadSnapshotResponse, err error) {
    if request == nil {
        request = NewDownloadSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainCustomLineRequest() (request *ModifyDomainCustomLineRequest) {
    request = &ModifyDomainCustomLineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainCustomLine")
    
    
    return
}

func NewModifyDomainCustomLineResponse() (response *ModifyDomainCustomLineResponse) {
    response = &ModifyDomainCustomLineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDomainCustomLine
// 修改域名的自定义线路
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"
//  INVALIDPARAMETER_IPALREADYEXIST = "InvalidParameter.IpAlreadyExist"
//  INVALIDPARAMETER_LINEINUSE = "InvalidParameter.LineInUse"
//  INVALIDPARAMETER_LINENAMEINVALID = "InvalidParameter.LineNameInvalid"
//  INVALIDPARAMETER_LINENAMEINVALIDCHARACTER = "InvalidParameter.LineNameInvalidCharacter"
//  INVALIDPARAMETER_LINENAMEOCCUPIED = "InvalidParameter.LineNameOccupied"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYIP = "InvalidParameterValue.IpAreaEmptyIp"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYLINENAME = "InvalidParameterValue.IpAreaEmptyLineName"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_CUSTOMLINELIMITED = "LimitExceeded.CustomLineLimited"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_EDITUSINGRECORDLINENOTALLOWED = "OperationDenied.EditUsingRecordLineNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDomainCustomLine(request *ModifyDomainCustomLineRequest) (response *ModifyDomainCustomLineResponse, err error) {
    return c.ModifyDomainCustomLineWithContext(context.Background(), request)
}

// ModifyDomainCustomLine
// 修改域名的自定义线路
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"
//  INVALIDPARAMETER_IPALREADYEXIST = "InvalidParameter.IpAlreadyExist"
//  INVALIDPARAMETER_LINEINUSE = "InvalidParameter.LineInUse"
//  INVALIDPARAMETER_LINENAMEINVALID = "InvalidParameter.LineNameInvalid"
//  INVALIDPARAMETER_LINENAMEINVALIDCHARACTER = "InvalidParameter.LineNameInvalidCharacter"
//  INVALIDPARAMETER_LINENAMEOCCUPIED = "InvalidParameter.LineNameOccupied"
//  INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"
//  INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"
//  INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYIP = "InvalidParameterValue.IpAreaEmptyIp"
//  INVALIDPARAMETERVALUE_IPAREAEMPTYLINENAME = "InvalidParameterValue.IpAreaEmptyLineName"
//  INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"
//  LIMITEXCEEDED_CUSTOMLINELIMITED = "LimitExceeded.CustomLineLimited"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_EDITUSINGRECORDLINENOTALLOWED = "OperationDenied.EditUsingRecordLineNotAllowed"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDomainCustomLineWithContext(ctx context.Context, request *ModifyDomainCustomLineRequest) (response *ModifyDomainCustomLineResponse, err error) {
    if request == nil {
        request = NewModifyDomainCustomLineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainCustomLine require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainCustomLineResponse()
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
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
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
    return c.ModifyDomainLockWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainLock require credential")
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_TRANSFERTOENTERPRISEDENIED = "FailedOperation.TransferToEnterpriseDenied"
//  FAILEDOPERATION_TRANSFERTOPERSONDENIED = "FailedOperation.TransferToPersonDenied"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_EMAILINVALID = "InvalidParameter.EmailInvalid"
//  INVALIDPARAMETER_EMAILORQQINVALID = "InvalidParameter.EmailOrQqInvalid"
//  INVALIDPARAMETER_EMAILSAME = "InvalidParameter.EmailSame"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_OTHERACCOUNTUNREALNAME = "InvalidParameter.OtherAccountUnrealName"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_USERAREAINVALID = "InvalidParameter.UserAreaInvalid"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) ModifyDomainOwner(request *ModifyDomainOwnerRequest) (response *ModifyDomainOwnerResponse, err error) {
    return c.ModifyDomainOwnerWithContext(context.Background(), request)
}

// ModifyDomainOwner
// 域名过户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_TRANSFERTOENTERPRISEDENIED = "FailedOperation.TransferToEnterpriseDenied"
//  FAILEDOPERATION_TRANSFERTOPERSONDENIED = "FailedOperation.TransferToPersonDenied"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_EMAILINVALID = "InvalidParameter.EmailInvalid"
//  INVALIDPARAMETER_EMAILORQQINVALID = "InvalidParameter.EmailOrQqInvalid"
//  INVALIDPARAMETER_EMAILSAME = "InvalidParameter.EmailSame"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_OTHERACCOUNTUNREALNAME = "InvalidParameter.OtherAccountUnrealName"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_USERAREAINVALID = "InvalidParameter.UserAreaInvalid"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
func (c *Client) ModifyDomainOwnerWithContext(ctx context.Context, request *ModifyDomainOwnerRequest) (response *ModifyDomainOwnerResponse, err error) {
    if request == nil {
        request = NewModifyDomainOwnerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainOwner require credential")
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
    return c.ModifyDomainRemarkWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainRemark require credential")
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
//  FAILEDOPERATION_DOMAINISKEYDOMAIN = "FailedOperation.DomainIsKeyDomain"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) ModifyDomainStatus(request *ModifyDomainStatusRequest) (response *ModifyDomainStatusResponse, err error) {
    return c.ModifyDomainStatusWithContext(context.Background(), request)
}

// ModifyDomainStatus
// 修改域名状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINISKEYDOMAIN = "FailedOperation.DomainIsKeyDomain"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) ModifyDomainStatusWithContext(ctx context.Context, request *ModifyDomainStatusRequest) (response *ModifyDomainStatusResponse, err error) {
    if request == nil {
        request = NewModifyDomainStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainStatus require credential")
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
    return c.ModifyDomainUnlockWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainUnlock require credential")
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
//  FAILEDOPERATION = "FailedOperation"
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
    return c.ModifyDynamicDNSWithContext(context.Background(), request)
}

// ModifyDynamicDNS
// 更新动态 DNS 记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDynamicDNS require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDynamicDNSResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPackageAutoRenewRequest() (request *ModifyPackageAutoRenewRequest) {
    request = &ModifyPackageAutoRenewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyPackageAutoRenew")
    
    
    return
}

func NewModifyPackageAutoRenewResponse() (response *ModifyPackageAutoRenewResponse) {
    response = &ModifyPackageAutoRenewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPackageAutoRenew
// DNS 解析套餐自动续费设置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTISLOCKED = "FailedOperation.AccountIsLocked"
//  FAILEDOPERATION_CONTAINSPERSONALVIP = "FailedOperation.ContainsPersonalVip"
//  FAILEDOPERATION_DOMAINISPERSONALTYPE = "FailedOperation.DomainIsPersonalType"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_EIPCHECKFAILED = "FailedOperation.EipCheckFailed"
//  FAILEDOPERATION_FUNCTIONNOTALLOWEDAPPLY = "FailedOperation.FunctionNotAllowedApply"
//  FAILEDOPERATION_GETWHOISFAILED = "FailedOperation.GetWhoisFailed"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_LOGINTIMEOUT = "FailedOperation.LoginTimeout"
//  FAILEDOPERATION_MOBILENOTVERIFIED = "FailedOperation.MobileNotVerified"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_ORDERCANNOTPAY = "FailedOperation.OrderCanNotPay"
//  FAILEDOPERATION_ORDERHASPAID = "FailedOperation.OrderHasPaid"
//  FAILEDOPERATION_RESOURCENOTBIND = "FailedOperation.ResourceNotBind"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  FAILEDOPERATION_VERIFYINGBILLEXISTS = "FailedOperation.VerifyingBillExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"
//  INVALIDPARAMETER_ACTIONSUCCESS = "InvalidParameter.ActionSuccess"
//  INVALIDPARAMETER_ACTIVITY = "InvalidParameter.Activity"
//  INVALIDPARAMETER_BILLNUMBERINVALID = "InvalidParameter.BillNumberInvalid"
//  INVALIDPARAMETER_COMMON = "InvalidParameter.Common"
//  INVALIDPARAMETER_DATAEXPIRED = "InvalidParameter.DataExpired"
//  INVALIDPARAMETER_DATAINVALID = "InvalidParameter.DataInvalid"
//  INVALIDPARAMETER_DEALTYPEINVALID = "InvalidParameter.DealTypeInvalid"
//  INVALIDPARAMETER_DNSDEALDOMAINUPGRADED = "InvalidParameter.DnsDealDomainUpgraded"
//  INVALIDPARAMETER_DNSDEALLOCKED = "InvalidParameter.DnsDealLocked"
//  INVALIDPARAMETER_DNSINVALIDDEAL = "InvalidParameter.DnsInvalidDeal"
//  INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"
//  INVALIDPARAMETER_DOMAINALIASIDINVALID = "InvalidParameter.DomainAliasIdInvalid"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINISMODIFYINGDNS = "InvalidParameter.DomainIsModifyingDns"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_DOMAINNOTVIP = "InvalidParameter.DomainNotVip"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_DOMAINTYPEINVALID = "InvalidParameter.DomainTypeInvalid"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_GOODSCHILDTYPEINVALID = "InvalidParameter.GoodsChildTypeInvalid"
//  INVALIDPARAMETER_GOODSNUMINVALID = "InvalidParameter.GoodsNumInvalid"
//  INVALIDPARAMETER_GOODSTYPEINVALID = "InvalidParameter.GoodsTypeInvalid"
//  INVALIDPARAMETER_GROUPNAMEINVALID = "InvalidParameter.GroupNameInvalid"
//  INVALIDPARAMETER_ILLEGALNEWDEAL = "InvalidParameter.IllegalNewDeal"
//  INVALIDPARAMETER_INVALIDDEALNAME = "InvalidParameter.InvalidDealName"
//  INVALIDPARAMETER_INVALIDSECRETID = "InvalidParameter.InvalidSecretId"
//  INVALIDPARAMETER_INVALIDSIGNATURE = "InvalidParameter.InvalidSignature"
//  INVALIDPARAMETER_INVALIDTIME = "InvalidParameter.InvalidTime"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SHAREDUSERSUNREALNAME = "InvalidParameter.SharedUsersUnrealName"
//  INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"
//  INVALIDPARAMETER_TIMESTAMPEXPIRED = "InvalidParameter.TimestampExpired"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"
//  OPERATIONDENIED_CANCELBILLNOTALLOWED = "OperationDenied.CancelBillNotAllowed"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_MONITORCALLBACKNOTENABLED = "OperationDenied.MonitorCallbackNotEnabled"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTGRANTEDBYOWNER = "OperationDenied.NotGrantedByOwner"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  OPERATIONDENIED_NOTORDEROWNER = "OperationDenied.NotOrderOwner"
//  OPERATIONDENIED_NOTRESOURCEOWNER = "OperationDenied.NotResourceOwner"
//  OPERATIONDENIED_POSTREQUESTACCEPTONLY = "OperationDenied.PostRequestAcceptOnly"
//  OPERATIONDENIED_VIPDOMAINALLOWED = "OperationDenied.VipDomainAllowed"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NODATAOFDOMAIN = "ResourceNotFound.NoDataOfDomain"
//  RESOURCENOTFOUND_NODATAOFDOMAINALIAS = "ResourceNotFound.NoDataOfDomainAlias"
//  RESOURCENOTFOUND_NODATAOFGIFT = "ResourceNotFound.NoDataOfGift"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPackageAutoRenew(request *ModifyPackageAutoRenewRequest) (response *ModifyPackageAutoRenewResponse, err error) {
    return c.ModifyPackageAutoRenewWithContext(context.Background(), request)
}

// ModifyPackageAutoRenew
// DNS 解析套餐自动续费设置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTISLOCKED = "FailedOperation.AccountIsLocked"
//  FAILEDOPERATION_CONTAINSPERSONALVIP = "FailedOperation.ContainsPersonalVip"
//  FAILEDOPERATION_DOMAINISPERSONALTYPE = "FailedOperation.DomainIsPersonalType"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_EIPCHECKFAILED = "FailedOperation.EipCheckFailed"
//  FAILEDOPERATION_FUNCTIONNOTALLOWEDAPPLY = "FailedOperation.FunctionNotAllowedApply"
//  FAILEDOPERATION_GETWHOISFAILED = "FailedOperation.GetWhoisFailed"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_LOGINTIMEOUT = "FailedOperation.LoginTimeout"
//  FAILEDOPERATION_MOBILENOTVERIFIED = "FailedOperation.MobileNotVerified"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_ORDERCANNOTPAY = "FailedOperation.OrderCanNotPay"
//  FAILEDOPERATION_ORDERHASPAID = "FailedOperation.OrderHasPaid"
//  FAILEDOPERATION_RESOURCENOTBIND = "FailedOperation.ResourceNotBind"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  FAILEDOPERATION_VERIFYINGBILLEXISTS = "FailedOperation.VerifyingBillExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"
//  INVALIDPARAMETER_ACTIONSUCCESS = "InvalidParameter.ActionSuccess"
//  INVALIDPARAMETER_ACTIVITY = "InvalidParameter.Activity"
//  INVALIDPARAMETER_BILLNUMBERINVALID = "InvalidParameter.BillNumberInvalid"
//  INVALIDPARAMETER_COMMON = "InvalidParameter.Common"
//  INVALIDPARAMETER_DATAEXPIRED = "InvalidParameter.DataExpired"
//  INVALIDPARAMETER_DATAINVALID = "InvalidParameter.DataInvalid"
//  INVALIDPARAMETER_DEALTYPEINVALID = "InvalidParameter.DealTypeInvalid"
//  INVALIDPARAMETER_DNSDEALDOMAINUPGRADED = "InvalidParameter.DnsDealDomainUpgraded"
//  INVALIDPARAMETER_DNSDEALLOCKED = "InvalidParameter.DnsDealLocked"
//  INVALIDPARAMETER_DNSINVALIDDEAL = "InvalidParameter.DnsInvalidDeal"
//  INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"
//  INVALIDPARAMETER_DOMAINALIASIDINVALID = "InvalidParameter.DomainAliasIdInvalid"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINISMODIFYINGDNS = "InvalidParameter.DomainIsModifyingDns"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_DOMAINNOTVIP = "InvalidParameter.DomainNotVip"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_DOMAINTYPEINVALID = "InvalidParameter.DomainTypeInvalid"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_GOODSCHILDTYPEINVALID = "InvalidParameter.GoodsChildTypeInvalid"
//  INVALIDPARAMETER_GOODSNUMINVALID = "InvalidParameter.GoodsNumInvalid"
//  INVALIDPARAMETER_GOODSTYPEINVALID = "InvalidParameter.GoodsTypeInvalid"
//  INVALIDPARAMETER_GROUPNAMEINVALID = "InvalidParameter.GroupNameInvalid"
//  INVALIDPARAMETER_ILLEGALNEWDEAL = "InvalidParameter.IllegalNewDeal"
//  INVALIDPARAMETER_INVALIDDEALNAME = "InvalidParameter.InvalidDealName"
//  INVALIDPARAMETER_INVALIDSECRETID = "InvalidParameter.InvalidSecretId"
//  INVALIDPARAMETER_INVALIDSIGNATURE = "InvalidParameter.InvalidSignature"
//  INVALIDPARAMETER_INVALIDTIME = "InvalidParameter.InvalidTime"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SHAREDUSERSUNREALNAME = "InvalidParameter.SharedUsersUnrealName"
//  INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"
//  INVALIDPARAMETER_TIMESTAMPEXPIRED = "InvalidParameter.TimestampExpired"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"
//  OPERATIONDENIED_CANCELBILLNOTALLOWED = "OperationDenied.CancelBillNotAllowed"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_MONITORCALLBACKNOTENABLED = "OperationDenied.MonitorCallbackNotEnabled"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTGRANTEDBYOWNER = "OperationDenied.NotGrantedByOwner"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  OPERATIONDENIED_NOTORDEROWNER = "OperationDenied.NotOrderOwner"
//  OPERATIONDENIED_NOTRESOURCEOWNER = "OperationDenied.NotResourceOwner"
//  OPERATIONDENIED_POSTREQUESTACCEPTONLY = "OperationDenied.PostRequestAcceptOnly"
//  OPERATIONDENIED_VIPDOMAINALLOWED = "OperationDenied.VipDomainAllowed"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NODATAOFDOMAIN = "ResourceNotFound.NoDataOfDomain"
//  RESOURCENOTFOUND_NODATAOFDOMAINALIAS = "ResourceNotFound.NoDataOfDomainAlias"
//  RESOURCENOTFOUND_NODATAOFGIFT = "ResourceNotFound.NoDataOfGift"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPackageAutoRenewWithContext(ctx context.Context, request *ModifyPackageAutoRenewRequest) (response *ModifyPackageAutoRenewResponse, err error) {
    if request == nil {
        request = NewModifyPackageAutoRenewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPackageAutoRenew require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPackageAutoRenewResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DNSSECINCOMPLETECLOSED = "FailedOperation.DNSSECIncompleteClosed"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
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
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
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
    return c.ModifyRecordWithContext(context.Background(), request)
}

// ModifyRecord
// 修改记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DNSSECINCOMPLETECLOSED = "FailedOperation.DNSSECIncompleteClosed"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
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
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecord require credential")
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
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
    return c.ModifyRecordBatchWithContext(context.Background(), request)
}

// ModifyRecordBatch
// 批量修改记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordBatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordFieldsRequest() (request *ModifyRecordFieldsRequest) {
    request = &ModifyRecordFieldsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecordFields")
    
    
    return
}

func NewModifyRecordFieldsResponse() (response *ModifyRecordFieldsResponse) {
    response = &ModifyRecordFieldsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRecordFields
// 修改记录可选字段
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
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
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
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
func (c *Client) ModifyRecordFields(request *ModifyRecordFieldsRequest) (response *ModifyRecordFieldsResponse, err error) {
    return c.ModifyRecordFieldsWithContext(context.Background(), request)
}

// ModifyRecordFields
// 修改记录可选字段
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
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
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
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
func (c *Client) ModifyRecordFieldsWithContext(ctx context.Context, request *ModifyRecordFieldsRequest) (response *ModifyRecordFieldsResponse, err error) {
    if request == nil {
        request = NewModifyRecordFieldsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordFields require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordFieldsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordGroupRequest() (request *ModifyRecordGroupRequest) {
    request = &ModifyRecordGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecordGroup")
    
    
    return
}

func NewModifyRecordGroupResponse() (response *ModifyRecordGroupResponse) {
    response = &ModifyRecordGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRecordGroup
// 修改记录分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) ModifyRecordGroup(request *ModifyRecordGroupRequest) (response *ModifyRecordGroupResponse, err error) {
    return c.ModifyRecordGroupWithContext(context.Background(), request)
}

// ModifyRecordGroup
// 修改记录分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) ModifyRecordGroupWithContext(ctx context.Context, request *ModifyRecordGroupRequest) (response *ModifyRecordGroupResponse, err error) {
    if request == nil {
        request = NewModifyRecordGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordGroupResponse()
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
    return c.ModifyRecordRemarkWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordRemark require credential")
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
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
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
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
    return c.ModifyRecordStatusWithContext(context.Background(), request)
}

// ModifyRecordStatus
// 修改解析记录的状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"
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
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordToGroupRequest() (request *ModifyRecordToGroupRequest) {
    request = &ModifyRecordToGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecordToGroup")
    
    
    return
}

func NewModifyRecordToGroupResponse() (response *ModifyRecordToGroupResponse) {
    response = &ModifyRecordToGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRecordToGroup
// 将记录添加到分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) ModifyRecordToGroup(request *ModifyRecordToGroupRequest) (response *ModifyRecordToGroupResponse, err error) {
    return c.ModifyRecordToGroupWithContext(context.Background(), request)
}

// ModifyRecordToGroup
// 将记录添加到分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) ModifyRecordToGroupWithContext(ctx context.Context, request *ModifyRecordToGroupRequest) (response *ModifyRecordToGroupResponse, err error) {
    if request == nil {
        request = NewModifyRecordToGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordToGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordToGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotConfigRequest() (request *ModifySnapshotConfigRequest) {
    request = &ModifySnapshotConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifySnapshotConfig")
    
    
    return
}

func NewModifySnapshotConfigResponse() (response *ModifySnapshotConfigResponse) {
    response = &ModifySnapshotConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySnapshotConfig
// 修改快照配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) ModifySnapshotConfig(request *ModifySnapshotConfigRequest) (response *ModifySnapshotConfigResponse, err error) {
    return c.ModifySnapshotConfigWithContext(context.Background(), request)
}

// ModifySnapshotConfig
// 修改快照配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) ModifySnapshotConfigWithContext(ctx context.Context, request *ModifySnapshotConfigRequest) (response *ModifySnapshotConfigResponse, err error) {
    if request == nil {
        request = NewModifySnapshotConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshotConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySnapshotConfigResponse()
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
    return c.ModifySubdomainStatusWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubdomainStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubdomainStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVasAutoRenewStatusRequest() (request *ModifyVasAutoRenewStatusRequest) {
    request = &ModifyVasAutoRenewStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyVasAutoRenewStatus")
    
    
    return
}

func NewModifyVasAutoRenewStatusResponse() (response *ModifyVasAutoRenewStatusResponse) {
    response = &ModifyVasAutoRenewStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVasAutoRenewStatus
// 增值服务自动续费设置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTISLOCKED = "FailedOperation.AccountIsLocked"
//  FAILEDOPERATION_CONTAINSPERSONALVIP = "FailedOperation.ContainsPersonalVip"
//  FAILEDOPERATION_DOMAINISPERSONALTYPE = "FailedOperation.DomainIsPersonalType"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_EIPCHECKFAILED = "FailedOperation.EipCheckFailed"
//  FAILEDOPERATION_FUNCTIONNOTALLOWEDAPPLY = "FailedOperation.FunctionNotAllowedApply"
//  FAILEDOPERATION_GETWHOISFAILED = "FailedOperation.GetWhoisFailed"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_LOGINTIMEOUT = "FailedOperation.LoginTimeout"
//  FAILEDOPERATION_MOBILENOTVERIFIED = "FailedOperation.MobileNotVerified"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_ORDERCANNOTPAY = "FailedOperation.OrderCanNotPay"
//  FAILEDOPERATION_ORDERHASPAID = "FailedOperation.OrderHasPaid"
//  FAILEDOPERATION_RESOURCENOTBIND = "FailedOperation.ResourceNotBind"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  FAILEDOPERATION_VERIFYINGBILLEXISTS = "FailedOperation.VerifyingBillExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"
//  INVALIDPARAMETER_ACTIONSUCCESS = "InvalidParameter.ActionSuccess"
//  INVALIDPARAMETER_ACTIVITY = "InvalidParameter.Activity"
//  INVALIDPARAMETER_BILLNUMBERINVALID = "InvalidParameter.BillNumberInvalid"
//  INVALIDPARAMETER_COMMON = "InvalidParameter.Common"
//  INVALIDPARAMETER_DATAEXPIRED = "InvalidParameter.DataExpired"
//  INVALIDPARAMETER_DATAINVALID = "InvalidParameter.DataInvalid"
//  INVALIDPARAMETER_DEALTYPEINVALID = "InvalidParameter.DealTypeInvalid"
//  INVALIDPARAMETER_DNSDEALDOMAINUPGRADED = "InvalidParameter.DnsDealDomainUpgraded"
//  INVALIDPARAMETER_DNSDEALLOCKED = "InvalidParameter.DnsDealLocked"
//  INVALIDPARAMETER_DNSINVALIDDEAL = "InvalidParameter.DnsInvalidDeal"
//  INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"
//  INVALIDPARAMETER_DOMAINALIASIDINVALID = "InvalidParameter.DomainAliasIdInvalid"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINISMODIFYINGDNS = "InvalidParameter.DomainIsModifyingDns"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_DOMAINNOTVIP = "InvalidParameter.DomainNotVip"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_DOMAINTYPEINVALID = "InvalidParameter.DomainTypeInvalid"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_GOODSCHILDTYPEINVALID = "InvalidParameter.GoodsChildTypeInvalid"
//  INVALIDPARAMETER_GOODSNUMINVALID = "InvalidParameter.GoodsNumInvalid"
//  INVALIDPARAMETER_GOODSTYPEINVALID = "InvalidParameter.GoodsTypeInvalid"
//  INVALIDPARAMETER_GROUPNAMEINVALID = "InvalidParameter.GroupNameInvalid"
//  INVALIDPARAMETER_ILLEGALNEWDEAL = "InvalidParameter.IllegalNewDeal"
//  INVALIDPARAMETER_INVALIDDEALNAME = "InvalidParameter.InvalidDealName"
//  INVALIDPARAMETER_INVALIDSECRETID = "InvalidParameter.InvalidSecretId"
//  INVALIDPARAMETER_INVALIDSIGNATURE = "InvalidParameter.InvalidSignature"
//  INVALIDPARAMETER_INVALIDTIME = "InvalidParameter.InvalidTime"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SHAREDUSERSUNREALNAME = "InvalidParameter.SharedUsersUnrealName"
//  INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"
//  INVALIDPARAMETER_TIMESTAMPEXPIRED = "InvalidParameter.TimestampExpired"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"
//  OPERATIONDENIED_CANCELBILLNOTALLOWED = "OperationDenied.CancelBillNotAllowed"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_MONITORCALLBACKNOTENABLED = "OperationDenied.MonitorCallbackNotEnabled"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTGRANTEDBYOWNER = "OperationDenied.NotGrantedByOwner"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  OPERATIONDENIED_NOTORDEROWNER = "OperationDenied.NotOrderOwner"
//  OPERATIONDENIED_NOTRESOURCEOWNER = "OperationDenied.NotResourceOwner"
//  OPERATIONDENIED_POSTREQUESTACCEPTONLY = "OperationDenied.PostRequestAcceptOnly"
//  OPERATIONDENIED_RESOURCENOTALLOWRENEW = "OperationDenied.ResourceNotAllowRenew"
//  OPERATIONDENIED_VIPDOMAINALLOWED = "OperationDenied.VipDomainAllowed"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NODATAOFDOMAIN = "ResourceNotFound.NoDataOfDomain"
//  RESOURCENOTFOUND_NODATAOFDOMAINALIAS = "ResourceNotFound.NoDataOfDomainAlias"
//  RESOURCENOTFOUND_NODATAOFGIFT = "ResourceNotFound.NoDataOfGift"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyVasAutoRenewStatus(request *ModifyVasAutoRenewStatusRequest) (response *ModifyVasAutoRenewStatusResponse, err error) {
    return c.ModifyVasAutoRenewStatusWithContext(context.Background(), request)
}

// ModifyVasAutoRenewStatus
// 增值服务自动续费设置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTISLOCKED = "FailedOperation.AccountIsLocked"
//  FAILEDOPERATION_CONTAINSPERSONALVIP = "FailedOperation.ContainsPersonalVip"
//  FAILEDOPERATION_DOMAINISPERSONALTYPE = "FailedOperation.DomainIsPersonalType"
//  FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"
//  FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"
//  FAILEDOPERATION_EIPCHECKFAILED = "FailedOperation.EipCheckFailed"
//  FAILEDOPERATION_FUNCTIONNOTALLOWEDAPPLY = "FailedOperation.FunctionNotAllowedApply"
//  FAILEDOPERATION_GETWHOISFAILED = "FailedOperation.GetWhoisFailed"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_LOGINTIMEOUT = "FailedOperation.LoginTimeout"
//  FAILEDOPERATION_MOBILENOTVERIFIED = "FailedOperation.MobileNotVerified"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_ORDERCANNOTPAY = "FailedOperation.OrderCanNotPay"
//  FAILEDOPERATION_ORDERHASPAID = "FailedOperation.OrderHasPaid"
//  FAILEDOPERATION_RESOURCENOTBIND = "FailedOperation.ResourceNotBind"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  FAILEDOPERATION_VERIFYINGBILLEXISTS = "FailedOperation.VerifyingBillExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"
//  INVALIDPARAMETER_ACTIONSUCCESS = "InvalidParameter.ActionSuccess"
//  INVALIDPARAMETER_ACTIVITY = "InvalidParameter.Activity"
//  INVALIDPARAMETER_BILLNUMBERINVALID = "InvalidParameter.BillNumberInvalid"
//  INVALIDPARAMETER_COMMON = "InvalidParameter.Common"
//  INVALIDPARAMETER_DATAEXPIRED = "InvalidParameter.DataExpired"
//  INVALIDPARAMETER_DATAINVALID = "InvalidParameter.DataInvalid"
//  INVALIDPARAMETER_DEALTYPEINVALID = "InvalidParameter.DealTypeInvalid"
//  INVALIDPARAMETER_DNSDEALDOMAINUPGRADED = "InvalidParameter.DnsDealDomainUpgraded"
//  INVALIDPARAMETER_DNSDEALLOCKED = "InvalidParameter.DnsDealLocked"
//  INVALIDPARAMETER_DNSINVALIDDEAL = "InvalidParameter.DnsInvalidDeal"
//  INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"
//  INVALIDPARAMETER_DOMAINALIASIDINVALID = "InvalidParameter.DomainAliasIdInvalid"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETER_DOMAINISMODIFYINGDNS = "InvalidParameter.DomainIsModifyingDns"
//  INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"
//  INVALIDPARAMETER_DOMAINNOTVIP = "InvalidParameter.DomainNotVip"
//  INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"
//  INVALIDPARAMETER_DOMAINTYPEINVALID = "InvalidParameter.DomainTypeInvalid"
//  INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"
//  INVALIDPARAMETER_GOODSCHILDTYPEINVALID = "InvalidParameter.GoodsChildTypeInvalid"
//  INVALIDPARAMETER_GOODSNUMINVALID = "InvalidParameter.GoodsNumInvalid"
//  INVALIDPARAMETER_GOODSTYPEINVALID = "InvalidParameter.GoodsTypeInvalid"
//  INVALIDPARAMETER_GROUPNAMEINVALID = "InvalidParameter.GroupNameInvalid"
//  INVALIDPARAMETER_ILLEGALNEWDEAL = "InvalidParameter.IllegalNewDeal"
//  INVALIDPARAMETER_INVALIDDEALNAME = "InvalidParameter.InvalidDealName"
//  INVALIDPARAMETER_INVALIDSECRETID = "InvalidParameter.InvalidSecretId"
//  INVALIDPARAMETER_INVALIDSIGNATURE = "InvalidParameter.InvalidSignature"
//  INVALIDPARAMETER_INVALIDTIME = "InvalidParameter.InvalidTime"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"
//  INVALIDPARAMETER_SHAREDUSERSUNREALNAME = "InvalidParameter.SharedUsersUnrealName"
//  INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"
//  INVALIDPARAMETER_TIMESTAMPEXPIRED = "InvalidParameter.TimestampExpired"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"
//  OPERATIONDENIED_CANCELBILLNOTALLOWED = "OperationDenied.CancelBillNotAllowed"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_MONITORCALLBACKNOTENABLED = "OperationDenied.MonitorCallbackNotEnabled"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"
//  OPERATIONDENIED_NOTGRANTEDBYOWNER = "OperationDenied.NotGrantedByOwner"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  OPERATIONDENIED_NOTORDEROWNER = "OperationDenied.NotOrderOwner"
//  OPERATIONDENIED_NOTRESOURCEOWNER = "OperationDenied.NotResourceOwner"
//  OPERATIONDENIED_POSTREQUESTACCEPTONLY = "OperationDenied.PostRequestAcceptOnly"
//  OPERATIONDENIED_RESOURCENOTALLOWRENEW = "OperationDenied.ResourceNotAllowRenew"
//  OPERATIONDENIED_VIPDOMAINALLOWED = "OperationDenied.VipDomainAllowed"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"
//  REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NODATAOFDOMAIN = "ResourceNotFound.NoDataOfDomain"
//  RESOURCENOTFOUND_NODATAOFDOMAINALIAS = "ResourceNotFound.NoDataOfDomainAlias"
//  RESOURCENOTFOUND_NODATAOFGIFT = "ResourceNotFound.NoDataOfGift"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyVasAutoRenewStatusWithContext(ctx context.Context, request *ModifyVasAutoRenewStatusRequest) (response *ModifyVasAutoRenewStatusResponse, err error) {
    if request == nil {
        request = NewModifyVasAutoRenewStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVasAutoRenewStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVasAutoRenewStatusResponse()
    err = c.Send(request, response)
    return
}

func NewPayOrderWithBalanceRequest() (request *PayOrderWithBalanceRequest) {
    request = &PayOrderWithBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "PayOrderWithBalance")
    
    
    return
}

func NewPayOrderWithBalanceResponse() (response *PayOrderWithBalanceResponse) {
    response = &PayOrderWithBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PayOrderWithBalance
// DNSPod商品余额支付
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTISLOCKED = "FailedOperation.AccountIsLocked"
//  FAILEDOPERATION_COUPONFORFREEDOMAIN = "FailedOperation.CouponForFreeDomain"
//  FAILEDOPERATION_COUPONNOTSUPPORTED = "FailedOperation.CouponNotSupported"
//  FAILEDOPERATION_COUPONTYPEALREADYUSED = "FailedOperation.CouponTypeAlreadyUsed"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_LOGINTIMEOUT = "FailedOperation.LoginTimeout"
//  FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_ORDERCANNOTPAY = "FailedOperation.OrderCanNotPay"
//  FAILEDOPERATION_ORDERHASPAID = "FailedOperation.OrderHasPaid"
//  FAILEDOPERATION_RESOURCENOTBIND = "FailedOperation.ResourceNotBind"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  FAILEDOPERATION_VERIFYINGBILLEXISTS = "FailedOperation.VerifyingBillExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"
//  INVALIDPARAMETER_ACTIONSUCCESS = "InvalidParameter.ActionSuccess"
//  INVALIDPARAMETER_ACTIVITY = "InvalidParameter.Activity"
//  INVALIDPARAMETER_BILLNUMBERINVALID = "InvalidParameter.BillNumberInvalid"
//  INVALIDPARAMETER_BROWSERNULL = "InvalidParameter.BrowserNull"
//  INVALIDPARAMETER_COMMON = "InvalidParameter.Common"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DATAEXPIRED = "InvalidParameter.DataExpired"
//  INVALIDPARAMETER_DATAINVALID = "InvalidParameter.DataInvalid"
//  INVALIDPARAMETER_DEALTYPEINVALID = "InvalidParameter.DealTypeInvalid"
//  INVALIDPARAMETER_DNSDEALDOMAINUPGRADED = "InvalidParameter.DnsDealDomainUpgraded"
//  INVALIDPARAMETER_DNSDEALLOCKED = "InvalidParameter.DnsDealLocked"
//  INVALIDPARAMETER_DNSINVALIDDEAL = "InvalidParameter.DnsInvalidDeal"
//  INVALIDPARAMETER_GOODSCHILDTYPEINVALID = "InvalidParameter.GoodsChildTypeInvalid"
//  INVALIDPARAMETER_GOODSNUMINVALID = "InvalidParameter.GoodsNumInvalid"
//  INVALIDPARAMETER_GOODSTYPEINVALID = "InvalidParameter.GoodsTypeInvalid"
//  INVALIDPARAMETER_GRADENOTCOPY = "InvalidParameter.GradeNotCopy"
//  INVALIDPARAMETER_HASPENDINGAPPLY = "InvalidParameter.HasPendingApply"
//  INVALIDPARAMETER_ILLEGALNEWDEAL = "InvalidParameter.IllegalNewDeal"
//  INVALIDPARAMETER_INVALIDCOUPON = "InvalidParameter.InvalidCoupon"
//  INVALIDPARAMETER_INVALIDDEALNAME = "InvalidParameter.InvalidDealName"
//  INVALIDPARAMETER_INVALIDSECRETID = "InvalidParameter.InvalidSecretId"
//  INVALIDPARAMETER_INVALIDSIGNATURE = "InvalidParameter.InvalidSignature"
//  INVALIDPARAMETER_IPSEXCEEDLIMIT = "InvalidParameter.IpsExceedLimit"
//  INVALIDPARAMETER_NEWPACKAGETYPEINVALID = "InvalidParameter.NewPackageTypeInvalid"
//  INVALIDPARAMETER_OPENIDINVALID = "InvalidParameter.OpenidInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_OPERATIONISTOOFREQUENT = "InvalidParameter.OperationIsTooFrequent"
//  INVALIDPARAMETER_OPTYPENOTSUPPORTED = "InvalidParameter.OptypeNotSupported"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_REALNAMEUSER = "InvalidParameter.RealNameUser"
//  INVALIDPARAMETER_TASKNOTCOMPLETED = "InvalidParameter.TaskNotCompleted"
//  INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"
//  INVALIDPARAMETER_TIMESTAMPEXPIRED = "InvalidParameter.TimestampExpired"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERALREADYLOCKED = "InvalidParameter.UserAlreadyLocked"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETER_UUIDINVALID = "InvalidParameter.UuidInvalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DOMAINGRADEINVALID = "InvalidParameterValue.DomainGradeInvalid"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_UPGRADETERMINVALID = "InvalidParameterValue.UpgradeTermInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"
//  OPERATIONDENIED_AGENTDENIED = "OperationDenied.AgentDenied"
//  OPERATIONDENIED_AGENTSUBORDINATEDENIED = "OperationDenied.AgentSubordinateDenied"
//  OPERATIONDENIED_CANCELBILLNOTALLOWED = "OperationDenied.CancelBillNotAllowed"
//  OPERATIONDENIED_DELETEUSINGRECORDLINENOTALLOWED = "OperationDenied.DeleteUsingRecordLineNotAllowed"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_MONITORCALLBACKNOTENABLED = "OperationDenied.MonitorCallbackNotEnabled"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTGRANTEDBYOWNER = "OperationDenied.NotGrantedByOwner"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  OPERATIONDENIED_NOTORDEROWNER = "OperationDenied.NotOrderOwner"
//  OPERATIONDENIED_PERSONALCOUPONNOTALLOWED = "OperationDenied.PersonalCouponNotAllowed"
//  OPERATIONDENIED_POSTREQUESTACCEPTONLY = "OperationDenied.PostRequestAcceptOnly"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PayOrderWithBalance(request *PayOrderWithBalanceRequest) (response *PayOrderWithBalanceResponse, err error) {
    return c.PayOrderWithBalanceWithContext(context.Background(), request)
}

// PayOrderWithBalance
// DNSPod商品余额支付
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTISLOCKED = "FailedOperation.AccountIsLocked"
//  FAILEDOPERATION_COUPONFORFREEDOMAIN = "FailedOperation.CouponForFreeDomain"
//  FAILEDOPERATION_COUPONNOTSUPPORTED = "FailedOperation.CouponNotSupported"
//  FAILEDOPERATION_COUPONTYPEALREADYUSED = "FailedOperation.CouponTypeAlreadyUsed"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_LOGINTIMEOUT = "FailedOperation.LoginTimeout"
//  FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"
//  FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"
//  FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"
//  FAILEDOPERATION_ORDERCANNOTPAY = "FailedOperation.OrderCanNotPay"
//  FAILEDOPERATION_ORDERHASPAID = "FailedOperation.OrderHasPaid"
//  FAILEDOPERATION_RESOURCENOTBIND = "FailedOperation.ResourceNotBind"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  FAILEDOPERATION_VERIFYINGBILLEXISTS = "FailedOperation.VerifyingBillExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"
//  INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"
//  INVALIDPARAMETER_ACTIONSUCCESS = "InvalidParameter.ActionSuccess"
//  INVALIDPARAMETER_ACTIVITY = "InvalidParameter.Activity"
//  INVALIDPARAMETER_BILLNUMBERINVALID = "InvalidParameter.BillNumberInvalid"
//  INVALIDPARAMETER_BROWSERNULL = "InvalidParameter.BrowserNull"
//  INVALIDPARAMETER_COMMON = "InvalidParameter.Common"
//  INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"
//  INVALIDPARAMETER_DATAEXPIRED = "InvalidParameter.DataExpired"
//  INVALIDPARAMETER_DATAINVALID = "InvalidParameter.DataInvalid"
//  INVALIDPARAMETER_DEALTYPEINVALID = "InvalidParameter.DealTypeInvalid"
//  INVALIDPARAMETER_DNSDEALDOMAINUPGRADED = "InvalidParameter.DnsDealDomainUpgraded"
//  INVALIDPARAMETER_DNSDEALLOCKED = "InvalidParameter.DnsDealLocked"
//  INVALIDPARAMETER_DNSINVALIDDEAL = "InvalidParameter.DnsInvalidDeal"
//  INVALIDPARAMETER_GOODSCHILDTYPEINVALID = "InvalidParameter.GoodsChildTypeInvalid"
//  INVALIDPARAMETER_GOODSNUMINVALID = "InvalidParameter.GoodsNumInvalid"
//  INVALIDPARAMETER_GOODSTYPEINVALID = "InvalidParameter.GoodsTypeInvalid"
//  INVALIDPARAMETER_GRADENOTCOPY = "InvalidParameter.GradeNotCopy"
//  INVALIDPARAMETER_HASPENDINGAPPLY = "InvalidParameter.HasPendingApply"
//  INVALIDPARAMETER_ILLEGALNEWDEAL = "InvalidParameter.IllegalNewDeal"
//  INVALIDPARAMETER_INVALIDCOUPON = "InvalidParameter.InvalidCoupon"
//  INVALIDPARAMETER_INVALIDDEALNAME = "InvalidParameter.InvalidDealName"
//  INVALIDPARAMETER_INVALIDSECRETID = "InvalidParameter.InvalidSecretId"
//  INVALIDPARAMETER_INVALIDSIGNATURE = "InvalidParameter.InvalidSignature"
//  INVALIDPARAMETER_IPSEXCEEDLIMIT = "InvalidParameter.IpsExceedLimit"
//  INVALIDPARAMETER_NEWPACKAGETYPEINVALID = "InvalidParameter.NewPackageTypeInvalid"
//  INVALIDPARAMETER_OPENIDINVALID = "InvalidParameter.OpenidInvalid"
//  INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"
//  INVALIDPARAMETER_OPERATIONISTOOFREQUENT = "InvalidParameter.OperationIsTooFrequent"
//  INVALIDPARAMETER_OPTYPENOTSUPPORTED = "InvalidParameter.OptypeNotSupported"
//  INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"
//  INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"
//  INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"
//  INVALIDPARAMETER_REALNAMEUSER = "InvalidParameter.RealNameUser"
//  INVALIDPARAMETER_TASKNOTCOMPLETED = "InvalidParameter.TaskNotCompleted"
//  INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"
//  INVALIDPARAMETER_TIMESTAMPEXPIRED = "InvalidParameter.TimestampExpired"
//  INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"
//  INVALIDPARAMETER_USERALREADYLOCKED = "InvalidParameter.UserAlreadyLocked"
//  INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"
//  INVALIDPARAMETER_UUIDINVALID = "InvalidParameter.UuidInvalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DOMAINGRADEINVALID = "InvalidParameterValue.DomainGradeInvalid"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  INVALIDPARAMETERVALUE_UPGRADETERMINVALID = "InvalidParameterValue.UpgradeTermInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"
//  OPERATIONDENIED_AGENTDENIED = "OperationDenied.AgentDenied"
//  OPERATIONDENIED_AGENTSUBORDINATEDENIED = "OperationDenied.AgentSubordinateDenied"
//  OPERATIONDENIED_CANCELBILLNOTALLOWED = "OperationDenied.CancelBillNotAllowed"
//  OPERATIONDENIED_DELETEUSINGRECORDLINENOTALLOWED = "OperationDenied.DeleteUsingRecordLineNotAllowed"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"
//  OPERATIONDENIED_MONITORCALLBACKNOTENABLED = "OperationDenied.MonitorCallbackNotEnabled"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
//  OPERATIONDENIED_NOTGRANTEDBYOWNER = "OperationDenied.NotGrantedByOwner"
//  OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"
//  OPERATIONDENIED_NOTORDEROWNER = "OperationDenied.NotOrderOwner"
//  OPERATIONDENIED_PERSONALCOUPONNOTALLOWED = "OperationDenied.PersonalCouponNotAllowed"
//  OPERATIONDENIED_POSTREQUESTACCEPTONLY = "OperationDenied.PostRequestAcceptOnly"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PayOrderWithBalanceWithContext(ctx context.Context, request *PayOrderWithBalanceRequest) (response *PayOrderWithBalanceResponse, err error) {
    if request == nil {
        request = NewPayOrderWithBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PayOrderWithBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewPayOrderWithBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackRecordSnapshotRequest() (request *RollbackRecordSnapshotRequest) {
    request = &RollbackRecordSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "RollbackRecordSnapshot")
    
    
    return
}

func NewRollbackRecordSnapshotResponse() (response *RollbackRecordSnapshotResponse) {
    response = &RollbackRecordSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackRecordSnapshot
// 重新回滚指定解析记录快照
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) RollbackRecordSnapshot(request *RollbackRecordSnapshotRequest) (response *RollbackRecordSnapshotResponse, err error) {
    return c.RollbackRecordSnapshotWithContext(context.Background(), request)
}

// RollbackRecordSnapshot
// 重新回滚指定解析记录快照
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) RollbackRecordSnapshotWithContext(ctx context.Context, request *RollbackRecordSnapshotRequest) (response *RollbackRecordSnapshotResponse, err error) {
    if request == nil {
        request = NewRollbackRecordSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackRecordSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackRecordSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackSnapshotRequest() (request *RollbackSnapshotRequest) {
    request = &RollbackSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dnspod", APIVersion, "RollbackSnapshot")
    
    
    return
}

func NewRollbackSnapshotResponse() (response *RollbackSnapshotResponse) {
    response = &RollbackSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackSnapshot
// 回滚快照
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) RollbackSnapshot(request *RollbackSnapshotRequest) (response *RollbackSnapshotResponse, err error) {
    return c.RollbackSnapshotWithContext(context.Background(), request)
}

// RollbackSnapshot
// 回滚快照
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"
//  INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"
//  INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"
//  INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"
//  INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"
//  OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"
//  OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"
func (c *Client) RollbackSnapshotWithContext(ctx context.Context, request *RollbackSnapshotRequest) (response *RollbackSnapshotResponse, err error) {
    if request == nil {
        request = NewRollbackSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackSnapshotResponse()
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
    err = c.Send(request, response)
    return
}
