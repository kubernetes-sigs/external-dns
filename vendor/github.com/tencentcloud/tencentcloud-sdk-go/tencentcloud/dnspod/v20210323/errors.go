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

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 该域名已在您的列表中，无需重复添加。
	FAILEDOPERATION_DOMAINEXISTS = "FailedOperation.DomainExists"

	// 该域名为腾讯云 DNSPod 重点保护资源，为了避免误操作造成的业务影响，域名禁止自行操作删除。如果您确认需要删除域名，请先联系您的客户经理，我们将竭诚为您提供技术支持。
	FAILEDOPERATION_DOMAINISKEYDOMAIN = "FailedOperation.DomainIsKeyDomain"

	// 锁定域名不能进行此操作。
	FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"

	// 封禁域名不能进行此操作。
	FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"

	// VIP域名不能进行此操作。
	FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"

	// 该域名已被其他账号添加，可在域名列表中添加取回。
	FAILEDOPERATION_DOMAINOWNEDBYOTHERUSER = "FailedOperation.DomainOwnedByOtherUser"

	// 帐号异地登录，请求被拒绝。
	FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"

	// 登录失败，请检查账号和密码是否正确。
	FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"

	// 权限错误，您无法查看该任务的详情。
	FAILEDOPERATION_NOTBATCHTASKOWNER = "FailedOperation.NotBatchTaskOwner"

	// 域名不在您的名下。
	FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"

	// 未实名认证用户，请先完成实名认证再操作。
	FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"

	// 不能转移到企业账号。
	FAILEDOPERATION_TRANSFERTOENTERPRISEDENIED = "FailedOperation.TransferToEnterpriseDenied"

	// 不能转移到个人账号。
	FAILEDOPERATION_TRANSFERTOPERSONDENIED = "FailedOperation.TransferToPersonDenied"

	// 操作未响应，请稍后重试。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 您的账号已被系统封禁，如果您有任何疑问请与我们联系。
	INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"

	// 创建批量域名任务失败，原因：内部错误。
	INVALIDPARAMETER_BATCHDOMAINCREATEACTIONERROR = "InvalidParameter.BatchDomainCreateActionError"

	// 创建批量记录任务失败，原因：内部错误。
	INVALIDPARAMETER_BATCHRECORDCREATEACTIONERROR = "InvalidParameter.BatchRecordCreateActionError"

	// 批量修改记录任务失败，原因：内部错误。
	INVALIDPARAMETER_BATCHRECORDMODIFYACTIONERROR = "InvalidParameter.BatchRecordModifyActionError"

	// 记录值无效。
	INVALIDPARAMETER_BATCHRECORDMODIFYACTIONINVALIDVALUE = "InvalidParameter.BatchRecordModifyActionInvalidValue"

	// 批量替换记录任务失败，原因：内部错误。
	INVALIDPARAMETER_BATCHRECORDREPLACEACTIONERROR = "InvalidParameter.BatchRecordReplaceActionError"

	// 超过批量任务数上限。
	INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"

	// 任务不存在，无法获取任务详情。
	INVALIDPARAMETER_BATCHTASKNOTEXIST = "InvalidParameter.BatchTaskNotExist"

	// 自定义错误信息。
	INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"

	// 该域名开启了 DNSSEC，不允许添加 @ 子域名 CNAME、显性 URL 或者隐性 URL 记录。
	INVALIDPARAMETER_DNSSECADDCNAMEERROR = "InvalidParameter.DnssecAddCnameError"

	// 别名已经存在。
	INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"

	// 别名编号错误。
	INVALIDPARAMETER_DOMAINALIASIDINVALID = "InvalidParameter.DomainAliasIdInvalid"

	// 域名编号不正确。
	INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"

	// 不允许操作生效中或失效中的域名。
	INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"

	// 域名不正确，请输入主域名，如 dnspod.cn。
	INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"

	// 此域名是其它域名的别名。
	INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"

	// 域名没有锁定。
	INVALIDPARAMETER_DOMAINISNOTLOCKED = "InvalidParameter.DomainIsNotlocked"

	// 暂停域名不支持锁定。
	INVALIDPARAMETER_DOMAINNOTALLOWEDLOCK = "InvalidParameter.DomainNotAllowedLock"

	// 处于生效中/失效中的域名，不允许变更解析记录。
	INVALIDPARAMETER_DOMAINNOTALLOWEDMODIFYRECORDS = "InvalidParameter.DomainNotAllowedModifyRecords"

	// 该域名未备案，无法添加 URL 记录。
	INVALIDPARAMETER_DOMAINNOTBEIAN = "InvalidParameter.DomainNotBeian"

	// 域名还没有注册，无法添加。
	INVALIDPARAMETER_DOMAINNOTREGED = "InvalidParameter.DomainNotReged"

	// 记录已经存在，无需再次添加。
	INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"

	// 域名自己无需进行复制。
	INVALIDPARAMETER_DOMAINSELFNOCOPY = "InvalidParameter.DomainSelfNoCopy"

	// 没有提交任何域名。
	INVALIDPARAMETER_DOMAINSEMPTY = "InvalidParameter.DomainsEmpty"

	// 邮箱地址不正确。
	INVALIDPARAMETER_EMAILINVALID = "InvalidParameter.EmailInvalid"

	// 抱歉，您的账户还没有通过邮箱验证。
	INVALIDPARAMETER_EMAILNOTVERIFIED = "InvalidParameter.EmailNotVerified"

	// 请输入合法的邮箱或者uin。
	INVALIDPARAMETER_EMAILORQQINVALID = "InvalidParameter.EmailOrQqInvalid"

	// 域名已经在该账号下。
	INVALIDPARAMETER_EMAILSAME = "InvalidParameter.EmailSame"

	// 分组编号不正确。
	INVALIDPARAMETER_GROUPIDINVALID = "InvalidParameter.GroupIdInvalid"

	// 同名分组已经存在。
	INVALIDPARAMETER_GROUPNAMEEXISTS = "InvalidParameter.GroupNameExists"

	// 分组名为1-17个字符。
	INVALIDPARAMETER_GROUPNAMEINVALID = "InvalidParameter.GroupNameInvalid"

	// 权重不合法。请输入0~100的整数。
	INVALIDPARAMETER_INVALIDWEIGHT = "InvalidParameter.InvalidWeight"

	// 单次任务数量超过上限。
	INVALIDPARAMETER_JOBGREATERTHANLIMIT = "InvalidParameter.JobGreaterThanLimit"

	// 锁定天数不正确。
	INVALIDPARAMETER_LOCKDAYSINVALID = "InvalidParameter.LockDaysInvalid"

	// Token 的 ID 不正确。
	INVALIDPARAMETER_LOGINTOKENIDERROR = "InvalidParameter.LoginTokenIdError"

	// 传入的 Token 不存在。
	INVALIDPARAMETER_LOGINTOKENNOTEXISTS = "InvalidParameter.LoginTokenNotExists"

	// Token 验证失败。
	INVALIDPARAMETER_LOGINTOKENVALIDATEFAILED = "InvalidParameter.LoginTokenValidateFailed"

	// 抱歉，您的账户还没有通过手机验证。
	INVALIDPARAMETER_MOBILENOTVERIFIED = "InvalidParameter.MobileNotVerified"

	// MX优先级不正确。
	INVALIDPARAMETER_MXINVALID = "InvalidParameter.MxInvalid"

	// 分页起始数量错误。
	INVALIDPARAMETER_OFFSETINVALID = "InvalidParameter.OffsetInvalid"

	// 操作失败，请稍候再试。
	INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"

	// 对方账号未实名认证，请先完成实名认证再操作。
	INVALIDPARAMETER_OTHERACCOUNTUNREALNAME = "InvalidParameter.OtherAccountUnrealName"

	// param 格式错误。
	INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"

	// 参数不合法，请求被拒绝。
	INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"

	// 参数错误。
	INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"

	// 用户UIN无效。
	INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"

	// 记录编号错误。
	INVALIDPARAMETER_RECORDIDINVALID = "InvalidParameter.RecordIdInvalid"

	// 记录线路不正确。
	INVALIDPARAMETER_RECORDLINEINVALID = "InvalidParameter.RecordLineInvalid"

	// 记录类型不正确。
	INVALIDPARAMETER_RECORDTYPEINVALID = "InvalidParameter.RecordTypeInvalid"

	// 记录的值不正确。
	INVALIDPARAMETER_RECORDVALUEINVALID = "InvalidParameter.RecordValueInvalid"

	// 解析记录值过长。
	INVALIDPARAMETER_RECORDVALUELENGTHINVALID = "InvalidParameter.RecordValueLengthInvalid"

	// 没有提交任何记录。
	INVALIDPARAMETER_RECORDSEMPTY = "InvalidParameter.RecordsEmpty"

	// 备注信息超出字符限制。
	INVALIDPARAMETER_REMARKLENGTHEXCEEDED = "InvalidParameter.RemarkLengthExceeded"

	// 备注过长（不能超过200个字）。
	INVALIDPARAMETER_REMARKTOOLONG = "InvalidParameter.RemarkTooLong"

	// 您的IP非法，请求被拒绝。
	INVALIDPARAMETER_REQUESTIPLIMITED = "InvalidParameter.RequestIpLimited"

	// 搜索结果大于500条，请增加关键字。
	INVALIDPARAMETER_RESULTMORETHAN500 = "InvalidParameter.ResultMoreThan500"

	// 状态代码不正确。
	INVALIDPARAMETER_STATUSCODEINVALID = "InvalidParameter.StatusCodeInvalid"

	// 子域名不正确。
	INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"

	// 当前帐号下的无效域名过多，暂时无法使用该功能。请将已有域名的 DNS 服务器正确指向 DNSPod 之后，再尝试添加。
	INVALIDPARAMETER_TOOMANYINVALIDDOMAINS = "InvalidParameter.TooManyInvalidDomains"

	// 域名无效。
	INVALIDPARAMETER_TOOLSDOMAININVALID = "InvalidParameter.ToolsDomainInvalid"

	// 解锁代码已失效。
	INVALIDPARAMETER_UNLOCKCODEEXPIRED = "InvalidParameter.UnLockCodeExpired"

	// 解锁代码不正确。
	INVALIDPARAMETER_UNLOCKCODEINVALID = "InvalidParameter.UnLockCodeInvalid"

	// 未实名认证用户，请先完成实名认证再操作。
	INVALIDPARAMETER_UNREALNAMEUSER = "InvalidParameter.UnrealNameUser"

	// 很抱歉，您要添加的URL的内容不符合DNSPod解析服务条款，URL添加/启用失败，如需帮助请联系技术支持。
	INVALIDPARAMETER_URLVALUEILLEGAL = "InvalidParameter.UrlValueIllegal"

	// 用户不存在。
	INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"

	// 域名等级不正确。
	INVALIDPARAMETERVALUE_DOMAINGRADEINVALID = "InvalidParameterValue.DomainGradeInvalid"

	// 当前域名有误，请返回重新操作。
	INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"

	// 分页长度数量错误。
	INVALIDPARAMETERVALUE_LIMITINVALID = "InvalidParameterValue.LimitInvalid"

	// 用户编号不正确。
	INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"

	// AAAA记录数量超出限制。
	LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"

	// @的NS记录只能设置为默认线路。
	LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"

	// 别名数量已经达到限制。
	LIMITEXCEEDED_DOMAINALIASCOUNTEXCEEDED = "LimitExceeded.DomainAliasCountExceeded"

	// 当前绑定别名数量已达到限制。
	LIMITEXCEEDED_DOMAINALIASNUMBERLIMIT = "LimitExceeded.DomainAliasNumberLimit"

	// 登录失败次数过多已被系统封禁。
	LIMITEXCEEDED_FAILEDLOGINLIMITEXCEEDED = "LimitExceeded.FailedLoginLimitExceeded"

	// 已经达到最大分组数量限制。
	LIMITEXCEEDED_GROUPNUMBERLIMIT = "LimitExceeded.GroupNumberLimit"

	// 该域名使用的套餐不支持隐性URL转发或数量已达上限，如需要使用，请去商城购买。
	LIMITEXCEEDED_HIDDENURLEXCEEDED = "LimitExceeded.HiddenUrlExceeded"

	// NS记录数量超出限制。
	LIMITEXCEEDED_NSCOUNTLIMIT = "LimitExceeded.NsCountLimit"

	// 记录的TTL值超出了限制。
	LIMITEXCEEDED_RECORDTTLLIMIT = "LimitExceeded.RecordTtlLimit"

	// SRV记录数量超出限制。
	LIMITEXCEEDED_SRVCOUNTLIMIT = "LimitExceeded.SrvCountLimit"

	// 子域名级数超出限制。
	LIMITEXCEEDED_SUBDOMAINLEVELLIMIT = "LimitExceeded.SubdomainLevelLimit"

	// 子域名负载均衡数量超出限制。
	LIMITEXCEEDED_SUBDOMAINROLLLIMIT = "LimitExceeded.SubdomainRollLimit"

	// 泛解析级数超出限制。
	LIMITEXCEEDED_SUBDOMAINWCARDLIMIT = "LimitExceeded.SubdomainWcardLimit"

	// 该域名的显性URL转发数量已达上限，如需继续使用，请去商城购买。
	LIMITEXCEEDED_URLCOUNTLIMIT = "LimitExceeded.UrlCountLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 您没有权限执行此操作。
	OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"

	// 仅域名所有者可进行此操作。
	OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"

	// 抱歉，不允许添加黑名单中的IP。
	OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"

	// 当前域名无权限，请返回域名列表。
	OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"

	// 您不是管理用户。
	OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"

	// 您不是代理用户。
	OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"

	// 不是您名下用户。
	OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 您的IP添加了过多任务，请稍后重试。
	REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"

	// 您的帐号在短时间内添加了大量的域名，请控制添加频率。
	REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"

	// API请求次数超出限制。
	REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 没有域名别名。
	RESOURCENOTFOUND_NODATAOFDOMAINALIAS = "ResourceNotFound.NoDataOfDomainAlias"

	// 记录列表为空。
	RESOURCENOTFOUND_NODATAOFRECORD = "ResourceNotFound.NoDataOfRecord"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
