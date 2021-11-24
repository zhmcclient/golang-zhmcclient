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
	"net/http"
	"runtime"
	"time"
)

// Global constants.
const (
	libraryName            = "zhmcclient"
	libraryVersion         = "v0.1"
	libraryUserAgentPrefix = "ZHMC (" + runtime.GOOS + "; " + runtime.GOARCH + ") "
	libraryUserAgent       = libraryUserAgentPrefix + libraryName + "/" + libraryVersion

	HMC_DEFAULT_PORT = "6794"

	DEFAULT_READ_RETRIES    = 0
	DEFAULT_CONNECT_RETRIES = 3

	DEFAULT_DIAL_TIMEOUT      = 10 * time.Second
	DEFAULT_HANDSHAKE_TIMEOUT = 10 * time.Second
	DEFAULT_CONNECT_TIMEOUT   = 30 * time.Second
	DEFAULT_READ_TIMEOUT      = 3600 * time.Second
	DEFAULT_MAX_REDIRECTS     = 30 * time.Second
	DEFAULT_OPERATION_TIMEOUT = 3600 * time.Second
	DEFAULT_STATUS_TIMEOUT    = 900 * time.Second

	APPLICATION_BODY_JSON         = "application/json"
	APPLICATION_BODY_OCTET_STREAM = "application/octet-stream"
)

// List of success status.
var KNOWN_SUCCESS_STATUS = []int{
	http.StatusOK,             // 200
	http.StatusCreated,        // 201
	http.StatusAccepted,       // 202
	http.StatusNoContent,      // 204
	http.StatusPartialContent, // 206
	http.StatusBadRequest,     // 400
	//http.StatusForbidden,    // 403, 403 susally caused by expired session header, we need handle it separately
	http.StatusNotFound,            // 404
	http.StatusConflict,            // 409
	http.StatusInternalServerError, // 500
	http.StatusServiceUnavailable,  // 503
}
