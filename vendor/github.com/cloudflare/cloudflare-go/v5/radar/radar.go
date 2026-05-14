// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// RadarService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarService] method instead.
type RadarService struct {
	Options           []option.RequestOption
	AI                *AIService
	Ct                *CtService
	Annotations       *AnnotationService
	BGP               *BGPService
	Bots              *BotService
	Datasets          *DatasetService
	DNS               *DNSService
	Netflows          *NetflowService
	Search            *SearchService
	VerifiedBots      *VerifiedBotService
	AS112             *AS112Service
	Email             *EmailService
	Attacks           *AttackService
	Entities          *EntityService
	HTTP              *HTTPService
	Quality           *QualityService
	Ranking           *RankingService
	TrafficAnomalies  *TrafficAnomalyService
	TCPResetsTimeouts *TCPResetsTimeoutService
	RobotsTXT         *RobotsTXTService
	LeakedCredentials *LeakedCredentialService
}

// NewRadarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRadarService(opts ...option.RequestOption) (r *RadarService) {
	r = &RadarService{}
	r.Options = opts
	r.AI = NewAIService(opts...)
	r.Ct = NewCtService(opts...)
	r.Annotations = NewAnnotationService(opts...)
	r.BGP = NewBGPService(opts...)
	r.Bots = NewBotService(opts...)
	r.Datasets = NewDatasetService(opts...)
	r.DNS = NewDNSService(opts...)
	r.Netflows = NewNetflowService(opts...)
	r.Search = NewSearchService(opts...)
	r.VerifiedBots = NewVerifiedBotService(opts...)
	r.AS112 = NewAS112Service(opts...)
	r.Email = NewEmailService(opts...)
	r.Attacks = NewAttackService(opts...)
	r.Entities = NewEntityService(opts...)
	r.HTTP = NewHTTPService(opts...)
	r.Quality = NewQualityService(opts...)
	r.Ranking = NewRankingService(opts...)
	r.TrafficAnomalies = NewTrafficAnomalyService(opts...)
	r.TCPResetsTimeouts = NewTCPResetsTimeoutService(opts...)
	r.RobotsTXT = NewRobotsTXTService(opts...)
	r.LeakedCredentials = NewLeakedCredentialService(opts...)
	return
}
