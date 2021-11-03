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
	ListLPARs(cpcURI string, query map[string]string) ([]LPAR, error)
	GetLparProperties(lparURI string) (*LparProperties, error)
	UpdateLparProperties(lparURI string, props *LparProperties) error
	StartLPAR(lparURI string) (string, error)
	StopLPAR(lparURI string) (string, error)

	MountIsoImage(lparURI string, image []byte, isoFile string, insFile string) error
	UnmountIsoImage(lparURI string) error

	ListNics(lparURI string) ([]string, error)
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
* @cpcURI is the cpc object-uri
* @query is a key, value pairs array,
*        currently, supports 'name=$name_reg_expression'
*                            'status=PartitionStatus'
*                            'type=PartitionType'
* @return lpar array
* Return: 200 and LPARs array
*     or: 400, 404, 409
 */
func (m *LparManager) ListLPARs(cpcURI string, query map[string]string) ([]LPAR, error) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, cpcURI, "/partitions")
	requestUrl, err := BuildUrlFromQuery(requestUrl, query)
	if err != nil {
		return nil, err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		lpars := &LPARsArray{}
		err = json.Unmarshal(responseBody, lpars)
		if err != nil {
			return nil, err
		}
		return lpars.LPARS, nil
	}

	return nil, GenerateErrorFromResponse(status, responseBody)
}

/**
* GET /api/partitions/{partition-id}
* @lparURI is the object-uri
* Return: 200 and LparProperties
*     or: 400, 404,
 */
func (m *LparManager) GetLparProperties(lparURI string) (*LparProperties, error) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI)

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
* @lparURI is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503,
 */
func (m *LparManager) UpdateLparProperties(lparURI string, props *LparProperties) error {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, props)
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
* @lparURI is the object-uri
* @return job-uri
* Return: 202 and job-uri
*     or: 400, 403, 404, 503,
 */
func (m *LparManager) StartLPAR(lparURI string) (string, error) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/start")

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
* @lparURI is the object-uri
* @return job-uri
* Return: 202 and job-uri
*     or: 400, 403, 404, 503,
 */
func (m *LparManager) StopLPAR(lparURI string) (string, error) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/stop")

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
* @lparURI is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *LparManager) MountIsoImage(lparURI string, image []byte, isoFile string, insFile string) error {
	query := map[string]string{
		"image-name":    isoFile,
		"ins-file-name": insFile,
	}

	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/mount-iso-image")
	requestUrl, err := BuildUrlFromQuery(requestUrl, query)

	if err != nil {
		return err
	}

	status, responseBody, err := m.client.UploadRequest(http.MethodPost, requestUrl, image)
	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(status, responseBody)
}

/**
* POST /api/partitions/{partition-id}/operations/unmount-iso-image
* @lparURI is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *LparManager) UnmountIsoImage(lparURI string) error {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/unmount-iso-image")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil)
	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(status, responseBody)
}

/**
* get_property('nic-uris') from LPAR
 */
func (m *LparManager) ListNics(lparURI string) ([]string, error) {
	props, err := m.GetLparProperties(lparURI)
	if err != nil {
		return nil, err
	}

	return props.NicUris, nil
}
