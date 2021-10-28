package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetInstances retrieves a list of instances
func (s *Service) GetInstances(parameters connection.APIRequestParameters) ([]Instance, error) {
	var instances []Instance

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstancesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, instance := range response.(*PaginatedInstance).Items {
			instances = append(instances, instance)
		}
	}

	return instances, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetInstancesPaginated retrieves a paginated list of instances
func (s *Service) GetInstancesPaginated(parameters connection.APIRequestParameters) (*PaginatedInstance, error) {
	body, err := s.getInstancesPaginatedResponseBody(parameters)

	return NewPaginatedInstance(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstancesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInstancesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetInstanceSliceResponseBody, error) {
	body := &GetInstanceSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/instances", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetInstance retrieves a single instance by id
func (s *Service) GetInstance(instanceID string) (Instance, error) {
	body, err := s.getInstanceResponseBody(instanceID)

	return body.Data, err
}

func (s *Service) getInstanceResponseBody(instanceID string) (*GetInstanceResponseBody, error) {
	body := &GetInstanceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/instances/%s", instanceID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// CreateInstance creates a new instance
func (s *Service) CreateInstance(req CreateInstanceRequest) (string, error) {
	body, err := s.createInstanceResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createInstanceResponseBody(req CreateInstanceRequest) (*GetInstanceResponseBody, error) {
	body := &GetInstanceResponseBody{}

	response, err := s.connection.Post("/ecloud/v2/instances", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchInstance updates an instance
func (s *Service) PatchInstance(instanceID string, req PatchInstanceRequest) error {
	_, err := s.patchInstanceResponseBody(instanceID, req)

	return err
}

func (s *Service) patchInstanceResponseBody(instanceID string, req PatchInstanceRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/instances/%s", instanceID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// DeleteInstance removes an instance
func (s *Service) DeleteInstance(instanceID string) error {
	_, err := s.deleteInstanceResponseBody(instanceID)

	return err
}

func (s *Service) deleteInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/instances/%s", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// LockInstance locks an instance from update/removal
func (s *Service) LockInstance(instanceID string) error {
	_, err := s.lockInstanceResponseBody(instanceID)

	return err
}

func (s *Service) lockInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/lock", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// UnlockInstance unlocks an instance
func (s *Service) UnlockInstance(instanceID string) error {
	_, err := s.unlockInstanceResponseBody(instanceID)

	return err
}

func (s *Service) unlockInstanceResponseBody(instanceID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/unlock", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// PowerOnInstance powers on an instance
func (s *Service) PowerOnInstance(instanceID string) (string, error) {
	body, err := s.powerOnInstanceResponseBody(instanceID)

	return body.Data.TaskID, err
}

func (s *Service) powerOnInstanceResponseBody(instanceID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/power-on", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// PowerOffInstance powers off an instance
func (s *Service) PowerOffInstance(instanceID string) (string, error) {
	body, err := s.powerOffInstanceResponseBody(instanceID)

	return body.Data.TaskID, err
}

func (s *Service) powerOffInstanceResponseBody(instanceID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/power-off", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// PowerResetInstance resets an instance
func (s *Service) PowerResetInstance(instanceID string) (string, error) {
	body, err := s.powerResetInstanceResponseBody(instanceID)

	return body.Data.TaskID, err
}

func (s *Service) powerResetInstanceResponseBody(instanceID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/power-reset", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// PowerShutdownInstance shuts down an instance
func (s *Service) PowerShutdownInstance(instanceID string) (string, error) {
	body, err := s.powerShutdownInstanceResponseBody(instanceID)

	return body.Data.TaskID, err
}

func (s *Service) powerShutdownInstanceResponseBody(instanceID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/power-shutdown", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// PowerRestartInstance restarts an instance
func (s *Service) PowerRestartInstance(instanceID string) (string, error) {
	body, err := s.powerRestartInstanceResponseBody(instanceID)

	return body.Data.TaskID, err
}

func (s *Service) powerRestartInstanceResponseBody(instanceID string) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/ecloud/v2/instances/%s/power-restart", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// MigrateInstance migrates an instance
func (s *Service) MigrateInstance(instanceID string, req MigrateInstanceRequest) (string, error) {
	body, err := s.migrateInstanceResponseBody(instanceID, req)

	return body.Data.TaskID, err
}

func (s *Service) migrateInstanceResponseBody(instanceID string, req MigrateInstanceRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/instances/%s/migrate", instanceID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// GetInstanceVolumes retrieves a list of instance volumes
func (s *Service) GetInstanceVolumes(instanceID string, parameters connection.APIRequestParameters) ([]Volume, error) {
	var volumes []Volume

	return volumes, connection.InvokeRequestAll(
		func(p connection.APIRequestParameters) (connection.Paginated, error) {
			return s.GetInstanceVolumesPaginated(instanceID, p)
		},
		func(response connection.Paginated) {
			for _, volume := range response.(*PaginatedVolume).Items {
				volumes = append(volumes, volume)
			}
		},
		parameters,
	)
}

// GetInstanceVolumesPaginated retrieves a paginated list of instance volumes
func (s *Service) GetInstanceVolumesPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedVolume, error) {
	body, err := s.getInstanceVolumesPaginatedResponseBody(instanceID, parameters)

	return NewPaginatedVolume(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstanceVolumesPaginated(instanceID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInstanceVolumesPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*GetVolumeSliceResponseBody, error) {
	body := &GetVolumeSliceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/instances/%s/volumes", instanceID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// GetInstanceCredentials retrieves a list of instance credentials
func (s *Service) GetInstanceCredentials(instanceID string, parameters connection.APIRequestParameters) ([]Credential, error) {
	var credentials []Credential

	return credentials, connection.InvokeRequestAll(
		func(p connection.APIRequestParameters) (connection.Paginated, error) {
			return s.GetInstanceCredentialsPaginated(instanceID, p)
		},
		func(response connection.Paginated) {
			for _, credential := range response.(*PaginatedCredential).Items {
				credentials = append(credentials, credential)
			}
		},
		parameters,
	)
}

// GetInstanceCredentialsPaginated retrieves a paginated list of instance credentials
func (s *Service) GetInstanceCredentialsPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedCredential, error) {
	body, err := s.getInstanceCredentialsPaginatedResponseBody(instanceID, parameters)

	return NewPaginatedCredential(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstanceCredentialsPaginated(instanceID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInstanceCredentialsPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*GetCredentialSliceResponseBody, error) {
	body := &GetCredentialSliceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/instances/%s/credentials", instanceID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// GetInstanceNICs retrieves a list of instance NICs
func (s *Service) GetInstanceNICs(instanceID string, parameters connection.APIRequestParameters) ([]NIC, error) {
	var nics []NIC

	return nics, connection.InvokeRequestAll(
		func(p connection.APIRequestParameters) (connection.Paginated, error) {
			return s.GetInstanceNICsPaginated(instanceID, p)
		},
		func(response connection.Paginated) {
			for _, nic := range response.(*PaginatedNIC).Items {
				nics = append(nics, nic)
			}
		},
		parameters,
	)
}

// GetInstanceNICsPaginated retrieves a paginated list of instance NICs
func (s *Service) GetInstanceNICsPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedNIC, error) {
	body, err := s.getInstanceNICsPaginatedResponseBody(instanceID, parameters)

	return NewPaginatedNIC(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstanceNICsPaginated(instanceID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInstanceNICsPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*GetNICSliceResponseBody, error) {
	body := &GetNICSliceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/instances/%s/nics", instanceID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// CreateInstanceConsoleSession creates an instance console session
func (s *Service) CreateInstanceConsoleSession(instanceID string) (ConsoleSession, error) {
	body, err := s.createInstanceConsoleSessionResponseBody(instanceID)

	return body.Data, err
}

func (s *Service) createInstanceConsoleSessionResponseBody(instanceID string) (*GetConsoleSessionResponseBody, error) {
	body := &GetConsoleSessionResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/instances/%s/console-session", instanceID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// GetInstanceTasks retrieves a list of Instance tasks
func (s *Service) GetInstanceTasks(instanceID string, parameters connection.APIRequestParameters) ([]Task, error) {
	var tasks []Task

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstanceTasksPaginated(instanceID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, task := range response.(*PaginatedTask).Items {
			tasks = append(tasks, task)
		}
	}

	return tasks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetInstanceTasksPaginated retrieves a paginated list of Instance tasks
func (s *Service) GetInstanceTasksPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedTask, error) {
	body, err := s.getInstanceTasksPaginatedResponseBody(instanceID, parameters)

	return NewPaginatedTask(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstanceTasksPaginated(instanceID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInstanceTasksPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*GetTaskSliceResponseBody, error) {
	body := &GetTaskSliceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/instances/%s/tasks", instanceID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// AttachInstanceVolume attaches a volume to an instance
func (s *Service) AttachInstanceVolume(instanceID string, req AttachDetachInstanceVolumeRequest) (string, error) {
	body, err := s.attachInstanceVolumeResponseBody(instanceID, req)

	return body.Data.TaskID, err
}

func (s *Service) attachInstanceVolumeResponseBody(instanceID string, req AttachDetachInstanceVolumeRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/instances/%s/volume-attach", instanceID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// DetachInstanceVolume detaches a volume from an instance
func (s *Service) DetachInstanceVolume(instanceID string, req AttachDetachInstanceVolumeRequest) (string, error) {
	body, err := s.detachInstanceVolumeResponseBody(instanceID, req)

	return body.Data.TaskID, err
}

func (s *Service) detachInstanceVolumeResponseBody(instanceID string, req AttachDetachInstanceVolumeRequest) (*GetTaskReferenceResponseBody, error) {
	body := &GetTaskReferenceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/instances/%s/volume-detach", instanceID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}

// GetInstanceFloatingIPs retrieves a list of instance fips
func (s *Service) GetInstanceFloatingIPs(instanceID string, parameters connection.APIRequestParameters) ([]FloatingIP, error) {
	var fips []FloatingIP

	return fips, connection.InvokeRequestAll(
		func(p connection.APIRequestParameters) (connection.Paginated, error) {
			return s.GetInstanceFloatingIPsPaginated(instanceID, p)
		},
		func(response connection.Paginated) {
			for _, fip := range response.(*PaginatedFloatingIP).Items {
				fips = append(fips, fip)
			}
		},
		parameters,
	)
}

// GetInstanceFloatingIPsPaginated retrieves a paginated list of instance floating IPs
func (s *Service) GetInstanceFloatingIPsPaginated(instanceID string, parameters connection.APIRequestParameters) (*PaginatedFloatingIP, error) {
	body, err := s.getInstanceFloatingIPsPaginatedResponseBody(instanceID, parameters)

	return NewPaginatedFloatingIP(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetInstanceFloatingIPsPaginated(instanceID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getInstanceFloatingIPsPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*GetFloatingIPSliceResponseBody, error) {
	body := &GetFloatingIPSliceResponseBody{}

	if instanceID == "" {
		return body, fmt.Errorf("invalid instance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/instances/%s/floating-ips", instanceID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &InstanceNotFoundError{ID: instanceID}
		}

		return nil
	})
}
