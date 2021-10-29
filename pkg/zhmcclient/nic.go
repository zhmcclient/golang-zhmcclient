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

// NicAPI defines an interface for issuing NIC requests to ZHMC
//go:generate counterfeiter -o fakes/nic.go --fake-name NicAPI . NicAPI
type NicAPI interface {
	CreateNic(lparID string, nic *NIC) (string, error)
	DeleteNic(lparID string, nicID string) error
	GetNicProperties(lparID string, nicID string) (*NIC, error)
}

type NicManager struct {
	client ClientAPI
}

func NewNicManager(client ClientAPI) *NicManager {
	return &NicManager{
		client: client,
	}
}

/**
* POST /api/partitions/{partition-id}/nics
* @ return element-uri
* Return: 201 and element-uri
*     or: 400, 403, 404, 409, 503,
 */
func (m *NicManager) CreateNic(lparID string, nic *NIC) (string, error) {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/partitions", lparID, "nics")
	requestUrl, err := BuildUrlFromUri(requestUri, nil)
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(nic)
	if err != nil {
		return "", err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, bytes)
	if err != nil {
		return "", err
	}

	if status == http.StatusCreated {
		uriObj := NicCreateResponse{}
		err = json.Unmarshal(responseBody, &uriObj)
		if err != nil {
			return "", err
		}
		return uriObj.URI, nil
	}

	return "", GenerateErrorFromResponse(status, responseBody)
}

/**
* DELETE /api/partitions/{partition-id}/nics/{nic-id}
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *NicManager) DeleteNic(lparID string, nicID string) error {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/partitions", lparID, "nics", nicID)
	requestUrl, err := BuildUrlFromUri(requestUri, nil)
	if err != nil {
		return err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodDelete, requestUrl, nil)
	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(status, responseBody)
}

/**
* GET /api/partitions/{partition-id}/nics/{nic-id}
* Return: 200 and LparProperties
*     or: 400, 404,
 */
func (m *NicManager) GetNicProperties(lparID string, nicID string) (*NIC, error) {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/partitions", lparID, "nics", nicID)
	requestUrl, err := BuildUrlFromUri(requestUri, nil)
	if err != nil {
		return nil, err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		nic := NIC{}
		err = json.Unmarshal(responseBody, &nic)
		if err != nil {
			return nil, err
		}

		return &nic, nil
	}

	return nil, GenerateErrorFromResponse(status, responseBody)
}
