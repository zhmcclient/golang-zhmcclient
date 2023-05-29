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
	"net/url"
	"os"

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
				ret = BuildUrlFromQuery(ret, nil)
				Expect(ret.String()).To(Equal(url.String()))
			})
		})

		Context("When empty query passed in", func() {
			It("returns same uri", func() {
				ret, _ := url.Parse(url.String())
				ret = BuildUrlFromQuery(ret, query0)
				Expect(ret.String()).To(Equal(url.String()))
			})
		})

		Context("When 1 query passed in", func() {
			It("returns correct uri", func() {
				ret, _ := url.Parse(url.String())
				ret = BuildUrlFromQuery(ret, query1)
				Expect(ret.String()).To(Equal(url.String() + "?name=lpar1"))
			})
		})

		Context("When 2 query passed in", func() {
			It("returns correct uri", func() {
				ret, _ := url.Parse(url.String())
				ret = BuildUrlFromQuery(ret, query2)
				Expect(ret.String()).To(Equal(url.String() + "?name=lpar1&type=dpm"))
			})
		})
	})

	Describe("RetrieveBytes", func() {
		var (
			filename string
		)

		BeforeEach(func() {
			filename = "data.txt"
			file, _ := os.Create(filename)
			_, _ = file.WriteString("test data")
		})

		Context("When no file is passed in", func() {
			It("returns error", func() {
				ret, err := RetrieveBytes("")

				Expect(err).ToNot(BeNil())
				Expect(ret).To(BeNil())
			})
		})

		Context("When file is passed in", func() {
			It("returns byte stream", func() {
				ret, err := RetrieveBytes(filename)
				Expect(err).To(BeNil())
				Expect(len(ret)).ToNot(Equal(0))
			})
		})
	})

	Describe("GenerateErrorFromResponse", func() {
		var (
			errMessage          string
			errFull, errUnknown *HmcError
			errByte             []byte
		)

		BeforeEach(func() {
			errMessage = "error message"

			errFull = &HmcError{
				Reason:  1,
				Message: errMessage,
			}
			errUnknown = &HmcError{
				Reason:  -1,
				Message: "Unknown error.",
			}

			errByte, _ = json.Marshal(errFull)
		})

		Context("message is normal", func() {
			It("returns message directly", func() {
				rets := GenerateErrorFromResponse(errByte)

				Expect(rets).ToNot(BeNil())
				Expect(rets).To(Equal(errFull))
			})
		})

		Context("message is not valid", func() {
			It("returns unknown error", func() {
				rets := GenerateErrorFromResponse([]byte("test"))

				Expect(rets).ToNot(BeNil())
				Expect(*rets).To(Equal(*errUnknown))
			})
		})
	})
})
