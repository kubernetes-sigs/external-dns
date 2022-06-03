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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddRecordBatch struct {

	// 记录类型, 详见 DescribeRecordType 接口。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 记录值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 子域名(主机记录)，默认为@。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 解析记录的线路，详见 DescribeRecordLineList 接口，RecordLine和RecordLineId都未填时，默认为「默认」线路。
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// 解析记录的线路 ID，RecordLine和RecordLineId都有时，系统优先取 RecordLineId。
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// 记录权重值(暂未支持)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 记录的 MX 记录值，非 MX 记录类型，默认为 0，MX记录则必选。
	MX *uint64 `json:"MX,omitempty" name:"MX"`

	// 记录的 TTL 值，默认600。
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 记录状态(暂未支持)。0表示禁用，1表示启用。默认启用。
	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`

	// 记录备注(暂未支持)。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type BatchRecordInfo struct {

	// 记录 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// 子域名(主机记录)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 记录类型, 详见 DescribeRecordType 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 解析记录的线路，详见 DescribeRecordLineList 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// 记录值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 记录的 TTL 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 记录添加状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 此条记录在列表中的ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 记录生效状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`

	// 记录的MX权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	MX *uint64 `json:"MX,omitempty" name:"MX"`
}

type CreateDomainAliasRequest struct {
	*tchttp.BaseRequest

	// 域名别名
	DomainAlias *string `json:"DomainAlias,omitempty" name:"DomainAlias"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID，参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain
	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *CreateDomainAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainAlias")
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名别名ID
		DomainAliasId *int64 `json:"DomainAliasId,omitempty" name:"DomainAliasId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainBatchDetail struct {

	// 见RecordInfoBatch
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordList []*CreateDomainBatchRecord `json:"RecordList,omitempty" name:"RecordList"`

	// 任务编号
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainGrade *string `json:"DomainGrade,omitempty" name:"DomainGrade"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 该条任务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

type CreateDomainBatchRecord struct {

	// 子域名(主机记录)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 记录类型, 详见 DescribeRecordType 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 解析记录的线路，详见 DescribeRecordLineList 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// 记录值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 记录的 TTL 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 记录添加状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 此条记录在列表中的ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type CreateDomainBatchRequest struct {
	*tchttp.BaseRequest

	// 域名数组
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList"`

	// 每个域名添加 @ 和 www 的 A 记录值，记录值为IP，如果不传此参数或者传空，将只添加域名，不添加记录。
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`
}

func (r *CreateDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainList")
	delete(f, "RecordValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批量添加域名信息
		DetailList []*CreateDomainBatchDetail `json:"DetailList,omitempty" name:"DetailList"`

		// 批量任务的ID
		JobId *uint64 `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainGroupRequest struct {
	*tchttp.BaseRequest

	// 域名分组
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

func (r *CreateDomainGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名分组ID
		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名分组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 是否星标域名，”yes”、”no” 分别代表是和否。
	IsMark *string `json:"IsMark,omitempty" name:"IsMark"`
}

func (r *CreateDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "GroupId")
	delete(f, "IsMark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名信息
		DomainInfo *DomainCreateInfo `json:"DomainInfo,omitempty" name:"DomainInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordBatchDetail struct {

	// 见RecordInfoBatch
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordList []*CreateRecordBatchRecord `json:"RecordList,omitempty" name:"RecordList"`

	// 任务编号
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainGrade *string `json:"DomainGrade,omitempty" name:"DomainGrade"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 该条任务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 域名ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

type CreateRecordBatchRecord struct {

	// 子域名(主机记录)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 记录类型, 详见 DescribeRecordType 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 解析记录的线路，详见 DescribeRecordLineList 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// 记录值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 记录的 TTL 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 记录添加状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 此条记录在列表中的ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 记录的MX权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	MX *uint64 `json:"MX,omitempty" name:"MX"`
}

type CreateRecordBatchRequest struct {
	*tchttp.BaseRequest

	// 域名ID，多个 domain_id 用英文逗号进行分割。
	DomainIdList []*string `json:"DomainIdList,omitempty" name:"DomainIdList"`

	// 记录数组
	RecordList []*AddRecordBatch `json:"RecordList,omitempty" name:"RecordList"`
}

func (r *CreateRecordBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainIdList")
	delete(f, "RecordList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecordBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批量添加域名信息
		DetailList []*CreateRecordBatchDetail `json:"DetailList,omitempty" name:"DetailList"`

		// 批量任务的ID
		JobId *uint64 `json:"JobId,omitempty" name:"JobId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRecordBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 记录类型，通过 API 记录类型获得，大写英文，比如：A 。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 记录线路，通过 API 记录线路获得，中文，比如：默认。
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// 记录值，如 IP : 200.200.200.200， CNAME : cname.dnspod.com.， MX : mail.dnspod.com.。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 线路的 ID，通过 API 记录线路获得，英文字符串，比如：10=1。参数RecordLineId优先级高于RecordLine，如果同时传递二者，优先使用RecordLineId参数。
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// MX 优先级，当记录类型是 MX 时有效，范围1-20，MX 记录时必选。
	MX *uint64 `json:"MX,omitempty" name:"MX"`

	// TTL，范围1-604800，不同等级域名最小值不同。
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 权重信息，0到100的整数。仅企业 VIP 域名可用，0 表示关闭，不传该参数，表示不设置权重信息。
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 记录初始状态，取值范围为 ENABLE 和 DISABLE 。默认为 ENABLE ，如果传入 DISABLE，解析不会生效，也不会验证负载均衡的限制。
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *CreateRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordType")
	delete(f, "RecordLine")
	delete(f, "Value")
	delete(f, "DomainId")
	delete(f, "SubDomain")
	delete(f, "RecordLineId")
	delete(f, "MX")
	delete(f, "TTL")
	delete(f, "Weight")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录ID
		RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainAliasRequest struct {
	*tchttp.BaseRequest

	// 域名别名ID
	DomainAliasId *int64 `json:"DomainAliasId,omitempty" name:"DomainAliasId"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID，参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain
	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DeleteDomainAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainAliasId")
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDomainAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DeleteDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 记录 ID 。
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DeleteRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShareDomainRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名共享的账号
	Account *string `json:"Account,omitempty" name:"Account"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DeleteShareDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Account")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShareDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShareDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteShareDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchTaskDetail struct {

	// 见BatchRecordInfo
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordList []*BatchRecordInfo `json:"RecordList,omitempty" name:"RecordList"`

	// 任务编号
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainGrade *string `json:"DomainGrade,omitempty" name:"DomainGrade"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 该条任务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 域名ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

type DescribeBatchTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	JobId *uint64 `json:"JobId,omitempty" name:"JobId"`
}

func (r *DescribeBatchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批量任务详情
		DetailList []*DescribeBatchTaskDetail `json:"DetailList,omitempty" name:"DetailList"`

		// 总任务条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 成功条数
		SuccessCount *uint64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

		// 失败条数
		FailCount *uint64 `json:"FailCount,omitempty" name:"FailCount"`

		// 批量任务类型
		JobType *string `json:"JobType,omitempty" name:"JobType"`

		// 任务创建时间
		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBatchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainAliasListRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名ID,域名ID，参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain
	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeDomainAliasListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainAliasListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainAliasListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainAliasListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名别名列表
		DomainAliasList []*DomainAliasInfo `json:"DomainAliasList,omitempty" name:"DomainAliasList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainAliasListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainAliasListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainListRequest struct {
	*tchttp.BaseRequest

	// 域名分组类型，默认为ALL。可取值为ALL，MINE，SHARE，ISMARK，PAUSE，VIP，RECENT，SHARE_OUT。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 记录开始的偏移, 第一条记录为 0, 依次类推。默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 要获取的域名数量, 比如获取20个, 则为20。默认值为3000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分组ID, 获取指定分组的域名
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 根据关键字搜索域名
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeDomainListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupId")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表页统计信息
		DomainCountInfo *DomainCountInfo `json:"DomainCountInfo,omitempty" name:"DomainCountInfo"`

		// 域名列表
		DomainList []*DomainListItem `json:"DomainList,omitempty" name:"DomainList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainLogListRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 记录开始的偏移，第一条记录为 0，依次类推，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 共要获取的日志条数，比如获取20条，则为20，默认为500条，单次最多获取500条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDomainLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainLogListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		LogList []*string `json:"LogList,omitempty" name:"LogList"`

		// 分页大小
		PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

		// 日志总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainPurviewRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeDomainPurviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainPurviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainPurviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainPurviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名权限列表
		PurviewList []*PurviewInfo `json:"PurviewList,omitempty" name:"PurviewList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainPurviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainPurviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名信息
		DomainInfo *DomainInfo `json:"DomainInfo,omitempty" name:"DomainInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainShareInfoRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeDomainShareInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainShareInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainShareInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainShareInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名共享信息
		ShareList []*DomainShareInfo `json:"ShareList,omitempty" name:"ShareList"`

		// 域名拥有者账号
		Owner *string `json:"Owner,omitempty" name:"Owner"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainShareInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainShareInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordLineListRequest struct {
	*tchttp.BaseRequest

	// 域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名等级。
	// + 旧套餐：D_FREE、D_PLUS、D_EXTRA、D_EXPERT、D_ULTRA 分别对应免费套餐、个人豪华、企业1、企业2、企业3。
	// + 新套餐：DP_FREE、DP_PLUS、DP_EXTRA、DP_EXPERT、DP_ULTRA 分别对应新免费、个人专业版、企业创业版、企业标准版、企业旗舰版。
	DomainGrade *string `json:"DomainGrade,omitempty" name:"DomainGrade"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeRecordLineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordLineListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainGrade")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordLineListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordLineListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 线路列表。
		LineList []*LineInfo `json:"LineList,omitempty" name:"LineList"`

		// 线路分组列表。
		LineGroupList []*LineGroupInfo `json:"LineGroupList,omitempty" name:"LineGroupList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordLineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordLineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordListRequest struct {
	*tchttp.BaseRequest

	// 要获取的解析记录所属的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 要获取的解析记录所属的域名Id，如果传了DomainId，系统将会忽略Domain参数
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 解析记录的主机头，如果传了此参数，则只会返回此主机头对应的解析记录
	Subdomain *string `json:"Subdomain,omitempty" name:"Subdomain"`

	// 获取某种类型的解析记录，如 A，CNAME，NS，AAAA，显性URL，隐性URL，CAA，SPF等
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 获取某条线路名称的解析记录
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// 获取某个线路Id对应的解析记录，如果传RecordLineId，系统会忽略RecordLine参数
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// 获取某个分组下的解析记录时，传这个分组Id
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 通过关键字搜索解析记录，当前支持搜索主机头和记录值
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 排序字段，支持 name,line,type,value,weight,mx,ttl,updated_on 几个字段。
	SortField *string `json:"SortField,omitempty" name:"SortField"`

	// 排序方式，正序：ASC，逆序：DESC。默认值为ASC。
	SortType *string `json:"SortType,omitempty" name:"SortType"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数量，当前Limit最大支持3000。默认值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Subdomain")
	delete(f, "RecordType")
	delete(f, "RecordLine")
	delete(f, "RecordLineId")
	delete(f, "GroupId")
	delete(f, "Keyword")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录的数量统计信息
		RecordCountInfo *RecordCountInfo `json:"RecordCountInfo,omitempty" name:"RecordCountInfo"`

		// 获取的记录列表
		RecordList []*RecordListItem `json:"RecordList,omitempty" name:"RecordList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 记录 ID 。
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录信息
		RecordInfo *RecordInfo `json:"RecordInfo,omitempty" name:"RecordInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordTypeRequest struct {
	*tchttp.BaseRequest

	// 域名等级。
	// + 旧套餐：D_FREE、D_PLUS、D_EXTRA、D_EXPERT、D_ULTRA 分别对应免费套餐、个人豪华、企业1、企业2、企业3。
	// + 新套餐：DP_FREE、DP_PLUS、DP_EXTRA、DP_EXPERT、DP_ULTRA 分别对应新免费、个人专业版、企业创业版、企业标准版、企业旗舰版。
	DomainGrade *string `json:"DomainGrade,omitempty" name:"DomainGrade"`
}

func (r *DescribeRecordTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainGrade")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录类型列表
		TypeList []*string `json:"TypeList,omitempty" name:"TypeList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDetailRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 帐户信息
		UserInfo *UserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainAliasInfo struct {

	// 域名别名ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 域名别名
	DomainAlias *string `json:"DomainAlias,omitempty" name:"DomainAlias"`
}

type DomainCountInfo struct {

	// 符合条件的域名数量
	DomainTotal *uint64 `json:"DomainTotal,omitempty" name:"DomainTotal"`

	// 用户可以查看的所有域名数量
	AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

	// 用户账号添加的域名数量
	MineTotal *uint64 `json:"MineTotal,omitempty" name:"MineTotal"`

	// 共享给用户的域名数量
	ShareTotal *uint64 `json:"ShareTotal,omitempty" name:"ShareTotal"`

	// 付费域名数量
	VipTotal *uint64 `json:"VipTotal,omitempty" name:"VipTotal"`

	// 暂停的域名数量
	PauseTotal *uint64 `json:"PauseTotal,omitempty" name:"PauseTotal"`

	// dns设置错误的域名数量
	ErrorTotal *uint64 `json:"ErrorTotal,omitempty" name:"ErrorTotal"`

	// 锁定的域名数量
	LockTotal *uint64 `json:"LockTotal,omitempty" name:"LockTotal"`

	// 封禁的域名数量
	SpamTotal *uint64 `json:"SpamTotal,omitempty" name:"SpamTotal"`

	// 30天内即将到期的域名数量
	VipExpire *uint64 `json:"VipExpire,omitempty" name:"VipExpire"`

	// 分享给其它人的域名数量
	ShareOutTotal *uint64 `json:"ShareOutTotal,omitempty" name:"ShareOutTotal"`

	// 指定分组内的域名数量
	GroupTotal *uint64 `json:"GroupTotal,omitempty" name:"GroupTotal"`
}

type DomainCreateInfo struct {

	// 域名ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名的punycode
	Punycode *string `json:"Punycode,omitempty" name:"Punycode"`

	// 域名的NS列表
	GradeNsList []*string `json:"GradeNsList,omitempty" name:"GradeNsList"`
}

type DomainInfo struct {

	// 域名ID
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 域名状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 域名套餐等级
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 域名分组ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 是否星标域名
	IsMark *string `json:"IsMark,omitempty" name:"IsMark"`

	// TTL(DNS记录缓存时间)
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// cname加速启用状态
	CnameSpeedup *string `json:"CnameSpeedup,omitempty" name:"CnameSpeedup"`

	// 域名备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 域名Punycode
	Punycode *string `json:"Punycode,omitempty" name:"Punycode"`

	// 域名DNS状态
	DnsStatus *string `json:"DnsStatus,omitempty" name:"DnsStatus"`

	// 域名的NS列表
	DnspodNsList []*string `json:"DnspodNsList,omitempty" name:"DnspodNsList"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名等级代号
	GradeLevel *uint64 `json:"GradeLevel,omitempty" name:"GradeLevel"`

	// 域名所属的用户ID
	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`

	// 是否为付费域名
	IsVip *string `json:"IsVip,omitempty" name:"IsVip"`

	// 域名所有者的账号
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 域名等级的描述
	GradeTitle *string `json:"GradeTitle,omitempty" name:"GradeTitle"`

	// 域名创建时间
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 最后操作时间
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// 腾讯云账户Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 域名实际使用的NS列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActualNsList []*string `json:"ActualNsList,omitempty" name:"ActualNsList"`

	// 域名的记录数量
	RecordCount *uint64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// 域名所有者的账户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerNick *string `json:"OwnerNick,omitempty" name:"OwnerNick"`
}

type DomainListItem struct {

	// 系统分配给域名的唯一标识
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 域名的原始格式
	Name *string `json:"Name,omitempty" name:"Name"`

	// 域名的状态，正常：ENABLE，暂停：PAUSE，封禁：SPAM
	Status *string `json:"Status,omitempty" name:"Status"`

	// 域名默认的解析记录默认TTL值
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 是否开启CNAME加速，开启：ENABLE，未开启：DISABLE
	CNAMESpeedup *string `json:"CNAMESpeedup,omitempty" name:"CNAMESpeedup"`

	// DNS 设置状态，错误：DNSERROR，正常：空字符串
	DNSStatus *string `json:"DNSStatus,omitempty" name:"DNSStatus"`

	// 域名的套餐等级代码
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 域名所属的分组Id
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 是否开启搜索引擎推送优化，是：YES，否：NO
	SearchEnginePush *string `json:"SearchEnginePush,omitempty" name:"SearchEnginePush"`

	// 域名备注说明
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 经过punycode编码后的域名格式
	Punycode *string `json:"Punycode,omitempty" name:"Punycode"`

	// 系统为域名分配的有效DNS
	EffectiveDNS []*string `json:"EffectiveDNS,omitempty" name:"EffectiveDNS"`

	// 域名套餐等级对应的序号
	GradeLevel *uint64 `json:"GradeLevel,omitempty" name:"GradeLevel"`

	// 套餐名称
	GradeTitle *string `json:"GradeTitle,omitempty" name:"GradeTitle"`

	// 是否是付费套餐
	IsVip *string `json:"IsVip,omitempty" name:"IsVip"`

	// 付费套餐开通时间
	VipStartAt *string `json:"VipStartAt,omitempty" name:"VipStartAt"`

	// 付费套餐到期时间
	VipEndAt *string `json:"VipEndAt,omitempty" name:"VipEndAt"`

	// 域名是否开通VIP自动续费，是：YES，否：NO，默认：DEFAULT
	VipAutoRenew *string `json:"VipAutoRenew,omitempty" name:"VipAutoRenew"`

	// 域名下的记录数量
	RecordCount *uint64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// 域名添加时间
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 域名更新时间
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// 域名所属账号
	Owner *string `json:"Owner,omitempty" name:"Owner"`
}

type DomainShareInfo struct {

	// 域名共享对象的账号
	ShareTo *string `json:"ShareTo,omitempty" name:"ShareTo"`

	// 共享模式，“rw”：可读写。 “r”:：只读
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 共享状态“enabled”：共享成功。“pending”：共享到的账号不存在, 等待注册
	Status *string `json:"Status,omitempty" name:"Status"`
}

type LineGroupInfo struct {

	// 线路分组ID
	LineId *string `json:"LineId,omitempty" name:"LineId"`

	// 线路分组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分组类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 线路分组包含的线路列表
	LineList []*string `json:"LineList,omitempty" name:"LineList"`
}

type LineInfo struct {

	// 线路名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 线路ID
	LineId *string `json:"LineId,omitempty" name:"LineId"`
}

type LockInfo struct {

	// 域名 ID
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 域名解锁码
	LockCode *string `json:"LockCode,omitempty" name:"LockCode"`

	// 域名自动解锁日期
	LockEnd *string `json:"LockEnd,omitempty" name:"LockEnd"`
}

type ModifyDomainLockRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名要锁定的天数，最多可锁定的天数可以通过获取域名权限接口获取。
	LockDays *uint64 `json:"LockDays,omitempty" name:"LockDays"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *ModifyDomainLockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainLockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "LockDays")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainLockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainLockResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名锁定信息
		LockInfo *LockInfo `json:"LockInfo,omitempty" name:"LockInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainLockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainLockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainOwnerRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名需要转入的账号，支持Uin或者邮箱格式
	Account *string `json:"Account,omitempty" name:"Account"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *ModifyDomainOwnerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainOwnerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Account")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainOwnerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainOwnerResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainOwnerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainOwnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainRemarkRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 域名备注，删除备注请提交空内容。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyDomainRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainStatusRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名状态，”enable” 、”disable” 分别代表启用和暂停
	Status *string `json:"Status,omitempty" name:"Status"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *ModifyDomainStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Status")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainUnlockRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名解锁码，锁定的时候会返回。
	LockCode *string `json:"LockCode,omitempty" name:"LockCode"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *ModifyDomainUnlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainUnlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "LockCode")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainUnlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainUnlockResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainUnlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainUnlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDynamicDNSRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 记录ID。
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// 记录线路，通过 API 记录线路获得，中文，比如：默认。
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// 记录值，如 IP : 200.200.200.200， CNAME : cname.dnspod.com.， MX : mail.dnspod.com.。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 线路的 ID，通过 API 记录线路获得，英文字符串，比如：10=1。参数RecordLineId优先级高于RecordLine，如果同时传递二者，优先使用RecordLineId参数。
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`
}

func (r *ModifyDynamicDNSRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDynamicDNSRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordId")
	delete(f, "RecordLine")
	delete(f, "Value")
	delete(f, "DomainId")
	delete(f, "SubDomain")
	delete(f, "RecordLineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDynamicDNSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDynamicDNSResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录ID
		RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDynamicDNSResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDynamicDNSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordBatchDetail struct {

	// 见RecordInfoBatchModify
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordList []*BatchRecordInfo `json:"RecordList,omitempty" name:"RecordList"`

	// 任务编号
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainGrade *string `json:"DomainGrade,omitempty" name:"DomainGrade"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 该条任务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 域名ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

type ModifyRecordBatchRequest struct {
	*tchttp.BaseRequest

	// 记录ID数组
	RecordIdList []*uint64 `json:"RecordIdList,omitempty" name:"RecordIdList"`

	// 要修改的字段，可选值为 [“sub_domain”、”record_type”、”area”、”value”、”mx”、”ttl”、”status”] 中的某一个。
	Change *string `json:"Change,omitempty" name:"Change"`

	// 修改为，具体依赖 change 字段，必填参数。
	ChangeTo *string `json:"ChangeTo,omitempty" name:"ChangeTo"`

	// 要修改到的记录值，仅当 change 字段为 “record_type” 时为必填参数。
	Value *string `json:"Value,omitempty" name:"Value"`

	// MX记录优先级，仅当修改为 MX 记录时为必填参数。
	MX *string `json:"MX,omitempty" name:"MX"`
}

func (r *ModifyRecordBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RecordIdList")
	delete(f, "Change")
	delete(f, "ChangeTo")
	delete(f, "Value")
	delete(f, "MX")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批量任务ID
		JobId *uint64 `json:"JobId,omitempty" name:"JobId"`

		// 见modifyRecordBatchDetail
		DetailList []*ModifyRecordBatchDetail `json:"DetailList,omitempty" name:"DetailList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRecordBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordRemarkRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 记录 ID 。
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 解析记录备注，删除备注请提交空内容。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyRecordRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordId")
	delete(f, "DomainId")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRecordRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 记录类型，通过 API 记录类型获得，大写英文，比如：A 。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 记录线路，通过 API 记录线路获得，中文，比如：默认。
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// 记录值，如 IP : 200.200.200.200， CNAME : cname.dnspod.com.， MX : mail.dnspod.com.。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 记录 ID 。
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 线路的 ID，通过 API 记录线路获得，英文字符串，比如：10=1。参数RecordLineId优先级高于RecordLine，如果同时传递二者，优先使用RecordLineId参数。
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// MX 优先级，当记录类型是 MX 时有效，范围1-20，MX 记录时必选。
	MX *uint64 `json:"MX,omitempty" name:"MX"`

	// TTL，范围1-604800，不同等级域名最小值不同。
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 权重信息，0到100的整数。仅企业 VIP 域名可用，0 表示关闭，不传该参数，表示不设置权重信息。
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 记录初始状态，取值范围为 ENABLE 和 DISABLE 。默认为 ENABLE ，如果传入 DISABLE，解析不会生效，也不会验证负载均衡的限制。
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordType")
	delete(f, "RecordLine")
	delete(f, "Value")
	delete(f, "RecordId")
	delete(f, "DomainId")
	delete(f, "SubDomain")
	delete(f, "RecordLineId")
	delete(f, "MX")
	delete(f, "TTL")
	delete(f, "Weight")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录ID
		RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordStatusRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 记录 ID 。
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// 记录的状态。取值范围为 ENABLE 和 DISABLE。如果传入 DISABLE，解析不会生效，也不会验证负载均衡的限制。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *ModifyRecordStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordId")
	delete(f, "Status")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录ID。
		RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRecordStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubdomainStatusRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 记录类型。允许的值为A、CNAME、MX、TXT、NS、AAAA、SPF、SRV、CAA、URL、URL1。若要传多个，用英文逗号分隔，例如A,TXT,CNAME。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 记录状态。允许的值为disable。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
}

func (r *ModifySubdomainStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubdomainStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordType")
	delete(f, "Status")
	delete(f, "DomainId")
	delete(f, "SubDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubdomainStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySubdomainStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubdomainStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubdomainStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PurviewInfo struct {

	// 权限名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 权限值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type RecordCountInfo struct {

	// 子域名数量
	SubdomainCount *uint64 `json:"SubdomainCount,omitempty" name:"SubdomainCount"`

	// 列表返回的记录数
	ListCount *uint64 `json:"ListCount,omitempty" name:"ListCount"`

	// 总的记录数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type RecordInfo struct {

	// 记录 ID 。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 子域名(主机记录)。
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 记录类型, 详见 DescribeRecordType 接口。
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 解析记录的线路，详见 DescribeRecordLineList 接口。
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// 解析记录的线路 ID ，详见 DescribeRecordLineList 接口。
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// 记录值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 记录权重值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 记录的 MX 记录值，非 MX 记录类型，默认为 0。
	MX *uint64 `json:"MX,omitempty" name:"MX"`

	// 记录的 TTL 值。
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// 记录状态。0表示禁用，1表示启用。
	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`

	// 该记录的 D 监控状态。
	// "Ok" : 服务器正常。
	// "Warn" : 该记录有报警, 服务器返回 4XX。
	// "Down" : 服务器宕机。
	// "" : 该记录未开启 D 监控。
	MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// 记录的备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 记录最后更新时间。
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// 域名 ID 。
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

type RecordListItem struct {

	// 记录Id
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// 记录值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 记录状态，启用：ENABLE，暂停：DISABLE
	Status *string `json:"Status,omitempty" name:"Status"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// 主机名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 记录线路
	Line *string `json:"Line,omitempty" name:"Line"`

	// 线路Id
	LineId *string `json:"LineId,omitempty" name:"LineId"`

	// 记录类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 记录权重，用于负载均衡记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 记录监控状态，正常：OK，告警：WARN，宕机：DOWN，未设置监控或监控暂停则为空
	MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// 记录备注说明
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 记录缓存时间
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// MX值，只有MX记录有
	// 注意：此字段可能返回 null，表示取不到有效值。
	MX *uint64 `json:"MX,omitempty" name:"MX"`
}

type UserInfo struct {

	// 用户昵称
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 用户ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 用户账号, 邮箱格式
	Email *string `json:"Email,omitempty" name:"Email"`

	// 账号状态：”enabled”: 正常；”disabled”: 被封禁
	Status *string `json:"Status,omitempty" name:"Status"`

	// 电话号码
	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`

	// 邮箱是否通过验证：”yes”: 通过；”no”: 未通过
	EmailVerified *string `json:"EmailVerified,omitempty" name:"EmailVerified"`

	// 手机是否通过验证：”yes”: 通过；”no”: 未通过
	TelephoneVerified *string `json:"TelephoneVerified,omitempty" name:"TelephoneVerified"`

	// 账号等级, 按照用户账号下域名等级排序, 选取一个最高等级为账号等级, 具体对应情况参见域名等级。
	UserGrade *string `json:"UserGrade,omitempty" name:"UserGrade"`

	// 用户名称, 企业用户对应为公司名称
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 是否绑定微信：”yes”: 通过；”no”: 未通过
	WechatBinded *string `json:"WechatBinded,omitempty" name:"WechatBinded"`

	// 用户UIN
	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
}
