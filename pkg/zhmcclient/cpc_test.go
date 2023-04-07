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

var _ = Describe("CPC", func() {
	Describe("ListCPCs", func() {
		var (
			manager    *CpcManager
			fakeClient *fakes.ClientAPI

			cpcs      []CPC
			cpcsArray = &CpcsArray{}
			url       *url.URL
			bytes     []byte

			hmcErr, unmarshalErr *HmcError
		)

		BeforeEach(func() {
			fakeClient = &fakes.ClientAPI{}
			manager = NewCpcManager(fakeClient)

			cpcs = []CPC{
				{
					URI:                 "object-uri1",
					Name:                "name1",
					Status:              "status1",
					HasAcceptableStatus: false,
					DpmEnabled:          false,
					SeVersion:           "se-version1",
				},
				{
					URI:                 "object-uri2",
					Name:                "name2",
					Status:              "status2",
					HasAcceptableStatus: false,
					DpmEnabled:          true,
					SeVersion:           "se-version2",
				},
			}

			cpcsArray = &CpcsArray{
				cpcs,
			}

			url, _ = url.Parse("https://127.0.0.1:443")
			bytes, _ = json.Marshal(cpcsArray)

			hmcErr = &HmcError{
				Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
				Message: "error message",
			}

			unmarshalErr = &HmcError{
				Reason:  int(ERR_CODE_HMC_UNMARSHAL_FAIL),
				Message: "invalid character 'i' looking for beginning of value",
			}
		})

		Context("When ListCPCs returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, _, err := manager.ListCPCs(nil)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets[0]).To(Equal(cpcs[0]))
			})
		})

		Context("When ListCPCs returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, hmcErr)
				rets, _, err := manager.ListCPCs(nil)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When ListCPCs returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.ListCPCs(nil)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When ListCPCs returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytes, nil)
				rets, _, err := manager.ListCPCs(nil)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})
})
