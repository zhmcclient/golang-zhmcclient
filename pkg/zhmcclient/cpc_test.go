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
	"errors"
	"net/http"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient/fakes"
)

var _ = Describe("CPC", func() {
	Describe("ListCPCs", func() {
		var (
			manager    *CpcManager
			fakeClient *fakes.ClientAPI

			cpcs  []CPC
			url   *url.URL
			bytes []byte
		)

		BeforeEach(func() {
			fakeClient = &fakes.ClientAPI{}
			manager = NewCpcManager(fakeClient)

			cpcs = []CPC{
				{
					URI:                 "object-uri1",
					Name:                "name1",
					Status:              "status1",
					HasAcceptableStatus: "has-unacceptable-status1",
					DpmEnabled:          false,
					SeVersion:           "se-version1",
				},
				{
					URI:                 "object-uri2",
					Name:                "name2",
					Status:              "status2",
					HasAcceptableStatus: "has-unacceptable-status2",
					DpmEnabled:          true,
					SeVersion:           "se-version2",
				},
			}

			url, _ = url.Parse("https://127.0.0.1:443")
			bytes, _ = json.Marshal(cpcs)
		})

		Context("When list cpcs and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, err := manager.ListCPCs(nil)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets[0]).To(Equal(cpcs[0]))
			})
		})

		Context("When list cpcs and returns error", func() {
			It("check the error happened", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, errors.New("error"))
				rets, err := manager.ListCPCs(nil)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When list cpcs and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), errors.New("error"))
				rets, err := manager.ListCPCs(nil)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When list cpcs and returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytes, nil)
				rets, err := manager.ListCPCs(nil)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})
})
