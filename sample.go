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
	//partitionId := os.Getenv("PAR_ID")
	// isofile := os.Getenv("ISO_FILE")
	// insfile := os.Getenv("INS_FILE")
	creds := &zhmcclient.Options{Username: username, Password: password, SkipCert: true, Trace: false}
	if endpoint == "" || username == "" || password == "" {
		fmt.Println("Please set HMC_ENDPOINT, HMC_USERNAME and HMC_PASSWORD")
		os.Exit(1)
	}
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
		*/
		/* #List All usage
		- List all LPAR's for the selected HMC endpoint
		*/
		ListAll(hmcManager)
	}
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
