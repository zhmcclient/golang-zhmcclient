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
	"net/http"
	"path"
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
	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		return nil, status, err
	}

	if status == http.StatusOK {
		storageGroups := &StorageGroupArray{}
		err := json.Unmarshal(responseBody, storageGroups)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return storageGroups.STORAGEGROUPS, status, nil
	}

	return nil, status, GenerateErrorFromResponse(responseBody)
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

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		return nil, status, err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroupProperties{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return storageGroup, status, nil
	}

	return nil, status, GenerateErrorFromResponse(responseBody)
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
	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		return nil, status, err
	}

	if status == http.StatusOK {
		storageVolumes := &StorageVolumeArray{}
		err := json.Unmarshal(responseBody, storageVolumes)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return storageVolumes.STORAGEVOLUMES, status, nil
	}

	return nil, status, GenerateErrorFromResponse(responseBody)
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

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		return nil, status, err
	}

	if status == http.StatusOK {
		storageVolume := &StorageVolume{}
		err := json.Unmarshal(responseBody, storageVolume)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return storageVolume, status, nil
	}

	return nil, status, GenerateErrorFromResponse(responseBody)
}

/**
 * POST /api/storage-groups/{storage-group-id}/operations/modify
 * Return: 200
 *     or: 400, 404, 409
 */
func (m *StorageGroupManager) UpdateStorageGroupProperties(storageGroupURI string, updateRequest *StorageGroupProperties) (int, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageGroupURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, updateRequest, "")
	if err != nil {
		return status, err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroup{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			return status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return status, nil
	}

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

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, request, "")

	if err != nil {
		return status, err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroup{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			return status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return status, nil
	}

	return status, nil
}
