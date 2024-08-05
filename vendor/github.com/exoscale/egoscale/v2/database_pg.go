package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// StopPgDatabaseMigration stops running Database migration.
func (c *Client) StopPgDatabaseMigration(ctx context.Context, zone string, name string) error {
	resp, err := c.StopDbaasPgMigrationWithResponse(apiv2.WithZone(ctx, zone), oapi.DbaasServiceName(name))
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
