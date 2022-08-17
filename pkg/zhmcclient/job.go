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
	"errors"
	"fmt"
	"net/http"
	"path"

	"github.ibm.com/genctl/shared-logger/genlog"
)

// JobAPI defines an interface for issuing Job requests to ZHMC
//go:generate counterfeiter -o fakes/job.go --fake-name JobAPI . JobAPI
type JobAPI interface {
	QueryJob(jobURI string) (*Job, int, *HmcError)
	DeleteJob(jobURI string) (int, *HmcError)
	CancelJob(jobURI string) (int, *HmcError)
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
func (m *JobManager) QueryJob(jobURI string) (*Job, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, jobURI)
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("Error on get on job uri",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return nil, status, err
	}

	if status == http.StatusOK {
		myjob := Job{}
		err := json.Unmarshal(responseBody, &myjob)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Status: %v, Adapters: %v", status, &myjob))
		return &myjob, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("Error on get on job uri",
		genlog.String("Status: ", fmt.Sprint(status)),
		genlog.Error(errors.New(errorResponseBody.Message)))
	return nil, status, errorResponseBody
}

/**
* DELETE /api/jobs/{job-id}
* Return: 204
*     or: 400, 404, 409
 */
func (m *JobManager) DeleteJob(jobURI string) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, jobURI)
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodDelete, requestUrl, nil, "")
	if err != nil {
		logger.Error("Error on delete job",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return status, err
	}

	if status == http.StatusNoContent {
		logger.Info(fmt.Sprintf("Job deleted, Status: %v", status))
		return status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("Error on delete job",
		genlog.String("Status: ", fmt.Sprint(status)),
		genlog.Error(errors.New(errorResponseBody.Message)))
	return status, errorResponseBody
}

/**
* POST /api/jobs/{job-id}/operations/cancel
* Return: 204
*     or: 400, 404, 409
 */
func (m *JobManager) CancelJob(jobURI string) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, jobURI, "operations/cancel")
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil, "")
	if err != nil {
		logger.Error("Error on cancel job",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return status, err
	}

	if status == http.StatusNoContent {
		logger.Info(fmt.Sprintf("Job cancelled, Status: %v", status))
		return status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("Error on cancel job",
		genlog.String("Status: ", fmt.Sprint(status)),
		genlog.Error(errors.New(errorResponseBody.Message)))
	return status, errorResponseBody
}
