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
	"crypto/tls"
	"io"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

var (
	client *Client
	hmcErr *HmcError
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

	Describe("ClientTesting", func() {
		BeforeEach(func() {
			client = &Client{}
			hmcErr = &HmcError{
				Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
				Message: "error message",
			}
		})
		Context("When client is empty", func() {
			It("check the result of CloneEndPointURL", func() {
				url := client.CloneEndpointURL()
				Expect(url).To(BeNil())
			})
		})

		Context("When client is empty, IsLogon is executed with true should throw  error as false", func() {
			It("checks the result of IsLogon", func() {
				err := client.IsLogon(true)
				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(false))
			})
		})
		Context("When client is empty, IsLogon is executed with false should throw error as false", func() {
			It("checks the result of IsLogon", func() {
				err := client.IsLogon(false)
				Expect(err).ToNot(BeNil())
				Expect(err).To(Equal(false))
			})
		})
		Context("When client is empty, Logon should throw error", func() {
			It("checks the result of Logon", func() {
				err := client.Logon()
				Expect(err).ToNot(BeNil())
			})
		})
		Context("When client is empty, Update TraceOn value", func() {
			It("check the result of Client", func() {
				var outputStream io.Writer
				client.TraceOn(outputStream)
				Expect(client).ToNot(BeNil())
			})
		})
		Context("When client is empty, Update Traceoff value without passing arguments", func() {
			It("check the result of Client", func() {
				client.TraceOff()
				Expect(client).ToNot(BeNil())
			})
		})
		Context("When client is empty, Update Traceoff value by passing arguments", func() {
			It("check the result of Client", func() {
				client.SetSkipCertVerify(false)
				Expect(client).ToNot(BeNil())
			})
		})
	})

	Describe("NewClient", func() {
		BeforeEach(func() {
			hmcErr = &HmcError{
				Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
				Message: "error message",
			}
		})
		Context("When NewClient is Executed", func() {
			It("Check the result of NewClient", func() {
				var endpoint string
				opts := &Options{
					SkipCert: false,
					Username: "",
					Password: "",
				}
				var logger Logger
				client, err := NewClient(endpoint, opts, logger)
				Expect(client).To(BeNil())
				Expect(err.Error()).ToNot(BeNil())
			})
		})
	})

	Describe("ChangePassword", func() {
		BeforeEach(func() {
			hmcErr = &HmcError{
				Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
				Message: "error message",
			}
		})
		Context("When ChangePassword is Executed", func() {
			It("Check the result of ChangePassword", func() {
				var endpoint string
				opts := &Options{
					SkipCert: false,
					Username: "",
					Password: "",
				}
				err := ChangePassword(endpoint, opts, "")
				Expect(err.Error()).ToNot(BeNil())
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
					CaCert:   "data.pem",
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
