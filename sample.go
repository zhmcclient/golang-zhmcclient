/*
 * =============================================================================
 * IBM Confidential
 * © Copyright IBM Corp. 2020, 2022
 *
 * The source code for this program is not published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with the
 * U.S. Copyright Office.
 * =============================================================================
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.ibm.com/genctl/shared-logger/genlog"
	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

var logger = genlog.New()

func main() {
	endpoint := os.Getenv("HMC_ENDPOINT") // "https://9.114.87.7:6794/", "https://192.168.195.118:6794"
	username := os.Getenv("HMC_USERNAME")
	password := os.Getenv("HMC_PASSWORD")
	cacert := os.Getenv("CA_CERT")
	skipCert := os.Getenv("SKIP_CERT")
	isSkipCert, _ := strconv.ParseBool(skipCert)

	args := os.Args[1:]
	//partitionId := os.Getenv("PAR_ID")
	// isofile := os.Getenv("ISO_FILE")
	// insfile := os.Getenv("INS_FILE")
	creds := &zhmcclient.Options{Username: username, Password: password, CaCert: cacert, SkipCert: isSkipCert, Trace: false}
	if endpoint == "" || username == "" || password == "" || (cacert == "" && !isSkipCert) {
		// The Fatal functions call os.Exit(1) after writing the log message
		logger.Fatal("Please set HMC_ENDPOINT, HMC_USERNAME, HMC_PASSWORD, CA_CERT")
	}
	if len(args) == 0 {
		logger.Fatal(`Usage: sample <Command>

			Please enter one of the below Command:
			
				"CreatePartition":
					- Creates the partition on the selected HMC
				
				"StartPartitionforHmc":
					- Starts the partition for the selected HMC
				
				"StopPartitionforHmc":
					- Stops the partition for the selected HMC

				"DeletePartitionforHmc":
					- Deletes the partition for the selected HMC
				
				"UpdateBootDeviceProperty":
					- Updates the partition as boot device="iso" for the selected HMC
				
				"MountIsoImageToPartition":
					- Mounts the iso image on the partition for the selected HMC
				
				"UnmountIsoImageToPartition":
					- Unmounts the iso image on the partition for the selected HMC
				
				"ListStorageGroupsforCPC":
					- List the storage groups of a given CPC for the selected HMC

				"ListStorageVolumesforCPC":
					- Get storage volumes of a given storage group for the selected HMC
				
				"AttachStorageGroupToPartitionofCPC":
					- Attach storage group to selected partition
				
				"DetachStorageGroupToPartitionofCPC":
					- Detach storage group from selected partition
				
				"GetAdapterPropsforCPC"
					- Get Adpater properties for adapter of given CPC

				"ListAdaptersofCPC"
					- List Adpaters for given CPC
				
				"FetchASCIIConsoleURI"
					- Get the URI to launch the Ascii Web Console

				"CreateStorageGroupsforCPC":
					- Create storage groups operation creates a new storage group object
					
				"GetStorageGroupPartitionsforCPC":
					- Get storage groups operation retrieves the properties of a single storage group object

				"DeleteStorageGroupforCPC":	
				    - Delete storage group operation deletes a storage group

				"GetStorageGroupPropertiesforCPC":
					- Get storage groups operation retrieves the properties of a single storage group object
	
					


		}`)
	} else {
		logger.Info("HMC_ENDPOINT: " + endpoint)
		logger.Info("HMC_USERNAME: " + username)
		logger.Info("HMC_PASSWORD: xxxxxx")
		client, err := zhmcclient.NewClient(endpoint, creds)
		if err != nil {
			logger.Error("Error getting client connection", genlog.Error(errors.New(err.Message)))
		}
		if client != nil {
			logger.Info("client initialized.")
			hmcManager := zhmcclient.NewManagerFromClient(client)
			/*
			 Create LPAR Base URI
			 partitionId := os.Getenv("PAR_ID")
			 lparURI := "api/partitions/" + partitionId
			*/
			/*
			 ### Usage Examples

			 #List All usage
			 - List all LPAR's for the selected HMC endpoint
			 hmcManager := zhmcclient.NewManagerFromClient(client)
			 ListAll(hmcManager)
			*/
			/*
				 ## Steps to update boot device for a partition

					 #1 Create a partition
					 - Following steps are done by ansible playbooks
						 - Create linux partition with resources reserved
						 - Create linux resources (storage group: boot vol, data vol)
							 - Create Storage group
							 - Create boot volume
							 - Create data volume
							 - Map Wwpns
							 - Create volumes on storage array (part of LUN sensing)
						 - Attach storage group to partition

					 #2 Mount Iso image Usage
					 - Mount Iso image on the partition
					 - isofile := os.Getenv("ISO_FILE")
					 - insfile := os.Getenv("INS_FILE")
					 @params:
					 - lparURI: Api endpoint for LPAR (type string)
					 - isofile: Iso file path to be mounted (type string)
					 - insfile: Ins file for the iso file (type string)
					 err := hmcManager.MountIsoImage(lparURI, isofile, insfile)

					 #3 Update Lpar Properties
					 - Update the boot device property of the partition to 'iso-image' to bring up
					 the partition with the mounted iso image
					 @params:
					 - lparURI: Api endpoint for LPAR (type string)
					 - props: Lpar properties to update (type *zhmcclient.LparProperties)
					 usage:
					 var bootDevice zhmcclient.PartitionBootDevice = zhmcclient.BOOT_DEVICE_ISO_IMAGE
					 props := &zhmcclient.LparProperties{BootDevice: bootDevice}
					 err := hmcManager.UpdateLparProperties(lparURI, props)

					 #4 Start Partition
					 - Start the partition on the selected HMC endpoint with iso image
					 @params:
					 - lparURI: Api endpoint for LPAR (type string)
					 err := hmcManager.StartPartition(lparURI)

					 printVirtualSwitches(hmcManager)

					 printVirtualSwitcheProperties(hmcManager)
			*/
			/* #List All usage
			- List all LPAR's for the selected HMC endpoint
			*/
			//ListAll(hmcManager)
			for _, arg := range args {
				switch arg {
				case "CreatePartition":
					CreatePartition(hmcManager)
				case "StartPartitionforHmc":
					StartPartitionforHmc(hmcManager)
				case "StopPartitionforHmc":
					StopPartitionforHmc(hmcManager)
				case "DeletePartitionforHmc":
					DeletePartitionforHmc(hmcManager)
				case "UpdateBootDeviceProperty":
					UpdateBootDeviceProperty(hmcManager)
				case "MountIsoImageToPartition":
					MountIsoImageToPartition(hmcManager)
				case "ListStorageGroupsforCPC":
					ListStorageGroupsforCPC(hmcManager)
				case "ListStorageVolumesforCPC":
					ListStorageVolumesforCPC(hmcManager)
				case "DetachStorageGroupToPartitionofCPC":
					DetachStorageGroupToPartitionofCPC(hmcManager)
				case "AttachStorageGroupToPartitionofCPC":
					AttachStorageGroupToPartitionofCPC(hmcManager)
				case "GetAdapterPropsforCPC":
					GetAdapterPropsforCPC(hmcManager)
				case "ListAdaptersofCPC":
					ListAdaptersofCPC(hmcManager)
				case "FetchASCIIConsoleURI":
					FetchASCIIConsoleURI(hmcManager)
				case "CreateStorageGroupsforCPC":
					CreateStorageGroupsforCPC(hmcManager)
				case "GetStorageGroupPartitionsforCPC":
					GetStorageGroupPartitionsforCPC(hmcManager)
				case "DeleteStorageGroupforCPC":
					DeleteStorageGroupforCPC(hmcManager)
				case "GetStorageGroupPropertiesforCPC":
					GetStorageGroupPropertiesforCPC(hmcManager)

				}

			}
		}
	}
}

func GetLPARURI() (lparURI string) {
	partitionId := os.Getenv("PAR_ID")
	lparURI = "api/partitions/" + partitionId

	return
}
func GetCpcURL(hmcManager zhmcclient.ZhmcAPI) (cpcuri string) {
	cpcURI := GetCPCURI(hmcManager)
	return cpcURI
}

func GetCPCURI(hmcManager zhmcclient.ZhmcAPI) string {
	query := map[string]string{}
	cpcID := ""
	cpcName := os.Getenv("CPC_NAME")
	cpcs, _, err := hmcManager.ListCPCs(query)
	if err != nil {
		logger.Error("Error: " + err.Message)
	} else {
		for _, cpc := range cpcs {
			logger.Info("########################################")
			logger.Info("cpc ENV Name: " + cpcName)
			logger.Info("cpc name: " + cpc.Name)
			logger.Info("cpc uri: " + cpc.URI)
			if cpc.Name == cpcName {
				cpcID = cpc.URI
			}
		}
	}
	return cpcID
}

func ListAdaptersofCPC(hmcManager zhmcclient.ZhmcAPI) {
	query := map[string]string{}
	CPCURI := "api/cpcs/" + os.Getenv("CPC_ID")
	adapters, _, err := hmcManager.ListAdapters(CPCURI, query)
	if err != nil {
		logger.Fatal("", genlog.Any("List Adapters error", err))
	} else {
		logger.Info("-----------------------")
		for _, adapter := range adapters {
			logger.Info("++++++++++++++++++++++++")
			logger.Info("Adapter Name:" + adapter.Name)
			logger.Info("Adapter Type:" + string(adapter.Type))
			logger.Info("Adapter Family:" + string(adapter.Family))
			logger.Info("Adapter Status:" + string(adapter.Status))
			logger.Info("Adapter URI:" + string(adapter.URI))
			adapter, _, _ := hmcManager.GetAdapterProperties(adapter.URI)
			logger.Info("********* Adapter properties **************")
			logger.Info("\n- NAME: " + adapter.Name)
			logger.Info("\n- Status: " + string(adapter.Status))
			logger.Info("\n- State: " + string(adapter.State))
			logger.Info("\n- Detected Card Type: " + string(adapter.DetectedCardType))
			logger.Info("\n- Port Count: " + fmt.Sprint(adapter.PortCount))
			logger.Info("\n- Adapter Family: " + string(adapter.Family))
			logger.Info("\n- PHYSICAL CONNECTION STATUS: " + string(adapter.PHYSICALCHANNELSTATUS))
			logger.Info("", genlog.Strings("\n- Network port URIS", adapter.NetworkAdapterPortURIs))
			logger.Info("", genlog.Strings("\n- Storage Port URIS", adapter.StoragePortURIs))
			logger.Info("*********************************************")
		}
		logger.Info("\n-----------------------")
	}
}

func GetAdapterPropsforCPC(hmcManager zhmcclient.ZhmcAPI) {
	adapterID := os.Getenv("ADAPTER_ID")
	adapterURI := "api/adapters/" + adapterID
	adapter, _, err := hmcManager.GetAdapterProperties(adapterURI)
	if err != nil {
		logger.Fatal("", genlog.Any("Get Adapter properties error", err))
	}
	logger.Info("Get properties operation successfull")
	logger.Info("********* Adapter properties **************")
	logger.Info("\n- NAME: " + adapter.Name)
	logger.Info("\n- Status: " + string(adapter.Status))
	logger.Info("\n- State: " + string(adapter.State))
	logger.Info("\n- Detected Card Type: " + string(adapter.DetectedCardType))
	logger.Info("\n- Port Count: " + fmt.Sprint(adapter.PortCount))
	logger.Info("\n- Adapter Family: " + string(adapter.Family))
	logger.Info("\n- PHYSICAL CONNECTION STATUS: " + string(adapter.PHYSICALCHANNELSTATUS))
	for i, v := range adapter.NetworkAdapterPortURIs {
		logger.Info(fmt.Sprint(i+1) + " Network Port URI: " + v)
		networkProps, _, err := hmcManager.GetAdapterProperties(v)
		if err != nil {
			logger.Fatal("", genlog.Any("Get Network Port properties error", err))
		}
		logger.Info("Network Port Name: " + networkProps.Name)
		logger.Info("Network Port Description: " + networkProps.Description)

	}
	logger.Info("*********************************************")

}

func CreatePartition(hmcManager zhmcclient.ZhmcAPI) {
	cpcURI := GetCpcURL(hmcManager)
	parName := os.Getenv("PAR_NAME")
	iflProcessor, _ := strconv.Atoi(os.Getenv("IFLs"))
	initialMemory, _ := strconv.Atoi(os.Getenv("INITIAL_MEMORY"))
	maxMemory, _ := strconv.Atoi(os.Getenv("MAX_MEMORY"))
	processorType := os.Getenv("PROCESSOR_TYPE")
	var processorMode zhmcclient.PartitionProcessorMode = ""
	if processorType == "shared" {
		processorMode = zhmcclient.PROCESSOR_MODE_SHARED
	} else {
		processorMode = zhmcclient.PROCESSOR_MODE_DEDICATED
	}
	props := &zhmcclient.LparProperties{
		Name:          parName,
		IflProcessors: iflProcessor,
		InitialMemory: initialMemory,
		MaximumMemory: maxMemory,
		ProcessorMode: processorMode,
	}
	_, _, err := hmcManager.CreateLPAR(cpcURI, props)
	if err != nil {
		logger.Fatal("", genlog.Any("Create Partition error", err))
	}
	logger.Info("Create partition successful")
}

func StopPartitionforHmc(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	_, _, err := hmcManager.StopLPAR(lparURI)
	if err != nil {
		logger.Fatal("", genlog.Any("Stop Partition error", err))
	}
	logger.Info("Stop partition successful")
}

func DeletePartitionforHmc(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	_, err := hmcManager.DeleteLPAR(lparURI)
	if err != nil {
		logger.Fatal("", genlog.Any("Delete Partition error", err))
	}
	logger.Info("Delete partition successfull")
}

func MountIsoImageToPartition(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	isofile := os.Getenv("ISO_FILE")
	insfile := os.Getenv("INS_FILE")
	_, err := hmcManager.MountIsoImage(lparURI, isofile, insfile)
	if err != nil {
		logger.Fatal("", genlog.Any("Mount iso error: ", err))
	}
	logger.Info("Mount iso image successful")
}

func UnmountIsoImageToPartition(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	_, err := hmcManager.UnmountIsoImage(lparURI)
	if err != nil {
		logger.Fatal("", genlog.Any("Unmount iso error", err))
	}
	logger.Info("Unmount iso image successful")
}

func UpdateBootDeviceProperty(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	var bootDevice zhmcclient.PartitionBootDevice = zhmcclient.BOOT_DEVICE_ISO_IMAGE
	props := &zhmcclient.LparProperties{BootDevice: bootDevice}
	_, err := hmcManager.UpdateLparProperties(lparURI, props)
	if err != nil {
		logger.Fatal("", genlog.Any("Update boot device error", err))
	}
	logger.Info("Update boot device successful")
}

func StartPartitionforHmc(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	_, _, err := hmcManager.StartLPAR(lparURI)
	if err != nil {
		logger.Fatal("", genlog.Any("Stop Partition error", err))
	}
	logger.Info("Start partition successfull")
}

func FetchASCIIConsoleURI(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	props := &zhmcclient.AsciiConsoleURIPayload{}

	response, _, err := hmcManager.FetchAsciiConsoleURI(lparURI, props)
	if err != nil {
		logger.Fatal("", genlog.Any("Fetch Ascii Console URI Error", err))
	}
	logger.Info("The URI to access the ASCII Console is :" + response.URI)
	logger.Info("The sessionID for the ASCII Console is :" + response.SessionID)
}

/*
List all function
*/
func ListStorageGroupsforCPC(hmcManager zhmcclient.ZhmcAPI) {
	cpcID := os.Getenv("CPC_ID")
	storageGroupURI := "api/storage-groups/"
	storageGroups, _, err := hmcManager.ListStorageGroups(storageGroupURI, "/api/cpcs/"+cpcID)
	if err != nil {
		logger.Error("", genlog.Any("List Storage Group Error", err))
	}
	for _, sg := range storageGroups {
		logger.Info("########################################")
		logger.Info("Storage group Name: " + sg.Name)
		logger.Info("Storage group URI: " + sg.ObjectURI)
		logger.Info("Storage group TYPE: " + sg.Type)
		logger.Info("Storage group Fullfillment state: " + string(sg.FulfillmentState))
		sgroup, _, _ := hmcManager.GetStorageGroupProperties(sg.ObjectURI)
		logger.Info("Storage Group Properties")
		logger.Info("", genlog.Any("Storage group unassigned wwpns", sgroup.UnAssignedWWPNs))
		logger.Info("", genlog.Any("  - Storage Group Volumes", sgroup.StorageVolumesURIs))
		logger.Info("", genlog.Any("  - Storage Group ObjectID", sgroup.ObjectID))
	}
	logger.Info("########################################")
}

func ListStorageVolumesforCPC(hmcManager zhmcclient.ZhmcAPI) {
	sgroupID := os.Getenv("SGROUP_ID")
	storageGroupURI := "/api/storage-groups/" + sgroupID + "/storage-volumes"
	storageVolumes, _, err := hmcManager.ListStorageVolumes(storageGroupURI)
	if err != nil {
		logger.Error("", genlog.Any("List Storage Group Error: ", err))
	}
	for _, sv := range storageVolumes {
		logger.Info("########################################")
		logger.Info("Storage Volume Name: " + sv.Name)
		logger.Info("", genlog.Any("Storage Volume Fullfillment state", sv.FulfillmentState))
		storageVolume, _, volErr := hmcManager.GetStorageVolumeProperties(sv.URI)
		if volErr != nil {
			logger.Fatal(volErr.Message)
		}
		logger.Info("Storage Volume Properties")
		logger.Info("  - Storage Volume ECKD Type: " + storageVolume.EckdType)
		logger.Info("  - Storage Volume Active Size: " + fmt.Sprint(storageVolume.ActiveSize))
		logger.Info("  - Storage Volume Device Number: " + storageVolume.DeviceNumber)
		logger.Info("  - Storage Volume Path Information: ")
		for index, path := range storageVolume.Paths {
			logger.Info(" " + fmt.Sprint(index+1) + "*****************************************")
			logger.Info("\tPath Device Number: " + path.DeviceNumber)
			logger.Info("\n\tPath PartitionURI: " + path.PartitionURI)
			logger.Info("\n\tPath LUN: " + path.LogicalUnitNumber)
			logger.Info("\n\tPath Target WWPN: " + path.TargetWWPN)
		}
	}
	logger.Info("########################################")
}

func AttachStorageGroupToPartitionofCPC(hmcManager zhmcclient.ZhmcAPI) {
	sgroupID := os.Getenv("SGROUP_ID")
	props := &zhmcclient.StorageGroupPayload{StorageGroupURI: "/api/storage-groups/" + sgroupID}
	_, err := hmcManager.AttachStorageGroupToPartition(GetLPARURI(), props)
	if err != nil {
		logger.Fatal("", genlog.Any("Attach storage group error", err))
	}
	logger.Info("Attach storage group operation successful")
}

func DetachStorageGroupToPartitionofCPC(hmcManager zhmcclient.ZhmcAPI) {
	sgroupID := os.Getenv("SGROUP_ID")
	props := &zhmcclient.StorageGroupPayload{StorageGroupURI: "/api/storage-groups/" + sgroupID}
	_, err := hmcManager.DetachStorageGroupToPartition(GetLPARURI(), props)
	if err != nil {
		logger.Fatal("", genlog.Any("Detach storage group error", err))
	}
	logger.Info("Detach storage group operation successful")
}

func CreateStorageGroupsforCPC(hmcManager zhmcclient.ZhmcAPI) {
	cpcURI := GetCpcURL(hmcManager)
	uri := "/api/storage-groups/"
	sgname := os.Getenv("SG_NAME")
	sgtype := os.Getenv("SG_TYPE")
	svusage := zhmcclient.BOOT_USAGE
	svname := os.Getenv("SVOLUME_NAME")
	svsize := os.Getenv("SVOLUME_SIZE")
	volsize, _ := strconv.ParseFloat(svsize, 64)
	svoperation := zhmcclient.STORAGE_VOLUME_CREATE
	sv := zhmcclient.StorageVolume{
		Operation: svoperation,
		Usage:     svusage,
		Size:      volsize,
		Name:      svname,
	}

	props := &zhmcclient.CreateStorageGroupProperties{
		CpcURI:         cpcURI,
		Name:           sgname,
		Type:           sgtype,
		StorageVolumes: []zhmcclient.StorageVolume{sv},
	}
	storagegroup, _, err := hmcManager.CreateStorageGroups(uri, props)
	if err != nil {
		logger.Fatal("", genlog.Any("Create StorageGroup error", err))
	}

	logger.Info("Create StorageGroup successful")
	logger.Info("storageGroup" + storagegroup.ObjectURI)

	for _, v := range storagegroup.URI {
		logger.Info("Storage Group URI:" + v)
		logger.Info("storage Group ObjectURI" + storagegroup.ObjectURI)

		storageVolume, _, volErr := hmcManager.GetStorageVolumeProperties(v)
		if volErr != nil {
			logger.Fatal(volErr.Message)
		}
		logger.Info("Storage Volume Properties")

		logger.Info("  - Storage Volume Active Size: " + fmt.Sprint(storageVolume.ActiveSize))
		logger.Info("  - Storage Volume Device Number: " + storageVolume.DeviceNumber)
		logger.Info("  - Storage Volume Path Information: ")
		for index, path := range storageVolume.Paths {
			logger.Info(" " + fmt.Sprint(index+1) + "*****************************************")
			logger.Info("\tPath Device Number: " + path.DeviceNumber)
			logger.Info("\n\tPath PartitionURI: " + path.PartitionURI)
			logger.Info("\n\tPath LUN: " + path.LogicalUnitNumber)
			logger.Info("\n\tPath Target WWPN: " + path.TargetWWPN)
		}
	}
	logger.Info("########################################")
}

func GetStorageGroupPropertiesforCPC(hmcManager zhmcclient.ZhmcAPI) {
	sgroupID := os.Getenv("SGROUP_ID")
	storageGroupURI := "api/storage-groups/" + sgroupID
	storagegroup, status, err := hmcManager.GetStorageGroupProperties(storageGroupURI)
	logger.Error("error message",
		genlog.String("status", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", err)))
	if err != nil {
		logger.Fatal("", genlog.Any("Get Storage Group Properties error", err))
	}
	logger.Info("Get Storage Group Properties successful")
	logger.Info("storaage group object uri" + storagegroup.ObjectURI)
	logger.Info("direct connect count" + fmt.Sprintln(storagegroup.DirectConnectionCount))
	logger.Info("max-partitions" + fmt.Sprintln(storagegroup.MaxPartitions))
	logger.Info("connectivity" + fmt.Sprintln(storagegroup.Connectivity))
}
func GetStorageGroupPartitionsforCPC(hmcManager zhmcclient.ZhmcAPI) {
	query := map[string]string{}
	sgroupID := os.Getenv("SGROUP_ID")
	storageGroupURI := "api/storage-groups/" + sgroupID

	storageGroupPartitions, status, err := hmcManager.GetStorageGroupPartitions(storageGroupURI, query)
	logger.Error("error message",
		genlog.String("status", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", err)))
	if err != nil {
		logger.Fatal("", genlog.Any("Get Storage Group Partitions error", err))
	}
	logger.Info("Get Storage Group Partitions successful")

	for _, v := range storageGroupPartitions.GetStorageGroups {
		logger.Info("LPAR Name" + v.Name)
		logger.Info("Storage Group ObjectURI" + v.URI)
		logger.Info("Storage Group Status:  " + string(v.Status))

	}
}

func DeleteStorageGroupforCPC(hmcManager zhmcclient.ZhmcAPI) {
	sgroupID := os.Getenv("SGROUP_ID")
	storageGroupURI := "api/storage-groups/" + sgroupID
	_, err := hmcManager.DeleteStorageGroup(storageGroupURI)
	if err != nil {
		logger.Fatal("", genlog.Any("Delete StorageGroup error", err))
	}
	logger.Info("Delete StorageGroup successfull")
}

/*
List all function
*/
func ListAll(hmcManager zhmcclient.ZhmcAPI) {
	query := map[string]string{}
	cpcs, _, err := hmcManager.ListCPCs(query)
	if err != nil {
		logger.Error("Error: " + err.Message)
	} else {
		for _, cpc := range cpcs {
			logger.Info("########################################")
			logger.Info("cpc name: " + cpc.Name)
			logger.Info("cpc uri: " + cpc.URI)

			adapters, _, err := hmcManager.ListAdapters(cpc.URI, query)
			if err != nil {
				logger.Error("Error: " + err.Message)
			} else {
				logger.Info("-----------------------")
				for _, adapter := range adapters {
					logger.Info("++++++++++++++++++++++++")
					logger.Info("", genlog.Any("adapter properties", adapter))
				}
			}

			lpars, _, err := hmcManager.ListLPARs(cpc.URI, query)
			if err != nil {
				logger.Error("Error: " + err.Message)
			} else {
				logger.Info("-----------------------")
				for _, lpar := range lpars {
					logger.Info("++++++++++++++++++++++++")
					logger.Info("lpar name: " + lpar.Name)
					logger.Info("lpar uri: " + lpar.URI)

					props, _, err := hmcManager.GetLparProperties(lpar.URI)
					if err != nil {
						logger.Error("Error getting lpart properties", genlog.Error(errors.New(err.Message)))
					} else {
						logger.Info("", genlog.Any("lpar properties", props))
					}

					logger.Info("++++++++++++++++++++++++")
					nics, _, err := hmcManager.ListNics(lpar.URI)
					if err != nil {
						logger.Error("Error listing nics ", genlog.Error(errors.New(err.Message)))
					} else {
						logger.Info("", genlog.Any("nics list", nics))
						for _, nicURI := range nics {
							nic, _, err := hmcManager.GetNicProperties(nicURI)
							if err != nil {
								logger.Error("Error getting nic properties", genlog.Error(errors.New(err.Message)))
							} else {
								logger.Info("", genlog.Any("nic properties", nic))
							}
						}
					}
				}
			}
		}
	}
}
