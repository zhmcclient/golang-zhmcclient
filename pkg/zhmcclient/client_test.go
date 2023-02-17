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
	"crypto/tls"
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

	Describe("SetCertificate", func() {
		Context("When skipcert is false", func() {
			It("returns tls config without CaCert", func() {
				opts := &Options{
					SkipCert: false,
				}
				tlsConfig, _ := SetCertificate(opts, &tls.Config{})
				Expect(tlsConfig).To(BeNil())
			})
		})

		Context("When skipcert is true", func() {
			It("returns tls config with root CaCert", func() {
				opts := &Options{
					SkipCert: true,
					CaCert:   "data.der",
				}
				tlsConfig, err := SetCertificate(opts, &tls.Config{})
				Expect(tlsConfig).ToNot(BeNil())
				Expect(err).To(BeNil())
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

	Describe("NeedLogon", func() {

		Context("status is 401", func() {
			It("returns true", func() {
				ret := NeedLogon(401, 0)
				Expect(ret).To(Equal(true))
			})
		})

		Context("status is 403", func() {
			It("returns true when reason is 4 0r 5", func() {
				ret := NeedLogon(403, 4)
				Expect(ret).To(Equal(true))

				ret = NeedLogon(403, 5)
				Expect(ret).To(Equal(true))
			})
		})

		Context("status is 403", func() {
			It("returns false when reason is not 4 or 5", func() {
				ret := NeedLogon(403, 0)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(403, -1)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(403, 1)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(403, 2)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(403, 3)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(403, 6)
				Expect(ret).To(Equal(false))
			})
		})

		Context("status is not 401 or 403", func() {
			It("returns false", func() {
				ret := NeedLogon(200, 0)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(201, 0)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(202, 0)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(204, 0)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(206, 0)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(400, 0)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(404, 0)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(409, 0)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(500, 0)
				Expect(ret).To(Equal(false))

				ret = NeedLogon(503, 0)
				Expect(ret).To(Equal(false))
			})
		})
	})
})
