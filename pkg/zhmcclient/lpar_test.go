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
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
	"github.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient/fakes"
)

var _ = Describe("LPAR", func() {
	var (
		manager              *LparManager
		fakeClient           *fakes.ClientAPI
		cpcid                string
		lparid               string
		url                  *url.URL
		hmcErr, unmarshalErr *HmcError
	)

	BeforeEach(func() {
		fakeClient = &fakes.ClientAPI{}
		manager = NewLparManager(fakeClient)

		url, _ = url.Parse("https://127.0.0.1:443")
		cpcid = "cpcid"
		lparid = "lparid"

		hmcErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
			Message: "error message",
		}

		unmarshalErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_UNMARSHAL_FAIL),
			Message: "invalid character 'i' looking for beginning of value",
		}
	})

	Describe("ListLPARs", func() {
		var (
			lpars      []LPAR
			lparsArray LPARsArray
			bytes      []byte
		)

		BeforeEach(func() {
			lpars = []LPAR{
				{
					URI:    "object-uri1",
					Name:   "name1",
					Status: "status",
					Type:   "type",
				},
				{
					URI:    "object-uri2",
					Name:   "name2",
					Status: "status",
					Type:   "type",
				},
			}
			lparsArray = LPARsArray{
				lpars,
			}
			bytes, _ = json.Marshal(lparsArray)
		})

		Context("When ListLPARs returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, _, err := manager.ListLPARs(cpcid, nil)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets[0]).To(Equal(lpars[0]))
			})
		})

		Context("When ListLPARs returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, hmcErr)
				rets, _, err := manager.ListLPARs(cpcid, nil)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When ListLPARs returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.ListLPARs(cpcid, nil)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When ListLPARs returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytes, nil)
				rets, _, err := manager.ListLPARs(cpcid, nil)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("GetEnergyDetailsforLPAR", func() {
		var (
			bytes []byte
		)

		BeforeEach(func() {
			jsonString := `{
				"wattage": [
					{"data": 53, "timestamp": 1680394193292},
					{"data": 52, "timestamp": 1680408593302}
				]
			}`
			bytes = []byte(jsonString)
		})

		Context("When GetEnergyDetailsforLPAR returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				props := &EnergyRequestPayload{
					Range:      "last-day",
					Resolution: "fifteen-minutes",
				}
				rets, _, err := manager.GetEnergyDetailsforLPAR(lparid, props)

				Expect(err).To(BeNil())
				Expect(rets).To(Equal(uint64(52)))
			})
		})

		Context("When GetEnergyDetailsforLPAR returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, hmcErr)
				rets, _, err := manager.GetEnergyDetailsforLPAR(lparid, nil)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(Equal(uint64(0)))

			})
		})
	})

	Describe("GetLparProperties", func() {
		var (
			response      *LparProperties
			bytesResponse []byte
		)

		BeforeEach(func() {
			response = &LparProperties{
				URI:            "uri",
				CpcURI:         "cpcuri",
				Class:          "partition",
				Name:           "lpar",
				Description:    "description",
				Status:         PARTITION_STATUS_STARTING,
				Type:           PARTITION_TYPE_LINUX,
				ShortName:      "short_name",
				ID:             "id",
				AutoGenerateID: true,
			}

			bytesResponse, _ = json.Marshal(response)
		})

		Context("When GetLparProperties and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytesResponse, nil)
				rets, _, err := manager.GetLparProperties(lparid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets.URI).To(Equal(response.URI))
			})
		})

		Context("When GetLparProperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.GetLparProperties(lparid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When GetLparProperties returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.GetLparProperties(lparid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When GetLparProperties returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytesResponse, nil)
				rets, _, err := manager.GetLparProperties(lparid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("CreateLPAR", func() {
		var (
			payload       *LparProperties
			bytesResponse []byte
		)

		BeforeEach(func() {
			payload = &LparProperties{
				URI:                        "uri",
				CpcURI:                     "cpcuri",
				Class:                      "partition",
				Name:                       "lpar",
				Description:                "description",
				InitialIflProcessingWeight: 4096,
				MaximumMemory:              1024,
				ProcessorMode:              "shared",
				Type:                       PARTITION_TYPE_LINUX,
				AutoGenerateID:             true,
			}

			bytesResponse, _ = json.Marshal(payload)
		})

		Context("When CreateLPAR and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusCreated, bytesResponse, nil)
				uri, status, err := manager.CreateLPAR(lparid, payload)
				Expect(uri).ToNot(BeNil())
				Expect(status).To(Equal(201))
				Expect(err).To(BeNil())
			})
		})

		Context("When CreateLPAR returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusCreated, []byte("incorrect json bytes"), nil)
				_, _, err := manager.CreateLPAR(lparid, payload)

				Expect(*err).To(Equal(*unmarshalErr))
			})
		})

		Context("When CreateLPAR and returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, _, err := manager.CreateLPAR(lparid, payload)

				Expect(*err).To(Equal(*hmcErr))
			})
		})
	})

	Describe("UpdateLparProperties", func() {
		var (
			payload       *LparProperties
			bytesResponse []byte
		)

		BeforeEach(func() {
			payload = &LparProperties{
				URI:            "uri",
				CpcURI:         "cpcuri",
				Class:          "partition",
				Name:           "lpar",
				Description:    "description",
				Status:         PARTITION_STATUS_STARTING,
				Type:           PARTITION_TYPE_LINUX,
				ShortName:      "short_name",
				ID:             "id",
				AutoGenerateID: true,
			}

			bytesResponse, _ = json.Marshal(payload)
		})

		Context("When UpdateLparProperties and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, bytesResponse, nil)
				_, err := manager.UpdateLparProperties(lparid, payload)

				Expect(err).To(BeNil())
			})
		})

		Context("When UpdateLparProperties returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, []byte("incorrect json bytes"), hmcErr)
				_, err := manager.UpdateLparProperties(lparid, payload)

				Expect(err.Error()).ToNot(BeNil())
			})
		})

		Context("When UpdateLparProperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, err := manager.UpdateLparProperties(lparid, payload)

				Expect(*err).To(Equal(*hmcErr))
			})
		})
	})

	Describe("StartLPAR", func() {
		var (
			response                *StartStopLparResponse
			responseWithoutURI      *StartStopLparResponse
			bytesResponse           []byte
			bytesResponseWithoutURI []byte
		)

		BeforeEach(func() {
			response = &StartStopLparResponse{
				URI:     "uri",
				Message: "message",
			}
			responseWithoutURI = &StartStopLparResponse{
				URI:     "",
				Message: "message",
			}

			bytesResponse, _ = json.Marshal(response)
			bytesResponseWithoutURI, _ = json.Marshal(responseWithoutURI)
		})

		Context("When StartLPAR returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, bytesResponse, nil)
				rets, _, err := manager.StartLPAR(lparid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets).To(Equal(response.URI))
			})
		})

		Context("When StartLPAR returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.StartLPAR(lparid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(Equal(""))
			})
		})

		Context("When StartLPAR returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.StartLPAR(lparid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(Equal(""))
			})
		})

		Context("When StartLPAR returns hmcErr for WithoutURI", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, bytesResponseWithoutURI, hmcErr)
				rets, _, err := manager.StartLPAR(lparid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(Equal(""))
			})
		})
	})

	Describe("StopLPAR", func() {
		var (
			response                *StartStopLparResponse
			responseWithoutURI      *StartStopLparResponse
			bytesResponse           []byte
			bytesResponseWithoutURI []byte
		)

		BeforeEach(func() {
			response = &StartStopLparResponse{
				URI:     "uri",
				Message: "message",
			}
			responseWithoutURI = &StartStopLparResponse{
				URI:     "",
				Message: "message",
			}

			bytesResponse, _ = json.Marshal(response)
			bytesResponseWithoutURI, _ = json.Marshal(responseWithoutURI)
		})

		Context("When StopLPAR returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, bytesResponse, nil)
				rets, _, err := manager.StopLPAR(lparid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets).To(Equal(response.URI))
			})
		})

		Context("When StopLPAR returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.StopLPAR(lparid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(Equal(""))
			})
		})

		Context("When StopLPAR returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.StopLPAR(lparid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(Equal(""))
			})
		})

		Context("When StopLPAR returns error due to WithoutURI", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, bytesResponseWithoutURI, nil)
				rets, _, err := manager.StopLPAR(lparid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(Equal(""))
			})
		})
	})

	Describe("DeleteLPAR", func() {

		BeforeEach(func() {
			hmcErr = &HmcError{
				Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
				Message: "error message",
			}
		})

		Context("When DeleteLPAR returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				status, err := manager.DeleteLPAR(lparid)
				Expect(err).To(BeNil())
				Expect(status).ToNot(BeNil())
				Expect(status).To(Equal(204))
			})
		})

		Context("When DeleteLPAR returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, nil, hmcErr)
				_, err := manager.DeleteLPAR(lparid)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("MountIsoImage", func() {

		var (
			imageFile string
			insFile   string
		)

		BeforeEach(func() {
			imageFile = "imageFileName"
			file, _ := os.Create(imageFile)
			_, _ = file.WriteString("test data")
			insFile = "insFileName"
		})

		Context("When MountIsoImage returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.UploadRequestReturns(http.StatusNoContent, nil, nil)
				_, err := manager.MountIsoImage(lparid, imageFile, insFile)
				Expect(err).To(BeNil())
			})
		})

		Context("When MountIsoImage returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.UploadRequestReturns(http.StatusBadRequest, nil, hmcErr)
				_, err := manager.MountIsoImage(lparid, imageFile, insFile)

				Expect(*err).To(Equal(*hmcErr))
			})
		})
	})

	Describe("UnmountIsoImage", func() {

		Context("When UnmountIsoImage returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				_, err := manager.UnmountIsoImage(lparid)

				Expect(err).To(BeNil())
			})
		})

		Context("When UnmountIsoImage returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, nil, hmcErr)
				_, err := manager.UnmountIsoImage(lparid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})
	})

	Describe("ListNics", func() {
		var (
			nicresponse      []string
			nicbytesResponse []byte
			lparProps        LparProperties
		)

		BeforeEach(func() {
			nicresponse = []string{"uri1", "uri2"}
			lparProps.NicUris = nicresponse
			nicbytesResponse, _ = json.Marshal(lparProps)
		})

		Context("When ListNics returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, nicbytesResponse, nil)
				rets, _, err := manager.ListNics(lparid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
			})
		})
		Context("When ListNics returns error due to unmarshalErr", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.ListNics(lparid)
				Expect(err).ToNot(BeNil())
				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

	})

	Describe("AttachStorageGroupToPartition", func() {
		var storage *StorageGroupPayload
		BeforeEach(func() {
			storage = &StorageGroupPayload{
				StorageGroupURI: "uri",
			}
		})
		Context("When AttachStorageGroupToPartition returns correctly", func() {
			It("Check the results Succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, nil, nil)
				rets, _ := manager.AttachStorageGroupToPartition(lparid, storage)
				Expect(rets).ToNot(BeNil())
			})

		})
		Context("When AttachStorageGroupToPartition returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, nil, hmcErr)
				rets, err := manager.AttachStorageGroupToPartition(lparid, storage)
				Expect(rets).To(Equal(200))
				Expect(err).To(Equal(hmcErr))
			})
		})
		Context("When AttachStorageGroupToPartition returns correctly with status 204", func() {
			It("check the response has no content", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				rets, err := manager.AttachStorageGroupToPartition(lparid, storage)
				Expect(err).To(BeNil())
				Expect(rets).To(Equal(204))
			})
		})
	})
	Describe("DetachStorageGroupToPartition", func() {
		var storage *StorageGroupPayload
		BeforeEach(func() {
			storage = &StorageGroupPayload{
				StorageGroupURI: "uri",
			}
		})
		Context("When DetachStorageGroupToPartition returns correctly", func() {
			It("Check the results Succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, nil, nil)
				rets, _ := manager.DetachStorageGroupToPartition(lparid, storage)
				Expect(rets).To(Equal(200))
			})

		})
		Context("When DetachStorageGroupToPartition returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadGateway, nil, hmcErr)
				rets, err := manager.DetachStorageGroupToPartition(lparid, storage)
				Expect(rets).To(Equal(502))
				Expect(err).To(Equal(hmcErr))
			})
		})
		Context("When DetachStorageGroupToPartition returns correctly with status 204", func() {
			It("check the response has no content", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				rets, err := manager.DetachStorageGroupToPartition(lparid, storage)
				Expect(err).To(BeNil())
				Expect(rets).To(Equal(204))
			})
		})
	})

	Describe("AttachCryptoToPartition", func() {
		var cryptoConfig *CryptoConfig
		BeforeEach(func() {
			cryptoConfig = &CryptoConfig{
				CryptoAdapterUris: []string{"uri"},
				CryptoDomainConfigurations: []DomainInfo{
					{
						DomainIdx:  1,
						AccessMode: "control",
					},
				},
			}
		})
		Context("When AttachCryptoToPartition returns correctly", func() {
			It("Check the results Succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, nil, nil)
				rets, _ := manager.AttachCryptoToPartition(lparid, cryptoConfig)
				Expect(rets).ToNot(BeNil())
			})

		})
		Context("When AttachCryptoToPartition returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, nil, hmcErr)
				rets, err := manager.AttachCryptoToPartition(lparid, cryptoConfig)
				Expect(rets).To(Equal(200))
				Expect(err).To(Equal(hmcErr))
			})
		})
		Context("When AttachCryptoToPartition returns correctly with status 204", func() {
			It("check the response has no content", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				rets, err := manager.AttachCryptoToPartition(lparid, cryptoConfig)
				Expect(err).To(BeNil())
				Expect(rets).To(Equal(204))
			})
		})
	})

})
