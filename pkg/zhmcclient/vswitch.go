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

// VirtualSwitchAPI defines an interface for issuing VirtualSwitch requests to ZHMC
//go:generate counterfeiter -o fakes/vswitch.go --fake-name VirtualSwitchAPI . VirtualSwitchAPI
type VirtualSwitchAPI interface {
	ListVirtualSwitches(cpcURI string, query map[string]string) ([]VirtualSwitch, int, *HmcError)
	GetVirtualSwitchProperties(vSwitchURI string) (*VirtualSwitchProperties, int, *HmcError)
}

type VirtualSwitchManager struct {
	client ClientAPI
}

func NewVirtualSwitchManager(client ClientAPI) *VirtualSwitchManager {
	return &VirtualSwitchManager{
		client: client,
	}
}

/**
 * GET /api/cpcs/{cpc-id}/virtual-switches
 * @cpcURI the URI of the CPC
 * @return adapter array
 * Return: 200 and VirtualSwitches array
 *     or: 400, 404, 409
 */
func (m *VirtualSwitchManager) ListVirtualSwitches(cpcURI string, query map[string]string) ([]VirtualSwitch, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, cpcURI, "virtual-switches")
	requestUrl = BuildUrlFromQuery(requestUrl, query)
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("Error listing virtual switches",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return nil, status, err
	}

	if status == http.StatusOK {
		virtualSwitches := &VirtualSwitchesArray{}
		err := json.Unmarshal(responseBody, virtualSwitches)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Status: %v, Virtual switches: %v", status, virtualSwitches.VIRTUALSWITCHES))
		return virtualSwitches.VIRTUALSWITCHES, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("Error on listing virtual switches",
		genlog.String("Status: ", fmt.Sprint(status)),
		genlog.Error(errors.New(errorResponseBody.Message)))
	return nil, status, errorResponseBody
}

/**
 * GET /api/virtual-switches/{vswitch-id}
 * @cpcURI the ID of the virtual switch
 * @return adapter array
 * Return: 200 and VirtualSwitchProperties
 *     or: 400, 404, 409
 */
func (m *VirtualSwitchManager) GetVirtualSwitchProperties(vSwitchURI string) (*VirtualSwitchProperties, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, vSwitchURI)
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("Error getting virtual switch properties",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return nil, status, err
	}

	if status == http.StatusOK {
		virtualSwitch := &VirtualSwitchProperties{}
		err := json.Unmarshal(responseBody, virtualSwitch)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Status: %v, Virtual switch properties: %v", status, virtualSwitch))
		return virtualSwitch, status, nil
	}

	return nil, status, GenerateErrorFromResponse(responseBody)
}
