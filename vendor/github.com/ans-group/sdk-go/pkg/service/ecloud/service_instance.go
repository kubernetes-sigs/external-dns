package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetInstances retrieves a list of instances
func (s *Service) GetInstances(parameters connection.APIRequestParameters) ([]Instance, error) {
	return connection.InvokeRequestAll(s.GetInstancesPaginated, parameters)
}

// GetInstancesPaginated retrieves a paginated list of instances
func (s *Service) GetInstancesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Instance], error) {
	body, err := s.getInstancesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetInstancesPaginated), err
}

func (s *Service) getInstancesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Instance], error) {
	body := &connection.APIResponseBodyData[[]Instance]{}

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

func (s *Service) getInstanceResponseBody(instanceID string) (*connection.APIResponseBodyData[Instance], error) {
	body := &connection.APIResponseBodyData[Instance]{}

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

func (s *Service) createInstanceResponseBody(req CreateInstanceRequest) (*connection.APIResponseBodyData[Instance], error) {
	body := &connection.APIResponseBodyData[Instance]{}

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

func (s *Service) powerOnInstanceResponseBody(instanceID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) powerOffInstanceResponseBody(instanceID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) powerResetInstanceResponseBody(instanceID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) powerShutdownInstanceResponseBody(instanceID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) powerRestartInstanceResponseBody(instanceID string) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) migrateInstanceResponseBody(instanceID string, req MigrateInstanceRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
		return s.GetInstanceVolumesPaginated(instanceID, p)
	}, parameters)
}

// GetInstanceVolumesPaginated retrieves a paginated list of instance volumes
func (s *Service) GetInstanceVolumesPaginated(instanceID string, parameters connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
	body, err := s.getInstanceVolumesPaginatedResponseBody(instanceID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Volume], error) {
		return s.GetInstanceVolumesPaginated(instanceID, p)
	}), err
}

func (s *Service) getInstanceVolumesPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Volume], error) {
	body := &connection.APIResponseBodyData[[]Volume]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Credential], error) {
		return s.GetInstanceCredentialsPaginated(instanceID, p)
	}, parameters)
}

// GetInstanceCredentialsPaginated retrieves a paginated list of instance credentials
func (s *Service) GetInstanceCredentialsPaginated(instanceID string, parameters connection.APIRequestParameters) (*connection.Paginated[Credential], error) {
	body, err := s.getInstanceCredentialsPaginatedResponseBody(instanceID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Credential], error) {
		return s.GetInstanceCredentialsPaginated(instanceID, p)
	}), err
}

func (s *Service) getInstanceCredentialsPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Credential], error) {
	body := &connection.APIResponseBodyData[[]Credential]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
		return s.GetInstanceNICsPaginated(instanceID, p)
	}, parameters)
}

// GetInstanceNICsPaginated retrieves a paginated list of instance NICs
func (s *Service) GetInstanceNICsPaginated(instanceID string, parameters connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
	body, err := s.getInstanceNICsPaginatedResponseBody(instanceID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[NIC], error) {
		return s.GetInstanceNICsPaginated(instanceID, p)
	}), err
}

func (s *Service) getInstanceNICsPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]NIC], error) {
	body := &connection.APIResponseBodyData[[]NIC]{}

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

func (s *Service) createInstanceConsoleSessionResponseBody(instanceID string) (*connection.APIResponseBodyData[ConsoleSession], error) {
	body := &connection.APIResponseBodyData[ConsoleSession]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetInstanceTasksPaginated(instanceID, p)
	}, parameters)
}

// GetInstanceTasksPaginated retrieves a paginated list of Instance tasks
func (s *Service) GetInstanceTasksPaginated(instanceID string, parameters connection.APIRequestParameters) (*connection.Paginated[Task], error) {
	body, err := s.getInstanceTasksPaginatedResponseBody(instanceID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Task], error) {
		return s.GetInstanceTasksPaginated(instanceID, p)
	}), err
}

func (s *Service) getInstanceTasksPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Task], error) {
	body := &connection.APIResponseBodyData[[]Task]{}

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

func (s *Service) attachInstanceVolumeResponseBody(instanceID string, req AttachDetachInstanceVolumeRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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

func (s *Service) detachInstanceVolumeResponseBody(instanceID string, req AttachDetachInstanceVolumeRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	body := &connection.APIResponseBodyData[TaskReference]{}

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
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[FloatingIP], error) {
		return s.GetInstanceFloatingIPsPaginated(instanceID, p)
	}, parameters)
}

// GetInstanceFloatingIPsPaginated retrieves a paginated list of instance floating IPs
func (s *Service) GetInstanceFloatingIPsPaginated(instanceID string, parameters connection.APIRequestParameters) (*connection.Paginated[FloatingIP], error) {
	body, err := s.getInstanceFloatingIPsPaginatedResponseBody(instanceID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[FloatingIP], error) {
		return s.GetInstanceFloatingIPsPaginated(instanceID, p)
	}), err
}

func (s *Service) getInstanceFloatingIPsPaginatedResponseBody(instanceID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]FloatingIP], error) {
	body := &connection.APIResponseBodyData[[]FloatingIP]{}

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

// CreateInstanceImage attaches a volume to an instance
func (s *Service) CreateInstanceImage(instanceID string, req CreateInstanceImageRequest) (TaskReference, error) {
	body, err := s.createInstanceImageResponseBody(instanceID, req)

	return body.Data, err
}

func (s *Service) createInstanceImageResponseBody(instanceID string, req CreateInstanceImageRequest) (*connection.APIResponseBodyData[TaskReference], error) {
	if instanceID == "" {
		return &connection.APIResponseBodyData[TaskReference]{}, fmt.Errorf("invalid instance id")
	}

	return connection.Post[TaskReference](s.connection, fmt.Sprintf("/ecloud/v2/instances/%s/create-image", instanceID), &req, connection.NotFoundResponseHandler(&InstanceNotFoundError{ID: instanceID}))
}
