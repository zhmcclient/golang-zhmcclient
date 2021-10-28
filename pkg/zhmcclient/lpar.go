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
	"net/http"
	"path"
)

// LparAPI defines an interface for issuing LPAR requests to ZHMC
//go:generate counterfeiter -o fakes/lpar.go --fake-name LparAPI . LparAPI
type LparAPI interface {
	ListLPARs(cpcID string, query map[string]string) ([]LPAR, error)
	UpdateLparProperties(lparID string, props map[string]string) (*LPAR, error)
	StartLPAR(lparID string) (string, error)
	StopLPAR(lparID string) (string, error)

	MountIsoImage(lparID string, isoFile string, insFile string) error
	UnmountIsoImage(lparID string) error
}

type PartitionType string

const (
	PARTITION_TYPE_SSC   PartitionType = "ssc"
	PARTITION_TYPE_LINUX               = "linux"
	PARTITION_TYPE_ZVM                 = "zvm"
)

type PartitionStatus string

const (
	PARTITION_STATUS_NOT_ACTIVE   PartitionStatus = "communications-not-active"
	PARTITION_STATUS_STATUS_CHECK                 = "status-check"
	PARTITION_STATUS_STOPPED                      = "stopped"
	PARTITION_STATUS_TERMINATED                   = "terminated"
	PARTITION_STATUS_STARTING                     = "starting"
	PARTITION_STATUS_ACTIVE                       = "active"
	PARTITION_STATUS_STOPPING                     = "stopping"
	PARTITION_STATUS_DEGRADED                     = "degraded"
	PARTITION_STATUS_REV_ERR                      = "reservation-error"
	PARTITION_STATUS_PAUSED                       = "paused"
)

/**
 */
type LPAR struct {
	URI    string          `json:"object-uri"`
	Name   string          `json:"name"`
	Status PartitionStatus `json:"status"`
	Type   PartitionStatus `json:"type"`
	cpc    *CPC
}

type StartStopLparResponse struct {
	URI     string `json:"job-uri"`
	Message string `json:"message"`
}

type LparManager struct {
	client ClientAPI
	lpars  []LPAR
}

func NewLparManager(client ClientAPI) *LparManager {
	return &LparManager{
		client: client,
		lpars:  nil,
	}
}

/**
* GET /api/cpcs/{cpc-id}/partitions
* @cpcID is the cpc object-uri
* @query is a key, value pairs array,
*        currently, supports 'name=$name_reg_expression'
* @return lpar array
* Return: 200 and LPARs array
*     or: 400, 404, 409
 */
func (m *LparManager) ListLPARs(cpcID string, query map[string]string) ([]LPAR, error) {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/cpcs", cpcID, "/partitions")
	requestUrl, err := BuildUrlFromUri(requestUri, query)
	if err != nil {
		return nil, err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		err = json.Unmarshal(responseBody, &m.lpars)
		if err != nil {
			return nil, err
		}
		return m.lpars, nil
	}

	return nil, GenerateErrorFromResponse(status, responseBody)
}

/**
* POST /api/partitions/{partition-id}
* @lparID is the object-uri
 */
func (m *LparManager) UpdateLparProperties(lparID string, props map[string]string) (*LPAR, error) {
	return nil, nil
}

/**
* POST /api/partitions/{partition-id}/operations/start
* @lparID is the object-uri
* @return job-uri
* Return: 202 and job-uri
*     or: 400, 403, 404, 503,
 */
func (m *LparManager) StartLPAR(lparID string) (string, error) {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/partitions", lparID, "/operations/start")
	requestUrl, err := BuildUrlFromUri(requestUri, nil)
	if err != nil {
		return "", err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return "", err
	}

	if status == http.StatusAccepted {
		responseObj := StartStopLparResponse{}
		err = json.Unmarshal(responseBody, &responseObj)
		if err != nil {
			return "", err
		}
		if responseObj.URI != "" {
			return responseObj.URI, nil
		}
		return "", errors.New("Succeeded start LPAR, but got empty job-uri.")
	}

	return "", GenerateErrorFromResponse(status, responseBody)
}

/**
* POST /api/partitions/{partition-id}/operations/stop
* @lparID is the object-uri
* @return job-uri
* Return: 202 and job-uri
*     or: 400, 403, 404, 503,
 */
func (m *LparManager) StopLPAR(lparID string) (string, error) {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/partitions", lparID, "/operations/stop")
	requestUrl, err := BuildUrlFromUri(requestUri, nil)
	if err != nil {
		return "", err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return "", err
	}

	if status == http.StatusAccepted {
		responseObj := StartStopLparResponse{}
		err = json.Unmarshal(responseBody, &responseObj)
		if err != nil {
			return "", err
		}
		if responseObj.URI != "" {
			return responseObj.URI, nil
		}
		return "", errors.New("Succeeded stop LPAR, but got empty job-uri.")
	}

	return "", GenerateErrorFromResponse(status, responseBody)
}

/**
* POST /api/partitions/{partition-id}/operations/mount-iso-image
* @lparID is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *LparManager) MountIsoImage(lparID string, isoFile string, insFile string) error {
	return nil
}

/**
* POST /api/partitions/{partition-id}/operations/unmount-iso-image
* @lparID is the object-uri
 */
func (m *LparManager) UnmountIsoImage(lparID string) error {
	return nil
}
