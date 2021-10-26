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

// LparAPI defines an interface for issuing LPAR requests to ZHMC
//go:generate counterfeiter -o fakes/lpar.go --fake-name LparAPI . LparAPI
type LparAPI interface {
	ListLPARs(cpcID string) ([]LPAR, error)
	UpdateLparProperties(lparID string, props map[string]string) (*LPAR, error)
	StartLPAR(lparID string) (string, error)
	StopLPAR(lparID string) (string, error)

	MountIsoImage(lparID string, isoFile string, insFile string) error

	ListNics(lparID string) ([]string, error)
	CreateNic(lparID string, nic *NIC) (*NIC, error)
	DeleteNic(lparID string, nicID string) error
	GetNic(lparID string, nicID string) (*NIC, error)
}

/**
* Sample:
* {
*    name: Required
*    description
*    network-adapter-port-uri
*    virtual-switch-uri
*    device-number
*    ssc-management-nic
*    ssc-ip-address-type
*    ssc-ip-address
*    vlan-id
*    ssc-mask-prefix
*    mac-address
*    vlan-type
*    element-uri: "/api/partitions/b4c4bf9e-97e0-11e5-9d1f-020000000192/nics/eb6887e4-
      97e8-11e5-9d1f-020000000192", Returned when create
* }
*/
type NIC struct {
	uri    string
	device string
	lpar   *LPAR
	Name   string
	Mac    string
}

/**
*/
type LPAR struct {
	uri       string
	cpc       *CPC
	Name      string
	Status    string
	Type      string
}

type LparManager struct {
	client *Client
	lpars  []LPAR
}

func NewLparManager(client *Client) LparAPI {
	return &LparManager{
		client: client,
		lpars:  nil,
	}
}

/**
* GET /api/cpcs/{cpc-id}/partitions
* Return: 200 and LPARs array
*     or: 400, 404, 409
*/
func (m *LparManager) ListLPARs(cpcID string) ([]LPAR, error) {
	return nil, nil
}

func (m *LparManager) UpdateLparProperties(lparID string, props map[string]string) (*LPAR, error) {
	return nil, nil
}

/**
* POST /api/partitions/{partition-id}/operations/start
* Return: 202 and job-uri
*     or: 400, 403, 404, 503, 
*/
func (m *LparManager) StartLPAR(lparID string) (string, error) {
	return "nil", nil
}

/**
* POST /api/partitions/{partition-id}/operations/stop
* Return: 202 and job-uri
*     or: 400, 403, 404, 503, 
*/
func (m *LparManager) StopLPAR(lparID string) (string, error) {
	return "nil", nil
}

/**
* POST /api/partitions/{partition-id}/operations/mount-iso-image
* Return: 204 
*     or: 400, 403, 404, 409, 503
*/
func (m *LparManager) MountIsoImage(lparID string, isoFile string, insFile string) error {
	return nil
}

/**
* get_property('nic-uris') from LPAR
*/
func (m *LparManager) ListNics(lparID string) ([]string, error) {
	return nil, nil
}

/**
* POST /api/partitions/{partition-id}/nics
*/
func (m *LparManager) CreateNic(lparID string, nic *NIC) (*NIC, error) {
	return nil, nil
}

/**
* DELETE /api/partitions/{partition-id}/nics/{nic-id}
* Return: 204
*     or: 400, 403, 404, 409, 503
*/
func (m *LparManager) DeleteNic(lparID string, nicID string) error {
	return nil
}

/**
* GET /api/partitions/{partition-id}/nics/{nic-id}
*/
func (m *LparManager) GetNic(lparID string, nicID string) (*NIC, error) {
	return nil, nil
}