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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path"
	"strings"

	"net/http"
	"net/http/httputil"

	"net/url"
)

// ClientAPI defines an interface for issuing client requests to ZHMC
//go:generate counterfeiter -o fakes/client.go --fake-name ClientAPI . ClientAPI
type ClientAPI interface {
	GetEndpointURL() *url.URL
	TraceOn(outputStream io.Writer)
	TraceOff()
	SetCertVerify(isVerify bool)
	Logon() error
	Logoff() error
	IsLogon(verify bool) bool
	ExecuteRequest(requestType string, url string, requestData interface{}) (responseStatusCode int, responseBodyStream []byte, err error)
}

// HTTP Client interface required for unit tests
type HTTPClient interface {
	Do(request *http.Request) (*http.Response, error)
}

const (
	SESSION_HEADER_NAME = "X-API-Session"
)

type Options struct {
	Username   string
	Password   string
	Trace      bool
	VerifyCert bool
}

type LogonData struct {
	Userid   string `json:"userid"`
	Password string `json:"password"`
}

type Session struct {
	SessionID    string `json:"api-session"`
	ObjectTopic  string `json:"notification-topic"`
	JobTopic     string `json:"job-notification-topic"`
	Credential   string `json:"session-credential"`
	MajorVersion string `json:"api-major-version"`
	MinorVersion string `json:"api-minor-version"`
}

type Client struct {
	endpointURL *url.URL
	httpClient  *http.Client
	logondata   *LogonData
	session     *Session

	isVerifyCert   bool
	isTraceEnabled bool
	traceOutput    io.Writer
}

func NewClient(endpoint string, opts *Options) (ClientAPI, error) {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: DialTimeout,
		}).Dial,
		TLSHandshakeTimeout: HandshakeTimeout,
	}

	httpclient := &http.Client{
		Timeout:   HttpClientTimeout,
		Transport: netTransport,
	}

	endpointurl, err := getEndpointURLFromString(endpoint)
	if err != nil {
		return nil, err
	}

	client := &Client{
		endpointURL: endpointurl,
		httpClient:  httpclient,
		logondata: &LogonData{
			Userid:   opts.Username,
			Password: opts.Password,
		},
	}

	err = client.Logon()
	if err != nil {
		return nil, err
	}

	if opts.Trace {
		client.TraceOn(nil)
	} else {
		client.TraceOff()
	}

	client.SetCertVerify(opts.VerifyCert)

	return client, nil
}

// TODO, validate the endpoint
func getEndpointURLFromString(endpoint string) (*url.URL, error) {
	return url.Parse(endpoint)
}

func (c *Client) GetEndpointURL() *url.URL {
	return c.endpointURL
}

func (c *Client) TraceOn(outputStream io.Writer) {
	if outputStream == nil {
		outputStream = os.Stdout
	}
	c.traceOutput = outputStream
	c.isTraceEnabled = true
}

func (c *Client) TraceOff() {
	c.traceOutput = os.Stdout
	c.isTraceEnabled = false
}

// TODO, check cert when request/logon
func (c *Client) SetCertVerify(isVerify bool) {
	c.isVerifyCert = isVerify
}

func (c *Client) clearSession() {
	c.session = nil
}

func (c *Client) Logon() error {
	c.clearSession()
	logonUri := path.Join(c.endpointURL.Path, "/api/sessions")
	status, responseBody, err := c.executeMethod(http.MethodPost, logonUri, c.logondata)
	if err != nil {
		return err
	}

	if status == http.StatusOK || status == http.StatusCreated {
		err = json.Unmarshal(responseBody, &c.session)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("Logon -- Unknown Error")
}

func (c *Client) Logoff() error {
	logoffUri := path.Join(c.endpointURL.Path, "/api/sessions/this-session")
	status, _, err := c.executeMethod(http.MethodDelete, logoffUri, nil)
	if status == http.StatusNoContent {
		c.clearSession()
		return nil
	}
	return err
}

func (c *Client) IsLogon(verify bool) bool {
	if verify {
		testUri := path.Join(c.endpointURL.Path, "/api/console")
		status, _, err := c.executeMethod(http.MethodGet, testUri, nil)
		if err != nil {
			return false
		} else if status == http.StatusOK || status == http.StatusBadRequest {
			return true
		}
		return false
	}

	if c.session != nil && c.session.SessionID != "" {
		return true
	}
	return false
}

func (c *Client) setUserAgent(req *http.Request) {
	req.Header.Set("User-Agent", libraryUserAgent)
}

// TODO, add "Content-Type" according to requestBody
func (c *Client) setRequestHeaders(req *http.Request) {
	c.setUserAgent(req)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(SESSION_HEADER_NAME, c.session.SessionID)
}

func (c *Client) ExecuteRequest(requestType string, url string, requestData interface{}) (responseStatusCode int, responseBodyStream []byte, err error) {
	return c.executeMethod(requestType, url, requestData)
}

// TODO, 1. Retry, 2. Logon when 401/403
func (c *Client) executeMethod(requestType string, url string, requestData interface{}) (responseStatusCode int, responseBodyStream []byte, err error) {
	var requestBody []byte

	if requestData != nil {
		requestBody, err = json.Marshal(requestData)
		if err != nil {
			return -1, nil, err
		}
	}

	request, err := http.NewRequest(requestType, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return -1, nil, err
	}

	c.setRequestHeaders(request)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return -1, nil, err
	}

	defer response.Body.Close()

	responseBodyStream, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return -1, nil, err
	}

	err = c.checkResponseStatus(response.StatusCode)

	if c.isTraceEnabled {
		err = c.traceHTTP(request, response)
		if err != nil {
			return response.StatusCode, nil, err
		}
	}

	return response.StatusCode, responseBodyStream, err
}

func (c *Client) checkResponseStatus(responseStatusCode int) error {
	if responseStatusCode != http.StatusAccepted && responseStatusCode != http.StatusOK && responseStatusCode != http.StatusNoContent {
		errorText := fmt.Sprintf("Invalid Status Code[%d]", responseStatusCode)
		return errors.New(errorText)
	}

	return nil
}

func (c *Client) traceHTTP(req *http.Request, resp *http.Response) error {
	_, err := fmt.Fprintln(c.traceOutput, "---------START-HTTP---------")
	if err != nil {
		return err
	}

	reqTrace, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(c.traceOutput, string(reqTrace))
	if err != nil {
		return err
	}

	var respTrace []byte
	if resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusPartialContent &&
		resp.StatusCode != http.StatusNoContent {
		respTrace, err = httputil.DumpResponse(resp, true)
		if err != nil {
			return err
		}
	} else {
		respTrace, err = httputil.DumpResponse(resp, false)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(c.traceOutput, strings.TrimSuffix(string(respTrace), "\r\n"))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(c.traceOutput, "---------END-HTTP---------")
	if err != nil {
		return err
	}

	return nil
}
