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

	"github.ibm.com/genctl/shared-logger/genlog"
)

// StorageGroupAPI defines an interface for issuing NIC requests to ZHMC
//go:generate counterfeiter -o fakes/sgroup.go --fake-name StorageGroupAPI . StorageGroupAPI

type StorageGroupAPI interface {
	ListStorageGroups(storageGroupURI string, cpcUri string) ([]StorageGroup, int, *HmcError)
	GetStorageGroupProperties(storageGroupURI string) (*StorageGroupProperties, int, *HmcError)
	ListStorageVolumes(storageGroupURI string) ([]StorageVolume, int, *HmcError)
	GetStorageVolumeProperties(storageVolumeURI string) (*StorageVolume, int, *HmcError)
	UpdateStorageGroupProperties(storageGroupURI string, updateRequest *StorageGroupProperties) (int, *HmcError)
	FulfillStorageGroup(storageGroupURI string, updateRequest *StorageGroupProperties) (int, *HmcError)
}

type StorageGroupManager struct {
	client ClientAPI
}

func NewStorageGroupManager(client ClientAPI) *StorageGroupManager {
	return &StorageGroupManager{
		client: client,
	}
}

/**
 * GET /api/storage-groups
 * @cpcURI the URI of the CPC
 * @return storage group array
 * Return: 200 and Storage Group array
 *     or: 400, 404, 409
 */
func (m *StorageGroupManager) ListStorageGroups(storageGroupURI string, cpcUri string) ([]StorageGroup, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageGroupURI)
	query := map[string]string{
		"cpc-uri": cpcUri,
	}
	requestUrl = BuildUrlFromQuery(requestUrl, query)

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodGet))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("error on list storage groups",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("method", http.MethodGet),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return nil, status, err
	}

	if status == http.StatusOK {
		storageGroups := &StorageGroupArray{}
		err := json.Unmarshal(responseBody, storageGroups)
		if err != nil {
			logger.Error("error on unmarshalling adapters",
				genlog.String("request url", fmt.Sprint(requestUrl)),
				genlog.String("method", http.MethodGet),
				genlog.Error(fmt.Errorf("%v", getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err))))
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Response: request url: %v, method: %v, status: %v, storage groups: %v", requestUrl, http.MethodGet, status, storageGroups.STORAGEGROUPS))
		return storageGroups.STORAGEGROUPS, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error on list storage groups",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("method", http.MethodGet),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return nil, status, errorResponseBody
}

/**
 * GET /api/storage-groups/{storage-group-id}
 * @cpcURI the ID of the virtual switch
 * @return adapter array
 * Return: 200 and Storage Group object
 *     or: 400, 404, 409
 */
func (m *StorageGroupManager) GetStorageGroupProperties(storageGroupURI string) (*StorageGroupProperties, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageGroupURI)

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodGet))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("error on get storage group properties",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("method", http.MethodGet),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return nil, status, err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroupProperties{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			logger.Error("error on unmarshalling adapters",
				genlog.String("request url", fmt.Sprint(requestUrl)),
				genlog.String("method", http.MethodGet),
				genlog.Error(fmt.Errorf("%v", getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err))))
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Response: request url: %v, method: %v, status: %v, storage group properties: %v", requestUrl, http.MethodGet, status, storageGroup))
		return storageGroup, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error on get storage group properties",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return nil, status, errorResponseBody
}

/**
 * GET /api/storage-groups/{storage-group-id}/storage-volumes
 * @storage-group-id the Object id of the storage group
 * @return storage volume array
 * Return: 200 and Storage Group array
 *     or: 400, 404, 409
 */
func (m *StorageGroupManager) ListStorageVolumes(storageGroupURI string) ([]StorageVolume, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageGroupURI)

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodGet))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("error on list storage volumes",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("method", http.MethodGet),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return nil, status, err
	}

	if status == http.StatusOK {
		storageVolumes := &StorageVolumeArray{}
		err := json.Unmarshal(responseBody, storageVolumes)
		if err != nil {
			logger.Error("error on unmarshalling adapters",
				genlog.String("request url", fmt.Sprint(requestUrl)),
				genlog.String("method", http.MethodGet),
				genlog.Error(fmt.Errorf("%v", getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err))))
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Response: request url: %v, method: %v, status: %v, storage volumes: %v", requestUrl, http.MethodGet, status, storageVolumes.STORAGEVOLUMES))
		return storageVolumes.STORAGEVOLUMES, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error on list storage volumes",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("method", http.MethodGet),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return nil, status, errorResponseBody
}

/**
 * GET /api/storage-groups/{storage-group-id}/storage-volumes/{storage-volume-id}
 * @return volume object
 * Return: 200 and Storage Volume object
 *     or: 400, 404, 409
 */
func (m *StorageGroupManager) GetStorageVolumeProperties(storageVolumeURI string) (*StorageVolume, int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageVolumeURI)

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodGet))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("error on get storage volume properties",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("method", http.MethodGet),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return nil, status, err
	}

	if status == http.StatusOK {
		storageVolume := &StorageVolume{}
		err := json.Unmarshal(responseBody, storageVolume)
		if err != nil {
			logger.Error("error on unmarshalling adapters",
				genlog.String("request url", fmt.Sprint(requestUrl)),
				genlog.String("method", http.MethodGet),
				genlog.Error(fmt.Errorf("%v", getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err))))
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Respone: request url: %v, method: %v, status: %v, storage volume properties: %v", http.MethodGet, requestUrl, status, storageVolume))
		return storageVolume, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("error on get storage volume properties",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("method", http.MethodGet),
		genlog.String("status: ", fmt.Sprint(status)),
		genlog.Error(fmt.Errorf("%v", errorResponseBody)))
	return nil, status, errorResponseBody
}

/**
 * POST /api/storage-groups/{storage-group-id}/operations/modify
 * Return: 200
 *     or: 400, 404, 409
 */
func (m *StorageGroupManager) UpdateStorageGroupProperties(storageGroupURI string, updateRequest *StorageGroupProperties) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageGroupURI)

	logger.Info(fmt.Sprintf("Request URL: %v,  Method: %v", requestUrl, http.MethodPost))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, updateRequest, "")
	if err != nil {
		logger.Error("error on update storage group properties",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("method", http.MethodPost),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return status, err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroup{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			logger.Error("error on unmarshalling adapters",
				genlog.String("request url", fmt.Sprint(requestUrl)),
				genlog.String("method", http.MethodPost),
				genlog.Error(fmt.Errorf("%v", getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err))))
			return status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Response: update storage group properties completed, request url: %v, method: %v, status: %v", requestUrl, http.MethodPost, status))
		return status, nil
	}
	logger.Error("error on update storage group properties",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("method", http.MethodPost),
		genlog.String("status: ", fmt.Sprint(status)))
	return status, nil
}

/**
* POST /api/storage-groups/{storage-group-id}/operations/accept-mismatched-
  storage-volumes
* Return: 200
*     or: 400, 404, 409
*/
func (m *StorageGroupManager) FulfillStorageGroup(storageGroupURI string, request *StorageGroupProperties) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageGroupURI)

	logger.Info(fmt.Sprintf("Request URL: %v, Method: %v", requestUrl, http.MethodPost))
	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, request, "")
	if err != nil {
		logger.Error("error on fulfill storage group",
			genlog.String("request url", fmt.Sprint(requestUrl)),
			genlog.String("method", http.MethodPost),
			genlog.String("status", fmt.Sprint(status)),
			genlog.Error(fmt.Errorf("%v", err)))
		return status, err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroup{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			logger.Error("error on unmarshalling adapters",
				genlog.String("request url", fmt.Sprint(requestUrl)),
				genlog.String("method", http.MethodPost),
				genlog.Error(fmt.Errorf("%v", getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err))))
			return status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Response: fulfill storage group completed, request url: %v, method: %v, status: %v", requestUrl, http.MethodPost, status))
		return status, nil
	}
	logger.Error("error on fulfill storage group",
		genlog.String("request url", fmt.Sprint(requestUrl)),
		genlog.String("method", http.MethodPost),
		genlog.String("status: ", fmt.Sprint(status)))
	return status, nil
}
