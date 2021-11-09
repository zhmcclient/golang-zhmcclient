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
	partitionId := os.Getenv("PAR_ID")
	creds := &zhmcclient.Options{
		Username: username,
		Password: password,
		SkipCert: true,
		Trace:    false,
	}
	if endpoint == "" || username == "" || password == "" {
		fmt.Println("Please set HMC_ENDPOINT, HMC_USERNAME and HMC_PASSWORD")
		os.Exit(1)
	}
	fmt.Println("HMC_ENDPOINT: ", endpoint)
	fmt.Println("HMC_USERNAME: ", username)
	fmt.Println("HMC_PASSWORD: xxxxxx")
	client, err := zhmcclient.NewClient(endpoint, creds)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	if client != nil {
		fmt.Println("client initialized.")
		hmcManager := zhmcclient.NewManagerFromClient(client)
		lparURI := "api/partitions/" + partitionId
		var bootDevice zhmcclient.PartitionBootDevice = zhmcclient.BOOT_DEVICE_ISO_IMAGE
		props := &zhmcclient.LparProperties{BootDevice: bootDevice}
		err := hmcManager.UpdateLparProperties(lparURI, props)
		if err != nil {
			fmt.Println("HMC Error: ", err.Error())
		} else {
			fmt.Println("Success: Boot device successfully updated")
		}
		_, startErr := hmcManager.StartLPAR(lparURI)
		if startErr != nil {
			fmt.Println("HMC Error: ", startErr.Error())
		} else {
			fmt.Println("Success: Partition successfully started")
		}
	}
}
