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
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("Error on list storage groups",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return nil, status, err
	}

	if status == http.StatusOK {
		storageGroups := &StorageGroupArray{}
		err := json.Unmarshal(responseBody, storageGroups)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Status: %v, Storage groups: %v", status, storageGroups.STORAGEGROUPS))
		return storageGroups.STORAGEGROUPS, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("Error on list storage groups",
		genlog.String("Status: ", fmt.Sprint(status)),
		genlog.Error(errors.New(errorResponseBody.Message)))
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
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("Error on get storage group properties",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return nil, status, err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroupProperties{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Status: %v, Storage group properties: %v", status, storageGroup))
		return storageGroup, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("Error on get storage group properties",
		genlog.String("Status: ", fmt.Sprint(status)),
		genlog.Error(errors.New(errorResponseBody.Message)))
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
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("Error on list storage volumes",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return nil, status, err
	}

	if status == http.StatusOK {
		storageVolumes := &StorageVolumeArray{}
		err := json.Unmarshal(responseBody, storageVolumes)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Status: %v, StorageVolumes: %v", status, storageVolumes.STORAGEVOLUMES))
		return storageVolumes.STORAGEVOLUMES, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("Error on list storage volumes",
		genlog.String("Status: ", fmt.Sprint(status)),
		genlog.Error(errors.New(errorResponseBody.Message)))
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
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil, "")
	if err != nil {
		logger.Error("Error on get storage volume properties",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return nil, status, err
	}

	if status == http.StatusOK {
		storageVolume := &StorageVolume{}
		err := json.Unmarshal(responseBody, storageVolume)
		if err != nil {
			return nil, status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Status: %v, Storage volume properties: %v", status, storageVolume))
		return storageVolume, status, nil
	}
	errorResponseBody := GenerateErrorFromResponse(responseBody)
	logger.Error("Error on get storage volume properties",
		genlog.String("Status: ", fmt.Sprint(status)),
		genlog.Error(errors.New(errorResponseBody.Message)))
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
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, updateRequest, "")
	if err != nil {
		logger.Error("Error on update storage group properties",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return status, err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroup{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			return status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Update storage group properties completed, Status: %v", status))
		return status, nil
	}
	logger.Error("Error on update storage group properties",
		genlog.String("Status: ", fmt.Sprint(status)))
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
	logger.Info(fmt.Sprintf("Request URL: %v", requestUrl))

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, request, "")
	if err != nil {
		logger.Error("Error on fulfill storage group",
			genlog.String("Status", fmt.Sprint(status)),
			genlog.Error(errors.New(err.Message)))
		return status, err
	}

	if status == http.StatusOK {
		storageGroup := &StorageGroup{}
		err := json.Unmarshal(responseBody, storageGroup)
		if err != nil {
			return status, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		logger.Info(fmt.Sprintf("Fulfill storage group completed, Status: %v", status))
		return status, nil
	}
	logger.Error("Error on fulfill storage group",
		genlog.String("Status: ", fmt.Sprint(status)))
	return status, nil
}
