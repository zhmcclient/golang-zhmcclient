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

package zhmcclient_test

import (
	"encoding/json"
	"net/http"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient/fakes"
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

		Context("When list storage groups request and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, _, err := manager.ListStorageGroups(sgroupid, cpcid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets[0]).To(Equal(storageGroups[0]))
			})
		})

		Context("When list storage groups request and returns error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, hmcErr)
				rets, _, err := manager.ListStorageGroups(sgroupid, cpcid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
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

		Context("When GetLparProperties and returns correctly", func() {
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

		Context("When get stoarge properties and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, _, err := manager.GetStorageGroupProperties(sgroupid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When GetLparProperties and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.GetStorageGroupProperties(sgroupid)

				Expect(err).ToNot(BeNil())
				Expect(*err).To(Equal(*unmarshalErr))
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
				URI:              "URI",
				Usage:            "usage",
				Name:             "lpar",
				Description:      "description",
				FulfillmentState: STORAGE_GROUP_COMPLETE,
				Paths:            volumePaths,
			}

			bytesResponse, _ = json.Marshal(response)
		})

		Context("When GetLparProperties and returns correctly", func() {
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

		Context("When get stoarge properties and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, _, err := manager.GetStorageGroupProperties(sgroupid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When GetLparProperties and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.GetStorageGroupProperties(sgroupid)

				Expect(*err).To(Equal(*unmarshalErr))
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

		Context("When list stoarge volumes request and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytesResponse, nil)
				_, _, err := manager.ListStorageVolumes(sgroupid)

				Expect(err).To(BeNil())
			})
		})

		Context("When list stoarge volumes and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, _, err := manager.ListStorageVolumes(sgroupid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})
	})
})
