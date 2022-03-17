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
	ListAdapters(cpcURI string, query map[string]string) ([]Adapter, int, *HmcError)
	GetAdapterProperties(adapterURI string) (*AdapterProperties, int, *HmcError)
	CreateHipersocket(cpcURI string, adaptor *HipersocketPayload) (string, int, *HmcError)
	DeleteHipersocket(adapterURI string) (int, *HmcError)
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
* @cpcURI the ID of the CPC
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
func (m *AdapterManager) ListAdapters(cpcURI string, query map[string]string) ([]Adapter, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, cpcURI, "/adapters")
	requestUrl = BuildUrlFromQuery(requestUrl, query)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		return nil, status, err
	}

	if status == http.StatusOK {
		adapters := &AdaptersArray{}
		err := json.Unmarshal(responseBody, adapters)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return adapters.ADAPTERS, status, nil
	}

	return nil, status, GenerateErrorFromResponse(responseBody)
}

/**
* GET /api/adapters/{adapter-id}
* GET /api/adapters/{adapter-id}/network-ports/{network-port-id}
* @adapterURI the adapter ID, network-port-id for which properties need to be fetched
* @return adapter properties
* Return: 200 and Adapters properties
*     or: 400, 404, 409
 */
func (m *AdapterManager) GetAdapterProperties(adapterURI string) (*AdapterProperties, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, adapterURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		return nil, status, err
	}

	if status == http.StatusOK {
		adapterProps := &AdapterProperties{}
		err := json.Unmarshal(responseBody, adapterProps)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return adapterProps, status, nil
	}

	return nil, status, GenerateErrorFromResponse(responseBody)
}

/**
* POST /api/cpcs/{cpc-id}/adapters
* @cpcURI the ID of the CPC
* @adaptor the payload includes properties when create Hipersocket
* Return: 201 and body with "object-uri"
*     or: 400, 403, 404, 409, 503
 */
func (m *AdapterManager) CreateHipersocket(cpcURI string, adaptor *HipersocketPayload) (string, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, cpcURI, "/adapters")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, adaptor, "")
	if err != nil {
		return "", status, err
	}

	if status == http.StatusCreated {
		uriObj := HipersocketCreateResponse{}
		err := json.Unmarshal(responseBody, &uriObj)
		if err != nil {
			return "", status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return uriObj.URI, status, nil
	}

	return "", status, GenerateErrorFromResponse(responseBody)
}

/**
* DELETE /api/adapters/{adapter-id}
* @adapterURI the adapter ID to be deleted
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *AdapterManager) DeleteHipersocket(adapterURI string) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, adapterURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodDelete, requestUrl, nil, "")
	if err != nil {
		return status, err
	}

	if status == http.StatusNoContent {
		return status, nil
	}

	return status, GenerateErrorFromResponse(responseBody)
}
