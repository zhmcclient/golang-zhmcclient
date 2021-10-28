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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

var _ = Describe("utils", func() {
	Describe("BuildUrlFromUri", func() {
		var (
			uri    string
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

			uri = "https://127.0.0.1:80/api"
		})

		Context("When nil query passed in", func() {
			It("returns same uri", func() {
				url, err := BuildUrlFromUri(uri, nil)

				Expect(err).To(BeNil())
				Expect(url.String()).To(Equal(uri))
			})
		})

		Context("When empty query passed in", func() {
			It("returns same uri", func() {
				url, err := BuildUrlFromUri(uri, query0)

				Expect(err).To(BeNil())
				Expect(url.String()).To(Equal(uri))
			})
		})

		Context("When 1 query passed in", func() {
			It("returns correct uri", func() {
				url, err := BuildUrlFromUri(uri, query1)

				Expect(err).To(BeNil())
				Expect(url.String()).To(Equal(uri + "?name=lpar1"))
			})
		})

		Context("When 2 query passed in", func() {
			It("returns correct uri", func() {
				url, err := BuildUrlFromUri(uri, query2)

				Expect(err).To(BeNil())
				Expect(url.String()).To(Equal(uri + "?name=lpar1&type=dpm"))
			})
		})
	})
})
