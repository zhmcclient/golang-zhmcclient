/*
 * =============================================================================
 * IBM Confidential
 * © Copyright IBM Corp. 2020, 2021
 *
 * The source code for this program is not published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with the
 * U.S. Copyright Office.
 * =============================================================================
 */

package zhmcclient

import (
	"encoding/json"
	"net/http"
	"path"
)

// JobAPI defines an interface for issuing Job requests to ZHMC
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fakes/job.go --fake-name JobAPI . JobAPI
type JobAPI interface {
	QueryJob(jobURI string) (*Job, *HmcError)
	DeleteJob(jobURI string) *HmcError
	CancelJob(jobURI string) *HmcError
}

type JobManager struct {
	client ClientAPI
}

func NewJobManager(client ClientAPI) *JobManager {
	return &JobManager{
		client: client,
	}
}

/**
* GET /api/jobs/{job-id}
* Return: 200 and job status
*     or: 400, 404
 */
func (m *JobManager) QueryJob(jobURI string) (*Job, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, jobURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		myjob := Job{}
		err := json.Unmarshal(responseBody, &myjob)
		if err != nil {
			return nil, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return &myjob, nil
	}

	return nil, GenerateErrorFromResponse(responseBody)
}

/**
* DELETE /api/jobs/{job-id}
* Return: 204
*     or: 400, 404, 409
 */
func (m *JobManager) DeleteJob(jobURI string) *HmcError {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, jobURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodDelete, requestUrl, nil)
	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(responseBody)
}

/**
* POST /api/jobs/{job-id}/operations/cancel
* Return: 204
*     or: 400, 404, 409
 */
func (m *JobManager) CancelJob(jobURI string) *HmcError {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, jobURI, "operations/cancel")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil)
	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(responseBody)
}
