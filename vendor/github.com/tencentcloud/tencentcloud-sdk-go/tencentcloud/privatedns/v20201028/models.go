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

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountVpcInfo struct {

	// VpcId： vpc-xadsafsdasd
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Vpc所属地区: ap-guangzhou, ap-shanghai
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// Vpc所属账号: 123456789
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// vpc资源名称：testname
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type AccountVpcInfoOut struct {

	// VpcId： vpc-xadsafsdasd
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Vpc所属地区: ap-guangzhou, ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`

	// Vpc所属账号: 123456789
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// vpc资源名称：testname
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type AccountVpcInfoOutput struct {

	// 关联账户的uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// vpcid
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`
}

type AuditLog struct {

	// 日志类型
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 日志表名
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 日志总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 日志列表
	DataSet []*AuditLogInfo `json:"DataSet,omitempty" name:"DataSet"`
}

type AuditLogInfo struct {

	// 时间
	Date *string `json:"Date,omitempty" name:"Date"`

	// 操作人uin
	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`

	// 日志内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreatePrivateDNSAccountRequest struct {
	*tchttp.BaseRequest

	// 私有域解析账号
	Account *PrivateDNSAccount `json:"Account,omitempty" name:"Account"`
}

func (r *CreatePrivateDNSAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateDNSAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateDNSAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateDNSAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrivateDNSAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateDNSAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateZoneRecordRequest struct {
	*tchttp.BaseRequest

	// 私有域ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 子域名，例如 "www", "m", "@"
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 记录值，例如 IP：192.168.10.2，CNAME：cname.qcloud.com.，MX：mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitempty" name:"MX"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`
}

func (r *CreatePrivateZoneRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordType")
	delete(f, "SubDomain")
	delete(f, "RecordValue")
	delete(f, "Weight")
	delete(f, "MX")
	delete(f, "TTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录Id
		RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrivateZoneRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateZoneRequest struct {
	*tchttp.BaseRequest

	// 域名，格式必须是标准的TLD
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 创建私有域的同时，为其打上标签
	TagSet []*TagInfo `json:"TagSet,omitempty" name:"TagSet"`

	// 创建私有域的同时，将其关联至VPC
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 是否开启子域名递归, ENABLED， DISABLED。默认值为DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`

	// 创建私有域的同时，将其关联至VPC
	Vpcs []*VpcInfo `json:"Vpcs,omitempty" name:"Vpcs"`

	// 创建私有域同时绑定关联账号的VPC
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
}

func (r *CreatePrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "TagSet")
	delete(f, "VpcSet")
	delete(f, "Remark")
	delete(f, "DnsForwardStatus")
	delete(f, "Vpcs")
	delete(f, "AccountVpcSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有域ID, zone-xxxxxx
		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

		// 私有域名
		Domain *string `json:"Domain,omitempty" name:"Domain"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatePoint struct {

	// 时间
	Date *string `json:"Date,omitempty" name:"Date"`

	// 值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DeletePrivateDNSAccountRequest struct {
	*tchttp.BaseRequest

	// 私有域解析账号
	Account *PrivateDNSAccount `json:"Account,omitempty" name:"Account"`
}

func (r *DeletePrivateDNSAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateDNSAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateDNSAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateDNSAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivateDNSAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateDNSAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateZoneRecordRequest struct {
	*tchttp.BaseRequest

	// 私有域ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 记录ID
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 记录ID数组，RecordId 优先
	RecordIdSet []*string `json:"RecordIdSet,omitempty" name:"RecordIdSet"`
}

func (r *DeletePrivateZoneRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordId")
	delete(f, "RecordIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivateZoneRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateZoneRequest struct {
	*tchttp.BaseRequest

	// 私有域ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 私有域ID数组，ZoneId 优先
	ZoneIdSet []*string `json:"ZoneIdSet,omitempty" name:"ZoneIdSet"`
}

func (r *DeletePrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ZoneIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountVpcListRequest struct {
	*tchttp.BaseRequest

	// 关联账号的uin
	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAccountVpcListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountVpcListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountVpcListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountVpcListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// VPC数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// VPC 列表
		VpcSet []*AccountVpcInfoOut `json:"VpcSet,omitempty" name:"VpcSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountVpcListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountVpcListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditLogRequest struct {
	*tchttp.BaseRequest

	// 请求量统计起始时间
	TimeRangeBegin *string `json:"TimeRangeBegin,omitempty" name:"TimeRangeBegin"`

	// 筛选参数：ZoneId：私有域ID；Domain：私有域；OperatorUin：操作者账号ID
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 请求量统计结束时间
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAuditLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeRangeBegin")
	delete(f, "Filters")
	delete(f, "TimeRangeEnd")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作日志列表
		Data []*AuditLog `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有域解析总数
		ZoneTotal *int64 `json:"ZoneTotal,omitempty" name:"ZoneTotal"`

		// 私有域关联VPC数量
		ZoneVpcCount *int64 `json:"ZoneVpcCount,omitempty" name:"ZoneVpcCount"`

		// 历史请求量总数
		RequestTotalCount *int64 `json:"RequestTotalCount,omitempty" name:"RequestTotalCount"`

		// 流量包用量
		FlowUsage []*FlowUsage `json:"FlowUsage,omitempty" name:"FlowUsage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateDNSAccountListRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrivateDNSAccountListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateDNSAccountListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateDNSAccountListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateDNSAccountListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有域解析账号数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 私有域解析账号列表
		AccountSet []*PrivateDNSAccount `json:"AccountSet,omitempty" name:"AccountSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateDNSAccountListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateDNSAccountListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneListRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrivateZoneListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有域数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 私有域列表
		PrivateZoneSet []*PrivateZone `json:"PrivateZoneSet,omitempty" name:"PrivateZoneSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneRecordListRequest struct {
	*tchttp.BaseRequest

	// 私有域ID: zone-xxxxxx
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 过滤参数（支持使用Value、RecordType过滤）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页偏移量，从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页限制数目， 最大100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePrivateZoneRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneRecordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneRecordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneRecordListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 解析记录数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 解析记录列表
		RecordSet []*PrivateZoneRecord `json:"RecordSet,omitempty" name:"RecordSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneRequest struct {
	*tchttp.BaseRequest

	// 域名，格式必须是标准的TLD
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribePrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有域详情
		PrivateZone *PrivateZone `json:"PrivateZone,omitempty" name:"PrivateZone"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePrivateZoneServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有域解析服务开通状态。ENABLED已开通，DISABLED未开通
		ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRequestDataRequest struct {
	*tchttp.BaseRequest

	// 请求量统计起始时间，格式：2020-11-22 00:00:00
	TimeRangeBegin *string `json:"TimeRangeBegin,omitempty" name:"TimeRangeBegin"`

	// 筛选参数：
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 请求量统计结束时间，格式：2020-11-22 23:59:59
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`
}

func (r *DescribeRequestDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRequestDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeRangeBegin")
	delete(f, "Filters")
	delete(f, "TimeRangeEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRequestDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRequestDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求量统计表
		Data []*MetricData `json:"Data,omitempty" name:"Data"`

		// 请求量单位时间: Day：天，Hour：小时
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRequestDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRequestDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 参数名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值数组
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FlowUsage struct {

	// 流量包类型：ZONE 私有域；TRAFFIC 解析流量包
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 流量包总额度
	TotalQuantity *int64 `json:"TotalQuantity,omitempty" name:"TotalQuantity"`

	// 流量包可用额度
	AvailableQuantity *int64 `json:"AvailableQuantity,omitempty" name:"AvailableQuantity"`
}

type MetricData struct {

	// 资源描述
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 表名
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 表数据
	DataSet []*DatePoint `json:"DataSet,omitempty" name:"DataSet"`
}

type ModifyPrivateZoneRecordRequest struct {
	*tchttp.BaseRequest

	// 私有域ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 记录ID
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 子域名，例如 "www", "m", "@"
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 记录值，例如 IP：192.168.10.2，CNAME：cname.qcloud.com.，MX：mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`

	// 记录权重，值为1-100
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	MX *int64 `json:"MX,omitempty" name:"MX"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`
}

func (r *ModifyPrivateZoneRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordId")
	delete(f, "RecordType")
	delete(f, "SubDomain")
	delete(f, "RecordValue")
	delete(f, "Weight")
	delete(f, "MX")
	delete(f, "TTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateZoneRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneRequest struct {
	*tchttp.BaseRequest

	// 私有域ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 是否开启子域名递归, ENABLED， DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
}

func (r *ModifyPrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Remark")
	delete(f, "DnsForwardStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneVpcRequest struct {
	*tchttp.BaseRequest

	// 私有域ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 私有域关联的全部VPC列表
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// 私有域账号关联的全部VPC列表
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
}

func (r *ModifyPrivateZoneVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "VpcSet")
	delete(f, "AccountVpcSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有域ID, zone-xxxxxx
		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

		// 解析域关联的VPC列表
		VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

		// 私有域账号关联的全部VPC列表
		AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateZoneVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrivateDNSAccount struct {

	// 主账号Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 主账号名称
	Account *string `json:"Account,omitempty" name:"Account"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
}

type PrivateZone struct {

	// 私有域id: zone-xxxxxxxx
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 域名所有者uin
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 私有域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 修改时间
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// 记录数
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 绑定的Vpc列表
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// 私有域状态：正常解析：ENABLED, 暂停解析：SUSPEND, 锁定：FROZEN
	Status *string `json:"Status,omitempty" name:"Status"`

	// 域名递归解析状态：开通：ENABLED, 关闭，DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`

	// 标签键值对集合
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 绑定的关联账号的vpc列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
}

type PrivateZoneRecord struct {

	// 记录id
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 私有域id: zone-xxxxxxxx
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 子域名
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 记录类型，可选的记录类型为："A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 记录值
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`

	// 记录缓存时间，数值越小生效越快，取值1-86400s, 默认 600
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// MX优先级：记录类型为MX时必填。取值范围：5,10,15,20,30,40,50
	// 注意：此字段可能返回 null，表示取不到有效值。
	MX *int64 `json:"MX,omitempty" name:"MX"`

	// 记录状态：ENABLED
	Status *string `json:"Status,omitempty" name:"Status"`

	// 记录权重，值为1-100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 记录创建时间
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 记录更新时间
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

type SubscribePrivateZoneServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *SubscribePrivateZoneServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubscribePrivateZoneServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubscribePrivateZoneServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SubscribePrivateZoneServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有域解析服务开通状态
		ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubscribePrivateZoneServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubscribePrivateZoneServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type VpcInfo struct {

	// VpcId： vpc-xadsafsdasd
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Vpc所属地区: ap-guangzhou, ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}
