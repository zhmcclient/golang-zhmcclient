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

//////////////////////////////////////////////////
// Adapter
//////////////////////////////////////////////////

type AdapterFamily string

const (
	ADAPTER_FAMILY_HIPERSOCKET AdapterFamily = "hipersockets"
	ADAPTER_FAMILY_OSA                       = "osa"
	ADAPTER_FAMILY_FICON                     = "ficon"
	ADAPTER_FAMILY_ROCE                      = "roce"
	ADAPTER_FAMILY_CRYPTO                    = "crypto"
	ADAPTER_FAMILY_ACCELERATOR               = "accelerator"
)

type AdapterType string

const (
	ADAPTER_TYPE_CRYPTO      AdapterType = "crypto"
	ADAPTER_TYPE_FCP                     = "fcp"
	ADAPTER_TYPE_HIPERSOCKET             = "hipersockets"
	ADAPTER_TYPE_OSD                     = "osd"
	ADAPTER_TYPE_OSM                     = "osm"
	ADAPTER_TYPE_ROCE                    = "roce"
	ADAPTER_TYPE_ZEDC                    = "zedc"
	ADAPTER_TYPE_FC                      = "fc"
	ADAPTER_TYPE_NOT_CFG                 = "not-configured"
)

type AdapterStatus string

const (
	ADAPTER_STATUS_ACTIVE       AdapterStatus = "active"
	ADAPTER_STATUS_NOT_ACTIVE                 = "not-active"
	ADAPTER_STATUS_NOT_DETECTED               = "not-detected"
	ADAPTER_STATUS_EXCEPTIONS                 = "exceptions"
)

type Adapter struct {
	URI    string        `json:"object-uri,omitempty"`
	Name   string        `json:"name,omitempty"`
	ID     string        `json:"adapter-id,omitempty"`
	Family AdapterFamily `json:"adapter-family,omitempty"`
	Type   AdapterType   `json:"type,omitempty"`
	Status AdapterStatus `json:"status,omitempty"`
}

type AdaptersArray struct {
	ADAPTERS []Adapter `json:"adapters"`
}

/**
* Sample
* {
* 	"name":
*   "description":
*   "port-description":
*   "maximum-transmission-unit-size":
* }
* @return *   "object-uri":"/api/adapters/542b9406-d033-11e5-9f39-020000000338"
 */
type HipersocketPayload struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	PortDescription string `json:"port-description,omitempty"`
	MaxUnitSize     int    `json:"maximum-transmission-unit-size,omitempty"`
}

type HipersocketCreateResponse struct {
	URI string `json:"object-uri"`
}

//////////////////////////////////////////////////
// Virtual Switches
//////////////////////////////////////////////////

type VirtualSwitchType string

const (
	VIRTUALSWITCH_TYPE_HIPERSOCKET VirtualSwitchType = "hipersockets"
)

type VirtualSwitchesArray struct {
	VIRTUALSWITCHES []VirtualSwitch `json:"virtual-switches"`
}

type VirtualSwitch struct {
	URI  string            `json:"object-uri,omitempty"`
	Name string            `json:"name,omitempty"`
	Type VirtualSwitchType `json:"type,omitempty"`
}

//////////////////////////////////////////////////
// CPC
//////////////////////////////////////////////////

type CpcStatus string

const (
	CPC_STATUS_ACTIVE           CpcStatus = "active"
	CPC_STATUS_OPERATING                  = "operating"
	CPC_STATUS_NO_COMMUNICATING           = "not-communicating"
	CPC_STATUS_EXCEPTIONS                 = "exceptions"
	CPC_STATUS_STATUS_CHECK               = "status-check"
	CPC_STATUS_SERVICE                    = "service"
	CPC_STATUS_NOT_OPERATING              = "not-operating"
	CPC_STATUS_NO_POWER                   = "no-power"
	CPC_STATUS_SERVICE_REQUIRED           = "service-required"
	CPC_STATUS_DEGRADED                   = "degraded"
)

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
	URI                 string    `json:"object-uri,omitempty"`
	Name                string    `json:"name,omitempty"`
	Status              CpcStatus `json:"status,omitempty"`
	HasAcceptableStatus bool      `json:"has-unacceptable-status,omitempty"`
	DpmEnabled          bool      `json:"dpm-enabled,omitempty"`
	SeVersion           string    `json:"se-version,omitempty"`
}

type CpcsArray struct {
	CPCS []CPC `json:"cpcs"`
}

//////////////////////////////////////////////////
// JOB
//////////////////////////////////////////////////

type JobStatus string

const (
	JOB_STATUS_RUNNING        JobStatus = "running"
	JOB_STATUS_CANCEL_PENDING           = "cancel-pending"
	JOB_STATUS_CANCELED                 = "canceled"
	JOB_STATUS_COMPLETE                 = "complete"
)

type JobResults struct {
	Message string `json:"message,omitempty"`
}

type Job struct {
	URI           string
	Status        JobStatus  `json:"status,omitempty"`
	JobStatusCode int        `json:"job-status-code,omitempty"`
	JobReasonCode int        `json:"job-reason-code,omitempty"`
	JobResults    JobResults `json:"job-results,omitempty"`
}

//////////////////////////////////////////////////
// LPAR
//////////////////////////////////////////////////

type PartitionType string

const (
	PARTITION_TYPE_SSC   PartitionType = "ssc"
	PARTITION_TYPE_LINUX               = "linux"
	PARTITION_TYPE_ZVM                 = "zvm"
)

type PartitionStatus string

const (
	PARTITION_STATUS_NOT_ACTIVE   PartitionStatus = "communications-not-active"
	PARTITION_STATUS_STATUS_CHECK                 = "status-check"
	PARTITION_STATUS_STOPPED                      = "stopped"
	PARTITION_STATUS_TERMINATED                   = "terminated"
	PARTITION_STATUS_STARTING                     = "starting"
	PARTITION_STATUS_ACTIVE                       = "active"
	PARTITION_STATUS_STOPPING                     = "stopping"
	PARTITION_STATUS_DEGRADED                     = "degraded"
	PARTITION_STATUS_REV_ERR                      = "reservation-error"
	PARTITION_STATUS_PAUSED                       = "paused"
)

type PartitionProcessorMode string

const (
	PROCESSOR_MODE_DEDICATED PartitionProcessorMode = "dedicated"
	PROCESSOR_MODE_SHARED                           = "shared"
)

type PartitionBootDevice string

const (
	BOOT_DEVICE_STORAGE_ADAPTER PartitionBootDevice = "storage-adapter"
	BOOT_DEVICE_STORAGE_VOLUME                      = "storage-volume"
	BOOT_DEVICE_NETWORK_ADAPTER                     = "network-adapter"
	BOOT_DEVICE_FTP                                 = "ftp"
	BOOT_DEVICE_REMOVABLE_MEDIA                     = "removable-media"
	BOOT_DEVICE_ISO_IMAGE                           = "iso-image"
	BOOT_DEVICE_NONE                                = "none" // default
)

type PartionBootRemovableMediaType string

const (
	BOOT_REMOVABLE_MEDIA_CDROM PartionBootRemovableMediaType = "cdrom"
	BOOT_REMOVABLE_MEDIA_USB                                 = "usb"
)

type SscBootSelection string

const (
	SSC_BOOT_SELECTION_INSTALLER SscBootSelection = "installer"
	SSC_BOOT_SELECTION_APPLIANCE                  = "appliance"
)

type LPAR struct {
	URI    string          `json:"object-uri,omitempty"`
	Name   string          `json:"name,omitempty"`
	Status PartitionStatus `json:"status,omitempty"`
	Type   PartitionType   `json:"type,omitempty"`
}

type LPARsArray struct {
	LPARS []LPAR `json:"partitions"`
}

type PartitionFeatureInfo struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State       bool   `json:"state,omitempty"`
}

type LparProperties struct {
	URI                              string                        `json:"object-uri,omitempty"`
	CpcURI                           string                        `json:"parent,omitempty"`
	Class                            string                        `json:"class,omitempty"`
	Name                             string                        `json:"name,omitempty"`
	Description                      string                        `json:"description,omitempty"`
	Status                           PartitionStatus               `json:"status,omitempty"`
	Type                             PartitionType                 `json:"type,omitempty"`
	ShortName                        string                        `json:"short-name,omitempty"`
	ID                               string                        `json:"partition-id,omitempty"`
	AutoGenerateID                   bool                          `json:"autogenerate-partition-id,omitempty"`
	OsName                           string                        `json:"os-name,omitempty"`
	OsType                           string                        `json:"os-type,omitempty"`
	OsVersion                        string                        `json:"os-version,omitempty"`
	ReserveResourves                 bool                          `json:"reserve-resources,omitempty"`
	DegradedAdapters                 []string                      `json:"degraded-adapters,omitempty"`
	ProcessorMode                    PartitionProcessorMode        `json:"processor-mode,omitempty"`
	CpProcessors                     int                           `json:"cp-processors,omitempty"`
	IflProcessors                    int                           `json:"ifl-processors,omitempty"`
	IflAbsoluteProcessorCapping      bool                          `json:"ifl-absolute-processor-capping,omitempty"`
	CpAbsoluteProcessorCapping       bool                          `json:"cp-absolute-processor-capping,omitempty"`
	IflAbsoluteProcessorCappingValue float64                       `json:"ifl-absolute-processor-capping-value,omitempty"`
	CpAbsoluteProcessorCappingValue  float64                       `json:"cp-absolute-processor-capping-value,omitempty"`
	IflProcessingWeightCapped        bool                          `json:"ifl-processing-weight-capped,omitempty"`
	CpProcessingWeightCapped         bool                          `json:"cp-processing-weight-capped,omitempty"`
	MinimumIflProcessingWeight       int                           `json:"minimum-ifl-processing-weight,omitempty"`
	MinimumCpProcessingWeight        int                           `json:"minimum-cp-processing-weight,omitempty"`
	InitialIflProcessingWeight       int                           `json:"initial-ifl-processing-weight,omitempty"`
	InitialCpProcessingWeight        int                           `json:"initial-cp-processing-weight,omitempty"`
	CurrentIflProcessingWeight       int                           `json:"current-ifl-processing-weight,omitempty"`
	CurrentCpProcessingWeight        int                           `json:"current-cp-processing-weight,omitempty"`
	MaximumIflProcessingWeight       int                           `json:"maximum-ifl-processing-weight,omitempty"`
	MaximumCpProcessingWeight        int                           `json:"maximum-cp-processing-weight,omitempty"`
	ProcessorManagementEnabled       bool                          `json:"processor-management-enabled,omitempty"`
	InitialMemory                    int                           `json:"initial-memory,omitempty"`
	ReservedMemory                   int                           `json:"reserved-memory,omitempty"`
	MaximumMemory                    int                           `json:"maximum-memory,omitempty"`
	AutoStart                        bool                          `json:"auto-start,omitempty"`
	BootDevice                       PartitionBootDevice           `json:"boot-device,omitempty"`
	BootNetworkDevice                string                        `json:"boot-network-device,omitempty"`
	BootFtpHost                      string                        `json:"boot-ftp-host,omitempty"`
	BootFtpUsername                  string                        `json:"boot-ftp-username,omitempty"`
	BootFtpPassword                  string                        `json:"boot-ftp-password,omitempty"`
	BootFtpInsfile                   string                        `json:"boot-ftp-insfile,omitempty"`
	BootRemovableMedia               string                        `json:"boot-removable-media,omitempty"`
	BootRemovableMediaType           PartionBootRemovableMediaType `json:"boot-removable-media-type,omitempty"`
	BootTimeout                      int                           `json:"boot-timeout,omitempty"`
	BootStorageDevice                string                        `json:"boot-storage-device,omitempty"`
	BootStorageVolume                string                        `json:"boot-storage-volume,omitempty"`
	BootLogicalUnitNumber            string                        `json:"boot-logical-unit-number,omitempty"`
	BootWorldWidePortName            string                        `json:"boot-world-wide-port-name,omitempty"`
	BootConfigurationSelector        int                           `json:"boot-configuration-selector,omitempty"`
	BootRecordLba                    string                        `json:"boot-record-lba,omitempty"`
	BootLoadParameters               string                        `json:"boot-load-parameters,omitempty"`
	BootOsSpecificParameters         string                        `json:"boot-os-specific-parameters,omitempty"`
	BootIsoImageName                 string                        `json:"boot-iso-image-name,omitempty"`
	BootIsoInsFile                   string                        `json:"boot-iso-ins-file,omitempty"`
	AccessGlobalPerformanceData      bool                          `json:"access-global-performance-data,omitempty"`
	PermitCrossPartitionCommands     bool                          `json:"permit-cross-partition-commands,omitempty"`
	AccessBasicCounterSet            bool                          `json:"access-basic-counter-set,omitempty"`
	AccessProblemStateCounterSet     bool                          `json:"access-problem-state-counter-set,omitempty"`
	AccessCryptoActivityCounterSet   bool                          `json:"access-crypto-activity-counter-set,omitempty"`
	AccessExtendedCounterSet         bool                          `json:"access-extended-counter-set,omitempty"`
	AccessCoprocessorGroupSet        bool                          `json:"access-coprocessor-group-set,omitempty"`
	AccessBasicSampling              bool                          `json:"access-basic-sampling,omitempty"`
	AccessDiagnosticSampling         bool                          `json:"access-diagnostic-sampling,omitempty"`
	PermitDesKeyImportFunctions      bool                          `json:"permit-des-key-import-functions,omitempty"`
	PermitAesKeyImportFunctions      bool                          `json:"permit-aes-key-import-functions,omitempty"`
	ThreadsPerProcessor              int                           `json:"threads-per-processor,omitempty"`
	VirtualFunctionUris              []string                      `json:"virtual-function-uris,omitempty"`
	NicUris                          []string                      `json:"nic-uris,omitempty"`
	HbaUris                          []string                      `json:"hba-uris,omitempty"`
	CryptoConfiguration              []byte                        `json:"crypto-configuration,omitempty"`
	SscHostName                      string                        `json:"ssc-host-name,omitempty"`
	SscBootSelection                 SscBootSelection              `json:"ssc-boot-selection,omitempty"`
	SscIpv4Gateway                   string                        `json:"ssc-ipv4-gateway,omitempty"`
	SscIpv6Gateway                   string                        `json:"ssc-ipv6-gateway,omitempty"`
	SscDnsServers                    []string                      `json:"ssc-dns-servers,omitempty"`
	SscMasterUserid                  string                        `json:"ssc-master-userid,omitempty"`
	SscMasterPw                      string                        `json:"ssc-master-pw,omitempty"`
	AvailableFeaturesList            []PartitionFeatureInfo        `json:"available-features-list,omitempty"`
}

type StartStopLparResponse struct {
	URI     string `json:"job-uri"`
	Message string `json:"message"`
}

//////////////////////////////////////////////////
// NIC
//////////////////////////////////////////////////

type NicType string

const (
	NIC_TYPE_ROCE NicType = "roce"
	NIC_TYPE_IQD          = "iqd"
	NIC_TYPE_OSD          = "osd"
)

type SscIpAddressType string

const (
	SSC_IP_TYPE_IPV4      SscIpAddressType = "ipv4"
	SSC_IP_TYPE_IPV6                       = "ipv6"
	SSC_IP_TYPE_LINKLOCAL                  = "linklocal"
	SSC_IP_TYPE_DHCP                       = "dhcp"
)

type VlanType string

const (
	VLAN_TYPE_ENFORCED VlanType = "enforced"
)

type NIC struct {
	ID     string `json:"element-id,omitempty"`
	URI    string `json:"element-uri,omitempty"`
	Parent string `json:"parent,omitempty"`
	Class  string `default:"nic,omitempty"`
	/* below are payloads when create a new Nic */
	Name                  string           `json:"name,omitempty"`
	Description           string           `json:"description,omitempty"`
	DeviceNumber          string           `json:"device-number,omitempty"`
	NetworkAdapterPortURI string           `json:"network-adapter-port-uri,omitempty"`
	VirtualSwitchUriType  string           `json:"virtual-switch-uri-type,omitempty"`
	VirtualSwitchURI      string           `json:"virtual-switch-uri,omitempty"`
	Type                  NicType          `json:"type,omitempty"`
	SscManagmentNIC       bool             `json:"ssc-management-nic,omitempty"`
	SscIpAddressType      SscIpAddressType `json:"ssc-ip-address-type,omitempty"`
	SscIpAddress          string           `json:"ssc-ip-address,omitempty"`
	VlanID                int              `json:"vlan-id,omitempty"`
	MacAddress            string           `json:"mac-address,omitempty"`
	SscMaskPrefix         string           `json:"ssc-mask-prefix,omitempty"`
	VlanType              VlanType         `json:"vlan-type,omitempty"`
}

type NicCreateResponse struct {
	URI string `json:"element-uri"`
}
