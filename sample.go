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
	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

func main() {
	endpoint := "https://192.168.195.118:9955"
	creds := &zhmcclient.Options{
		Username:   "name",
		Password:   "psw",
		VerifyCert: false,
		Trace:      false,
	}
	client, _ := zhmcclient.NewClient(endpoint, creds)
	if client != nil {
		hmcManager := zhmcclient.NewManagerFromClient(client)
		hmcManager.ListCPCs()
	}
}
