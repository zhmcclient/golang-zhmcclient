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
	"fmt"
	"net/http"
	"path"
)

// VirtualSwitchAPI defines an interface for issuing VirtualSwitch requests to ZHMC
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fakes/vswitch.go --fake-name VirtualSwitchAPI . VirtualSwitchAPI
type VirtualSwitchAPI interface {
	ListVirtualSwitches(cpcURI string) ([]VirtualSwitch, error)
	GetVirtualSwitchProperties(vsSwitchURI string) (*VirtualSwitch, error)
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
 * Return: 200 and Adapters array
 *     or: 400, 404, 409
 */
func (m *VirtualSwitchManager) ListVirtualSwitches(cpcURI string) ([]VirtualSwitch, error) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, cpcURI, "virtual-switches")
	fmt.Println(requestUrl)
	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		virtualSwitches := &VirtualSwitchesArray{}
		err = json.Unmarshal(responseBody, virtualSwitches)
		if err != nil {
			return nil, err
		}
		return virtualSwitches.VIRTUALSWITCHES, nil
	}

	return nil, GenerateErrorFromResponse(status, responseBody)
}

/**
 * GET /api/virtual-switches/{vswitch-id}
 * @cpcURI the ID of the virtual switch
 * @return adapter array
 * Return: 200 and Adapters array
 *     or: 400, 404, 409
 */
func (m *VirtualSwitchManager) GetVirtualSwitchProperties(vsSwitchURI string) (*VirtualSwitch, error) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, vsSwitchURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		virtualSwitch := &VirtualSwitch{}
		err = json.Unmarshal(responseBody, virtualSwitch)
		if err != nil {
			return nil, err
		}
		return virtualSwitch, nil
	}

	return nil, GenerateErrorFromResponse(status, responseBody)
}
