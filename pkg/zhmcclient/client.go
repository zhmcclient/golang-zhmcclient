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
	//"net"
	"net/http"
	//"net/http/cookiejar"
	//"net/http/httputil"
	"net/url"

)

type Client struct {
	endpointURL	*url.URL
	httpClient	*http.Client
	sessionID	string
	objectTopic string
	jobTopic	string
}

type Options struct {
	Username 	string
	Password 	string
	VerifyCert  bool
	Trace		bool
}

func NewClient(endpoint string, opts *Options) (*Client, error) {
	return nil, nil
}

func (c *Client) setUserAgent(req *http.Request) {
	req.Header.Set("User-Agent", libraryUserAgent)
}