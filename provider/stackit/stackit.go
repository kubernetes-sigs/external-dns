package stackit

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/antihax/optional"
	log "github.com/sirupsen/logrus"

	stackitdnsclient "github.com/stackitcloud/stackit-dns-api-client-go"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	CREATE = "CREATE"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

// Config is used to configure the creation of the StackitDNSProvider.
type Config struct {
	BasePath  string
	Token     string
	ProjectId string
}

// StackitDNSProvider implements the DNS provider for STACKIT DNS.
type StackitDNSProvider struct {
	provider.BaseProvider
	Client       Client
	ProjectId    string
	DomainFilter endpoint.DomainFilter
	DryRun       bool
	Workers      int
}

// ErrorMessage is the error message returned by the API.
type ErrorMessage struct {
	Message string `json:"message"`
}

// changeTask is a task that is passed to the worker.
type changeTask struct {
	change *endpoint.Endpoint
	action string
}

// endpointError is a list of endpoints and an error to pass to workers.
type endpointError struct {
	endpoints []*endpoint.Endpoint
	err       error
}

// Client interface to wrap the stackit dns api client, specifically for testing.
type Client interface {
	GetZones(
		ctx context.Context,
		projectId string,
		localVarOptionals *stackitdnsclient.ZoneApiV1ProjectsProjectIdZonesGetOpts,
	) (stackitdnsclient.ZoneResponseZoneAll, error)
	GetRRSets(
		ctx context.Context,
		projectId string,
		zoneId string,
		localVarOptionals *stackitdnsclient.RecordSetApiV1ProjectsProjectIdZonesZoneIdRrsetsGetOpts,
	) (stackitdnsclient.RrsetResponseRrSetAll, error)
	CreateRRSet(
		ctx context.Context,
		body stackitdnsclient.RrsetRrSetPost,
		projectId string,
		zoneId string,
	) (stackitdnsclient.RrsetResponseRrSet, error)
	UpdateRRSet(
		ctx context.Context,
		body stackitdnsclient.RrsetRrSetPatch,
		projectId string,
		zoneId string,
		rrsetId string,
	) (stackitdnsclient.RrsetResponseRrSet, error)
	DeleteRRSet(
		ctx context.Context,
		projectId string,
		zoneId string,
		rrsetId string,
	) (stackitdnsclient.SerializerMessage, error)
}

var _ Client = client{}

// client to implement Client.
type client struct {
	client *stackitdnsclient.APIClient
}

// GetRRSets returns all record sets for a given zone.
func (c client) GetRRSets(
	ctx context.Context,
	projectId string,
	zoneId string,
	localVarOptionals *stackitdnsclient.RecordSetApiV1ProjectsProjectIdZonesZoneIdRrsetsGetOpts,
) (stackitdnsclient.RrsetResponseRrSetAll, error) {
	res, _, err := c.client.RecordSetApi.V1ProjectsProjectIdZonesZoneIdRrsetsGet(
		ctx,
		projectId,
		zoneId,
		localVarOptionals,
	)

	return res, err
}

// CreateRRSet creates a new record set.
func (c client) CreateRRSet(
	ctx context.Context,
	body stackitdnsclient.RrsetRrSetPost,
	projectId string,
	zoneId string,
) (stackitdnsclient.RrsetResponseRrSet, error) {
	res, _, err := c.client.RecordSetApi.V1ProjectsProjectIdZonesZoneIdRrsetsPost(
		ctx,
		body,
		projectId,
		zoneId,
	)

	return res, err
}

// UpdateRRSet updates an existing record set.
func (c client) UpdateRRSet(
	ctx context.Context,
	body stackitdnsclient.RrsetRrSetPatch,
	projectId string,
	zoneId string,
	rrsetId string,
) (stackitdnsclient.RrsetResponseRrSet, error) {
	res, _, err := c.client.RecordSetApi.V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdPatch(
		ctx,
		body,
		projectId,
		zoneId,
		rrsetId,
	)

	return res, err
}

// DeleteRRSet deletes an existing record set.
func (c client) DeleteRRSet(
	ctx context.Context,
	projectId string,
	zoneId string,
	rrsetId string,
) (stackitdnsclient.SerializerMessage, error) {
	res, _, err := c.client.RecordSetApi.V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdDelete(
		ctx,
		projectId,
		zoneId,
		rrsetId,
	)

	return res, err
}

// GetZones returns all zones for a given project.
func (c client) GetZones(
	ctx context.Context,
	projectId string,
	localVarOptionals *stackitdnsclient.ZoneApiV1ProjectsProjectIdZonesGetOpts,
) (stackitdnsclient.ZoneResponseZoneAll, error) {
	res, _, err := c.client.ZoneApi.V1ProjectsProjectIdZonesGet(ctx, projectId, localVarOptionals)

	return res, err
}

// NewStackitDNSProvider creates a new STACKIT DNS provider.
func NewStackitDNSProvider(
	domainFilter endpoint.DomainFilter,
	dryRun bool,
	config Config,
) (*StackitDNSProvider, error) {
	configClient := stackitdnsclient.NewConfiguration()

	token := config.Token
	if token == "" {
		token = os.Getenv("EXTERNAL_DNS_STACKIT_CLIENT_TOKEN")
	}
	if token == "" {
		return nil, fmt.Errorf("no token found")
	}

	configClient.DefaultHeader["Authorization"] = fmt.Sprintf("Bearer %s", token)
	configClient.BasePath = config.BasePath
	apiClient := stackitdnsclient.NewAPIClient(configClient)
	clientWrapper := client{client: apiClient}

	provider := &StackitDNSProvider{
		Client:       clientWrapper,
		DomainFilter: domainFilter,
		DryRun:       dryRun,
		ProjectId:    config.ProjectId,
		Workers:      10,
	}

	return provider, nil
}

// Records returns resource records
func (d *StackitDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := d.zones(ctx)
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint
	endpointsErrorChannel := make(chan endpointError, len(zones))
	zoneIdsChannel := make(chan string, len(zones))

	for i := 0; i < d.Workers; i++ {
		go d.fetchRecordsWorker(ctx, zoneIdsChannel, endpointsErrorChannel)
	}

	for _, zone := range zones {
		zoneIdsChannel <- zone.Id
	}

	for i := 0; i < len(zones); i++ {
		endpointsErrorList := <-endpointsErrorChannel
		if endpointsErrorList.err != nil {
			close(zoneIdsChannel)

			return nil, endpointsErrorList.err
		}
		endpoints = append(endpoints, endpointsErrorList.endpoints...)
	}

	close(zoneIdsChannel)

	return endpoints, nil
}

// fetchRecordsWorker fetches all records from a given zone.
func (d *StackitDNSProvider) fetchRecordsWorker(
	ctx context.Context,
	zoneIdChannel chan string,
	endpointsErrorChannel chan<- endpointError,
) {
	for zoneId := range zoneIdChannel {
		var endpoints []*endpoint.Endpoint
		rrSets, err := d.fetchRecords(ctx, zoneId, nil)
		if err != nil {
			endpointsErrorChannel <- endpointError{
				endpoints: nil,
				err:       err,
			}

			continue
		}

		for _, r := range rrSets {
			if provider.SupportedRecordType(r.Type_) {
				for _, _r := range r.Records {
					endpoints = append(
						endpoints,
						endpoint.NewEndpointWithTTL(
							r.Name,
							r.Type_,
							endpoint.TTL(r.Ttl),
							_r.Content,
						),
					)
				}
			}
		}

		endpointsErrorChannel <- endpointError{
			endpoints: endpoints,
			err:       nil,
		}
	}

	log.Debug("fetch record set worker finished")
}

// ApplyChanges applies a given set of changes in a given zone.
func (d *StackitDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// create rr set. POST /v1/projects/{projectId}/zones/{zoneId}/rrsets
	err := d.createRRSets(ctx, changes.Create)
	if err != nil {
		return err
	}

	// update rr set. PATCH /v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId}
	err = d.updateRRSets(ctx, changes.UpdateNew)
	if err != nil {
		return err
	}

	// delete rr set. DELETE /v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId}
	err = d.deleteRRSets(ctx, changes.Delete)
	if err != nil {
		return err
	}

	return nil
}

// zones returns filtered list of stackitdnsclient.DomainZone if filter is set
func (d *StackitDNSProvider) zones(ctx context.Context) ([]stackitdnsclient.DomainZone, error) {
	if len(d.DomainFilter.Filters) == 0 {
		// no filters, return all zones
		queryParams := stackitdnsclient.ZoneApiV1ProjectsProjectIdZonesGetOpts{
			ActiveEq: optional.NewBool(true),
		}
		zones, err := d.fetchZones(ctx, queryParams)
		if err != nil {
			return nil, err
		}
		return zones, nil
	}

	var result []stackitdnsclient.DomainZone
	// send one request per filter
	for _, filter := range d.DomainFilter.Filters {
		queryParams := stackitdnsclient.ZoneApiV1ProjectsProjectIdZonesGetOpts{
			DnsNameLike: optional.NewString(filter),
			ActiveEq:    optional.NewBool(true),
		}
		zones, err := d.fetchZones(ctx, queryParams)
		if err != nil {
			return nil, err
		}
		result = append(result, zones...)
	}

	return result, nil
}

// fetchZones fetches all []stackitdnsclient.DomainZone from STACKIT DNS API.
func (d *StackitDNSProvider) fetchZones(
	ctx context.Context,
	queryParams stackitdnsclient.ZoneApiV1ProjectsProjectIdZonesGetOpts,
) ([]stackitdnsclient.DomainZone, error) {
	var result []stackitdnsclient.DomainZone
	queryParams.Page = optional.NewInt32(1)
	queryParams.PageSize = optional.NewInt32(10000)

	zoneResponse, err := d.Client.GetZones(
		ctx,
		d.ProjectId,
		&queryParams,
	)
	if err != nil {
		return nil, err
	}
	result = append(result, zoneResponse.Zones...)

	page := int32(2)
	for page <= zoneResponse.TotalPages {
		zoneResponse, err := d.Client.GetZones(
			ctx,
			d.ProjectId,
			&queryParams,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, zoneResponse.Zones...)
		page++
	}

	return result, nil
}

// fetchRecords fetches all []stackitdnsclient.DomainRrSet from STACKIT DNS API for given zone id.
func (d *StackitDNSProvider) fetchRecords(
	ctx context.Context,
	zoneId string,
	nameFilter *string,
) ([]stackitdnsclient.DomainRrSet, error) {
	var result []stackitdnsclient.DomainRrSet
	queryParams := stackitdnsclient.RecordSetApiV1ProjectsProjectIdZonesZoneIdRrsetsGetOpts{
		Page:     optional.NewInt32(1),
		PageSize: optional.NewInt32(10000),
		ActiveEq: optional.NewBool(true),
	}

	if nameFilter != nil {
		queryParams.NameLike = optional.NewString(*nameFilter)
	}

	rrSetResponse, err := d.Client.GetRRSets(
		ctx,
		d.ProjectId,
		zoneId,
		&queryParams,
	)
	if err != nil {
		return nil, err
	}
	result = append(result, rrSetResponse.RrSets...)

	queryParams.Page = optional.NewInt32(2)
	for queryParams.Page.Value() <= rrSetResponse.TotalPages {
		rrSetResponse, err := d.Client.GetRRSets(
			ctx,
			d.ProjectId,
			zoneId,
			&queryParams,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, rrSetResponse.RrSets...)
		queryParams.Page = optional.NewInt32(queryParams.Page.Value() + 1)
	}

	return result, nil
}

// handleRRSetWithWorkers handles the given endpoints with workers to optimize speed.
func (d *StackitDNSProvider) handleRRSetWithWorkers(
	ctx context.Context,
	endpoints []*endpoint.Endpoint,
	zones []stackitdnsclient.DomainZone,
	action string,
) {
	workerChannel := make(chan changeTask, len(endpoints))
	wg := new(sync.WaitGroup)

	// create workers
	for i := 0; i < d.Workers; i++ {
		wg.Add(1)
		go d.changeWorker(ctx, wg, workerChannel, zones)
	}

	for _, change := range endpoints {
		workerChannel <- changeTask{
			action: action,
			change: change,
		}
	}

	close(workerChannel)
	wg.Wait()
}

// changeWorker is a worker that handles changes passed by a channel.
func (d *StackitDNSProvider) changeWorker(
	ctx context.Context,
	wg *sync.WaitGroup,
	changes chan changeTask,
	zones []stackitdnsclient.DomainZone,
) {
	defer wg.Done()

	for change := range changes {
		switch change.action {
		case CREATE:
			_ = d.createRRSet(ctx, change.change, zones)
		case UPDATE:
			_ = d.updateRRSet(ctx, change.change, zones)
		case DELETE:
			_ = d.deleteRRSet(ctx, change.change, zones)
		}
	}

	log.Debug("change worker finished")
}

// createRRSets creates new record sets in the provider for the given endpoints that are in the
// creation field.
func (d *StackitDNSProvider) createRRSets(
	ctx context.Context,
	endpoints []*endpoint.Endpoint,
) error {
	if len(endpoints) == 0 {
		return nil
	}

	zones, err := d.zones(ctx)
	if err != nil {
		return err
	}

	d.handleRRSetWithWorkers(ctx, endpoints, zones, CREATE)

	return nil
}

// createRRSet creates a new record set in the provider for the given endpoint.
func (d *StackitDNSProvider) createRRSet(
	ctx context.Context,
	change *endpoint.Endpoint,
	zones []stackitdnsclient.DomainZone,
) error {
	resultZone, found := findBestMatchingZone(change.DNSName, zones)
	if !found {
		return fmt.Errorf("no matching zone found for %s", change.DNSName)
	}

	logFields := getLogFields(change, CREATE, resultZone.Id)
	log.WithFields(logFields).Info("create record set")

	if d.DryRun {
		log.WithFields(logFields).Debug("dry run, skipping")

		return nil
	}

	modifyChange(change)

	rrSet := getStackitRRSetRecordPost(change)

	// ignore all errors to just retry on next run
	_, err := d.Client.CreateRRSet(
		ctx,
		rrSet,
		d.ProjectId,
		resultZone.Id,
	)
	if err != nil {
		message := getSwaggerErrorMessage(err)
		log.Error(fmt.Sprintf("error creating record set: %s", message))

		return err
	}

	log.WithFields(logFields).Info("create record set successfully")

	return nil
}

// updateRRSets patches (overrides) contents in the record sets in the provider for the given
// endpoints that are in the update new field.
func (d *StackitDNSProvider) updateRRSets(
	ctx context.Context,
	endpoints []*endpoint.Endpoint,
) error {
	if len(endpoints) == 0 {
		return nil
	}

	zones, err := d.zones(ctx)
	if err != nil {
		return err
	}

	d.handleRRSetWithWorkers(ctx, endpoints, zones, UPDATE)

	return nil
}

// updateRRSet patches (overrides) contents in the record set in the provider.
func (d *StackitDNSProvider) updateRRSet(
	ctx context.Context,
	change *endpoint.Endpoint,
	zones []stackitdnsclient.DomainZone,
) error {
	modifyChange(change)

	resultZone, resultRRSet, err := d.getRRSetForUpdateDeletion(ctx, change, zones)
	if err != nil {
		return fmt.Errorf("no matching zone found for %s", change.DNSName)
	}

	logFields := getLogFields(change, UPDATE, resultRRSet.Id)
	log.WithFields(logFields).Info("update record set")

	if d.DryRun {
		log.WithFields(logFields).Debug("dry run, skipping")

		return nil
	}

	rrSet := getStackitRRSetRecordPatch(change)

	_, err = d.Client.UpdateRRSet(
		ctx,
		rrSet,
		d.ProjectId,
		resultZone.Id,
		resultRRSet.Id,
	)
	if err != nil {
		message := getSwaggerErrorMessage(err)
		log.Error(fmt.Sprintf("error updating record set: %s", message))

		return err
	}

	log.WithFields(logFields).Info("update record set successfully")

	return nil
}

// deleteRRSets delete record sets in the provider for the given endpoints that are in the
// deletion field.
func (d *StackitDNSProvider) deleteRRSets(
	ctx context.Context,
	endpoints []*endpoint.Endpoint,
) error {
	if len(endpoints) == 0 {
		log.Debug("no endpoints to delete")

		return nil
	}

	log.Info(fmt.Sprintf("records to delete %v", endpoints))

	zones, err := d.zones(ctx)
	if err != nil {
		return err
	}

	d.handleRRSetWithWorkers(ctx, endpoints, zones, DELETE)

	return nil
}

// deleteRRSet deletes a record set in the provider for the given endpoint.
func (d *StackitDNSProvider) deleteRRSet(
	ctx context.Context,
	change *endpoint.Endpoint,
	zones []stackitdnsclient.DomainZone,
) error {
	modifyChange(change)

	resultZone, resultRRSet, err := d.getRRSetForUpdateDeletion(ctx, change, zones)
	if err != nil {
		return fmt.Errorf("no matching zone found for %s", change.DNSName)
	}

	logFields := getLogFields(change, DELETE, resultRRSet.Id)
	log.WithFields(logFields).Info("delete record set")

	if d.DryRun {
		log.WithFields(logFields).Debug("dry run, skipping")

		return nil
	}

	_, err = d.Client.DeleteRRSet(
		ctx,
		d.ProjectId,
		resultZone.Id,
		resultRRSet.Id,
	)
	if err != nil {
		message := getSwaggerErrorMessage(err)
		log.Error(fmt.Sprintf("error delete record set: %s", message))

		return err
	}

	log.WithFields(logFields).Info("delete record set successfully")

	return nil
}

// findBestMatchingZone finds the best matching zone for a given record set name. The criteria are
// that the zone name is contained in the record set name and that the zone name is the longest
// possible match. Eg foo.bar.com. would have prejudice over bar.com. if rr set name is foo.bar.com.
func findBestMatchingZone(
	rrSetName string,
	zones []stackitdnsclient.DomainZone,
) (*stackitdnsclient.DomainZone, bool) {
	count := 0
	var domainZone stackitdnsclient.DomainZone
	for _, zone := range zones {
		if len(zone.DnsName) > count && strings.Contains(rrSetName, zone.DnsName) {
			count = len(zone.DnsName)
			domainZone = zone
		}
	}

	if count == 0 {
		return nil, false
	}

	return &domainZone, true
}

// findRRSet finds a record set by name and type in a list of record sets.
func findRRSet(
	rrSetName, recordType string,
	rrSets []stackitdnsclient.DomainRrSet,
) (*stackitdnsclient.DomainRrSet, bool) {
	for _, rrSet := range rrSets {
		if rrSet.Name == rrSetName && rrSet.Type_ == recordType {
			return &rrSet, true
		}
	}

	return nil, false
}

// appendDotIfNotExists appends a dot to the end of a string if it doesn't already end with a dot.
func appendDotIfNotExists(s string) string {
	if !strings.HasSuffix(s, ".") {
		return s + "."
	}

	return s
}

// modifyChange modifies a change to ensure it is valid for this provider.
func modifyChange(change *endpoint.Endpoint) {
	change.DNSName = appendDotIfNotExists(change.DNSName)

	if change.RecordTTL == 0 {
		change.RecordTTL = 300
	}
}

// getRRSetForUpdateDeletion returns the record set to be deleted and the zone it belongs to.
func (d *StackitDNSProvider) getRRSetForUpdateDeletion(
	ctx context.Context,
	change *endpoint.Endpoint,
	zones []stackitdnsclient.DomainZone,
) (*stackitdnsclient.DomainZone, *stackitdnsclient.DomainRrSet, error) {
	resultZone, found := findBestMatchingZone(change.DNSName, zones)
	if !found {
		log.Info(fmt.Sprintf("record set name %s contains no zone dns name", change.DNSName))

		return nil, nil, fmt.Errorf("record set name contains no zone dns name")
	}

	domainRrSets, err := d.fetchRecords(ctx, resultZone.Id, &change.DNSName)
	if err != nil {
		return nil, nil, err
	}

	resultRRSet, found := findRRSet(change.DNSName, change.RecordType, domainRrSets)
	if !found {
		log.Info(fmt.Sprintf("record %s not found on record sets", change.DNSName))

		return nil, nil, fmt.Errorf("record not found on record sets")
	}

	return resultZone, resultRRSet, nil
}

// getStackitRRSetRecordPost returns a stackitdnsclient.RrsetRrSetPost from a change for the api client.
func getStackitRRSetRecordPost(change *endpoint.Endpoint) stackitdnsclient.RrsetRrSetPost {
	records := make([]stackitdnsclient.RrsetRecordPost, len(change.Targets))
	for i, target := range change.Targets {
		records[i] = stackitdnsclient.RrsetRecordPost{
			Content: target,
		}
	}

	return stackitdnsclient.RrsetRrSetPost{
		Name:    change.DNSName,
		Records: records,
		Ttl:     int32(change.RecordTTL),
		Type_:   change.RecordType,
	}
}

// getStackitRRSetRecordPatch returns a stackitdnsclient.RrsetRrSetPatch from a change for the api client.
func getStackitRRSetRecordPatch(change *endpoint.Endpoint) stackitdnsclient.RrsetRrSetPatch {
	records := make([]stackitdnsclient.RrsetRecordPost, len(change.Targets))
	for i, target := range change.Targets {
		records[i] = stackitdnsclient.RrsetRecordPost{
			Content: target,
		}
	}

	return stackitdnsclient.RrsetRrSetPatch{
		Name:    change.DNSName,
		Records: records,
		Ttl:     int32(change.RecordTTL),
	}
}

// getLogFields returns a log.Fields object for a change.
func getLogFields(change *endpoint.Endpoint, action string, id string) log.Fields {
	return log.Fields{
		"record":  change.DNSName,
		"content": strings.Join(change.Targets, ","),
		"type":    change.RecordType,
		"action":  action,
		"id":      id,
	}
}

// getSwaggerErrorMessage returns the error message from a swagger error.
func getSwaggerErrorMessage(err error) string {
	message := err.Error()
	if v, ok := err.(stackitdnsclient.GenericSwaggerError); ok {
		if v2, ok := v.Model().(stackitdnsclient.SerializerMessage); ok {
			message = v2.Message
		}
	}

	return message
}
