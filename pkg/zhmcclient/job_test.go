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

var _ = Describe("JOB", func() {
	var (
		manager    *JobManager
		fakeClient *fakes.ClientAPI

		job   *Job
		url   *url.URL
		bytes []byte
		jobid string

		hmcErr, unmarshalErr *HmcError
	)

	BeforeEach(func() {
		fakeClient = &fakes.ClientAPI{}
		manager = NewJobManager(fakeClient)

		jobResults := &JobResults{
			Message: "aaa",
		}

		job = &Job{
			URI:           "uri",
			Status:        JOB_STATUS_RUNNING,
			JobStatusCode: 200,
			JobReasonCode: 200,
			JobResults:    *jobResults,
		}

		url, _ = url.Parse("https://127.0.0.1:443")
		bytes, _ = json.Marshal(job)
		jobid = "jobid"

		hmcErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
			Message: "error message",
		}

		unmarshalErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_UNMARSHAL_FAIL),
			Message: "invalid character 'i' looking for beginning of value",
		}
	})

	Describe("QueryJob", func() {

		Context("When query job and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
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
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, bytes, hmcErr)
				rets, err := manager.QueryJob(jobid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When query job and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, err := manager.QueryJob(jobid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When query job and returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
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
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				err := manager.DeleteJob(jobid)

				Expect(err).To(BeNil())
			})
		})

		Context("When delete job  and returns error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, hmcErr)
				err := manager.DeleteJob(jobid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When delete job  and returns incorrect status", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, nil)
				err := manager.DeleteJob(jobid)

				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("CancelJob", func() {

		Context("When cancel job and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				err := manager.CancelJob(jobid)

				Expect(err).To(BeNil())
			})
		})

		Context("When cancel job  and returns error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, hmcErr)
				err := manager.CancelJob(jobid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When cancel job  and returns incorrect status", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, nil)
				err := manager.CancelJob(jobid)

				Expect(err).ToNot(BeNil())
			})
		})
	})
})
