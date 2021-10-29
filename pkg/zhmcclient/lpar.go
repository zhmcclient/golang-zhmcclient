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
	GetLparProperties(lparID string) (*LparProperties, error)
	UpdateLparProperties(lparID string, props *LparProperties) error
	StartLPAR(lparID string) (string, error)
	StopLPAR(lparID string) (string, error)

	MountIsoImage(lparID string, isoFile string, insFile string) error
	UnmountIsoImage(lparID string) error

	ListNics(lparID string) ([]string, error)
}

type LparManager struct {
	client ClientAPI
}

func NewLparManager(client ClientAPI) *LparManager {
	return &LparManager{
		client: client,
	}
}

/**
* GET /api/cpcs/{cpc-id}/partitions
* @cpcID is the cpc object-uri
* @query is a key, value pairs array,
*        currently, supports 'name=$name_reg_expression'
*                            'status=PartitionStatus'
*                            'type=PartitionType'
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
		lpars := []LPAR{}
		err = json.Unmarshal(responseBody, &lpars)
		if err != nil {
			return nil, err
		}
		return lpars, nil
	}

	return nil, GenerateErrorFromResponse(status, responseBody)
}

/**
* GET /api/partitions/{partition-id}
* @lparID is the object-uri
* Return: 200 and LparProperties
*     or: 400, 404,
 */
func (m *LparManager) GetLparProperties(lparID string) (*LparProperties, error) {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/partitions", lparID)
	requestUrl, err := BuildUrlFromUri(requestUri, nil)
	if err != nil {
		return nil, err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		lparProps := LparProperties{}
		err = json.Unmarshal(responseBody, &lparProps)
		if err != nil {
			return nil, err
		}

		return &lparProps, nil
	}

	return nil, GenerateErrorFromResponse(status, responseBody)
}

/**
* POST /api/partitions/{partition-id}
* @lparID is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503,
 */
func (m *LparManager) UpdateLparProperties(lparID string, props *LparProperties) error {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/partitions", lparID)
	requestUrl, err := BuildUrlFromUri(requestUri, nil)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(props)
	if err != nil {
		return err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, bytes)
	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(status, responseBody)
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

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil)
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

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil)
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

/**
* get_property('nic-uris') from LPAR
 */
func (m *LparManager) ListNics(lparID string) ([]string, error) {
	props, err := m.GetLparProperties(lparID)
	if err != nil {
		return nil, err
	}

	return props.NicUris, nil
}
