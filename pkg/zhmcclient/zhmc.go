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
}

type ZhmcManager struct {
	client         ClientAPI
	cpcManager     CpcAPI
	lparManager    LparAPI
	adapterManager AdapterAPI
	nicManager     NicAPI
	jobManager     JobAPI
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

// CPC
func (m *ZhmcManager) ListCPCs(query map[string]string) ([]CPC, error) {
	return m.cpcManager.ListCPCs(query)
}

// LPAR
func (m *ZhmcManager) ListLPARs(cpcURI string, query map[string]string) ([]LPAR, error) {
	return m.lparManager.ListLPARs(cpcURI, query)
}
func (m *ZhmcManager) GetLparProperties(lparURI string) (*LparProperties, error) {
	return m.lparManager.GetLparProperties(lparURI)
}
func (m *ZhmcManager) UpdateLparProperties(lparURI string, props *LparProperties) error {
	return m.lparManager.UpdateLparProperties(lparURI, props)
}
func (m *ZhmcManager) StartLPAR(lparURI string) (string, error) {
	return m.lparManager.StartLPAR(lparURI)
}
func (m *ZhmcManager) StopLPAR(lparURI string) (string, error) {
	return m.lparManager.StopLPAR(lparURI)
}
func (m *ZhmcManager) MountIsoImage(lparURI string, isoFile string, insFile string) error {
	return m.lparManager.MountIsoImage(lparURI, isoFile, insFile)
}
func (m *ZhmcManager) UnmountIsoImage(lparURI string) error {
	return m.lparManager.UnmountIsoImage(lparURI)
}
func (m *ZhmcManager) ListNics(lparURI string) ([]string, error) {
	return m.lparManager.ListNics(lparURI)
}

// Adapter
func (m *ZhmcManager) ListAdapters(cpcURI string, query map[string]string) ([]Adapter, error) {
	return m.adapterManager.ListAdapters(cpcURI, query)
}
func (m *ZhmcManager) CreateHipersocket(cpcURI string, adaptor *HipersocketPayload) (string, error) {
	return m.adapterManager.CreateHipersocket(cpcURI, adaptor)
}
func (m *ZhmcManager) DeleteHipersocket(adapterURI string) error {
	return m.adapterManager.DeleteHipersocket(adapterURI)
}

// NIC
func (m *ZhmcManager) CreateNic(lparURI string, nic *NIC) (string, error) {
	return m.nicManager.CreateNic(lparURI, nic)
}
func (m *ZhmcManager) DeleteNic(nicURI string) error {
	return m.nicManager.DeleteNic(nicURI)
}
func (m *ZhmcManager) GetNicProperties(nicURI string) (*NIC, error) {
	return m.nicManager.GetNicProperties(nicURI)
}

// JOB
func (m *ZhmcManager) QueryJob(jobURI string) (*Job, error) {
	return m.jobManager.QueryJob(jobURI)
}
func (m *ZhmcManager) DeleteJob(jobURI string) error {
	return m.jobManager.DeleteJob(jobURI)
}
func (m *ZhmcManager) CancelJob(jobURI string) error {
	return m.jobManager.CancelJob(jobURI)
}
