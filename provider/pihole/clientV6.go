package pihole

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/linki/instrumented_http"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider"
)

// piholeClient implements the piholeAPI.
type piholeClientV6 struct {
	cfg        PiholeConfig
	httpClient *http.Client
	token      string
}

// newPiholeClient creates a new Pihole API V6 client.
func newPiholeClientV6(cfg PiholeConfig) (piholeAPI, error) {
	if cfg.Server == "" {
		return nil, ErrNoPiholeServer
	}

	// Setup a persistent cookiejar for storing PHP session information
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return nil, err
	}
	// Setup an HTTP client using the cookiejar
	httpClient := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: cfg.TLSInsecureSkipVerify,
			},
		},
	}
	cl := instrumented_http.NewClient(httpClient, &instrumented_http.Callbacks{})

	p := &piholeClientV6{
		cfg:        cfg,
		httpClient: cl,
	}

	if cfg.Password != "" {
		if err := p.retrieveNewToken(context.Background()); err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p *piholeClientV6) listRecords(ctx context.Context, rtype string) ([]*endpoint.Endpoint, error) {
	apiUrl, err := p.urlForRecordType(rtype)
	if err != nil {
		return nil, err
	}

	log.Debugf("Listing %s records from %s", rtype, apiUrl)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	jRes, err := p.do(req)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var apiResponse ApiRecordsResponse
	err = json.Unmarshal(jRes, &apiResponse)
	if err != nil {
		log.Errorf("error reading response: %s", err)
	}

	// Pi-Hole does not allow for a record to have multiple targets.
	out := make([]*endpoint.Endpoint, 0)
	var results []string
	if endpoint.RecordTypeCNAME == rtype {
		results = apiResponse.Config.DNS.CnameRecords
	} else {
		results = apiResponse.Config.DNS.Hosts
	}

	for _, rec := range results {
		recs := strings.FieldsFunc(rec, func(r rune) bool {
			return r == ' ' || r == ','
		})
		var DNSName string
		var Target string
		// A/AAAA record format is target(IP) DNSName
		DNSName = recs[1]
		Target = recs[0]
		switch rtype {
		case endpoint.RecordTypeA:
			if strings.Contains(recs[0], ":") {
				continue
			}
			break
		case endpoint.RecordTypeAAAA:
			if strings.Contains(recs[0], ".") {
				continue
			}
			break
		case endpoint.RecordTypeCNAME:
			// CNAME format is DNSName,target
			DNSName = recs[0]
			Target = recs[1]
			break
		}

		out = append(out, &endpoint.Endpoint{
			DNSName:    DNSName,
			Targets:    []string{Target},
			RecordType: rtype,
		})
	}
	return out, nil
}

func (p *piholeClientV6) createRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	return p.apply(ctx, "add", ep)
}

func (p *piholeClientV6) deleteRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	return p.apply(ctx, "delete", ep)
}

func (p *piholeClientV6) aRecordsScript() string {
	return fmt.Sprintf("%s/api/config/dns%2Fhosts", p.cfg.Server)
}

func (p *piholeClientV6) cnameRecordsScript() string {
	return fmt.Sprintf("%s/api/config/dns%2FcnameRecords", p.cfg.Server)
}

func (p *piholeClientV6) urlForRecordType(rtype string) (string, error) {
	switch rtype {
	case endpoint.RecordTypeA, endpoint.RecordTypeAAAA:
		return p.aRecordsScript(), nil
	case endpoint.RecordTypeCNAME:
		return p.cnameRecordsScript(), nil
	default:
		return "", fmt.Errorf("unsupported record type: %s", rtype)
	}
}

// ApiAuthResponse Define a struct to match the JSON response /auth/app structure
type ApiAuthResponse struct {
	Session struct {
		Valid    bool   `json:"valid"`
		TOTP     bool   `json:"totp"`
		SID      string `json:"sid"`
		CSRF     string `json:"csrf"`
		Validity int    `json:"validity"`
		Message  string `json:"message"`
	} `json:"session"`
	Took float64 `json:"took"`
}

// ApiErrorResponse Define struct to match the JSON structure
type ApiErrorResponse struct {
	Error struct {
		Key     string `json:"key"`
		Message string `json:"message"`
		Hint    string ` json:"hint"`
	} `json:"error"`
	Took float64 `json:"took"`
}

// ApiRecordsResponse Define struct to match JSON structure
type ApiRecordsResponse struct {
	Config struct {
		DNS struct {
			Hosts        []string `json:"hosts"`
			CnameRecords []string `json:"cnameRecords"`
		} `json:"dns"`
	} `json:"config"`
	Took float64 `json:"took"`
}

func (p *piholeClientV6) apply(ctx context.Context, action string, ep *endpoint.Endpoint) error {
	if !p.cfg.DomainFilter.Match(ep.DNSName) {
		log.Debugf("Skipping %s %s that does not match domain filter", action, ep.DNSName)
		return nil
	}
	url, err := p.urlForRecordType(ep.RecordType)
	if err != nil {
		log.Warnf("Skipping unsupported endpoint %s %s %v", ep.DNSName, ep.RecordType, ep.Targets)
		return nil
	}

	if p.cfg.DryRun {
		log.Infof("DRY RUN: %s %s IN %s -> %s", action, ep.DNSName, ep.RecordType, ep.Targets[0])
		return nil
	}

	log.Infof("%s %s IN %s -> %s", action, ep.DNSName, ep.RecordType, ep.Targets[0])

	form := p.newDNSActionForm(action, ep)
	if strings.Contains(ep.DNSName, "*") {
		return provider.NewSoftError(errors.New("UNSUPPORTED: Pihole DNS names cannot return wildcard"))
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	body, err := p.do(req)
	if err != nil {
		return err
	}
	defer body.Close()

	raw, err := io.ReadAll(body)
	if err != nil {
		return nil
	}

	var res actionResponse
	if err := json.Unmarshal(raw, &res); err != nil {
		// Unfortunately this could also be a generic server or auth error.
		err = errors.New(string(raw))
		if strings.Contains(err.Error(), "expired") && p.cfg.Password != "" {
			// Try to fetch a new token and redo the request.
			log.Info("Pihole token has expired, fetching a new one")
			if err := p.retrieveNewToken(ctx); err != nil {
				return err
			}
			return p.apply(ctx, action, ep)
		}
		// Return raw body as error.
		return err
	}

	if !res.Success {
		return errors.New(res.Message)
	}

	return nil
}

func (p *piholeClientV6) retrieveNewToken(ctx context.Context) error {
	if p.cfg.Password == "" {
		return nil
	}

	apiUrl := fmt.Sprintf("%s/api/auth", p.cfg.Server)
	log.Debugf("Fetching new token from %s", apiUrl)

	// Define the JSON payload
	jsonData := []byte(`{"password":"` + p.cfg.Password + `"}`)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	jRes, err := p.do(req)
	if err != nil {
		return err
	}

	// Parse JSON response
	var apiResponse ApiAuthResponse
	err = json.Unmarshal(jRes, &apiResponse)
	if err != nil {
		fmt.Println("Error reading response:", err)
	} else {
		// Set the token
		if apiResponse.Session.SID == "" {
			p.token = apiResponse.Session.SID
		}
	}
	return err
}

func (p *piholeClientV6) checkTokenValidity(ctx context.Context) (bool, error) {
	if p.token == "" {
		return false, nil
	}

	apiUrl := fmt.Sprintf("%s/api/auth", p.cfg.Server)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		return false, nil
	}

	jRes, err := p.do(req)
	if err != nil {
		return false, nil
	}

	// Parse JSON response
	var apiResponse ApiAuthResponse
	err = json.Unmarshal(jRes, &apiResponse)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return false, err
	}
	return apiResponse.Session.Valid, nil
}

func (p *piholeClientV6) do(req *http.Request) ([]byte, error) {
	req.Header.Add("content-type", "application/json")
	if p.token != "" {
		req.Header.Add("X-FTL-SID", p.token)
	}
	res, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	jRes, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		// Parse JSON response
		var apiError ApiErrorResponse
		err = json.Unmarshal(jRes, &apiError)
		log.Debugf("Error on request %s", req.Body)
		if res.StatusCode == http.StatusUnauthorized && p.token != "" {
			// Try to fetch a new token and redo the request.
			valid, err := p.checkTokenValidity(req.Context())
			if err != nil {
				return nil, err
			}
			if !valid {
				log.Info("Pihole token has expired, fetching a new one")
				if err := p.retrieveNewToken(req.Context()); err != nil {
					return nil, err
				}
				return p.do(req)
			}
		}
		return nil, fmt.Errorf("received %s status code from request: [%s] %s (%s) - %fs", res.Status, apiError.Error.Key, apiError.Error.Message, apiError.Error.Hint, apiError.Took)
	}
	return jRes, nil
}

func (p *piholeClientV6) newDNSActionForm(action string, ep *endpoint.Endpoint) *url.Values {
	form := &url.Values{}
	form.Add("action", action)
	form.Add("domain", ep.DNSName)
	switch ep.RecordType {
	case endpoint.RecordTypeA, endpoint.RecordTypeAAAA:
		form.Add("ip", ep.Targets[0])
	case endpoint.RecordTypeCNAME:
		form.Add("target", ep.Targets[0])
	}
	if p.token != "" {
		form.Add("token", p.token)
	}
	return form
}
