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

import (
	"runtime"
	"time"
)

// Global constants.
const (
	libraryName            = "zhmcclient"
	libraryVersion         = "v0.1"
	libraryUserAgentPrefix = "ZHMC (" + runtime.GOOS + "; " + runtime.GOARCH + ") "
	libraryUserAgent       = libraryUserAgentPrefix + libraryName + "/" + libraryVersion

	HttpClientTimeout = 10 * time.Second
	HandshakeTimeout  = 5 * time.Second
	DialTimeout       = 5 * time.Second
)
