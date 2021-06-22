package egoscale

import (
	"context"
	"fmt"
	"time"
)

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// RunstatusEvent is a runstatus event
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//RunstatusEvent is a runstatus event
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
//RunstatusEvent is a runstatus event
=======
// RunstatusEvent is a runstatus event
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//RunstatusEvent is a runstatus event
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//RunstatusEvent is a runstatus event
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//RunstatusEvent is a runstatus event
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
type RunstatusEvent struct {
	Created *time.Time `json:"created,omitempty"`
	State   string     `json:"state,omitempty"`
	Status  string     `json:"status"`
	Text    string     `json:"text"`
}

// UpdateRunstatusIncident create runstatus incident event
// Events can be updates or final message with status completed.
func (client *Client) UpdateRunstatusIncident(ctx context.Context, incident RunstatusIncident, event RunstatusEvent) error {
	if incident.EventsURL == "" {
		return fmt.Errorf("empty Events URL for %#v", incident)
	}

	_, err := client.runstatusRequest(ctx, incident.EventsURL, event, "POST")
	return err
}

// UpdateRunstatusMaintenance adds a event to a maintenance.
// Events can be updates or final message with status completed.
func (client *Client) UpdateRunstatusMaintenance(ctx context.Context, maintenance RunstatusMaintenance, event RunstatusEvent) error {
	if maintenance.EventsURL == "" {
		return fmt.Errorf("empty Events URL for %#v", maintenance)
	}

	_, err := client.runstatusRequest(ctx, maintenance.EventsURL, event, "POST")
	return err
}
