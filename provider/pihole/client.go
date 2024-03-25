/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pihole

import (
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
	"golang.org/x/net/html"

	"sigs.k8s.io/external-dns/endpoint"
)

// piholeAPI declares the "API" actions performed against the Pihole server.
type piholeAPI interface {
	// listRecords returns endpoints for the given record type (A or CNAME).
	listRecords(ctx context.Context, rtype string) ([]*endpoint.Endpoint, error)
	// createRecord will create a new record for the given endpoint.
	createRecord(ctx context.Context, ep *endpoint.Endpoint) error
	// deleteRecord will delete the given record.
	deleteRecord(ctx context.Context, ep *endpoint.Endpoint) error
}

// piholeClient implements the piholeAPI.
type piholeClient struct {
	cfg        PiholeConfig
	httpClient *http.Client
	token      string
}

// newPiholeClient creates a new Pihole API client.
func newPiholeClient(cfg PiholeConfig) (piholeAPI, error) {
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

	p := &piholeClient{
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

func (p *piholeClient) listRecords(ctx context.Context, rtype string) ([]*endpoint.Endpoint, error) {
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

func (p *piholeClient) createRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	return p.apply(ctx, "add", ep)
}

func (p *piholeClient) deleteRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	return p.apply(ctx, "delete", ep)
}

func (p *piholeClient) aRecordsScript() string {
	return fmt.Sprintf("%s/admin/scripts/pi-hole/php/customdns.php", p.cfg.Server)
}

func (p *piholeClient) cnameRecordsScript() string {
	return fmt.Sprintf("%s/admin/scripts/pi-hole/php/customcname.php", p.cfg.Server)
}

func (p *piholeClient) urlForRecordType(rtype string) (string, error) {
	switch rtype {
	case endpoint.RecordTypeA, endpoint.RecordTypeAAAA:
		return p.aRecordsScript(), nil
	case endpoint.RecordTypeCNAME:
		return p.cnameRecordsScript(), nil
	default:
		return "", fmt.Errorf("unsupported record type: %s", rtype)
	}
}

type actionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (p *piholeClient) apply(ctx context.Context, action string, ep *endpoint.Endpoint) error {
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

func (p *piholeClient) retrieveNewToken(ctx context.Context) error {
	if p.cfg.Password == "" {
		return nil
	}

	form := &url.Values{}
	form.Add("pw", p.cfg.Password)
	url := fmt.Sprintf("%s/admin/index.php?login", p.cfg.Server)
	log.Debugf("Fetching new token from %s", url)

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

	// If successful the request will redirect us to an HTML page with a hidden
	// div containing the token...The token gives us access to other PHP
	// endpoints via a form value.
	p.token, err = parseTokenFromLogin(body)
	return err
}

func (p *piholeClient) newDNSActionForm(action string, ep *endpoint.Endpoint) *url.Values {
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

func (p *piholeClient) do(req *http.Request) (io.ReadCloser, error) {
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

func parseTokenFromLogin(body io.ReadCloser) (string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	tokenNode := getElementById(doc, "token")
	if tokenNode == nil {
		return "", errors.New("could not parse token from login response")
	}

	return tokenNode.FirstChild.Data, nil
}

func getAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func hasID(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		s, ok := getAttribute(n, "id")
		if ok && s == id {
			return true
		}
	}
	return false
}

func traverse(n *html.Node, id string) *html.Node {
	if hasID(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := traverse(c, id)
		if result != nil {
			return result
		}
	}

	return nil
}

func getElementById(n *html.Node, id string) *html.Node {
	return traverse(n, id)
}
