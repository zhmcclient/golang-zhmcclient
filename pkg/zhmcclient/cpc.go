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
	uri                 string
	Name                string
	Status              string
	HasAcceptableStatus string
	DpmEnabled          bool
	SeVersion           string
}

type CpcManager struct {
	client *Client
	cpcs   []CPC
}

func NewCpcManager(client *Client) CpcAPI {
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
	return nil, nil
}
