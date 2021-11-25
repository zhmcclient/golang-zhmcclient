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
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fakes/nic.go --fake-name NicAPI . NicAPI
type NicAPI interface {
	CreateNic(lparURI string, nic *NIC) (string, *HmcError)
	DeleteNic(nicURI string) *HmcError
	GetNicProperties(nicURI string) (*NIC, *HmcError)
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
func (m *NicManager) CreateNic(lparURI string, nic *NIC) (string, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "nics")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nic)
	if err != nil {
		return "", err
	}

	if status == http.StatusCreated {
		uriObj := NicCreateResponse{}
		err := json.Unmarshal(responseBody, &uriObj)
		if err != nil {
			return "", getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return uriObj.URI, nil
	}

	return "", GenerateErrorFromResponse(responseBody)
}

/**
* DELETE /api/partitions/{partition-id}/nics/{nic-id}
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *NicManager) DeleteNic(nicURI string) *HmcError {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, nicURI)

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
* GET /api/partitions/{partition-id}/nics/{nic-id}
* Return: 200 and LparProperties
*     or: 400, 404,
 */
func (m *NicManager) GetNicProperties(nicURI string) (*NIC, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, nicURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		nic := NIC{}
		err := json.Unmarshal(responseBody, &nic)
		if err != nil {
			return nil, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}

		return &nic, nil
	}

	return nil, GenerateErrorFromResponse(responseBody)
}
