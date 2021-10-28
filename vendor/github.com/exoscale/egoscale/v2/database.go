package v2

import (
	"context"
	"net/url"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// DatabaseBackupConfig represents a Database Backup configuration.
type DatabaseBackupConfig struct {
	FrequentIntervalMinutes    *int64
	FrequentOldestAgeMinutes   *int64
	InfrequentIntervalMinutes  *int64
	InfrequentOldestAgeMinutes *int64
	Interval                   *int64
	MaxCount                   *int64
	RecoveryMode               *string
}

func databaseBackupConfigFromAPI(c *oapi.DbaasBackupConfig) *DatabaseBackupConfig {
	return &DatabaseBackupConfig{
		FrequentIntervalMinutes:    c.FrequentIntervalMinutes,
		FrequentOldestAgeMinutes:   c.FrequentOldestAgeMinutes,
		InfrequentIntervalMinutes:  c.InfrequentIntervalMinutes,
		InfrequentOldestAgeMinutes: c.InfrequentOldestAgeMinutes,
		Interval:                   c.Interval,
		MaxCount:                   c.MaxCount,
		RecoveryMode:               c.RecoveryMode,
	}
}

// DatabasePlan represents a Database Plan.
type DatabasePlan struct {
	BackupConfig     *DatabaseBackupConfig
	DiskSpace        *int64
	MaxMemoryPercent *int64
	Name             *string
	Nodes            *int64
	NodeCPUs         *int64
	NodeMemory       *int64
}

func databasePlanFromAPI(p *oapi.DbaasPlan) *DatabasePlan {
	return &DatabasePlan{
		BackupConfig:     databaseBackupConfigFromAPI(p.BackupConfig),
		DiskSpace:        p.DiskSpace,
		MaxMemoryPercent: p.MaxMemoryPercent,
		Name:             p.Name,
		Nodes:            p.NodeCount,
		NodeCPUs:         p.NodeCpuCount,
		NodeMemory:       p.NodeMemory,
	}
}

// DatabaseServiceType represents a Database Service type.
type DatabaseServiceType struct {
	DefaultVersion   *string
	Description      *string
	LatestVersion    *string
	Name             *string
	Plans            []*DatabasePlan
	UserConfigSchema map[string]interface{}
}

func databaseServiceTypeFromAPI(t *oapi.DbaasServiceType) *DatabaseServiceType {
	return &DatabaseServiceType{
		DefaultVersion: t.DefaultVersion,
		Description:    t.Description,
		LatestVersion:  t.LatestVersion,
		Name:           (*string)(t.Name),
		Plans: func() []*DatabasePlan {
			plans := make([]*DatabasePlan, 0)
			if t.Plans != nil {
				for _, plan := range *t.Plans {
					plan := plan
					plans = append(plans, databasePlanFromAPI(&plan))
				}
			}
			return plans
		}(),
		UserConfigSchema: func() map[string]interface{} {
			if t.UserConfigSchema != nil {
				return t.UserConfigSchema.AdditionalProperties
			}
			return nil
		}(),
	}
}

// DatabaseServiceBackup represents a Database Service backup.
type DatabaseServiceBackup struct {
	Name *string
	Size *int64
	Date *time.Time
}

func databaseServiceBackupFromAPI(b *oapi.DbaasServiceBackup) *DatabaseServiceBackup {
	return &DatabaseServiceBackup{
		Name: &b.BackupName,
		Size: &b.DataSize,
		Date: &b.BackupTime,
	}
}

// DatabaseServiceComponent represents a Database Service component.
type DatabaseServiceComponent struct {
	Name *string
	Info map[string]interface{}
}

func databaseServiceComponentFromAPI(c *oapi.DbaasServiceComponents) *DatabaseServiceComponent {
	info := map[string]interface{}{
		"host":  c.Host,
		"port":  c.Port,
		"route": c.Route,
		"usage": c.Usage,
	}

	if c.KafkaAuthenticationMethod != nil {
		info["kafka_authenticated_method"] = c.KafkaAuthenticationMethod
	}

	if c.Path != nil {
		info["path"] = c.Path
	}

	if c.Ssl != nil {
		info["ssl"] = c.Ssl
	}

	return &DatabaseServiceComponent{
		Name: &c.Component,
		Info: info,
	}
}

// DatabaseServiceMaintenance represents a Database Service maintenance.
type DatabaseServiceMaintenance struct {
	DOW  string
	Time string
}

func databaseServiceMaintenanceFromAPI(m *oapi.DbaasServiceMaintenance) *DatabaseServiceMaintenance {
	return &DatabaseServiceMaintenance{
		DOW:  string(m.Dow),
		Time: m.Time,
	}
}

// DatabaseServiceUser represents a Database Service user.
type DatabaseServiceUser struct {
	Password *string
	Type     *string
	UserName *string
}

func databaseServiceUserFromAPI(u *oapi.DbaasServiceUser) *DatabaseServiceUser {
	return &DatabaseServiceUser{
		Password: u.Password,
		UserName: &u.Username,
		Type:     &u.Type,
	}
}

// DatabaseService represents a Database Service.
type DatabaseService struct {
	Backups               []*DatabaseServiceBackup
	Components            []*DatabaseServiceComponent
	ConnectionInfo        map[string]interface{}
	CreatedAt             *time.Time
	DiskSize              *int64
	Features              map[string]interface{}
	Maintenance           *DatabaseServiceMaintenance
	Metadata              map[string]interface{}
	Name                  *string
	Nodes                 *int64
	NodeCPUs              *int64
	NodeMemory            *int64
	Plan                  *string
	State                 *string
	TerminationProtection *bool
	Type                  *string
	UpdatedAt             *time.Time
	URI                   *url.URL
	UserConfig            *map[string]interface{}
	Users                 []*DatabaseServiceUser
}

func databaseServiceFromAPI(s *oapi.DbaasService) *DatabaseService {
	return &DatabaseService{
		Backups: func() []*DatabaseServiceBackup {
			backups := make([]*DatabaseServiceBackup, 0)
			if s.Backups != nil {
				for _, b := range *s.Backups {
					backups = append(backups, databaseServiceBackupFromAPI(&b))
				}
			}
			return backups
		}(),
		Components: func() []*DatabaseServiceComponent {
			components := make([]*DatabaseServiceComponent, 0)
			if s.Components != nil {
				for _, c := range *s.Components {
					components = append(components, databaseServiceComponentFromAPI(&c))
				}
			}
			return components
		}(),
		ConnectionInfo: func() (v map[string]interface{}) {
			if s.ConnectionInfo != nil {
				v = s.ConnectionInfo.AdditionalProperties
			}
			return
		}(),
		CreatedAt: s.CreatedAt,
		DiskSize:  s.DiskSize,
		Features: func() (v map[string]interface{}) {
			if s.Features != nil {
				v = s.Features.AdditionalProperties
			}
			return
		}(),
		Maintenance: func() (v *DatabaseServiceMaintenance) {
			if s.Maintenance != nil {
				return databaseServiceMaintenanceFromAPI(s.Maintenance)
			}
			return
		}(),
		Metadata: func() (v map[string]interface{}) {
			if s.Metadata != nil {
				v = s.Metadata.AdditionalProperties
			}
			return
		}(),
		Name:                  (*string)(&s.Name),
		Nodes:                 s.NodeCount,
		NodeCPUs:              s.NodeCpuCount,
		NodeMemory:            s.NodeMemory,
		Plan:                  &s.Plan,
		State:                 (*string)(s.State),
		TerminationProtection: s.TerminationProtection,
		Type:                  (*string)(&s.Type),
		UpdatedAt:             s.UpdatedAt,
		URI: func() *url.URL {
			if s.Uri != nil {
				if u, _ := url.Parse(*s.Uri); u != nil {
					return u
				}
			}
			return nil
		}(),
		UserConfig: func() (v *map[string]interface{}) {
			if s.UserConfig != nil {
				v = &s.UserConfig.AdditionalProperties
			}
			return
		}(),
		Users: func() []*DatabaseServiceUser {
			users := make([]*DatabaseServiceUser, 0)
			if s.Users != nil {
				for _, u := range *s.Users {
					users = append(users, databaseServiceUserFromAPI(&u))
				}
			}
			return users
		}(),
	}
}

// GetDatabaseCACertificate returns the CA certificate required to access Database Services using a TLS connection.
func (c *Client) GetDatabaseCACertificate(ctx context.Context, zone string) (string, error) {
	resp, err := c.GetDbaasCaCertificateWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return "", err
	}

	return *resp.JSON200.Certificate, nil
}

// GetDatabaseServiceType returns the Database Service type corresponding to the specified name.
func (c *Client) GetDatabaseServiceType(ctx context.Context, zone, name string) (*DatabaseServiceType, error) {
	resp, err := c.GetDbaasServiceTypeWithResponse(apiv2.WithZone(ctx, zone), name)
	if err != nil {
		return nil, err
	}

	return databaseServiceTypeFromAPI(resp.JSON200), nil
}

// ListDatabaseServiceTypes returns the list of existing Database Service types.
func (c *Client) ListDatabaseServiceTypes(ctx context.Context, zone string) ([]*DatabaseServiceType, error) {
	list := make([]*DatabaseServiceType, 0)

	resp, err := c.ListDbaasServiceTypesWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.DbaasServiceTypes != nil {
		for i := range *resp.JSON200.DbaasServiceTypes {
			list = append(list, databaseServiceTypeFromAPI(&(*resp.JSON200.DbaasServiceTypes)[i]))
		}
	}

	return list, nil
}

// CreateDatabaseService creates a Database Service.
func (c *Client) CreateDatabaseService(
	ctx context.Context,
	zone string,
	databaseService *DatabaseService,
) (*DatabaseService, error) {
	_, err := c.CreateDbaasServiceWithResponse(
		apiv2.WithZone(ctx, zone),
		oapi.CreateDbaasServiceJSONRequestBody{
			Maintenance: func() (v *struct {
				Dow  oapi.CreateDbaasServiceJSONBodyMaintenanceDow `json:"dow"`
				Time string                                        `json:"time"`
			}) {
				if databaseService.Maintenance != nil {
					v = &struct {
						Dow  oapi.CreateDbaasServiceJSONBodyMaintenanceDow `json:"dow"`
						Time string                                        `json:"time"`
					}{
						Dow:  oapi.CreateDbaasServiceJSONBodyMaintenanceDow(databaseService.Maintenance.DOW),
						Time: databaseService.Maintenance.Time,
					}
				}
				return
			}(),
			Name:                  oapi.DbaasServiceName(*databaseService.Name),
			Plan:                  *databaseService.Plan,
			TerminationProtection: databaseService.TerminationProtection,
			Type:                  oapi.DbaasServiceTypeName(*databaseService.Type),
			UserConfig: func() (v *oapi.CreateDbaasServiceJSONBody_UserConfig) {
				if databaseService.UserConfig != nil {
					v = &oapi.CreateDbaasServiceJSONBody_UserConfig{
						AdditionalProperties: *databaseService.UserConfig,
					}
				}
				return
			}(),
		})
	if err != nil {
		return nil, err
	}

	return c.GetDatabaseService(ctx, zone, *databaseService.Name)
}

// DeleteDatabaseService deletes a Database Service.
func (c *Client) DeleteDatabaseService(ctx context.Context, zone string, databaseService *DatabaseService) error {
	_, err := c.TerminateDbaasServiceWithResponse(apiv2.WithZone(ctx, zone), *databaseService.Name)
	if err != nil {
		return err
	}

	return nil
}

// GetDatabaseService returns the Database Service corresponding to the specified name.
func (c *Client) GetDatabaseService(ctx context.Context, zone, name string) (*DatabaseService, error) {
	resp, err := c.GetDbaasServiceWithResponse(apiv2.WithZone(ctx, zone), name)
	if err != nil {
		return nil, err
	}

	return databaseServiceFromAPI(resp.JSON200), nil
}

// ListDatabaseServices returns the list of Database Services.
func (c *Client) ListDatabaseServices(ctx context.Context, zone string) ([]*DatabaseService, error) {
	list := make([]*DatabaseService, 0)

	resp, err := c.ListDbaasServicesWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.DbaasServices != nil {
		for i := range *resp.JSON200.DbaasServices {
			list = append(list, databaseServiceFromAPI(&(*resp.JSON200.DbaasServices)[i]))
		}
	}

	return list, nil
}

// UpdateDatabaseService updates a Database Service.
func (c *Client) UpdateDatabaseService(ctx context.Context, zone string, databaseService *DatabaseService) error {
	_, err := c.UpdateDbaasServiceWithResponse(
		apiv2.WithZone(ctx, zone),
		*databaseService.Name,
		oapi.UpdateDbaasServiceJSONRequestBody{
			Maintenance: func() (v *struct {
				Dow  oapi.UpdateDbaasServiceJSONBodyMaintenanceDow `json:"dow"`
				Time string                                        `json:"time"`
			}) {
				if databaseService.Maintenance != nil {
					v = &struct {
						Dow  oapi.UpdateDbaasServiceJSONBodyMaintenanceDow `json:"dow"`
						Time string                                        `json:"time"`
					}{
						Dow:  oapi.UpdateDbaasServiceJSONBodyMaintenanceDow(databaseService.Maintenance.DOW),
						Time: databaseService.Maintenance.Time,
					}
				}
				return
			}(),
			Plan:                  databaseService.Plan,
			TerminationProtection: databaseService.TerminationProtection,
			UserConfig: func() (v *oapi.UpdateDbaasServiceJSONBody_UserConfig) {
				if databaseService.UserConfig != nil {
					v = &oapi.UpdateDbaasServiceJSONBody_UserConfig{
						AdditionalProperties: *databaseService.UserConfig,
					}
				}
				return
			}(),
		})
	if err != nil {
		return err
	}

	return nil
}
