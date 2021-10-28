package ibclient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

type AuthConfig struct {
	Username string
	Password string

	ClientCert []byte
	ClientKey  []byte
}

type HostConfig struct {
	Host    string
	Version string
	Port    string
}

type TransportConfig struct {
	SslVerify           bool
	certPool            *x509.CertPool
	HttpRequestTimeout  time.Duration // in seconds
	HttpPoolConnections int
	ProxyUrl            *url.URL
}

func NewTransportConfig(sslVerify string, httpRequestTimeout int, httpPoolConnections int) (cfg TransportConfig) {
	switch {
	case "false" == strings.ToLower(sslVerify):
		cfg.SslVerify = false
	case "true" == strings.ToLower(sslVerify):
		cfg.SslVerify = true
	default:
		caPool := x509.NewCertPool()
		cert, err := ioutil.ReadFile(sslVerify)
		if err != nil {
			log.Printf("Cannot load certificate file '%s'", sslVerify)
			return
		}
		if !caPool.AppendCertsFromPEM(cert) {
			err = fmt.Errorf("cannot append certificate from file '%s'", sslVerify)
			return
		}
		cfg.certPool = caPool
		cfg.SslVerify = true
	}

	cfg.HttpPoolConnections = httpPoolConnections
	cfg.HttpRequestTimeout = time.Duration(httpRequestTimeout)

	return
}

type HttpRequestBuilder interface {
	Init(HostConfig, AuthConfig)
	BuildUrl(r RequestType, objType string, ref string, returnFields []string, queryParams *QueryParams) (urlStr string)
	BuildBody(r RequestType, obj IBObject) (jsonStr []byte)
	BuildRequest(r RequestType, obj IBObject, ref string, queryParams *QueryParams) (req *http.Request, err error)
}

type HttpRequestor interface {
	Init(AuthConfig, TransportConfig)
	SendRequest(*http.Request) ([]byte, error)
}

type WapiRequestBuilder struct {
	hostCfg HostConfig
	authCfg AuthConfig
}

type WapiHttpRequestor struct {
	client http.Client
}

type IBConnector interface {
	CreateObject(obj IBObject) (ref string, err error)
	GetObject(obj IBObject, ref string, queryParams *QueryParams, res interface{}) error
	DeleteObject(ref string) (refRes string, err error)
	UpdateObject(obj IBObject, ref string) (refRes string, err error)
}

type Connector struct {
	hostCfg        HostConfig
	authCfg        AuthConfig
	transportCfg   TransportConfig
	requestBuilder HttpRequestBuilder
	requestor      HttpRequestor
}

type RequestType int

const (
	CREATE RequestType = iota
	GET
	DELETE
	UPDATE
)

func (r RequestType) toMethod() string {
	switch r {
	case CREATE:
		return "POST"
	case GET:
		return "GET"
	case DELETE:
		return "DELETE"
	case UPDATE:
		return "PUT"
	}

	return ""
}

func getHTTPResponseError(resp *http.Response) error {
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	msg := fmt.Sprintf("WAPI request error: %d('%s')\nContents:\n%s\n", resp.StatusCode, resp.Status, content)
	log.Printf(msg)
	if resp.StatusCode == http.StatusNotFound {
		return NewNotFoundError(msg)
	}
	return errors.New(msg)
}

func (whr *WapiHttpRequestor) Init(authCfg AuthConfig, trCfg TransportConfig) {
	var certList []tls.Certificate

	clientAuthType := tls.NoClientCert

	if authCfg.ClientKey != nil && authCfg.ClientCert != nil {
		cert, err := tls.X509KeyPair(authCfg.ClientCert, authCfg.ClientKey)
		if err != nil {
			log.Fatal(err)
		}

		certList = []tls.Certificate{cert}
		clientAuthType = tls.RequestClientCert
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:            trCfg.certPool,
			ClientAuth:         clientAuthType,
			Certificates:       certList,
			InsecureSkipVerify: !trCfg.SslVerify,
			Renegotiation:      tls.RenegotiateOnceAsClient,
		},
		MaxIdleConnsPerHost: trCfg.HttpPoolConnections,
		Proxy:               http.ProxyFromEnvironment,
	}

	if trCfg.ProxyUrl != nil {
		tr.Proxy = http.ProxyURL(trCfg.ProxyUrl)
	}

	// All users of cookiejar should import "golang.org/x/net/publicsuffix"
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	whr.client = http.Client{
		Jar:       jar,
		Transport: tr,
		Timeout:   trCfg.HttpRequestTimeout * time.Second,
	}
}

func (whr *WapiHttpRequestor) SendRequest(req *http.Request) (res []byte, err error) {
	var resp *http.Response
	resp, err = whr.client.Do(req)
	if err != nil {
		return
	} else if !(resp.StatusCode == http.StatusOK ||
		(resp.StatusCode == http.StatusCreated &&
			req.Method == CREATE.toMethod())) {
		err := getHTTPResponseError(resp)
		return nil, err
	}
	defer resp.Body.Close()
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Http Reponse ioutil.ReadAll() Error: '%s'", err)
		return
	}

	return
}

func NewWapiRequestBuilder(hostCfg HostConfig, authCfg AuthConfig) (*WapiRequestBuilder, error) {
	wrb := WapiRequestBuilder{
		hostCfg: hostCfg,
		authCfg: authCfg,
	}

	return &wrb, nil
}

func (wrb *WapiRequestBuilder) Init(hostCfg HostConfig, authCfg AuthConfig) {
	wrb.hostCfg = hostCfg
	wrb.authCfg = authCfg
}

func (wrb *WapiRequestBuilder) BuildUrl(t RequestType, objType string, ref string, returnFields []string, queryParams *QueryParams) (urlStr string) {
	path := []string{"wapi", "v" + wrb.hostCfg.Version}
	if len(ref) > 0 {
		path = append(path, ref)
	} else {
		path = append(path, objType)
	}

	qry := ""
	vals := url.Values{}
	if t == GET {
		if len(returnFields) > 0 {
			vals.Set("_return_fields", strings.Join(returnFields, ","))
		}
		if queryParams != nil {
			// TODO need to get this from individual objects in future
			if queryParams.forceProxy {
				vals.Set("_proxy_search", "GM")
			}
			for k, v := range queryParams.searchFields {
				vals.Set(k, v)
			}
		}

		qry = vals.Encode()
	}

	u := url.URL{
		Scheme:   "https",
		Host:     wrb.hostCfg.Host + ":" + wrb.hostCfg.Port,
		Path:     strings.Join(path, "/"),
		RawQuery: qry,
	}

	return u.String()
}

func (wrb *WapiRequestBuilder) BuildBody(t RequestType, obj IBObject) []byte {
	var objJSON []byte
	var err error

	objJSON, err = json.Marshal(obj)
	if err != nil {
		log.Printf("Cannot marshal object '%s': %s", obj, err)
		return nil
	}

	eaSearch := obj.EaSearch()
	if t == GET && len(eaSearch) > 0 {
		eaSearchJSON, err := json.Marshal(eaSearch)
		if err != nil {
			log.Printf("Cannot marshal EA Search attributes. '%s'\n", err)
			return nil
		}
		objJSON = append(append(objJSON[:len(objJSON)-1], byte(',')), eaSearchJSON[1:]...)
	}

	return objJSON
}

func (wrb *WapiRequestBuilder) BuildRequest(t RequestType, obj IBObject, ref string, queryParams *QueryParams) (req *http.Request, err error) {
	var (
		objType      string
		returnFields []string
	)
	if obj != nil {
		objType = obj.ObjectType()
		returnFields = obj.ReturnFields()
	}
	urlStr := wrb.BuildUrl(t, objType, ref, returnFields, queryParams)

	var bodyStr []byte
	if obj != nil && (t == CREATE || t == UPDATE) {
		bodyStr = wrb.BuildBody(t, obj)
	}

	req, err = http.NewRequest(t.toMethod(), urlStr, bytes.NewBuffer(bodyStr))
	if err != nil {
		log.Printf("err1: '%s'", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if wrb.authCfg.Username != "" {
		req.SetBasicAuth(wrb.authCfg.Username, wrb.authCfg.Password)
	}

	return
}

func (c *Connector) makeRequest(t RequestType, obj IBObject, ref string, queryParams *QueryParams) (res []byte, err error) {
	var req *http.Request
	req, err = c.requestBuilder.BuildRequest(t, obj, ref, queryParams)
	if err != nil {
		return
	}
	res, err = c.requestor.SendRequest(req)
	if err != nil {
		if queryParams != nil {
			/* Forcing the request to redirect to Grid Master by making forcedProxy=true */
			queryParams.forceProxy = true
			req, err = c.requestBuilder.BuildRequest(t, obj, ref, queryParams)
			if err != nil {
				return
			}
			res, err = c.requestor.SendRequest(req)
		} else {
			return nil, err
		}
	}

	return
}

func (c *Connector) CreateObject(obj IBObject) (ref string, err error) {
	ref = ""
	queryParams := NewQueryParams(false, nil)
	resp, err := c.makeRequest(CREATE, obj, "", queryParams)
	if err != nil || len(resp) == 0 {
		log.Printf("CreateObject request error: '%s'\n", err)
		return
	}

	err = json.Unmarshal(resp, &ref)
	if err != nil {
		log.Printf("cannot unmarshall '%s', err: '%s'\n", string(resp), err)
		return
	}

	return
}

func (c *Connector) GetObject(
	// TODO: distinguish between "not found" and other kinds of errors.

	obj IBObject, ref string,
	queryParams *QueryParams, res interface{}) (err error) {

	resp, err := c.makeRequest(GET, obj, ref, queryParams)
	if err != nil {
		return
	}
	//to check empty underlying value of interface
	var result interface{}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		log.Printf("cannot unmarshall to check empty value '%s': '%s'\n", string(resp), err)
	}

	var data []interface{}
	if resp == nil || (reflect.TypeOf(result) == reflect.TypeOf(data) && len(result.([]interface{})) == 0) {
		if queryParams == nil {
			err = NewNotFoundError("requested object not found")
			return
		}
		queryParams.forceProxy = true
		resp, err = c.makeRequest(GET, obj, ref, queryParams)
	}
	if err != nil {
		log.Printf("GetObject request error: '%s'\n", err)
	}
	err = json.Unmarshal(resp, res)
	if err != nil {
		log.Printf("cannot unmarshall '%s', err: '%s'\n", string(resp), err)
		return
	}

	if string(resp) == "[]" {
		return NewNotFoundError("not found")
	}

	return
}

func (c *Connector) DeleteObject(ref string) (refRes string, err error) {
	refRes = ""
	queryParams := NewQueryParams(false, nil)
	resp, err := c.makeRequest(DELETE, nil, ref, queryParams)
	if err != nil {
		log.Printf("DeleteObject request error: '%s'\n", err)
		return
	}

	err = json.Unmarshal(resp, &refRes)
	if err != nil {
		log.Printf("cannot unmarshall '%s': '%s'\n", string(resp), err)
		return
	}

	return
}

func (c *Connector) UpdateObject(obj IBObject, ref string) (refRes string, err error) {
	queryParams := NewQueryParams(false, nil)
	refRes = ""
	resp, err := c.makeRequest(UPDATE, obj, ref, queryParams)
	if err != nil {
		log.Printf("failed to update object %s: %s", obj.ObjectType(), err)
		return
	}

	err = json.Unmarshal(resp, &refRes)
	if err != nil {
		log.Printf("cannot unmarshall update object response'%s', err: '%s'\n", string(resp), err)
		return
	}
	return
}

// Logout sends a request to invalidate the ibapauth cookie and should
// be used in a defer statement after the Connector has been successfully
// initialized.
func (c *Connector) Logout() (err error) {
	queryParams := NewQueryParams(false, nil)
	_, err = c.makeRequest(CREATE, nil, "logout", queryParams)
	if err != nil {
		log.Printf("Logout request error: '%s'\n", err)
	}

	return
}

var ValidateConnector = validateConnector

func validateConnector(c *Connector) (err error) {
	// GET UserProfile request is used here to validate connector's basic auth and reachability.
	// TODO: It seems to be broken, needs to be fixed.
	//var response []UserProfile
	//userprofile := NewUserProfile(UserProfile{})
	//err = c.GetObject(userprofile, "", &response)
	//if err != nil {
	//	log.Printf("Failed to connect to the Grid, err: %s \n", err)
	//}
	return
}

func NewConnector(hostConfig HostConfig, authCfg AuthConfig, transportConfig TransportConfig,
	requestBuilder HttpRequestBuilder, requestor HttpRequestor) (res *Connector, err error) {
	res = nil

	connector := &Connector{
		hostCfg:      hostConfig,
		authCfg:      authCfg,
		transportCfg: transportConfig,
	}

	//connector.requestBuilder = WapiRequestBuilder{WaipHostConfig: connector.hostCfg}
	connector.requestBuilder = requestBuilder
	connector.requestBuilder.Init(connector.hostCfg, connector.authCfg)

	connector.requestor = requestor
	connector.requestor.Init(connector.authCfg, connector.transportCfg)

	res = connector
	err = ValidateConnector(connector)
	return
}
