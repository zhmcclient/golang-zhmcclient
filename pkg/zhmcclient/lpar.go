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

// LparAPI defines an interface for issuing LPAR requests to ZHMC
//go:generate counterfeiter -o fakes/lpar.go --fake-name LparAPI . LparAPI
type LparAPI interface {
	ListLPARs(cpcID string) ([]LPAR, error)
	UpdateLparProperties(lparID string, props map[string]string) (*LPAR, error)
	StartLPAR(lparID string) (string, error)
	StopLPAR(lparID string) (string, error)

	MountIsoImage(lparID string, isoFile string, insFile string) error
	UnmountIsoImage(lparID string) error
}

/**
 */
type LPAR struct {
	uri    string
	cpc    *CPC
	Name   string
	Status string
	Type   string
}

type LparManager struct {
	client ClientAPI
	lpars  []LPAR
}

func NewLparManager(client ClientAPI) *LparManager {
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
* POST /api/partitions/{partition-id}/operations/unmount-iso-image
 */
func (m *LparManager) UnmountIsoImage(lparID string) error {
	return nil
}
