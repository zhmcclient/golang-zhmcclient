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

package main

import (
	"fmt"
	"os"

	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

func main() {
	endpoint := os.Getenv("HMC_ENDPOINT") // "https://9.114.87.7:6794/", "https://192.168.195.118:6794"
	username := os.Getenv("HMC_USERNAME")
	password := os.Getenv("HMC_PASSWORD")
	args := os.Args[1:]
	//partitionId := os.Getenv("PAR_ID")
	// isofile := os.Getenv("ISO_FILE")
	// insfile := os.Getenv("INS_FILE")
	creds := &zhmcclient.Options{Username: username, Password: password, SkipCert: true, Trace: false}
	if endpoint == "" || username == "" || password == "" {
		fmt.Println("Please set HMC_ENDPOINT, HMC_USERNAME and HMC_PASSWORD")
		os.Exit(1)
	}
	if len(args) == 0 {
		fmt.Println(`
		    	Usage: sample <Command>

			Please enter one of the below Command:

				"StartPartitionforHmc":
					- Starts the partition for the selected HMC
				
				"StopPartitionforHmc":
					- Stops the partition for the selected HMC
				
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

		}`)
		os.Exit(1)
	} else {
		fmt.Println("HMC_ENDPOINT: ", endpoint)
		fmt.Println("HMC_USERNAME: ", username)
		fmt.Println("HMC_PASSWORD: xxxxxx")
		client, err := zhmcclient.NewClient(endpoint, creds)
		if err != nil {
			fmt.Println("Error: ", err.Message)
		}
		if client != nil {
			fmt.Println("client initialized.")
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
				case "StartPartitionforHmc":
					StartPartitionforHmc(hmcManager)
				case "StopPartitionforHmc":
					StopPartitionforHmc(hmcManager)
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

func ListAdaptersofCPC(hmcManager zhmcclient.ZhmcAPI) {
	query := map[string]string{}
	CPCURI := "api/cpcs/" + os.Getenv("CPC_ID")
	adapters,
		err := hmcManager.ListAdapters(CPCURI, query)

	if err != nil {
		fmt.Println("List Adapters error: ", err.Message)
		os.Exit(1)
	} else {
		fmt.Println("-----------------------")
		for _, adapter := range adapters {
			fmt.Println("++++++++++++++++++++++++")
			fmt.Println("Adapter Name:", adapter.Name)
			fmt.Println("Adapter Type:", adapter.Type)
			fmt.Println("Adapter Family:", adapter.Family)
			fmt.Println("Adapter Status:", adapter.Status)
			fmt.Println("Adapter URI:", adapter.URI)
			adapter, _ := hmcManager.GetAdapterProperties(adapter.URI)
			fmt.Println("********* Adapter properties **************")
			fmt.Println("\n- NAME: " + adapter.Name)
			fmt.Println("\n- Status: " + adapter.Status)
			fmt.Println("\n- State: " + adapter.State)
			fmt.Println("\n- Detected Card Type: " + adapter.DetectedCardType)
			fmt.Println("\n- Port Count: ", adapter.PortCount)
			fmt.Println("\n- Adapter Family: " + adapter.Family)
			fmt.Println("\n- PHYSICAL CONNECTION STATUS: " + adapter.PHYSICALCHANNELSTATUS)
			fmt.Println("\n- Network Port URIS: ", adapter.NetworkAdapterPortURIs)
			fmt.Println("\n- Storage Port URIS: ", adapter.StoragePortURIs)
			fmt.Println("*********************************************")
		}
		fmt.Println("\n-----------------------")
	}
}

func GetAdapterPropsforCPC(hmcManager zhmcclient.ZhmcAPI) {
	adapterID := os.Getenv("ADAPTER_ID")
	adapterURI := "api/adapters/" + adapterID
	adapter, err := hmcManager.GetAdapterProperties(adapterURI)
	if err != nil {
		fmt.Println("Get Adapter properties error: ", err.Message)
		os.Exit(1)
	}
	fmt.Println("Get properties operation successfull")
	fmt.Println("********* Adapter properties **************")
	fmt.Println("\n- NAME: " + adapter.Name)
	fmt.Println("\n- Status: " + adapter.Status)
	fmt.Println("\n- State: " + adapter.State)
	fmt.Println("\n- Detected Card Type: " + adapter.DetectedCardType)
	fmt.Println("\n- Port Count: ", adapter.PortCount)
	fmt.Println("\n- Adapter Family: " + adapter.Family)
	fmt.Println("\n- PHYSICAL CONNECTION STATUS: " + adapter.PHYSICALCHANNELSTATUS)
	for i, v := range adapter.NetworkAdapterPortURIs {
		fmt.Println(i+1, " Network Port URI: ", v)
		networkProps, err := hmcManager.GetAdapterProperties(v)
		if err != nil {
			fmt.Println("Get Network Port properties error: ", err.Message)
			os.Exit(1)
		}
		fmt.Println("Network Port Name: ", networkProps.Name)
		fmt.Println("Network Port Description: ", networkProps.Description)

	}
	fmt.Println("*********************************************")

}

func StopPartitionforHmc(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	_, err := hmcManager.StopLPAR(lparURI)
	if err != nil {
		fmt.Println("Stop Partition error: ", err.Message)
		os.Exit(1)
	}
	fmt.Println("Stop partition successfull")
}

func MountIsoImageToPartition(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	isofile := os.Getenv("ISO_FILE")
	insfile := os.Getenv("INS_FILE")
	err := hmcManager.MountIsoImage(lparURI, isofile, insfile)
	if err != nil {
		fmt.Println("Mount iso error: ", err.Message)
		os.Exit(1)
	}
	fmt.Println("Mount iso image successfull")
}

func UnmountIsoImageToPartition(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	err := hmcManager.UnmountIsoImage(lparURI)
	if err != nil {
		fmt.Println("Unmount iso error: ", err.Message)
		os.Exit(1)
	}
	fmt.Println("Unmount iso image successfull")
}

func UpdateBootDeviceProperty(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	var bootDevice zhmcclient.PartitionBootDevice = zhmcclient.BOOT_DEVICE_ISO_IMAGE
	props := &zhmcclient.LparProperties{BootDevice: bootDevice}
	err := hmcManager.UpdateLparProperties(lparURI, props)
	if err != nil {
		fmt.Println("Update boot device error: ", err.Message)
		os.Exit(1)
	}
	fmt.Println("Update boot device successfull")
}

func StartPartitionforHmc(hmcManager zhmcclient.ZhmcAPI) {
	lparURI := GetLPARURI()
	_, err := hmcManager.StartLPAR(lparURI)
	if err != nil {
		fmt.Println("Stop Partition error: ", err.Message)
		os.Exit(1)
	}
	fmt.Println("Start partition successfull")
}

/*
	List all function
*/
func ListStorageGroupsforCPC(hmcManager zhmcclient.ZhmcAPI) {
	cpcID := os.Getenv("CPC_ID")
	storageGroupURI := "api/storage-groups/"
	storageGroups, err := hmcManager.ListStorageGroups(storageGroupURI, "/api/cpcs/"+cpcID)
	if err != nil {
		fmt.Println("List Storage Group Error: ", err.Message)
	}
	for _, sg := range storageGroups {
		fmt.Println("########################################")
		fmt.Println("Storage group Name: ", sg.Name)
		fmt.Println("Storage group URI: ", sg.ObjectURI)
		fmt.Println("Storage group TYPE: ", sg.Type)
		fmt.Println("Storage group Fullfillment state: ", sg.FulfillmentState)
		sgroup, _ := hmcManager.GetStorageGroupProperties(sg.ObjectURI)
		fmt.Println("Storage Group Properties")
		fmt.Println("Storage group unassigned wwpns: ", sgroup.UnAssignedWWPNs)
		fmt.Println("  - Storage Group Volumes: ", sgroup.StorageVolumesURIs)
		fmt.Println("  - Storage Group ObjectID: ", sgroup.ObjectID)
	}
	fmt.Println("########################################")
}

func ListStorageVolumesforCPC(hmcManager zhmcclient.ZhmcAPI) {
	sgroupID := os.Getenv("SGROUP_ID")
	storageGroupURI := "/api/storage-groups/" + sgroupID + "/storage-volumes"
	storageVolumes, err := hmcManager.ListStorageVolumes(storageGroupURI)
	if err != nil {
		fmt.Println("List Storage Group Error: ", err.Message)
	}
	for _, sv := range storageVolumes {
		fmt.Println("########################################")
		fmt.Println("Storage Volume Name: ", sv.Name)
		fmt.Println("Storage Volume Fullfillment state: ", sv.FulfillmentState)
		fmt.Println("Storage volume usage: ", sv.Usage)
		storageVolume, volErr := hmcManager.GetStorageVolumeProperties(sv.URI)
		if volErr != nil {
			fmt.Println(volErr.Message)
			os.Exit(1)
		}
		fmt.Println("Storage Volume Properties")
		fmt.Println("  - Storage Volume ECKD Type: ", storageVolume.EckdType)
		fmt.Println("  - Storage Volume Active Size: ", storageVolume.ActiveSize)
		fmt.Println("  - Storage Volume Device Number: ", storageVolume.DeviceNumber)
		fmt.Println("  - Storage Volume Path Information: ")
		for index, path := range storageVolume.Paths {
			fmt.Println(" ", (index + 1), "*****************************************")
			fmt.Println("\tPath Device Number: ", path.DeviceNumber)
			fmt.Println("\n\tPath PartitionURI: ", path.PartitionURI)
			fmt.Println("\n\tPath LUN: ", path.LogicalUnitNumber)
			fmt.Println("\n\tPath Target WWPN: ", path.TargetWWPN)
		}
	}
	fmt.Println("########################################")
}

func AttachStorageGroupToPartitionofCPC(hmcManager zhmcclient.ZhmcAPI) {
	sgroupID := os.Getenv("SGROUP_ID")
	props := &zhmcclient.StorageGroupPayload{StorageGroupURI: "/api/storage-groups/" + sgroupID}
	err := hmcManager.AttachStorageGroupToPartition(GetLPARURI(), props)
	if err != nil {
		fmt.Println("Attach storage group error: ", err.Message)
		os.Exit(1)
	}
	fmt.Println("Attach storage group operation successfull")
}

func DetachStorageGroupToPartitionofCPC(hmcManager zhmcclient.ZhmcAPI) {
	sgroupID := os.Getenv("SGROUP_ID")
	props := &zhmcclient.StorageGroupPayload{StorageGroupURI: "/api/storage-groups/" + sgroupID}
	err := hmcManager.DetachStorageGroupToPartition(GetLPARURI(), props)
	if err != nil {
		fmt.Println("Detach storage group error: ", err.Message)
		os.Exit(1)
	}
	fmt.Println("Detach storage group operation successfull")
}

/*
	List all function
*/
func ListAll(hmcManager zhmcclient.ZhmcAPI) {
	query := map[string]string{}
	cpcs,
		err := hmcManager.ListCPCs(query)
	if err != nil {
		fmt.Println("Error: ", err.Message)
	} else {
		for _, cpc := range cpcs {
			fmt.Println("########################################")
			fmt.Println("cpc name: ", cpc.Name)
			fmt.Println("cpc uri: ", cpc.URI)

			adapters,
				err := hmcManager.ListAdapters(cpc.URI, query)
			if err != nil {
				fmt.Println("Error: ", err.Message)
			} else {
				fmt.Println("-----------------------")
				for _, adapter := range adapters {
					fmt.Println("++++++++++++++++++++++++")
					fmt.Println("adapter properties: ", adapter)
				}
			}

			lpars,
				err := hmcManager.ListLPARs(cpc.URI, query)
			if err != nil {
				fmt.Println("Error: ", err.Message)
			} else {
				fmt.Println("-----------------------")
				for _, lpar := range lpars {
					fmt.Println("++++++++++++++++++++++++")
					fmt.Println("lpar name: ", lpar.Name)
					fmt.Println("lpar uri: ", lpar.URI)

					props,
						err := hmcManager.GetLparProperties(lpar.URI)
					if err != nil {
						fmt.Println("Error: ", err.Message)
					} else {
						fmt.Println("lpar properties: ", props)
					}

					fmt.Println("++++++++++++++++++++++++")
					nics,
						err := hmcManager.ListNics(lpar.URI)
					if err != nil {
						fmt.Println("Error: ", err.Message)
					} else {
						fmt.Println("nics list: ", nics)
						for _, nicURI := range nics {
							nic,
								err := hmcManager.GetNicProperties(nicURI)
							if err != nil {
								fmt.Println("Error: ", err.Message)
							} else {
								fmt.Println("nic properties: ", nic)
							}
						}
					}
				}
			}
		}
	}
}
