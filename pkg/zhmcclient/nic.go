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

// NicAPI defines an interface for issuing NIC requests to ZHMC
//go:generate counterfeiter -o fakes/nic.go --fake-name NicAPI . NicAPI
type NicAPI interface {
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

type NicManager struct {
	client *Client
	nics   []NIC
}

func NewNicManager(client *Client) NicAPI {
	return &NicManager{
		client: client,
		nics:   nil,
	}
}

/**
* get_property('nic-uris') from LPAR
 */
func (m *NicManager) ListNics(lparID string) ([]string, error) {
	return nil, nil
}

/**
* POST /api/partitions/{partition-id}/nics
 */
func (m *NicManager) CreateNic(lparID string, nic *NIC) (*NIC, error) {
	return nil, nil
}

/**
* DELETE /api/partitions/{partition-id}/nics/{nic-id}
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *NicManager) DeleteNic(lparID string, nicID string) error {
	return nil
}

/**
* GET /api/partitions/{partition-id}/nics/{nic-id}
 */
func (m *NicManager) GetNic(lparID string, nicID string) (*NIC, error) {
	return nil, nil
}
