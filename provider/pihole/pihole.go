package pihole

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// Environment variables.
const (
	PiholeServerEnvVar     = "PIHOLE_SERVER"
	PiholePasswordEnvVar   = "PIHOLE_PASSWORD"
	PiholeSkipVerifyEnvVar = "PIHOLE_TLS_INSECURE_SKIP_VERIFY"
)

// ErrNoPiholeServer is returned when there is no Pihole server configured
// in the environment.
var ErrNoPiholeServer = fmt.Errorf("no %s found in the environment", PiholeServerEnvVar)

// PiholeProvider is an implementation of Provider for Pi-hole Local DNS.
type PiholeProvider struct {
	provider.BaseProvider

	server, passw string
	token         string
	client        *http.Client
	domainFilter  endpoint.DomainFilter
	dryRun        bool
}

// NewPiholeProvider initializes a new Pi-hole Local DNS based Provider.
func NewPiholeProvider(domainFilter endpoint.DomainFilter, dryRun bool) (*PiholeProvider, error) {
	server, ok := os.LookupEnv(PiholeServerEnvVar)
	if !ok {
		return nil, ErrNoPiholeServer
	}

	// Setup a persistent cookiejar for storing PHP session information
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return nil, err
	}
	// Setup an HTTP client using the cookiejar
	cl := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: os.Getenv(PiholeSkipVerifyEnvVar) == "true",
			},
		},
	}

	return &PiholeProvider{
		server:       server,
		passw:        os.Getenv(PiholePasswordEnvVar), // This can be blank to signify an unprotected Pi-hole DNS server.
		client:       cl,
		domainFilter: domainFilter,
		dryRun:       dryRun,
	}, nil
}

func (p *PiholeProvider) aRecordsScript() string {
	return fmt.Sprintf("%s/admin/scripts/pi-hole/php/customdns.php", p.server)
}

func (p *PiholeProvider) cnameRecordsScript() string {
	return fmt.Sprintf("%s/admin/scripts/pi-hole/php/customcname.php", p.server)
}

func (p *PiholeProvider) urlForRecordType(rtype string) (string, error) {
	switch rtype {
	case endpoint.RecordTypeA:
		return p.aRecordsScript(), nil
	case endpoint.RecordTypeCNAME:
		return p.cnameRecordsScript(), nil
	default:
		return "", fmt.Errorf("unsupported record type: %s", rtype)
	}
}

// Records implements Provider, populating a slice of endpoints from
// Pi-Hole local DNS.
func (p *PiholeProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	if p.passw != "" {
		// Retrieve a new token on every request until catch and retry logic
		// is implemented for expired tokens.
		if err := p.retrieveNewToken(ctx); err != nil {
			return nil, err
		}
	}

	aRecords, err := p.listRecords(ctx, endpoint.RecordTypeA)
	if err != nil {
		return nil, err
	}
	cnameRecords, err := p.listRecords(ctx, endpoint.RecordTypeCNAME)
	if err != nil {
		return nil, err
	}

	return append(aRecords, cnameRecords...), nil
}

func (p *PiholeProvider) listRecords(ctx context.Context, rtype string) ([]*endpoint.Endpoint, error) {
	form := &url.Values{}
	form.Add("action", "get")
	form.Add("token", p.token)

	url, err := p.urlForRecordType(rtype)
	if err != nil {
		return nil, err
	}

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
	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	// Response is a map of "data" to a list of lists where the first element in each
	// list is the dns name and the second is the target.
	// Pi-Hole does not allow for a record to have multiple targets.
	var res map[string][][]string
	if err := json.Unmarshal(raw, &res); err != nil {
		// Unfortunately this could also just mean we needed to authenticate (still returns a 200).
		// Return raw body as error.
		return nil, errors.New(string(raw))
	}

	out := make([]*endpoint.Endpoint, 0)
	data, ok := res["data"]
	if !ok {
		return out, nil
	}
	for _, rec := range data {
		out = append(out, &endpoint.Endpoint{
			DNSName:    rec[0],
			Targets:    []string{rec[1]},
			RecordType: rtype,
			RecordTTL:  300, // It actually is ignored
		})
	}

	return out, nil
}

// ApplyChanges implements Provider, syncing desired state with the Pi-hole server Local DNS.
func (p *PiholeProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if p.passw != "" {
		// Retrieve a new token on every request until catch and retry logic
		// is implemented for expired tokens.
		if err := p.retrieveNewToken(ctx); err != nil {
			return err
		}
	}

	// Handle deletions first - there are no endpoints for updating in place.
	for _, ep := range changes.Delete {
		if !p.domainFilter.MatchParent(ep.Targets[0]) {
			continue
		}
		if err := p.apply(ctx, "delete", ep); err != nil {
			return err
		}
	}
	for _, ep := range changes.UpdateOld {
		if !p.domainFilter.MatchParent(ep.Targets[0]) {
			continue
		}
		if err := p.apply(ctx, "delete", ep); err != nil {
			return err
		}
	}

	// Handle desired state
	for _, ep := range changes.Create {
		if !p.domainFilter.MatchParent(ep.Targets[0]) {
			continue
		}
		if err := p.apply(ctx, "add", ep); err != nil {
			return err
		}
	}
	for _, ep := range changes.UpdateNew {
		if !p.domainFilter.MatchParent(ep.Targets[0]) {
			continue
		}
		if err := p.apply(ctx, "add", ep); err != nil {
			return err
		}
	}
	return nil
}

type actionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (p *PiholeProvider) apply(ctx context.Context, action string, ep *endpoint.Endpoint) error {
	if p.dryRun {
		return nil
	}

	uraw, err := p.urlForRecordType(ep.RecordType)
	if err != nil {
		log.Warnf("Skipping unsupported endpoint %s %s %v", ep.DNSName, ep.RecordType, ep.Targets)
		return nil
	}
	form := p.newDNSActionForm(action, ep)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uraw, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	body, err := p.do(req)
	if err != nil {
		return err
	}
	defer body.Close()

	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return nil
	}

	var res actionResponse
	if err := json.Unmarshal(raw, &res); err != nil {
		// Unfortunately this could also just mean a generic server error. Return the raw body.
		return errors.New(string(raw))
	}

	if !res.Success {
		return errors.New(res.Message)
	}

	return nil
}

func (p *PiholeProvider) newDNSActionForm(action string, ep *endpoint.Endpoint) *url.Values {
	form := &url.Values{}
	form.Add("action", action)
	form.Add("domain", ep.DNSName)
	switch ep.RecordType {
	case endpoint.RecordTypeA:
		form.Add("ip", ep.Targets[0])
	case endpoint.RecordTypeCNAME:
		form.Add("target", ep.Targets[0])
	}
	form.Add("token", p.token)
	return form
}

func (p *PiholeProvider) do(req *http.Request) (io.ReadCloser, error) {
	res, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		return nil, fmt.Errorf("received non-200 status code from request: %s", res.Status)
	}
	return res.Body, nil
}
