package common

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type OIDCRoleArnProvider struct {
	region           string
	providerId       string
	webIdentityToken string
	roleArn          string
	roleSessionName  string
	durationSeconds  int64
	Endpoint         string
	beforeRefresh    func(provider *OIDCRoleArnProvider) error
}

type oidcStsRsp struct {
	Response struct {
		Credentials struct {
			Token        string `json:"Token"`
			TmpSecretId  string `json:"TmpSecretId"`
			TmpSecretKey string `json:"TmpSecretKey"`
		} `json:"Credentials"`
		ExpiredTime int       `json:"ExpiredTime"`
		Expiration  time.Time `json:"Expiration"`
		RequestId   string    `json:"RequestId"`
	} `json:"Response"`
}

func NewOIDCRoleArnProvider(region, providerId, webIdentityToken, roleArn, roleSessionName string, durationSeconds int64) *OIDCRoleArnProvider {
	return &OIDCRoleArnProvider{
		region:           region,
		providerId:       providerId,
		webIdentityToken: webIdentityToken,
		roleArn:          roleArn,
		roleSessionName:  roleSessionName,
		durationSeconds:  durationSeconds,
	}
}

// DefaultTkeOIDCRoleArnProvider returns a OIDCRoleArnProvider with some default options:
//  1. providerId will be environment var TKE_PROVIDER_ID
//  2. webIdentityToken will be the content of file specified by env TKE_WEB_IDENTITY_TOKEN_FILE
//  3. roleArn will be env TKE_ROLE_ARN
//  4. roleSessionName will be "tencentcloud-go-sdk-" + timestamp
//  5. durationSeconds will be 7200s
func DefaultTkeOIDCRoleArnProvider() (*OIDCRoleArnProvider, error) {
	beforeRefresh := func(provider *OIDCRoleArnProvider) error {
		reg := os.Getenv("TKE_REGION")
		if reg == "" {
			return errors.New("env TKE_REGION not exist")
		}

		providerId := os.Getenv("TKE_PROVIDER_ID")
		if providerId == "" {
			return errors.New("env TKE_PROVIDER_ID not exist")
		}

		tokenFile := os.Getenv("TKE_WEB_IDENTITY_TOKEN_FILE")
		if tokenFile == "" {
			return errors.New("env TKE_WEB_IDENTITY_TOKEN_FILE not exist")
		}
		tokenBytes, err := ioutil.ReadFile(tokenFile)
		if err != nil {
			return err
		}

		roleArn := os.Getenv("TKE_ROLE_ARN")
		if roleArn == "" {
			return errors.New("env TKE_ROLE_ARN not exist")
		}

		sessionName := defaultSessionName + strconv.FormatInt(time.Now().UnixNano()/1000, 10)

		provider.region = region
		provider.providerId = providerId
		provider.webIdentityToken = string(tokenBytes)
		provider.roleArn = roleArn
		provider.roleSessionName = sessionName
		return nil
	}

	provider := &OIDCRoleArnProvider{
		beforeRefresh:   beforeRefresh,
		durationSeconds: defaultDurationSeconds,
	}
	return provider, provider.beforeRefresh(provider)
}

func (r *OIDCRoleArnProvider) GetCredential() (CredentialIface, error) {
	const (
		service = "sts"
		version = "2018-08-13"
		action  = "AssumeRoleWithWebIdentity"
	)

	if r.beforeRefresh != nil {
		err := r.beforeRefresh(r)
		if err != nil {
			return nil, err
		}
	}

	if r.durationSeconds > 43200 || r.durationSeconds <= 0 {
		return nil, tcerr.NewTencentCloudSDKError(creErr, "AssumeRoleWithWebIdentity durationSeconds should be in the range of 0~43200s", "")
	}
	providerEndpoint := r.Endpoint
	if providerEndpoint == "" {
		providerEndpoint = endpoint
	}
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = providerEndpoint
	cpf.HttpProfile.ReqMethod = "POST"

	client := NewCommonClient(nil, r.region, cpf)
	request := tchttp.NewCommonRequest(service, version, action)
	request.SetSkipSign(true)

	params := map[string]interface{}{
		"ProviderId":       r.providerId,
		"WebIdentityToken": r.webIdentityToken,
		"RoleArn":          r.roleArn,
		"RoleSessionName":  r.roleSessionName,
		"DurationSeconds":  r.durationSeconds,
	}
	err := request.SetActionParameters(params)
	if err != nil {
		return nil, err
	}

	response := tchttp.NewCommonResponse()
	err = client.Send(request, response)
	if err != nil {
		return nil, err
	}
	rspSt := new(oidcStsRsp)

	if err = json.Unmarshal(response.GetBody(), rspSt); err != nil {
		return nil, tcerr.NewTencentCloudSDKError(creErr, err.Error(), "")
	}

	return &RoleArnCredential{
		roleArn:         r.roleArn,
		roleSessionName: r.roleSessionName,
		durationSeconds: r.durationSeconds,
		expiredTime:     int64(rspSt.Response.ExpiredTime) - r.durationSeconds/10,
		token:           rspSt.Response.Credentials.Token,
		tmpSecretId:     rspSt.Response.Credentials.TmpSecretId,
		tmpSecretKey:    rspSt.Response.Credentials.TmpSecretKey,
		source:          r,
	}, nil
}
