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
	StorageGroupAPI
	VirtualSwitchAPI
	JobAPI
}

type ZhmcManager struct {
	client               ClientAPI
	cpcManager           CpcAPI
	lparManager          LparAPI
	adapterManager       AdapterAPI
	storageGroupManager  StorageGroupAPI
	virtualSwitchManager VirtualSwitchAPI
	nicManager           NicAPI
	jobManager           JobAPI
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
		client:               client,
		cpcManager:           NewCpcManager(client),
		lparManager:          NewLparManager(client),
		adapterManager:       NewAdapterManager(client),
		storageGroupManager:  NewStorageGroupManager(client),
		virtualSwitchManager: NewVirtualSwitchManager(client),
		nicManager:           NewNicManager(client),
		jobManager:           NewJobManager(client),
	}
}

// CPC
func (m *ZhmcManager) ListCPCs(query map[string]string) ([]CPC, int, *HmcError) {
	return m.cpcManager.ListCPCs(query)
}

// LPAR
func (m *ZhmcManager) ListLPARs(cpcURI string, query map[string]string) ([]LPAR, int, *HmcError) {
	return m.lparManager.ListLPARs(cpcURI, query)
}
func (m *ZhmcManager) GetLparProperties(lparURI string) (*LparProperties, int, *HmcError) {
	return m.lparManager.GetLparProperties(lparURI)
}
func (m *ZhmcManager) UpdateLparProperties(lparURI string, props *LparProperties) (int, *HmcError) {
	return m.lparManager.UpdateLparProperties(lparURI, props)
}
func (m *ZhmcManager) CreateLPAR(cpcURI string, props *LparProperties) (string, int, *HmcError) {
	return m.lparManager.CreateLPAR(cpcURI, props)
}
func (m *ZhmcManager) StartLPAR(lparURI string) (string, int, *HmcError) {
	return m.lparManager.StartLPAR(lparURI)
}
func (m *ZhmcManager) StopLPAR(lparURI string) (string, int, *HmcError) {
	return m.lparManager.StopLPAR(lparURI)
}
func (m *ZhmcManager) DeleteLPAR(lparURI string) (int, *HmcError) {
	return m.lparManager.DeleteLPAR(lparURI)
}
func (m *ZhmcManager) MountIsoImage(lparURI string, isoFile string, insFile string) (int, *HmcError) {
	return m.lparManager.MountIsoImage(lparURI, isoFile, insFile)
}
func (m *ZhmcManager) UnmountIsoImage(lparURI string) (int, *HmcError) {
	return m.lparManager.UnmountIsoImage(lparURI)
}
func (m *ZhmcManager) ListNics(lparURI string) ([]string, int, *HmcError) {
	return m.lparManager.ListNics(lparURI)
}

func (m *ZhmcManager) AttachStorageGroupToPartition(lparURI string, request *StorageGroupPayload) (int, *HmcError) {
	return m.lparManager.AttachStorageGroupToPartition(lparURI, request)
}

func (m *ZhmcManager) DetachStorageGroupToPartition(lparURI string, request *StorageGroupPayload) (int, *HmcError) {
	return m.lparManager.DetachStorageGroupToPartition(lparURI, request)
}

func (m *ZhmcManager) FetchAsciiConsoleURI(lparURI string, request *AsciiConsoleURIPayload) (*AsciiConsoleURIResponse, int, *HmcError) {
	return m.lparManager.FetchAsciiConsoleURI(lparURI, request)
}

// Adapter
func (m *ZhmcManager) ListAdapters(cpcURI string, query map[string]string) ([]Adapter, int, *HmcError) {
	return m.adapterManager.ListAdapters(cpcURI, query)
}
func (m *ZhmcManager) GetAdapterProperties(adapterURI string) (*AdapterProperties, int, *HmcError) {
	return m.adapterManager.GetAdapterProperties(adapterURI)
}
func (m *ZhmcManager) CreateHipersocket(cpcURI string, adaptor *HipersocketPayload) (string, int, *HmcError) {
	return m.adapterManager.CreateHipersocket(cpcURI, adaptor)
}
func (m *ZhmcManager) DeleteHipersocket(adapterURI string) (int, *HmcError) {
	return m.adapterManager.DeleteHipersocket(adapterURI)
}

// Storage groups

func (m *ZhmcManager) ListStorageGroups(storageGroupURI string, cpc string) ([]StorageGroup, int, *HmcError) {
	return m.storageGroupManager.ListStorageGroups(storageGroupURI, cpc)
}

func (m *ZhmcManager) GetStorageGroupProperties(storageGroupURI string) (*StorageGroupProperties, int, *HmcError) {
	return m.storageGroupManager.GetStorageGroupProperties(storageGroupURI)
}

func (m *ZhmcManager) ListStorageVolumes(storageGroupURI string) ([]StorageVolume, int, *HmcError) {
	return m.storageGroupManager.ListStorageVolumes(storageGroupURI)
}

func (m *ZhmcManager) GetStorageVolumeProperties(storageGroupURI string) (*StorageVolume, int, *HmcError) {
	return m.storageGroupManager.GetStorageVolumeProperties(storageGroupURI)
}

func (m *ZhmcManager) UpdateStorageGroupProperties(storageGroupURI string, uploadRequest *StorageGroupProperties) (int, *HmcError) {
	return m.storageGroupManager.UpdateStorageGroupProperties(storageGroupURI, uploadRequest)
}

func (m *ZhmcManager) FulfillStorageGroup(storageGroupURI string, updateRequest *StorageGroupProperties) (int, *HmcError) {
	return m.storageGroupManager.FulfillStorageGroup(storageGroupURI, updateRequest)
}

// Virtual Switches

func (m *ZhmcManager) ListVirtualSwitches(cpcURI string, query map[string]string) ([]VirtualSwitch, int, *HmcError) {
	return m.virtualSwitchManager.ListVirtualSwitches(cpcURI, query)
}

func (m *ZhmcManager) GetVirtualSwitchProperties(vSwitchURI string) (*VirtualSwitchProperties, int, *HmcError) {
	return m.virtualSwitchManager.GetVirtualSwitchProperties(vSwitchURI)
}

// NIC
func (m *ZhmcManager) CreateNic(lparURI string, nic *NIC) (string, int, *HmcError) {
	return m.nicManager.CreateNic(lparURI, nic)
}
func (m *ZhmcManager) DeleteNic(nicURI string) (int, *HmcError) {
	return m.nicManager.DeleteNic(nicURI)
}
func (m *ZhmcManager) GetNicProperties(nicURI string) (*NIC, int, *HmcError) {
	return m.nicManager.GetNicProperties(nicURI)
}

// JOB
func (m *ZhmcManager) QueryJob(jobURI string) (*Job, int, *HmcError) {
	return m.jobManager.QueryJob(jobURI)
}
func (m *ZhmcManager) DeleteJob(jobURI string) (int, *HmcError) {
	return m.jobManager.DeleteJob(jobURI)
}
func (m *ZhmcManager) CancelJob(jobURI string) (int, *HmcError) {
	return m.jobManager.CancelJob(jobURI)
}
