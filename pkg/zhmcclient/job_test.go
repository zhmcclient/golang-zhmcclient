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

type TempNestedObj struct {
	aaa string
	bbb string
}

var _ = Describe("JOB", func() {
	var (
		manager    *JobManager
		fakeClient *fakes.ClientAPI

		job         *Job
		url         *url.URL
		nestedbytes []byte
		bytes       []byte
		jobid       string
	)

	BeforeEach(func() {
		fakeClient = &fakes.ClientAPI{}
		manager = NewJobManager(fakeClient)

		nested := &TempNestedObj{
			aaa: "aaa",
			bbb: "bbb",
		}

		nestedbytes, _ = json.Marshal(nested)

		job = &Job{
			URI:           "uri",
			Status:        JOB_STATUS_RUNNING,
			JobStatusCode: 200,
			JobReasonCode: 200,
			JobResults:    nestedbytes,
		}

		url, _ = url.Parse("https://127.0.0.1:443")
		bytes, _ = json.Marshal(job)
		jobid = "jobid"
	})

	Describe("QueryJob", func() {

		Context("When query job and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, err := manager.QueryJob(jobid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets.URI).To(Equal(job.URI))
				Expect(rets.Status).To(Equal(job.Status))
			})
		})

		Context("When query job  and returns error", func() {
			It("check the error happened", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, bytes, errors.New("error"))
				rets, err := manager.QueryJob(jobid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When query job and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), errors.New("error"))
				rets, err := manager.QueryJob(jobid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When query job and returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, bytes, nil)
				rets, err := manager.QueryJob(jobid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("DeleteJob", func() {

		Context("When delete job and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				err := manager.DeleteJob(jobid)

				Expect(err).To(BeNil())
			})
		})

		Context("When delete job  and returns error", func() {
			It("check the error happened", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, errors.New("error"))
				err := manager.DeleteJob(jobid)

				Expect(err).ToNot(BeNil())
			})
		})

		Context("When delete job  and returns incorrect status", func() {
			It("check the error happened", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, nil)
				err := manager.DeleteJob(jobid)

				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("CancelJob", func() {

		Context("When cancel job and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				err := manager.CancelJob(jobid)

				Expect(err).To(BeNil())
			})
		})

		Context("When cancel job  and returns error", func() {
			It("check the error happened", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, errors.New("error"))
				err := manager.CancelJob(jobid)

				Expect(err).ToNot(BeNil())
			})
		})

		Context("When cancel job  and returns incorrect status", func() {
			It("check the error happened", func() {
				fakeClient.GetEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, nil)
				err := manager.CancelJob(jobid)

				Expect(err).ToNot(BeNil())
			})
		})
	})
})
