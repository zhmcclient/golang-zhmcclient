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

var _ = Describe("Adapter", func() {
	var (
		manager              *AdapterManager
		fakeClient           *fakes.ClientAPI
		cpcid                string
		adapterid            string
		url                  *url.URL
		hmcErr, unmarshalErr *HmcError
	)

	BeforeEach(func() {
		fakeClient = &fakes.ClientAPI{}
		manager = NewAdapterManager(fakeClient)

		url, _ = url.Parse("https://127.0.0.1:443")
		cpcid = "cpcid"
		adapterid = "adapterid"

		hmcErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
			Message: "error message",
		}

		unmarshalErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_UNMARSHAL_FAIL),
			Message: "invalid character 'i' looking for beginning of value",
		}
	})

	Describe("ListAdapters", func() {
		var (
			adapters      []Adapter
			adaptersArray AdaptersArray
			bytes         []byte
		)

		BeforeEach(func() {
			adapters = []Adapter{
				{
					URI:    "uri1",
					Name:   "name1",
					ID:     "id1",
					Family: ADAPTER_FAMILY_HIPERSOCKET,
					Type:   ADAPTER_TYPE_HIPERSOCKET,
					Status: ADAPTER_STATUS_ACTIVE,
				},
				{
					URI:    "uri2",
					Name:   "name2",
					ID:     "id2",
					Family: ADAPTER_FAMILY_HIPERSOCKET,
					Type:   ADAPTER_TYPE_HIPERSOCKET,
					Status: ADAPTER_STATUS_ACTIVE,
				},
			}

			adaptersArray = AdaptersArray{
				adapters,
			}
			bytes, _ = json.Marshal(adaptersArray)
		})

		Context("When ListAdapters returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, _, err := manager.ListAdapters(cpcid, nil)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets[0]).To(Equal(adapters[0]))
			})
		})

		Context("When ListAdapters returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, hmcErr)
				rets, _, err := manager.ListAdapters(cpcid, nil)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When ListAdapters returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.ListAdapters(cpcid, nil)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When ListAdapters returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytes, nil)
				rets, _, err := manager.ListAdapters(cpcid, nil)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("GetAdapterProperties", func() {
		var (
			response      *AdapterProperties
			bytesResponse []byte
		)

		BeforeEach(func() {
			response = &AdapterProperties{
				URI:              "uri",
				Name:             "adapter",
				ID:               "id",
				ObjectID:         "object-id",
				Description:      "description",
				Family:           ADAPTER_FAMILY_HIPERSOCKET,
				Type:             ADAPTER_TYPE_HIPERSOCKET,
				Status:           ADAPTER_STATUS_ACTIVE,
				DetectedCardType: ADPATER_CARD_HIPERSOCKETS,
			}
			bytesResponse, _ = json.Marshal(response)
		})
		Context("When GetAdapterProperties returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytesResponse, nil)
				rets, _, err := manager.GetAdapterProperties(adapterid)
				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets.URI).To(Equal(response.URI))
			})
		})
		Context("When GetAdapterProperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.GetAdapterProperties(adapterid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})
		Context("When GetAdapterProperties returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.GetAdapterProperties(adapterid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})
		Context("When GetAdapaterProperties returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytesResponse, nil)
				rets, _, err := manager.GetAdapterProperties(adapterid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

	})

	Describe("CreateHipersocket", func() {
		var (
			response                *HipersocketCreateResponse
			responseWithoutURI      *HipersocketCreateResponse
			payload                 *HipersocketPayload
			bytesResponse           []byte
			bytesResponseWithoutURI []byte
		)

		BeforeEach(func() {
			response = &HipersocketCreateResponse{
				URI: "uri",
			}
			responseWithoutURI = &HipersocketCreateResponse{
				URI: "",
			}
			payload = &HipersocketPayload{
				Name:            "name",
				Description:     "description",
				PortDescription: "port_description",
				MaxUnitSize:     1024,
			}

			bytesResponse, _ = json.Marshal(response)
			bytesResponseWithoutURI, _ = json.Marshal(responseWithoutURI)
		})

		Context("When CreateHipersocket and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusCreated, bytesResponse, nil)
				rets, _, err := manager.CreateHipersocket(cpcid, payload)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(Equal(""))
				Expect(rets).To(Equal(response.URI))
			})
		})

		Context("When CreateHipersocket returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.CreateHipersocket(cpcid, payload)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(Equal(""))
			})
		})

		Context("When CreateHipersocket returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusCreated, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.CreateHipersocket(cpcid, payload)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(Equal(""))
			})
		})

		Context("When CreateHipersocket returns hmcErr for WithoutURI", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusCreated, bytesResponseWithoutURI, nil)
				rets, _, err := manager.CreateHipersocket(cpcid, payload)

				Expect(err).To(BeNil())
				Expect(rets).To(Equal(""))
			})
		})
	})

	Describe("DeleteHipersocket", func() {

		Context("When DeleteHipersocket returns correctly with status 204", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				_, err := manager.DeleteHipersocket(adapterid)

				Expect(err).To(BeNil())
			})
		})

		Context("When DeleteHipersocket returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, nil, hmcErr)
				_, err := manager.DeleteHipersocket(adapterid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})
	})

	Describe("GetNetworkAdapterPort", func() {
		var (
			response      *NetworkAdapterPort
			bytesResponse []byte
		)

		BeforeEach(func() {
			response = &NetworkAdapterPort{
				URI:         "uri",
				Name:        "adapter",
				ID:          "id",
				Parent:      "parent",
				Description: "description",
				Class:       "class",
				Index:       1,
			}
			bytesResponse, _ = json.Marshal(response)
		})
		Context("When GetNetworkAdapterPortProperties returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytesResponse, nil)
				rets, _, err := manager.GetNetworkAdapterPortProperties(adapterid)
				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets.URI).To(Equal(response.URI))
			})
		})
		Context("When GetNetworkAdapterPortProperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.GetNetworkAdapterPortProperties(adapterid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})
		Context("When GetNetworkAdapterPortProperties returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.GetNetworkAdapterPortProperties(adapterid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})
		Context("When GetNetworkAdapterPortProperties returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytesResponse, nil)
				rets, _, err := manager.GetNetworkAdapterPortProperties(adapterid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

	})

	Describe("GetStorageAdapterPort", func() {
		var (
			response      *StorageAdapterPort
			bytesResponse []byte
		)

		BeforeEach(func() {
			response = &StorageAdapterPort{
				URI:                     "uri",
				Name:                    "adapter",
				ID:                      "id",
				Parent:                  "parent",
				Description:             "description",
				Class:                   "class",
				Index:                   1,
				FabricID:                "fabricID",
				ConnectionEndpointURI:   "connection-uri",
				ConnectionEndpointClass: ADAPTER_CONNECTION_STORAGE_SUBSYSTEM,
			}
			bytesResponse, _ = json.Marshal(response)
		})
		Context("When GetStorageAdapterPortProperties returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytesResponse, nil)
				rets, _, err := manager.GetStorageAdapterPortProperties(adapterid)
				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets.URI).To(Equal(response.URI))
			})
		})
		Context("When GetStorageAdapterPortProperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.GetStorageAdapterPortProperties(adapterid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})
		Context("When GetStorageAdapterPortProperties returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.GetStorageAdapterPortProperties(adapterid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})
		Context("When GetStorageAdapterPortProperties returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytesResponse, nil)
				rets, _, err := manager.GetStorageAdapterPortProperties(adapterid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

	})

})
