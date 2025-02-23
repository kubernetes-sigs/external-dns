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
	form := &url.Values{}
	form.Add("action", "get")
	if p.token != "" {
		form.Add("token", p.token)
	}

	url, err := p.urlForRecordType(rtype)
	if err != nil {
		return nil, err
	}

	log.Debugf("Listing %s records from %s", rtype, url)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	body, err := p.do(req)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	raw, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	// Response is a map of "data" to a list of lists where the first element in each
	// list is the dns name and the second is the target.
	// Pi-Hole does not allow for a record to have multiple targets.
	var res map[string][][]string
	if err := json.Unmarshal(raw, &res); err != nil {
		// Unfortunately this could also just mean we needed to authenticate (still returns a 200).
		// Thankfully the body is a short and concise error.
		err = errors.New(string(raw))
		if strings.Contains(err.Error(), "expired") && p.cfg.Password != "" {
			// Try to fetch a new token and redo the request.
			// Full error message at time of writing:
			// "Not allowed (login session invalid or expired, please relogin on the Pi-hole dashboard)!"
			log.Info("Pihole token has expired, fetching a new one")
			if err := p.retrieveNewToken(ctx); err != nil {
				return nil, err
			}
			return p.listRecords(ctx, rtype)
		}
		// Return raw body as error.
		return nil, err
	}

	out := make([]*endpoint.Endpoint, 0)
	data, ok := res["data"]
	if !ok {
		return out, nil
	}
loop:
	for _, rec := range data {
		name := rec[0]
		target := rec[1]
		if !p.cfg.DomainFilter.Match(name) {
			log.Debugf("Skipping %s that does not match domain filter", name)
			continue
		}
		switch rtype {
		case endpoint.RecordTypeA:
			if strings.Contains(target, ":") {
				continue loop
			}
		case endpoint.RecordTypeAAAA:
			if strings.Contains(target, ".") {
				continue loop
			}
		}
		out = append(out, &endpoint.Endpoint{
			DNSName:    name,
			Targets:    []string{target},
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
	return fmt.Sprintf("%s/admin/scripts/pi-hole/php/customdns.php", p.cfg.Server)
}

func (p *piholeClientV6) cnameRecordsScript() string {
	return fmt.Sprintf("%s/admin/scripts/pi-hole/php/customcname.php", p.cfg.Server)
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
	} `json:"session"`
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

	form := &url.Values{}
	form.Add("pw", p.cfg.Password)
	url := fmt.Sprintf("%s/api/auth", p.cfg.Server)
	log.Debugf("Fetching new token from %s", url)

	// Define the JSON payload
	jsonData := []byte(`{"password":"` + p.cfg.Password + `"}`)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")

	body, err := p.do(req)
	if err != nil {
		return err
	}
	defer body.Close()

	jRes, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}
	// Parse JSON response
	var apiResponse ApiAuthResponse
	err = json.Unmarshal(jRes, &apiResponse)
	// Set the token
	p.token = apiResponse.Session.SID

	return err
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

func (p *piholeClientV6) do(req *http.Request) (io.ReadCloser, error) {
	res, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		return nil, fmt.Errorf("received non-200 status code from request: %s", res.Status)
	}
	return res.Body, nil
}
