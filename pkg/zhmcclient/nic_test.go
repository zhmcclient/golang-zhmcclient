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

		Context("When CreateNic and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusCreated, bytesResponse, nil)
				rets, _, err := manager.CreateNic(lparid, payload)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(Equal(""))
				Expect(rets).To(Equal(response.URI))
			})
		})

		Context("When CreateNic and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, hmcErr)
				rets, _, err := manager.CreateNic(lparid, payload)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(Equal(""))
			})
		})

		Context("When CreateNic and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusCreated, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.CreateNic(lparid, payload)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(Equal(""))
			})
		})

		Context("When CreateNic and no URI responded", func() {
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

		Context("When DeleteNic and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				_, err := manager.DeleteNic(nicid)

				Expect(err).To(BeNil())
			})
		})

		Context("When DeleteNic and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, nil, hmcErr)
				_, err := manager.DeleteNic(nicid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})
	})
})
