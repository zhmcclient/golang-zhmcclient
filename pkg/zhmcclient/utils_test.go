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
	"fmt"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

var _ = Describe("utils", func() {
	Describe("BuildUrlFromQuery", func() {
		var (
			url    *url.URL
			query0 map[string]string
			query1 map[string]string
			query2 map[string]string
		)

		BeforeEach(func() {
			query0 = map[string]string{}
			query1 = map[string]string{
				"name": "lpar1",
			}
			query2 = map[string]string{
				"name": "lpar1",
				"type": "dpm",
			}

			url, _ = url.Parse("https://127.0.0.1:80/api")
		})

		Context("When nil query passed in", func() {
			It("returns same uri", func() {
				ret, _ := url.Parse(url.String())
				ret, err := BuildUrlFromQuery(ret, nil)

				Expect(err).To(BeNil())
				Expect(ret.String()).To(Equal(url.String()))
			})
		})

		Context("When empty query passed in", func() {
			It("returns same uri", func() {
				ret, _ := url.Parse(url.String())
				ret, err := BuildUrlFromQuery(ret, query0)

				Expect(err).To(BeNil())
				Expect(ret.String()).To(Equal(url.String()))
			})
		})

		Context("When 1 query passed in", func() {
			It("returns correct uri", func() {
				ret, _ := url.Parse(url.String())
				ret, err := BuildUrlFromQuery(ret, query1)

				Expect(err).To(BeNil())
				Expect(ret.String()).To(Equal(url.String() + "?name=lpar1"))
			})
		})

		Context("When 2 query passed in", func() {
			It("returns correct uri", func() {
				ret, _ := url.Parse(url.String())
				ret, err := BuildUrlFromQuery(ret, query2)

				Expect(err).To(BeNil())
				Expect(ret.String()).To(Equal(url.String() + "?name=lpar1&type=dpm"))
			})
		})
	})

	Describe("GenerateErrorFromResponse", func() {
		var (
			status     int
			errMessage string

			errFull           *ErrorBody
			errWithoutMessage *ErrorBody

			errByte               []byte
			errWithoutMessageByte []byte
		)

		BeforeEach(func() {
			errMessage = "error message"

			errFull = &ErrorBody{
				Reason:  1,
				Message: errMessage,
			}
			errWithoutMessage = &ErrorBody{
				Reason: 1,
			}
			errByte, _ = json.Marshal(errFull)
			errWithoutMessageByte, _ = json.Marshal(errWithoutMessage)
		})

		Context("message is not empty", func() {
			It("returns message directly", func() {
				status = 400
				rets := GenerateErrorFromResponse(status, errByte)

				Expect(rets).ToNot(BeNil())
				Expect(rets.Error()).To(Equal(errMessage))
			})
		})

		Context("message is empty", func() {
			It("returns message according to reason for 400, 1", func() {
				status = 400
				rets := GenerateErrorFromResponse(status, errWithoutMessageByte)

				Expect(rets).ToNot(BeNil())
				Expect(rets.Error()).To(Equal("The request included an unrecognized or unsupported query parameter."))
			})

			It("returns message according to reason for 403, 1", func() {
				status = 403
				rets := GenerateErrorFromResponse(status, errWithoutMessageByte)

				Expect(rets).ToNot(BeNil())
				Expect(rets.Error()).To(Equal("The user under which the API request was authenticated does not have the required authority to perform the requested action."))
			})

			It("returns message according to reason for 404, 1", func() {
				status = 404
				rets := GenerateErrorFromResponse(status, errWithoutMessageByte)

				Expect(rets).ToNot(BeNil())
				Expect(rets.Error()).To(Equal("The request URI does not designate an existing resource of the expected type, or designates a resource for which the API user does not have object-access permission. For URIs that contain object ID and/or element ID components, this reason code may be used for issues accessing the resource identified by the first (leftmost) such ID in the URI."))
			})

			It("returns message according to reason for 409, 1", func() {
				status = 409
				rets := GenerateErrorFromResponse(status, errWithoutMessageByte)

				Expect(rets).ToNot(BeNil())
				Expect(rets.Error()).To(Equal("The operation cannot be performed because the object designated by the request URI is not in the correct state."))
			})

			It("returns message according to reason for 500", func() {
				status = 500
				rets := GenerateErrorFromResponse(status, errWithoutMessageByte)

				Expect(rets).ToNot(BeNil())
				Expect(rets.Error()).To(Equal("An internal processing error has occurred and no additional details are documented."))
			})

			It("returns message according to reason for 503, 1", func() {
				status = 503
				rets := GenerateErrorFromResponse(status, errWithoutMessageByte)

				Expect(rets).ToNot(BeNil())
				Expect(rets.Error()).To(Equal("The request could not be processed because the HMC is not currently communicating with an SE needed to perform the requested operation."))
			})

			It("returns message according to unknown reason or status code", func() {
				status = 401
				rets := GenerateErrorFromResponse(status, errWithoutMessageByte)

				Expect(rets).ToNot(BeNil())
				Expect(rets.Error()).To(Equal("HTTP Error: " + fmt.Sprint(status)))
			})
		})

	})
})
