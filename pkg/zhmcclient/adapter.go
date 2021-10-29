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

// AdapterAPI defines an interface for issuing Adapter requests to ZHMC
//go:generate counterfeiter -o fakes/adapter.go --fake-name AdapterAPI . AdapterAPI
type AdapterAPI interface {
	ListAdapters(cpcID string, query map[string]string) ([]Adapter, error)
	CreateHipersocket(cpcID string, adaptor *HypersocketPayload) (string, error)
	DeleteHipersocket(adapterID string) error
}

type AdapterManager struct {
	client ClientAPI
}

func NewAdapterManager(client ClientAPI) *AdapterManager {
	return &AdapterManager{
		client: client,
	}
}

/**
* GET /api/cpcs/{cpc-id}/adapters
* @cpcID the ID of the CPC
* @query the fields can be queried include:
*        name,
*        adapter-id,
*        adapter-family,
*        type,
*        status
* @return adapter array
* Return: 200 and Adapters array
*     or: 400, 404, 409
 */
func (m *AdapterManager) ListAdapters(cpcID string, query map[string]string) ([]Adapter, error) {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/cpcs", cpcID, "/adapters")
	requestUrl, err := BuildUrlFromUri(requestUri, query)
	if err != nil {
		return nil, err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		adapters := []Adapter{}
		err = json.Unmarshal(responseBody, &adapters)
		if err != nil {
			return nil, err
		}
		return adapters, nil
	}

	return nil, GenerateErrorFromResponse(status, responseBody)
}

/**
* POST /api/cpcs/{cpc-id}/adapters
* @cpcID the ID of the CPC
* @adaptor the payload includes properties when create Hipersocket
* Return: 201 and body with "object-uri"
*     or: 400, 403, 404, 409, 503
 */
func (m *AdapterManager) CreateHipersocket(cpcID string, adaptor *HypersocketPayload) (string, error) {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/cpcs", cpcID, "adapters")
	requestUrl, err := BuildUrlFromUri(requestUri, nil)
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(adaptor)
	if err != nil {
		return "", err
	}

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, bytes)
	if err != nil {
		return "", err
	}

	if status == http.StatusCreated {
		uriObj := HipersocketCreateResponse{}
		err = json.Unmarshal(responseBody, &uriObj)
		if err != nil {
			return "", err
		}
		return uriObj.URI, nil
	}

	return "", GenerateErrorFromResponse(status, responseBody)
}

/**
* DELETE /api/adapters/{adapter-id}
* @adapterID the adapter ID to be deleted
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *AdapterManager) DeleteHipersocket(adapterID string) error {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/adapters", adapterID)
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
