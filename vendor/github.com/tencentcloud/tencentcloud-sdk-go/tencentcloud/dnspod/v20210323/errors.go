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

<<<<<<< HEAD
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
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	// 操作未授权。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 抱歉，该账户已经被锁定。
	FAILEDOPERATION_ACCOUNTISLOCKED = "FailedOperation.AccountIsLocked"

	// 您的账户下包含个人豪华域名，不能直接升级，请联系销售。
	FAILEDOPERATION_CONTAINSPERSONALVIP = "FailedOperation.ContainsPersonalVip"

	// 此优惠券只能被免费域名使用。
	FAILEDOPERATION_COUPONFORFREEDOMAIN = "FailedOperation.CouponForFreeDomain"

	// 您的账户不满足使用此优惠券的条件。
	FAILEDOPERATION_COUPONNOTSUPPORTED = "FailedOperation.CouponNotSupported"

	// 域名已经使用过该类型的礼券了，不能重复使用。
	FAILEDOPERATION_COUPONTYPEALREADYUSED = "FailedOperation.CouponTypeAlreadyUsed"

	// DNSSEC 未完全关闭，不允许添加 @ 子域名 CNAME、显性 URL 或者隐性 URL 记录。
	FAILEDOPERATION_DNSSECINCOMPLETECLOSED = "FailedOperation.DNSSECIncompleteClosed"

	// 该域名已在您的列表中，无需重复添加。
	FAILEDOPERATION_DOMAINEXISTS = "FailedOperation.DomainExists"

	// 该域名属于企业邮用户
	FAILEDOPERATION_DOMAININENTERPRISEMAILACCOUNT = "FailedOperation.DomainInEnterpriseMailAccount"

	// 该域名为腾讯云 DNSPod 重点保护资源，为了避免误操作造成的业务影响，域名禁止自行操作删除。如果您确认需要删除域名，请先联系您的客户经理，我们将竭诚为您提供技术支持。
	FAILEDOPERATION_DOMAINISKEYDOMAIN = "FailedOperation.DomainIsKeyDomain"

	// 锁定域名不能进行此操作。
	FAILEDOPERATION_DOMAINISLOCKED = "FailedOperation.DomainIsLocked"

	// 域名已升级为个人套餐，但目前位于企业账号，请与销售联系。
	FAILEDOPERATION_DOMAINISPERSONALTYPE = "FailedOperation.DomainIsPersonalType"

	// 封禁域名不能进行此操作。
	FAILEDOPERATION_DOMAINISSPAM = "FailedOperation.DomainIsSpam"

	// VIP域名不能进行此操作。
	FAILEDOPERATION_DOMAINISVIP = "FailedOperation.DomainIsVip"

	// 当前域名还未使用 DNSPod 的解析服务，我们无法获取解析量数据。
	FAILEDOPERATION_DOMAINNOTINSERVICE = "FailedOperation.DomainNotInService"

	// 该域名已被其他账号添加，可在域名列表中添加取回。
	FAILEDOPERATION_DOMAINOWNEDBYOTHERUSER = "FailedOperation.DomainOwnedByOtherUser"

	// 校验公网 IP 发生异常。
	FAILEDOPERATION_EIPCHECKFAILED = "FailedOperation.EipCheckFailed"

	// 您操作过于频繁，请稍后重试
	FAILEDOPERATION_FREQUENCYLIMIT = "FailedOperation.FrequencyLimit"

	// 此功能暂停申请，请稍候重试。
	FAILEDOPERATION_FUNCTIONNOTALLOWEDAPPLY = "FailedOperation.FunctionNotAllowedApply"

	// 获取不到域名信息，可能域名非法或网络故障，请稍后再试。
	FAILEDOPERATION_GETWHOISFAILED = "FailedOperation.GetWhoisFailed"

	// 账户余额不足。
	FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"

	// 账号异地登录，请求被拒绝。
	FAILEDOPERATION_LOGINAREANOTALLOWED = "FailedOperation.LoginAreaNotAllowed"

	// 登录失败，请检查账号和密码是否正确。
	FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"

	// 登录已经超时，请重新登录。
	FAILEDOPERATION_LOGINTIMEOUT = "FailedOperation.LoginTimeout"

	// 用户手机没有通过验证。
	FAILEDOPERATION_MOBILENOTVERIFIED = "FailedOperation.MobileNotVerified"

	// 权限错误，您无法查看该任务的详情。
	FAILEDOPERATION_NOTBATCHTASKOWNER = "FailedOperation.NotBatchTaskOwner"

	// 域名不在您的名下。
	FAILEDOPERATION_NOTDOMAINOWNER = "FailedOperation.NotDomainOwner"

	// 未实名认证用户，请先完成实名认证再操作。
	FAILEDOPERATION_NOTREALNAMEDUSER = "FailedOperation.NotRealNamedUser"

	// 您没有权限操作此资源。
	FAILEDOPERATION_NOTRESOURCEOWNER = "FailedOperation.NotResourceOwner"

	// 您不能付款此订单。
	FAILEDOPERATION_ORDERCANNOTPAY = "FailedOperation.OrderCanNotPay"

	// 此订单已经付过款。
	FAILEDOPERATION_ORDERHASPAID = "FailedOperation.OrderHasPaid"

	// 资源未绑定域名。
	FAILEDOPERATION_RESOURCENOTBIND = "FailedOperation.ResourceNotBind"

	// 请求量统计数据暂时不可用，请稍后再试。
	FAILEDOPERATION_TEMPORARYERROR = "FailedOperation.TemporaryError"

	// 不能转移到企业账号。
	FAILEDOPERATION_TRANSFERTOENTERPRISEDENIED = "FailedOperation.TransferToEnterpriseDenied"

	// 不能转移到个人账号。
	FAILEDOPERATION_TRANSFERTOPERSONDENIED = "FailedOperation.TransferToPersonDenied"

	// 操作未响应，请稍后重试。
	FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnknowError"

	// 域名已经提交过订单并且正在审核中，请稍候。
	FAILEDOPERATION_VERIFYINGBILLEXISTS = "FailedOperation.VerifyingBillExists"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 您的账号已被系统封禁，如果您有任何疑问请与我们联系。
	INVALIDPARAMETER_ACCOUNTISBANNED = "InvalidParameter.AccountIsBanned"

	// 无效的操作。
	INVALIDPARAMETER_ACTIONINVALID = "InvalidParameter.ActionInvalid"

	// 操作已经成功完成。
	INVALIDPARAMETER_ACTIONSUCCESS = "InvalidParameter.ActionSuccess"

	// 用户未实名。
	INVALIDPARAMETER_ACTIVITY = "InvalidParameter.Activity"

	// 该域名已在您的域名列表中，请删除后再添加到别名列表
	INVALIDPARAMETER_ALIASISMYDOMAIN = "InvalidParameter.AliasIsMyDomain"

	// 创建批量域名任务失败，原因：内部错误。
	INVALIDPARAMETER_BATCHDOMAINCREATEACTIONERROR = "InvalidParameter.BatchDomainCreateActionError"

	// 创建批量记录任务失败，原因：内部错误。
	INVALIDPARAMETER_BATCHRECORDCREATEACTIONERROR = "InvalidParameter.BatchRecordCreateActionError"

	// 批量修改记录任务失败，原因：内部错误。
	INVALIDPARAMETER_BATCHRECORDMODIFYACTIONERROR = "InvalidParameter.BatchRecordModifyActionError"

	// 记录值无效。
	INVALIDPARAMETER_BATCHRECORDMODIFYACTIONINVALIDVALUE = "InvalidParameter.BatchRecordModifyActionInvalidValue"

	// 批量删除记录任务失败，原因：内部错误。
	INVALIDPARAMETER_BATCHRECORDREMOVEACTIONERROR = "InvalidParameter.BatchRecordRemoveActionError"

	// 批量替换记录任务失败，原因：内部错误。
	INVALIDPARAMETER_BATCHRECORDREPLACEACTIONERROR = "InvalidParameter.BatchRecordReplaceActionError"

	// 超过批量任务数上限。
	INVALIDPARAMETER_BATCHTASKCOUNTLIMIT = "InvalidParameter.BatchTaskCountLimit"

	// 任务不存在，无法获取任务详情。
	INVALIDPARAMETER_BATCHTASKNOTEXIST = "InvalidParameter.BatchTaskNotExist"

	// 订单号码不正确。
	INVALIDPARAMETER_BILLNUMBERINVALID = "InvalidParameter.BillNumberInvalid"

	// 浏览器字段为空。
	INVALIDPARAMETER_BROWSERNULL = "InvalidParameter.BrowserNull"

	// 您操作过于频繁，请稍后重试。
	INVALIDPARAMETER_COMMON = "InvalidParameter.Common"

	// 自定义错误信息。
	INVALIDPARAMETER_CUSTOMMESSAGE = "InvalidParameter.CustomMessage"

	// 数据过期,请重新提交。
	INVALIDPARAMETER_DATAEXPIRED = "InvalidParameter.DataExpired"

	// data 无效。
	INVALIDPARAMETER_DATAINVALID = "InvalidParameter.DataInvalid"

	// 订单类型无效。
	INVALIDPARAMETER_DEALTYPEINVALID = "InvalidParameter.DealTypeInvalid"

	// 域名已升级至付费套餐，无法完成下单。
	INVALIDPARAMETER_DNSDEALDOMAINUPGRADED = "InvalidParameter.DnsDealDomainUpgraded"

	// 相关服务已有其他未完成的订单，请先将原订单支付或取消后，才可提交新的订单。
	INVALIDPARAMETER_DNSDEALLOCKED = "InvalidParameter.DnsDealLocked"

	// 订单数据非法。
	INVALIDPARAMETER_DNSINVALIDDEAL = "InvalidParameter.DnsInvalidDeal"

	// 该域名开启了 DNSSEC，不允许添加 @ 子域名 CNAME、显性 URL 或者隐性 URL 记录。
	INVALIDPARAMETER_DNSSECADDCNAMEERROR = "InvalidParameter.DnssecAddCnameError"

	// 别名已经存在。
	INVALIDPARAMETER_DOMAINALIASEXISTS = "InvalidParameter.DomainAliasExists"

	// 别名编号错误。
	INVALIDPARAMETER_DOMAINALIASIDINVALID = "InvalidParameter.DomainAliasIdInvalid"

	// 域名编号不正确。
	INVALIDPARAMETER_DOMAINIDINVALID = "InvalidParameter.DomainIdInvalid"

	// 该域名涉及违法违规黑名单，无法进行该操作
	INVALIDPARAMETER_DOMAININBLACKLIST = "InvalidParameter.DomainInBlackList"

	// 不允许操作生效中或失效中的域名。
	INVALIDPARAMETER_DOMAININEFFECTORINVALIDATED = "InvalidParameter.DomainInEffectOrInvalidated"

	// 域名不正确，请输入主域名，如 dnspod.cn。
	INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"

	// 此域名是其它域名的别名。
	INVALIDPARAMETER_DOMAINISALIASER = "InvalidParameter.DomainIsAliaser"

	// 该域名已有同类型操作未完成，无法执行该操作。
	INVALIDPARAMETER_DOMAINISMODIFYINGDNS = "InvalidParameter.DomainIsModifyingDns"

	// 此域名是自己域名的别名
	INVALIDPARAMETER_DOMAINISMYALIAS = "InvalidParameter.DomainIsMyAlias"

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

	// 原域名不是VIP域名，无法替换。
	INVALIDPARAMETER_DOMAINNOTVIP = "InvalidParameter.DomainNotVip"

	// 记录已经存在，无需再次添加。
	INVALIDPARAMETER_DOMAINRECORDEXIST = "InvalidParameter.DomainRecordExist"

	// 域名自己无需进行复制。
	INVALIDPARAMETER_DOMAINSELFNOCOPY = "InvalidParameter.DomainSelfNoCopy"

	// 域名过长。
	INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"

	// 域名类型错误。
	INVALIDPARAMETER_DOMAINTYPEINVALID = "InvalidParameter.DomainTypeInvalid"

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

	// 商品子类型无效。
	INVALIDPARAMETER_GOODSCHILDTYPEINVALID = "InvalidParameter.GoodsChildTypeInvalid"

	// 商品数量无效。
	INVALIDPARAMETER_GOODSNUMINVALID = "InvalidParameter.GoodsNumInvalid"

	// 商品类型无效。
	INVALIDPARAMETER_GOODSTYPEINVALID = "InvalidParameter.GoodsTypeInvalid"

	// 当前域名等级低于源域名的等级，无法进行复制。
	INVALIDPARAMETER_GRADENOTCOPY = "InvalidParameter.GradeNotCopy"

	// 分组编号不正确。
	INVALIDPARAMETER_GROUPIDINVALID = "InvalidParameter.GroupIdInvalid"

	// 同名分组已经存在。
	INVALIDPARAMETER_GROUPNAMEEXISTS = "InvalidParameter.GroupNameExists"

	// 分组名为1-17个字符。
	INVALIDPARAMETER_GROUPNAMEINVALID = "InvalidParameter.GroupNameInvalid"

	// 您已经提交过申请并且正在审核中，请耐心等候。
	INVALIDPARAMETER_HASPENDINGAPPLY = "InvalidParameter.HasPendingApply"

	// 该订单存在冲突或参数有误，无法完成支付，请重新购买。
	INVALIDPARAMETER_ILLEGALNEWDEAL = "InvalidParameter.IllegalNewDeal"

	// 任务不存在。
	INVALIDPARAMETER_INNERTASKNOTEXIST = "InvalidParameter.InnerTaskNotExist"

	// 礼券代码无效。
	INVALIDPARAMETER_INVALIDCOUPON = "InvalidParameter.InvalidCoupon"

	// 请输入正确的订单号。
	INVALIDPARAMETER_INVALIDDEALNAME = "InvalidParameter.InvalidDealName"

	// 不是合法的IP段。
	INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"

	// 无效密钥 ID。
	INVALIDPARAMETER_INVALIDSECRETID = "InvalidParameter.InvalidSecretId"

	// 无效签名。
	INVALIDPARAMETER_INVALIDSIGNATURE = "InvalidParameter.InvalidSignature"

	// 无效的时间。
	INVALIDPARAMETER_INVALIDTIME = "InvalidParameter.InvalidTime"

	// 权重不合法。请输入0~100的整数。
	INVALIDPARAMETER_INVALIDWEIGHT = "InvalidParameter.InvalidWeight"

	// IP已经存在。
	INVALIDPARAMETER_IPALREADYEXIST = "InvalidParameter.IpAlreadyExist"

	// 删除自定义线路失败，原因：线路不存在或者已删除。
	INVALIDPARAMETER_IPAREA = "InvalidParameter.IpArea"

	// ips 过长。
	INVALIDPARAMETER_IPSEXCEEDLIMIT = "InvalidParameter.IpsExceedLimit"

	// 单次任务数量超过上限。
	INVALIDPARAMETER_JOBGREATERTHANLIMIT = "InvalidParameter.JobGreaterThanLimit"

	// 线路正在使用当中，无法修改名称。
	INVALIDPARAMETER_LINEINUSE = "InvalidParameter.LineInUse"

	// 线路名称的长度不能超过17个字符。
	INVALIDPARAMETER_LINENAMEINVALID = "InvalidParameter.LineNameInvalid"

	// 线路名称包含不被接受的字符。
	INVALIDPARAMETER_LINENAMEINVALIDCHARACTER = "InvalidParameter.LineNameInvalidCharacter"

	// 线路名是系统内置线路或用户自定义分组线路，无法使用该名称。
	INVALIDPARAMETER_LINENAMEOCCUPIED = "InvalidParameter.LineNameOccupied"

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

	// 新套餐类型无效。
	INVALIDPARAMETER_NEWPACKAGETYPEINVALID = "InvalidParameter.NewPackageTypeInvalid"

	// 分页起始数量错误。
	INVALIDPARAMETER_OFFSETINVALID = "InvalidParameter.OffsetInvalid"

	// 无效的openid。
	INVALIDPARAMETER_OPENIDINVALID = "InvalidParameter.OpenidInvalid"

	// 操作失败，请稍后再试。
	INVALIDPARAMETER_OPERATEFAILED = "InvalidParameter.OperateFailed"

	// 当前操作过于频繁，请 1 分钟后重试。
	INVALIDPARAMETER_OPERATIONISTOOFREQUENT = "InvalidParameter.OperationIsTooFrequent"

	// 不支持的操作类型。
	INVALIDPARAMETER_OPTYPENOTSUPPORTED = "InvalidParameter.OptypeNotSupported"

	// 对方账号未实名认证，请先完成实名认证再操作。
	INVALIDPARAMETER_OTHERACCOUNTUNREALNAME = "InvalidParameter.OtherAccountUnrealName"

	// param 格式错误。
	INVALIDPARAMETER_PARAMINVALID = "InvalidParameter.ParamInvalid"

	// 参数不合法，请求被拒绝。
	INVALIDPARAMETER_PARAMSILLEGAL = "InvalidParameter.ParamsIllegal"

	// 参数错误。
	INVALIDPARAMETER_PARAMSMISSING = "InvalidParameter.ParamsMissing"

	// 鉴权失败。
	INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"

	// 用户UIN无效。
	INVALIDPARAMETER_QCLOUDUININVALID = "InvalidParameter.QcloudUinInvalid"

	// TXT 记录无法匹配，请确认记录值是否准确，并再次验证。
	INVALIDPARAMETER_QUHUITXTNOTMATCH = "InvalidParameter.QuhuiTxtNotMatch"

	// TXT 记录未设置或暂未生效，请稍后重试。
	INVALIDPARAMETER_QUHUITXTRECORDWAIT = "InvalidParameter.QuhuiTxtRecordWait"

	// 已实名用户。
	INVALIDPARAMETER_REALNAMEUSER = "InvalidParameter.RealNameUser"

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

	// 共享用户中包含未实名认证用户。
	INVALIDPARAMETER_SHAREDUSERSUNREALNAME = "InvalidParameter.SharedUsersUnrealName"

	// 状态代码不正确。
	INVALIDPARAMETER_STATUSCODEINVALID = "InvalidParameter.StatusCodeInvalid"

	// 子域名不正确。
	INVALIDPARAMETER_SUBDOMAININVALID = "InvalidParameter.SubdomainInvalid"

	// 任务未完成。
	INVALIDPARAMETER_TASKNOTCOMPLETED = "InvalidParameter.TaskNotCompleted"

	// 时长无效。
	INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"

	// 时间戳已过期。
	INVALIDPARAMETER_TIMESTAMPEXPIRED = "InvalidParameter.TimestampExpired"

	// 当前账号下的无效域名过多，暂时无法使用该功能。请将已有域名的 DNS 服务器正确指向 DNSPod 之后，再尝试添加。
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

	// 账号已经被锁定。
	INVALIDPARAMETER_USERALREADYLOCKED = "InvalidParameter.UserAlreadyLocked"

	// 对方账户非国内站用户。
	INVALIDPARAMETER_USERAREAINVALID = "InvalidParameter.UserAreaInvalid"

	// 用户不存在。
	INVALIDPARAMETER_USERNOTEXISTS = "InvalidParameter.UserNotExists"

	// 无效的uuid。
	INVALIDPARAMETER_UUIDINVALID = "InvalidParameter.UuidInvalid"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 域名等级不正确。
	INVALIDPARAMETERVALUE_DOMAINGRADEINVALID = "InvalidParameterValue.DomainGradeInvalid"

	// 当前域名有误，请返回重新操作。
	INVALIDPARAMETERVALUE_DOMAINNOTEXISTS = "InvalidParameterValue.DomainNotExists"

	// 自定义线路IP段不能为空。
	INVALIDPARAMETERVALUE_IPAREAEMPTYIP = "InvalidParameterValue.IpAreaEmptyIp"

	// 自定义线路名称不能为空。
	INVALIDPARAMETERVALUE_IPAREAEMPTYLINENAME = "InvalidParameterValue.IpAreaEmptyLineName"

	// 分页长度数量错误。
	INVALIDPARAMETERVALUE_LIMITINVALID = "InvalidParameterValue.LimitInvalid"

	// 升级年限不正确。
	INVALIDPARAMETERVALUE_UPGRADETERMINVALID = "InvalidParameterValue.UpgradeTermInvalid"

	// 用户编号不正确。
	INVALIDPARAMETERVALUE_USERIDINVALID = "InvalidParameterValue.UserIdInvalid"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// AAAA记录数量超出限制。
	LIMITEXCEEDED_AAAACOUNTLIMIT = "LimitExceeded.AAAACountLimit"

	// @的NS记录只能设置为默认线路。
	LIMITEXCEEDED_ATNSRECORDLIMIT = "LimitExceeded.AtNsRecordLimit"

	// 抱歉，自定义线路个数超过限制，请联系销售进行购买。
	LIMITEXCEEDED_CUSTOMLINELIMITED = "LimitExceeded.CustomLineLimited"

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

	// 分页起始数量过大。
	LIMITEXCEEDED_OFFSETEXCEEDED = "LimitExceeded.OffsetExceeded"

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

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 您没有权限执行此操作。
	OPERATIONDENIED_ACCESSDENIED = "OperationDenied.AccessDenied"

	// 代理不能使用此功能。
	OPERATIONDENIED_AGENTDENIED = "OperationDenied.AgentDenied"

	// 代理名下的用户不能使用此功能。
	OPERATIONDENIED_AGENTSUBORDINATEDENIED = "OperationDenied.AgentSubordinateDenied"

	// 此订单不能取消。
	OPERATIONDENIED_CANCELBILLNOTALLOWED = "OperationDenied.CancelBillNotAllowed"

	// 该线路正在使用当中，不能删除。
	OPERATIONDENIED_DELETEUSINGRECORDLINENOTALLOWED = "OperationDenied.DeleteUsingRecordLineNotAllowed"

	// 仅域名所有者可进行此操作。
	OPERATIONDENIED_DOMAINOWNERALLOWEDONLY = "OperationDenied.DomainOwnerAllowedOnly"

	// 该线路正在使用当中，不能编辑。
	OPERATIONDENIED_EDITUSINGRECORDLINENOTALLOWED = "OperationDenied.EditUsingRecordLineNotAllowed"

	// 抱歉，不允许添加黑名单中的IP。
	OPERATIONDENIED_IPINBLACKLISTNOTALLOWED = "OperationDenied.IPInBlacklistNotAllowed"

	// 抱歉，您的域名等级不支持D监控通知回调功能。
	OPERATIONDENIED_MONITORCALLBACKNOTENABLED = "OperationDenied.MonitorCallbackNotEnabled"

	// 当前域名无权限，请返回域名列表。
	OPERATIONDENIED_NOPERMISSIONTOOPERATEDOMAIN = "OperationDenied.NoPermissionToOperateDomain"

	// 您不是管理用户。
	OPERATIONDENIED_NOTADMIN = "OperationDenied.NotAdmin"

	// 您不是代理用户。
	OPERATIONDENIED_NOTAGENT = "OperationDenied.NotAgent"

	// 您还没有获取到授权，无法执行此操作。
	OPERATIONDENIED_NOTGRANTEDBYOWNER = "OperationDenied.NotGrantedByOwner"

	// 不是您名下用户。
	OPERATIONDENIED_NOTMANAGEDUSER = "OperationDenied.NotManagedUser"

	// 您没有权限操作此订单。
	OPERATIONDENIED_NOTORDEROWNER = "OperationDenied.NotOrderOwner"

	// 您没有权限操作此资源。
	OPERATIONDENIED_NOTRESOURCEOWNER = "OperationDenied.NotResourceOwner"

	// 此礼券为个人礼券，请使用企业礼券。
	OPERATIONDENIED_PERSONALCOUPONNOTALLOWED = "OperationDenied.PersonalCouponNotAllowed"

	// 只支持 POST 方法提交数据。
	OPERATIONDENIED_POSTREQUESTACCEPTONLY = "OperationDenied.PostRequestAcceptOnly"

	// 该资源不允许续费。
	OPERATIONDENIED_RESOURCENOTALLOWRENEW = "OperationDenied.ResourceNotAllowRenew"

	// 企业用户的域名需要升级到VIP才能解析。
	OPERATIONDENIED_VIPDOMAINALLOWED = "OperationDenied.VipDomainAllowed"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 您的IP添加了过多任务，请稍后重试。
	REQUESTLIMITEXCEEDED_BATCHTASKLIMIT = "RequestLimitExceeded.BatchTaskLimit"

	// 您的账号在短时间内添加了大量的域名，请控制添加频率。
	REQUESTLIMITEXCEEDED_CREATEDOMAINLIMIT = "RequestLimitExceeded.CreateDomainLimit"

	// API请求次数超出限制。
	REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDED = "RequestLimitExceeded.RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 域名列表为空。
	RESOURCENOTFOUND_NODATAOFDOMAIN = "ResourceNotFound.NoDataOfDomain"

	// 没有域名别名。
	RESOURCENOTFOUND_NODATAOFDOMAINALIAS = "ResourceNotFound.NoDataOfDomainAlias"

	// 还没有任何礼券。
	RESOURCENOTFOUND_NODATAOFGIFT = "ResourceNotFound.NoDataOfGift"

	// 记录列表为空。
	RESOURCENOTFOUND_NODATAOFRECORD = "ResourceNotFound.NoDataOfRecord"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
)
