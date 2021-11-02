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
	endpoint := os.Getenv("HMC_ENDPOINT") // "https://192.168.195.118:9955"
	username := os.Getenv("HMC_USERNAME")
	password := os.Getenv("HMC_PASSWORD")
	creds := &zhmcclient.Options{
		Username:   username,
		Password:   password,
		VerifyCert: false,
		Trace:      true,
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
		fmt.Println("client initiliazed.")
		hmcManager := zhmcclient.NewManagerFromClient(client)

		query := map[string]string{}
		cpcs, err := hmcManager.ListCPCs(query)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		} else {
			for _, cpc := range cpcs {
				fmt.Println("########################################")
				fmt.Println("cpc name: ", cpc.Name)
				fmt.Println("cpc id: ", cpc.URI)

				adapters, err := hmcManager.ListAdapters(cpc.URI, query)
				if err != nil {
					fmt.Println("Error: ", err.Error())
				} else {
					fmt.Println("-----------------------")
					for _, adapter := range adapters {
						fmt.Println("--")
						fmt.Println("adapter properties: ", adapter)
					}
				}

				lpars, err := hmcManager.ListLPARs(cpc.URI, query)
				if err != nil {
					fmt.Println("Error: ", err.Error())
				} else {
					fmt.Println("-----------------------")
					for _, lpar := range lpars {
						fmt.Println("--")
						fmt.Println("lpar name: ", lpar.Name)
						fmt.Println("lpar id: ", lpar.URI)

						props, err := hmcManager.GetLparProperties(lpar.URI)
						if err != nil {
							fmt.Println("Error: ", err.Error())
						} else {
							fmt.Println("lpar properties: ", props)
						}
					}
				}
			}
		}
	}
}
