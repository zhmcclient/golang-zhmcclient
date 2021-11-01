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
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

var _ = Describe("client", func() {
	Describe("IsExpectedHttpStatus", func() {

		Context(`when status is http.StatusOK
								http.StatusCreated,
								http.StatusAccepted,
								http.StatusNoContent,
								http.StatusPartialContent,
								http.StatusBadRequest,
								http.StatusForbidden,
								http.StatusNotFound,
								http.StatusConflict,
								http.StatusInternalServerError,
								http.StatusServiceUnavailable,
				`, func() {

			It("returns true", func() {
				var staus = http.StatusOK
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
			It("returns true", func() {
				var staus = http.StatusCreated
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
			It("returns true", func() {
				var staus = http.StatusAccepted
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
			It("returns true", func() {
				var staus = http.StatusNoContent
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
			It("returns true", func() {
				var staus = http.StatusPartialContent
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
			It("returns true", func() {
				var staus = http.StatusBadRequest
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
			It("returns true", func() {
				var staus = http.StatusForbidden
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
			It("returns true", func() {
				var staus = http.StatusNotFound
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
			It("returns true", func() {
				var staus = http.StatusConflict
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
			It("returns true", func() {
				var staus = http.StatusInternalServerError
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
			It("returns true", func() {
				var staus = http.StatusServiceUnavailable
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(true))
			})
		})

		Context("when status is http.StatusUnauthorized", func() {
			It("returns false", func() {
				var staus = http.StatusUnauthorized
				ret := IsExpectedHttpStatus(staus)
				Expect(ret).To(Equal(false))
			})
		})
	})

	Describe("GetEndpointURLFromString", func() {

		Context("When url string is correct", func() {
			It("returns a url", func() {
				urlStr := "https://127.0.0.1:433/api"
				url, err := GetEndpointURLFromString(urlStr)

				Expect(err).To(BeNil())
				Expect(url).ToNot(BeNil())
			})
		})

		Context("When url string does not have schema", func() {
			It("returns error", func() {
				urlStr := "/api"
				url, err := GetEndpointURLFromString(urlStr)

				Expect(err).ToNot(BeNil())
				Expect(url).To(BeNil())
			})
		})

		Context("When url string has insecure schema", func() {
			It("returns error", func() {
				urlStr := "http://127.0.0.1:433/api"
				url, err := GetEndpointURLFromString(urlStr)

				Expect(err).ToNot(BeNil())
				Expect(url).To(BeNil())
			})
		})

		Context("When url string has incorrect schema", func() {
			It("returns error", func() {
				urlStr := "ftp://127.0.0.1:433/api"
				url, err := GetEndpointURLFromString(urlStr)

				Expect(err).ToNot(BeNil())
				Expect(url).To(BeNil())
			})
		})

	})
})
