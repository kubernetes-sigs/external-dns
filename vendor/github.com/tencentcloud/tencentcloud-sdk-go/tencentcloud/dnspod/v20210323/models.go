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
<<<<<<< HEAD
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
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AddRecordBatch struct {
	// 记录类型, 详见 DescribeRecordType 接口。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 子域名(主机记录)，默认为@。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 解析记录的线路，详见 DescribeRecordLineList 接口，RecordLine和RecordLineId都未填时，默认为「默认」线路。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 解析记录的线路 ID，RecordLine和RecordLineId都有时，系统优先取 RecordLineId。
	RecordLineId *string `json:"RecordLineId,omitnil,omitempty" name:"RecordLineId"`

	// 记录权重值(暂未支持)。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 记录的 MX 记录值，非 MX 记录类型，默认为 0，MX记录则必选。
	MX *uint64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 记录的 TTL 值，默认600。
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 记录状态(暂未支持)。0表示禁用，1表示启用。默认启用。
	Enabled *uint64 `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 记录备注(暂未支持)。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type BatchRecordInfo struct {
	// 记录 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 子域名(主机记录)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录类型, 详见 DescribeRecordType 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 解析记录的线路，详见 DescribeRecordLineList 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 记录值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 记录的 TTL 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 记录添加状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 此条记录在列表中的ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 记录生效状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *uint64 `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 记录的MX权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	MX *uint64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 记录权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

// Predefined struct for user
type CheckRecordSnapshotRollbackRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 解析记录信息
	Record *SnapshotRecord `json:"Record,omitnil,omitempty" name:"Record"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type CheckRecordSnapshotRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 解析记录信息
	Record *SnapshotRecord `json:"Record,omitnil,omitempty" name:"Record"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *CheckRecordSnapshotRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRecordSnapshotRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SnapshotId")
	delete(f, "Record")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckRecordSnapshotRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckRecordSnapshotRollbackResponseParams struct {
	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckRecordSnapshotRollbackResponse struct {
	*tchttp.BaseResponse
	Response *CheckRecordSnapshotRollbackResponseParams `json:"Response"`
}

func (r *CheckRecordSnapshotRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckRecordSnapshotRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckSnapshotRollbackRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照记录 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type CheckSnapshotRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照记录 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *CheckSnapshotRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckSnapshotRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SnapshotId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckSnapshotRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckSnapshotRollbackResponseParams struct {
	// 快照记录 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 回滚时长（分钟）
	CostMinutes *uint64 `json:"CostMinutes,omitnil,omitempty" name:"CostMinutes"`

	// 快照所属域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 解析记录总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 值为 1，表示超时
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 检查失败数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Failed *uint64 `json:"Failed,omitnil,omitempty" name:"Failed"`

	// 失败记录信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedRecordList []*SnapshotRecord `json:"FailedRecordList,omitnil,omitempty" name:"FailedRecordList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckSnapshotRollbackResponse struct {
	*tchttp.BaseResponse
	Response *CheckSnapshotRollbackResponseParams `json:"Response"`
}

func (r *CheckSnapshotRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckSnapshotRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDealRequestParams struct {
	// 询价类型，1 新购，2 续费，3 套餐升级（增值服务暂时只支持新购）
	DealType *uint64 `json:"DealType,omitnil,omitempty" name:"DealType"`

	// 商品类型，1 域名套餐 2 增值服务
	GoodsType *uint64 `json:"GoodsType,omitnil,omitempty" name:"GoodsType"`

	// 套餐类型：
	// DP_PLUS：专业版
	// DP_EXPERT：企业版
	// DP_ULTRA：尊享版
	// 
	// 增值服务类型
	// LB：负载均衡
	// URL：URL转发
	// DMONITOR_TASKS：D监控任务数
	// DMONITOR_IP：D监控备用 IP 数
	// CUSTOMLINE：自定义线路数
	GoodsChildType *string `json:"GoodsChildType,omitnil,omitempty" name:"GoodsChildType"`

	// 增值服务购买数量，如果是域名套餐固定为1，如果是增值服务则按以下规则：
	// 负载均衡、D监控任务数、D监控备用 IP 数、自定义线路数、URL 转发（必须是5的正整数倍，如 5、10、15 等）
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 是否开启自动续费，1 开启，2 不开启（增值服务暂不支持自动续费），默认值为 2 不开启
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 需要绑定套餐的域名，如 dnspod.cn，如果是续费或升级，domain 参数必须要传，新购可不传。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 套餐时长：
	// 1. 套餐以月为单位（按月只能是 3、6 还有 12 的倍数），套餐例如购买一年则传12，最大120 。（续费最低一年）
	// 2. 升级套餐时不需要传。
	// 3. 增值服务的时长单位为年，买一年传1（增值服务新购按年只能是 1，增值服务续费最大为 10）
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 套餐类型，需要升级到的套餐类型，只有升级时需要。
	NewPackageType *string `json:"NewPackageType,omitnil,omitempty" name:"NewPackageType"`
}

type CreateDealRequest struct {
	*tchttp.BaseRequest
	
	// 询价类型，1 新购，2 续费，3 套餐升级（增值服务暂时只支持新购）
	DealType *uint64 `json:"DealType,omitnil,omitempty" name:"DealType"`

	// 商品类型，1 域名套餐 2 增值服务
	GoodsType *uint64 `json:"GoodsType,omitnil,omitempty" name:"GoodsType"`

	// 套餐类型：
	// DP_PLUS：专业版
	// DP_EXPERT：企业版
	// DP_ULTRA：尊享版
	// 
	// 增值服务类型
	// LB：负载均衡
	// URL：URL转发
	// DMONITOR_TASKS：D监控任务数
	// DMONITOR_IP：D监控备用 IP 数
	// CUSTOMLINE：自定义线路数
	GoodsChildType *string `json:"GoodsChildType,omitnil,omitempty" name:"GoodsChildType"`

	// 增值服务购买数量，如果是域名套餐固定为1，如果是增值服务则按以下规则：
	// 负载均衡、D监控任务数、D监控备用 IP 数、自定义线路数、URL 转发（必须是5的正整数倍，如 5、10、15 等）
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 是否开启自动续费，1 开启，2 不开启（增值服务暂不支持自动续费），默认值为 2 不开启
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 需要绑定套餐的域名，如 dnspod.cn，如果是续费或升级，domain 参数必须要传，新购可不传。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 套餐时长：
	// 1. 套餐以月为单位（按月只能是 3、6 还有 12 的倍数），套餐例如购买一年则传12，最大120 。（续费最低一年）
	// 2. 升级套餐时不需要传。
	// 3. 增值服务的时长单位为年，买一年传1（增值服务新购按年只能是 1，增值服务续费最大为 10）
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 套餐类型，需要升级到的套餐类型，只有升级时需要。
	NewPackageType *string `json:"NewPackageType,omitnil,omitempty" name:"NewPackageType"`
}

func (r *CreateDealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealType")
	delete(f, "GoodsType")
	delete(f, "GoodsChildType")
	delete(f, "GoodsNum")
	delete(f, "AutoRenew")
	delete(f, "Domain")
	delete(f, "TimeSpan")
	delete(f, "NewPackageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDealResponseParams struct {
	// 大订单号，一个大订单号下可以有多个子订单，说明是同一次下单
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 子订单列表
	DealList []*Deals `json:"DealList,omitnil,omitempty" name:"DealList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDealResponse struct {
	*tchttp.BaseResponse
	Response *CreateDealResponseParams `json:"Response"`
}

func (r *CreateDealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainAliasRequestParams struct {
	// 域名别名
	DomainAlias *string `json:"DomainAlias,omitnil,omitempty" name:"DomainAlias"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID，参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type CreateDomainAliasRequest struct {
	*tchttp.BaseRequest
	
	// 域名别名
	DomainAlias *string `json:"DomainAlias,omitnil,omitempty" name:"DomainAlias"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID，参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type CreateDomainAliasResponseParams struct {
	// 域名别名ID
	DomainAliasId *int64 `json:"DomainAliasId,omitnil,omitempty" name:"DomainAliasId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainAliasResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainAliasResponseParams `json:"Response"`
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
	RecordList []*CreateDomainBatchRecord `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 任务编号
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainGrade *string `json:"DomainGrade,omitnil,omitempty" name:"DomainGrade"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 该条任务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

type CreateDomainBatchRecord struct {
	// 子域名(主机记录)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录类型, 详见 DescribeRecordType 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 解析记录的线路，详见 DescribeRecordLineList 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 记录值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 记录的 TTL 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 记录添加状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 此条记录在列表中的ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

// Predefined struct for user
type CreateDomainBatchRequestParams struct {
	// 域名数组
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 每个域名添加 @ 和 www 的 A 记录值，记录值为IP，如果不传此参数或者传空，将只添加域名，不添加记录。
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`
}

type CreateDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// 域名数组
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 每个域名添加 @ 和 www 的 A 记录值，记录值为IP，如果不传此参数或者传空，将只添加域名，不添加记录。
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`
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

// Predefined struct for user
type CreateDomainBatchResponseParams struct {
	// 批量添加域名信息
	DetailList []*CreateDomainBatchDetail `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// 批量任务的ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainBatchResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateDomainCustomLineRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 自定义线路名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义线路IP段，用-分割
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 域名ID，如果传了DomainId，系统将会忽略Domain参数，优先使用DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type CreateDomainCustomLineRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 自定义线路名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义线路IP段，用-分割
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 域名ID，如果传了DomainId，系统将会忽略Domain参数，优先使用DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *CreateDomainCustomLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainCustomLineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Name")
	delete(f, "Area")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainCustomLineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainCustomLineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainCustomLineResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainCustomLineResponseParams `json:"Response"`
}

func (r *CreateDomainCustomLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainCustomLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainGroupRequestParams struct {
	// 域名分组
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type CreateDomainGroupRequest struct {
	*tchttp.BaseRequest
	
	// 域名分组
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
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

// Predefined struct for user
type CreateDomainGroupResponseParams struct {
	// 域名分组ID
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名分组ID。可以通过接口DescribeDomainGroupList查看当前域名分组信息
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 是否星标域名，”yes”、”no” 分别代表是和否。
	IsMark *string `json:"IsMark,omitnil,omitempty" name:"IsMark"`

	// 添加子域名时，是否迁移相关父域名的解析记录。不传默认为 true
	TransferSubDomain *bool `json:"TransferSubDomain,omitnil,omitempty" name:"TransferSubDomain"`

	// 域名绑定的标签
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名分组ID。可以通过接口DescribeDomainGroupList查看当前域名分组信息
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 是否星标域名，”yes”、”no” 分别代表是和否。
	IsMark *string `json:"IsMark,omitnil,omitempty" name:"IsMark"`

	// 添加子域名时，是否迁移相关父域名的解析记录。不传默认为 true
	TransferSubDomain *bool `json:"TransferSubDomain,omitnil,omitempty" name:"TransferSubDomain"`

	// 域名绑定的标签
	Tags []*TagItem `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "TransferSubDomain")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainResponseParams struct {
	// 域名信息
	DomainInfo *DomainCreateInfo `json:"DomainInfo,omitnil,omitempty" name:"DomainInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainResponseParams `json:"Response"`
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
	RecordList []*CreateRecordBatchRecord `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 任务编号
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainGrade *string `json:"DomainGrade,omitnil,omitempty" name:"DomainGrade"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 该条任务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 域名ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type CreateRecordBatchRecord struct {
	// 子域名(主机记录)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录类型, 详见 DescribeRecordType 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 解析记录的线路，详见 DescribeRecordLineList 接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 记录值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 记录的 TTL 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 记录添加状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 此条记录在列表中的ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 记录的MX权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	MX *uint64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 记录的权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

// Predefined struct for user
type CreateRecordBatchRequestParams struct {
	// 域名ID，多个 domain_id 用英文逗号进行分割。
	DomainIdList []*string `json:"DomainIdList,omitnil,omitempty" name:"DomainIdList"`

	// 记录数组
	RecordList []*AddRecordBatch `json:"RecordList,omitnil,omitempty" name:"RecordList"`
}

type CreateRecordBatchRequest struct {
	*tchttp.BaseRequest
	
	// 域名ID，多个 domain_id 用英文逗号进行分割。
	DomainIdList []*string `json:"DomainIdList,omitnil,omitempty" name:"DomainIdList"`

	// 记录数组
	RecordList []*AddRecordBatch `json:"RecordList,omitnil,omitempty" name:"RecordList"`
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

// Predefined struct for user
type CreateRecordBatchResponseParams struct {
	// 批量添加域名信息
	DetailList []*CreateRecordBatchDetail `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// 批量任务的ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRecordBatchResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecordBatchResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRecordGroupRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type CreateRecordGroupRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *CreateRecordGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "GroupName")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecordGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecordGroupResponseParams struct {
	// 新增的分组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRecordGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecordGroupResponseParams `json:"Response"`
}

func (r *CreateRecordGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecordRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录类型，通过 API 记录类型获得，大写英文，比如：A 。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录线路，通过 API 记录线路获得，中文，比如：默认。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 记录值，如 IP : 200.200.200.200， CNAME : cname.dnspod.com.， MX : mail.dnspod.com.。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 线路的 ID，通过 API 记录线路获得，英文字符串，比如：10=1。参数RecordLineId优先级高于RecordLine，如果同时传递二者，优先使用RecordLineId参数。
	RecordLineId *string `json:"RecordLineId,omitnil,omitempty" name:"RecordLineId"`

	// MX 优先级，当记录类型是 MX 时有效，范围1-20，MX 记录时必选。
	MX *uint64 `json:"MX,omitnil,omitempty" name:"MX"`

	// TTL，范围1-604800，不同套餐域名最小值不同。
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 权重信息，0到100的整数。0 表示关闭，不传该参数，表示不设置权重信息。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 记录初始状态，取值范围为 ENABLE 和 DISABLE 。默认为 ENABLE ，如果传入 DISABLE，解析不会生效，也不会验证负载均衡的限制。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 开启DNSSEC时，强制添加CNAME/URL记录
	DnssecConflictMode *string `json:"DnssecConflictMode,omitnil,omitempty" name:"DnssecConflictMode"`
}

type CreateRecordRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录类型，通过 API 记录类型获得，大写英文，比如：A 。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录线路，通过 API 记录线路获得，中文，比如：默认。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 记录值，如 IP : 200.200.200.200， CNAME : cname.dnspod.com.， MX : mail.dnspod.com.。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 线路的 ID，通过 API 记录线路获得，英文字符串，比如：10=1。参数RecordLineId优先级高于RecordLine，如果同时传递二者，优先使用RecordLineId参数。
	RecordLineId *string `json:"RecordLineId,omitnil,omitempty" name:"RecordLineId"`

	// MX 优先级，当记录类型是 MX 时有效，范围1-20，MX 记录时必选。
	MX *uint64 `json:"MX,omitnil,omitempty" name:"MX"`

	// TTL，范围1-604800，不同套餐域名最小值不同。
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 权重信息，0到100的整数。0 表示关闭，不传该参数，表示不设置权重信息。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 记录初始状态，取值范围为 ENABLE 和 DISABLE 。默认为 ENABLE ，如果传入 DISABLE，解析不会生效，也不会验证负载均衡的限制。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 开启DNSSEC时，强制添加CNAME/URL记录
	DnssecConflictMode *string `json:"DnssecConflictMode,omitnil,omitempty" name:"DnssecConflictMode"`
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
	delete(f, "Remark")
	delete(f, "DnssecConflictMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecordResponseParams struct {
	// 记录ID
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateSnapshotRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type CreateSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *CreateSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotResponseParams `json:"Response"`
}

func (r *CreateSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomLineInfo struct {
	// 域名ID
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 自定义线路名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义线路IP段
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 已使用IP段个数
	UseCount *uint64 `json:"UseCount,omitnil,omitempty" name:"UseCount"`

	// 允许使用IP段最大个数
	MaxCount *uint64 `json:"MaxCount,omitnil,omitempty" name:"MaxCount"`
}

type Deals struct {
	// 子订单ID
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 子订单号
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

// Predefined struct for user
type DeleteDomainAliasRequestParams struct {
	// 域名别名ID。可以通过接口DescribeDomainAliasList查到所有的域名别名列表以及对应的ID
	DomainAliasId *int64 `json:"DomainAliasId,omitnil,omitempty" name:"DomainAliasId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID，参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DeleteDomainAliasRequest struct {
	*tchttp.BaseRequest
	
	// 域名别名ID。可以通过接口DescribeDomainAliasList查到所有的域名别名列表以及对应的ID
	DomainAliasId *int64 `json:"DomainAliasId,omitnil,omitempty" name:"DomainAliasId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID，参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type DeleteDomainAliasResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDomainAliasResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainAliasResponseParams `json:"Response"`
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

type DeleteDomainBatchDetail struct {
	// 域名 ID
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`

	// 删除状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

// Predefined struct for user
type DeleteDomainBatchRequestParams struct {
	// 域名数组
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

type DeleteDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// 域名数组
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

func (r *DeleteDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainBatchResponseParams struct {
	// 任务 ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务详情数组
	DetailList []*DeleteDomainBatchDetail `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainBatchResponseParams `json:"Response"`
}

func (r *DeleteDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainCustomLineRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 自定义线路名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 域名ID，如果传了DomainId，系统将会忽略Domain参数，优先使用DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DeleteDomainCustomLineRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 自定义线路名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 域名ID，如果传了DomainId，系统将会忽略Domain参数，优先使用DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DeleteDomainCustomLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainCustomLineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Name")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainCustomLineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainCustomLineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDomainCustomLineResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainCustomLineResponseParams `json:"Response"`
}

func (r *DeleteDomainCustomLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainCustomLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DeleteDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type DeleteDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainResponseParams `json:"Response"`
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

type DeleteRecordBatchDetail struct {
	// 域名 ID
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`

	// 删除状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 解析记录列表，json 序列化之后的字符串形式
	RecordList *string `json:"RecordList,omitnil,omitempty" name:"RecordList"`
}

// Predefined struct for user
type DeleteRecordBatchRequestParams struct {
	// 解析记录 ID 数组
	RecordIdList []*uint64 `json:"RecordIdList,omitnil,omitempty" name:"RecordIdList"`
}

type DeleteRecordBatchRequest struct {
	*tchttp.BaseRequest
	
	// 解析记录 ID 数组
	RecordIdList []*uint64 `json:"RecordIdList,omitnil,omitempty" name:"RecordIdList"`
}

func (r *DeleteRecordBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RecordIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordBatchResponseParams struct {
	// 批量任务 ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务详情
	DetailList []*DeleteRecordBatchDetail `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRecordBatchResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordBatchResponseParams `json:"Response"`
}

func (r *DeleteRecordBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordGroupRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DeleteRecordGroupRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DeleteRecordGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "GroupId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRecordGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordGroupResponseParams `json:"Response"`
}

func (r *DeleteRecordGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录 ID 。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DeleteRecordRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录 ID 。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type DeleteRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteShareDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名共享的账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DeleteShareDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名共享的账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type DeleteShareDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteShareDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShareDomainResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteSnapshotRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照记录 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DeleteSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照记录 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DeleteSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SnapshotId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotResponseParams `json:"Response"`
}

func (r *DeleteSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchTaskDetail struct {
	// 见BatchRecordInfo
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordList []*BatchRecordInfo `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 任务编号
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainGrade *string `json:"DomainGrade,omitnil,omitempty" name:"DomainGrade"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 该条任务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 域名ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

// Predefined struct for user
type DescribeBatchTaskRequestParams struct {
	// 任务ID。操作批量接口时会返回JobId
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeBatchTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。操作批量接口时会返回JobId
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`
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

// Predefined struct for user
type DescribeBatchTaskResponseParams struct {
	// 批量任务详情
	DetailList []*DescribeBatchTaskDetail `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// 总任务条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 成功条数
	SuccessCount *uint64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 失败条数
	FailCount *uint64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`

	// 批量任务类型
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 任务创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchTaskResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDomainAliasListRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID,域名ID，参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeDomainAliasListRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID,域名ID，参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *int64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type DescribeDomainAliasListResponseParams struct {
	// 域名别名列表
	DomainAliasList []*DomainAliasInfo `json:"DomainAliasList,omitnil,omitempty" name:"DomainAliasList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainAliasListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainAliasListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDomainAnalyticsRequestParams struct {
	// 要查询解析量的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询的开始时间，格式：YYYY-MM-DD
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询的结束时间，格式：YYYY-MM-DD
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// DATE:按天维度统计 HOUR:按小时维度统计
	DnsFormat *string `json:"DnsFormat,omitnil,omitempty" name:"DnsFormat"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeDomainAnalyticsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询解析量的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询的开始时间，格式：YYYY-MM-DD
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询的结束时间，格式：YYYY-MM-DD
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// DATE:按天维度统计 HOUR:按小时维度统计
	DnsFormat *string `json:"DnsFormat,omitnil,omitempty" name:"DnsFormat"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeDomainAnalyticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainAnalyticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "DnsFormat")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainAnalyticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainAnalyticsResponseParams struct {
	// 当前统计维度解析量小计
	Data []*DomainAnalyticsDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 域名解析量统计查询信息
	Info *DomainAnalyticsInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// 域名别名解析量统计信息
	AliasData []*DomainAliasAnalyticsItem `json:"AliasData,omitnil,omitempty" name:"AliasData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainAnalyticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainAnalyticsResponseParams `json:"Response"`
}

func (r *DescribeDomainAnalyticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainAnalyticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainCustomLineListRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID，如果传了DomainId，系统将会忽略Domain参数，优先使用DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeDomainCustomLineListRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID，如果传了DomainId，系统将会忽略Domain参数，优先使用DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeDomainCustomLineListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainCustomLineListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainCustomLineListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainCustomLineListResponseParams struct {
	// 自定义线路列表
	LineList []*CustomLineInfo `json:"LineList,omitnil,omitempty" name:"LineList"`

	// 可添加的自定义线路条数
	AvailableCount *uint64 `json:"AvailableCount,omitnil,omitempty" name:"AvailableCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainCustomLineListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainCustomLineListResponseParams `json:"Response"`
}

func (r *DescribeDomainCustomLineListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainCustomLineListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainFilterListRequestParams struct {
	// 根据域名分组类型获取域名。可取值为 ALL，MINE，SHARE，RECENT。
	// ALL：全部
	// MINE：我的域名
	// SHARE：共享给我的域名
	// RECENT：最近操作过的域名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 记录开始的偏移, 第一条记录为 0, 依次类推。默认值为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 要获取的域名数量, 比如获取 20 个, 则为 20。默认值为 5000。如果账户中的域名数量超过了 5000, 将会强制分页并且只返回前 5000 条, 这时需要通过 Offset 和 Limit 参数去获取其它域名。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据域名分组 id 获取域名，可通过 DescribeDomain 或 DescribeDomainList 接口 GroupId 字段获取。
	GroupId []*int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 根据关键字获取域名。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序字段。可取值为 NAME，STATUS，RECORDS，GRADE，UPDATED_ON。
	// NAME：域名名称
	// STATUS：域名状态
	// RECORDS：记录数量
	// GRADE：套餐等级
	// UPDATED_ON：更新时间
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序类型，升序：ASC，降序：DESC。
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 根据域名状态获取域名。可取值为 ENABLE，LOCK，PAUSE，SPAM。
	// ENABLE：正常
	// LOCK：锁定
	// PAUSE：暂停
	// SPAM：封禁
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 根据套餐获取域名，可通过 DescribeDomain 或 DescribeDomainList 接口 Grade 字段获取。
	Package []*string `json:"Package,omitnil,omitempty" name:"Package"`

	// 根据备注信息获取域名。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 要获取域名的更新时间起始时间点，如 '2021-05-01 03:00:00'。
	UpdatedAtBegin *string `json:"UpdatedAtBegin,omitnil,omitempty" name:"UpdatedAtBegin"`

	// 要获取域名的更新时间终止时间点，如 '2021-05-10 20:00:00'。
	UpdatedAtEnd *string `json:"UpdatedAtEnd,omitnil,omitempty" name:"UpdatedAtEnd"`

	// 要获取域名的记录数查询区间起点。
	RecordCountBegin *uint64 `json:"RecordCountBegin,omitnil,omitempty" name:"RecordCountBegin"`

	// 要获取域名的记录数查询区间终点。
	RecordCountEnd *uint64 `json:"RecordCountEnd,omitnil,omitempty" name:"RecordCountEnd"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 标签过滤
	Tags []*TagItemFilter `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeDomainFilterListRequest struct {
	*tchttp.BaseRequest
	
	// 根据域名分组类型获取域名。可取值为 ALL，MINE，SHARE，RECENT。
	// ALL：全部
	// MINE：我的域名
	// SHARE：共享给我的域名
	// RECENT：最近操作过的域名
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 记录开始的偏移, 第一条记录为 0, 依次类推。默认值为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 要获取的域名数量, 比如获取 20 个, 则为 20。默认值为 5000。如果账户中的域名数量超过了 5000, 将会强制分页并且只返回前 5000 条, 这时需要通过 Offset 和 Limit 参数去获取其它域名。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据域名分组 id 获取域名，可通过 DescribeDomain 或 DescribeDomainList 接口 GroupId 字段获取。
	GroupId []*int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 根据关键字获取域名。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序字段。可取值为 NAME，STATUS，RECORDS，GRADE，UPDATED_ON。
	// NAME：域名名称
	// STATUS：域名状态
	// RECORDS：记录数量
	// GRADE：套餐等级
	// UPDATED_ON：更新时间
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序类型，升序：ASC，降序：DESC。
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 根据域名状态获取域名。可取值为 ENABLE，LOCK，PAUSE，SPAM。
	// ENABLE：正常
	// LOCK：锁定
	// PAUSE：暂停
	// SPAM：封禁
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 根据套餐获取域名，可通过 DescribeDomain 或 DescribeDomainList 接口 Grade 字段获取。
	Package []*string `json:"Package,omitnil,omitempty" name:"Package"`

	// 根据备注信息获取域名。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 要获取域名的更新时间起始时间点，如 '2021-05-01 03:00:00'。
	UpdatedAtBegin *string `json:"UpdatedAtBegin,omitnil,omitempty" name:"UpdatedAtBegin"`

	// 要获取域名的更新时间终止时间点，如 '2021-05-10 20:00:00'。
	UpdatedAtEnd *string `json:"UpdatedAtEnd,omitnil,omitempty" name:"UpdatedAtEnd"`

	// 要获取域名的记录数查询区间起点。
	RecordCountBegin *uint64 `json:"RecordCountBegin,omitnil,omitempty" name:"RecordCountBegin"`

	// 要获取域名的记录数查询区间终点。
	RecordCountEnd *uint64 `json:"RecordCountEnd,omitnil,omitempty" name:"RecordCountEnd"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 标签过滤
	Tags []*TagItemFilter `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeDomainFilterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainFilterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "GroupId")
	delete(f, "Keyword")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "Status")
	delete(f, "Package")
	delete(f, "Remark")
	delete(f, "UpdatedAtBegin")
	delete(f, "UpdatedAtEnd")
	delete(f, "RecordCountBegin")
	delete(f, "RecordCountEnd")
	delete(f, "ProjectId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainFilterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainFilterListResponseParams struct {
	// 列表页统计信息
	DomainCountInfo *DomainCountInfo `json:"DomainCountInfo,omitnil,omitempty" name:"DomainCountInfo"`

	// 域名列表
	DomainList []*DomainListItem `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainFilterListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainFilterListResponseParams `json:"Response"`
}

func (r *DescribeDomainFilterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainFilterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainGroupListRequestParams struct {

}

type DescribeDomainGroupListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDomainGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainGroupListResponseParams struct {
	// 分组列表
	GroupList []*GroupInfo `json:"GroupList,omitnil,omitempty" name:"GroupList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainGroupListResponseParams `json:"Response"`
}

func (r *DescribeDomainGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainListRequestParams struct {
	// 域名分组类型，默认为ALL。可取值为ALL，MINE，SHARE，ISMARK，PAUSE，VIP，RECENT，SHARE_OUT，FREE。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 记录开始的偏移, 第一条记录为 0, 依次类推。默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 要获取的域名数量, 比如获取20个, 则为20。默认值为3000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分组ID, 获取指定分组的域名
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 根据关键字搜索域名
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 标签过滤
	Tags []*TagItemFilter `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeDomainListRequest struct {
	*tchttp.BaseRequest
	
	// 域名分组类型，默认为ALL。可取值为ALL，MINE，SHARE，ISMARK，PAUSE，VIP，RECENT，SHARE_OUT，FREE。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 记录开始的偏移, 第一条记录为 0, 依次类推。默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 要获取的域名数量, 比如获取20个, 则为20。默认值为3000。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分组ID, 获取指定分组的域名
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 根据关键字搜索域名
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 标签过滤
	Tags []*TagItemFilter `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainListResponseParams struct {
	// 列表页统计信息
	DomainCountInfo *DomainCountInfo `json:"DomainCountInfo,omitnil,omitempty" name:"DomainCountInfo"`

	// 域名列表
	DomainList []*DomainListItem `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDomainLogListRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 记录开始的偏移，第一条记录为 0，依次类推，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 共要获取的日志条数，比如获取20条，则为20，默认为500条，单次最多获取500条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDomainLogListRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 记录开始的偏移，第一条记录为 0，依次类推，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 共要获取的日志条数，比如获取20条，则为20，默认为500条，单次最多获取500条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeDomainLogListResponseParams struct {
	// 域名信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogList []*string `json:"LogList,omitnil,omitempty" name:"LogList"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 日志总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainLogListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDomainPreviewRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeDomainPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeDomainPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainPreviewResponseParams struct {
	// 域名概览信息
	Domain *PreviewDetail `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainPreviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainPreviewResponseParams `json:"Response"`
}

func (r *DescribeDomainPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainPurviewRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeDomainPurviewRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type DescribeDomainPurviewResponseParams struct {
	// 域名权限列表
	PurviewList []*PurviewInfo `json:"PurviewList,omitnil,omitempty" name:"PurviewList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainPurviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainPurviewResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type DescribeDomainResponseParams struct {
	// 域名信息
	DomainInfo *DomainInfo `json:"DomainInfo,omitnil,omitempty" name:"DomainInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDomainShareInfoRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeDomainShareInfoRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type DescribeDomainShareInfoResponseParams struct {
	// 域名共享信息
	ShareList []*DomainShareInfo `json:"ShareList,omitnil,omitempty" name:"ShareList"`

	// 域名拥有者账号
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainShareInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainShareInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDomainWhoisRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeDomainWhoisRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *DescribeDomainWhoisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainWhoisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainWhoisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainWhoisResponseParams struct {
	// 域名Whois信息
	Info *WhoisInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainWhoisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainWhoisResponseParams `json:"Response"`
}

func (r *DescribeDomainWhoisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainWhoisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackageDetailRequestParams struct {

}

type DescribePackageDetailRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePackageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePackageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackageDetailResponseParams struct {
	// 套餐配置详情
	Info []*PackageDetailItem `json:"Info,omitnil,omitempty" name:"Info"`

	// 套餐代码列表
	LevelMap []*string `json:"LevelMap,omitnil,omitempty" name:"LevelMap"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePackageDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePackageDetailResponseParams `json:"Response"`
}

func (r *DescribePackageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordExistExceptDefaultNSRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeRecordExistExceptDefaultNSRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeRecordExistExceptDefaultNSRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordExistExceptDefaultNSRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordExistExceptDefaultNSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordExistExceptDefaultNSResponseParams struct {
	// true 是 false 否
	Exist *bool `json:"Exist,omitnil,omitempty" name:"Exist"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordExistExceptDefaultNSResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordExistExceptDefaultNSResponseParams `json:"Response"`
}

func (r *DescribeRecordExistExceptDefaultNSResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordExistExceptDefaultNSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordFilterListRequestParams struct {
	// 要获取的解析记录所属的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 要获取的解析记录所属的域名 Id，如果传了 DomainId，系统将会忽略 Domain 参数。 可以通过接口 DescribeDomainList 查到所有的 Domain 以及 DomainId。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 根据解析记录的主机头获取解析记录。默认模糊匹配。可以通过设置 IsExactSubdomain 参数为 true 进行精确查找。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 获取某些类型的解析记录，如 A，CNAME，NS，AAAA，显性URL，隐性URL，CAA，SPF等。
	RecordType []*string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 获取某些线路ID的解析记录。可以通过接口 DescribeRecordLineList 查看当前域名允许的线路信息。
	RecordLine []*string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 获取某些分组下的解析记录时，传这个分组 Id。可以通过接口 DescribeRecordGroupList 接口 GroupId 字段获取。
	GroupId []*uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 通过关键字搜索解析记录，当前支持搜索主机头和记录值
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序字段，支持 NAME，LINE，TYPE，VALUE，WEIGHT，MX，TTL，UPDATED_ON 几个字段。
	// NAME：解析记录的主机头
	// LINE：解析记录线路
	// TYPE：解析记录类型
	// VALUE：解析记录值
	// WEIGHT：权重
	// MX：MX 优先级
	// TTL：解析记录缓存时间
	// UPDATED_ON：解析记录更新时间
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序方式，升序：ASC，降序：DESC。默认值为ASC。
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数量，当前Limit最大支持3000。默认值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据解析记录的值获取解析记录
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`

	// 根据解析记录的状态获取解析记录。可取值为 ENABLE，DISABLE。
	// ENABLE：正常 
	// DISABLE：暂停 
	RecordStatus []*string `json:"RecordStatus,omitnil,omitempty" name:"RecordStatus"`

	// 要获取解析记录权重查询区间起点。
	WeightBegin *uint64 `json:"WeightBegin,omitnil,omitempty" name:"WeightBegin"`

	// 要获取解析记录权重查询区间终点。
	WeightEnd *uint64 `json:"WeightEnd,omitnil,omitempty" name:"WeightEnd"`

	// 要获取解析记录 MX 优先级查询区间起点。
	MXBegin *uint64 `json:"MXBegin,omitnil,omitempty" name:"MXBegin"`

	// 要获取解析记录 MX 优先级查询区间终点。
	MXEnd *uint64 `json:"MXEnd,omitnil,omitempty" name:"MXEnd"`

	// 要获取解析记录 TTL 查询区间起点。
	TTLBegin *uint64 `json:"TTLBegin,omitnil,omitempty" name:"TTLBegin"`

	// 要获取解析记录 TTL 查询区间终点。
	TTLEnd *uint64 `json:"TTLEnd,omitnil,omitempty" name:"TTLEnd"`

	// 要获取解析记录更新时间查询区间起点。
	UpdatedAtBegin *string `json:"UpdatedAtBegin,omitnil,omitempty" name:"UpdatedAtBegin"`

	// 要获取解析记录更新时间查询区间终点。
	UpdatedAtEnd *string `json:"UpdatedAtEnd,omitnil,omitempty" name:"UpdatedAtEnd"`

	// 根据解析记录的备注获取解析记录。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否根据 Subdomain 参数进行精确查找。
	IsExactSubDomain *bool `json:"IsExactSubDomain,omitnil,omitempty" name:"IsExactSubDomain"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeRecordFilterListRequest struct {
	*tchttp.BaseRequest
	
	// 要获取的解析记录所属的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 要获取的解析记录所属的域名 Id，如果传了 DomainId，系统将会忽略 Domain 参数。 可以通过接口 DescribeDomainList 查到所有的 Domain 以及 DomainId。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 根据解析记录的主机头获取解析记录。默认模糊匹配。可以通过设置 IsExactSubdomain 参数为 true 进行精确查找。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 获取某些类型的解析记录，如 A，CNAME，NS，AAAA，显性URL，隐性URL，CAA，SPF等。
	RecordType []*string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 获取某些线路ID的解析记录。可以通过接口 DescribeRecordLineList 查看当前域名允许的线路信息。
	RecordLine []*string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 获取某些分组下的解析记录时，传这个分组 Id。可以通过接口 DescribeRecordGroupList 接口 GroupId 字段获取。
	GroupId []*uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 通过关键字搜索解析记录，当前支持搜索主机头和记录值
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序字段，支持 NAME，LINE，TYPE，VALUE，WEIGHT，MX，TTL，UPDATED_ON 几个字段。
	// NAME：解析记录的主机头
	// LINE：解析记录线路
	// TYPE：解析记录类型
	// VALUE：解析记录值
	// WEIGHT：权重
	// MX：MX 优先级
	// TTL：解析记录缓存时间
	// UPDATED_ON：解析记录更新时间
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序方式，升序：ASC，降序：DESC。默认值为ASC。
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数量，当前Limit最大支持3000。默认值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据解析记录的值获取解析记录
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`

	// 根据解析记录的状态获取解析记录。可取值为 ENABLE，DISABLE。
	// ENABLE：正常 
	// DISABLE：暂停 
	RecordStatus []*string `json:"RecordStatus,omitnil,omitempty" name:"RecordStatus"`

	// 要获取解析记录权重查询区间起点。
	WeightBegin *uint64 `json:"WeightBegin,omitnil,omitempty" name:"WeightBegin"`

	// 要获取解析记录权重查询区间终点。
	WeightEnd *uint64 `json:"WeightEnd,omitnil,omitempty" name:"WeightEnd"`

	// 要获取解析记录 MX 优先级查询区间起点。
	MXBegin *uint64 `json:"MXBegin,omitnil,omitempty" name:"MXBegin"`

	// 要获取解析记录 MX 优先级查询区间终点。
	MXEnd *uint64 `json:"MXEnd,omitnil,omitempty" name:"MXEnd"`

	// 要获取解析记录 TTL 查询区间起点。
	TTLBegin *uint64 `json:"TTLBegin,omitnil,omitempty" name:"TTLBegin"`

	// 要获取解析记录 TTL 查询区间终点。
	TTLEnd *uint64 `json:"TTLEnd,omitnil,omitempty" name:"TTLEnd"`

	// 要获取解析记录更新时间查询区间起点。
	UpdatedAtBegin *string `json:"UpdatedAtBegin,omitnil,omitempty" name:"UpdatedAtBegin"`

	// 要获取解析记录更新时间查询区间终点。
	UpdatedAtEnd *string `json:"UpdatedAtEnd,omitnil,omitempty" name:"UpdatedAtEnd"`

	// 根据解析记录的备注获取解析记录。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否根据 Subdomain 参数进行精确查找。
	IsExactSubDomain *bool `json:"IsExactSubDomain,omitnil,omitempty" name:"IsExactSubDomain"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeRecordFilterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordFilterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "SubDomain")
	delete(f, "RecordType")
	delete(f, "RecordLine")
	delete(f, "GroupId")
	delete(f, "Keyword")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RecordValue")
	delete(f, "RecordStatus")
	delete(f, "WeightBegin")
	delete(f, "WeightEnd")
	delete(f, "MXBegin")
	delete(f, "MXEnd")
	delete(f, "TTLBegin")
	delete(f, "TTLEnd")
	delete(f, "UpdatedAtBegin")
	delete(f, "UpdatedAtEnd")
	delete(f, "Remark")
	delete(f, "IsExactSubDomain")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordFilterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordFilterListResponseParams struct {
	// 记录的数量统计信息
	RecordCountInfo *RecordCountInfo `json:"RecordCountInfo,omitnil,omitempty" name:"RecordCountInfo"`

	// 获取的记录列表
	RecordList []*RecordListItem `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordFilterListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordFilterListResponseParams `json:"Response"`
}

func (r *DescribeRecordFilterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordFilterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordGroupListRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 分页开始位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页每页数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRecordGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 分页开始位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页每页数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRecordGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordGroupListResponseParams struct {
	// 分组列表
	GroupList []*RecordGroupInfo `json:"GroupList,omitnil,omitempty" name:"GroupList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordGroupListResponseParams `json:"Response"`
}

func (r *DescribeRecordGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordLineCategoryListRequestParams struct {
	// 要查询线路列表的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 要查询线路列表的域名 ID。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain。可以通过接口 DescribeDomainList 查到所有的 Domain 以及 DomainId。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeRecordLineCategoryListRequest struct {
	*tchttp.BaseRequest
	
	// 要查询线路列表的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 要查询线路列表的域名 ID。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain。可以通过接口 DescribeDomainList 查到所有的 Domain 以及 DomainId。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeRecordLineCategoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordLineCategoryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordLineCategoryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordLineCategoryListResponseParams struct {
	// 按分类返回的线路列表。
	LineList []*LineItem `json:"LineList,omitnil,omitempty" name:"LineList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordLineCategoryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordLineCategoryListResponseParams `json:"Response"`
}

func (r *DescribeRecordLineCategoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordLineCategoryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordLineListRequestParams struct {
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名等级。
	// + 旧套餐：D_FREE、D_PLUS、D_EXTRA、D_EXPERT、D_ULTRA 分别对应免费套餐、个人豪华、企业1、企业2、企业3。
	// + 新套餐：DP_FREE、DP_PLUS、DP_EXTRA、DP_EXPERT、DP_ULTRA 分别对应新免费、个人专业版、企业创业版、企业标准版、企业旗舰版。
	DomainGrade *string `json:"DomainGrade,omitnil,omitempty" name:"DomainGrade"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeRecordLineListRequest struct {
	*tchttp.BaseRequest
	
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名等级。
	// + 旧套餐：D_FREE、D_PLUS、D_EXTRA、D_EXPERT、D_ULTRA 分别对应免费套餐、个人豪华、企业1、企业2、企业3。
	// + 新套餐：DP_FREE、DP_PLUS、DP_EXTRA、DP_EXPERT、DP_ULTRA 分别对应新免费、个人专业版、企业创业版、企业标准版、企业旗舰版。
	DomainGrade *string `json:"DomainGrade,omitnil,omitempty" name:"DomainGrade"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type DescribeRecordLineListResponseParams struct {
	// 线路列表。
	LineList []*LineInfo `json:"LineList,omitnil,omitempty" name:"LineList"`

	// 线路分组列表。
	LineGroupList []*LineGroupInfo `json:"LineGroupList,omitnil,omitempty" name:"LineGroupList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordLineListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordLineListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRecordListRequestParams struct {
	// 要获取的解析记录所属的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 要获取的解析记录所属的域名Id，如果传了DomainId，系统将会忽略Domain参数。 可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 解析记录的主机头，如果传了此参数，则只会返回此主机头对应的解析记录
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// 获取某种类型的解析记录，如 A，CNAME，NS，AAAA，显性URL，隐性URL，CAA，SPF等
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 获取某条线路名称的解析记录。可以通过接口DescribeRecordLineList查看当前域名允许的线路信息
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 获取某个线路Id对应的解析记录，如果传RecordLineId，系统会忽略RecordLine参数。可以通过接口DescribeRecordLineList查看当前域名允许的线路信息
	RecordLineId *string `json:"RecordLineId,omitnil,omitempty" name:"RecordLineId"`

	// 获取某个分组下的解析记录时，传这个分组Id。
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 通过关键字搜索解析记录，当前支持搜索主机头和记录值
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序字段，支持 name,line,type,value,weight,mx,ttl,updated_on 几个字段。
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序方式，正序：ASC，逆序：DESC。默认值为ASC。
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数量，当前Limit最大支持3000。默认值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRecordListRequest struct {
	*tchttp.BaseRequest
	
	// 要获取的解析记录所属的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 要获取的解析记录所属的域名Id，如果传了DomainId，系统将会忽略Domain参数。 可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 解析记录的主机头，如果传了此参数，则只会返回此主机头对应的解析记录
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// 获取某种类型的解析记录，如 A，CNAME，NS，AAAA，显性URL，隐性URL，CAA，SPF等
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 获取某条线路名称的解析记录。可以通过接口DescribeRecordLineList查看当前域名允许的线路信息
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 获取某个线路Id对应的解析记录，如果传RecordLineId，系统会忽略RecordLine参数。可以通过接口DescribeRecordLineList查看当前域名允许的线路信息
	RecordLineId *string `json:"RecordLineId,omitnil,omitempty" name:"RecordLineId"`

	// 获取某个分组下的解析记录时，传这个分组Id。
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 通过关键字搜索解析记录，当前支持搜索主机头和记录值
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序字段，支持 name,line,type,value,weight,mx,ttl,updated_on 几个字段。
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序方式，正序：ASC，逆序：DESC。默认值为ASC。
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数量，当前Limit最大支持3000。默认值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeRecordListResponseParams struct {
	// 记录的数量统计信息
	RecordCountInfo *RecordCountInfo `json:"RecordCountInfo,omitnil,omitempty" name:"RecordCountInfo"`

	// 获取的记录列表
	RecordList []*RecordListItem `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRecordRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录 ID 。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeRecordRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录 ID 。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type DescribeRecordResponseParams struct {
	// 记录信息
	RecordInfo *RecordInfo `json:"RecordInfo,omitnil,omitempty" name:"RecordInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRecordSnapshotRollbackResultRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 回滚任务 ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeRecordSnapshotRollbackResultRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 回滚任务 ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeRecordSnapshotRollbackResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordSnapshotRollbackResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "JobId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordSnapshotRollbackResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordSnapshotRollbackResultResponseParams struct {
	// 回滚任务 ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 回滚状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 失败的记录信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedRecordList []*SnapshotRecord `json:"FailedRecordList,omitnil,omitempty" name:"FailedRecordList"`

	// 所属域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 回滚进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 回滚剩余时间（单位：分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftMinutes *uint64 `json:"LeftMinutes,omitnil,omitempty" name:"LeftMinutes"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 失败记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Failed *uint64 `json:"Failed,omitnil,omitempty" name:"Failed"`

	// 成功记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Success *uint64 `json:"Success,omitnil,omitempty" name:"Success"`

	// 快照下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordSnapshotRollbackResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordSnapshotRollbackResultResponseParams `json:"Response"`
}

func (r *DescribeRecordSnapshotRollbackResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordSnapshotRollbackResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordTypeRequestParams struct {
	// 域名等级。
	// + 旧套餐：D_FREE、D_PLUS、D_EXTRA、D_EXPERT、D_ULTRA 分别对应免费套餐、个人豪华、企业1、企业2、企业3。
	// + 新套餐：DP_FREE、DP_PLUS、DP_EXTRA、DP_EXPERT、DP_ULTRA 分别对应新免费、个人专业版、企业创业版、企业标准版、企业旗舰版。
	DomainGrade *string `json:"DomainGrade,omitnil,omitempty" name:"DomainGrade"`
}

type DescribeRecordTypeRequest struct {
	*tchttp.BaseRequest
	
	// 域名等级。
	// + 旧套餐：D_FREE、D_PLUS、D_EXTRA、D_EXPERT、D_ULTRA 分别对应免费套餐、个人豪华、企业1、企业2、企业3。
	// + 新套餐：DP_FREE、DP_PLUS、DP_EXTRA、DP_EXPERT、DP_ULTRA 分别对应新免费、个人专业版、企业创业版、企业标准版、企业旗舰版。
	DomainGrade *string `json:"DomainGrade,omitnil,omitempty" name:"DomainGrade"`
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

// Predefined struct for user
type DescribeRecordTypeResponseParams struct {
	// 记录类型列表
	TypeList []*string `json:"TypeList,omitnil,omitempty" name:"TypeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordTypeResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSnapshotConfigRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeSnapshotConfigRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeSnapshotConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotConfigResponseParams struct {
	// 解析快照配置
	SnapshotConfig *SnapshotConfig `json:"SnapshotConfig,omitnil,omitempty" name:"SnapshotConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotConfigResponseParams `json:"Response"`
}

func (r *DescribeSnapshotConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotListRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeSnapshotListRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeSnapshotListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotListResponseParams struct {
	// 分页信息
	Info *SnapshotPageInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// 快照列表
	SnapshotList []*SnapshotInfo `json:"SnapshotList,omitnil,omitempty" name:"SnapshotList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotListResponseParams `json:"Response"`
}

func (r *DescribeSnapshotListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotRollbackResultRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照回滚任务 ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeSnapshotRollbackResultRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照回滚任务 ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeSnapshotRollbackResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotRollbackResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "TaskId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotRollbackResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotRollbackResultResponseParams struct {
	// 快照所属域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 回滚剩余时间（分钟）
	LeftMinutes *uint64 `json:"LeftMinutes,omitnil,omitempty" name:"LeftMinutes"`

	// 回滚进度百分比
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 快照 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 回滚状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 快照回滚任务 ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 成功数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Success *uint64 `json:"Success,omitnil,omitempty" name:"Success"`

	// 失败数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Failed *uint64 `json:"Failed,omitnil,omitempty" name:"Failed"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 失败详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedRecordList []*SnapshotRecord `json:"FailedRecordList,omitnil,omitempty" name:"FailedRecordList"`

	// 快照的下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotRollbackResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotRollbackResultResponseParams `json:"Response"`
}

func (r *DescribeSnapshotRollbackResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotRollbackResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotRollbackTaskRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeSnapshotRollbackTaskRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeSnapshotRollbackTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotRollbackTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotRollbackTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotRollbackTaskResponseParams struct {
	// 快照所属域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 回滚状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 快照回滚任务 ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 总数量
	RecordCount *uint64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 开始回滚时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotRollbackTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotRollbackTaskResponseParams `json:"Response"`
}

func (r *DescribeSnapshotRollbackTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotRollbackTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubdomainAnalyticsRequestParams struct {
	// 要查询解析量的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询的开始时间，格式：YYYY-MM-DD
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询的结束时间，格式：YYYY-MM-DD
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 要查询解析量的子域名
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// DATE:按天维度统计 HOUR:按小时维度统计
	DnsFormat *string `json:"DnsFormat,omitnil,omitempty" name:"DnsFormat"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeSubdomainAnalyticsRequest struct {
	*tchttp.BaseRequest
	
	// 要查询解析量的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询的开始时间，格式：YYYY-MM-DD
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询的结束时间，格式：YYYY-MM-DD
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 要查询解析量的子域名
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// DATE:按天维度统计 HOUR:按小时维度统计
	DnsFormat *string `json:"DnsFormat,omitnil,omitempty" name:"DnsFormat"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeSubdomainAnalyticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubdomainAnalyticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Subdomain")
	delete(f, "DnsFormat")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubdomainAnalyticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubdomainAnalyticsResponseParams struct {
	// 当前统计维度解析量小计
	Data []*DomainAnalyticsDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 子域名解析量统计查询信息
	Info *SubdomainAnalyticsInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// 子域名别名解析量统计信息
	AliasData []*SubdomainAliasAnalyticsItem `json:"AliasData,omitnil,omitempty" name:"AliasData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubdomainAnalyticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubdomainAnalyticsResponseParams `json:"Response"`
}

func (r *DescribeSubdomainAnalyticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubdomainAnalyticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDetailRequestParams struct {

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

// Predefined struct for user
type DescribeUserDetailResponseParams struct {
	// 账户信息
	UserInfo *UserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeVASStatisticRequestParams struct {
	// 域名ID
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DescribeVASStatisticRequest struct {
	*tchttp.BaseRequest
	
	// 域名ID
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DescribeVASStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVASStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVASStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVASStatisticResponseParams struct {
	// 增值服务用量列表
	VASList []*VASStatisticItem `json:"VASList,omitnil,omitempty" name:"VASList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVASStatisticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVASStatisticResponseParams `json:"Response"`
}

func (r *DescribeVASStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVASStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainAliasAnalyticsItem struct {
	// 域名解析量统计查询信息
	Info *DomainAnalyticsInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// 当前统计维度解析量小计
	Data []*DomainAnalyticsDetail `json:"Data,omitnil,omitempty" name:"Data"`
}

type DomainAliasInfo struct {
	// 域名别名ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 域名别名
	DomainAlias *string `json:"DomainAlias,omitnil,omitempty" name:"DomainAlias"`

	// 别名状态：1-DNS不正确；2-正常；3-封禁。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DomainAnalyticsDetail struct {
	// 当前统计维度解析量小计
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 按天统计时，为统计日期
	DateKey *string `json:"DateKey,omitnil,omitempty" name:"DateKey"`

	// 按小时统计时，为统计的当前时间的小时数(0-23)，例：HourKey为23时，统计周期为22点-23点的解析量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HourKey *uint64 `json:"HourKey,omitnil,omitempty" name:"HourKey"`
}

type DomainAnalyticsInfo struct {
	// DATE:按天维度统计 HOUR:按小时维度统计
	DnsFormat *string `json:"DnsFormat,omitnil,omitempty" name:"DnsFormat"`

	// 当前统计周期解析量总计
	DnsTotal *uint64 `json:"DnsTotal,omitnil,omitempty" name:"DnsTotal"`

	// 当前查询的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 当前统计周期开始时间
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 当前统计周期结束时间
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type DomainCountInfo struct {
	// 符合条件的域名数量
	DomainTotal *uint64 `json:"DomainTotal,omitnil,omitempty" name:"DomainTotal"`

	// 用户可以查看的所有域名数量
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// 用户账号添加的域名数量
	MineTotal *uint64 `json:"MineTotal,omitnil,omitempty" name:"MineTotal"`

	// 共享给用户的域名数量
	ShareTotal *uint64 `json:"ShareTotal,omitnil,omitempty" name:"ShareTotal"`

	// 付费域名数量
	VipTotal *uint64 `json:"VipTotal,omitnil,omitempty" name:"VipTotal"`

	// 暂停的域名数量
	PauseTotal *uint64 `json:"PauseTotal,omitnil,omitempty" name:"PauseTotal"`

	// dns设置错误的域名数量
	ErrorTotal *uint64 `json:"ErrorTotal,omitnil,omitempty" name:"ErrorTotal"`

	// 锁定的域名数量
	LockTotal *uint64 `json:"LockTotal,omitnil,omitempty" name:"LockTotal"`

	// 封禁的域名数量
	SpamTotal *uint64 `json:"SpamTotal,omitnil,omitempty" name:"SpamTotal"`

	// 30天内即将到期的域名数量
	VipExpire *uint64 `json:"VipExpire,omitnil,omitempty" name:"VipExpire"`

	// 分享给其它人的域名数量
	ShareOutTotal *uint64 `json:"ShareOutTotal,omitnil,omitempty" name:"ShareOutTotal"`

	// 指定分组内的域名数量
	GroupTotal *uint64 `json:"GroupTotal,omitnil,omitempty" name:"GroupTotal"`
}

type DomainCreateInfo struct {
	// 域名ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名的punycode
	Punycode *string `json:"Punycode,omitnil,omitempty" name:"Punycode"`

	// 域名的NS列表
	GradeNsList []*string `json:"GradeNsList,omitnil,omitempty" name:"GradeNsList"`
}

type DomainInfo struct {
	// 域名ID
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名套餐等级
	Grade *string `json:"Grade,omitnil,omitempty" name:"Grade"`

	// 域名分组ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 是否星标域名
	IsMark *string `json:"IsMark,omitnil,omitempty" name:"IsMark"`

	// TTL(DNS记录缓存时间)
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// cname加速启用状态
	CnameSpeedup *string `json:"CnameSpeedup,omitnil,omitempty" name:"CnameSpeedup"`

	// 域名备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 域名Punycode
	Punycode *string `json:"Punycode,omitnil,omitempty" name:"Punycode"`

	// 域名DNS状态
	DnsStatus *string `json:"DnsStatus,omitnil,omitempty" name:"DnsStatus"`

	// 域名的NS列表
	DnspodNsList []*string `json:"DnspodNsList,omitnil,omitempty" name:"DnspodNsList"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名等级代号
	GradeLevel *uint64 `json:"GradeLevel,omitnil,omitempty" name:"GradeLevel"`

	// 域名所属的用户ID
	UserId *uint64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 是否为付费域名
	IsVip *string `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// 域名所有者的账号
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 域名等级的描述
	GradeTitle *string `json:"GradeTitle,omitnil,omitempty" name:"GradeTitle"`

	// 域名创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 最后操作时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 腾讯云账户Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 域名实际使用的NS列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActualNsList []*string `json:"ActualNsList,omitnil,omitempty" name:"ActualNsList"`

	// 域名的记录数量
	RecordCount *uint64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 域名所有者的账户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerNick *string `json:"OwnerNick,omitnil,omitempty" name:"OwnerNick"`

	// 是否在付费套餐宽限期
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsGracePeriod *string `json:"IsGracePeriod,omitnil,omitempty" name:"IsGracePeriod"`

	// 是否在付费套餐缓冲期
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipBuffered *string `json:"VipBuffered,omitnil,omitempty" name:"VipBuffered"`

	// VIP套餐有效期开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipStartAt *string `json:"VipStartAt,omitnil,omitempty" name:"VipStartAt"`

	// VIP套餐有效期结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipEndAt *string `json:"VipEndAt,omitnil,omitempty" name:"VipEndAt"`

	// VIP套餐自动续费标识。可能的值为：default-默认；no-不自动续费；yes-自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipAutoRenew *string `json:"VipAutoRenew,omitnil,omitempty" name:"VipAutoRenew"`

	// VIP套餐资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipResourceId *string `json:"VipResourceId,omitnil,omitempty" name:"VipResourceId"`

	// 是否是子域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSubDomain *bool `json:"IsSubDomain,omitnil,omitempty" name:"IsSubDomain"`

	// 域名关联的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*TagItem `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 是否启用搜索引擎推送
	SearchEnginePush *string `json:"SearchEnginePush,omitnil,omitempty" name:"SearchEnginePush"`

	// 是否开启辅助 DNS
	SlaveDNS *string `json:"SlaveDNS,omitnil,omitempty" name:"SlaveDNS"`
}

type DomainListItem struct {
	// 系统分配给域名的唯一标识
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名的原始格式
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 域名的状态，正常：ENABLE，暂停：PAUSE，封禁：SPAM
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名默认的解析记录默认TTL值
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 是否开启CNAME加速，开启：ENABLE，未开启：DISABLE
	CNAMESpeedup *string `json:"CNAMESpeedup,omitnil,omitempty" name:"CNAMESpeedup"`

	// DNS 设置状态，错误：DNSERROR，正常：空字符串
	DNSStatus *string `json:"DNSStatus,omitnil,omitempty" name:"DNSStatus"`

	// 域名的套餐等级代码
	Grade *string `json:"Grade,omitnil,omitempty" name:"Grade"`

	// 域名所属的分组Id
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 是否开启搜索引擎推送优化，是：YES，否：NO
	SearchEnginePush *string `json:"SearchEnginePush,omitnil,omitempty" name:"SearchEnginePush"`

	// 域名备注说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 经过punycode编码后的域名格式
	Punycode *string `json:"Punycode,omitnil,omitempty" name:"Punycode"`

	// 系统为域名分配的有效DNS
	EffectiveDNS []*string `json:"EffectiveDNS,omitnil,omitempty" name:"EffectiveDNS"`

	// 域名套餐等级对应的序号
	GradeLevel *uint64 `json:"GradeLevel,omitnil,omitempty" name:"GradeLevel"`

	// 套餐名称
	GradeTitle *string `json:"GradeTitle,omitnil,omitempty" name:"GradeTitle"`

	// 是否是付费套餐
	IsVip *string `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// 付费套餐开通时间
	VipStartAt *string `json:"VipStartAt,omitnil,omitempty" name:"VipStartAt"`

	// 付费套餐到期时间
	VipEndAt *string `json:"VipEndAt,omitnil,omitempty" name:"VipEndAt"`

	// 域名是否开通VIP自动续费，是：YES，否：NO，默认：DEFAULT
	VipAutoRenew *string `json:"VipAutoRenew,omitnil,omitempty" name:"VipAutoRenew"`

	// 域名下的记录数量
	RecordCount *uint64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 域名添加时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 域名更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 域名所属账号
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 域名关联的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*TagItem `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type DomainShareInfo struct {
	// 域名共享对象的账号
	ShareTo *string `json:"ShareTo,omitnil,omitempty" name:"ShareTo"`

	// 共享模式，“rw”：可读写。 “r”:：只读
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 共享状态“enabled”：共享成功。“pending”：共享到的账号不存在, 等待注册
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DownloadSnapshotRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照记录 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type DownloadSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照记录 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *DownloadSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SnapshotId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadSnapshotResponseParams struct {
	// 快照下载链接
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DownloadSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *DownloadSnapshotResponseParams `json:"Response"`
}

func (r *DownloadSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfo struct {
	// 分组ID
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 分组类型
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 该分组中域名个数
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type KeyValue struct {
	// 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LineGroupInfo struct {
	// 线路分组ID
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`

	// 线路分组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分组类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 线路分组包含的线路列表
	LineList []*string `json:"LineList,omitnil,omitempty" name:"LineList"`
}

type LineInfo struct {
	// 线路名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 线路ID
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`
}

type LineItem struct {
	// 解析线路名称。
	LineName *string `json:"LineName,omitnil,omitempty" name:"LineName"`

	// 解析线路 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`

	// 当前线路在当前域名下是否可用。
	Useful *bool `json:"Useful,omitnil,omitempty" name:"Useful"`

	// 当前线路最低套餐等级要求。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *string `json:"Grade,omitnil,omitempty" name:"Grade"`

	// 当前线路分类下的子线路列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubGroup []*LineItem `json:"SubGroup,omitnil,omitempty" name:"SubGroup"`

	// 自定义线路分组内包含的线路。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lines []*string `json:"Lines,omitnil,omitempty" name:"Lines"`
}

type LockInfo struct {
	// 域名 ID
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名解锁码
	LockCode *string `json:"LockCode,omitnil,omitempty" name:"LockCode"`

	// 域名自动解锁日期
	LockEnd *string `json:"LockEnd,omitnil,omitempty" name:"LockEnd"`
}

// Predefined struct for user
type ModifyDomainCustomLineRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 修改后的自定义线路名称，如果不修改名称，需要和PreName保持一致
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义线路IP段，用-分割
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 修改前的自定义线路名称
	PreName *string `json:"PreName,omitnil,omitempty" name:"PreName"`

	// 域名ID，如果传了DomainId，系统将会忽略Domain参数，优先使用DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifyDomainCustomLineRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 修改后的自定义线路名称，如果不修改名称，需要和PreName保持一致
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自定义线路IP段，用-分割
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 修改前的自定义线路名称
	PreName *string `json:"PreName,omitnil,omitempty" name:"PreName"`

	// 域名ID，如果传了DomainId，系统将会忽略Domain参数，优先使用DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *ModifyDomainCustomLineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainCustomLineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Name")
	delete(f, "Area")
	delete(f, "PreName")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainCustomLineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainCustomLineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainCustomLineResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainCustomLineResponseParams `json:"Response"`
}

func (r *ModifyDomainCustomLineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainCustomLineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainLockRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名要锁定的天数，最多可锁定的天数可以通过获取域名权限接口获取。
	LockDays *uint64 `json:"LockDays,omitnil,omitempty" name:"LockDays"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifyDomainLockRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名要锁定的天数，最多可锁定的天数可以通过获取域名权限接口获取。
	LockDays *uint64 `json:"LockDays,omitnil,omitempty" name:"LockDays"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type ModifyDomainLockResponseParams struct {
	// 域名锁定信息
	LockInfo *LockInfo `json:"LockInfo,omitnil,omitempty" name:"LockInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainLockResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainLockResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDomainOwnerRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名需要转入的账号，支持Uin或者邮箱格式
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifyDomainOwnerRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名需要转入的账号，支持Uin或者邮箱格式
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type ModifyDomainOwnerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainOwnerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainOwnerResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDomainRemarkRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名备注，删除备注请提交空内容。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyDomainRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名备注，删除备注请提交空内容。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
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

// Predefined struct for user
type ModifyDomainRemarkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainRemarkResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDomainStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名状态，”enable” 、”disable” 分别代表启用和暂停
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifyDomainStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名状态，”enable” 、”disable” 分别代表启用和暂停
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type ModifyDomainStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDomainUnlockRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名解锁码，锁定的时候会返回。
	LockCode *string `json:"LockCode,omitnil,omitempty" name:"LockCode"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifyDomainUnlockRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名解锁码，锁定的时候会返回。
	LockCode *string `json:"LockCode,omitnil,omitempty" name:"LockCode"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type ModifyDomainUnlockResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainUnlockResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainUnlockResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyDynamicDNSRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录ID。 可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 记录线路，通过 API 记录线路获得，中文，比如：默认。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 记录值，如 IP : 200.200.200.200， CNAME : cname.dnspod.com.， MX : mail.dnspod.com.。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 线路的 ID，通过 API 记录线路获得，英文字符串，比如：10=1。参数RecordLineId优先级高于RecordLine，如果同时传递二者，优先使用RecordLineId参数。
	RecordLineId *string `json:"RecordLineId,omitnil,omitempty" name:"RecordLineId"`

	// TTL值，如果不传，默认为域名的TTL值。
	Ttl *uint64 `json:"Ttl,omitnil,omitempty" name:"Ttl"`
}

type ModifyDynamicDNSRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录ID。 可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 记录线路，通过 API 记录线路获得，中文，比如：默认。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 记录值，如 IP : 200.200.200.200， CNAME : cname.dnspod.com.， MX : mail.dnspod.com.。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 线路的 ID，通过 API 记录线路获得，英文字符串，比如：10=1。参数RecordLineId优先级高于RecordLine，如果同时传递二者，优先使用RecordLineId参数。
	RecordLineId *string `json:"RecordLineId,omitnil,omitempty" name:"RecordLineId"`

	// TTL值，如果不传，默认为域名的TTL值。
	Ttl *uint64 `json:"Ttl,omitnil,omitempty" name:"Ttl"`
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
	delete(f, "Ttl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDynamicDNSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDynamicDNSResponseParams struct {
	// 记录ID
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDynamicDNSResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDynamicDNSResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyPackageAutoRenewRequestParams struct {
	// 资源ID。可以在控制台查看所有的资源
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// enable 开启自动续费；disable 关闭自动续费
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyPackageAutoRenewRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID。可以在控制台查看所有的资源
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// enable 开启自动续费；disable 关闭自动续费
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyPackageAutoRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPackageAutoRenewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPackageAutoRenewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPackageAutoRenewResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPackageAutoRenewResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPackageAutoRenewResponseParams `json:"Response"`
}

func (r *ModifyPackageAutoRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPackageAutoRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordBatchDetail struct {
	// 见RecordInfoBatchModify
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordList []*BatchRecordInfo `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 任务编号
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainGrade *string `json:"DomainGrade,omitnil,omitempty" name:"DomainGrade"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 该条任务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 域名ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

// Predefined struct for user
type ModifyRecordBatchRequestParams struct {
	// 记录ID数组。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordIdList []*uint64 `json:"RecordIdList,omitnil,omitempty" name:"RecordIdList"`

	// 要修改的字段，可选值为 [“sub_domain”、”record_type”、”area”、”value”、”mx”、”ttl”、”status”] 中的某一个。
	Change *string `json:"Change,omitnil,omitempty" name:"Change"`

	// 修改为，具体依赖 change 字段，必填参数。
	ChangeTo *string `json:"ChangeTo,omitnil,omitempty" name:"ChangeTo"`

	// 要修改到的记录值，仅当 change 字段为 “record_type” 时为必填参数。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// MX记录优先级，仅当修改为 MX 记录时为必填参数。
	MX *string `json:"MX,omitnil,omitempty" name:"MX"`
}

type ModifyRecordBatchRequest struct {
	*tchttp.BaseRequest
	
	// 记录ID数组。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordIdList []*uint64 `json:"RecordIdList,omitnil,omitempty" name:"RecordIdList"`

	// 要修改的字段，可选值为 [“sub_domain”、”record_type”、”area”、”value”、”mx”、”ttl”、”status”] 中的某一个。
	Change *string `json:"Change,omitnil,omitempty" name:"Change"`

	// 修改为，具体依赖 change 字段，必填参数。
	ChangeTo *string `json:"ChangeTo,omitnil,omitempty" name:"ChangeTo"`

	// 要修改到的记录值，仅当 change 字段为 “record_type” 时为必填参数。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// MX记录优先级，仅当修改为 MX 记录时为必填参数。
	MX *string `json:"MX,omitnil,omitempty" name:"MX"`
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

// Predefined struct for user
type ModifyRecordBatchResponseParams struct {
	// 批量任务ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 见modifyRecordBatchDetail
	DetailList []*ModifyRecordBatchDetail `json:"DetailList,omitnil,omitempty" name:"DetailList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRecordBatchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordBatchResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRecordFieldsRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录 ID 。
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 要修改的记录属性和值，支持：sub_domain，record_line，record_line_id，record_type，value，ttl，status，mx，weight
	FieldList []*KeyValue `json:"FieldList,omitnil,omitempty" name:"FieldList"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifyRecordFieldsRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录 ID 。
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 要修改的记录属性和值，支持：sub_domain，record_line，record_line_id，record_type，value，ttl，status，mx，weight
	FieldList []*KeyValue `json:"FieldList,omitnil,omitempty" name:"FieldList"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *ModifyRecordFieldsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordFieldsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordId")
	delete(f, "FieldList")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordFieldsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordFieldsResponseParams struct {
	// 记录ID
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRecordFieldsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordFieldsResponseParams `json:"Response"`
}

func (r *ModifyRecordFieldsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordFieldsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordGroupRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 要修改的分组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifyRecordGroupRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 要修改的分组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *ModifyRecordGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "GroupName")
	delete(f, "GroupId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordGroupResponseParams struct {
	// 修改的分组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRecordGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordGroupResponseParams `json:"Response"`
}

func (r *ModifyRecordGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordRemarkRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录 ID 。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 解析记录备注，删除备注请提交空内容。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyRecordRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录 ID 。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 解析记录备注，删除备注请提交空内容。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
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

// Predefined struct for user
type ModifyRecordRemarkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRecordRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordRemarkResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRecordRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录类型，通过 API 记录类型获得，大写英文，比如：A 。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录线路，通过 API 记录线路获得，中文，比如：默认。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 记录值，如 IP : 200.200.200.200， CNAME : cname.dnspod.com.， MX : mail.dnspod.com.。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 记录 ID 。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 线路的 ID，通过 API 记录线路获得，英文字符串，比如：10=1。参数RecordLineId优先级高于RecordLine，如果同时传递二者，优先使用RecordLineId参数。
	RecordLineId *string `json:"RecordLineId,omitnil,omitempty" name:"RecordLineId"`

	// MX 优先级，当记录类型是 MX 时有效，范围1-20，MX 记录时必选。
	MX *uint64 `json:"MX,omitnil,omitempty" name:"MX"`

	// TTL，范围1-604800，不同等级域名最小值不同。
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 权重信息，0到100的整数。0 表示关闭，不传该参数，表示不设置权重信息。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 记录初始状态，取值范围为 ENABLE 和 DISABLE 。默认为 ENABLE ，如果传入 DISABLE，解析不会生效，也不会验证负载均衡的限制。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 记录的备注信息。传空删除备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 开启DNSSEC时，强制将其它记录修改为CNAME/URL记录
	DnssecConflictMode *string `json:"DnssecConflictMode,omitnil,omitempty" name:"DnssecConflictMode"`
}

type ModifyRecordRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录类型，通过 API 记录类型获得，大写英文，比如：A 。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录线路，通过 API 记录线路获得，中文，比如：默认。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 记录值，如 IP : 200.200.200.200， CNAME : cname.dnspod.com.， MX : mail.dnspod.com.。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 记录 ID 。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 线路的 ID，通过 API 记录线路获得，英文字符串，比如：10=1。参数RecordLineId优先级高于RecordLine，如果同时传递二者，优先使用RecordLineId参数。
	RecordLineId *string `json:"RecordLineId,omitnil,omitempty" name:"RecordLineId"`

	// MX 优先级，当记录类型是 MX 时有效，范围1-20，MX 记录时必选。
	MX *uint64 `json:"MX,omitnil,omitempty" name:"MX"`

	// TTL，范围1-604800，不同等级域名最小值不同。
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 权重信息，0到100的整数。0 表示关闭，不传该参数，表示不设置权重信息。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 记录初始状态，取值范围为 ENABLE 和 DISABLE 。默认为 ENABLE ，如果传入 DISABLE，解析不会生效，也不会验证负载均衡的限制。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 记录的备注信息。传空删除备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 开启DNSSEC时，强制将其它记录修改为CNAME/URL记录
	DnssecConflictMode *string `json:"DnssecConflictMode,omitnil,omitempty" name:"DnssecConflictMode"`
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
	delete(f, "Remark")
	delete(f, "DnssecConflictMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordResponseParams struct {
	// 记录ID
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRecordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRecordStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录 ID 。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 记录的状态。取值范围为 ENABLE 和 DISABLE。如果传入 DISABLE，解析不会生效，也不会验证负载均衡的限制。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifyRecordStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录 ID 。可以通过接口DescribeRecordList查到所有的解析记录列表以及对应的RecordId
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 记录的状态。取值范围为 ENABLE 和 DISABLE。如果传入 DISABLE，解析不会生效，也不会验证负载均衡的限制。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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

// Predefined struct for user
type ModifyRecordStatusResponseParams struct {
	// 记录ID。
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRecordStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRecordToGroupRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 记录 ID，多个 ID 用竖线“|”分割
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifyRecordToGroupRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 记录 ID，多个 ID 用竖线“|”分割
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *ModifyRecordToGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordToGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "GroupId")
	delete(f, "RecordId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordToGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordToGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRecordToGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordToGroupResponseParams `json:"Response"`
}

func (r *ModifyRecordToGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordToGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotConfigRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 备件间隔：空字符串-不备份，half_hour-每半小时，hourly-每小时，daily-每天，monthly-每月
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifySnapshotConfigRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 备件间隔：空字符串-不备份，half_hour-每半小时，hourly-每小时，daily-每天，monthly-每月
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *ModifySnapshotConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Period")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySnapshotConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotConfigResponseParams `json:"Response"`
}

func (r *ModifySnapshotConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubdomainStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录类型。允许的值为A、CNAME、MX、TXT、NS、AAAA、SPF、SRV、CAA、URL、URL1。若要传多个，用英文逗号分隔，例如A,TXT,CNAME。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录状态。允许的值为disable。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`
}

type ModifySubdomainStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 记录类型。允许的值为A、CNAME、MX、TXT、NS、AAAA、SPF、SRV、CAA、URL、URL1。若要传多个，用英文逗号分隔，例如A,TXT,CNAME。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 记录状态。允许的值为disable。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。可以通过接口DescribeDomainList查到所有的Domain以及DomainId
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 主机记录，如 www，如果不传，默认为 @。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`
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

// Predefined struct for user
type ModifySubdomainStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySubdomainStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubdomainStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyVasAutoRenewStatusRequestParams struct {
	// 资源ID。可以从控制台查看所有的资源
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// enable 开启自动续费；disable 关闭自动续费
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyVasAutoRenewStatusRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID。可以从控制台查看所有的资源
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// enable 开启自动续费；disable 关闭自动续费
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyVasAutoRenewStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVasAutoRenewStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVasAutoRenewStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVasAutoRenewStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVasAutoRenewStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVasAutoRenewStatusResponseParams `json:"Response"`
}

func (r *ModifyVasAutoRenewStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVasAutoRenewStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PackageDetailItem struct {
	// 套餐原价
	RealPrice *uint64 `json:"RealPrice,omitnil,omitempty" name:"RealPrice"`

	// 可更换域名次数
	ChangedTimes *uint64 `json:"ChangedTimes,omitnil,omitempty" name:"ChangedTimes"`

	// 允许设置的最小 TTL 值
	MinTtl *uint64 `json:"MinTtl,omitnil,omitempty" name:"MinTtl"`

	// 负载均衡数量
	RecordRoll *uint64 `json:"RecordRoll,omitnil,omitempty" name:"RecordRoll"`

	// 子域名级数
	SubDomainLevel *uint64 `json:"SubDomainLevel,omitnil,omitempty" name:"SubDomainLevel"`

	// 泛解析级数
	MaxWildcard *uint64 `json:"MaxWildcard,omitnil,omitempty" name:"MaxWildcard"`

	// DNS 服务集群个数
	DnsServerRegion *string `json:"DnsServerRegion,omitnil,omitempty" name:"DnsServerRegion"`

	// 套餐名称
	DomainGradeCn *string `json:"DomainGradeCn,omitnil,omitempty" name:"DomainGradeCn"`

	// 套餐代号
	GradeLevel *uint64 `json:"GradeLevel,omitnil,omitempty" name:"GradeLevel"`

	// 套餐对应的 NS
	Ns []*string `json:"Ns,omitnil,omitempty" name:"Ns"`

	// 套餐代码
	DomainGrade *string `json:"DomainGrade,omitnil,omitempty" name:"DomainGrade"`
}

// Predefined struct for user
type PayOrderWithBalanceRequestParams struct {
	// 需要支付的大订单号数组
	BigDealIdList []*string `json:"BigDealIdList,omitnil,omitempty" name:"BigDealIdList"`

	// 代金券ID数组。可以从控制台查到拥有的代金券
	VoucherIdList []*string `json:"VoucherIdList,omitnil,omitempty" name:"VoucherIdList"`
}

type PayOrderWithBalanceRequest struct {
	*tchttp.BaseRequest
	
	// 需要支付的大订单号数组
	BigDealIdList []*string `json:"BigDealIdList,omitnil,omitempty" name:"BigDealIdList"`

	// 代金券ID数组。可以从控制台查到拥有的代金券
	VoucherIdList []*string `json:"VoucherIdList,omitnil,omitempty" name:"VoucherIdList"`
}

func (r *PayOrderWithBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PayOrderWithBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BigDealIdList")
	delete(f, "VoucherIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PayOrderWithBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PayOrderWithBalanceResponseParams struct {
	// 此次操作支付成功的订单id数组
	DealIdList []*string `json:"DealIdList,omitnil,omitempty" name:"DealIdList"`

	// 此次操作支付成功的大订单号数组
	BigDealIdList []*string `json:"BigDealIdList,omitnil,omitempty" name:"BigDealIdList"`

	// 此次操作支付成功的订单号数组
	DealNameList []*string `json:"DealNameList,omitnil,omitempty" name:"DealNameList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PayOrderWithBalanceResponse struct {
	*tchttp.BaseResponse
	Response *PayOrderWithBalanceResponseParams `json:"Response"`
}

func (r *PayOrderWithBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PayOrderWithBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreviewDetail struct {
	// 域名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 域名套餐代码
	Grade *string `json:"Grade,omitnil,omitempty" name:"Grade"`

	// 域名套餐名称
	GradeTitle *string `json:"GradeTitle,omitnil,omitempty" name:"GradeTitle"`

	// 域名记录数
	Records *uint64 `json:"Records,omitnil,omitempty" name:"Records"`

	// 域名停靠状态。0 未开启 1 已开启 2 已暂停
	DomainParkingStatus *uint64 `json:"DomainParkingStatus,omitnil,omitempty" name:"DomainParkingStatus"`

	// 自定义线路数量
	LineCount *uint64 `json:"LineCount,omitnil,omitempty" name:"LineCount"`

	// 自定义线路分组数量
	LineGroupCount *uint64 `json:"LineGroupCount,omitnil,omitempty" name:"LineGroupCount"`

	// 域名别名数量
	AliasCount *uint64 `json:"AliasCount,omitnil,omitempty" name:"AliasCount"`

	// 允许添加的最大域名别名数量
	MaxAliasCount *uint64 `json:"MaxAliasCount,omitnil,omitempty" name:"MaxAliasCount"`

	// 昨天的解析量
	ResolveCount *uint64 `json:"ResolveCount,omitnil,omitempty" name:"ResolveCount"`

	// 增值服务数量
	VASCount *uint64 `json:"VASCount,omitnil,omitempty" name:"VASCount"`
}

type PurviewInfo struct {
	// 权限名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权限值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RecordCountInfo struct {
	// 子域名数量
	SubdomainCount *uint64 `json:"SubdomainCount,omitnil,omitempty" name:"SubdomainCount"`

	// 列表返回的记录数
	ListCount *uint64 `json:"ListCount,omitnil,omitempty" name:"ListCount"`

	// 总的记录数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type RecordGroupInfo struct {
	// 分组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 分组类型：system-系统；user-用户
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`
}

type RecordInfo struct {
	// 记录 ID 。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 子域名(主机记录)。
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录类型, 详见 DescribeRecordType 接口。
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 解析记录的线路，详见 DescribeRecordLineList 接口。
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 解析记录的线路 ID ，详见 DescribeRecordLineList 接口。
	RecordLineId *string `json:"RecordLineId,omitnil,omitempty" name:"RecordLineId"`

	// 记录值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 记录权重值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 记录的 MX 记录值，非 MX 记录类型，默认为 0。
	MX *uint64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 记录的 TTL 值。
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 记录状态。0表示禁用，1表示启用。
	Enabled *uint64 `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 该记录的 D 监控状态。
	// "Ok" : 服务器正常。
	// "Warn" : 该记录有报警, 服务器返回 4XX。
	// "Down" : 服务器宕机。
	// "" : 该记录未开启 D 监控。
	MonitorStatus *string `json:"MonitorStatus,omitnil,omitempty" name:"MonitorStatus"`

	// 记录的备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 记录最后更新时间。
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 域名 ID 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type RecordListItem struct {
	// 记录Id
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 记录值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 记录状态，启用：ENABLE，暂停：DISABLE
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 主机名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 记录线路
	Line *string `json:"Line,omitnil,omitempty" name:"Line"`

	// 线路Id
	LineId *string `json:"LineId,omitnil,omitempty" name:"LineId"`

	// 记录类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 记录权重，用于负载均衡记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 记录监控状态，正常：OK，告警：WARN，宕机：DOWN，未设置监控或监控暂停则为空
	MonitorStatus *string `json:"MonitorStatus,omitnil,omitempty" name:"MonitorStatus"`

	// 记录备注说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 记录缓存时间
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// MX值，只有MX记录有
	// 注意：此字段可能返回 null，表示取不到有效值。
	MX *uint64 `json:"MX,omitnil,omitempty" name:"MX"`

	// 是否是默认的ns记录
	DefaultNS *bool `json:"DefaultNS,omitnil,omitempty" name:"DefaultNS"`
}

// Predefined struct for user
type RollbackRecordSnapshotRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 解析记录信息
	RecordList []*SnapshotRecord `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 之前的快照回滚任务 ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type RollbackRecordSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 解析记录信息
	RecordList []*SnapshotRecord `json:"RecordList,omitnil,omitempty" name:"RecordList"`

	// 之前的快照回滚任务 ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *RollbackRecordSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackRecordSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SnapshotId")
	delete(f, "RecordList")
	delete(f, "TaskId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackRecordSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackRecordSnapshotResponseParams struct {
	// 回滚任务 ID
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackRecordSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *RollbackRecordSnapshotResponseParams `json:"Response"`
}

func (r *RollbackRecordSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackRecordSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackSnapshotRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照记录 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 指定需要回滚的记录
	RecordList []*SnapshotRecord `json:"RecordList,omitnil,omitempty" name:"RecordList"`
}

type RollbackSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照记录 ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 域名 ID 。参数 DomainId 优先级比参数 Domain 高，如果传递参数 DomainId 将忽略参数 Domain 。
	DomainId *uint64 `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 指定需要回滚的记录
	RecordList []*SnapshotRecord `json:"RecordList,omitnil,omitempty" name:"RecordList"`
}

func (r *RollbackSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SnapshotId")
	delete(f, "DomainId")
	delete(f, "RecordList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackSnapshotResponseParams struct {
	// 回滚任务 ID，用来查询回滚状态
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *RollbackSnapshotResponseParams `json:"Response"`
}

func (r *RollbackSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotConfig struct {
	// 配置类型：空字符串-不备份，half_hour-每半小时，hourly-每小时，daily-每天，monthly-每月
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 添加时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 所属域名 ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 配置 ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 快照数量
	SnapshotCount *uint64 `json:"SnapshotCount,omitnil,omitempty" name:"SnapshotCount"`

	// 状态：enable-启用，disable-禁用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`
}

type SnapshotInfo struct {
	// 快照的对象存储地址
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 添加时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 所属域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 快照记录 ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 域名解析记录数
	RecordCount *string `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 状态：normal-正常，create-备份中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type SnapshotPageInfo struct {
	// 快照总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type SnapshotRecord struct {
	// 子域名
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// 记录类型
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// 解析线路
	RecordLine *string `json:"RecordLine,omitnil,omitempty" name:"RecordLine"`

	// 解析值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// TTL(秒)
	TTL *string `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 解析记录 ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// MX优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	MX *string `json:"MX,omitnil,omitempty" name:"MX"`

	// 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *string `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type SubdomainAliasAnalyticsItem struct {
	// 子域名解析量统计查询信息
	Info *SubdomainAnalyticsInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// 当前统计维度解析量小计
	Data []*DomainAnalyticsDetail `json:"Data,omitnil,omitempty" name:"Data"`
}

type SubdomainAnalyticsInfo struct {
	// DATE:按天维度统计 HOUR:按小时维度统计
	DnsFormat *string `json:"DnsFormat,omitnil,omitempty" name:"DnsFormat"`

	// 当前统计周期解析量总计
	DnsTotal *uint64 `json:"DnsTotal,omitnil,omitempty" name:"DnsTotal"`

	// 当前查询的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 当前统计周期开始时间
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 当前统计周期结束时间
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 当前统计的子域名
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`
}

type TagItem struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagItemFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签键
	TagValue []*string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type UserInfo struct {
	// 用户昵称
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// 用户ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户账号, 邮箱格式
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 账号状态：”enabled”: 正常；”disabled”: 被封禁
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 电话号码
	Telephone *string `json:"Telephone,omitnil,omitempty" name:"Telephone"`

	// 邮箱是否通过验证：”yes”: 通过；”no”: 未通过
	EmailVerified *string `json:"EmailVerified,omitnil,omitempty" name:"EmailVerified"`

	// 手机是否通过验证：”yes”: 通过；”no”: 未通过
	TelephoneVerified *string `json:"TelephoneVerified,omitnil,omitempty" name:"TelephoneVerified"`

	// 账号等级, 按照用户账号下域名等级排序, 选取一个最高等级为账号等级, 具体对应情况参见域名等级。
	UserGrade *string `json:"UserGrade,omitnil,omitempty" name:"UserGrade"`

	// 用户名称, 企业用户对应为公司名称
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 是否绑定微信：”yes”: 通过；”no”: 未通过
	WechatBinded *string `json:"WechatBinded,omitnil,omitempty" name:"WechatBinded"`

	// 用户UIN
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 所属 DNS 服务器
	FreeNs []*string `json:"FreeNs,omitnil,omitempty" name:"FreeNs"`
}

type VASStatisticItem struct {
	// 增值服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 增值服务标识
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 增值服务最大用量
	LimitCount *uint64 `json:"LimitCount,omitnil,omitempty" name:"LimitCount"`

	// 增值服务已使用的用量
	UseCount *uint64 `json:"UseCount,omitnil,omitempty" name:"UseCount"`
}

type WhoisContact struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Admin *WhoisContactAddress `json:"Admin,omitnil,omitempty" name:"Admin"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Billing *WhoisContactAddress `json:"Billing,omitnil,omitempty" name:"Billing"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Registrant *WhoisContactAddress `json:"Registrant,omitnil,omitempty" name:"Registrant"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tech *WhoisContactAddress `json:"Tech,omitnil,omitempty" name:"Tech"`
}

type WhoisContactAddress struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fax *string `json:"Fax,omitnil,omitempty" name:"Fax"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaxExt *string `json:"FaxExt,omitnil,omitempty" name:"FaxExt"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Handle *string `json:"Handle,omitnil,omitempty" name:"Handle"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Street *string `json:"Street,omitnil,omitempty" name:"Street"`
}

type WhoisInfo struct {
	// 联系信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contacts *WhoisContact `json:"Contacts,omitnil,omitempty" name:"Contacts"`

	// 域名注册时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationDate *string `json:"CreationDate,omitnil,omitempty" name:"CreationDate"`

	// 域名到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

	// 是否是在腾讯云注册的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQcloud *bool `json:"IsQcloud,omitnil,omitempty" name:"IsQcloud"`

	// 是否当前操作账号注册的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQcloudOwner *bool `json:"IsQcloudOwner,omitnil,omitempty" name:"IsQcloudOwner"`

	// 域名配置的NS
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`

	// Whois原始信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Raw []*string `json:"Raw,omitnil,omitempty" name:"Raw"`

	// 域名注册商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Registrar []*string `json:"Registrar,omitnil,omitempty" name:"Registrar"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 更新日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedDate *string `json:"UpdatedDate,omitnil,omitempty" name:"UpdatedDate"`

	// dnssec
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dnssec *string `json:"Dnssec,omitnil,omitempty" name:"Dnssec"`
}
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
