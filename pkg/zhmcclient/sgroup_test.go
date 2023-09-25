// Copyright 2021-2023 IBM Corp. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zhmcclient_test

import (
	"encoding/json"

	"net/http"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
	"github.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient/fakes"
)

var _ = Describe("Storage Group", func() {
	var (
		manager              *StorageGroupManager
		fakeClient           *fakes.ClientAPI
		cpcid                string
		sgroupid             string
		url                  *url.URL
		hmcErr, unmarshalErr *HmcError
	)

	BeforeEach(func() {
		fakeClient = &fakes.ClientAPI{}
		manager = NewStorageGroupManager(fakeClient)

		url, _ = url.Parse("https://127.0.0.1:443")
		cpcid = "cpcid"
		sgroupid = "sgroupid"

		hmcErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
			Message: "error message",
		}

		unmarshalErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_UNMARSHAL_FAIL),
			Message: "invalid character 'i' looking for beginning of value",
		}
	})

	Describe("ListStorageGroups", func() {
		var (
			storageGroups      []StorageGroup
			storageGroupsArray StorageGroupArray
			bytes              []byte
		)

		BeforeEach(func() {
			storageGroups = []StorageGroup{
				{
					ObjectID:         "object-id",
					Name:             "name1",
					FulfillmentState: "status",
					Type:             "type",
				},
				{
					ObjectID:         "object-id",
					Name:             "name1",
					FulfillmentState: "status",
					Type:             "type",
				},
			}
			storageGroupsArray = StorageGroupArray{
				storageGroups,
			}
			bytes, _ = json.Marshal(storageGroupsArray)
		})

		Context("When ListStorageGroups returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, _, err := manager.ListStorageGroups(sgroupid, cpcid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets[0]).To(Equal(storageGroups[0]))
			})
		})

		Context("When ListStorageGroups returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, hmcErr)
				rets, _, err := manager.ListStorageGroups(sgroupid, cpcid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When ListStorageGroups returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.ListStorageGroups(sgroupid, cpcid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})
		Context("When ListStorageGroups returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytes, nil)
				rets, _, err := manager.ListStorageGroups(sgroupid, cpcid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("UpdateStorageGroupProperties", func() {
		var (
			response           *StorageGroupProperties
			storageVolumeURIs  []string
			virtualStorageURIs []string
			bytesResponse      []byte
		)

		BeforeEach(func() {
			response = &StorageGroupProperties{
				Class:                      "class",
				CpcURI:                     "cpcuri",
				Connectivity:               4,
				Name:                       "lpar",
				Description:                "description",
				FulfillmentState:           PARTITION_STATUS_STARTING,
				Type:                       PARTITION_TYPE_LINUX,
				StorageVolumesURIs:         storageVolumeURIs,
				ActiveMaxPartitions:        1,
				MaxPartitions:              2,
				VirtualStorageResourceURIs: virtualStorageURIs,
			}
			bytesResponse, _ = json.Marshal(response)
		})

		Context("When UpdateStorageGroupProperties returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, bytesResponse, nil)
				_, err := manager.UpdateStorageGroupProperties(sgroupid, response)

				Expect(err).To(BeNil())
			})
		})

		Context("When UpdateStorageGroupProperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, err := manager.UpdateStorageGroupProperties(sgroupid, response)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When UpdateStorageGroupProperties returns error due to unmarshalerr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				_, err := manager.UpdateStorageGroupProperties(sgroupid, response)

				Expect(*err).To(Equal(*unmarshalErr))
			})
		})
	})

	Describe("FulfillStorageGroup", func() {
		var (
			response           *StorageGroupProperties
			storageVolumeURIs  []string
			virtualStorageURIs []string
			bytesResponse      []byte
		)

		BeforeEach(func() {
			response = &StorageGroupProperties{
				Class:                      "class",
				CpcURI:                     "cpcuri",
				Connectivity:               4,
				Name:                       "lpar",
				Description:                "description",
				FulfillmentState:           PARTITION_STATUS_STARTING,
				Type:                       PARTITION_TYPE_LINUX,
				StorageVolumesURIs:         storageVolumeURIs,
				ActiveMaxPartitions:        1,
				MaxPartitions:              2,
				VirtualStorageResourceURIs: virtualStorageURIs,
			}
			bytesResponse, _ = json.Marshal(response)
		})

		Context("When FulfillStorageGroup returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, bytesResponse, nil)
				_, err := manager.FulfillStorageGroup(sgroupid, response)

				Expect(err).To(BeNil())
			})
		})

		Context("When FulfillStorageGroup returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, err := manager.FulfillStorageGroup(sgroupid, response)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When FulfillStorageGroup returns error due to unmarshalerr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				_, err := manager.FulfillStorageGroup(sgroupid, response)

				Expect(*err).To(Equal(*unmarshalErr))
			})
		})

		Context("When FulfillStorageGroup returns incorrect status", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytesResponse, hmcErr)
				_, err := manager.FulfillStorageGroup(sgroupid, response)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("GetStorageGroupProperties", func() {
		var (
			response           *StorageGroupProperties
			storageVolumeURIs  []string
			virtualStorageURIs []string
			bytesResponse      []byte
		)

		BeforeEach(func() {
			storageVolumeURIs = make([]string, 0)
			storageVolumeURIs = append(storageVolumeURIs, "Volume1-URI")
			storageVolumeURIs = append(storageVolumeURIs, "Volume2-URI")

			response = &StorageGroupProperties{
				Class:                      "class",
				CpcURI:                     "cpcuri",
				Connectivity:               4,
				Name:                       "lpar",
				Description:                "description",
				FulfillmentState:           PARTITION_STATUS_STARTING,
				Type:                       PARTITION_TYPE_LINUX,
				StorageVolumesURIs:         storageVolumeURIs,
				ActiveMaxPartitions:        1,
				MaxPartitions:              2,
				VirtualStorageResourceURIs: virtualStorageURIs,
			}

			bytesResponse, _ = json.Marshal(response)
		})

		Context("When GetStorageGroupProperties and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytesResponse, nil)
				rets, _, err := manager.GetStorageGroupProperties(sgroupid)

				Expect(rets.Class).To(Equal("class"))
				Expect(rets.MaxPartitions).To(Equal(2))
				Expect(rets.Connectivity).To(Equal(4))
				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
			})
		})

		Context("When GetStorageGroupproperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, _, err := manager.GetStorageGroupProperties(sgroupid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When GetStorageGroupProperties returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.GetStorageGroupProperties(sgroupid)

				Expect(err).ToNot(BeNil())
				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When GetStorageGroupProperties returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytesResponse, nil)
				rets, _, err := manager.GetStorageGroupProperties(sgroupid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("GetStorageVolumeProperties", func() {
		var (
			response      *StorageVolume
			volumePaths   []VolumePath
			bytesResponse []byte
		)

		BeforeEach(func() {

			volumePaths = []VolumePath{
				{
					PartitionURI:      "PartitionURI",
					DeviceNumber:      "DeviceNumber",
					TargetWWPN:        "TargetWWPN",
					LogicalUnitNumber: "LUN",
				},
			}

			response = &StorageVolume{
				Class:            "class",
				URI:              "StorageVolume",
				Usage:            "usage",
				Name:             "lpar",
				Description:      "description",
				FulfillmentState: STORAGE_GROUP_COMPLETE,
				Paths:            volumePaths,
			}

			bytesResponse, _ = json.Marshal(response)
		})

		Context("When GetStorageVolumeProperties returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytesResponse, nil)
				rets, _, err := manager.GetStorageVolumeProperties(sgroupid)

				Expect(rets.Class).To(Equal("class"))
				Expect(string(rets.FulfillmentState)).To(Equal("complete"))
				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
			})
		})

		Context("When GetStorageVolumeProperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, _, err := manager.GetStorageVolumeProperties(sgroupid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When GetStorageVolumeProperties returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.GetStorageVolumeProperties(sgroupid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When GetStorageVolumeProperties returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytesResponse, nil)
				rets, _, err := manager.GetStorageVolumeProperties(sgroupid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("ListStorageVolumes", func() {
		var (
			storagevolumes     []StorageVolume
			storagevolumearray StorageVolumeArray
			bytesResponse      []byte
		)

		BeforeEach(func() {
			storagevolumes = []StorageVolume{
				{
					Class:            "class",
					Parent:           "parent",
					URI:              "uri",
					Name:             "name",
					Description:      "description",
					FulfillmentState: STORAGE_GROUP_COMPLETE,
				},
			}
			storagevolumearray = StorageVolumeArray{
				storagevolumes,
			}
			bytesResponse, _ = json.Marshal(storagevolumearray)
		})

		Context("When ListStorageVolumes returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytesResponse, nil)
				_, _, err := manager.ListStorageVolumes(sgroupid)

				Expect(err).To(BeNil())
			})
		})

		Context("When ListStorageVolumes returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, _, err := manager.ListStorageVolumes(sgroupid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When ListStorageVolumes returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytesResponse, nil)
				rets, _, err := manager.ListStorageVolumes(sgroupid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When ListStorageVolumes returns error due to unmarshalerr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.ListStorageVolumes(sgroupid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})
	})

	// New Entries

	Describe("CreateStorageGroups", func() {
		var (
			response                *StorageGroupCreateResponse
			responseWithoutURI      *StorageGroupCreateResponse
			payload                 *CreateStorageGroupProperties
			response1, response2    *StorageVolume
			volumePaths             []VolumePath
			bytesResponse           []byte
			bytesResponse1          []byte
			bytesResponse2          []byte
			bytesResponseWithoutURI []byte
		)

		BeforeEach(func() {

			volumePaths = []VolumePath{
				{
					PartitionURI:      "partition-uri",
					DeviceNumber:      "09",
					TargetWWPN:        "target-wwpn",
					LogicalUnitNumber: "logicalunit-number",
				},
			}
			response = &StorageGroupCreateResponse{
				URI:       []string{"Storagevolume1"},
				ObjectURI: "object_uri",
				SvPaths: []StorageGroupVolumePath{
					{

						URI:   "Storagevolume1",
						Paths: volumePaths,
					},
				},
			}

			response1 = &StorageVolume{

				URI:   "Storagevolume1",
				Paths: volumePaths,
			}
			response2 = &StorageVolume{

				URI:   "Storagevolume2",
				Paths: volumePaths,
			}
			responseWithoutURI = &StorageGroupCreateResponse{
				ObjectURI: "",
				URI:       nil,
				SvPaths:   nil,
			}
			payload = &CreateStorageGroupProperties{
				CpcURI: "cpc_uri",
				Name:   "name",
				Type:   "type",
			}

			bytesResponse, _ = json.Marshal(response)
			bytesResponse1, _ = json.Marshal(response1)
			bytesResponse2, _ = json.Marshal(response2)
			bytesResponseWithoutURI, _ = json.Marshal(responseWithoutURI)
		})

		Context("When CreateStorageGroup and returns correctly", func() {
			It("check the results succeed", func() {

				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturnsOnCall(0, http.StatusCreated, bytesResponse, nil)
				fakeClient.ExecuteRequestReturnsOnCall(1, http.StatusOK, bytesResponse1, nil)

				rets, status, err := manager.CreateStorageGroups(sgroupid, payload)

				Expect(status).To(Equal(201))

				Expect(err).To(BeNil())
				Expect(rets).ToNot(Equal(""))
				Expect(rets.URI).To(Equal(response.URI))
				Expect(rets.ObjectURI).To(Equal(response.ObjectURI))
				Expect(rets.SvPaths).To(Equal(response.SvPaths))

			})
		})

		Context("When CreateStorageGroup and ExecuteRequest error", func() {
			It("check the error happened", func() {

				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.CreateStorageGroups(sgroupid, payload)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When CreateStorageGroup and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusCreated, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.CreateStorageGroups(sgroupid, payload)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When CreateStorageGroup and no URI responded", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, bytesResponseWithoutURI, hmcErr)
				rets, _, err := manager.CreateStorageGroups(sgroupid, payload)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When CreateStorageGroup creates Multiple StorageVolumes", func() {
			BeforeEach(func() {
				response.URI = []string{"storageVolume1", "storageVolume2"}
				response.SvPaths = []StorageGroupVolumePath{
					{

						URI:   "Storagevolume1",
						Paths: volumePaths,
					},
					{

						URI:   "Storagevolume2",
						Paths: volumePaths,
					},
				}
				bytesResponse, _ = json.Marshal(response)

			})
			It("check the result success", func() {

				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturnsOnCall(0, http.StatusCreated, bytesResponse, nil)
				fakeClient.ExecuteRequestReturnsOnCall(1, http.StatusOK, bytesResponse1, nil)
				fakeClient.ExecuteRequestReturnsOnCall(2, http.StatusOK, bytesResponse2, nil)
				rets, status, err := manager.CreateStorageGroups(sgroupid, payload)
				Expect(status).To(Equal(201))
				Expect(err).To(BeNil())
				Expect(rets).ToNot(Equal(""))
				Expect(len(rets.SvPaths)).To(Equal(2))
				Expect(rets.URI).To(Equal(response.URI))
				Expect(rets.SvPaths).To(Equal(response.SvPaths))
			})
		})

	})

	Describe("GetStorageGroupPartitionss", func() {
		var (
			storageGroups      []LparProperties
			storageGroupsArray *StorageGroupPartitions
			bytes              []byte
		)

		BeforeEach(func() {
			storageGroups = []LparProperties{
				{
					URI:    "object-uri",
					Name:   "name1",
					Status: "status",
				},
				{
					URI:    "object-uri",
					Name:   "name2",
					Status: "status",
				},
			}
			storageGroupsArray = &StorageGroupPartitions{
				storageGroups,
			}
			bytes, _ = json.Marshal(storageGroupsArray)
		})

		Context("When get storage groups partitions request and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, _, err := manager.GetStorageGroupPartitions(sgroupid, nil)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets).To(Equal(storageGroupsArray))
			})
		})

		Context("When get storage groups partitions request and returns error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, hmcErr)
				rets, _, err := manager.GetStorageGroupPartitions(sgroupid, nil)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("DeleteStorageGroup", func() {

		BeforeEach(func() {
			hmcErr = &HmcError{
				Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
				Message: "error message",
			}
		})

		Context("When delete lpar and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				status, err := manager.DeleteStorageGroup(sgroupid)
				Expect(err).To(BeNil())
				Expect(status).ToNot(BeNil())
				Expect(status).To(Equal(204))
			})
		})

		Context("When delete lpar and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, nil, hmcErr)
				_, err := manager.DeleteStorageGroup(sgroupid)
				Expect(err).ToNot(BeNil())
			})
		})

		Context("When delete lpar and returns incorrect status", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, nil)
				_, err := manager.DeleteStorageGroup(sgroupid)
				Expect(err).ToNot(BeNil())
			})
		})
	})

})
