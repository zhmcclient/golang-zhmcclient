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

var _ = Describe("LPAR", func() {
	var (
		manager    *LparManager
		fakeClient *fakes.ClientAPI
		cpcid      string
		lparid     string
		url        *url.URL
	)

	BeforeEach(func() {
		fakeClient = &fakes.ClientAPI{}
		manager = NewLparManager(fakeClient)

		url, _ = url.Parse("https://127.0.0.1:443")
		cpcid = "cpcid"
		lparid = "lparid"
	})

	Describe("ListLPARs", func() {
		var (
			lpars []LPAR
			bytes []byte
		)

		BeforeEach(func() {
			lpars = []LPAR{
				{
					URI:    "object-uri1",
					Name:   "name1",
					Status: "status",
					Type:   "type",
				},
				{
					URI:    "object-uri2",
					Name:   "name2",
					Status: "status",
					Type:   "type",
				},
			}
			bytes, _ = json.Marshal(lpars)
		})

		Context("When list lpars and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, err := manager.ListLPARs(cpcid, nil)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets[0]).To(Equal(lpars[0]))
			})
		})

		Context("When list lpars and returns error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, errors.New("error"))
				rets, err := manager.ListLPARs(cpcid, nil)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When list lpars and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), errors.New("error"))
				rets, err := manager.ListLPARs(cpcid, nil)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When list lpars and returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytes, nil)
				rets, err := manager.ListLPARs(cpcid, nil)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("GetLparProperties", func() {
		var (
			response      *LparProperties
			bytesResponse []byte
		)

		BeforeEach(func() {
			response = &LparProperties{
				URI:            "uri",
				CpcURI:         "cpcuri",
				Class:          "partition",
				Name:           "lpar",
				Description:    "description",
				Status:         PARTITION_STATUS_STARTING,
				Type:           PARTITION_TYPE_LINUX,
				ShortName:      "short_name",
				ID:             "id",
				AutoGenerateID: true,
			}

			bytesResponse, _ = json.Marshal(response)
		})

		Context("When GetLparProperties and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytesResponse, nil)
				rets, err := manager.GetLparProperties(lparid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets.URI).To(Equal(response.URI))
			})
		})

		Context("When GetLparProperties and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, errors.New("error"))
				rets, err := manager.GetLparProperties(lparid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When GetLparProperties and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), errors.New("error"))
				rets, err := manager.GetLparProperties(lparid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("UpdateLparProperties", func() {
		var (
			payload       *LparProperties
			bytesResponse []byte
		)

		BeforeEach(func() {
			payload = &LparProperties{
				URI:            "uri",
				CpcURI:         "cpcuri",
				Class:          "partition",
				Name:           "lpar",
				Description:    "description",
				Status:         PARTITION_STATUS_STARTING,
				Type:           PARTITION_TYPE_LINUX,
				ShortName:      "short_name",
				ID:             "id",
				AutoGenerateID: true,
			}

			bytesResponse, _ = json.Marshal(payload)
		})

		Context("When UpdateLparProperties and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, bytesResponse, nil)
				err := manager.UpdateLparProperties(lparid, payload)

				Expect(err).To(BeNil())
			})
		})

		Context("When UpdateLparProperties and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, errors.New("error"))
				err := manager.UpdateLparProperties(lparid, payload)

				Expect(err).ToNot(BeNil())
			})
		})

		Context("When UpdateLparProperties and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), errors.New("error"))
				err := manager.UpdateLparProperties(lparid, nil)

				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("StartLPAR", func() {
		var (
			response                *StartStopLparResponse
			responseWithoutURI      *StartStopLparResponse
			bytesResponse           []byte
			bytesResponseWithoutURI []byte
		)

		BeforeEach(func() {
			response = &StartStopLparResponse{
				URI:     "uri",
				Message: "message",
			}
			responseWithoutURI = &StartStopLparResponse{
				URI:     "",
				Message: "message",
			}

			bytesResponse, _ = json.Marshal(response)
			bytesResponseWithoutURI, _ = json.Marshal(responseWithoutURI)
		})

		Context("When start lpar and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, bytesResponse, nil)
				rets, err := manager.StartLPAR(lparid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets).To(Equal(response.URI))
			})
		})

		Context("When start lpar and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, errors.New("error"))
				rets, err := manager.StartLPAR(lparid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(Equal(""))
			})
		})

		Context("When start lpar and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, []byte("incorrect json bytes"), errors.New("error"))
				rets, err := manager.StartLPAR(lparid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(Equal(""))
			})
		})

		Context("When start lpar and no URI responded", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, bytesResponseWithoutURI, errors.New("error"))
				rets, err := manager.StartLPAR(lparid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(Equal(""))
			})
		})
	})

	Describe("StopLPAR", func() {
		var (
			response                *StartStopLparResponse
			responseWithoutURI      *StartStopLparResponse
			bytesResponse           []byte
			bytesResponseWithoutURI []byte
		)

		BeforeEach(func() {
			response = &StartStopLparResponse{
				URI:     "uri",
				Message: "message",
			}
			responseWithoutURI = &StartStopLparResponse{
				URI:     "",
				Message: "message",
			}

			bytesResponse, _ = json.Marshal(response)
			bytesResponseWithoutURI, _ = json.Marshal(responseWithoutURI)
		})

		Context("When stop lpar and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, bytesResponse, nil)
				rets, err := manager.StopLPAR(lparid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets).To(Equal(response.URI))
			})
		})

		Context("When stop lpar and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, bytesResponse, errors.New("error"))
				rets, err := manager.StopLPAR(lparid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(Equal(""))
			})
		})

		Context("When stop lpar and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, []byte("incorrect json bytes"), errors.New("error"))
				rets, err := manager.StopLPAR(lparid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(Equal(""))
			})
		})

		Context("When stop lpar and no URI responded", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusAccepted, bytesResponseWithoutURI, errors.New("error"))
				rets, err := manager.StopLPAR(lparid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(Equal(""))
			})
		})
	})

	Describe("MountIsoImage", func() {

		var (
			bytes     []byte
			imageFile string
			insFile   string
		)

		BeforeEach(func() {
			imageFile = "imageFileName"
			insFile = "insFileName"
			bytes = []byte("Here is a bytes array represents an image file contents....")
		})

		Context("When mount iso image and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.UploadRequestReturns(http.StatusNoContent, nil, nil)
				err := manager.MountIsoImage(lparid, bytes, imageFile, insFile)

				Expect(err).To(BeNil())
			})
		})

		Context("When mount iso image and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.UploadRequestReturns(http.StatusBadRequest, nil, errors.New("error"))
				err := manager.MountIsoImage(lparid, bytes, imageFile, insFile)

				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UnmountIsoImage", func() {

		Context("When unmount iso image and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				err := manager.UnmountIsoImage(lparid)

				Expect(err).To(BeNil())
			})
		})

		Context("When unmount iso image and ExecuteRequest error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusBadRequest, nil, errors.New("error"))
				err := manager.UnmountIsoImage(lparid)

				Expect(err).ToNot(BeNil())
			})
		})
	})
})
