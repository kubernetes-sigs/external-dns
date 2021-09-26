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
	"strings"

	"github.com/linki/instrumented_http"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// ErrNoPiholeServer is returned when there is no Pihole server configured
// in the environment.
var ErrNoPiholeServer = fmt.Errorf("no pihole server found in the environment or flags")

// PiholeProvider is an implementation of Provider for Pi-hole Local DNS.
type PiholeProvider struct {
	provider.BaseProvider
	cfg    PiholeConfig
	token  string
	client *http.Client
}

// PiholeConfig is used for configuring a PiholeProvider.
type PiholeConfig struct {
	// The root URL of the Pi-hole server.
	Server string
	// An optional password if the server is protected.
	Password string
	// Disable verification of TLS certificates.
	TLSInsecureSkipVerify bool
	// A filter to apply when looking up and applying records.
	DomainFilter endpoint.DomainFilter
	// Do nothing and log what would have changed to stdout.
	DryRun bool
}

// NewPiholeProvider initializes a new Pi-hole Local DNS based Provider.
func NewPiholeProvider(cfg PiholeConfig) (*PiholeProvider, error) {
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

	p := &PiholeProvider{
		cfg:    cfg,
		client: cl,
	}

	if cfg.Password != "" {
		if err := p.retrieveNewToken(context.Background()); err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p *PiholeProvider) aRecordsScript() string {
	return fmt.Sprintf("%s/admin/scripts/pi-hole/php/customdns.php", p.cfg.Server)
}

func (p *PiholeProvider) cnameRecordsScript() string {
	return fmt.Sprintf("%s/admin/scripts/pi-hole/php/customcname.php", p.cfg.Server)
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
	for _, rec := range data {
		name := rec[0]
		target := rec[1]
		if !p.cfg.DomainFilter.MatchParent(target) {
			log.Debugf("Skipping %s target that does not match domain filter", target)
			continue
		}
		out = append(out, &endpoint.Endpoint{
			DNSName:    name,
			Targets:    []string{target},
			RecordType: rtype,
		})
	}

	return out, nil
}

// ApplyChanges implements Provider, syncing desired state with the Pi-hole server Local DNS.
func (p *PiholeProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// Handle deletions first - there are no endpoints for updating in place.
	for _, ep := range changes.Delete {
		if !p.cfg.DomainFilter.MatchParent(ep.Targets[0]) {
			log.Debugf("Skipping delete %s that does not match domain filter", ep.Targets[0])
			continue
		}
		if err := p.apply(ctx, "delete", ep); err != nil {
			return err
		}
	}
	for _, ep := range changes.UpdateOld {
		if !p.cfg.DomainFilter.MatchParent(ep.Targets[0]) {
			log.Debugf("Skipping delete %s that does not match domain filter", ep.Targets[0])
			continue
		}
		if err := p.apply(ctx, "delete", ep); err != nil {
			return err
		}
	}

	// Handle desired state
	for _, ep := range changes.Create {
		if !p.cfg.DomainFilter.MatchParent(ep.Targets[0]) {
			log.Debugf("Skipping create %s that does not match domain filter", ep.Targets[0])
			continue
		}
		if err := p.apply(ctx, "add", ep); err != nil {
			return err
		}
	}
	for _, ep := range changes.UpdateNew {
		if !p.cfg.DomainFilter.MatchParent(ep.Targets[0]) {
			log.Debugf("Skipping create %s that does not match domain filter", ep.Targets[0])
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
	url, err := p.urlForRecordType(ep.RecordType)
	if err != nil {
		log.Warnf("Skipping unsupported endpoint %s %s %v", ep.DNSName, ep.RecordType, ep.Targets)
		return nil
	}

	if p.cfg.DryRun {
		log.Infof("DRY RUN: %s %s IN %s -> %s", strings.Title(action), ep.DNSName, ep.RecordType, ep.Targets[0])
		return nil
	}

	log.Infof("%s %s IN %s -> %s", strings.Title(action), ep.DNSName, ep.RecordType, ep.Targets[0])

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

	raw, err := ioutil.ReadAll(body)
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
	if p.token != "" {
		form.Add("token", p.token)
	}
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
