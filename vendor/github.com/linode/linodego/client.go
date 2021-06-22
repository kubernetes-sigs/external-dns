package linodego

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
<<<<<<< HEAD
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	// APIConfigEnvVar environment var to get path to Linode config
	APIConfigEnvVar = "LINODE_CONFIG"
	// APIConfigProfileEnvVar specifies the profile to use when loading from a Linode config
	APIConfigProfileEnvVar = "LINODE_PROFILE"
	// APIHost Linode API hostname
	APIHost = "api.linode.com"
	// APIHostVar environment var to check for alternate API URL
	APIHostVar = "LINODE_URL"
	// APIHostCert environment var containing path to CA cert to validate against
	APIHostCert = "LINODE_CA"
	// APIVersion Linode API version
	APIVersion = "v4"
	// APIVersionVar environment var to check for alternate API Version
	APIVersionVar = "LINODE_API_VERSION"
	// APIProto connect to API with http(s)
	APIProto = "https"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	// APIEnvVar environment var to check for API token
	APIEnvVar = "LINODE_TOKEN"
	// APISecondsPerPoll how frequently to poll for new Events or Status in WaitFor functions
	APISecondsPerPoll = 3
	// Maximum wait time for retries
	APIRetryMaxWaitTime = time.Duration(30) * time.Second
)

var envDebug = false

// Client is a wrapper around the Resty client
type Client struct {
	resty             *resty.Client
	userAgent         string
	resources         map[string]*Resource
	debug             bool
	retryConditionals []RetryConditional

	millisecondsPerPoll time.Duration

	baseURL         string
	apiVersion      string
	apiProto        string
	selectedProfile string
	loadedProfile   string

	configProfiles map[string]ConfigProfile

	Account                *Resource
	AccountSettings        *Resource
	Databases              *Resource
	DomainRecords          *Resource
	Domains                *Resource
	Events                 *Resource
	Firewalls              *Resource
	FirewallDevices        *Resource
	FirewallRules          *Resource
	IPAddresses            *Resource
	IPv6Pools              *Resource
	IPv6Ranges             *Resource
	Images                 *Resource
	InstanceConfigs        *Resource
	InstanceDisks          *Resource
	InstanceIPs            *Resource
	InstanceSnapshots      *Resource
	InstanceStats          *Resource
	InstanceVolumes        *Resource
	Instances              *Resource
	InvoiceItems           *Resource
	Invoices               *Resource
	Kernels                *Resource
	LKEClusters            *Resource
	LKEClusterAPIEndpoints *Resource

	// Deprecated: Please use LKENodePools
	LKEClusterPools *Resource

	LKENodePools              *Resource
	LKEVersions               *Resource
	Longview                  *Resource
	LongviewClients           *Resource
	LongviewSubscriptions     *Resource
	Managed                   *Resource
	DatabaseMySQLInstances    *Resource
	DatabaseMongoInstances    *Resource
	DatabasePostgresInstances *Resource
	NodeBalancerConfigs       *Resource
	NodeBalancerNodes         *Resource
	NodeBalancerStats         *Resource
	NodeBalancers             *Resource
	Notifications             *Resource
	OAuthClients              *Resource
	ObjectStorageBuckets      *Resource
	ObjectStorageBucketCerts  *Resource
	ObjectStorageClusters     *Resource
	ObjectStorageKeys         *Resource
	ObjectStorage             *Resource
	Payments                  *Resource
	Profile                   *Resource
	ProfilePhoneNumber        *Resource
	ProfileSecurityQuestions  *Resource
	Regions                   *Resource
	SSHKeys                   *Resource
	StackScripts              *Resource
	Tags                      *Resource
	Tickets                   *Resource
	Token                     *Resource
	Tokens                    *Resource
	Types                     *Resource
	UserGrants                *Resource
	Users                     *Resource
	VLANs                     *Resource
	Volumes                   *Resource
}

type EnvDefaults struct {
	Token   string
	Profile string
}

type Request = resty.Request

func init() {
	// Wether or not we will enable Resty debugging output
	if apiDebug, ok := os.LookupEnv("LINODE_DEBUG"); ok {
		if parsed, err := strconv.ParseBool(apiDebug); err == nil {
			envDebug = parsed
			log.Println("[INFO] LINODE_DEBUG being set to", envDebug)
		} else {
			log.Println("[WARN] LINODE_DEBUG should be an integer, 0 or 1")
		}
	}
}

// SetUserAgent sets a custom user-agent for HTTP requests
func (c *Client) SetUserAgent(ua string) *Client {
	c.userAgent = ua
	c.resty.SetHeader("User-Agent", c.userAgent)

	return c
}

// R wraps resty's R method
func (c *Client) R(ctx context.Context) *resty.Request {
	return c.resty.R().
		ExpectContentType("application/json").
		SetHeader("Content-Type", "application/json").
		SetContext(ctx).
		SetError(APIError{})
}

// SetDebug sets the debug on resty's client
func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	c.resty.SetDebug(debug)

	return c
}

// OnBeforeRequest adds a handler to the request body to run before the request is sent
func (c *Client) OnBeforeRequest(m func(request *Request) error) {
	c.resty.OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
		return m(req)
	})
}

// SetBaseURL sets the base URL of the Linode v4 API (https://api.linode.com/v4)
func (c *Client) SetBaseURL(baseURL string) *Client {
	baseURLPath, _ := url.Parse(baseURL)

	c.baseURL = path.Join(baseURLPath.Host, baseURLPath.Path)
	c.apiProto = baseURLPath.Scheme

	c.updateHostURL()

	return c
}

// SetAPIVersion sets the version of the API to interface with
func (c *Client) SetAPIVersion(apiVersion string) *Client {
	c.apiVersion = apiVersion

	c.updateHostURL()

	return c
}

func (c *Client) updateHostURL() {
	apiProto := APIProto
	baseURL := APIHost
	apiVersion := APIVersion

	if c.baseURL != "" {
		baseURL = c.baseURL
	}

	if c.apiVersion != "" {
		apiVersion = c.apiVersion
	}

	if c.apiProto != "" {
		apiProto = c.apiProto
	}

	c.resty.SetHostURL(fmt.Sprintf("%s://%s/%s", apiProto, baseURL, apiVersion))
}

// SetRootCertificate adds a root certificate to the underlying TLS client config
func (c *Client) SetRootCertificate(path string) *Client {
	c.resty.SetRootCertificate(path)
	return c
}

// SetToken sets the API token for all requests from this client
// Only necessary if you haven't already provided an http client to NewClient() configured with the token.
func (c *Client) SetToken(token string) *Client {
	c.resty.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	return c
}

// SetRetries adds retry conditions for "Linode Busy." errors and 429s.
func (c *Client) SetRetries() *Client {
	c.
		addRetryConditional(linodeBusyRetryCondition).
		addRetryConditional(tooManyRequestsRetryCondition).
		addRetryConditional(serviceUnavailableRetryCondition).
		addRetryConditional(requestTimeoutRetryCondition).
		SetRetryMaxWaitTime(APIRetryMaxWaitTime)
	configureRetries(c)
	return c
}

// AddRetryCondition adds a RetryConditional function to the Client
func (c *Client) AddRetryCondition(retryCondition RetryConditional) *Client {
	c.resty.AddRetryCondition(resty.RetryConditionFunc(retryCondition))
	return c
}

func (c *Client) addRetryConditional(retryConditional RetryConditional) *Client {
	c.retryConditionals = append(c.retryConditionals, retryConditional)
	return c
}

// SetRetryMaxWaitTime sets the maximum delay before retrying a request.
func (c *Client) SetRetryMaxWaitTime(max time.Duration) *Client {
	c.resty.SetRetryMaxWaitTime(max)
	return c
}

// SetRetryWaitTime sets the default (minimum) delay before retrying a request.
func (c *Client) SetRetryWaitTime(min time.Duration) *Client {
	c.resty.SetRetryWaitTime(min)
	return c
}

// SetRetryAfter sets the callback function to be invoked with a failed request
// to determine wben it should be retried.
func (c *Client) SetRetryAfter(callback RetryAfter) *Client {
	c.resty.SetRetryAfter(resty.RetryAfterFunc(callback))
	return c
}

// SetRetryCount sets the maximum retry attempts before aborting.
func (c *Client) SetRetryCount(count int) *Client {
	c.resty.SetRetryCount(count)
	return c
}

// SetPollDelay sets the number of milliseconds to wait between events or status polls.
// Affects all WaitFor* functions and retries.
func (c *Client) SetPollDelay(delay time.Duration) *Client {
	c.millisecondsPerPoll = delay
	return c
}

// GetPollDelay gets the number of milliseconds to wait between events or status polls.
// Affects all WaitFor* functions and retries.
func (c *Client) GetPollDelay() time.Duration {
	return c.millisecondsPerPoll
}

// Resource looks up a resource by name
func (c Client) Resource(resourceName string) *Resource {
	selectedResource, ok := c.resources[resourceName]
	if !ok {
		log.Fatalf("Could not find resource named '%s', exiting.", resourceName)
	}

	return selectedResource
}

// NewClient factory to create new Client struct
func NewClient(hc *http.Client) (client Client) {
	if hc != nil {
		client.resty = resty.NewWithClient(hc)
	} else {
		client.resty = resty.New()
	}

	client.SetUserAgent(DefaultUserAgent)

	baseURL, baseURLExists := os.LookupEnv(APIHostVar)

	if baseURLExists {
		client.SetBaseURL(baseURL)
	}
	apiVersion, apiVersionExists := os.LookupEnv(APIVersionVar)
	if apiVersionExists {
		client.SetAPIVersion(apiVersion)
	} else {
		client.SetAPIVersion(APIVersion)
	}

	certPath, certPathExists := os.LookupEnv(APIHostCert)

	if certPathExists {
		cert, err := ioutil.ReadFile(certPath)
		if err != nil {
			log.Fatalf("[ERROR] Error when reading cert at %s: %s\n", certPath, err.Error())
		}

		client.SetRootCertificate(certPath)

		if envDebug {
			log.Printf("[DEBUG] Set API root certificate to %s with contents %s\n", certPath, cert)
		}
	}

	client.
		SetRetryWaitTime((1000 * APISecondsPerPoll) * time.Millisecond).
		SetPollDelay(1000 * APISecondsPerPoll).
		SetRetries().
		SetDebug(envDebug)

	addResources(&client)

	return
}

// NewClientFromEnv creates a Client and initializes it with values
// from the LINODE_CONFIG file and the LINODE_TOKEN environment variable.
func NewClientFromEnv(hc *http.Client) (*Client, error) {
	client := NewClient(hc)

	// Users are expected to chain NewClient(...) and LoadConfig(...) to customize these options
	configPath, err := resolveValidConfigPath()
	if err != nil {
		return nil, err
	}

	// Populate the token from the environment.
	// Tokens should be first priority to maintain backwards compatibility
	if token, ok := os.LookupEnv(APIEnvVar); ok && token != "" {
		client.SetToken(token)
		return &client, nil
	}

	if p, ok := os.LookupEnv(APIConfigEnvVar); ok {
		configPath = p
	} else if !ok && configPath == "" {
		return nil, fmt.Errorf("no linode config file or token found")
	}

	configProfile := DefaultConfigProfile

	if p, ok := os.LookupEnv(APIConfigProfileEnvVar); ok {
		configProfile = p
	}

	client.selectedProfile = configProfile

	// We should only load the config if the config file exists
	if _, err := os.Stat(configPath); err != nil {
		return nil, fmt.Errorf("error loading config file %s: %s", configPath, err)
	}

	err = client.preLoadConfig(configPath)
	return &client, err
}

func (c *Client) preLoadConfig(configPath string) error {
	if envDebug {
		log.Printf("[INFO] Loading profile from %s\n", configPath)
	}

	if err := c.LoadConfig(&LoadConfigOptions{
		Path:            configPath,
		SkipLoadProfile: true,
	}); err != nil {
		return err
	}

	// We don't want to load the profile until the user is actually making requests
	c.OnBeforeRequest(func(request *Request) error {
		if c.loadedProfile != c.selectedProfile {
			if err := c.UseProfile(c.selectedProfile); err != nil {
				return err
			}
		}

		return nil
	})

	return nil
}

// nolint
func addResources(client *Client) {
	resources := map[string]*Resource{
		accountName:                  NewResource(client, accountName, accountEndpoint, false, Account{}, nil),                         // really?
		accountSettingsName:          NewResource(client, accountSettingsName, accountSettingsEndpoint, false, AccountSettings{}, nil), // really?
		databasesName:                NewResource(client, databasesName, databasesEndpoint, false, Database{}, nil),
		domainRecordsName:            NewResource(client, domainRecordsName, domainRecordsEndpoint, true, DomainRecord{}, DomainRecordsPagedResponse{}),
		domainsName:                  NewResource(client, domainsName, domainsEndpoint, false, Domain{}, DomainsPagedResponse{}),
		eventsName:                   NewResource(client, eventsName, eventsEndpoint, false, Event{}, EventsPagedResponse{}),
		firewallsName:                NewResource(client, firewallsName, firewallsEndpoint, false, Firewall{}, FirewallsPagedResponse{}),
		firewallDevicesName:          NewResource(client, firewallDevicesName, firewallDevicesEndpoint, true, FirewallDevice{}, FirewallDevicesPagedResponse{}),
		firewallRulesName:            NewResource(client, firewallRulesName, firewallRulesEndpoint, true, FirewallRule{}, nil),
		imagesName:                   NewResource(client, imagesName, imagesEndpoint, false, Image{}, ImagesPagedResponse{}),
		instanceConfigsName:          NewResource(client, instanceConfigsName, instanceConfigsEndpoint, true, InstanceConfig{}, InstanceConfigsPagedResponse{}),
		instanceDisksName:            NewResource(client, instanceDisksName, instanceDisksEndpoint, true, InstanceDisk{}, InstanceDisksPagedResponse{}),
		instanceIPsName:              NewResource(client, instanceIPsName, instanceIPsEndpoint, true, InstanceIP{}, nil), // really?
		instanceSnapshotsName:        NewResource(client, instanceSnapshotsName, instanceSnapshotsEndpoint, true, InstanceSnapshot{}, nil),
		instanceStatsName:            NewResource(client, instanceStatsName, instanceStatsEndpoint, true, InstanceStats{}, nil),
		instanceVolumesName:          NewResource(client, instanceVolumesName, instanceVolumesEndpoint, true, nil, InstanceVolumesPagedResponse{}), // really?
		instancesName:                NewResource(client, instancesName, instancesEndpoint, false, Instance{}, InstancesPagedResponse{}),
		invoiceItemsName:             NewResource(client, invoiceItemsName, invoiceItemsEndpoint, true, InvoiceItem{}, InvoiceItemsPagedResponse{}),
		invoicesName:                 NewResource(client, invoicesName, invoicesEndpoint, false, Invoice{}, InvoicesPagedResponse{}),
		ipaddressesName:              NewResource(client, ipaddressesName, ipaddressesEndpoint, false, nil, IPAddressesPagedResponse{}), // really?
		ipv6poolsName:                NewResource(client, ipv6poolsName, ipv6poolsEndpoint, false, nil, IPv6PoolsPagedResponse{}),       // really?
		ipv6rangesName:               NewResource(client, ipv6rangesName, ipv6rangesEndpoint, false, IPv6Range{}, IPv6RangesPagedResponse{}),
		kernelsName:                  NewResource(client, kernelsName, kernelsEndpoint, false, LinodeKernel{}, LinodeKernelsPagedResponse{}),
		lkeClusterAPIEndpointsName:   NewResource(client, lkeClusterAPIEndpointsName, lkeClusterAPIEndpointsEndpoint, true, LKEClusterAPIEndpoint{}, LKEClusterAPIEndpointsPagedResponse{}),
		lkeClustersName:              NewResource(client, lkeClustersName, lkeClustersEndpoint, false, LKECluster{}, LKEClustersPagedResponse{}),
		lkeClusterPoolsName:          NewResource(client, lkeClusterPoolsName, lkeClusterPoolsEndpoint, true, LKEClusterPool{}, LKEClusterPoolsPagedResponse{}),
		lkeNodePoolsName:             NewResource(client, lkeNodePoolsName, lkeNodePoolsEndpoint, true, LKENodePool{}, LKENodePoolsPagedResponse{}),
		lkeVersionsName:              NewResource(client, lkeVersionsName, lkeVersionsEndpoint, false, LKEVersion{}, LKEVersionsPagedResponse{}),
		longviewName:                 NewResource(client, longviewName, longviewEndpoint, false, nil, nil), // really?
		longviewclientsName:          NewResource(client, longviewclientsName, longviewclientsEndpoint, false, LongviewClient{}, LongviewClientsPagedResponse{}),
		longviewsubscriptionsName:    NewResource(client, longviewsubscriptionsName, longviewsubscriptionsEndpoint, false, LongviewSubscription{}, LongviewSubscriptionsPagedResponse{}),
		managedName:                  NewResource(client, managedName, managedEndpoint, false, nil, nil), // really?
		mysqlName:                    NewResource(client, mysqlName, mysqlEndpoint, false, MySQLDatabase{}, MySQLDatabasesPagedResponse{}),
		mongoName:                    NewResource(client, mongoName, mongoEndpoint, false, MongoDatabase{}, MongoDatabasesPagedResponse{}),
		postgresName:                 NewResource(client, postgresName, postgresEndpoint, false, PostgresDatabase{}, PostgresDatabasesPagedResponse{}),
		nodebalancerconfigsName:      NewResource(client, nodebalancerconfigsName, nodebalancerconfigsEndpoint, true, NodeBalancerConfig{}, NodeBalancerConfigsPagedResponse{}),
		nodebalancernodesName:        NewResource(client, nodebalancernodesName, nodebalancernodesEndpoint, true, NodeBalancerNode{}, NodeBalancerNodesPagedResponse{}),
		nodebalancerStatsName:        NewResource(client, nodebalancerStatsName, nodebalancerStatsEndpoint, true, NodeBalancerStats{}, nil),
		nodebalancersName:            NewResource(client, nodebalancersName, nodebalancersEndpoint, false, NodeBalancer{}, NodeBalancerConfigsPagedResponse{}),
		notificationsName:            NewResource(client, notificationsName, notificationsEndpoint, false, Notification{}, NotificationsPagedResponse{}),
		oauthClientsName:             NewResource(client, oauthClientsName, oauthClientsEndpoint, false, OAuthClient{}, OAuthClientsPagedResponse{}),
		objectStorageBucketsName:     NewResource(client, objectStorageBucketsName, objectStorageBucketsEndpoint, false, ObjectStorageBucket{}, ObjectStorageBucketsPagedResponse{}),
		objectStorageBucketCertsName: NewResource(client, objectStorageBucketCertsName, objectStorageBucketCertsEndpoint, true, ObjectStorageBucketCert{}, nil),
		objectStorageClustersName:    NewResource(client, objectStorageClustersName, objectStorageClustersEndpoint, false, ObjectStorageCluster{}, ObjectStorageClustersPagedResponse{}),
		objectStorageKeysName:        NewResource(client, objectStorageKeysName, objectStorageKeysEndpoint, false, ObjectStorageKey{}, ObjectStorageKeysPagedResponse{}),
		objectStorageName:            NewResource(client, objectStorageName, objectStorageEndpoint, false, nil, nil),
		paymentsName:                 NewResource(client, paymentsName, paymentsEndpoint, false, Payment{}, PaymentsPagedResponse{}),
		profileName:                  NewResource(client, profileName, profileEndpoint, false, nil, nil), // really?
		profilePhoneNumberName:       NewResource(client, profilePhoneNumberName, profilePhoneNumberEndpoint, false, nil, nil),
		profileSecurityQuestionsName: NewResource(client, profileSecurityQuestionsName, profileSecurityQuestionsEndpoint, false, nil, nil),
		regionsName:                  NewResource(client, regionsName, regionsEndpoint, false, Region{}, RegionsPagedResponse{}),
		sshkeysName:                  NewResource(client, sshkeysName, sshkeysEndpoint, false, SSHKey{}, SSHKeysPagedResponse{}),
		stackscriptsName:             NewResource(client, stackscriptsName, stackscriptsEndpoint, false, Stackscript{}, StackscriptsPagedResponse{}),
		tagsName:                     NewResource(client, tagsName, tagsEndpoint, false, Tag{}, TagsPagedResponse{}),
		ticketsName:                  NewResource(client, ticketsName, ticketsEndpoint, false, Ticket{}, TicketsPagedResponse{}),
		tokensName:                   NewResource(client, tokensName, tokensEndpoint, false, Token{}, TokensPagedResponse{}),
		typesName:                    NewResource(client, typesName, typesEndpoint, false, LinodeType{}, LinodeTypesPagedResponse{}),
		userGrantsName:               NewResource(client, typesName, userGrantsEndpoint, true, UserGrants{}, nil),
		usersName:                    NewResource(client, usersName, usersEndpoint, false, User{}, UsersPagedResponse{}),
		vlansName:                    NewResource(client, vlansName, vlansEndpoint, false, VLAN{}, VLANsPagedResponse{}),
		volumesName:                  NewResource(client, volumesName, volumesEndpoint, false, Volume{}, VolumesPagedResponse{}),
	}

	client.resources = resources

	client.Account = resources[accountName]
	client.Databases = resources[databasesName]
	client.DomainRecords = resources[domainRecordsName]
	client.Domains = resources[domainsName]
	client.Events = resources[eventsName]
	client.Firewalls = resources[firewallsName]
	client.FirewallDevices = resources[firewallDevicesName]
	client.FirewallRules = resources[firewallRulesName]
	client.IPAddresses = resources[ipaddressesName]
	client.IPv6Pools = resources[ipv6poolsName]
	client.IPv6Ranges = resources[ipv6rangesName]
	client.Images = resources[imagesName]
	client.InstanceConfigs = resources[instanceConfigsName]
	client.InstanceDisks = resources[instanceDisksName]
	client.InstanceIPs = resources[instanceIPsName]
	client.InstanceSnapshots = resources[instanceSnapshotsName]
	client.InstanceStats = resources[instanceStatsName]
	client.InstanceVolumes = resources[instanceVolumesName]
	client.Instances = resources[instancesName]
	client.Invoices = resources[invoicesName]
	client.Kernels = resources[kernelsName]
	client.LKEClusterAPIEndpoints = resources[lkeClusterAPIEndpointsName]
	client.LKEClusters = resources[lkeClustersName]
	client.LKEClusterPools = resources[lkeClusterPoolsName]
	client.LKENodePools = resources[lkeNodePoolsName]
	client.LKEVersions = resources[lkeVersionsName]
	client.Longview = resources[longviewName]
	client.LongviewSubscriptions = resources[longviewsubscriptionsName]
	client.Managed = resources[managedName]
	client.DatabaseMySQLInstances = resources[mysqlName]
	client.DatabaseMongoInstances = resources[mongoName]
	client.DatabasePostgresInstances = resources[postgresName]
	client.NodeBalancerConfigs = resources[nodebalancerconfigsName]
	client.NodeBalancerNodes = resources[nodebalancernodesName]
	client.NodeBalancerStats = resources[nodebalancerStatsName]
	client.NodeBalancers = resources[nodebalancersName]
	client.Notifications = resources[notificationsName]
	client.OAuthClients = resources[oauthClientsName]
	client.ObjectStorageBuckets = resources[objectStorageBucketsName]
	client.ObjectStorageBucketCerts = resources[objectStorageBucketCertsName]
	client.ObjectStorageClusters = resources[objectStorageClustersName]
	client.ObjectStorageKeys = resources[objectStorageKeysName]
	client.ObjectStorage = resources[objectStorageName]
	client.Payments = resources[paymentsName]
	client.Profile = resources[profileName]
	client.ProfilePhoneNumber = resources[profilePhoneNumberName]
	client.ProfileSecurityQuestions = resources[profileSecurityQuestionsName]
	client.Regions = resources[regionsName]
	client.SSHKeys = resources[sshkeysName]
	client.StackScripts = resources[stackscriptsName]
	client.Tags = resources[tagsName]
	client.Tickets = resources[ticketsName]
	client.Tokens = resources[tokensName]
	client.Types = resources[typesName]
	client.UserGrants = resources[userGrantsName]
	client.Users = resources[usersName]
	client.VLANs = resources[vlansName]
	client.Volumes = resources[volumesName]
}

func copyBool(bPtr *bool) *bool {
	if bPtr == nil {
		return nil
	}

	t := *bPtr

	return &t
}

func copyInt(iPtr *int) *int {
	if iPtr == nil {
		return nil
	}

	t := *iPtr

	return &t
}

func copyString(sPtr *string) *string {
	if sPtr == nil {
		return nil
	}

	t := *sPtr

	return &t
}

func copyTime(tPtr *time.Time) *time.Time {
	if tPtr == nil {
		return nil
	}

	t := *tPtr
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Version of linodego
	Version = "0.12.0"
||||||| parent of 5ce8c7613 (update vendored files)
	// Version of linodego
	Version = "0.12.0"
=======
>>>>>>> 5ce8c7613 (update vendored files)
	// APIEnvVar environment var to check for API token
	APIEnvVar = "LINODE_TOKEN"
	// APISecondsPerPoll how frequently to poll for new Events or Status in WaitFor functions
	APISecondsPerPoll = 3
	// Maximum wait time for retries
	APIRetryMaxWaitTime = time.Duration(30) * time.Second
)

var envDebug = false

// Client is a wrapper around the Resty client
type Client struct {
	resty             *resty.Client
	userAgent         string
	resources         map[string]*Resource
	debug             bool
	retryConditionals []RetryConditional

	millisecondsPerPoll time.Duration

	Account                  *Resource
	AccountSettings          *Resource
	DomainRecords            *Resource
	Domains                  *Resource
	Events                   *Resource
	Firewalls                *Resource
	FirewallDevices          *Resource
	FirewallRules            *Resource
	IPAddresses              *Resource
	IPv6Pools                *Resource
	IPv6Ranges               *Resource
	Images                   *Resource
	InstanceConfigs          *Resource
	InstanceDisks            *Resource
	InstanceIPs              *Resource
	InstanceSnapshots        *Resource
	InstanceStats            *Resource
	InstanceVolumes          *Resource
	Instances                *Resource
	InvoiceItems             *Resource
	Invoices                 *Resource
	Kernels                  *Resource
	LKEClusters              *Resource
	LKEClusterAPIEndpoints   *Resource
	LKEClusterPools          *Resource
	LKEVersions              *Resource
	Longview                 *Resource
	LongviewClients          *Resource
	LongviewSubscriptions    *Resource
	Managed                  *Resource
	NodeBalancerConfigs      *Resource
	NodeBalancerNodes        *Resource
	NodeBalancerStats        *Resource
	NodeBalancers            *Resource
	Notifications            *Resource
	OAuthClients             *Resource
	ObjectStorageBuckets     *Resource
	ObjectStorageBucketCerts *Resource
	ObjectStorageClusters    *Resource
	ObjectStorageKeys        *Resource
	Payments                 *Resource
	Profile                  *Resource
	Regions                  *Resource
	SSHKeys                  *Resource
	StackScripts             *Resource
	Tags                     *Resource
	Tickets                  *Resource
	Token                    *Resource
	Tokens                   *Resource
	Types                    *Resource
	UserGrants               *Resource
	Users                    *Resource
	VLANs                    *Resource
	Volumes                  *Resource
}

func init() {
	// Wether or not we will enable Resty debugging output
	if apiDebug, ok := os.LookupEnv("LINODE_DEBUG"); ok {
		if parsed, err := strconv.ParseBool(apiDebug); err == nil {
			envDebug = parsed
			log.Println("[INFO] LINODE_DEBUG being set to", envDebug)
		} else {
			log.Println("[WARN] LINODE_DEBUG should be an integer, 0 or 1")
		}
	}
}

// SetUserAgent sets a custom user-agent for HTTP requests
func (c *Client) SetUserAgent(ua string) *Client {
	c.userAgent = ua
	c.resty.SetHeader("User-Agent", c.userAgent)

	return c
}

// R wraps resty's R method
func (c *Client) R(ctx context.Context) *resty.Request {
	return c.resty.R().
		ExpectContentType("application/json").
		SetHeader("Content-Type", "application/json").
		SetContext(ctx).
		SetError(APIError{})
}

// SetDebug sets the debug on resty's client
func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	c.resty.SetDebug(debug)

	return c
}

// SetBaseURL sets the base URL of the Linode v4 API (https://api.linode.com/v4)
func (c *Client) SetBaseURL(url string) *Client {
	c.resty.SetHostURL(url)
	return c
}

// SetAPIVersion sets the version of the API to interface with
func (c *Client) SetAPIVersion(apiVersion string) *Client {
	c.SetBaseURL(fmt.Sprintf("%s://%s/%s", APIProto, APIHost, apiVersion))
	return c
}

// SetRootCertificate adds a root certificate to the underlying TLS client config
func (c *Client) SetRootCertificate(path string) *Client {
	c.resty.SetRootCertificate(path)
	return c
}

// SetToken sets the API token for all requests from this client
// Only necessary if you haven't already provided an http client to NewClient() configured with the token.
func (c *Client) SetToken(token string) *Client {
	c.resty.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	return c
}

// SetRetries adds retry conditions for "Linode Busy." errors and 429s.
func (c *Client) SetRetries() *Client {
	c.
		addRetryConditional(linodeBusyRetryCondition).
		addRetryConditional(tooManyRequestsRetryCondition).
		addRetryConditional(serviceUnavailableRetryCondition).
		addRetryConditional(requestTimeoutRetryCondition).
		SetRetryMaxWaitTime(APIRetryMaxWaitTime)
	configureRetries(c)
	return c
}

func (c *Client) addRetryConditional(retryConditional RetryConditional) *Client {
	c.retryConditionals = append(c.retryConditionals, retryConditional)
	return c
}

// SetRetryMaxWaitTime sets the maximum delay before retrying a request.
func (c *Client) SetRetryMaxWaitTime(max time.Duration) *Client {
	c.resty.SetRetryMaxWaitTime(max)
	return c
}

// SetRetryWaitTime sets the default (minimum) delay before retrying a request.
func (c *Client) SetRetryWaitTime(min time.Duration) *Client {
	c.resty.SetRetryWaitTime(min)
	return c
}

// SetRetryAfter sets the callback function to be invoked with a failed request
// to determine wben it should be retried.
func (c *Client) SetRetryAfter(callback RetryAfter) *Client {
	c.resty.SetRetryAfter(resty.RetryAfterFunc(callback))
	return c
}

// SetRetryCount sets the maximum retry attempts before aborting.
func (c *Client) SetRetryCount(count int) *Client {
	c.resty.SetRetryCount(count)
	return c
}

// SetPollDelay sets the number of milliseconds to wait between events or status polls.
// Affects all WaitFor* functions and retries.
func (c *Client) SetPollDelay(delay time.Duration) *Client {
	c.millisecondsPerPoll = delay
	return c
}

// Resource looks up a resource by name
func (c Client) Resource(resourceName string) *Resource {
	selectedResource, ok := c.resources[resourceName]
	if !ok {
		log.Fatalf("Could not find resource named '%s', exiting.", resourceName)
	}

	return selectedResource
}

// NewClient factory to create new Client struct
func NewClient(hc *http.Client) (client Client) {
	if hc != nil {
		client.resty = resty.NewWithClient(hc)
	} else {
		client.resty = resty.New()
	}

	client.SetUserAgent(DefaultUserAgent)

	baseURL, baseURLExists := os.LookupEnv(APIHostVar)

	if baseURLExists {
		client.SetBaseURL(baseURL)
	} else {
		apiVersion, apiVersionExists := os.LookupEnv(APIVersionVar)
		if apiVersionExists {
			client.SetAPIVersion(apiVersion)
		} else {
			client.SetAPIVersion(APIVersion)
		}
	}

	certPath, certPathExists := os.LookupEnv(APIHostCert)

	if certPathExists {
		cert, err := ioutil.ReadFile(certPath)
		if err != nil {
			log.Fatalf("[ERROR] Error when reading cert at %s: %s\n", certPath, err.Error())
		}

		client.SetRootCertificate(certPath)

		if envDebug {
			log.Printf("[DEBUG] Set API root certificate to %s with contents %s\n", certPath, cert)
		}
	}

	client.
		SetRetryWaitTime((1000 * APISecondsPerPoll) * time.Millisecond).
		SetPollDelay(1000 * APISecondsPerPoll).
		SetRetries().
		SetDebug(envDebug)

	addResources(&client)

	return
}

// nolint
func addResources(client *Client) {
	resources := map[string]*Resource{
		accountName:                  NewResource(client, accountName, accountEndpoint, false, Account{}, nil),                         // really?
		accountSettingsName:          NewResource(client, accountSettingsName, accountSettingsEndpoint, false, AccountSettings{}, nil), // really?
		domainRecordsName:            NewResource(client, domainRecordsName, domainRecordsEndpoint, true, DomainRecord{}, DomainRecordsPagedResponse{}),
		domainsName:                  NewResource(client, domainsName, domainsEndpoint, false, Domain{}, DomainsPagedResponse{}),
		eventsName:                   NewResource(client, eventsName, eventsEndpoint, false, Event{}, EventsPagedResponse{}),
		firewallsName:                NewResource(client, firewallsName, firewallsEndpoint, false, Firewall{}, FirewallsPagedResponse{}),
		firewallDevicesName:          NewResource(client, firewallDevicesName, firewallDevicesEndpoint, true, FirewallDevice{}, FirewallDevicesPagedResponse{}),
		firewallRulesName:            NewResource(client, firewallRulesName, firewallRulesEndpoint, true, FirewallRule{}, nil),
		imagesName:                   NewResource(client, imagesName, imagesEndpoint, false, Image{}, ImagesPagedResponse{}),
		instanceConfigsName:          NewResource(client, instanceConfigsName, instanceConfigsEndpoint, true, InstanceConfig{}, InstanceConfigsPagedResponse{}),
		instanceDisksName:            NewResource(client, instanceDisksName, instanceDisksEndpoint, true, InstanceDisk{}, InstanceDisksPagedResponse{}),
		instanceIPsName:              NewResource(client, instanceIPsName, instanceIPsEndpoint, true, InstanceIP{}, nil), // really?
		instanceSnapshotsName:        NewResource(client, instanceSnapshotsName, instanceSnapshotsEndpoint, true, InstanceSnapshot{}, nil),
		instanceStatsName:            NewResource(client, instanceStatsName, instanceStatsEndpoint, true, InstanceStats{}, nil),
		instanceVolumesName:          NewResource(client, instanceVolumesName, instanceVolumesEndpoint, true, nil, InstanceVolumesPagedResponse{}), // really?
		instancesName:                NewResource(client, instancesName, instancesEndpoint, false, Instance{}, InstancesPagedResponse{}),
		invoiceItemsName:             NewResource(client, invoiceItemsName, invoiceItemsEndpoint, true, InvoiceItem{}, InvoiceItemsPagedResponse{}),
		invoicesName:                 NewResource(client, invoicesName, invoicesEndpoint, false, Invoice{}, InvoicesPagedResponse{}),
		ipaddressesName:              NewResource(client, ipaddressesName, ipaddressesEndpoint, false, nil, IPAddressesPagedResponse{}), // really?
		ipv6poolsName:                NewResource(client, ipv6poolsName, ipv6poolsEndpoint, false, nil, IPv6PoolsPagedResponse{}),       // really?
		ipv6rangesName:               NewResource(client, ipv6rangesName, ipv6rangesEndpoint, false, IPv6Range{}, IPv6RangesPagedResponse{}),
		kernelsName:                  NewResource(client, kernelsName, kernelsEndpoint, false, LinodeKernel{}, LinodeKernelsPagedResponse{}),
		lkeClusterAPIEndpointsName:   NewResource(client, lkeClusterAPIEndpointsName, lkeClusterAPIEndpointsEndpoint, true, LKEClusterAPIEndpoint{}, LKEClusterAPIEndpointsPagedResponse{}),
		lkeClustersName:              NewResource(client, lkeClustersName, lkeClustersEndpoint, false, LKECluster{}, LKEClustersPagedResponse{}),
		lkeClusterPoolsName:          NewResource(client, lkeClusterPoolsName, lkeClusterPoolsEndpoint, true, LKEClusterPool{}, LKEClusterPoolsPagedResponse{}),
		lkeVersionsName:              NewResource(client, lkeVersionsName, lkeVersionsEndpoint, false, LKEVersion{}, LKEVersionsPagedResponse{}),
		longviewName:                 NewResource(client, longviewName, longviewEndpoint, false, nil, nil), // really?
		longviewclientsName:          NewResource(client, longviewclientsName, longviewclientsEndpoint, false, LongviewClient{}, LongviewClientsPagedResponse{}),
		longviewsubscriptionsName:    NewResource(client, longviewsubscriptionsName, longviewsubscriptionsEndpoint, false, LongviewSubscription{}, LongviewSubscriptionsPagedResponse{}),
		managedName:                  NewResource(client, managedName, managedEndpoint, false, nil, nil), // really?
		nodebalancerconfigsName:      NewResource(client, nodebalancerconfigsName, nodebalancerconfigsEndpoint, true, NodeBalancerConfig{}, NodeBalancerConfigsPagedResponse{}),
		nodebalancernodesName:        NewResource(client, nodebalancernodesName, nodebalancernodesEndpoint, true, NodeBalancerNode{}, NodeBalancerNodesPagedResponse{}),
		nodebalancerStatsName:        NewResource(client, nodebalancerStatsName, nodebalancerStatsEndpoint, true, NodeBalancerStats{}, nil),
		nodebalancersName:            NewResource(client, nodebalancersName, nodebalancersEndpoint, false, NodeBalancer{}, NodeBalancerConfigsPagedResponse{}),
		notificationsName:            NewResource(client, notificationsName, notificationsEndpoint, false, Notification{}, NotificationsPagedResponse{}),
		oauthClientsName:             NewResource(client, oauthClientsName, oauthClientsEndpoint, false, OAuthClient{}, OAuthClientsPagedResponse{}),
		objectStorageBucketsName:     NewResource(client, objectStorageBucketsName, objectStorageBucketsEndpoint, false, ObjectStorageBucket{}, ObjectStorageBucketsPagedResponse{}),
		objectStorageBucketCertsName: NewResource(client, objectStorageBucketCertsName, objectStorageBucketCertsEndpoint, true, ObjectStorageBucketCert{}, nil),
		objectStorageClustersName:    NewResource(client, objectStorageClustersName, objectStorageClustersEndpoint, false, ObjectStorageCluster{}, ObjectStorageClustersPagedResponse{}),
		objectStorageKeysName:        NewResource(client, objectStorageKeysName, objectStorageKeysEndpoint, false, ObjectStorageKey{}, ObjectStorageKeysPagedResponse{}),
		paymentsName:                 NewResource(client, paymentsName, paymentsEndpoint, false, Payment{}, PaymentsPagedResponse{}),
		profileName:                  NewResource(client, profileName, profileEndpoint, false, nil, nil), // really?
		regionsName:                  NewResource(client, regionsName, regionsEndpoint, false, Region{}, RegionsPagedResponse{}),
		sshkeysName:                  NewResource(client, sshkeysName, sshkeysEndpoint, false, SSHKey{}, SSHKeysPagedResponse{}),
		stackscriptsName:             NewResource(client, stackscriptsName, stackscriptsEndpoint, false, Stackscript{}, StackscriptsPagedResponse{}),
		tagsName:                     NewResource(client, tagsName, tagsEndpoint, false, Tag{}, TagsPagedResponse{}),
		ticketsName:                  NewResource(client, ticketsName, ticketsEndpoint, false, Ticket{}, TicketsPagedResponse{}),
		tokensName:                   NewResource(client, tokensName, tokensEndpoint, false, Token{}, TokensPagedResponse{}),
		typesName:                    NewResource(client, typesName, typesEndpoint, false, LinodeType{}, LinodeTypesPagedResponse{}),
		userGrantsName:               NewResource(client, typesName, userGrantsEndpoint, true, UserGrants{}, nil),
		usersName:                    NewResource(client, usersName, usersEndpoint, false, User{}, UsersPagedResponse{}),
		vlansName:                    NewResource(client, vlansName, vlansEndpoint, false, VLAN{}, VLANsPagedResponse{}),
		volumesName:                  NewResource(client, volumesName, volumesEndpoint, false, Volume{}, VolumesPagedResponse{}),
	}

	client.resources = resources

	client.Account = resources[accountName]
	client.DomainRecords = resources[domainRecordsName]
	client.Domains = resources[domainsName]
	client.Events = resources[eventsName]
	client.Firewalls = resources[firewallsName]
	client.FirewallDevices = resources[firewallDevicesName]
	client.FirewallRules = resources[firewallRulesName]
	client.IPAddresses = resources[ipaddressesName]
	client.IPv6Pools = resources[ipv6poolsName]
	client.IPv6Ranges = resources[ipv6rangesName]
	client.Images = resources[imagesName]
	client.InstanceConfigs = resources[instanceConfigsName]
	client.InstanceDisks = resources[instanceDisksName]
	client.InstanceIPs = resources[instanceIPsName]
	client.InstanceSnapshots = resources[instanceSnapshotsName]
	client.InstanceStats = resources[instanceStatsName]
	client.InstanceVolumes = resources[instanceVolumesName]
	client.Instances = resources[instancesName]
	client.Invoices = resources[invoicesName]
	client.Kernels = resources[kernelsName]
	client.LKEClusterAPIEndpoints = resources[lkeClusterAPIEndpointsName]
	client.LKEClusters = resources[lkeClustersName]
	client.LKEClusterPools = resources[lkeClusterPoolsName]
	client.LKEVersions = resources[lkeVersionsName]
	client.Longview = resources[longviewName]
	client.LongviewSubscriptions = resources[longviewsubscriptionsName]
	client.Managed = resources[managedName]
	client.NodeBalancerConfigs = resources[nodebalancerconfigsName]
	client.NodeBalancerNodes = resources[nodebalancernodesName]
	client.NodeBalancerStats = resources[nodebalancerStatsName]
	client.NodeBalancers = resources[nodebalancersName]
	client.Notifications = resources[notificationsName]
	client.OAuthClients = resources[oauthClientsName]
	client.ObjectStorageBuckets = resources[objectStorageBucketsName]
	client.ObjectStorageBucketCerts = resources[objectStorageBucketCertsName]
	client.ObjectStorageClusters = resources[objectStorageClustersName]
	client.ObjectStorageKeys = resources[objectStorageKeysName]
	client.Payments = resources[paymentsName]
	client.Profile = resources[profileName]
	client.Regions = resources[regionsName]
	client.SSHKeys = resources[sshkeysName]
	client.StackScripts = resources[stackscriptsName]
	client.Tags = resources[tagsName]
	client.Tickets = resources[ticketsName]
	client.Tokens = resources[tokensName]
	client.Types = resources[typesName]
	client.UserGrants = resources[userGrantsName]
	client.Users = resources[usersName]
	client.VLANs = resources[vlansName]
	client.Volumes = resources[volumesName]
}

func copyBool(bPtr *bool) *bool {
	if bPtr == nil {
		return nil
	}

	t := *bPtr

	return &t
}

func copyInt(iPtr *int) *int {
	if iPtr == nil {
		return nil
	}

	t := *iPtr

	return &t
}

func copyString(sPtr *string) *string {
	if sPtr == nil {
		return nil
	}

	t := *sPtr

	return &t
}

func copyTime(tPtr *time.Time) *time.Time {
	if tPtr == nil {
		return nil
	}

<<<<<<< HEAD
	var t = *tPtr
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	var t = *tPtr
=======
	t := *tPtr
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Version of linodego
	Version = "0.12.0"
||||||| parent of 6b7ce455e (update vendored files)
	// Version of linodego
	Version = "0.12.0"
=======
>>>>>>> 6b7ce455e (update vendored files)
	// APIEnvVar environment var to check for API token
	APIEnvVar = "LINODE_TOKEN"
	// APISecondsPerPoll how frequently to poll for new Events or Status in WaitFor functions
	APISecondsPerPoll = 3
	// Maximum wait time for retries
	APIRetryMaxWaitTime = time.Duration(30) * time.Second
)

var envDebug = false

// Client is a wrapper around the Resty client
type Client struct {
	resty             *resty.Client
	userAgent         string
	resources         map[string]*Resource
	debug             bool
	retryConditionals []RetryConditional

	millisecondsPerPoll time.Duration

	Account                  *Resource
	AccountSettings          *Resource
	DomainRecords            *Resource
	Domains                  *Resource
	Events                   *Resource
	Firewalls                *Resource
	FirewallDevices          *Resource
	FirewallRules            *Resource
	IPAddresses              *Resource
	IPv6Pools                *Resource
	IPv6Ranges               *Resource
	Images                   *Resource
	InstanceConfigs          *Resource
	InstanceDisks            *Resource
	InstanceIPs              *Resource
	InstanceSnapshots        *Resource
	InstanceStats            *Resource
	InstanceVolumes          *Resource
	Instances                *Resource
	InvoiceItems             *Resource
	Invoices                 *Resource
	Kernels                  *Resource
	LKEClusters              *Resource
	LKEClusterAPIEndpoints   *Resource
	LKEClusterPools          *Resource
	LKEVersions              *Resource
	Longview                 *Resource
	LongviewClients          *Resource
	LongviewSubscriptions    *Resource
	Managed                  *Resource
	NodeBalancerConfigs      *Resource
	NodeBalancerNodes        *Resource
	NodeBalancerStats        *Resource
	NodeBalancers            *Resource
	Notifications            *Resource
	OAuthClients             *Resource
	ObjectStorageBuckets     *Resource
	ObjectStorageBucketCerts *Resource
	ObjectStorageClusters    *Resource
	ObjectStorageKeys        *Resource
	Payments                 *Resource
	Profile                  *Resource
	Regions                  *Resource
	SSHKeys                  *Resource
	StackScripts             *Resource
	Tags                     *Resource
	Tickets                  *Resource
	Token                    *Resource
	Tokens                   *Resource
	Types                    *Resource
	UserGrants               *Resource
	Users                    *Resource
	VLANs                    *Resource
	Volumes                  *Resource
}

func init() {
	// Wether or not we will enable Resty debugging output
	if apiDebug, ok := os.LookupEnv("LINODE_DEBUG"); ok {
		if parsed, err := strconv.ParseBool(apiDebug); err == nil {
			envDebug = parsed
			log.Println("[INFO] LINODE_DEBUG being set to", envDebug)
		} else {
			log.Println("[WARN] LINODE_DEBUG should be an integer, 0 or 1")
		}
	}
}

// SetUserAgent sets a custom user-agent for HTTP requests
func (c *Client) SetUserAgent(ua string) *Client {
	c.userAgent = ua
	c.resty.SetHeader("User-Agent", c.userAgent)

	return c
}

// R wraps resty's R method
func (c *Client) R(ctx context.Context) *resty.Request {
	return c.resty.R().
		ExpectContentType("application/json").
		SetHeader("Content-Type", "application/json").
		SetContext(ctx).
		SetError(APIError{})
}

// SetDebug sets the debug on resty's client
func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	c.resty.SetDebug(debug)

	return c
}

// SetBaseURL sets the base URL of the Linode v4 API (https://api.linode.com/v4)
func (c *Client) SetBaseURL(url string) *Client {
	c.resty.SetHostURL(url)
	return c
}

// SetAPIVersion sets the version of the API to interface with
func (c *Client) SetAPIVersion(apiVersion string) *Client {
	c.SetBaseURL(fmt.Sprintf("%s://%s/%s", APIProto, APIHost, apiVersion))
	return c
}

// SetRootCertificate adds a root certificate to the underlying TLS client config
func (c *Client) SetRootCertificate(path string) *Client {
	c.resty.SetRootCertificate(path)
	return c
}

// SetToken sets the API token for all requests from this client
// Only necessary if you haven't already provided an http client to NewClient() configured with the token.
func (c *Client) SetToken(token string) *Client {
	c.resty.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	return c
}

// SetRetries adds retry conditions for "Linode Busy." errors and 429s.
func (c *Client) SetRetries() *Client {
	c.
		addRetryConditional(linodeBusyRetryCondition).
		addRetryConditional(tooManyRequestsRetryCondition).
		addRetryConditional(serviceUnavailableRetryCondition).
		addRetryConditional(requestTimeoutRetryCondition).
		SetRetryMaxWaitTime(APIRetryMaxWaitTime)
	configureRetries(c)
	return c
}

func (c *Client) addRetryConditional(retryConditional RetryConditional) *Client {
	c.retryConditionals = append(c.retryConditionals, retryConditional)
	return c
}

// SetRetryMaxWaitTime sets the maximum delay before retrying a request.
func (c *Client) SetRetryMaxWaitTime(max time.Duration) *Client {
	c.resty.SetRetryMaxWaitTime(max)
	return c
}

// SetRetryWaitTime sets the default (minimum) delay before retrying a request.
func (c *Client) SetRetryWaitTime(min time.Duration) *Client {
	c.resty.SetRetryWaitTime(min)
	return c
}

// SetRetryAfter sets the callback function to be invoked with a failed request
// to determine wben it should be retried.
func (c *Client) SetRetryAfter(callback RetryAfter) *Client {
	c.resty.SetRetryAfter(resty.RetryAfterFunc(callback))
	return c
}

// SetRetryCount sets the maximum retry attempts before aborting.
func (c *Client) SetRetryCount(count int) *Client {
	c.resty.SetRetryCount(count)
	return c
}

// SetPollDelay sets the number of milliseconds to wait between events or status polls.
// Affects all WaitFor* functions and retries.
func (c *Client) SetPollDelay(delay time.Duration) *Client {
	c.millisecondsPerPoll = delay
	return c
}

// Resource looks up a resource by name
func (c Client) Resource(resourceName string) *Resource {
	selectedResource, ok := c.resources[resourceName]
	if !ok {
		log.Fatalf("Could not find resource named '%s', exiting.", resourceName)
	}

	return selectedResource
}

// NewClient factory to create new Client struct
func NewClient(hc *http.Client) (client Client) {
	if hc != nil {
		client.resty = resty.NewWithClient(hc)
	} else {
		client.resty = resty.New()
	}

	client.SetUserAgent(DefaultUserAgent)

	baseURL, baseURLExists := os.LookupEnv(APIHostVar)

	if baseURLExists {
		client.SetBaseURL(baseURL)
	} else {
		apiVersion, apiVersionExists := os.LookupEnv(APIVersionVar)
		if apiVersionExists {
			client.SetAPIVersion(apiVersion)
		} else {
			client.SetAPIVersion(APIVersion)
		}
	}

	certPath, certPathExists := os.LookupEnv(APIHostCert)

	if certPathExists {
		cert, err := ioutil.ReadFile(certPath)
		if err != nil {
			log.Fatalf("[ERROR] Error when reading cert at %s: %s\n", certPath, err.Error())
		}

		client.SetRootCertificate(certPath)

		if envDebug {
			log.Printf("[DEBUG] Set API root certificate to %s with contents %s\n", certPath, cert)
		}
	}

	client.
		SetRetryWaitTime((1000 * APISecondsPerPoll) * time.Millisecond).
		SetPollDelay(1000 * APISecondsPerPoll).
		SetRetries().
		SetDebug(envDebug)

	addResources(&client)

	return
}

// nolint
func addResources(client *Client) {
	resources := map[string]*Resource{
		accountName:                  NewResource(client, accountName, accountEndpoint, false, Account{}, nil),                         // really?
		accountSettingsName:          NewResource(client, accountSettingsName, accountSettingsEndpoint, false, AccountSettings{}, nil), // really?
		domainRecordsName:            NewResource(client, domainRecordsName, domainRecordsEndpoint, true, DomainRecord{}, DomainRecordsPagedResponse{}),
		domainsName:                  NewResource(client, domainsName, domainsEndpoint, false, Domain{}, DomainsPagedResponse{}),
		eventsName:                   NewResource(client, eventsName, eventsEndpoint, false, Event{}, EventsPagedResponse{}),
		firewallsName:                NewResource(client, firewallsName, firewallsEndpoint, false, Firewall{}, FirewallsPagedResponse{}),
		firewallDevicesName:          NewResource(client, firewallDevicesName, firewallDevicesEndpoint, true, FirewallDevice{}, FirewallDevicesPagedResponse{}),
		firewallRulesName:            NewResource(client, firewallRulesName, firewallRulesEndpoint, true, FirewallRule{}, nil),
		imagesName:                   NewResource(client, imagesName, imagesEndpoint, false, Image{}, ImagesPagedResponse{}),
		instanceConfigsName:          NewResource(client, instanceConfigsName, instanceConfigsEndpoint, true, InstanceConfig{}, InstanceConfigsPagedResponse{}),
		instanceDisksName:            NewResource(client, instanceDisksName, instanceDisksEndpoint, true, InstanceDisk{}, InstanceDisksPagedResponse{}),
		instanceIPsName:              NewResource(client, instanceIPsName, instanceIPsEndpoint, true, InstanceIP{}, nil), // really?
		instanceSnapshotsName:        NewResource(client, instanceSnapshotsName, instanceSnapshotsEndpoint, true, InstanceSnapshot{}, nil),
		instanceStatsName:            NewResource(client, instanceStatsName, instanceStatsEndpoint, true, InstanceStats{}, nil),
		instanceVolumesName:          NewResource(client, instanceVolumesName, instanceVolumesEndpoint, true, nil, InstanceVolumesPagedResponse{}), // really?
		instancesName:                NewResource(client, instancesName, instancesEndpoint, false, Instance{}, InstancesPagedResponse{}),
		invoiceItemsName:             NewResource(client, invoiceItemsName, invoiceItemsEndpoint, true, InvoiceItem{}, InvoiceItemsPagedResponse{}),
		invoicesName:                 NewResource(client, invoicesName, invoicesEndpoint, false, Invoice{}, InvoicesPagedResponse{}),
		ipaddressesName:              NewResource(client, ipaddressesName, ipaddressesEndpoint, false, nil, IPAddressesPagedResponse{}), // really?
		ipv6poolsName:                NewResource(client, ipv6poolsName, ipv6poolsEndpoint, false, nil, IPv6PoolsPagedResponse{}),       // really?
		ipv6rangesName:               NewResource(client, ipv6rangesName, ipv6rangesEndpoint, false, IPv6Range{}, IPv6RangesPagedResponse{}),
		kernelsName:                  NewResource(client, kernelsName, kernelsEndpoint, false, LinodeKernel{}, LinodeKernelsPagedResponse{}),
		lkeClusterAPIEndpointsName:   NewResource(client, lkeClusterAPIEndpointsName, lkeClusterAPIEndpointsEndpoint, true, LKEClusterAPIEndpoint{}, LKEClusterAPIEndpointsPagedResponse{}),
		lkeClustersName:              NewResource(client, lkeClustersName, lkeClustersEndpoint, false, LKECluster{}, LKEClustersPagedResponse{}),
		lkeClusterPoolsName:          NewResource(client, lkeClusterPoolsName, lkeClusterPoolsEndpoint, true, LKEClusterPool{}, LKEClusterPoolsPagedResponse{}),
		lkeVersionsName:              NewResource(client, lkeVersionsName, lkeVersionsEndpoint, false, LKEVersion{}, LKEVersionsPagedResponse{}),
		longviewName:                 NewResource(client, longviewName, longviewEndpoint, false, nil, nil), // really?
		longviewclientsName:          NewResource(client, longviewclientsName, longviewclientsEndpoint, false, LongviewClient{}, LongviewClientsPagedResponse{}),
		longviewsubscriptionsName:    NewResource(client, longviewsubscriptionsName, longviewsubscriptionsEndpoint, false, LongviewSubscription{}, LongviewSubscriptionsPagedResponse{}),
		managedName:                  NewResource(client, managedName, managedEndpoint, false, nil, nil), // really?
		nodebalancerconfigsName:      NewResource(client, nodebalancerconfigsName, nodebalancerconfigsEndpoint, true, NodeBalancerConfig{}, NodeBalancerConfigsPagedResponse{}),
		nodebalancernodesName:        NewResource(client, nodebalancernodesName, nodebalancernodesEndpoint, true, NodeBalancerNode{}, NodeBalancerNodesPagedResponse{}),
		nodebalancerStatsName:        NewResource(client, nodebalancerStatsName, nodebalancerStatsEndpoint, true, NodeBalancerStats{}, nil),
		nodebalancersName:            NewResource(client, nodebalancersName, nodebalancersEndpoint, false, NodeBalancer{}, NodeBalancerConfigsPagedResponse{}),
		notificationsName:            NewResource(client, notificationsName, notificationsEndpoint, false, Notification{}, NotificationsPagedResponse{}),
		oauthClientsName:             NewResource(client, oauthClientsName, oauthClientsEndpoint, false, OAuthClient{}, OAuthClientsPagedResponse{}),
		objectStorageBucketsName:     NewResource(client, objectStorageBucketsName, objectStorageBucketsEndpoint, false, ObjectStorageBucket{}, ObjectStorageBucketsPagedResponse{}),
		objectStorageBucketCertsName: NewResource(client, objectStorageBucketCertsName, objectStorageBucketCertsEndpoint, true, ObjectStorageBucketCert{}, nil),
		objectStorageClustersName:    NewResource(client, objectStorageClustersName, objectStorageClustersEndpoint, false, ObjectStorageCluster{}, ObjectStorageClustersPagedResponse{}),
		objectStorageKeysName:        NewResource(client, objectStorageKeysName, objectStorageKeysEndpoint, false, ObjectStorageKey{}, ObjectStorageKeysPagedResponse{}),
		paymentsName:                 NewResource(client, paymentsName, paymentsEndpoint, false, Payment{}, PaymentsPagedResponse{}),
		profileName:                  NewResource(client, profileName, profileEndpoint, false, nil, nil), // really?
		regionsName:                  NewResource(client, regionsName, regionsEndpoint, false, Region{}, RegionsPagedResponse{}),
		sshkeysName:                  NewResource(client, sshkeysName, sshkeysEndpoint, false, SSHKey{}, SSHKeysPagedResponse{}),
		stackscriptsName:             NewResource(client, stackscriptsName, stackscriptsEndpoint, false, Stackscript{}, StackscriptsPagedResponse{}),
		tagsName:                     NewResource(client, tagsName, tagsEndpoint, false, Tag{}, TagsPagedResponse{}),
		ticketsName:                  NewResource(client, ticketsName, ticketsEndpoint, false, Ticket{}, TicketsPagedResponse{}),
		tokensName:                   NewResource(client, tokensName, tokensEndpoint, false, Token{}, TokensPagedResponse{}),
		typesName:                    NewResource(client, typesName, typesEndpoint, false, LinodeType{}, LinodeTypesPagedResponse{}),
		userGrantsName:               NewResource(client, typesName, userGrantsEndpoint, true, UserGrants{}, nil),
		usersName:                    NewResource(client, usersName, usersEndpoint, false, User{}, UsersPagedResponse{}),
		vlansName:                    NewResource(client, vlansName, vlansEndpoint, false, VLAN{}, VLANsPagedResponse{}),
		volumesName:                  NewResource(client, volumesName, volumesEndpoint, false, Volume{}, VolumesPagedResponse{}),
	}

	client.resources = resources

	client.Account = resources[accountName]
	client.DomainRecords = resources[domainRecordsName]
	client.Domains = resources[domainsName]
	client.Events = resources[eventsName]
	client.Firewalls = resources[firewallsName]
	client.FirewallDevices = resources[firewallDevicesName]
	client.FirewallRules = resources[firewallRulesName]
	client.IPAddresses = resources[ipaddressesName]
	client.IPv6Pools = resources[ipv6poolsName]
	client.IPv6Ranges = resources[ipv6rangesName]
	client.Images = resources[imagesName]
	client.InstanceConfigs = resources[instanceConfigsName]
	client.InstanceDisks = resources[instanceDisksName]
	client.InstanceIPs = resources[instanceIPsName]
	client.InstanceSnapshots = resources[instanceSnapshotsName]
	client.InstanceStats = resources[instanceStatsName]
	client.InstanceVolumes = resources[instanceVolumesName]
	client.Instances = resources[instancesName]
	client.Invoices = resources[invoicesName]
	client.Kernels = resources[kernelsName]
	client.LKEClusterAPIEndpoints = resources[lkeClusterAPIEndpointsName]
	client.LKEClusters = resources[lkeClustersName]
	client.LKEClusterPools = resources[lkeClusterPoolsName]
	client.LKEVersions = resources[lkeVersionsName]
	client.Longview = resources[longviewName]
	client.LongviewSubscriptions = resources[longviewsubscriptionsName]
	client.Managed = resources[managedName]
	client.NodeBalancerConfigs = resources[nodebalancerconfigsName]
	client.NodeBalancerNodes = resources[nodebalancernodesName]
	client.NodeBalancerStats = resources[nodebalancerStatsName]
	client.NodeBalancers = resources[nodebalancersName]
	client.Notifications = resources[notificationsName]
	client.OAuthClients = resources[oauthClientsName]
	client.ObjectStorageBuckets = resources[objectStorageBucketsName]
	client.ObjectStorageBucketCerts = resources[objectStorageBucketCertsName]
	client.ObjectStorageClusters = resources[objectStorageClustersName]
	client.ObjectStorageKeys = resources[objectStorageKeysName]
	client.Payments = resources[paymentsName]
	client.Profile = resources[profileName]
	client.Regions = resources[regionsName]
	client.SSHKeys = resources[sshkeysName]
	client.StackScripts = resources[stackscriptsName]
	client.Tags = resources[tagsName]
	client.Tickets = resources[ticketsName]
	client.Tokens = resources[tokensName]
	client.Types = resources[typesName]
	client.UserGrants = resources[userGrantsName]
	client.Users = resources[usersName]
	client.VLANs = resources[vlansName]
	client.Volumes = resources[volumesName]
}

func copyBool(bPtr *bool) *bool {
	if bPtr == nil {
		return nil
	}

	t := *bPtr

	return &t
}

func copyInt(iPtr *int) *int {
	if iPtr == nil {
		return nil
	}

	t := *iPtr

	return &t
}

func copyString(sPtr *string) *string {
	if sPtr == nil {
		return nil
	}

	t := *sPtr

	return &t
}

func copyTime(tPtr *time.Time) *time.Time {
	if tPtr == nil {
		return nil
	}

<<<<<<< HEAD
	var t = *tPtr
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	var t = *tPtr
=======
	t := *tPtr
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// Version of linodego
	Version = "0.12.0"
||||||| parent of 4d7e5ad26 (update vendored files)
	// Version of linodego
	Version = "0.12.0"
=======
>>>>>>> 4d7e5ad26 (update vendored files)
	// APIEnvVar environment var to check for API token
	APIEnvVar = "LINODE_TOKEN"
	// APISecondsPerPoll how frequently to poll for new Events or Status in WaitFor functions
	APISecondsPerPoll = 3
	// Maximum wait time for retries
	APIRetryMaxWaitTime = time.Duration(30) * time.Second
)

var envDebug = false

// Client is a wrapper around the Resty client
type Client struct {
	resty             *resty.Client
	userAgent         string
	resources         map[string]*Resource
	debug             bool
	retryConditionals []RetryConditional

	millisecondsPerPoll time.Duration

	Account                  *Resource
	AccountSettings          *Resource
	DomainRecords            *Resource
	Domains                  *Resource
	Events                   *Resource
	Firewalls                *Resource
	FirewallDevices          *Resource
	FirewallRules            *Resource
	IPAddresses              *Resource
	IPv6Pools                *Resource
	IPv6Ranges               *Resource
	Images                   *Resource
	InstanceConfigs          *Resource
	InstanceDisks            *Resource
	InstanceIPs              *Resource
	InstanceSnapshots        *Resource
	InstanceStats            *Resource
	InstanceVolumes          *Resource
	Instances                *Resource
	InvoiceItems             *Resource
	Invoices                 *Resource
	Kernels                  *Resource
	LKEClusters              *Resource
	LKEClusterAPIEndpoints   *Resource
	LKEClusterPools          *Resource
	LKEVersions              *Resource
	Longview                 *Resource
	LongviewClients          *Resource
	LongviewSubscriptions    *Resource
	Managed                  *Resource
	NodeBalancerConfigs      *Resource
	NodeBalancerNodes        *Resource
	NodeBalancerStats        *Resource
	NodeBalancers            *Resource
	Notifications            *Resource
	OAuthClients             *Resource
	ObjectStorageBuckets     *Resource
	ObjectStorageBucketCerts *Resource
	ObjectStorageClusters    *Resource
	ObjectStorageKeys        *Resource
	Payments                 *Resource
	Profile                  *Resource
	Regions                  *Resource
	SSHKeys                  *Resource
	StackScripts             *Resource
	Tags                     *Resource
	Tickets                  *Resource
	Token                    *Resource
	Tokens                   *Resource
	Types                    *Resource
	UserGrants               *Resource
	Users                    *Resource
	VLANs                    *Resource
	Volumes                  *Resource
}

func init() {
	// Wether or not we will enable Resty debugging output
	if apiDebug, ok := os.LookupEnv("LINODE_DEBUG"); ok {
		if parsed, err := strconv.ParseBool(apiDebug); err == nil {
			envDebug = parsed
			log.Println("[INFO] LINODE_DEBUG being set to", envDebug)
		} else {
			log.Println("[WARN] LINODE_DEBUG should be an integer, 0 or 1")
		}
	}
}

// SetUserAgent sets a custom user-agent for HTTP requests
func (c *Client) SetUserAgent(ua string) *Client {
	c.userAgent = ua
	c.resty.SetHeader("User-Agent", c.userAgent)

	return c
}

// R wraps resty's R method
func (c *Client) R(ctx context.Context) *resty.Request {
	return c.resty.R().
		ExpectContentType("application/json").
		SetHeader("Content-Type", "application/json").
		SetContext(ctx).
		SetError(APIError{})
}

// SetDebug sets the debug on resty's client
func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	c.resty.SetDebug(debug)

	return c
}

// SetBaseURL sets the base URL of the Linode v4 API (https://api.linode.com/v4)
func (c *Client) SetBaseURL(url string) *Client {
	c.resty.SetHostURL(url)
	return c
}

// SetAPIVersion sets the version of the API to interface with
func (c *Client) SetAPIVersion(apiVersion string) *Client {
	c.SetBaseURL(fmt.Sprintf("%s://%s/%s", APIProto, APIHost, apiVersion))
	return c
}

// SetRootCertificate adds a root certificate to the underlying TLS client config
func (c *Client) SetRootCertificate(path string) *Client {
	c.resty.SetRootCertificate(path)
	return c
}

// SetToken sets the API token for all requests from this client
// Only necessary if you haven't already provided an http client to NewClient() configured with the token.
func (c *Client) SetToken(token string) *Client {
	c.resty.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	return c
}

// SetRetries adds retry conditions for "Linode Busy." errors and 429s.
func (c *Client) SetRetries() *Client {
	c.
		addRetryConditional(linodeBusyRetryCondition).
		addRetryConditional(tooManyRequestsRetryCondition).
		addRetryConditional(serviceUnavailableRetryCondition).
		addRetryConditional(requestTimeoutRetryCondition).
		SetRetryMaxWaitTime(APIRetryMaxWaitTime)
	configureRetries(c)
	return c
}

func (c *Client) addRetryConditional(retryConditional RetryConditional) *Client {
	c.retryConditionals = append(c.retryConditionals, retryConditional)
	return c
}

// SetRetryMaxWaitTime sets the maximum delay before retrying a request.
func (c *Client) SetRetryMaxWaitTime(max time.Duration) *Client {
	c.resty.SetRetryMaxWaitTime(max)
	return c
}

// SetRetryWaitTime sets the default (minimum) delay before retrying a request.
func (c *Client) SetRetryWaitTime(min time.Duration) *Client {
	c.resty.SetRetryWaitTime(min)
	return c
}

// SetRetryAfter sets the callback function to be invoked with a failed request
// to determine wben it should be retried.
func (c *Client) SetRetryAfter(callback RetryAfter) *Client {
	c.resty.SetRetryAfter(resty.RetryAfterFunc(callback))
	return c
}

// SetRetryCount sets the maximum retry attempts before aborting.
func (c *Client) SetRetryCount(count int) *Client {
	c.resty.SetRetryCount(count)
	return c
}

// SetPollDelay sets the number of milliseconds to wait between events or status polls.
// Affects all WaitFor* functions and retries.
func (c *Client) SetPollDelay(delay time.Duration) *Client {
	c.millisecondsPerPoll = delay
	return c
}

// Resource looks up a resource by name
func (c Client) Resource(resourceName string) *Resource {
	selectedResource, ok := c.resources[resourceName]
	if !ok {
		log.Fatalf("Could not find resource named '%s', exiting.", resourceName)
	}

	return selectedResource
}

// NewClient factory to create new Client struct
func NewClient(hc *http.Client) (client Client) {
	if hc != nil {
		client.resty = resty.NewWithClient(hc)
	} else {
		client.resty = resty.New()
	}

	client.SetUserAgent(DefaultUserAgent)

	baseURL, baseURLExists := os.LookupEnv(APIHostVar)

	if baseURLExists {
		client.SetBaseURL(baseURL)
	} else {
		apiVersion, apiVersionExists := os.LookupEnv(APIVersionVar)
		if apiVersionExists {
			client.SetAPIVersion(apiVersion)
		} else {
			client.SetAPIVersion(APIVersion)
		}
	}

	certPath, certPathExists := os.LookupEnv(APIHostCert)

	if certPathExists {
		cert, err := ioutil.ReadFile(certPath)
		if err != nil {
			log.Fatalf("[ERROR] Error when reading cert at %s: %s\n", certPath, err.Error())
		}

		client.SetRootCertificate(certPath)

		if envDebug {
			log.Printf("[DEBUG] Set API root certificate to %s with contents %s\n", certPath, cert)
		}
	}

	client.
		SetRetryWaitTime((1000 * APISecondsPerPoll) * time.Millisecond).
		SetPollDelay(1000 * APISecondsPerPoll).
		SetRetries().
		SetDebug(envDebug)

	addResources(&client)

	return
}

// nolint
func addResources(client *Client) {
	resources := map[string]*Resource{
		accountName:                  NewResource(client, accountName, accountEndpoint, false, Account{}, nil),                         // really?
		accountSettingsName:          NewResource(client, accountSettingsName, accountSettingsEndpoint, false, AccountSettings{}, nil), // really?
		domainRecordsName:            NewResource(client, domainRecordsName, domainRecordsEndpoint, true, DomainRecord{}, DomainRecordsPagedResponse{}),
		domainsName:                  NewResource(client, domainsName, domainsEndpoint, false, Domain{}, DomainsPagedResponse{}),
		eventsName:                   NewResource(client, eventsName, eventsEndpoint, false, Event{}, EventsPagedResponse{}),
		firewallsName:                NewResource(client, firewallsName, firewallsEndpoint, false, Firewall{}, FirewallsPagedResponse{}),
		firewallDevicesName:          NewResource(client, firewallDevicesName, firewallDevicesEndpoint, true, FirewallDevice{}, FirewallDevicesPagedResponse{}),
		firewallRulesName:            NewResource(client, firewallRulesName, firewallRulesEndpoint, true, FirewallRule{}, nil),
		imagesName:                   NewResource(client, imagesName, imagesEndpoint, false, Image{}, ImagesPagedResponse{}),
		instanceConfigsName:          NewResource(client, instanceConfigsName, instanceConfigsEndpoint, true, InstanceConfig{}, InstanceConfigsPagedResponse{}),
		instanceDisksName:            NewResource(client, instanceDisksName, instanceDisksEndpoint, true, InstanceDisk{}, InstanceDisksPagedResponse{}),
		instanceIPsName:              NewResource(client, instanceIPsName, instanceIPsEndpoint, true, InstanceIP{}, nil), // really?
		instanceSnapshotsName:        NewResource(client, instanceSnapshotsName, instanceSnapshotsEndpoint, true, InstanceSnapshot{}, nil),
		instanceStatsName:            NewResource(client, instanceStatsName, instanceStatsEndpoint, true, InstanceStats{}, nil),
		instanceVolumesName:          NewResource(client, instanceVolumesName, instanceVolumesEndpoint, true, nil, InstanceVolumesPagedResponse{}), // really?
		instancesName:                NewResource(client, instancesName, instancesEndpoint, false, Instance{}, InstancesPagedResponse{}),
		invoiceItemsName:             NewResource(client, invoiceItemsName, invoiceItemsEndpoint, true, InvoiceItem{}, InvoiceItemsPagedResponse{}),
		invoicesName:                 NewResource(client, invoicesName, invoicesEndpoint, false, Invoice{}, InvoicesPagedResponse{}),
		ipaddressesName:              NewResource(client, ipaddressesName, ipaddressesEndpoint, false, nil, IPAddressesPagedResponse{}), // really?
		ipv6poolsName:                NewResource(client, ipv6poolsName, ipv6poolsEndpoint, false, nil, IPv6PoolsPagedResponse{}),       // really?
		ipv6rangesName:               NewResource(client, ipv6rangesName, ipv6rangesEndpoint, false, IPv6Range{}, IPv6RangesPagedResponse{}),
		kernelsName:                  NewResource(client, kernelsName, kernelsEndpoint, false, LinodeKernel{}, LinodeKernelsPagedResponse{}),
		lkeClusterAPIEndpointsName:   NewResource(client, lkeClusterAPIEndpointsName, lkeClusterAPIEndpointsEndpoint, true, LKEClusterAPIEndpoint{}, LKEClusterAPIEndpointsPagedResponse{}),
		lkeClustersName:              NewResource(client, lkeClustersName, lkeClustersEndpoint, false, LKECluster{}, LKEClustersPagedResponse{}),
		lkeClusterPoolsName:          NewResource(client, lkeClusterPoolsName, lkeClusterPoolsEndpoint, true, LKEClusterPool{}, LKEClusterPoolsPagedResponse{}),
		lkeVersionsName:              NewResource(client, lkeVersionsName, lkeVersionsEndpoint, false, LKEVersion{}, LKEVersionsPagedResponse{}),
		longviewName:                 NewResource(client, longviewName, longviewEndpoint, false, nil, nil), // really?
		longviewclientsName:          NewResource(client, longviewclientsName, longviewclientsEndpoint, false, LongviewClient{}, LongviewClientsPagedResponse{}),
		longviewsubscriptionsName:    NewResource(client, longviewsubscriptionsName, longviewsubscriptionsEndpoint, false, LongviewSubscription{}, LongviewSubscriptionsPagedResponse{}),
		managedName:                  NewResource(client, managedName, managedEndpoint, false, nil, nil), // really?
		nodebalancerconfigsName:      NewResource(client, nodebalancerconfigsName, nodebalancerconfigsEndpoint, true, NodeBalancerConfig{}, NodeBalancerConfigsPagedResponse{}),
		nodebalancernodesName:        NewResource(client, nodebalancernodesName, nodebalancernodesEndpoint, true, NodeBalancerNode{}, NodeBalancerNodesPagedResponse{}),
		nodebalancerStatsName:        NewResource(client, nodebalancerStatsName, nodebalancerStatsEndpoint, true, NodeBalancerStats{}, nil),
		nodebalancersName:            NewResource(client, nodebalancersName, nodebalancersEndpoint, false, NodeBalancer{}, NodeBalancerConfigsPagedResponse{}),
		notificationsName:            NewResource(client, notificationsName, notificationsEndpoint, false, Notification{}, NotificationsPagedResponse{}),
		oauthClientsName:             NewResource(client, oauthClientsName, oauthClientsEndpoint, false, OAuthClient{}, OAuthClientsPagedResponse{}),
		objectStorageBucketsName:     NewResource(client, objectStorageBucketsName, objectStorageBucketsEndpoint, false, ObjectStorageBucket{}, ObjectStorageBucketsPagedResponse{}),
		objectStorageBucketCertsName: NewResource(client, objectStorageBucketCertsName, objectStorageBucketCertsEndpoint, true, ObjectStorageBucketCert{}, nil),
		objectStorageClustersName:    NewResource(client, objectStorageClustersName, objectStorageClustersEndpoint, false, ObjectStorageCluster{}, ObjectStorageClustersPagedResponse{}),
		objectStorageKeysName:        NewResource(client, objectStorageKeysName, objectStorageKeysEndpoint, false, ObjectStorageKey{}, ObjectStorageKeysPagedResponse{}),
		paymentsName:                 NewResource(client, paymentsName, paymentsEndpoint, false, Payment{}, PaymentsPagedResponse{}),
		profileName:                  NewResource(client, profileName, profileEndpoint, false, nil, nil), // really?
		regionsName:                  NewResource(client, regionsName, regionsEndpoint, false, Region{}, RegionsPagedResponse{}),
		sshkeysName:                  NewResource(client, sshkeysName, sshkeysEndpoint, false, SSHKey{}, SSHKeysPagedResponse{}),
		stackscriptsName:             NewResource(client, stackscriptsName, stackscriptsEndpoint, false, Stackscript{}, StackscriptsPagedResponse{}),
		tagsName:                     NewResource(client, tagsName, tagsEndpoint, false, Tag{}, TagsPagedResponse{}),
		ticketsName:                  NewResource(client, ticketsName, ticketsEndpoint, false, Ticket{}, TicketsPagedResponse{}),
		tokensName:                   NewResource(client, tokensName, tokensEndpoint, false, Token{}, TokensPagedResponse{}),
		typesName:                    NewResource(client, typesName, typesEndpoint, false, LinodeType{}, LinodeTypesPagedResponse{}),
		userGrantsName:               NewResource(client, typesName, userGrantsEndpoint, true, UserGrants{}, nil),
		usersName:                    NewResource(client, usersName, usersEndpoint, false, User{}, UsersPagedResponse{}),
		vlansName:                    NewResource(client, vlansName, vlansEndpoint, false, VLAN{}, VLANsPagedResponse{}),
		volumesName:                  NewResource(client, volumesName, volumesEndpoint, false, Volume{}, VolumesPagedResponse{}),
	}

	client.resources = resources

	client.Account = resources[accountName]
	client.DomainRecords = resources[domainRecordsName]
	client.Domains = resources[domainsName]
	client.Events = resources[eventsName]
	client.Firewalls = resources[firewallsName]
	client.FirewallDevices = resources[firewallDevicesName]
	client.FirewallRules = resources[firewallRulesName]
	client.IPAddresses = resources[ipaddressesName]
	client.IPv6Pools = resources[ipv6poolsName]
	client.IPv6Ranges = resources[ipv6rangesName]
	client.Images = resources[imagesName]
	client.InstanceConfigs = resources[instanceConfigsName]
	client.InstanceDisks = resources[instanceDisksName]
	client.InstanceIPs = resources[instanceIPsName]
	client.InstanceSnapshots = resources[instanceSnapshotsName]
	client.InstanceStats = resources[instanceStatsName]
	client.InstanceVolumes = resources[instanceVolumesName]
	client.Instances = resources[instancesName]
	client.Invoices = resources[invoicesName]
	client.Kernels = resources[kernelsName]
	client.LKEClusterAPIEndpoints = resources[lkeClusterAPIEndpointsName]
	client.LKEClusters = resources[lkeClustersName]
	client.LKEClusterPools = resources[lkeClusterPoolsName]
	client.LKEVersions = resources[lkeVersionsName]
	client.Longview = resources[longviewName]
	client.LongviewSubscriptions = resources[longviewsubscriptionsName]
	client.Managed = resources[managedName]
	client.NodeBalancerConfigs = resources[nodebalancerconfigsName]
	client.NodeBalancerNodes = resources[nodebalancernodesName]
	client.NodeBalancerStats = resources[nodebalancerStatsName]
	client.NodeBalancers = resources[nodebalancersName]
	client.Notifications = resources[notificationsName]
	client.OAuthClients = resources[oauthClientsName]
	client.ObjectStorageBuckets = resources[objectStorageBucketsName]
	client.ObjectStorageBucketCerts = resources[objectStorageBucketCertsName]
	client.ObjectStorageClusters = resources[objectStorageClustersName]
	client.ObjectStorageKeys = resources[objectStorageKeysName]
	client.Payments = resources[paymentsName]
	client.Profile = resources[profileName]
	client.Regions = resources[regionsName]
	client.SSHKeys = resources[sshkeysName]
	client.StackScripts = resources[stackscriptsName]
	client.Tags = resources[tagsName]
	client.Tickets = resources[ticketsName]
	client.Tokens = resources[tokensName]
	client.Types = resources[typesName]
	client.UserGrants = resources[userGrantsName]
	client.Users = resources[usersName]
	client.VLANs = resources[vlansName]
	client.Volumes = resources[volumesName]
}

func copyBool(bPtr *bool) *bool {
	if bPtr == nil {
		return nil
	}

	t := *bPtr

	return &t
}

func copyInt(iPtr *int) *int {
	if iPtr == nil {
		return nil
	}

	t := *iPtr

	return &t
}

func copyString(sPtr *string) *string {
	if sPtr == nil {
		return nil
	}

	t := *sPtr

	return &t
}

func copyTime(tPtr *time.Time) *time.Time {
	if tPtr == nil {
		return nil
	}

<<<<<<< HEAD
	var t = *tPtr
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	var t = *tPtr
=======
	t := *tPtr
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	// APIHost Linode API hostname
	APIHost = "api.linode.com"
	// APIHostVar environment var to check for alternate API URL
	APIHostVar = "LINODE_URL"
	// APIHostCert environment var containing path to CA cert to validate against
	APIHostCert = "LINODE_CA"
	// APIVersion Linode API version
	APIVersion = "v4"
	// APIVersionVar environment var to check for alternate API Version
	APIVersionVar = "LINODE_API_VERSION"
	// APIProto connect to API with http(s)
	APIProto = "https"
	// Version of linodego
	Version = "0.12.0"
	// APIEnvVar environment var to check for API token
	APIEnvVar = "LINODE_TOKEN"
	// APISecondsPerPoll how frequently to poll for new Events or Status in WaitFor functions
	APISecondsPerPoll = 3
	// Maximum wait time for retries
	APIRetryMaxWaitTime = time.Duration(30) * time.Second
	// DefaultUserAgent is the default User-Agent sent in HTTP request headers
	DefaultUserAgent = "linodego " + Version + " https://github.com/linode/linodego"
)

var (
	envDebug = false
)

// Client is a wrapper around the Resty client
type Client struct {
	resty             *resty.Client
	userAgent         string
	resources         map[string]*Resource
	debug             bool
	retryConditionals []RetryConditional

	millisecondsPerPoll time.Duration

	Account                *Resource
	AccountSettings        *Resource
	DomainRecords          *Resource
	Domains                *Resource
	Events                 *Resource
	Firewalls              *Resource
	FirewallDevices        *Resource
	FirewallRules          *Resource
	IPAddresses            *Resource
	IPv6Pools              *Resource
	IPv6Ranges             *Resource
	Images                 *Resource
	InstanceConfigs        *Resource
	InstanceDisks          *Resource
	InstanceIPs            *Resource
	InstanceSnapshots      *Resource
	InstanceStats          *Resource
	InstanceVolumes        *Resource
	Instances              *Resource
	InvoiceItems           *Resource
	Invoices               *Resource
	Kernels                *Resource
	LKEClusters            *Resource
	LKEClusterAPIEndpoints *Resource
	LKEClusterPools        *Resource
	LKEVersions            *Resource
	Longview               *Resource
	LongviewClients        *Resource
	LongviewSubscriptions  *Resource
	Managed                *Resource
	NodeBalancerConfigs    *Resource
	NodeBalancerNodes      *Resource
	NodeBalancerStats      *Resource
	NodeBalancers          *Resource
	Notifications          *Resource
	OAuthClients           *Resource
	ObjectStorageBuckets   *Resource
	ObjectStorageClusters  *Resource
	ObjectStorageKeys      *Resource
	Payments               *Resource
	Profile                *Resource
	Regions                *Resource
	SSHKeys                *Resource
	StackScripts           *Resource
	Tags                   *Resource
	Tickets                *Resource
	Token                  *Resource
	Tokens                 *Resource
	Types                  *Resource
	Users                  *Resource
	Volumes                *Resource
}

func init() {
	// Wether or not we will enable Resty debugging output
	if apiDebug, ok := os.LookupEnv("LINODE_DEBUG"); ok {
		if parsed, err := strconv.ParseBool(apiDebug); err == nil {
			envDebug = parsed
			log.Println("[INFO] LINODE_DEBUG being set to", envDebug)
		} else {
			log.Println("[WARN] LINODE_DEBUG should be an integer, 0 or 1")
		}
	}
}

// SetUserAgent sets a custom user-agent for HTTP requests
func (c *Client) SetUserAgent(ua string) *Client {
	c.userAgent = ua
	c.resty.SetHeader("User-Agent", c.userAgent)

	return c
}

// R wraps resty's R method
func (c *Client) R(ctx context.Context) *resty.Request {
	return c.resty.R().
		ExpectContentType("application/json").
		SetHeader("Content-Type", "application/json").
		SetContext(ctx).
		SetError(APIError{})
}

// SetDebug sets the debug on resty's client
func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	c.resty.SetDebug(debug)

	return c
}

// SetBaseURL sets the base URL of the Linode v4 API (https://api.linode.com/v4)
func (c *Client) SetBaseURL(url string) *Client {
	c.resty.SetHostURL(url)
	return c
}

// SetAPIVersion sets the version of the API to interface with
func (c *Client) SetAPIVersion(apiVersion string) *Client {
	c.SetBaseURL(fmt.Sprintf("%s://%s/%s", APIProto, APIHost, apiVersion))
	return c
}

// SetRootCertificate adds a root certificate to the underlying TLS client config
func (c *Client) SetRootCertificate(path string) *Client {
	c.resty.SetRootCertificate(path)
	return c
}

// SetToken sets the API token for all requests from this client
// Only necessary if you haven't already provided an http client to NewClient() configured with the token.
func (c *Client) SetToken(token string) *Client {
	c.resty.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	return c
}

// SetRetries adds retry conditions for "Linode Busy." errors and 429s.
func (c *Client) SetRetries() *Client {
	c.
		addRetryConditional(linodeBusyRetryCondition).
		addRetryConditional(tooManyRequestsRetryCondition).
		addRetryConditional(serviceUnavailableRetryCondition).
		SetRetryMaxWaitTime(APIRetryMaxWaitTime)
	configureRetries(c)
	return c
}

func (c *Client) addRetryConditional(retryConditional RetryConditional) *Client {
	c.retryConditionals = append(c.retryConditionals, retryConditional)
	return c
}

func (c *Client) SetRetryMaxWaitTime(max time.Duration) *Client {
	c.resty.SetRetryMaxWaitTime(max)
	return c
}

// SetPollDelay sets the number of milliseconds to wait between events or status polls.
// Affects all WaitFor* functions and retries.
func (c *Client) SetPollDelay(delay time.Duration) *Client {
	c.millisecondsPerPoll = delay
	c.resty.SetRetryWaitTime(delay * time.Millisecond)
	return c
}

// Resource looks up a resource by name
func (c Client) Resource(resourceName string) *Resource {
	selectedResource, ok := c.resources[resourceName]
	if !ok {
		log.Fatalf("Could not find resource named '%s', exiting.", resourceName)
	}

	return selectedResource
}

// NewClient factory to create new Client struct
func NewClient(hc *http.Client) (client Client) {
	if hc != nil {
		client.resty = resty.NewWithClient(hc)
	} else {
		client.resty = resty.New()
	}

	client.SetUserAgent(DefaultUserAgent)

	baseURL, baseURLExists := os.LookupEnv(APIHostVar)

	if baseURLExists {
		client.SetBaseURL(baseURL)
	} else {
		apiVersion, apiVersionExists := os.LookupEnv(APIVersionVar)
		if apiVersionExists {
			client.SetAPIVersion(apiVersion)
		} else {
			client.SetAPIVersion(APIVersion)
		}
	}

	certPath, certPathExists := os.LookupEnv(APIHostCert)

	if certPathExists {
		cert, err := ioutil.ReadFile(certPath)
		if err != nil {
			log.Fatalf("[ERROR] Error when reading cert at %s: %s\n", certPath, err.Error())
		}

		client.SetRootCertificate(certPath)

		if envDebug {
			log.Printf("[DEBUG] Set API root certificate to %s with contents %s\n", certPath, cert)
		}
	}

	client.
		SetPollDelay(1000 * APISecondsPerPoll).
		SetRetries().
		SetDebug(envDebug)

	addResources(&client)

	return
}

// nolint
func addResources(client *Client) {
	resources := map[string]*Resource{
		accountName:                NewResource(client, accountName, accountEndpoint, false, Account{}, nil),                         // really?
		accountSettingsName:        NewResource(client, accountSettingsName, accountSettingsEndpoint, false, AccountSettings{}, nil), // really?
		domainRecordsName:          NewResource(client, domainRecordsName, domainRecordsEndpoint, true, DomainRecord{}, DomainRecordsPagedResponse{}),
		domainsName:                NewResource(client, domainsName, domainsEndpoint, false, Domain{}, DomainsPagedResponse{}),
		eventsName:                 NewResource(client, eventsName, eventsEndpoint, false, Event{}, EventsPagedResponse{}),
		firewallsName:              NewResource(client, firewallsName, firewallsEndpoint, false, Firewall{}, FirewallsPagedResponse{}),
		firewallDevicesName:        NewResource(client, firewallDevicesName, firewallDevicesEndpoint, true, FirewallDevice{}, FirewallDevicesPagedResponse{}),
		firewallRulesName:          NewResource(client, firewallRulesName, firewallRulesEndpoint, true, FirewallRule{}, nil),
		imagesName:                 NewResource(client, imagesName, imagesEndpoint, false, Image{}, ImagesPagedResponse{}),
		instanceConfigsName:        NewResource(client, instanceConfigsName, instanceConfigsEndpoint, true, InstanceConfig{}, InstanceConfigsPagedResponse{}),
		instanceDisksName:          NewResource(client, instanceDisksName, instanceDisksEndpoint, true, InstanceDisk{}, InstanceDisksPagedResponse{}),
		instanceIPsName:            NewResource(client, instanceIPsName, instanceIPsEndpoint, true, InstanceIP{}, nil), // really?
		instanceSnapshotsName:      NewResource(client, instanceSnapshotsName, instanceSnapshotsEndpoint, true, InstanceSnapshot{}, nil),
		instanceStatsName:          NewResource(client, instanceStatsName, instanceStatsEndpoint, true, InstanceStats{}, nil),
		instanceVolumesName:        NewResource(client, instanceVolumesName, instanceVolumesEndpoint, true, nil, InstanceVolumesPagedResponse{}), // really?
		instancesName:              NewResource(client, instancesName, instancesEndpoint, false, Instance{}, InstancesPagedResponse{}),
		invoiceItemsName:           NewResource(client, invoiceItemsName, invoiceItemsEndpoint, true, InvoiceItem{}, InvoiceItemsPagedResponse{}),
		invoicesName:               NewResource(client, invoicesName, invoicesEndpoint, false, Invoice{}, InvoicesPagedResponse{}),
		ipaddressesName:            NewResource(client, ipaddressesName, ipaddressesEndpoint, false, nil, IPAddressesPagedResponse{}), // really?
		ipv6poolsName:              NewResource(client, ipv6poolsName, ipv6poolsEndpoint, false, nil, IPv6PoolsPagedResponse{}),       // really?
		ipv6rangesName:             NewResource(client, ipv6rangesName, ipv6rangesEndpoint, false, IPv6Range{}, IPv6RangesPagedResponse{}),
		kernelsName:                NewResource(client, kernelsName, kernelsEndpoint, false, LinodeKernel{}, LinodeKernelsPagedResponse{}),
		lkeClusterAPIEndpointsName: NewResource(client, lkeClusterAPIEndpointsName, lkeClusterAPIEndpointsEndpoint, true, LKEClusterAPIEndpoint{}, LKEClusterAPIEndpointsPagedResponse{}),
		lkeClustersName:            NewResource(client, lkeClustersName, lkeClustersEndpoint, false, LKECluster{}, LKEClustersPagedResponse{}),
		lkeClusterPoolsName:        NewResource(client, lkeClusterPoolsName, lkeClusterPoolsEndpoint, true, LKEClusterPool{}, LKEClusterPoolsPagedResponse{}),
		lkeVersionsName:            NewResource(client, lkeVersionsName, lkeVersionsEndpoint, false, LKEVersion{}, LKEVersionsPagedResponse{}),
		longviewName:               NewResource(client, longviewName, longviewEndpoint, false, nil, nil), // really?
		longviewclientsName:        NewResource(client, longviewclientsName, longviewclientsEndpoint, false, LongviewClient{}, LongviewClientsPagedResponse{}),
		longviewsubscriptionsName:  NewResource(client, longviewsubscriptionsName, longviewsubscriptionsEndpoint, false, LongviewSubscription{}, LongviewSubscriptionsPagedResponse{}),
		managedName:                NewResource(client, managedName, managedEndpoint, false, nil, nil), // really?
		nodebalancerconfigsName:    NewResource(client, nodebalancerconfigsName, nodebalancerconfigsEndpoint, true, NodeBalancerConfig{}, NodeBalancerConfigsPagedResponse{}),
		nodebalancernodesName:      NewResource(client, nodebalancernodesName, nodebalancernodesEndpoint, true, NodeBalancerNode{}, NodeBalancerNodesPagedResponse{}),
		nodebalancerStatsName:      NewResource(client, nodebalancerStatsName, nodebalancerStatsEndpoint, true, NodeBalancerStats{}, nil),
		nodebalancersName:          NewResource(client, nodebalancersName, nodebalancersEndpoint, false, NodeBalancer{}, NodeBalancerConfigsPagedResponse{}),
		notificationsName:          NewResource(client, notificationsName, notificationsEndpoint, false, Notification{}, NotificationsPagedResponse{}),
		oauthClientsName:           NewResource(client, oauthClientsName, oauthClientsEndpoint, false, OAuthClient{}, OAuthClientsPagedResponse{}),
		objectStorageBucketsName:   NewResource(client, objectStorageBucketsName, objectStorageBucketsEndpoint, false, ObjectStorageBucket{}, ObjectStorageBucketsPagedResponse{}),
		objectStorageClustersName:  NewResource(client, objectStorageClustersName, objectStorageClustersEndpoint, false, ObjectStorageCluster{}, ObjectStorageClustersPagedResponse{}),
		objectStorageKeysName:      NewResource(client, objectStorageKeysName, objectStorageKeysEndpoint, false, ObjectStorageKey{}, ObjectStorageKeysPagedResponse{}),
		paymentsName:               NewResource(client, paymentsName, paymentsEndpoint, false, Payment{}, PaymentsPagedResponse{}),
		profileName:                NewResource(client, profileName, profileEndpoint, false, nil, nil), // really?
		regionsName:                NewResource(client, regionsName, regionsEndpoint, false, Region{}, RegionsPagedResponse{}),
		sshkeysName:                NewResource(client, sshkeysName, sshkeysEndpoint, false, SSHKey{}, SSHKeysPagedResponse{}),
		stackscriptsName:           NewResource(client, stackscriptsName, stackscriptsEndpoint, false, Stackscript{}, StackscriptsPagedResponse{}),
		tagsName:                   NewResource(client, tagsName, tagsEndpoint, false, Tag{}, TagsPagedResponse{}),
		ticketsName:                NewResource(client, ticketsName, ticketsEndpoint, false, Ticket{}, TicketsPagedResponse{}),
		tokensName:                 NewResource(client, tokensName, tokensEndpoint, false, Token{}, TokensPagedResponse{}),
		typesName:                  NewResource(client, typesName, typesEndpoint, false, LinodeType{}, LinodeTypesPagedResponse{}),
		usersName:                  NewResource(client, usersName, usersEndpoint, false, User{}, UsersPagedResponse{}),
		volumesName:                NewResource(client, volumesName, volumesEndpoint, false, Volume{}, VolumesPagedResponse{}),
	}

	client.resources = resources

	client.Account = resources[accountName]
	client.DomainRecords = resources[domainRecordsName]
	client.Domains = resources[domainsName]
	client.Events = resources[eventsName]
	client.Firewalls = resources[firewallsName]
	client.FirewallDevices = resources[firewallDevicesName]
	client.FirewallRules = resources[firewallRulesName]
	client.IPAddresses = resources[ipaddressesName]
	client.IPv6Pools = resources[ipv6poolsName]
	client.IPv6Ranges = resources[ipv6rangesName]
	client.Images = resources[imagesName]
	client.InstanceConfigs = resources[instanceConfigsName]
	client.InstanceDisks = resources[instanceDisksName]
	client.InstanceIPs = resources[instanceIPsName]
	client.InstanceSnapshots = resources[instanceSnapshotsName]
	client.InstanceStats = resources[instanceStatsName]
	client.InstanceVolumes = resources[instanceVolumesName]
	client.Instances = resources[instancesName]
	client.Invoices = resources[invoicesName]
	client.Kernels = resources[kernelsName]
	client.LKEClusterAPIEndpoints = resources[lkeClusterAPIEndpointsName]
	client.LKEClusters = resources[lkeClustersName]
	client.LKEClusterPools = resources[lkeClusterPoolsName]
	client.LKEVersions = resources[lkeVersionsName]
	client.Longview = resources[longviewName]
	client.LongviewSubscriptions = resources[longviewsubscriptionsName]
	client.Managed = resources[managedName]
	client.NodeBalancerConfigs = resources[nodebalancerconfigsName]
	client.NodeBalancerNodes = resources[nodebalancernodesName]
	client.NodeBalancerStats = resources[nodebalancerStatsName]
	client.NodeBalancers = resources[nodebalancersName]
	client.Notifications = resources[notificationsName]
	client.OAuthClients = resources[oauthClientsName]
	client.ObjectStorageBuckets = resources[objectStorageBucketsName]
	client.ObjectStorageClusters = resources[objectStorageClustersName]
	client.ObjectStorageKeys = resources[objectStorageKeysName]
	client.Payments = resources[paymentsName]
	client.Profile = resources[profileName]
	client.Regions = resources[regionsName]
	client.SSHKeys = resources[sshkeysName]
	client.StackScripts = resources[stackscriptsName]
	client.Tags = resources[tagsName]
	client.Tickets = resources[ticketsName]
	client.Tokens = resources[tokensName]
	client.Types = resources[typesName]
	client.Users = resources[usersName]
	client.Volumes = resources[volumesName]
}

func copyBool(bPtr *bool) *bool {
	if bPtr == nil {
		return nil
	}

	var t = *bPtr

	return &t
}

func copyInt(iPtr *int) *int {
	if iPtr == nil {
		return nil
	}

	var t = *iPtr

	return &t
}

func copyString(sPtr *string) *string {
	if sPtr == nil {
		return nil
	}

	var t = *sPtr

	return &t
}

func copyTime(tPtr *time.Time) *time.Time {
	if tPtr == nil {
		return nil
	}

	var t = *tPtr
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

	return &t
}
