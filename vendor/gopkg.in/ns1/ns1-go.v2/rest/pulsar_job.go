package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/pulsar"
)

// PulsarJobsService handles 'pulsar/apps/APPID/jobs/JOBID' endpoint.
type PulsarJobsService service

// List takes an Application ID and returns all Jobs inside said Application.
//
// NS1 API docs: https://ns1.com/api/#getlist-jobs-within-an-app
func (s *PulsarJobsService) List(appID string) ([]*pulsar.Job, *http.Response, error) {
	path := fmt.Sprintf("pulsar/apps/%s/jobs", appID)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	jl := []*pulsar.Job{}
	var resp *http.Response
	resp, err = s.client.Do(req, &jl)
	if err != nil {
		switch errorType := err.(type) {
		case *Error:
			if errorType.Resp.StatusCode == 404 {
				return nil, resp, ErrAppMissing
			}
		}
		return nil, resp, err
	}

	return jl, resp, nil
}

// Get takes an Application ID and Job Id and returns full configuration for a pulsar Job.
//
// NS1 API docs: https://ns1.com/api/#getview-job-details
func (s *PulsarJobsService) Get(appID string, jobID string) (*pulsar.Job, *http.Response, error) {
	path := fmt.Sprintf("pulsar/apps/%s/jobs/%s", appID, jobID)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var job pulsar.Job

	resp, err := s.client.Do(req, &job)
	if err != nil {
		switch errorType := err.(type) {
		case *Error:
			jobNotFound := fmt.Sprintf("pulsar job %s not found for appid %s", jobID, appID)
			switch errorType.Message {
			case jobNotFound:
				return nil, resp, ErrJobMissing

			case "pulsar app not found":
				return nil, resp, ErrAppMissing
			}
		}
		return nil, resp, err
	}

	return &job, resp, nil
}

// Create takes a *PulsarJob and an AppId and creates a new Pulsar Job in the specified Application with the specific name, typeid, host and url_path.
//
// NS1 API docs: https://ns1.com/api/#putcreate-a-pulsar-job
func (s *PulsarJobsService) Create(j *pulsar.Job) (*http.Response, error) {
	path := fmt.Sprintf("pulsar/apps/%s/jobs", j.AppID)

	req, err := s.client.NewRequest("PUT", path, j)
	if err != nil {
		return nil, err
	}

	// Update job fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, j)
	if err != nil {
		switch errorType := err.(type) {
		case *Error:
			if errorType.Resp.StatusCode == 404 {
				return resp, ErrAppMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

// Update takes a *PulsarJob and modifies configuration details for an existing Pulsar job.
//
// Only the fields to be updated are required in the given job.
// NS1 API docs: https://ns1.com/api/#postmodify-a-pulsar-job
func (s *PulsarJobsService) Update(j *pulsar.Job) (*http.Response, error) {
	path := fmt.Sprintf("pulsar/apps/%s/jobs/%s", j.AppID, j.JobID)

	req, err := s.client.NewRequest("POST", path, j)
	if err != nil {
		return nil, err
	}

	// Update jobs fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, j)
	if err != nil {
		switch errorType := err.(type) {
		case *Error:
			jobNotFound := fmt.Sprintf("pulsar job %s not found for appid %s", j.JobID, j.AppID)
			switch errorType.Message {
			case "pulsar app not found":
				return resp, ErrAppMissing
			case jobNotFound:
				return resp, ErrJobMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

// Delete takes a appId and jobId and removes an existing Pulsar job .
//
// NS1 API docs: https://ns1.com/api/#deletedelete-a-pulsar-job
func (s *PulsarJobsService) Delete(pulsarJob *pulsar.Job) (*http.Response, error) {
	path := fmt.Sprintf("pulsar/apps/%s/jobs/%s", pulsarJob.AppID, pulsarJob.JobID)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch errorType := err.(type) {
		case *Error:
			jobNotFound := fmt.Sprintf("pulsar job %s not found for appid %s", pulsarJob.JobID, pulsarJob.AppID)
			switch errorType.Message {
			case jobNotFound:
				return resp, ErrJobMissing

			case "pulsar app not found":
				return resp, ErrAppMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

var (
	// ErrAppMissing bundles GET/PUT/POST/DELETE
	ErrAppMissing = errors.New("pulsar application does not exist")
	// ErrJobMissing bundles GET/POST/DELETE error.
	ErrJobMissing = errors.New("pulsar job does not exist")
)
