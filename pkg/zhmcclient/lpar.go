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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"

	"github.ibm.com/genctl/shared-logger/genlog"
)

// LparAPI defines an interface for issuing LPAR requests to ZHMC
//go:generate counterfeiter -o fakes/lpar.go --fake-name LparAPI . LparAPI
type LparAPI interface {
	ListLPARs(cpcURI string, query map[string]string) ([]LPAR, int, *HmcError)
	GetLparProperties(lparURI string) (*LparProperties, int, *HmcError)
	UpdateLparProperties(lparURI string, props *LparProperties) (int, *HmcError)
	StartLPAR(lparURI string) (string, int, *HmcError)
	StopLPAR(lparURI string) (string, int, *HmcError)
	AttachStorageGroupToPartition(storageGroupURI string, request *StorageGroupPayload) (int, *HmcError)
	DetachStorageGroupToPartition(storageGroupURI string, request *StorageGroupPayload) (int, *HmcError)
	MountIsoImage(lparURI string, isoFile string, insFile string) (int, *HmcError)
	UnmountIsoImage(lparURI string) (int, *HmcError)

	ListNics(lparURI string) ([]string, int, *HmcError)
	FetchAsciiConsoleURI(lparURI string, request *AsciiConsoleURIPayload) (*AsciiConsoleURIResponse, int, *HmcError)
}

type LparManager struct {
	client ClientAPI
}

func NewLparManager(client ClientAPI) *LparManager {
	return &LparManager{
		client: client,
	}
}

/**
* GET /api/cpcs/{cpc-id}/partitions
* @cpcURI is the cpc object-uri
* @query is a key, value pairs array,
*        currently, supports 'name=$name_reg_expression'
*                            'status=PartitionStatus'
*                            'type=PartitionType'
* @return lpar array
* Return: 200 and LPARs array
*     or: 400, 404, 409
 */
func (m *LparManager) ListLPARs(cpcURI string, query map[string]string) ([]LPAR, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, cpcURI, "/partitions")
	requestUrl = BuildUrlFromQuery(requestUrl, query)

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodGet))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("error on getting lpar's",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return nil, status, err
	}

	if status == http.StatusOK {
		lpars := &LPARsArray{}
		err := json.Unmarshal(responseBody, lpars)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Response: status: %v", status))
		return lpars.LPARS, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error listing lpar's",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return nil, status, errorResponseBody
}

/**
* GET /api/partitions/{partition-id}
* @lparURI is the object-uri
* Return: 200 and LparProperties
*     or: 400, 404,
 */
func (m *LparManager) GetLparProperties(lparURI string) (*LparProperties, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI)

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodGet))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("error on getting lpar properties",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return nil, status, err
	}

	if status == http.StatusOK {
		lparProps := LparProperties{}
		err := json.Unmarshal(responseBody, &lparProps)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Response: request url: %v, status: %v, lpar properties: %v", requestUrl, status, lparProps))
		return &lparProps, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error getting lpar properties",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return nil, status, errorResponseBody
}

/**
* POST /api/partitions/{partition-id}
* @lparURI is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503,
 */
func (m *LparManager) UpdateLparProperties(lparURI string, props *LparProperties) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI)

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v, Parameters: %v", requestUrl, http.MethodPost, props))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, props, "")
	if err != nil {
		logger.Error("error on getting lpar properties",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return status, err
	}

	if status == http.StatusNoContent {
		logger.Info(fmt.Sprintf("update lpar properties completed, request url: %v, status: %v", requestUrl, status))
		return status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error updating lpar properties",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return status, errorResponseBody
}

/**
* POST /api/partitions/{partition-id}/operations/start
* @lparURI is the object-uri
* @return job-uri
* Return: 202 and job-uri
*     or: 400, 403, 404, 503,
 */
func (m *LparManager) StartLPAR(lparURI string) (string, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/start")

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodPost))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil, "")
	if err != nil {
		logger.Error("error on starting lpar",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return "", status, err
	}

	if status == http.StatusAccepted {
		responseObj := StartStopLparResponse{}
		err := json.Unmarshal(responseBody, &responseObj)
		if err != nil {
			return "", status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		if responseObj.URI != "" {
			logger.Info(fmt.Sprintf("Response: request url: %v, status: %v, lpar uri: %v", requestUrl, status, responseObj.URI))
			return responseObj.URI, status, nil
		}
		logger.Error("error on starting lpar",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(errors.New("empty job uri")))
		return "", status, getHmcErrorFromMsg(ERR_CODE_EMPTY_JOB_URI, ERR_MSG_EMPTY_JOB_URI)
	}

	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error starting lpar",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return "", status, errorResponseBody
}

/**
* POST /api/partitions/{partition-id}/operations/stop
* @lparURI is the object-uri
* @return job-uri
* Return: 202 and job-uri
*     or: 400, 403, 404, 503,
 */
func (m *LparManager) StopLPAR(lparURI string) (string, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/stop")

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodPost))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil, "")
	if err != nil {
		logger.Error("error on stopping lpar",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return "", status, err
	}

	if status == http.StatusAccepted {
		responseObj := StartStopLparResponse{}
		err := json.Unmarshal(responseBody, &responseObj)
		if err != nil {
			return "", status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		if responseObj.URI != "" {
			logger.Info(fmt.Sprintf("Response: request url: %v, status: %v, lpar uri: %v", requestUrl, status, responseObj.URI))
			return responseObj.URI, status, nil
		}
		logger.Error("error on stopping lpar",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(errors.New("empty job uri")))
		return "", status, getHmcErrorFromMsg(ERR_CODE_EMPTY_JOB_URI, ERR_MSG_EMPTY_JOB_URI)
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error stopping lpar",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return "", status, errorResponseBody
}

/**
* POST /api/partitions/{partition-id}/operations/mount-iso-image
* @lparURI is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *LparManager) MountIsoImage(lparURI string, isoFile string, insFile string) (int, *HmcError) {
	pureIsoName := path.Base(isoFile)
	pureInsName := path.Base(insFile)
	query := map[string]string{
		"image-name":    pureIsoName,
		"ins-file-name": "/" + pureInsName,
	}
	imageData, byteErr := RetrieveBytes(isoFile)
	if byteErr != nil {
		logger.Error("error on retrieving iso file", genlog.Error(byteErr))
	}
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/mount-iso-image")
	requestUrl = BuildUrlFromQuery(requestUrl, query)

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v, Parameters: iso file %v, insfile: %v", requestUrl, http.MethodPost, isoFile, insFile))

	status, responseBody, err := m.client.UploadRequest(http.MethodPost, requestUrl, imageData)
	if err != nil {
		logger.Error("error mounting iso image",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return status, err
	}

	if status == http.StatusNoContent {
		logger.Info(fmt.Sprintf("Response: mounting iso image completed, request url: %v, status: %v", requestUrl, status))
		return status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error mounting iso image",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return status, errorResponseBody
}

/**
* POST /api/partitions/{partition-id}/operations/unmount-iso-image
* @lparURI is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *LparManager) UnmountIsoImage(lparURI string) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/unmount-iso-image")

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodPost))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil, "")
	if err != nil {
		logger.Error("error unmounting iso image",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return status, err
	}

	if status == http.StatusNoContent {
		logger.Info(fmt.Sprintf("Response: iso image unmounted. request url: %v, status: %v", requestUrl, status))
		return status, nil
	}

	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error unmounting iso image",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return status, errorResponseBody
}

/**
* get_property('nic-uris') from LPAR
 */
func (m *LparManager) ListNics(lparURI string) ([]string, int, *HmcError) {
	props, status, err := m.GetLparProperties(lparURI)
	if err != nil {
		logger.Error("error listing nics",
			genlog.String("request url", fmt.Sprint(lparURI)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return nil, status, err
	}
	logger.Info(fmt.Sprintf("request url: %v, status: %v, nic uri's: %v", lparURI, status, props.NicUris))
	return props.NicUris, status, nil
}

// AttachStorageGroupToPartition

/**
* POST /api/partitions/{partition-id}/operations/attach-storage-group
* Return: 200
*     or: 400, 404, 409
 */
func (m *LparManager) AttachStorageGroupToPartition(lparURI string, request *StorageGroupPayload) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/attach-storage-group")

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodPost))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, request, "")
	if err != nil {
		logger.Error("error on attach storage group to partition",
			genlog.String("request url", fmt.Sprint(lparURI)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return status, err
	}

	if status == http.StatusNoContent {
		logger.Info(fmt.Sprintf("Response: attach storage group to partition successfull, request url: %v, status: %v", lparURI, status))
		return status, nil
	}

	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error attaching storage group to partition",
		genlog.String("request url", fmt.Sprint(lparURI)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return status, errorResponseBody
}

// DetachStorageGroupToPartition
/**
* POST /api/partitions/{partition-id}/operations/detach-storage-group
* Return: 200
*     or: 400, 404, 409
 */
func (m *LparManager) DetachStorageGroupToPartition(lparURI string, request *StorageGroupPayload) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/detach-storage-group")

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodPost))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, request, "")

	if err != nil {
		logger.Error("error on detach storage group to partition",
			genlog.String("request url", fmt.Sprint(lparURI)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return status, err
	}

	if status == http.StatusNoContent {
		logger.Info(fmt.Sprintf("Response: detach storage group to partition successfull, request url: %v, status: %v", lparURI, status))
		return status, nil
	}

	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error detaching storage group to partition",
		genlog.String("request url", fmt.Sprint(lparURI)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return status, errorResponseBody
}

// FetchAsciiConsoleURI
/**
* POST /api/partitions/{partition-id}/operations/get-ascii-console-websocket-uri
* Return: 200 and ascii-console-websocket-uri and sessionID for the given lpar
*     or: 400, 404, 409
 */
func (m *LparManager) FetchAsciiConsoleURI(lparURI string, request *AsciiConsoleURIPayload) (*AsciiConsoleURIResponse, int, *HmcError) {
	// Start a new session for each ascii console URI
	consoleSessionID, status, err := m.client.LogonConsole()
	if err != nil {
		return nil, status, err
	}
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/get-ascii-console-websocket-uri")

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodPost))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, request, consoleSessionID)

	if err != nil {
		logger.Error("error on fetch ascii console uri",
			genlog.String("request url", fmt.Sprint(lparURI)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return nil, status, err
	}

	if status == http.StatusOK {
		responseObj := &AsciiConsoleURIResponse{}

		err := json.Unmarshal(responseBody, &responseObj)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		if responseObj.URI != "" {
			newResponseObj := &AsciiConsoleURIResponse{
				URI:       path.Join(requestUrl.Host, responseObj.URI),
				SessionID: consoleSessionID,
			}

			logger.Info(fmt.Sprintf("Response: request url: %v, status: %v, ascii console object: %v", lparURI, status, responseObj))
			return newResponseObj, status, nil
		}
		logger.Error("error on fetch ascii console uri",
			genlog.String("request url", fmt.Sprint(lparURI)),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(errors.New("empty job uri")))
		return responseObj, status, getHmcErrorFromMsg(ERR_CODE_EMPTY_JOB_URI, ERR_MSG_EMPTY_JOB_URI)
	}

	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error on fetch ascii console uri",
		genlog.String("request url", fmt.Sprint(lparURI)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return nil, status, errorResponseBody
}
