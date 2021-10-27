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

// CpcAPI defines an interface for issuing CPC requests to ZHMC
//go:generate counterfeiter -o fakes/cpc.go --fake-name CpcAPI . CpcAPI
type CpcAPI interface {
	ListCPCs() ([]CPC, error)
}

/**
* Sample:
* {
*    "dpm-enabled": true,
*    "has-unacceptable-status": true,
*    "name": "P0LXSMOZ",
*    "object-uri": "/api/cpcs/e8753ff5-8ea6-35d9-b047-83c2624ba8da",
*    "se-version": "2.13.1"
*    "status": "not-operating"
* }
 */
type CPC struct {
	URI                 string `json:"object-uri"`
	Name                string `json:"name"`
	Status              string `json:"status"`
	HasAcceptableStatus string `json:"has-unacceptable-status"`
	DpmEnabled          bool   `json:"dpm-enabled"`
	SeVersion           string `json:"se-version"`
}

type CpcManager struct {
	client ClientAPI
	cpcs   []CPC
}

func NewCpcManager(client ClientAPI) *CpcManager {
	return &CpcManager{
		client: client,
		cpcs:   nil,
	}
}

/**
* GET /api/cpcs
* Return: 200 and CPCs array
*     or: 400
 */
func (m *CpcManager) ListCPCs() ([]CPC, error) {
	requestUri := path.Join(m.client.GetEndpointURL().Path, "/api/cpcs")
	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUri, nil)

	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		err = json.Unmarshal(responseBody, &m.cpcs)
		if err != nil {
			return nil, err
		}
		return m.cpcs, nil
	}

	return nil, errors.New("ListCPCs -- Unknown Error")
}
