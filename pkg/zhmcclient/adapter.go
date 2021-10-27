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

// AdapterAPI defines an interface for issuing Adapter requests to ZHMC
//go:generate counterfeiter -o fakes/adapter.go --fake-name AdapterAPI . AdapterAPI
type AdapterAPI interface {
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
	Name                        string // Required
	Description                 string
	PortDescription             string
	MaximumTransmissionUnitSize int
}

type AdapterManager struct {
	client   *Client
	adapters []Adapter
}

func NewAdapterManager(client *Client) AdapterAPI {
	return &AdapterManager{
		client:   client,
		adapters: nil,
	}
}

/**
* GET /api/cpcs/{cpc-id}/adapters
 */
func (m *AdapterManager) ListAdapters(cpcID string) ([]Adapter, error) {
	return nil, nil
}

/**
* POST /api/cpcs/{cpc-id}/adapters
* Return: 201 and "object-uri"
*     or: 400, 403, 404, 409, 503
 */
func (m *AdapterManager) CreateAdapter(cpcID string, adaptor *Adapter) (*Adapter, error) {
	return nil, nil
}

/**
* DELETE /api/adapters/{adapter-id}
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *AdapterManager) DeleteAdapter(cpcID string) error {
	return nil
}
