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

// ZhmcAPI defines an interface for issuing requests to ZHMC
//go:generate counterfeiter -o fakes/zhmc.go --fake-name ZhmcAPI . ZhmcAPI
type ZhmcAPI interface {
	CpcAPI
	LparAPI
	NicAPI
	AdapterAPI
	JobAPI

	// TODO, interface for caches
	//GetCPCs() ([]CPC, error)
	//GetLPARs(cpcName string) ([]LPAR, error)
	//GetAdapters(cpcName string) ([]Adapter, error)
	//GetNics(cpcName string, lparName string) ([]NIC, error)
}

type ZhmcManager struct {
	client         ClientAPI
	cpcManager     CpcAPI
	lparManager    LparAPI
	adapterManager AdapterAPI
	nicManager     NicAPI
	jobManager     JobAPI
	// TODO model to cache objects...
	//cpcs     []CPC
	//lpars    []LPAR
	//adpaters []Adapter
	//nics     []NIC
}

func NewManagerFromOptions(endpoint string, creds *Options) ZhmcAPI {
	client, _ := NewClient(endpoint, creds)
	if client != nil {
		return NewManagerFromClient(client)
	}
	return nil
}

func NewManagerFromClient(client ClientAPI) ZhmcAPI {
	return &ZhmcManager{
		client:         client,
		cpcManager:     NewCpcManager(client),
		lparManager:    NewLparManager(client),
		adapterManager: NewAdapterManager(client),
		nicManager:     NewNicManager(client),
		jobManager:     NewJobManager(client),
	}
}

// ZHMC
//func (m *ZhmcManager) GetCPCs() []CPC {
//	return nil
//}
//func (m *ZhmcManager) GetLPARs(cpcName string) []LPAR {
//	return nil
//}
//func (m *ZhmcManager) GetAdapters(cpcName string) []Adapter {
//	return nil
//}
//func (m *ZhmcManager) GetNics(cpcName string, lparName string) []NIC {
//	return nil
//}

// CPC
func (m *ZhmcManager) ListCPCs(query map[string]string) ([]CPC, error) {
	return m.cpcManager.ListCPCs(query)
}

// LPAR
func (m *ZhmcManager) ListLPARs(cpcID string, query map[string]string) ([]LPAR, error) {
	return m.lparManager.ListLPARs(cpcID, query)
}
func (m *ZhmcManager) GetLparProperties(lparID string) (*LparProperties, error) {
	return m.lparManager.GetLparProperties(lparID)
}
func (m *ZhmcManager) UpdateLparProperties(lparID string, props *LparProperties) error {
	return m.lparManager.UpdateLparProperties(lparID, props)
}
func (m *ZhmcManager) StartLPAR(lparID string) (string, error) {
	return m.lparManager.StartLPAR(lparID)
}
func (m *ZhmcManager) StopLPAR(lparID string) (string, error) {
	return m.lparManager.StopLPAR(lparID)
}
func (m *ZhmcManager) MountIsoImage(lparID string, image []byte, isoFile string, insFile string) error {
	return m.lparManager.MountIsoImage(lparID, image, isoFile, insFile)
}
func (m *ZhmcManager) UnmountIsoImage(lparID string) error {
	return m.lparManager.UnmountIsoImage(lparID)
}
func (m *ZhmcManager) ListNics(lparID string) ([]string, error) {
	return m.lparManager.ListNics(lparID)
}

// Adapter
func (m *ZhmcManager) ListAdapters(cpcID string, query map[string]string) ([]Adapter, error) {
	return m.adapterManager.ListAdapters(cpcID, query)
}
func (m *ZhmcManager) CreateHipersocket(cpcID string, adaptor *HypersocketPayload) (string, error) {
	return m.adapterManager.CreateHipersocket(cpcID, adaptor)
}
func (m *ZhmcManager) DeleteHipersocket(adapterID string) error {
	return m.adapterManager.DeleteHipersocket(adapterID)
}

// NIC
func (m *ZhmcManager) CreateNic(lparID string, nic *NIC) (string, error) {
	return m.nicManager.CreateNic(lparID, nic)
}
func (m *ZhmcManager) DeleteNic(lparID string, nicID string) error {
	return m.nicManager.DeleteNic(lparID, nicID)
}
func (m *ZhmcManager) GetNicProperties(lparID string, nicID string) (*NIC, error) {
	return m.nicManager.GetNicProperties(lparID, nicID)
}

// JOB
func (m *ZhmcManager) QueryJob(jobID string) (*Job, error) {
	return m.jobManager.QueryJob(jobID)
}
func (m *ZhmcManager) DeleteJob(jobID string) error {
	return m.jobManager.DeleteJob(jobID)
}
func (m *ZhmcManager) CancelJob(jobID string) error {
	return m.jobManager.CancelJob(jobID)
}
