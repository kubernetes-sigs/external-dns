package linodego

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty"
)

// Event represents an action taken on the Account.
type Event struct {
	CreatedStr string `json:"created"`

	// The unique ID of this Event.
	ID int

	// Current status of the Event, Enum: "failed" "finished" "notification" "scheduled" "started"
	Status EventStatus

	// The action that caused this Event. New actions may be added in the future.
	Action EventAction

	// A percentage estimating the amount of time remaining for an Event. Returns null for notification events.
	PercentComplete int `json:"percent_complete"`

	// The rate of completion of the Event. Only some Events will return rate; for example, migration and resize Events.
	Rate string

	// If this Event has been read.
	Read bool

	// If this Event has been seen.
	Seen bool

	// The estimated time remaining until the completion of this Event. This value is only returned for in-progress events.
	TimeRemaining int

	// The username of the User who caused the Event.
	Username string

	// Detailed information about the Event's entity, including ID, type, label, and URL used to access it.
	Entity *EventEntity

	// When this Event was created.
	Created *time.Time `json:"-"`
}

// EventAction constants start with Action and include all known Linode API Event Actions.
type EventAction string

const (
	ActionBackupsEnable            EventAction = "backups_enable"
	ActionBackupsCancel            EventAction = "backups_cancel"
	ActionBackupsRestore           EventAction = "backups_restore"
	ActionCommunityQuestionReply   EventAction = "community_question_reply"
	ActionCreateCardUpdated        EventAction = "credit_card_updated"
	ActionDiskCreate               EventAction = "disk_create"
	ActionDiskDelete               EventAction = "disk_delete"
	ActionDiskDuplicate            EventAction = "disk_duplicate"
	ActionDiskImagize              EventAction = "disk_imagize"
	ActionDiskResize               EventAction = "disk_resize"
	ActionDNSRecordCreate          EventAction = "dns_record_create"
	ActionDNSRecordDelete          EventAction = "dns_record_delete"
	ActionDNSZoneCreate            EventAction = "dns_zone_create"
	ActionDNSZoneDelete            EventAction = "dns_zone_delete"
	ActionImageDelete              EventAction = "image_delete"
	ActionLinodeAddIP              EventAction = "linode_addip"
	ActionLinodeBoot               EventAction = "linode_boot"
	ActionLinodeClone              EventAction = "linode_clone"
	ActionLinodeCreate             EventAction = "linode_create"
	ActionLinodeDelete             EventAction = "linode_delete"
	ActionLinodeDeleteIP           EventAction = "linode_deleteip"
	ActionLinodeMigrate            EventAction = "linode_migrate"
	ActionLinodeMutate             EventAction = "linode_mutate"
	ActionLinodeReboot             EventAction = "linode_reboot"
	ActionLinodeRebuild            EventAction = "linode_rebuild"
	ActionLinodeResize             EventAction = "linode_resize"
	ActionLinodeShutdown           EventAction = "linode_shutdown"
	ActionLinodeSnapshot           EventAction = "linode_snapshot"
	ActionLongviewClientCreate     EventAction = "longviewclient_create"
	ActionLongviewClientDelete     EventAction = "longviewclient_delete"
	ActionManagedDisabled          EventAction = "managed_disabled"
	ActionManagedEnabled           EventAction = "managed_enabled"
	ActionManagedServiceCreate     EventAction = "managed_service_create"
	ActionManagedServiceDelete     EventAction = "managed_service_delete"
	ActionNodebalancerCreate       EventAction = "nodebalancer_create"
	ActionNodebalancerDelete       EventAction = "nodebalancer_delete"
	ActionNodebalancerConfigCreate EventAction = "nodebalancer_config_create"
	ActionNodebalancerConfigDelete EventAction = "nodebalancer_config_delete"
	ActionPasswordReset            EventAction = "password_reset"
	ActionPaymentSubmitted         EventAction = "payment_submitted"
	ActionStackScriptCreate        EventAction = "stackscript_create"
	ActionStackScriptDelete        EventAction = "stackscript_delete"
	ActionStackScriptPublicize     EventAction = "stackscript_publicize"
	ActionStackScriptRevise        EventAction = "stackscript_revise"
	ActionTFADisabled              EventAction = "tfa_disabled"
	ActionTFAEnabled               EventAction = "tfa_enabled"
	ActionTicketAttachmentUpload   EventAction = "ticket_attachment_upload"
	ActionTicketCreate             EventAction = "ticket_create"
	ActionTicketReply              EventAction = "ticket_reply"
	ActionVolumeAttach             EventAction = "volume_attach"
	ActionVolumeClone              EventAction = "volume_clone"
	ActionVolumeCreate             EventAction = "volume_create"
	ActionVolumeDelte              EventAction = "volume_delete"
	ActionVolumeDetach             EventAction = "volume_detach"
	ActionVolumeResize             EventAction = "volume_resize"
)

// EntityType constants start with Entity and include Linode API Event Entity Types
type EntityType string

const (
	EntityLinode EntityType = "linode"
	EntityDisk   EntityType = "disk"
)

// EventStatus constants start with Event and include Linode API Event Status values
type EventStatus string

const (
	EventFailed       EventStatus = "failed"
	EventFinished     EventStatus = "finished"
	EventNotification EventStatus = "notification"
	EventScheduled    EventStatus = "scheduled"
	EventStarted      EventStatus = "started"
)

// EventEntity provides detailed information about the Event's
// associated entity, including ID, Type, Label, and a URL that
// can be used to access it.
type EventEntity struct {
	// ID may be a string or int, it depends on the EntityType
	ID    interface{}
	Label string
	Type  EntityType
	URL   string
}

// EventsPagedResponse represents a paginated Events API response
type EventsPagedResponse struct {
	*PageOptions
	Data []*Event
}

// endpoint gets the endpoint URL for Event
func (EventsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.Events.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// endpointWithID gets the endpoint URL for a specific Event
func (e Event) endpointWithID(c *Client) string {
	endpoint, err := c.Events.Endpoint()
	if err != nil {
		panic(err)
	}
	endpoint = fmt.Sprintf("%s/%d", endpoint, e.ID)
	return endpoint
}

// appendData appends Events when processing paginated Event responses
func (resp *EventsPagedResponse) appendData(r *EventsPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

// setResult sets the Resty response type of Events
func (EventsPagedResponse) setResult(r *resty.Request) {
	r.SetResult(EventsPagedResponse{})
}

// ListEvents gets a collection of Event objects representing actions taken
// on the Account. The Events returned depend on the token grants and the grants
// of the associated user.
func (c *Client) ListEvents(ctx context.Context, opts *ListOptions) ([]*Event, error) {
	response := EventsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	for _, el := range response.Data {
		el.fixDates()
	}
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// GetEvent gets the Event with the Event ID
func (c *Client) GetEvent(ctx context.Context, id int) (*Event, error) {
	e, err := c.Events.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := c.R(ctx).SetResult(&Event{}).Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*Event).fixDates(), nil
}

// fixDates converts JSON timestamps to Go time.Time values
func (v *Event) fixDates() *Event {
	v.Created, _ = parseDates(v.CreatedStr)
	return v
}

// MarkEventRead marks a single Event as read.
func (c *Client) MarkEventRead(ctx context.Context, event *Event) error {
	e := event.endpointWithID(c)
	e = fmt.Sprintf("%s/read", e)

	if _, err := coupleAPIErrors(c.R(ctx).Post(e)); err != nil {
		return err
	}

	return nil
}

// MarkEventsSeen marks all Events up to and including this Event by ID as seen.
func (c *Client) MarkEventsSeen(ctx context.Context, event *Event) error {
	e := event.endpointWithID(c)
	e = fmt.Sprintf("%s/seen", e)

	if _, err := coupleAPIErrors(c.R(ctx).Post(e)); err != nil {
		return err
	}

	return nil
}
