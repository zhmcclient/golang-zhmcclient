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
	URI    string        `json:"object-uri"`
	Name   string        `json:"name"`
	ID     string        `json:"adapter-id"`
	Family AdapterFamily `json:"adapter-family"`
	Type   AdapterType   `json:"type"`
	Status AdapterStatus `json:"status"`
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
type HypersocketPayload struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	PortDescription string `json:"port-description"`
	MaxUnitSize     int    `json:"maximum-transmission-unit-size"`
}

type HipersocketCreateResponse struct {
	URI string `json:"object-uri"`
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
	URI                 string    `json:"object-uri"`
	Name                string    `json:"name"`
	Status              CpcStatus `json:"status"`
	HasAcceptableStatus bool      `json:"has-unacceptable-status"`
	DpmEnabled          bool      `json:"dpm-enabled"`
	SeVersion           string    `json:"se-version"`
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

type Job struct {
	URI           string
	Status        JobStatus `json:"status"`
	JobStatusCode int       `json:"job-status-code"`
	JobReasonCode int       `json:"job-reason-code"`
	JobResults    []byte    `json:"job-results"`
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
	URI    string          `json:"object-uri"`
	Name   string          `json:"name"`
	Status PartitionStatus `json:"status"`
	Type   PartitionType   `json:"type"`
}

type LPARsArray struct {
	LPARS []LPAR `json:"partitions"`
}

type PartitionFeatureInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	State       bool   `json:"state"`
}

type LparProperties struct {
	URI                              string                        `json:"object-uri"`
	CpcURI                           string                        `json:"parent"`
	Class                            string                        `default:"partition"`
	Name                             string                        `json:"name"`
	Description                      string                        `json:"description"`
	Status                           PartitionStatus               `json:"status"`
	Type                             PartitionType                 `json:"type"`
	ShortName                        string                        `json:"short-name"`
	ID                               string                        `json:"partition-id"`
	AutoGenerateID                   bool                          `json:"autogenerate-partition-id"`
	OsName                           string                        `json:"os-name"`
	OsType                           string                        `json:"os-type"`
	OsVersion                        string                        `json:"os-version"`
	ReserveResourves                 bool                          `json:"reserve-resources"`
	DegradedAdapters                 []string                      `json:"degraded-adapters"`
	ProcessorMode                    PartitionProcessorMode        `json:"processor-mode"`
	CpProcessors                     int                           `json:"cp-processors"`
	IflProcessors                    int                           `json:"ifl-processors"`
	IflAbsoluteProcessorCapping      bool                          `json:"ifl-absolute-processor-capping"`
	CpAbsoluteProcessorCapping       bool                          `json:"cp-absolute-processor-capping"`
	IflAbsoluteProcessorCappingValue float64                       `json:"ifl-absolute-processor-capping-value"`
	CpAbsoluteProcessorCappingValue  float64                       `json:"cp-absolute-processor-capping-value"`
	IflProcessingWeightCapped        bool                          `json:"ifl-processing-weight-capped"`
	CpProcessingWeightCapped         bool                          `json:"cp-processing-weight-capped"`
	MinimumIflProcessingWeight       int                           `json:"minimum-ifl-processing-weight"`
	MinimumCpProcessingWeight        int                           `json:"minimum-cp-processing-weight"`
	InitialIflProcessingWeight       int                           `json:"initial-ifl-processing-weight"`
	InitialCpProcessingWeight        int                           `json:"initial-cp-processing-weight"`
	CurrentIflProcessingWeight       int                           `json:"current-ifl-processing-weight"`
	CurrentCpProcessingWeight        int                           `json:"current-cp-processing-weight"`
	MaximumIflProcessingWeight       int                           `json:"maximum-ifl-processing-weight"`
	MaximumCpProcessingWeight        int                           `json:"maximum-cp-processing-weight"`
	ProcessorManagementEnabled       bool                          `json:"processor-management-enabled"`
	InitialMemory                    int                           `json:"initial-memory"`
	ReservedMemory                   int                           `json:"reserved-memory"`
	MaximumMemory                    int                           `json:"maximum-memory"`
	AutoStart                        bool                          `json:"auto-start"`
	BootDevice                       PartitionBootDevice           `json:"boot-device"`
	BootNetworkDevice                string                        `json:"boot-network-device"`
	BootFtpHost                      string                        `json:"boot-ftp-host"`
	BootFtpUsername                  string                        `json:"boot-ftp-username"`
	BootFtpPassword                  string                        `json:"boot-ftp-password"`
	BootFtpInsfile                   string                        `json:"boot-ftp-insfile"`
	BootRemovableMedia               string                        `json:"boot-removable-media"`
	BootRemovableMediaType           PartionBootRemovableMediaType `json:"boot-removable-media-type"`
	BootTimeout                      int                           `json:"boot-timeout"`
	BootStorageDevice                string                        `json:"boot-storage-device"`
	BootStorageVolume                string                        `json:"boot-storage-volume"`
	BootLogicalUnitNumber            string                        `json:"boot-logical-unit-number"`
	BootWorldWidePortName            string                        `json:"boot-world-wide-port-name"`
	BootConfigurationSelector        int                           `json:"boot-configuration-selector"`
	BootRecordLba                    string                        `json:"boot-record-lba"`
	BootLoadParameters               string                        `json:"boot-load-parameters"`
	BootOsSpecificParameters         string                        `json:"boot-os-specific-parameters"`
	BootIsoImageName                 string                        `json:"boot-iso-image-name"`
	BootIsoInsFile                   string                        `json:"boot-iso-ins-file"`
	AccessGlobalPerformanceData      bool                          `json:"access-global-performance-data"`
	PermitCrossPartitionCommands     bool                          `json:"permit-cross-partition-commands"`
	AccessBasicCounterSet            bool                          `json:"access-basic-counter-set"`
	AccessProblemStateCounterSet     bool                          `json:"access-problem-state-counter-set"`
	AccessCryptoActivityCounterSet   bool                          `json:"access-crypto-activity-counter-set"`
	AccessExtendedCounterSet         bool                          `json:"access-extended-counter-set"`
	AccessCoprocessorGroupSet        bool                          `json:"access-coprocessor-group-set"`
	AccessBasicSampling              bool                          `json:"access-basic-sampling"`
	AccessDiagnosticSampling         bool                          `json:"access-diagnostic-sampling"`
	PermitDesKeyImportFunctions      bool                          `json:"permit-des-key-import-functions"`
	PermitAesKeyImportFunctions      bool                          `json:"permit-aes-key-import-functions"`
	ThreadsPerProcessor              int                           `json:"threads-per-processor"`
	VirtualFunctionUris              []string                      `json:"virtual-function-uris"`
	NicUris                          []string                      `json:"nic-uris"`
	HbaUris                          []string                      `json:"hba-uris"`
	CryptoConfiguration              []byte                        `json:"crypto-configuration"`
	SscHostName                      string                        `json:"ssc-host-name"`
	SscBootSelection                 SscBootSelection              `json:"ssc-boot-selection"`
	SscIpv4Gateway                   string                        `json:"ssc-ipv4-gateway"`
	SscIpv6Gateway                   string                        `json:"ssc-ipv6-gateway"`
	SscDnsServers                    []string                      `json:"ssc-dns-servers"`
	SscMasterUserid                  string                        `json:"ssc-master-userid"`
	SscMasterPw                      string                        `json:"ssc-master-pw"`
	AvailableFeaturesList            []PartitionFeatureInfo        `json:"available-features-list"`
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
	ID     string `json:"element-id"`
	URI    string `json:"element-uri"`
	Parent string `json:"parent"`
	Class  string `default:"nic"`
	/* below are payloads when create a new Nic */
	Name                  string           `json:"name"`
	Description           string           `json:"description"`
	DeviceNumber          string           `json:"device-number"`
	NetworkAdapterPortURI string           `json:"network-adapter-port-uri"`
	VirtualSwitchUriType  string           `json:"virtual-switch-uri-type"`
	VirtualSwitchURI      string           `json:"virtual-switch-uri"`
	Type                  NicType          `json:"type"`
	SscManagmentNIC       bool             `json:"ssc-management-nic"`
	SscIpAddressType      SscIpAddressType `json:"ssc-ip-address-type"`
	SscIpAddress          string           `json:"ssc-ip-address"`
	VlanID                int              `json:"vlan-id"`
	MacAddress            string           `json:"mac-address"`
	SscMaskPrefix         string           `json:"ssc-mask-prefix"`
	VlanType              VlanType         `json:"vlan-type"`
}

type NicCreateResponse struct {
	URI string `json:"element-uri"`
}
