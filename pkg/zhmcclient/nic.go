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

// NicAPI defines an interface for issuing NIC requests to ZHMC
//go:generate counterfeiter -o fakes/nic.go --fake-name NicAPI . NicAPI
type NicAPI interface {
	CreateNic(lparURI string, nic *NIC) (string, int, *HmcError)
	DeleteNic(nicURI string) (int, *HmcError)
	GetNicProperties(nicURI string) (*NIC, int, *HmcError)
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
func (m *NicManager) CreateNic(lparURI string, nic *NIC) (string, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "nics")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nic, "")
	if err != nil {
		logger.Error("error on create nic",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return "", status, err
	}

	if status == http.StatusCreated {
		uriObj := NicCreateResponse{}
		err := json.Unmarshal(responseBody, &uriObj)
		if err != nil {
			return "", status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("request url: %v, status: %v, nic uri: %v", requestUrl, status, uriObj.URI))
		return uriObj.URI, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error on create nic",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(errors.New(errorResponseBody.Message)))
	return "", status, errorResponseBody

}

/**
* DELETE /api/partitions/{partition-id}/nics/{nic-id}
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *NicManager) DeleteNic(nicURI string) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, nicURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodDelete, requestUrl, nil, "")
	if err != nil {
		logger.Error("error on delete nic",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return status, err
	}

	if status == http.StatusNoContent {
		logger.Info(fmt.Sprintf("nic deleted, request url: %v, status: %v", requestUrl, status))
		return status, nil
	}

	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error on delete nic",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return status, errorResponseBody
}

/**
* GET /api/partitions/{partition-id}/nics/{nic-id}
* Return: 200 and LparProperties
*     or: 400, 404,
 */
func (m *NicManager) GetNicProperties(nicURI string) (*NIC, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, nicURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("error on get nic properties",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return nil, status, err
	}

	if status == http.StatusOK {
		nic := NIC{}
		err := json.Unmarshal(responseBody, &nic)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("request url: %v, status: %v, nic propeties: %v", requestUrl, status, &nic))
		return &nic, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error on get nic properties",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return nil, status, errorResponseBody
}
