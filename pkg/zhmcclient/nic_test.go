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

var _ = Describe("Nic", func() {
	var (
		manager              *NicManager
		fakeClient           *fakes.ClientAPI
		lparid               string
		nicid                string
		url                  *url.URL
		hmcErr, unmarshalErr *HmcError
	)

	BeforeEach(func() {
		fakeClient = &fakes.ClientAPI{}
		manager = NewNicManager(fakeClient)

		url, _ = url.Parse("https://127.0.0.1:443")
		lparid = "lparid"
		nicid = "nicid"

		hmcErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
			Message: "error message",
		}

		unmarshalErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_UNMARSHAL_FAIL),
			Message: "invalid character 'i' looking for beginning of value",
		}
	})

	Describe("CreateNic", func() {
		var (
			response                *NicCreateResponse
			responseWithoutURI      *NicCreateResponse
			payload                 *NIC
			bytesResponse           []byte
			bytesResponseWithoutURI []byte
		)

		BeforeEach(func() {
			response = &NicCreateResponse{
				URI: "uri",
			}
			responseWithoutURI = &NicCreateResponse{
				URI: "",
			}
			payload = &NIC{
				ID:                    "id",
				URI:                   "uri",
				Parent:                "parent_uri",
				Class:                 "nic",
				Name:                  "name",
				Description:           "description",
				DeviceNumber:          "device_number",
				NetworkAdapterPortURI: "adapter_uri",
				VirtualSwitchUriType:  "",
				VirtualSwitchURI:      "",
				Type:                  NIC_TYPE_ROCE,
				SscManagmentNIC:       false,
				SscIpAddressType:      SSC_IP_TYPE_IPV4,
				SscIpAddress:          "",
				VlanID:                1024,
				MacAddress:            "",
				SscMaskPrefix:         "",
				VlanType:              VLAN_TYPE_ENFORCED,
			}

			bytesResponse, _ = json.Marshal(response)
			bytesResponseWithoutURI, _ = json.Marshal(responseWithoutURI)
		})

		Context("When CreateNic returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusCreated, bytesResponse, nil)
				rets, _, err := manager.CreateNic(lparid, payload)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(Equal(""))
				Expect(rets).To(Equal(response.URI))
			})
		})

		Context("When CreateNic returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.CreateNic(lparid, payload)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(Equal(""))
			})
		})

		Context("When CreateNic returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusCreated, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.CreateNic(lparid, payload)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(Equal(""))
			})
		})

		Context("When CreateNic returns hmcErr for WithoutURI", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, bytesResponseWithoutURI, hmcErr)
				rets, _, err := manager.CreateNic(lparid, payload)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(Equal(""))
			})
		})
	})

	Describe("DeleteNic", func() {

		Context("When DeleteNic returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				_, err := manager.DeleteNic(nicid)

				Expect(err).To(BeNil())
			})
		})

		Context("When DeleteNic returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, nil, hmcErr)
				_, err := manager.DeleteNic(nicid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})
	})

	Describe("GetNicProperties", func() {
		var (
			response      *NIC
			bytesResponse []byte
		)

		BeforeEach(func() {
			response = &NIC{
				ID:                    "id",
				URI:                   "uri",
				Parent:                "parent_uri",
				Class:                 "nic",
				Name:                  "name",
				Description:           "description",
				DeviceNumber:          "device_number",
				NetworkAdapterPortURI: "adapter_uri",
				VirtualSwitchUriType:  "",
				VirtualSwitchURI:      "",
				Type:                  NIC_TYPE_ROCE,
				SscManagmentNIC:       false,
				SscIpAddressType:      SSC_IP_TYPE_IPV4,
				SscIpAddress:          "",
				VlanID:                1024,
				MacAddress:            "",
				SscMaskPrefix:         "",
				VlanType:              VLAN_TYPE_ENFORCED,
			}
			bytesResponse, _ = json.Marshal(response)
		})
		Context("When GetNicProperties returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytesResponse, nil)
				rets, _, err := manager.GetNicProperties(nicid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets.URI).To(Equal(response.URI))
			})
		})
		Context("When GetNicProperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.GetNicProperties(nicid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})
		Context("When GetNicProperties returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.GetNicProperties(nicid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})
		Context("When GetNicProperties returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytesResponse, nil)
				rets, _, err := manager.GetNicProperties(nicid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

	})

	Describe("UpdateNicProperties", func() {
		var (
			response      *NIC
			bytesResponse []byte
			updateProps   *NIC
		)

		BeforeEach(func() {
			updateProps = &NIC{
				DeviceNumber: "device_number_updated",
			}
			response = &NIC{
				ID:                    "id",
				URI:                   "uri",
				Parent:                "parent_uri",
				Class:                 "nic",
				Name:                  "name",
				Description:           "description",
				DeviceNumber:          "device_number",
				NetworkAdapterPortURI: "adapter_uri",
				VirtualSwitchUriType:  "",
				VirtualSwitchURI:      "",
				Type:                  NIC_TYPE_ROCE,
				SscManagmentNIC:       false,
				SscIpAddressType:      SSC_IP_TYPE_IPV4,
				SscIpAddress:          "",
				VlanID:                1024,
				MacAddress:            "",
				SscMaskPrefix:         "",
				VlanType:              VLAN_TYPE_ENFORCED,
			}
			bytesResponse, _ = json.Marshal(response)
		})
		Context("When UpdateNicProperties returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, bytesResponse, nil)
				_, err := manager.UpdateNicProperties(nicid, updateProps)

				Expect(err).To(BeNil())
			})
		})
		Context("When UpdateNicProperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				_, err := manager.UpdateNicProperties(nicid, updateProps)
				Expect(*err).To(Equal(*hmcErr))
			})
		})
		Context("When UpdateNicProperties returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytesResponse, nil)
				_, err := manager.UpdateNicProperties(nicid, updateProps)

				Expect(err).ToNot(BeNil())
			})
		})

	})
})
