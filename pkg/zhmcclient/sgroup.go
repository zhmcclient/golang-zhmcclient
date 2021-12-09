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

// StorageGroupAPI defines an interface for issuing NIC requests to ZHMC
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fakes/nic.go --fake-name StorageGroupAPI . StorageGroupAPI

type StorageGroupAPI interface {
	ListStorageGroups(storageGroupURI string, cpc string) ([]StorageGroup, error)
	GetStorageGroupProperties(storageGroupURI string) (*StorageGroupProperties, error)
	UpdateStorageGroupProperties(storageGroupURI string, updateRequest *StorageGroupProperties) error
	FulfillStorageGroup(storageGroupURI string, updateRequest *StorageGroupProperties) error
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
func (m *StorageGroupManager) ListStorageGroups(storageGroupURI string, cpc string) ([]StorageGroup, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageGroupURI)
	query := map[string]string{
		"cpc-uri": cpc,
	}
	requestUrl = BuildUrlFromQuery(requestUrl, query)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		storageGroups := &StorageGroupArray{}
		err := json.Unmarshal(responseBody, storageGroups)
		if err != nil {
			return nil, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return storageGroups.STORAGEGROUPS, nil
	}

	return nil, GenerateErrorFromResponse(responseBody)
}

/**
 * GET /api/storage-groups/{storage-group-id}
 * @cpcURI the ID of the virtual switch
 * @return adapter array
 * Return: 200 and Storage Group object
 *     or: 400, 404, 409
 */
func (m *StorageGroupManager) GetStorageGroupProperties(storageGroupURI string) (*StorageGroupProperties, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageGroupURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroupProperties{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			return nil, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return storageGroup, nil
	}

	return nil, GenerateErrorFromResponse(responseBody)
}

/**
 * POST /api/storage-groups/{storage-group-id}/operations/modify
 * @cpcURI the ID of the virtual switch
 * @return adapter array
 * Return: 200
 *     or: 400, 404, 409
 */
func (m *StorageGroupManager) UpdateStorageGroupProperties(storageGroupURI string, updateRequest *StorageGroupProperties) *HmcError {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageGroupURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, updateRequest)
	fmt.Println("Status:", status)
	fmt.Println("Update Response:", string(responseBody))
	if err != nil {
		return err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroup{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			return getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return nil
	}

	return nil
}

/**
* POST /api/storage-groups/{storage-group-id}/operations/accept-mismatched-
  storage-volumes
* @cpcURI the ID of the virtual switch
* @return adapter array
* Return: 200 and Adapters array
*     or: 400, 404, 409
*/
func (m *StorageGroupManager) FulfillStorageGroup(storageGroupURI string, request *StorageGroupProperties) *HmcError {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, storageGroupURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, request)
	fmt.Println(requestUrl)
	if err != nil {
		return err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroup{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			return getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return nil
	}

	return nil
}
