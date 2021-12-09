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
	"fmt"
	"net/http"
	"path"
)

// LparAPI defines an interface for issuing LPAR requests to ZHMC
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fakes/lpar.go --fake-name LparAPI . LparAPI
type LparAPI interface {
	ListLPARs(cpcURI string, query map[string]string) ([]LPAR, *HmcError)
	GetLparProperties(lparURI string) (*LparProperties, *HmcError)
	UpdateLparProperties(lparURI string, props *LparProperties) *HmcError
	StartLPAR(lparURI string) (string, *HmcError)
	StopLPAR(lparURI string) (string, *HmcError)
	AttachStorageGroupToPartition(storageGroupURI string, request *StorageGroupPayload) *HmcError
	DetachStorageGroupToPartition(storageGroupURI string, request *StorageGroupPayload) *HmcError
	MountIsoImage(lparURI string, isoFile string, insFile string) *HmcError
	UnmountIsoImage(lparURI string) *HmcError

	ListNics(lparURI string) ([]string, *HmcError)
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
func (m *LparManager) ListLPARs(cpcURI string, query map[string]string) ([]LPAR, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, cpcURI, "/partitions")
	requestUrl = BuildUrlFromQuery(requestUrl, query)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		lpars := &LPARsArray{}
		err := json.Unmarshal(responseBody, lpars)
		if err != nil {
			return nil, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return lpars.LPARS, nil
	}

	return nil, GenerateErrorFromResponse(responseBody)
}

/**
* GET /api/partitions/{partition-id}
* @lparURI is the object-uri
* Return: 200 and LparProperties
*     or: 400, 404,
 */
func (m *LparManager) GetLparProperties(lparURI string) (*LparProperties, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		lparProps := LparProperties{}
		err := json.Unmarshal(responseBody, &lparProps)
		if err != nil {
			return nil, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}

		return &lparProps, nil
	}

	return nil, GenerateErrorFromResponse(responseBody)
}

/**
* POST /api/partitions/{partition-id}
* @lparURI is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503,
 */
func (m *LparManager) UpdateLparProperties(lparURI string, props *LparProperties) *HmcError {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, props)
	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(responseBody)
}

/**
* POST /api/partitions/{partition-id}/operations/start
* @lparURI is the object-uri
* @return job-uri
* Return: 202 and job-uri
*     or: 400, 403, 404, 503,
 */
func (m *LparManager) StartLPAR(lparURI string) (string, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/start")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil)
	if err != nil {
		return "", err
	}

	if status == http.StatusAccepted {
		responseObj := StartStopLparResponse{}
		err := json.Unmarshal(responseBody, &responseObj)
		if err != nil {
			return "", getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		if responseObj.URI != "" {
			return responseObj.URI, nil
		}
		return "", getHmcErrorFromMsg(ERR_CODE_EMPTY_JOB_URI, ERR_MSG_EMPTY_JOB_URI)
	}

	return "", GenerateErrorFromResponse(responseBody)
}

/**
* POST /api/partitions/{partition-id}/operations/stop
* @lparURI is the object-uri
* @return job-uri
* Return: 202 and job-uri
*     or: 400, 403, 404, 503,
 */
func (m *LparManager) StopLPAR(lparURI string) (string, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/stop")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil)
	if err != nil {
		return "", err
	}

	if status == http.StatusAccepted {
		responseObj := StartStopLparResponse{}
		err := json.Unmarshal(responseBody, &responseObj)
		if err != nil {
			return "", getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		if responseObj.URI != "" {
			return responseObj.URI, nil
		}
		return "", getHmcErrorFromMsg(ERR_CODE_EMPTY_JOB_URI, ERR_MSG_EMPTY_JOB_URI)
	}

	return "", GenerateErrorFromResponse(responseBody)
}

/**
* POST /api/partitions/{partition-id}/operations/mount-iso-image
* @lparURI is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *LparManager) MountIsoImage(lparURI string, isoFile string, insFile string) *HmcError {
	pureIsoName := path.Base(isoFile)
	pureInsName := path.Base(insFile)
	query := map[string]string{
		"image-name":    pureIsoName,
		"ins-file-name": "/boot/" + pureInsName,
	}
	imageData, byteErr := RetrieveBytes(isoFile)
	if byteErr != nil {
		fmt.Println("Error: ", byteErr.Error())
	}
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/mount-iso-image")
	requestUrl = BuildUrlFromQuery(requestUrl, query)

	status, responseBody, err := m.client.UploadRequest(http.MethodPost, requestUrl, imageData)

	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}
	return GenerateErrorFromResponse(responseBody)
}

/**
* POST /api/partitions/{partition-id}/operations/unmount-iso-image
* @lparURI is the object-uri
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *LparManager) UnmountIsoImage(lparURI string) *HmcError {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/unmount-iso-image")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, nil)
	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(responseBody)
}

/**
* get_property('nic-uris') from LPAR
 */
func (m *LparManager) ListNics(lparURI string) ([]string, *HmcError) {
	props, err := m.GetLparProperties(lparURI)
	if err != nil {
		return nil, err
	}

	return props.NicUris, nil
}

// AttachStorageGroupToPartition

/**
* POST /api/partitions/{partition-id}/operations/attach-storage-group
* Return: 200
*     or: 400, 404, 409
 */
func (m *LparManager) AttachStorageGroupToPartition(lparURI string, request *StorageGroupPayload) *HmcError {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "“/operations/attach-storage-group”")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, request)

	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(responseBody)
}

// DetachStorageGroupToPartition
/**
* POST /api/partitions/{partition-id}/operations/detach-storage-group
* Return: 200
*     or: 400, 404, 409
 */
func (m *LparManager) DetachStorageGroupToPartition(lparURI string, request *StorageGroupPayload) *HmcError {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, lparURI, "/operations/detach-storage-group")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, request)

	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(responseBody)
}
