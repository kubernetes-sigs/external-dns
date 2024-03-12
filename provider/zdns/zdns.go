package zdns

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	//"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

var resourceTxtMap = make(map[string]*endpoint.Endpoint)
var lowerTypeMap = map[string]string{
	"CNAME": "cname",
	"A":     "a",
	"AAAA":  "aaaa",
}

const (
	OPCREATE = "create"
	OPDELETE = "delete"

	EXTERNALDNSOWNER = "external-dns/owner="
)

type ZDNSProvider struct {
	provider.BaseProvider
	client ZDNSAPIClient
}

type ZDNSAPIClient struct {
	config ZDNSConfig
	httpcf httpConfig
}

type ZDNSConfig struct {
	Host  string
	Port  string
	View  string
	Zones string
	Auth  string
	Owner string
}

type httpConfig struct {
	ip      string
	header  map[string]string
	zones   []zone
	hasRoot bool
}

type zone struct {
	name  string
	len   int
	level int
}

type zdnsRRss struct {
	Resources []zdnsRRS `json:"resources"`
}

type zdnsRRS struct {
	Name    string `json:"name"`
	RrsType string `json:"type"`
	Ttl     int    `json:"ttl"`
	Rdata   string `json:"rdata"`
	Id      string `json:"id"`
}

type zdnsResources struct {
	Resources []zdnsResource `json:"resources"`
}

type zdnsResource struct {
	Name string `json:"name"`
}

func (z *zdnsRRS) tostring(name string) []string {
	rrs_strings := []string{}
	datas := strings.Split(z.Rdata, " ")
	pre := name + " " + strconv.Itoa(z.Ttl) + " " + z.RrsType + " "
	for _, data := range datas {
		rrs_strings = append(rrs_strings, pre+data)
	}
	return rrs_strings
}

func NewZDNSProvider(ctx context.Context, config ZDNSConfig) (*ZDNSProvider, error) {
	//2.发送请求进行验证
	reqPath := "https://" + config.Host + ":" + config.Port + "/views/" + config.View + "/zones?"
	reqHead := map[string]string{
		"Authorization": "Basic " + config.Auth,
	}
	var status int
	var respbody []byte
	var err error

	//3.校验服务器是否正常
	for i := 0; i <= 2; i++ {
		log.Info("循环次数:", i)
		status, respbody, err = sendHTTPReqest(http.MethodGet, reqPath, "", reqHead)
		if err != nil {
			log.Error(err.Error())
			if i == 2 {
				return nil, fmt.Errorf("failed to connect to zdns api: %v", err)
			} else {
				continue
			}
		}
		break
	}

	//4.校验权限是否正常
	if status == http.StatusUnauthorized {
		log.Warn("auth 认证失败.")
		return nil, errors.New("auth 认证失败.")
	}

	//5.弱检查，zones是否都存在。
	var newresources zdnsResources
	err = json.Unmarshal(respbody, &newresources)
	if err != nil {
		log.Warn(err.Error())
	}
	getZonesMap := make(map[string]bool)
	for _, resource := range newresources.Resources {
		getZonesMap[resource.Name] = true
	}

	tagetZones := strings.Split(config.Zones, ",")
	noHaveZone := []string{}
	for _, targetZone := range tagetZones {
		if _, ok := getZonesMap[targetZone]; !ok {
			noHaveZone = append(noHaveZone, targetZone)
		}
	}
	if len(noHaveZone) > 0 {
		log.Warn(noHaveZone, "区未创建.")
	}

	zones := []zone{}
	hasRoot := false
	for _, targetZone := range tagetZones {
		var zone zone
		zone.name = targetZone
		if targetZone == "@" {
			hasRoot = true
		}
		zone.len = len(targetZone)
		zone.level = strings.Count(targetZone, ".") + 1
		zones = append(zones, zone)
	}
	log.Debug("zones : ", zones)
	httpcfg := httpConfig{
		ip: "https://" + config.Host + ":" + config.Port,
		header: map[string]string{
			"Authorization": "Basic " + config.Auth,
		},
		zones:   zones,
		hasRoot: hasRoot,
	}

	return &ZDNSProvider{
		client: ZDNSAPIClient{
			config: config,
			httpcf: httpcfg,
		},
	}, nil
}

func sendHTTPReqest(method, url, body string, header map[string]string) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return -1, nil, err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	if method != http.MethodGet {
		req.Header.Add("Content-Type", "application/json")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := http.Client{
		Timeout:   time.Duration(3 * time.Second),
		Transport: tr,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	respByte, _ := ioutil.ReadAll(resp.Body)
	return status, respByte, nil
}

func (z *ZDNSProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, _ error) {
	endpoints = []*endpoint.Endpoint{}
	resourceTxtMap = make(map[string]*endpoint.Endpoint)
	resourceMap := make(map[string]*endpoint.Endpoint)
	fmt.Println(z.client.httpcf.zones)
	for _, zone := range z.client.httpcf.zones {
		respbody := z.getZoneRecords(zone.name)
		fmt.Println(string(respbody))
		var zdnsRrss zdnsRRss
		err := json.Unmarshal(respbody, &zdnsRrss)
		if err != nil {
			log.Warn(err.Error())
		}
		fmt.Println(zdnsRrss)
		for _, rrs := range zdnsRrss.Resources {
			is_ok, ep := rrsToEndpoint(rrs)
			if is_ok {
				//txt 和 A，AAAA，CNAME分别处理
				//txt类型为说明记录包含ownerid
				//A,AAAA,CNAME为正式记录
				if ep.RecordType == "TXT" {
					resourceTxtMap[ep.DNSName] = &ep
				} else {
					lowerType := lowerTypeMap[ep.RecordType]
					if _, ok := resourceMap[lowerType+"-"+ep.DNSName]; ok {
						resourceMap[lowerType+"-"+ep.DNSName].Targets = append(resourceMap[lowerType+"-"+ep.DNSName].Targets, ep.Targets...)
						resourceMap[lowerType+"-"+ep.DNSName].ProviderSpecific = append(resourceMap[lowerType+"-"+ep.DNSName].ProviderSpecific, ep.ProviderSpecific...)
					} else {
						resourceMap[lowerType+"-"+ep.DNSName] = &ep
					}
				}
			}
		}
	}

	owner := ""
	for txtkey, txtvalue := range resourceTxtMap {
		owner = ""
		txtvalueString := strings.Join(txtvalue.Targets, ",")
		txtvalueSlice := strings.Split(txtvalueString, ",")
		for _, txt := range txtvalueSlice {
			if strings.Index(txt, EXTERNALDNSOWNER) != -1 {
				owner = strings.TrimPrefix(txt, EXTERNALDNSOWNER)
			}
		}
		if rrs, ok := resourceMap[txtkey]; ok {
			rrs.Labels["owner"] = owner
			resourceMap[txtkey+"#"+"A"] = rrs
		}

		//给TXT记录填上lable,方便后面判断是否为该集群控制的记录
		if owner != "" {
			txtRRS := resourceTxtMap[txtkey]
			txtRRS.Labels["owner"] = owner
			resourceTxtMap[txtkey] = txtRRS
		}

	}

	for _, ep := range resourceMap {
		if value, ok := ep.Labels["owner"]; ok {
			if value == z.client.config.Owner {
				endpoints = append(endpoints, ep)
			}
		}
	}
	return endpoints, nil
}

func (z *ZDNSProvider) getZoneRecords(zone string) []byte {
	url := z.client.httpcf.ip + "/views/" + z.client.config.View + "/zones/" + zone + "/rrs?"
	status, respbody, err := sendHTTPReqest(http.MethodGet, url, "", z.client.httpcf.header)
	if err != nil {
		log.Warn(err.Error())
		return nil
	}

	if status == http.StatusOK {
		return respbody
	} else {
		log.Warn(url, string(respbody))
	}
	return nil
}

func (z *ZDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	//1.删除
	if len(changes.Delete) > 0 {
		z.deleteRRS(changes.Delete)
	}

	//2.更新
	create, delete := z.analyzeUpdate(changes)
	if len(delete) > 0 {
		z.deleteRRS(delete)
	}
	if len(create) > 0 {
		z.createRRS(create)
	}

	//3.创建
	if len(changes.Create) > 0 {
		z.createRRS(changes.Create)
	}

	return nil
}

func (z *ZDNSProvider) analyzeUpdate(changes *plan.Changes) ([]*endpoint.Endpoint, []*endpoint.Endpoint) {
	var create []*endpoint.Endpoint
	var delete []*endpoint.Endpoint
	newTxt := make(map[string]bool)
	oldTxt := make(map[string]bool)
	for i, epNew := range changes.UpdateNew {
		epOld := changes.UpdateOld[i]
		if epNew.RecordType != "TXT" {
			if epNew.RecordType == "CNAME" {
				//old cname : aaa.kube.
				//new cname : aaa
				//原因我们系统创建cname记录会自动添加区，因为external会删除target记录最后的.
				zone := z.getZoneFromName(epOld.DNSName)
				if zone.name != "" {
					if zone.name == "@" {
						if len(epOld.Targets) > 0 {
							epOld.Targets[0] = epOld.Targets[0][:len(epOld.Targets[0])-1]
						}
					} else {
						if len(epOld.Targets) > 0 {
							epOld.Targets[0] = epOld.Targets[0][:len(epOld.Targets[0])-len(zone.name)-2]
						}
					}
				}
			}

			if z.compareNewOld(epNew, epOld) {
				continue
			}
			newTxtTag := strings.ToLower(epNew.RecordType) + "-" + epNew.DNSName
			if newTxtTag[len(newTxtTag)-1:] != "." {
				newTxtTag += "."
			}
			oldTxtTag := strings.ToLower(epOld.RecordType) + "-" + epOld.DNSName
			if oldTxtTag[len(oldTxtTag)-1:] != "." {
				oldTxtTag += "."
			}
			newTxt[newTxtTag] = true
			oldTxt[oldTxtTag] = true
			create = append(create, epNew)
			delete = append(delete, epOld)
		} else {
			newTxtTag := epNew.DNSName
			if newTxtTag[len(newTxtTag)-1:] != "." {
				newTxtTag += "."
			}
			oldTxtTag := epOld.DNSName
			if oldTxtTag[len(oldTxtTag)-1:] != "." {
				oldTxtTag += "."
			}
			if newTxt[newTxtTag] {
				create = append(create, epNew)
			}
			if oldTxt[oldTxtTag] {
				delete = append(delete, epOld)
			}
		}
	}
	return create, delete
}

func (z *ZDNSProvider) compareNewOld(epNew, epOld *endpoint.Endpoint) bool {
	if epNew.RecordType != epOld.RecordType {
		return false
	}
	if epNew.RecordTTL != epOld.RecordTTL {
		return false
	}
	z.sortSlice(epNew)
	z.sortSlice(epOld)
	if epNew.Targets.String() != epOld.Targets.String() {
		return false
	}
	return true
}

func (z *ZDNSProvider) sortSlice(ep *endpoint.Endpoint) {
	sort.Slice(ep.Targets, func(i, j int) bool {
		targetOne := net.ParseIP(ep.Targets[i])
		targetTwo := net.ParseIP(ep.Targets[j])
		ep.Targets[i] = targetOne.String()
		ep.Targets[j] = targetTwo.String()
		return ep.Targets[i] < ep.Targets[j]
	})
}

func (z *ZDNSProvider) createRRS(endpoints []*endpoint.Endpoint) {
	zoneRRSMap := make(map[string][]string)
	for _, zone := range z.client.httpcf.zones {
		zoneRRSMap[zone.name] = []string{}
	}
	for _, endpoint := range endpoints {
		if !hasAuth(*endpoint, z.client.config.Owner) {
			log.Warn("已有记录，禁止操作 : ", endpoint)
			continue
		}
		rrs := endpointToRRS(*endpoint)
		zone := z.getZoneFromName(rrs.Name)
		if zone.name == "" {
			continue
		}
		if _, ok := zoneRRSMap[zone.name]; ok {
			index := 0
			if zone.name == "@" {
				index = len(rrs.Name)
			} else {
				index = len(rrs.Name) - zone.len - 1
			}
			if index <= 0 {
				log.Error("target name error : ", rrs)
				continue
			}
			zoneRRSMap[zone.name] = append(zoneRRSMap[zone.name], rrs.tostring(rrs.Name[:index])...)
		}
	}
	for zone, rrs := range zoneRRSMap {
		content := ""
		index := 0
		for _, rr := range rrs {
			content += rr + "\n"
			index += 1
			if index >= 200 {
				content64 := base64.StdEncoding.EncodeToString([]byte(content))
				z.sendCreate(zone, content64)
				index = 0
				content = ""
			}
		}
		if content != "" {
			content64 := base64.StdEncoding.EncodeToString([]byte(content))
			z.sendCreate(zone, content64)
		}
	}
}

func (z *ZDNSProvider) deleteRRS(endpoints []*endpoint.Endpoint) {
	zoneRRSMap := make(map[string][]string)
	for _, zone := range z.client.httpcf.zones {
		zoneRRSMap[zone.name] = []string{}
	}
	for _, endpoint := range endpoints {
		rrs := endpointToRRS(*endpoint)
		if rrs.Id == "" {
			continue
		}
		zone := z.getZoneFromName(rrs.Name)
		if zone.name == "" {
			continue
		}
		if _, ok := zoneRRSMap[zone.name]; ok {
			zoneRRSMap[zone.name] = append(zoneRRSMap[zone.name], strings.Split(rrs.Id, ",")...)
		}
	}
	for zone, ids := range zoneRRSMap {
		if len(ids) > 0 {
			z.sendDelete(zone, ids)
		}
	}
}

func (z *ZDNSProvider) sendCreate(zone, content64 string) {
	url := z.client.httpcf.ip + "/views/" + z.client.config.View + "/zones/" + zone + "/rrs"
	body := make(map[string]string)
	body["is_enable"] = "yes"
	body["zone_content"] = content64
	bodyByte, err := json.Marshal(body)
	if err != nil {
		log.Error(err.Error())
	}

	status, respBody, err := sendHTTPReqest(http.MethodPost, url, string(bodyByte), z.client.httpcf.header)
	if status != http.StatusOK || err != nil {
		log.Error("url : ", url, ",status : ", status, ", respbody : ", string(respBody))
	}
	if err != nil {
		log.Error("url : ", url, ",status : ", status, ",err : ", err.Error(), ", respbody : ", string(respBody))

	}
}

func (z *ZDNSProvider) sendDelete(zone string, ids []string) {
	url := z.client.httpcf.ip + "/views/" + z.client.config.View + "/zones/" + zone + "/rrs"
	body := make(map[string]interface{})
	body["link_ptr"] = "no"
	body["link_cname"] = "no"
	body["link_srv"] = "no"
	body["link_mx"] = "no"
	body["ids"] = ids
	bodyByte, err := json.Marshal(body)
	if err != nil {
		log.Error(err.Error())
	}
	status, respBody, err := sendHTTPReqest(http.MethodDelete, url, string(bodyByte), z.client.httpcf.header)
	if status != http.StatusOK || err != nil {
		if err != nil {
			log.Error("url : ", url, ",status : ", status, ",err : ", err.Error(), ", respbody : ", string(respBody))
		} else {
			log.Error("url : ", url, ",status : ", status, ", respbody : ", string(respBody))
		}
	}
}

func (z *ZDNSProvider) getZoneFromName(name string) zone {
	var finalZone zone
	if name[len(name)-1:] == "." {
		name = name[:len(name)-1]
	}

	for sliceIndex, zone := range z.client.httpcf.zones {
		index := strings.LastIndex(name, zone.name)
		if index == -1 {
			continue
		}
		if len(name) == zone.len+index {
			if zone.level > finalZone.level {
				finalZone = z.client.httpcf.zones[sliceIndex]
			}
		}
	}
	if finalZone.name == "" && z.client.httpcf.hasRoot {
		finalZone.name = "@"
		finalZone.len = 1
	}
	return finalZone
}

func rrsToEndpoint(rrs zdnsRRS) (bool, endpoint.Endpoint) {
	var ep endpoint.Endpoint
	if rrs.RrsType != "A" && rrs.RrsType != "AAAA" && rrs.RrsType != "TXT" && rrs.RrsType != "CNAME" {
		return false, ep
	}
	ep.RecordType = rrs.RrsType
	ep.DNSName = rrs.Name
	ep.RecordTTL = endpoint.TTL(rrs.Ttl)
	ep.Targets = append(ep.Targets, rrs.Rdata)
	ps := endpoint.ProviderSpecificProperty{
		Name:  "id",
		Value: rrs.Id,
	}
	ep.ProviderSpecific = append(ep.ProviderSpecific, ps)
	if ep.Labels == nil {
		ep.Labels = endpoint.NewLabels()
	}
	return true, ep
}

func endpointToRRS(endpoint endpoint.Endpoint) zdnsRRS {
	var rrs zdnsRRS
	rrs.Name = endpoint.DNSName
	rrs.RrsType = endpoint.RecordType
	rrs.Rdata = strings.Join(endpoint.Targets, " ")
	rrs.Ttl = int(endpoint.RecordTTL)
	if rrs.RrsType == "TXT" {
		name := endpoint.DNSName
		if name[len(name)-1:] != "." {
			name = name + "."
		}
		if oldTxt, ok := resourceTxtMap[name]; ok {
			psp, is_ok := oldTxt.GetProviderSpecificProperty("id")
			if is_ok {
				rrs.Id = psp
			}
		}
	} else {
		for _, providerSpecific := range endpoint.ProviderSpecific {
			if providerSpecific.Value != "" {
				if rrs.Id == "" {
					rrs.Id = providerSpecific.Value
				} else {
					rrs.Id = rrs.Id + "," + providerSpecific.Value
				}
			}
		}
	}
	return rrs
}

// hasAuth checks if the endpoint has permission for the given owner ID.
//
// Parameters:
// - endpoint: the endpoint to check permission for
// - ownerId: the ID of the owner to compare with
// Return type: bool
func hasAuth(endpoint endpoint.Endpoint, ownerId string) bool {
	//判断本集群是否有权限操作该记录
	name := endpoint.DNSName
	if name[len(name)-1:] != "." {
		name = name + "."
	}
	owner := ""
	txtMapKey := ""
	if endpoint.RecordType == "TXT" {
		txtMapKey = name
	} else {
		txtMapKey = lowerTypeMap[endpoint.RecordType] + "-" + name
	}
	if oldTxt, ok := resourceTxtMap[txtMapKey]; ok {
		owner = oldTxt.Labels["owner"]
	}
	if owner != ownerId && owner != "" {
		return false
	}
	return true
}
