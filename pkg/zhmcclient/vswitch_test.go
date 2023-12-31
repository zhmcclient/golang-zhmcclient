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

var _ = Describe("Virtual Switch", func() {
	var (
		manager              *VirtualSwitchManager
		fakeClient           *fakes.ClientAPI
		cpcid                string
		vswitchId            string
		url                  *url.URL
		hmcErr, unmarshalErr *HmcError
	)

	BeforeEach(func() {
		fakeClient = &fakes.ClientAPI{}
		manager = NewVirtualSwitchManager(fakeClient)

		url, _ = url.Parse("https://127.0.0.1:443")
		cpcid = "cpcid"
		vswitchId = "vswitchId"

		hmcErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
			Message: "error message",
		}

		unmarshalErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_UNMARSHAL_FAIL),
			Message: "invalid character 'i' looking for beginning of value",
		}
	})

	Describe("ListVirtualSwitches", func() {
		var (
			virtualSwitches      []VirtualSwitch
			virtualSwitchesArray VirtualSwitchesArray
			bytes                []byte
		)

		BeforeEach(func() {
			virtualSwitches = []VirtualSwitch{
				{
					URI:  "uri1",
					Name: "name1",
					Type: VIRTUALSWITCH_TYPE_HIPERSOCKET,
				},
				{
					URI:  "uri2",
					Name: "name2",
					Type: VIRTUALSWITCH_TYPE_HIPERSOCKET,
				},
			}

			virtualSwitchesArray = VirtualSwitchesArray{
				virtualSwitches,
			}
			bytes, _ = json.Marshal(virtualSwitchesArray)
		})

		Context("When ListVirtualSwitches returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, _, err := manager.ListVirtualSwitches(cpcid, nil)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets[0]).To(Equal(virtualSwitches[0]))
			})
		})

		Context("When ListVirtualSwitches returns error due hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, hmcErr)
				rets, _, err := manager.ListVirtualSwitches(cpcid, nil)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When ListVirtualSwitches returns error due to unmarshalerr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.ListVirtualSwitches(cpcid, nil)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When ListVirtualSwitches returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytes, nil)
				rets, _, err := manager.ListVirtualSwitches(cpcid, nil)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("GetVirtualSwitchProperties", func() {
		var (
			virtualSwitch VirtualSwitchProperties
			bytes         []byte
		)

		BeforeEach(func() {
			virtualSwitch = VirtualSwitchProperties{
				URI:         "uri1",
				Name:        "name1",
				Type:        VIRTUALSWITCH_TYPE_HIPERSOCKET,
				ID:          "id",
				Parent:      "parent",
				Class:       "virtual-switch",
				Description: "This is a test hipersocket",
				AdapterURI:  "adapter_uri",
				Port:        1234567,
				VNicUris:    []string{"uri1", "uri2"},
			}
			bytes, _ = json.Marshal(virtualSwitch)
		})

		Context("When GetVirtualSwitchProperties return correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, _, err := manager.GetVirtualSwitchProperties(vswitchId)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(*rets).To(Equal(virtualSwitch))
			})
		})

		Context("When GetVirtualSwitchProperties returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, hmcErr)
				rets, _, err := manager.GetVirtualSwitchProperties(vswitchId)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When GetVirtualSwitchProperties returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.GetVirtualSwitchProperties(vswitchId)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When GetVirtualSwitchProperties returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytes, nil)
				rets, _, err := manager.GetVirtualSwitchProperties(vswitchId)
				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})
})
