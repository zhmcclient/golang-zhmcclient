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

)

// CpcAPI defines an interface for issuing CPC requests to ZHMC
//go:generate counterfeiter -o fakes/cpc.go --fake-name CpcAPI . CpcAPI
type CpcAPI interface {
	ListCPCs() ([]CPC, error)

	ListAdapters(cpcID string) ([]Adapter, error)
	CreateAdapter(cpcID string, adaptor *Adapter) (*Adapter, error)
	DeleteAdapter(cpcID string) error
}

/**
* Sample
* {
* 	"name":
*   "description":
*   "port-description":
*   "maximum-transmission-unit-size":
*   "object-uri":"/api/adapters/542b9406-d033-11e5-9f39-020000000338", retured when creating
* }
*/
type Adapter struct {
	uri                         string
	Name 						string  // Required
	Description 				string
	PortDescription             string
	MaximumTransmissionUnitSize int
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

/**
* GET /api/cpcs/{cpc-id}/adapters
*/
func (m *CpcManager) ListAdapters(cpcID string) ([]Adapter, error) {
	return nil, nil
}

/**
* POST /api/cpcs/{cpc-id}/adapters
* Return: 201 and "object-uri"
*     or: 400, 403, 404, 409, 503
*/
func (m *CpcManager) CreateAdapter(cpcID string, adaptor *Adapter) (*Adapter, error) {
	return nil, nil
}

/**
* DELETE /api/adapters/{adapter-id}
* Return: 204
*     or: 400, 403, 404, 409, 503
*/
func (m *CpcManager) DeleteAdapter(cpcID string) error {
	return nil
}