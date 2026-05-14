// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package abuse_reports

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AbuseReportService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAbuseReportService] method instead.
type AbuseReportService struct {
	Options []option.RequestOption
}

// NewAbuseReportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAbuseReportService(opts ...option.RequestOption) (r *AbuseReportService) {
	r = &AbuseReportService{}
	r.Options = opts
	return
}

// Submit the Abuse Report of a particular type
func (r *AbuseReportService) New(ctx context.Context, reportType string, params AbuseReportNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env AbuseReportNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if reportType == "" {
		err = errors.New("missing required report_type parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/abuse-reports/%s", params.AccountID, reportType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AbuseReportNewParams struct {
	AccountID param.Field[string]           `path:"account_id,required"`
	Body      AbuseReportNewParamsBodyUnion `json:"body,required"`
}

func (r AbuseReportNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AbuseReportNewParamsBody struct {
	// The abuse report type.
	Act param.Field[AbuseReportNewParamsBodyAct] `json:"act,required"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email,required"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyOwnerNotification] `json:"owner_notification,required"`
	// A list of valid URLs separated by ‘\n’ (new line character). The list of the
	// URLs should not exceed 250 URLs. All URLs should have the same hostname. Each
	// URL should be unique. This field may be released by Cloudflare to third parties
	// such as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls,required"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAgree] `json:"agree"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// A list of IP addresses separated by ‘\n’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique.
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyHostNotification] `json:"host_notification"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters.
	Justification param.Field[string] `json:"justification"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyNcmecNotification] `json:"ncmec_notification"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	OriginalWork param.Field[string] `json:"original_work"`
	// A comma separated list of ports and protocols e.g. 80/TCP, 22/UDP. The total
	// size of the field should not exceed 2000 characters. Each individual
	// port/protocol should not exceed 100 characters. The list should not have more
	// than 30 unique ports and protocols.
	PortsProtocols param.Field[string] `json:"ports_protocols"`
	// Text containing 2 characters
	ReportedCountry param.Field[string] `json:"reported_country"`
	// Text not exceeding 255 characters
	ReportedUserAgent param.Field[string] `json:"reported_user_agent"`
	// Required for DMCA reports, should be same as Name. An affirmation that all
	// information in the report is true and accurate while agreeing to the policies of
	// Cloudflare's abuse reports
	Signature param.Field[string] `json:"signature"`
	// A list of IP addresses separated by ‘\n’ (new line character). The list of
	// source IPs should not exceed 30 IP addresses. Each one of the IP addresses ought
	// to be unique.
	SourceIPs param.Field[string] `json:"source_ips"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol"`
}

func (r AbuseReportNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBody) implementsAbuseReportNewParamsBodyUnion() {}

// Satisfied by [abuse_reports.AbuseReportNewParamsBodyAbuseReportsDmcaReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsTrademarkReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsGeneralReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsPhishingReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsCsamReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsThreatReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReport],
// [abuse_reports.AbuseReportNewParamsBodyAbuseReportsNcseiReport],
// [AbuseReportNewParamsBody].
type AbuseReportNewParamsBodyUnion interface {
	implementsAbuseReportNewParamsBodyUnion()
}

type AbuseReportNewParamsBodyAbuseReportsDmcaReport struct {
	// The abuse report type.
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsDmcaReportAct] `json:"act,required"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Address1 param.Field[string] `json:"address1,required"`
	// The name of the copyright holder. Text not exceeding 60 characters. This field
	// may be released by Cloudflare to third parties such as the Lumen Database
	// (https://lumendatabase.org/).
	AgentName param.Field[string] `json:"agent_name,required"`
	// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
	Agree param.Field[AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree] `json:"agree,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	City param.Field[string] `json:"city,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country,required"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email,required"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotification] `json:"host_notification,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	OriginalWork param.Field[string] `json:"original_work,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotification] `json:"owner_notification,required"`
	// Required for DMCA reports, should be same as Name. An affirmation that all
	// information in the report is true and accurate while agreeing to the policies of
	// Cloudflare's abuse reports
	Signature param.Field[string] `json:"signature,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	State param.Field[string] `json:"state,required"`
	// A list of valid URLs separated by ‘\n’ (new line character). The list of the
	// URLs should not exceed 250 URLs. All URLs should have the same hostname. Each
	// URL should be unique. This field may be released by Cloudflare to third parties
	// such as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls,required"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text containing 2 characters
	ReportedCountry param.Field[string] `json:"reported_country"`
	// Text not exceeding 255 characters
	ReportedUserAgent param.Field[string] `json:"reported_user_agent"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
}

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReport) implementsAbuseReportNewParamsBodyUnion() {}

// The abuse report type.
type AbuseReportNewParamsBodyAbuseReportsDmcaReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseDmca AbuseReportNewParamsBodyAbuseReportsDmcaReportAct = "abuse_dmca"
)

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsDmcaReportActAbuseDmca:
		return true
	}
	return false
}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree int64

const (
	AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree1 AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree = 1
)

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsDmcaReportAgree1:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotificationSend AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotification = "send"
)

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsDmcaReportHostNotificationSend:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotificationSend AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotification = "send"
)

func (r AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsDmcaReportOwnerNotificationSend:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsTrademarkReport struct {
	// The abuse report type.
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct] `json:"act,required"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email,required"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotification] `json:"host_notification,required"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters.
	Justification param.Field[string] `json:"justification,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotification] `json:"owner_notification,required"`
	// Text not exceeding 1000 characters
	TrademarkNumber param.Field[string] `json:"trademark_number,required"`
	// Text not exceeding 1000 characters
	TrademarkOffice param.Field[string] `json:"trademark_office,required"`
	// Text not exceeding 1000 characters
	TrademarkSymbol param.Field[string] `json:"trademark_symbol,required"`
	// A list of valid URLs separated by ‘\n’ (new line character). The list of the
	// URLs should not exceed 250 URLs. All URLs should have the same hostname. Each
	// URL should be unique. This field may be released by Cloudflare to third parties
	// such as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls,required"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text containing 2 characters
	ReportedCountry param.Field[string] `json:"reported_country"`
	// Text not exceeding 255 characters
	ReportedUserAgent param.Field[string] `json:"reported_user_agent"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
}

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReport) implementsAbuseReportNewParamsBodyUnion() {
}

// The abuse report type.
type AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseTrademark AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct = "abuse_trademark"
)

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsTrademarkReportActAbuseTrademark:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotificationSend AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotification = "send"
)

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsTrademarkReportHostNotificationSend:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotificationSend AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotification = "send"
)

func (r AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsTrademarkReportOwnerNotificationSend:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsGeneralReport struct {
	// The abuse report type.
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsGeneralReportAct] `json:"act,required"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email,required"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification] `json:"host_notification,required"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters.
	Justification param.Field[string] `json:"justification,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification] `json:"owner_notification,required"`
	// A list of valid URLs separated by ‘\n’ (new line character). The list of the
	// URLs should not exceed 250 URLs. All URLs should have the same hostname. Each
	// URL should be unique. This field may be released by Cloudflare to third parties
	// such as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls,required"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// A list of IP addresses separated by ‘\n’ (new line character). The list of
	// destination IPs should not exceed 30 IP addresses. Each one of the IP addresses
	// ought to be unique.
	DestinationIPs param.Field[string] `json:"destination_ips"`
	// A comma separated list of ports and protocols e.g. 80/TCP, 22/UDP. The total
	// size of the field should not exceed 2000 characters. Each individual
	// port/protocol should not exceed 100 characters. The list should not have more
	// than 30 unique ports and protocols.
	PortsProtocols param.Field[string] `json:"ports_protocols"`
	// Text containing 2 characters
	ReportedCountry param.Field[string] `json:"reported_country"`
	// Text not exceeding 255 characters
	ReportedUserAgent param.Field[string] `json:"reported_user_agent"`
	// A list of IP addresses separated by ‘\n’ (new line character). The list of
	// source IPs should not exceed 30 IP addresses. Each one of the IP addresses ought
	// to be unique.
	SourceIPs param.Field[string] `json:"source_ips"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
}

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReport) implementsAbuseReportNewParamsBodyUnion() {
}

// The abuse report type.
type AbuseReportNewParamsBodyAbuseReportsGeneralReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseGeneral AbuseReportNewParamsBodyAbuseReportsGeneralReportAct = "abuse_general"
)

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsGeneralReportActAbuseGeneral:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification = "send-anon"
)

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsGeneralReportHostNotificationSendAnon:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsGeneralReportOwnerNotificationNone:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsPhishingReport struct {
	// The abuse report type.
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsPhishingReportAct] `json:"act,required"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email,required"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification] `json:"host_notification,required"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters.
	Justification param.Field[string] `json:"justification,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification] `json:"owner_notification,required"`
	// A list of valid URLs separated by ‘\n’ (new line character). The list of the
	// URLs should not exceed 250 URLs. All URLs should have the same hostname. Each
	// URL should be unique. This field may be released by Cloudflare to third parties
	// such as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls,required"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	OriginalWork param.Field[string] `json:"original_work"`
	// Text containing 2 characters
	ReportedCountry param.Field[string] `json:"reported_country"`
	// Text not exceeding 255 characters
	ReportedUserAgent param.Field[string] `json:"reported_user_agent"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
}

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReport) implementsAbuseReportNewParamsBodyUnion() {
}

// The abuse report type.
type AbuseReportNewParamsBodyAbuseReportsPhishingReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbusePhishing AbuseReportNewParamsBodyAbuseReportsPhishingReportAct = "abuse_phishing"
)

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsPhishingReportActAbusePhishing:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification = "send-anon"
)

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsPhishingReportHostNotificationSendAnon:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification = "send-anon"
)

func (r AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsPhishingReportOwnerNotificationSendAnon:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsCsamReport struct {
	// The abuse report type.
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsCsamReportAct] `json:"act,required"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email,required"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsCsamReportHostNotification] `json:"host_notification,required"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters.
	Justification param.Field[string] `json:"justification,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	NcmecNotification param.Field[AbuseReportNewParamsBodyAbuseReportsCsamReportNcmecNotification] `json:"ncmec_notification,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotification] `json:"owner_notification,required"`
	// A list of valid URLs separated by ‘\n’ (new line character). The list of the
	// URLs should not exceed 250 URLs. All URLs should have the same hostname. Each
	// URL should be unique. This field may be released by Cloudflare to third parties
	// such as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls,required"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// Text containing 2 characters
	ReportedCountry param.Field[string] `json:"reported_country"`
	// Text not exceeding 255 characters
	ReportedUserAgent param.Field[string] `json:"reported_user_agent"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
}

func (r AbuseReportNewParamsBodyAbuseReportsCsamReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsCsamReport) implementsAbuseReportNewParamsBodyUnion() {}

// The abuse report type.
type AbuseReportNewParamsBodyAbuseReportsCsamReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsCsamReportActAbuseChildren AbuseReportNewParamsBodyAbuseReportsCsamReportAct = "abuse_children"
)

func (r AbuseReportNewParamsBodyAbuseReportsCsamReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsCsamReportActAbuseChildren:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsCsamReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsCsamReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsCsamReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsCsamReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsCsamReportHostNotification = "send-anon"
)

func (r AbuseReportNewParamsBodyAbuseReportsCsamReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsCsamReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsCsamReportHostNotificationSendAnon:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsCsamReportNcmecNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsCsamReportNcmecNotificationSend     AbuseReportNewParamsBodyAbuseReportsCsamReportNcmecNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsCsamReportNcmecNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsCsamReportNcmecNotification = "send-anon"
)

func (r AbuseReportNewParamsBodyAbuseReportsCsamReportNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsCsamReportNcmecNotificationSend, AbuseReportNewParamsBodyAbuseReportsCsamReportNcmecNotificationSendAnon:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsCsamReportOwnerNotificationNone:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsThreatReport struct {
	// The abuse report type.
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsThreatReportAct] `json:"act,required"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email,required"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification] `json:"host_notification,required"`
	// A detailed description of the infringement, including any necessary access
	// details and the exact steps needed to view the content, not exceeding 5000
	// characters.
	Justification param.Field[string] `json:"justification,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification] `json:"owner_notification,required"`
	// A list of valid URLs separated by ‘\n’ (new line character). The list of the
	// URLs should not exceed 250 URLs. All URLs should have the same hostname. Each
	// URL should be unique. This field may be released by Cloudflare to third parties
	// such as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls,required"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text containing 2 characters
	ReportedCountry param.Field[string] `json:"reported_country"`
	// Text not exceeding 255 characters
	ReportedUserAgent param.Field[string] `json:"reported_user_agent"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
}

func (r AbuseReportNewParamsBodyAbuseReportsThreatReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsThreatReport) implementsAbuseReportNewParamsBodyUnion() {}

// The abuse report type.
type AbuseReportNewParamsBodyAbuseReportsThreatReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseThreat AbuseReportNewParamsBodyAbuseReportsThreatReportAct = "abuse_threat"
)

func (r AbuseReportNewParamsBodyAbuseReportsThreatReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsThreatReportActAbuseThreat:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification = "send-anon"
)

func (r AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsThreatReportHostNotificationSendAnon:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification = "send-anon"
)

func (r AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsThreatReportOwnerNotificationSendAnon:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReport struct {
	// The abuse report type.
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct] `json:"act,required"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email,required"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification] `json:"owner_notification,required"`
	// A list of valid URLs separated by ‘\n’ (new line character). The list of the
	// URLs should not exceed 250 URLs. All URLs should have the same hostname. Each
	// URL should be unique. This field may be released by Cloudflare to third parties
	// such as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls,required"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text containing 2 characters
	ReportedCountry param.Field[string] `json:"reported_country"`
	// Text not exceeding 255 characters
	ReportedUserAgent param.Field[string] `json:"reported_user_agent"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
}

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReport) implementsAbuseReportNewParamsBodyUnion() {
}

// The abuse report type.
type AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseRegistrarWhois AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct = "abuse_registrar_whois"
)

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportActAbuseRegistrarWhois:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsRegistrarWhoisReportOwnerNotificationNone:
		return true
	}
	return false
}

type AbuseReportNewParamsBodyAbuseReportsNcseiReport struct {
	// The abuse report type.
	Act param.Field[AbuseReportNewParamsBodyAbuseReportsNcseiReportAct] `json:"act,required"`
	// A valid email of the abuse reporter. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Email param.Field[string] `json:"email,required"`
	// Should match the value provided in `email`
	Email2 param.Field[string] `json:"email2,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	HostNotification param.Field[AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification] `json:"host_notification,required"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Name param.Field[string] `json:"name,required"`
	// If the submitter is the target of NCSEI in the URLs of the abuse report.
	NcseiSubjectRepresentation param.Field[bool] `json:"ncsei_subject_representation,required"`
	// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
	// reports cannot be anonymous.
	OwnerNotification param.Field[AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification] `json:"owner_notification,required"`
	// A list of valid URLs separated by ‘\n’ (new line character). The list of the
	// URLs should not exceed 250 URLs. All URLs should have the same hostname. Each
	// URL should be unique. This field may be released by Cloudflare to third parties
	// such as the Lumen Database (https://lumendatabase.org/).
	URLs param.Field[string] `json:"urls,required"`
	// Any additional comments about the infringement not exceeding 2000 characters
	Comments param.Field[string] `json:"comments"`
	// Text not exceeding 100 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Company param.Field[string] `json:"company"`
	// Text not exceeding 255 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Country param.Field[string] `json:"country"`
	// Text containing 2 characters
	ReportedCountry param.Field[string] `json:"reported_country"`
	// Text not exceeding 255 characters
	ReportedUserAgent param.Field[string] `json:"reported_user_agent"`
	// Text not exceeding 20 characters. This field may be released by Cloudflare to
	// third parties such as the Lumen Database (https://lumendatabase.org/).
	Tele param.Field[string] `json:"tele"`
	// Text not exceeding 255 characters
	Title param.Field[string] `json:"title"`
}

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReport) implementsAbuseReportNewParamsBodyUnion() {}

// The abuse report type.
type AbuseReportNewParamsBodyAbuseReportsNcseiReportAct string

const (
	AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseNcsei AbuseReportNewParamsBodyAbuseReportsNcseiReportAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReportAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsNcseiReportActAbuseNcsei:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotificationSend     AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification = "send-anon"
)

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotificationSend, AbuseReportNewParamsBodyAbuseReportsNcseiReportHostNotificationSendAnon:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification string

const (
	AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationSend     AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification = "send"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationSendAnon AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationNone     AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationSend, AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationSendAnon, AbuseReportNewParamsBodyAbuseReportsNcseiReportOwnerNotificationNone:
		return true
	}
	return false
}

// The abuse report type.
type AbuseReportNewParamsBodyAct string

const (
	AbuseReportNewParamsBodyActAbuseDmca           AbuseReportNewParamsBodyAct = "abuse_dmca"
	AbuseReportNewParamsBodyActAbuseTrademark      AbuseReportNewParamsBodyAct = "abuse_trademark"
	AbuseReportNewParamsBodyActAbuseGeneral        AbuseReportNewParamsBodyAct = "abuse_general"
	AbuseReportNewParamsBodyActAbusePhishing       AbuseReportNewParamsBodyAct = "abuse_phishing"
	AbuseReportNewParamsBodyActAbuseChildren       AbuseReportNewParamsBodyAct = "abuse_children"
	AbuseReportNewParamsBodyActAbuseThreat         AbuseReportNewParamsBodyAct = "abuse_threat"
	AbuseReportNewParamsBodyActAbuseRegistrarWhois AbuseReportNewParamsBodyAct = "abuse_registrar_whois"
	AbuseReportNewParamsBodyActAbuseNcsei          AbuseReportNewParamsBodyAct = "abuse_ncsei"
)

func (r AbuseReportNewParamsBodyAct) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyActAbuseDmca, AbuseReportNewParamsBodyActAbuseTrademark, AbuseReportNewParamsBodyActAbuseGeneral, AbuseReportNewParamsBodyActAbusePhishing, AbuseReportNewParamsBodyActAbuseChildren, AbuseReportNewParamsBodyActAbuseThreat, AbuseReportNewParamsBodyActAbuseRegistrarWhois, AbuseReportNewParamsBodyActAbuseNcsei:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyOwnerNotification string

const (
	AbuseReportNewParamsBodyOwnerNotificationSend     AbuseReportNewParamsBodyOwnerNotification = "send"
	AbuseReportNewParamsBodyOwnerNotificationSendAnon AbuseReportNewParamsBodyOwnerNotification = "send-anon"
	AbuseReportNewParamsBodyOwnerNotificationNone     AbuseReportNewParamsBodyOwnerNotification = "none"
)

func (r AbuseReportNewParamsBodyOwnerNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyOwnerNotificationSend, AbuseReportNewParamsBodyOwnerNotificationSendAnon, AbuseReportNewParamsBodyOwnerNotificationNone:
		return true
	}
	return false
}

// Can be `0` for false or `1` for true. Must be value: 1 for DMCA reports
type AbuseReportNewParamsBodyAgree int64

const (
	AbuseReportNewParamsBodyAgree1 AbuseReportNewParamsBodyAgree = 1
)

func (r AbuseReportNewParamsBodyAgree) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyAgree1:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyHostNotification string

const (
	AbuseReportNewParamsBodyHostNotificationSend     AbuseReportNewParamsBodyHostNotification = "send"
	AbuseReportNewParamsBodyHostNotificationSendAnon AbuseReportNewParamsBodyHostNotification = "send-anon"
)

func (r AbuseReportNewParamsBodyHostNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyHostNotificationSend, AbuseReportNewParamsBodyHostNotificationSendAnon:
		return true
	}
	return false
}

// Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark
// reports cannot be anonymous.
type AbuseReportNewParamsBodyNcmecNotification string

const (
	AbuseReportNewParamsBodyNcmecNotificationSend     AbuseReportNewParamsBodyNcmecNotification = "send"
	AbuseReportNewParamsBodyNcmecNotificationSendAnon AbuseReportNewParamsBodyNcmecNotification = "send-anon"
)

func (r AbuseReportNewParamsBodyNcmecNotification) IsKnown() bool {
	switch r {
	case AbuseReportNewParamsBodyNcmecNotificationSend, AbuseReportNewParamsBodyNcmecNotificationSendAnon:
		return true
	}
	return false
}

type AbuseReportNewResponseEnvelope struct {
	// The identifier for the submitted abuse report.
	AbuseRand string                                `json:"abuse_rand,required"`
	Request   AbuseReportNewResponseEnvelopeRequest `json:"request,required"`
	// The result should be 'success' for successful response
	Result string                             `json:"result,required"`
	JSON   abuseReportNewResponseEnvelopeJSON `json:"-"`
}

// abuseReportNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AbuseReportNewResponseEnvelope]
type abuseReportNewResponseEnvelopeJSON struct {
	AbuseRand   apijson.Field
	Request     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AbuseReportNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r abuseReportNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AbuseReportNewResponseEnvelopeRequest struct {
	// The abuse report type.
	Act  string                                    `json:"act,required"`
	JSON abuseReportNewResponseEnvelopeRequestJSON `json:"-"`
}

// abuseReportNewResponseEnvelopeRequestJSON contains the JSON metadata for the
// struct [AbuseReportNewResponseEnvelopeRequest]
type abuseReportNewResponseEnvelopeRequestJSON struct {
	Act         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AbuseReportNewResponseEnvelopeRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r abuseReportNewResponseEnvelopeRequestJSON) RawJSON() string {
	return r.raw
}
