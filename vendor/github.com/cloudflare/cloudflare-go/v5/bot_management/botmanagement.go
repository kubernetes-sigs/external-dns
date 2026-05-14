// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// BotManagementService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBotManagementService] method instead.
type BotManagementService struct {
	Options []option.RequestOption
}

// NewBotManagementService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBotManagementService(opts ...option.RequestOption) (r *BotManagementService) {
	r = &BotManagementService{}
	r.Options = opts
	return
}

// Updates the Bot Management configuration for a zone.
//
// This API is used to update:
//
// - **Bot Fight Mode**
// - **Super Bot Fight Mode**
// - **Bot Management for Enterprise**
//
// See [Bot Plans](https://developers.cloudflare.com/bots/plans/) for more
// information on the different plans \
// If you recently upgraded or downgraded your plan, refer to the following examples
// to clean up old configurations. Copy and paste the example body to remove old zone
// configurations based on your current plan.
//
// #### Clean up configuration for Bot Fight Mode plan
//
// ```json
//
//	{
//	  "sbfm_likely_automated": "allow",
//	  "sbfm_definitely_automated": "allow",
//	  "sbfm_verified_bots": "allow",
//	  "sbfm_static_resource_protection": false,
//	  "optimize_wordpress": false,
//	  "suppress_session_score": false
//	}
//
// ```
//
// #### Clean up configuration for SBFM Pro plan
//
// ```json
//
//	{
//	  "sbfm_likely_automated": "allow",
//	  "fight_mode": false
//	}
//
// ```
//
// #### Clean up configuration for SBFM Biz plan
//
// ```json
//
//	{
//	  "fight_mode": false
//	}
//
// ```
//
// #### Clean up configuration for BM Enterprise Subscription plan
//
// It is strongly recommended that you ensure you have
// [custom rules](https://developers.cloudflare.com/waf/custom-rules/) in place to
// protect your zone before disabling the SBFM rules. Without these protections,
// your zone is vulnerable to attacks.
//
// ```json
//
//	{
//	  "sbfm_likely_automated": "allow",
//	  "sbfm_definitely_automated": "allow",
//	  "sbfm_verified_bots": "allow",
//	  "sbfm_static_resource_protection": false,
//	  "optimize_wordpress": false,
//	  "fight_mode": false
//	}
//
// ```
func (r *BotManagementService) Update(ctx context.Context, params BotManagementUpdateParams, opts ...option.RequestOption) (res *BotManagementUpdateResponse, err error) {
	var env BotManagementUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/bot_management", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve a zone's Bot Management Config
func (r *BotManagementService) Get(ctx context.Context, query BotManagementGetParams, opts ...option.RequestOption) (res *BotManagementGetResponse, err error) {
	var env BotManagementGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/bot_management", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BotFightModeConfiguration struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection BotFightModeConfigurationAIBotsProtection `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection BotFightModeConfigurationCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged bool `json:"is_robots_txt_managed"`
	// A read-only field that shows which unauthorized settings are currently active on
	// the zone. These settings typically result from upgrades or downgrades.
	StaleZoneConfiguration BotFightModeConfigurationStaleZoneConfiguration `json:"stale_zone_configuration"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                          `json:"using_latest_model"`
	JSON             botFightModeConfigurationJSON `json:"-"`
}

// botFightModeConfigurationJSON contains the JSON metadata for the struct
// [BotFightModeConfiguration]
type botFightModeConfigurationJSON struct {
	AIBotsProtection       apijson.Field
	CrawlerProtection      apijson.Field
	EnableJS               apijson.Field
	FightMode              apijson.Field
	IsRobotsTXTManaged     apijson.Field
	StaleZoneConfiguration apijson.Field
	UsingLatestModel       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BotFightModeConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botFightModeConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r BotFightModeConfiguration) implementsBotManagementUpdateResponse() {}

func (r BotFightModeConfiguration) implementsBotManagementGetResponse() {}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type BotFightModeConfigurationAIBotsProtection string

const (
	BotFightModeConfigurationAIBotsProtectionBlock         BotFightModeConfigurationAIBotsProtection = "block"
	BotFightModeConfigurationAIBotsProtectionDisabled      BotFightModeConfigurationAIBotsProtection = "disabled"
	BotFightModeConfigurationAIBotsProtectionOnlyOnADPages BotFightModeConfigurationAIBotsProtection = "only_on_ad_pages"
)

func (r BotFightModeConfigurationAIBotsProtection) IsKnown() bool {
	switch r {
	case BotFightModeConfigurationAIBotsProtectionBlock, BotFightModeConfigurationAIBotsProtectionDisabled, BotFightModeConfigurationAIBotsProtectionOnlyOnADPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type BotFightModeConfigurationCrawlerProtection string

const (
	BotFightModeConfigurationCrawlerProtectionEnabled  BotFightModeConfigurationCrawlerProtection = "enabled"
	BotFightModeConfigurationCrawlerProtectionDisabled BotFightModeConfigurationCrawlerProtection = "disabled"
)

func (r BotFightModeConfigurationCrawlerProtection) IsKnown() bool {
	switch r {
	case BotFightModeConfigurationCrawlerProtectionEnabled, BotFightModeConfigurationCrawlerProtectionDisabled:
		return true
	}
	return false
}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type BotFightModeConfigurationStaleZoneConfiguration struct {
	// Indicates that the zone's wordpress optimization for SBFM is turned on.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Indicates that the zone's definitely automated requests are being blocked or
	// challenged.
	SBFMDefinitelyAutomated string `json:"sbfm_definitely_automated"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SBFMLikelyAutomated string `json:"sbfm_likely_automated"`
	// Indicates that the zone's static resource protection is turned on.
	SBFMStaticResourceProtection string `json:"sbfm_static_resource_protection"`
	// Indicates that the zone's verified bot requests are being blocked.
	SBFMVerifiedBots string `json:"sbfm_verified_bots"`
	// Indicates that the zone's session score tracking is disabled.
	SuppressSessionScore bool                                                `json:"suppress_session_score"`
	JSON                 botFightModeConfigurationStaleZoneConfigurationJSON `json:"-"`
}

// botFightModeConfigurationStaleZoneConfigurationJSON contains the JSON metadata
// for the struct [BotFightModeConfigurationStaleZoneConfiguration]
type botFightModeConfigurationStaleZoneConfigurationJSON struct {
	OptimizeWordpress            apijson.Field
	SBFMDefinitelyAutomated      apijson.Field
	SBFMLikelyAutomated          apijson.Field
	SBFMStaticResourceProtection apijson.Field
	SBFMVerifiedBots             apijson.Field
	SuppressSessionScore         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *BotFightModeConfigurationStaleZoneConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botFightModeConfigurationStaleZoneConfigurationJSON) RawJSON() string {
	return r.raw
}

type BotFightModeConfigurationParam struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection param.Field[BotFightModeConfigurationAIBotsProtection] `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection param.Field[BotFightModeConfigurationCrawlerProtection] `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged param.Field[bool] `json:"is_robots_txt_managed"`
}

func (r BotFightModeConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BotFightModeConfigurationParam) implementsBotManagementUpdateParamsBodyUnion() {}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type BotFightModeConfigurationStaleZoneConfigurationParam struct {
	// Indicates that the zone's wordpress optimization for SBFM is turned on.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Indicates that the zone's definitely automated requests are being blocked or
	// challenged.
	SBFMDefinitelyAutomated param.Field[string] `json:"sbfm_definitely_automated"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SBFMLikelyAutomated param.Field[string] `json:"sbfm_likely_automated"`
	// Indicates that the zone's static resource protection is turned on.
	SBFMStaticResourceProtection param.Field[string] `json:"sbfm_static_resource_protection"`
	// Indicates that the zone's verified bot requests are being blocked.
	SBFMVerifiedBots param.Field[string] `json:"sbfm_verified_bots"`
	// Indicates that the zone's session score tracking is disabled.
	SuppressSessionScore param.Field[bool] `json:"suppress_session_score"`
}

func (r BotFightModeConfigurationStaleZoneConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionConfiguration struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection SubscriptionConfigurationAIBotsProtection `json:"ai_bots_protection"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection SubscriptionConfigurationCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged bool `json:"is_robots_txt_managed"`
	// A read-only field that shows which unauthorized settings are currently active on
	// the zone. These settings typically result from upgrades or downgrades.
	StaleZoneConfiguration SubscriptionConfigurationStaleZoneConfiguration `json:"stale_zone_configuration"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool `json:"suppress_session_score"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                          `json:"using_latest_model"`
	JSON             subscriptionConfigurationJSON `json:"-"`
}

// subscriptionConfigurationJSON contains the JSON metadata for the struct
// [SubscriptionConfiguration]
type subscriptionConfigurationJSON struct {
	AIBotsProtection       apijson.Field
	AutoUpdateModel        apijson.Field
	CrawlerProtection      apijson.Field
	EnableJS               apijson.Field
	IsRobotsTXTManaged     apijson.Field
	StaleZoneConfiguration apijson.Field
	SuppressSessionScore   apijson.Field
	UsingLatestModel       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SubscriptionConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionConfiguration) implementsBotManagementUpdateResponse() {}

func (r SubscriptionConfiguration) implementsBotManagementGetResponse() {}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type SubscriptionConfigurationAIBotsProtection string

const (
	SubscriptionConfigurationAIBotsProtectionBlock         SubscriptionConfigurationAIBotsProtection = "block"
	SubscriptionConfigurationAIBotsProtectionDisabled      SubscriptionConfigurationAIBotsProtection = "disabled"
	SubscriptionConfigurationAIBotsProtectionOnlyOnADPages SubscriptionConfigurationAIBotsProtection = "only_on_ad_pages"
)

func (r SubscriptionConfigurationAIBotsProtection) IsKnown() bool {
	switch r {
	case SubscriptionConfigurationAIBotsProtectionBlock, SubscriptionConfigurationAIBotsProtectionDisabled, SubscriptionConfigurationAIBotsProtectionOnlyOnADPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type SubscriptionConfigurationCrawlerProtection string

const (
	SubscriptionConfigurationCrawlerProtectionEnabled  SubscriptionConfigurationCrawlerProtection = "enabled"
	SubscriptionConfigurationCrawlerProtectionDisabled SubscriptionConfigurationCrawlerProtection = "disabled"
)

func (r SubscriptionConfigurationCrawlerProtection) IsKnown() bool {
	switch r {
	case SubscriptionConfigurationCrawlerProtectionEnabled, SubscriptionConfigurationCrawlerProtectionDisabled:
		return true
	}
	return false
}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type SubscriptionConfigurationStaleZoneConfiguration struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode bool `json:"fight_mode"`
	// Indicates that the zone's wordpress optimization for SBFM is turned on.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Indicates that the zone's definitely automated requests are being blocked or
	// challenged.
	SBFMDefinitelyAutomated string `json:"sbfm_definitely_automated"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SBFMLikelyAutomated string `json:"sbfm_likely_automated"`
	// Indicates that the zone's static resource protection is turned on.
	SBFMStaticResourceProtection string `json:"sbfm_static_resource_protection"`
	// Indicates that the zone's verified bot requests are being blocked.
	SBFMVerifiedBots string                                              `json:"sbfm_verified_bots"`
	JSON             subscriptionConfigurationStaleZoneConfigurationJSON `json:"-"`
}

// subscriptionConfigurationStaleZoneConfigurationJSON contains the JSON metadata
// for the struct [SubscriptionConfigurationStaleZoneConfiguration]
type subscriptionConfigurationStaleZoneConfigurationJSON struct {
	FightMode                    apijson.Field
	OptimizeWordpress            apijson.Field
	SBFMDefinitelyAutomated      apijson.Field
	SBFMLikelyAutomated          apijson.Field
	SBFMStaticResourceProtection apijson.Field
	SBFMVerifiedBots             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SubscriptionConfigurationStaleZoneConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionConfigurationStaleZoneConfigurationJSON) RawJSON() string {
	return r.raw
}

type SubscriptionConfigurationParam struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection param.Field[SubscriptionConfigurationAIBotsProtection] `json:"ai_bots_protection"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel param.Field[bool] `json:"auto_update_model"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection param.Field[SubscriptionConfigurationCrawlerProtection] `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS param.Field[bool] `json:"enable_js"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged param.Field[bool] `json:"is_robots_txt_managed"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore param.Field[bool] `json:"suppress_session_score"`
}

func (r SubscriptionConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionConfigurationParam) implementsBotManagementUpdateParamsBodyUnion() {}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type SubscriptionConfigurationStaleZoneConfigurationParam struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode param.Field[bool] `json:"fight_mode"`
	// Indicates that the zone's wordpress optimization for SBFM is turned on.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Indicates that the zone's definitely automated requests are being blocked or
	// challenged.
	SBFMDefinitelyAutomated param.Field[string] `json:"sbfm_definitely_automated"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SBFMLikelyAutomated param.Field[string] `json:"sbfm_likely_automated"`
	// Indicates that the zone's static resource protection is turned on.
	SBFMStaticResourceProtection param.Field[string] `json:"sbfm_static_resource_protection"`
	// Indicates that the zone's verified bot requests are being blocked.
	SBFMVerifiedBots param.Field[string] `json:"sbfm_verified_bots"`
}

func (r SubscriptionConfigurationStaleZoneConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperBotFightModeDefinitelyConfiguration struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection SuperBotFightModeDefinitelyConfigurationAIBotsProtection `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection SuperBotFightModeDefinitelyConfigurationCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged bool `json:"is_robots_txt_managed"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that shows which unauthorized settings are currently active on
	// the zone. These settings typically result from upgrades or downgrades.
	StaleZoneConfiguration SuperBotFightModeDefinitelyConfigurationStaleZoneConfiguration `json:"stale_zone_configuration"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                         `json:"using_latest_model"`
	JSON             superBotFightModeDefinitelyConfigurationJSON `json:"-"`
}

// superBotFightModeDefinitelyConfigurationJSON contains the JSON metadata for the
// struct [SuperBotFightModeDefinitelyConfiguration]
type superBotFightModeDefinitelyConfigurationJSON struct {
	AIBotsProtection             apijson.Field
	CrawlerProtection            apijson.Field
	EnableJS                     apijson.Field
	IsRobotsTXTManaged           apijson.Field
	OptimizeWordpress            apijson.Field
	SBFMDefinitelyAutomated      apijson.Field
	SBFMStaticResourceProtection apijson.Field
	SBFMVerifiedBots             apijson.Field
	StaleZoneConfiguration       apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SuperBotFightModeDefinitelyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superBotFightModeDefinitelyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r SuperBotFightModeDefinitelyConfiguration) implementsBotManagementUpdateResponse() {}

func (r SuperBotFightModeDefinitelyConfiguration) implementsBotManagementGetResponse() {}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type SuperBotFightModeDefinitelyConfigurationAIBotsProtection string

const (
	SuperBotFightModeDefinitelyConfigurationAIBotsProtectionBlock         SuperBotFightModeDefinitelyConfigurationAIBotsProtection = "block"
	SuperBotFightModeDefinitelyConfigurationAIBotsProtectionDisabled      SuperBotFightModeDefinitelyConfigurationAIBotsProtection = "disabled"
	SuperBotFightModeDefinitelyConfigurationAIBotsProtectionOnlyOnADPages SuperBotFightModeDefinitelyConfigurationAIBotsProtection = "only_on_ad_pages"
)

func (r SuperBotFightModeDefinitelyConfigurationAIBotsProtection) IsKnown() bool {
	switch r {
	case SuperBotFightModeDefinitelyConfigurationAIBotsProtectionBlock, SuperBotFightModeDefinitelyConfigurationAIBotsProtectionDisabled, SuperBotFightModeDefinitelyConfigurationAIBotsProtectionOnlyOnADPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type SuperBotFightModeDefinitelyConfigurationCrawlerProtection string

const (
	SuperBotFightModeDefinitelyConfigurationCrawlerProtectionEnabled  SuperBotFightModeDefinitelyConfigurationCrawlerProtection = "enabled"
	SuperBotFightModeDefinitelyConfigurationCrawlerProtectionDisabled SuperBotFightModeDefinitelyConfigurationCrawlerProtection = "disabled"
)

func (r SuperBotFightModeDefinitelyConfigurationCrawlerProtection) IsKnown() bool {
	switch r {
	case SuperBotFightModeDefinitelyConfigurationCrawlerProtectionEnabled, SuperBotFightModeDefinitelyConfigurationCrawlerProtectionDisabled:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated string

const (
	SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedAllow            SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated = "allow"
	SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedBlock            SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated = "block"
	SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedManagedChallenge SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated = "managed_challenge"
)

func (r SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated) IsKnown() bool {
	switch r {
	case SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedAllow, SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedBlock, SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots string

const (
	SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBotsAllow SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots = "allow"
	SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBotsBlock SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots = "block"
)

func (r SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots) IsKnown() bool {
	switch r {
	case SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBotsAllow, SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBotsBlock:
		return true
	}
	return false
}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type SuperBotFightModeDefinitelyConfigurationStaleZoneConfiguration struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode bool `json:"fight_mode"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SBFMLikelyAutomated string                                                             `json:"sbfm_likely_automated"`
	JSON                superBotFightModeDefinitelyConfigurationStaleZoneConfigurationJSON `json:"-"`
}

// superBotFightModeDefinitelyConfigurationStaleZoneConfigurationJSON contains the
// JSON metadata for the struct
// [SuperBotFightModeDefinitelyConfigurationStaleZoneConfiguration]
type superBotFightModeDefinitelyConfigurationStaleZoneConfigurationJSON struct {
	FightMode           apijson.Field
	SBFMLikelyAutomated apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SuperBotFightModeDefinitelyConfigurationStaleZoneConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superBotFightModeDefinitelyConfigurationStaleZoneConfigurationJSON) RawJSON() string {
	return r.raw
}

type SuperBotFightModeDefinitelyConfigurationParam struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection param.Field[SuperBotFightModeDefinitelyConfigurationAIBotsProtection] `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection param.Field[SuperBotFightModeDefinitelyConfigurationCrawlerProtection] `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS param.Field[bool] `json:"enable_js"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged param.Field[bool] `json:"is_robots_txt_managed"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated param.Field[SuperBotFightModeDefinitelyConfigurationSBFMDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots param.Field[SuperBotFightModeDefinitelyConfigurationSBFMVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r SuperBotFightModeDefinitelyConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperBotFightModeDefinitelyConfigurationParam) implementsBotManagementUpdateParamsBodyUnion() {
}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type SuperBotFightModeDefinitelyConfigurationStaleZoneConfigurationParam struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode param.Field[bool] `json:"fight_mode"`
	// Indicates that the zone's likely automated requests are being blocked or
	// challenged.
	SBFMLikelyAutomated param.Field[string] `json:"sbfm_likely_automated"`
}

func (r SuperBotFightModeDefinitelyConfigurationStaleZoneConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperBotFightModeLikelyConfiguration struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection SuperBotFightModeLikelyConfigurationAIBotsProtection `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection SuperBotFightModeLikelyConfigurationCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged bool `json:"is_robots_txt_managed"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SBFMLikelyAutomated SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots SuperBotFightModeLikelyConfigurationSBFMVerifiedBots `json:"sbfm_verified_bots"`
	// A read-only field that shows which unauthorized settings are currently active on
	// the zone. These settings typically result from upgrades or downgrades.
	StaleZoneConfiguration SuperBotFightModeLikelyConfigurationStaleZoneConfiguration `json:"stale_zone_configuration"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                                     `json:"using_latest_model"`
	JSON             superBotFightModeLikelyConfigurationJSON `json:"-"`
}

// superBotFightModeLikelyConfigurationJSON contains the JSON metadata for the
// struct [SuperBotFightModeLikelyConfiguration]
type superBotFightModeLikelyConfigurationJSON struct {
	AIBotsProtection             apijson.Field
	CrawlerProtection            apijson.Field
	EnableJS                     apijson.Field
	IsRobotsTXTManaged           apijson.Field
	OptimizeWordpress            apijson.Field
	SBFMDefinitelyAutomated      apijson.Field
	SBFMLikelyAutomated          apijson.Field
	SBFMStaticResourceProtection apijson.Field
	SBFMVerifiedBots             apijson.Field
	StaleZoneConfiguration       apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *SuperBotFightModeLikelyConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superBotFightModeLikelyConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r SuperBotFightModeLikelyConfiguration) implementsBotManagementUpdateResponse() {}

func (r SuperBotFightModeLikelyConfiguration) implementsBotManagementGetResponse() {}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type SuperBotFightModeLikelyConfigurationAIBotsProtection string

const (
	SuperBotFightModeLikelyConfigurationAIBotsProtectionBlock         SuperBotFightModeLikelyConfigurationAIBotsProtection = "block"
	SuperBotFightModeLikelyConfigurationAIBotsProtectionDisabled      SuperBotFightModeLikelyConfigurationAIBotsProtection = "disabled"
	SuperBotFightModeLikelyConfigurationAIBotsProtectionOnlyOnADPages SuperBotFightModeLikelyConfigurationAIBotsProtection = "only_on_ad_pages"
)

func (r SuperBotFightModeLikelyConfigurationAIBotsProtection) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationAIBotsProtectionBlock, SuperBotFightModeLikelyConfigurationAIBotsProtectionDisabled, SuperBotFightModeLikelyConfigurationAIBotsProtectionOnlyOnADPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type SuperBotFightModeLikelyConfigurationCrawlerProtection string

const (
	SuperBotFightModeLikelyConfigurationCrawlerProtectionEnabled  SuperBotFightModeLikelyConfigurationCrawlerProtection = "enabled"
	SuperBotFightModeLikelyConfigurationCrawlerProtectionDisabled SuperBotFightModeLikelyConfigurationCrawlerProtection = "disabled"
)

func (r SuperBotFightModeLikelyConfigurationCrawlerProtection) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationCrawlerProtectionEnabled, SuperBotFightModeLikelyConfigurationCrawlerProtectionDisabled:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated string

const (
	SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedAllow            SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated = "allow"
	SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedBlock            SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated = "block"
	SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedManagedChallenge SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated = "managed_challenge"
)

func (r SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedAllow, SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedBlock, SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated string

const (
	SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedAllow            SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated = "allow"
	SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedBlock            SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated = "block"
	SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedManagedChallenge SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated = "managed_challenge"
)

func (r SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedAllow, SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedBlock, SuperBotFightModeLikelyConfigurationSBFMLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type SuperBotFightModeLikelyConfigurationSBFMVerifiedBots string

const (
	SuperBotFightModeLikelyConfigurationSBFMVerifiedBotsAllow SuperBotFightModeLikelyConfigurationSBFMVerifiedBots = "allow"
	SuperBotFightModeLikelyConfigurationSBFMVerifiedBotsBlock SuperBotFightModeLikelyConfigurationSBFMVerifiedBots = "block"
)

func (r SuperBotFightModeLikelyConfigurationSBFMVerifiedBots) IsKnown() bool {
	switch r {
	case SuperBotFightModeLikelyConfigurationSBFMVerifiedBotsAllow, SuperBotFightModeLikelyConfigurationSBFMVerifiedBotsBlock:
		return true
	}
	return false
}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type SuperBotFightModeLikelyConfigurationStaleZoneConfiguration struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode bool                                                           `json:"fight_mode"`
	JSON      superBotFightModeLikelyConfigurationStaleZoneConfigurationJSON `json:"-"`
}

// superBotFightModeLikelyConfigurationStaleZoneConfigurationJSON contains the JSON
// metadata for the struct
// [SuperBotFightModeLikelyConfigurationStaleZoneConfiguration]
type superBotFightModeLikelyConfigurationStaleZoneConfigurationJSON struct {
	FightMode   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuperBotFightModeLikelyConfigurationStaleZoneConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superBotFightModeLikelyConfigurationStaleZoneConfigurationJSON) RawJSON() string {
	return r.raw
}

type SuperBotFightModeLikelyConfigurationParam struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection param.Field[SuperBotFightModeLikelyConfigurationAIBotsProtection] `json:"ai_bots_protection"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection param.Field[SuperBotFightModeLikelyConfigurationCrawlerProtection] `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS param.Field[bool] `json:"enable_js"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged param.Field[bool] `json:"is_robots_txt_managed"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated param.Field[SuperBotFightModeLikelyConfigurationSBFMDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SBFMLikelyAutomated param.Field[SuperBotFightModeLikelyConfigurationSBFMLikelyAutomated] `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots param.Field[SuperBotFightModeLikelyConfigurationSBFMVerifiedBots] `json:"sbfm_verified_bots"`
}

func (r SuperBotFightModeLikelyConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperBotFightModeLikelyConfigurationParam) implementsBotManagementUpdateParamsBodyUnion() {}

// A read-only field that shows which unauthorized settings are currently active on
// the zone. These settings typically result from upgrades or downgrades.
type SuperBotFightModeLikelyConfigurationStaleZoneConfigurationParam struct {
	// Indicates that the zone's Bot Fight Mode is turned on.
	FightMode param.Field[bool] `json:"fight_mode"`
}

func (r SuperBotFightModeLikelyConfigurationStaleZoneConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BotManagementUpdateResponse struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection BotManagementUpdateResponseAIBotsProtection `json:"ai_bots_protection"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection BotManagementUpdateResponseCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged bool `json:"is_robots_txt_managed"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated BotManagementUpdateResponseSBFMDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SBFMLikelyAutomated BotManagementUpdateResponseSBFMLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots BotManagementUpdateResponseSBFMVerifiedBots `json:"sbfm_verified_bots"`
	// This field can have the runtime type of
	// [BotFightModeConfigurationStaleZoneConfiguration],
	// [SuperBotFightModeDefinitelyConfigurationStaleZoneConfiguration],
	// [SuperBotFightModeLikelyConfigurationStaleZoneConfiguration],
	// [SubscriptionConfigurationStaleZoneConfiguration].
	StaleZoneConfiguration interface{} `json:"stale_zone_configuration"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool `json:"suppress_session_score"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                            `json:"using_latest_model"`
	JSON             botManagementUpdateResponseJSON `json:"-"`
	union            BotManagementUpdateResponseUnion
}

// botManagementUpdateResponseJSON contains the JSON metadata for the struct
// [BotManagementUpdateResponse]
type botManagementUpdateResponseJSON struct {
	AIBotsProtection             apijson.Field
	AutoUpdateModel              apijson.Field
	CrawlerProtection            apijson.Field
	EnableJS                     apijson.Field
	FightMode                    apijson.Field
	IsRobotsTXTManaged           apijson.Field
	OptimizeWordpress            apijson.Field
	SBFMDefinitelyAutomated      apijson.Field
	SBFMLikelyAutomated          apijson.Field
	SBFMStaticResourceProtection apijson.Field
	SBFMVerifiedBots             apijson.Field
	StaleZoneConfiguration       apijson.Field
	SuppressSessionScore         apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r botManagementUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BotManagementUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BotManagementUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BotManagementUpdateResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [BotFightModeConfiguration],
// [SuperBotFightModeDefinitelyConfiguration],
// [SuperBotFightModeLikelyConfiguration], [SubscriptionConfiguration].
func (r BotManagementUpdateResponse) AsUnion() BotManagementUpdateResponseUnion {
	return r.union
}

// Union satisfied by [BotFightModeConfiguration],
// [SuperBotFightModeDefinitelyConfiguration],
// [SuperBotFightModeLikelyConfiguration] or [SubscriptionConfiguration].
type BotManagementUpdateResponseUnion interface {
	implementsBotManagementUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BotManagementUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BotFightModeConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperBotFightModeDefinitelyConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperBotFightModeLikelyConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionConfiguration{}),
		},
	)
}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type BotManagementUpdateResponseAIBotsProtection string

const (
	BotManagementUpdateResponseAIBotsProtectionBlock         BotManagementUpdateResponseAIBotsProtection = "block"
	BotManagementUpdateResponseAIBotsProtectionDisabled      BotManagementUpdateResponseAIBotsProtection = "disabled"
	BotManagementUpdateResponseAIBotsProtectionOnlyOnADPages BotManagementUpdateResponseAIBotsProtection = "only_on_ad_pages"
)

func (r BotManagementUpdateResponseAIBotsProtection) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseAIBotsProtectionBlock, BotManagementUpdateResponseAIBotsProtectionDisabled, BotManagementUpdateResponseAIBotsProtectionOnlyOnADPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type BotManagementUpdateResponseCrawlerProtection string

const (
	BotManagementUpdateResponseCrawlerProtectionEnabled  BotManagementUpdateResponseCrawlerProtection = "enabled"
	BotManagementUpdateResponseCrawlerProtectionDisabled BotManagementUpdateResponseCrawlerProtection = "disabled"
)

func (r BotManagementUpdateResponseCrawlerProtection) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseCrawlerProtectionEnabled, BotManagementUpdateResponseCrawlerProtectionDisabled:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateResponseSBFMDefinitelyAutomated string

const (
	BotManagementUpdateResponseSBFMDefinitelyAutomatedAllow            BotManagementUpdateResponseSBFMDefinitelyAutomated = "allow"
	BotManagementUpdateResponseSBFMDefinitelyAutomatedBlock            BotManagementUpdateResponseSBFMDefinitelyAutomated = "block"
	BotManagementUpdateResponseSBFMDefinitelyAutomatedManagedChallenge BotManagementUpdateResponseSBFMDefinitelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateResponseSBFMDefinitelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseSBFMDefinitelyAutomatedAllow, BotManagementUpdateResponseSBFMDefinitelyAutomatedBlock, BotManagementUpdateResponseSBFMDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementUpdateResponseSBFMLikelyAutomated string

const (
	BotManagementUpdateResponseSBFMLikelyAutomatedAllow            BotManagementUpdateResponseSBFMLikelyAutomated = "allow"
	BotManagementUpdateResponseSBFMLikelyAutomatedBlock            BotManagementUpdateResponseSBFMLikelyAutomated = "block"
	BotManagementUpdateResponseSBFMLikelyAutomatedManagedChallenge BotManagementUpdateResponseSBFMLikelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateResponseSBFMLikelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseSBFMLikelyAutomatedAllow, BotManagementUpdateResponseSBFMLikelyAutomatedBlock, BotManagementUpdateResponseSBFMLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateResponseSBFMVerifiedBots string

const (
	BotManagementUpdateResponseSBFMVerifiedBotsAllow BotManagementUpdateResponseSBFMVerifiedBots = "allow"
	BotManagementUpdateResponseSBFMVerifiedBotsBlock BotManagementUpdateResponseSBFMVerifiedBots = "block"
)

func (r BotManagementUpdateResponseSBFMVerifiedBots) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseSBFMVerifiedBotsAllow, BotManagementUpdateResponseSBFMVerifiedBotsBlock:
		return true
	}
	return false
}

type BotManagementGetResponse struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection BotManagementGetResponseAIBotsProtection `json:"ai_bots_protection"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel bool `json:"auto_update_model"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection BotManagementGetResponseCrawlerProtection `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS bool `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode bool `json:"fight_mode"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged bool `json:"is_robots_txt_managed"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress bool `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated BotManagementGetResponseSBFMDefinitelyAutomated `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SBFMLikelyAutomated BotManagementGetResponseSBFMLikelyAutomated `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection bool `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots BotManagementGetResponseSBFMVerifiedBots `json:"sbfm_verified_bots"`
	// This field can have the runtime type of
	// [BotFightModeConfigurationStaleZoneConfiguration],
	// [SuperBotFightModeDefinitelyConfigurationStaleZoneConfiguration],
	// [SuperBotFightModeLikelyConfigurationStaleZoneConfiguration],
	// [SubscriptionConfigurationStaleZoneConfiguration].
	StaleZoneConfiguration interface{} `json:"stale_zone_configuration"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore bool `json:"suppress_session_score"`
	// A read-only field that indicates whether the zone currently is running the
	// latest ML model.
	UsingLatestModel bool                         `json:"using_latest_model"`
	JSON             botManagementGetResponseJSON `json:"-"`
	union            BotManagementGetResponseUnion
}

// botManagementGetResponseJSON contains the JSON metadata for the struct
// [BotManagementGetResponse]
type botManagementGetResponseJSON struct {
	AIBotsProtection             apijson.Field
	AutoUpdateModel              apijson.Field
	CrawlerProtection            apijson.Field
	EnableJS                     apijson.Field
	FightMode                    apijson.Field
	IsRobotsTXTManaged           apijson.Field
	OptimizeWordpress            apijson.Field
	SBFMDefinitelyAutomated      apijson.Field
	SBFMLikelyAutomated          apijson.Field
	SBFMStaticResourceProtection apijson.Field
	SBFMVerifiedBots             apijson.Field
	StaleZoneConfiguration       apijson.Field
	SuppressSessionScore         apijson.Field
	UsingLatestModel             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r botManagementGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BotManagementGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BotManagementGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BotManagementGetResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [BotFightModeConfiguration],
// [SuperBotFightModeDefinitelyConfiguration],
// [SuperBotFightModeLikelyConfiguration], [SubscriptionConfiguration].
func (r BotManagementGetResponse) AsUnion() BotManagementGetResponseUnion {
	return r.union
}

// Union satisfied by [BotFightModeConfiguration],
// [SuperBotFightModeDefinitelyConfiguration],
// [SuperBotFightModeLikelyConfiguration] or [SubscriptionConfiguration].
type BotManagementGetResponseUnion interface {
	implementsBotManagementGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BotManagementGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BotFightModeConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperBotFightModeDefinitelyConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperBotFightModeLikelyConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionConfiguration{}),
		},
	)
}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type BotManagementGetResponseAIBotsProtection string

const (
	BotManagementGetResponseAIBotsProtectionBlock         BotManagementGetResponseAIBotsProtection = "block"
	BotManagementGetResponseAIBotsProtectionDisabled      BotManagementGetResponseAIBotsProtection = "disabled"
	BotManagementGetResponseAIBotsProtectionOnlyOnADPages BotManagementGetResponseAIBotsProtection = "only_on_ad_pages"
)

func (r BotManagementGetResponseAIBotsProtection) IsKnown() bool {
	switch r {
	case BotManagementGetResponseAIBotsProtectionBlock, BotManagementGetResponseAIBotsProtectionDisabled, BotManagementGetResponseAIBotsProtectionOnlyOnADPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type BotManagementGetResponseCrawlerProtection string

const (
	BotManagementGetResponseCrawlerProtectionEnabled  BotManagementGetResponseCrawlerProtection = "enabled"
	BotManagementGetResponseCrawlerProtectionDisabled BotManagementGetResponseCrawlerProtection = "disabled"
)

func (r BotManagementGetResponseCrawlerProtection) IsKnown() bool {
	switch r {
	case BotManagementGetResponseCrawlerProtectionEnabled, BotManagementGetResponseCrawlerProtectionDisabled:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementGetResponseSBFMDefinitelyAutomated string

const (
	BotManagementGetResponseSBFMDefinitelyAutomatedAllow            BotManagementGetResponseSBFMDefinitelyAutomated = "allow"
	BotManagementGetResponseSBFMDefinitelyAutomatedBlock            BotManagementGetResponseSBFMDefinitelyAutomated = "block"
	BotManagementGetResponseSBFMDefinitelyAutomatedManagedChallenge BotManagementGetResponseSBFMDefinitelyAutomated = "managed_challenge"
)

func (r BotManagementGetResponseSBFMDefinitelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementGetResponseSBFMDefinitelyAutomatedAllow, BotManagementGetResponseSBFMDefinitelyAutomatedBlock, BotManagementGetResponseSBFMDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementGetResponseSBFMLikelyAutomated string

const (
	BotManagementGetResponseSBFMLikelyAutomatedAllow            BotManagementGetResponseSBFMLikelyAutomated = "allow"
	BotManagementGetResponseSBFMLikelyAutomatedBlock            BotManagementGetResponseSBFMLikelyAutomated = "block"
	BotManagementGetResponseSBFMLikelyAutomatedManagedChallenge BotManagementGetResponseSBFMLikelyAutomated = "managed_challenge"
)

func (r BotManagementGetResponseSBFMLikelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementGetResponseSBFMLikelyAutomatedAllow, BotManagementGetResponseSBFMLikelyAutomatedBlock, BotManagementGetResponseSBFMLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementGetResponseSBFMVerifiedBots string

const (
	BotManagementGetResponseSBFMVerifiedBotsAllow BotManagementGetResponseSBFMVerifiedBots = "allow"
	BotManagementGetResponseSBFMVerifiedBotsBlock BotManagementGetResponseSBFMVerifiedBots = "block"
)

func (r BotManagementGetResponseSBFMVerifiedBots) IsKnown() bool {
	switch r {
	case BotManagementGetResponseSBFMVerifiedBotsAllow, BotManagementGetResponseSBFMVerifiedBotsBlock:
		return true
	}
	return false
}

type BotManagementUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string]                `path:"zone_id,required"`
	Body   BotManagementUpdateParamsBodyUnion `json:"body,required"`
}

func (r BotManagementUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type BotManagementUpdateParamsBody struct {
	// Enable rule to block AI Scrapers and Crawlers. Please note the value
	// `only_on_ad_pages` is currently not available for Enterprise customers.
	AIBotsProtection param.Field[BotManagementUpdateParamsBodyAIBotsProtection] `json:"ai_bots_protection"`
	// Automatically update to the newest bot detection models created by Cloudflare as
	// they are released.
	// [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)
	AutoUpdateModel param.Field[bool] `json:"auto_update_model"`
	// Enable rule to punish AI Scrapers and Crawlers via a link maze.
	CrawlerProtection param.Field[BotManagementUpdateParamsBodyCrawlerProtection] `json:"crawler_protection"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management.
	// [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	EnableJS param.Field[bool] `json:"enable_js"`
	// Whether to enable Bot Fight Mode.
	FightMode param.Field[bool] `json:"fight_mode"`
	// Enable cloudflare managed robots.txt. If an existing robots.txt is detected,
	// then managed robots.txt will be prepended to the existing robots.txt.
	IsRobotsTXTManaged param.Field[bool] `json:"is_robots_txt_managed"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	OptimizeWordpress param.Field[bool] `json:"optimize_wordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	SBFMDefinitelyAutomated param.Field[BotManagementUpdateParamsBodySBFMDefinitelyAutomated] `json:"sbfm_definitely_automated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	SBFMLikelyAutomated param.Field[BotManagementUpdateParamsBodySBFMLikelyAutomated] `json:"sbfm_likely_automated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if
	// static resources on your application need bot protection. Note: Static resource
	// protection can also result in legitimate traffic being blocked.
	SBFMStaticResourceProtection param.Field[bool] `json:"sbfm_static_resource_protection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	SBFMVerifiedBots       param.Field[BotManagementUpdateParamsBodySBFMVerifiedBots] `json:"sbfm_verified_bots"`
	StaleZoneConfiguration param.Field[interface{}]                                   `json:"stale_zone_configuration"`
	// Whether to disable tracking the highest bot score for a session in the Bot
	// Management cookie.
	SuppressSessionScore param.Field[bool] `json:"suppress_session_score"`
}

func (r BotManagementUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BotManagementUpdateParamsBody) implementsBotManagementUpdateParamsBodyUnion() {}

// Satisfied by [bot_management.BotFightModeConfigurationParam],
// [bot_management.SuperBotFightModeDefinitelyConfigurationParam],
// [bot_management.SuperBotFightModeLikelyConfigurationParam],
// [bot_management.SubscriptionConfigurationParam],
// [BotManagementUpdateParamsBody].
type BotManagementUpdateParamsBodyUnion interface {
	implementsBotManagementUpdateParamsBodyUnion()
}

// Enable rule to block AI Scrapers and Crawlers. Please note the value
// `only_on_ad_pages` is currently not available for Enterprise customers.
type BotManagementUpdateParamsBodyAIBotsProtection string

const (
	BotManagementUpdateParamsBodyAIBotsProtectionBlock         BotManagementUpdateParamsBodyAIBotsProtection = "block"
	BotManagementUpdateParamsBodyAIBotsProtectionDisabled      BotManagementUpdateParamsBodyAIBotsProtection = "disabled"
	BotManagementUpdateParamsBodyAIBotsProtectionOnlyOnADPages BotManagementUpdateParamsBodyAIBotsProtection = "only_on_ad_pages"
)

func (r BotManagementUpdateParamsBodyAIBotsProtection) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsBodyAIBotsProtectionBlock, BotManagementUpdateParamsBodyAIBotsProtectionDisabled, BotManagementUpdateParamsBodyAIBotsProtectionOnlyOnADPages:
		return true
	}
	return false
}

// Enable rule to punish AI Scrapers and Crawlers via a link maze.
type BotManagementUpdateParamsBodyCrawlerProtection string

const (
	BotManagementUpdateParamsBodyCrawlerProtectionEnabled  BotManagementUpdateParamsBodyCrawlerProtection = "enabled"
	BotManagementUpdateParamsBodyCrawlerProtectionDisabled BotManagementUpdateParamsBodyCrawlerProtection = "disabled"
)

func (r BotManagementUpdateParamsBodyCrawlerProtection) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsBodyCrawlerProtectionEnabled, BotManagementUpdateParamsBodyCrawlerProtectionDisabled:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
type BotManagementUpdateParamsBodySBFMDefinitelyAutomated string

const (
	BotManagementUpdateParamsBodySBFMDefinitelyAutomatedAllow            BotManagementUpdateParamsBodySBFMDefinitelyAutomated = "allow"
	BotManagementUpdateParamsBodySBFMDefinitelyAutomatedBlock            BotManagementUpdateParamsBodySBFMDefinitelyAutomated = "block"
	BotManagementUpdateParamsBodySBFMDefinitelyAutomatedManagedChallenge BotManagementUpdateParamsBodySBFMDefinitelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateParamsBodySBFMDefinitelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsBodySBFMDefinitelyAutomatedAllow, BotManagementUpdateParamsBodySBFMDefinitelyAutomatedBlock, BotManagementUpdateParamsBodySBFMDefinitelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
type BotManagementUpdateParamsBodySBFMLikelyAutomated string

const (
	BotManagementUpdateParamsBodySBFMLikelyAutomatedAllow            BotManagementUpdateParamsBodySBFMLikelyAutomated = "allow"
	BotManagementUpdateParamsBodySBFMLikelyAutomatedBlock            BotManagementUpdateParamsBodySBFMLikelyAutomated = "block"
	BotManagementUpdateParamsBodySBFMLikelyAutomatedManagedChallenge BotManagementUpdateParamsBodySBFMLikelyAutomated = "managed_challenge"
)

func (r BotManagementUpdateParamsBodySBFMLikelyAutomated) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsBodySBFMLikelyAutomatedAllow, BotManagementUpdateParamsBodySBFMLikelyAutomatedBlock, BotManagementUpdateParamsBodySBFMLikelyAutomatedManagedChallenge:
		return true
	}
	return false
}

// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
type BotManagementUpdateParamsBodySBFMVerifiedBots string

const (
	BotManagementUpdateParamsBodySBFMVerifiedBotsAllow BotManagementUpdateParamsBodySBFMVerifiedBots = "allow"
	BotManagementUpdateParamsBodySBFMVerifiedBotsBlock BotManagementUpdateParamsBodySBFMVerifiedBots = "block"
)

func (r BotManagementUpdateParamsBodySBFMVerifiedBots) IsKnown() bool {
	switch r {
	case BotManagementUpdateParamsBodySBFMVerifiedBotsAllow, BotManagementUpdateParamsBodySBFMVerifiedBotsBlock:
		return true
	}
	return false
}

type BotManagementUpdateResponseEnvelope struct {
	Errors   []BotManagementUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BotManagementUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success BotManagementUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  BotManagementUpdateResponse                `json:"result"`
	JSON    botManagementUpdateResponseEnvelopeJSON    `json:"-"`
}

// botManagementUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [BotManagementUpdateResponseEnvelope]
type botManagementUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BotManagementUpdateResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           BotManagementUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             botManagementUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// botManagementUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BotManagementUpdateResponseEnvelopeErrors]
type botManagementUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotManagementUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BotManagementUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    botManagementUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// botManagementUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [BotManagementUpdateResponseEnvelopeErrorsSource]
type botManagementUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BotManagementUpdateResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           BotManagementUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             botManagementUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// botManagementUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BotManagementUpdateResponseEnvelopeMessages]
type botManagementUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotManagementUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BotManagementUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    botManagementUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// botManagementUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [BotManagementUpdateResponseEnvelopeMessagesSource]
type botManagementUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BotManagementUpdateResponseEnvelopeSuccess bool

const (
	BotManagementUpdateResponseEnvelopeSuccessTrue BotManagementUpdateResponseEnvelopeSuccess = true
)

func (r BotManagementUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BotManagementUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BotManagementGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type BotManagementGetResponseEnvelope struct {
	Errors   []BotManagementGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BotManagementGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success BotManagementGetResponseEnvelopeSuccess `json:"success,required"`
	Result  BotManagementGetResponse                `json:"result"`
	JSON    botManagementGetResponseEnvelopeJSON    `json:"-"`
}

// botManagementGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BotManagementGetResponseEnvelope]
type botManagementGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BotManagementGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           BotManagementGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             botManagementGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// botManagementGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BotManagementGetResponseEnvelopeErrors]
type botManagementGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BotManagementGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    botManagementGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// botManagementGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [BotManagementGetResponseEnvelopeErrorsSource]
type botManagementGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BotManagementGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           BotManagementGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             botManagementGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// botManagementGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BotManagementGetResponseEnvelopeMessages]
type botManagementGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BotManagementGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    botManagementGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// botManagementGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [BotManagementGetResponseEnvelopeMessagesSource]
type botManagementGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotManagementGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botManagementGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BotManagementGetResponseEnvelopeSuccess bool

const (
	BotManagementGetResponseEnvelopeSuccessTrue BotManagementGetResponseEnvelopeSuccess = true
)

func (r BotManagementGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BotManagementGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
