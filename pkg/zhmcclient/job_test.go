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

		Context("When QueryJob returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, _, err := manager.QueryJob(jobid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets.URI).To(Equal(job.URI))
				Expect(rets.Status).To(Equal(job.Status))
			})
		})

		Context("When QueryJob returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, bytes, hmcErr)
				rets, _, err := manager.QueryJob(jobid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When QueryJob returns error due to unmarshalErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), nil)
				rets, _, err := manager.QueryJob(jobid)

				Expect(*err).To(Equal(*unmarshalErr))
				Expect(rets).To(BeNil())
			})
		})

		Context("When QueryJob returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, bytes, nil)
				rets, _, err := manager.QueryJob(jobid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("DeleteJob", func() {

		Context("When DeleteJob returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				_, err := manager.DeleteJob(jobid)

				Expect(err).To(BeNil())
			})
		})

		Context("When DeleteJob returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, hmcErr)
				_, err := manager.DeleteJob(jobid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When DeleteJob returns incorrect status", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, nil)
				_, err := manager.DeleteJob(jobid)

				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("CancelJob", func() {

		Context("When CancelJob returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusNoContent, nil, nil)
				_, err := manager.CancelJob(jobid)

				Expect(err).To(BeNil())
			})
		})

		Context("When CancelJob returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, hmcErr)
				_, err := manager.CancelJob(jobid)

				Expect(*err).To(Equal(*hmcErr))
			})
		})

		Context("When CancelJob returns incorrect status", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusInternalServerError, nil, nil)
				_, err := manager.CancelJob(jobid)

				Expect(err).ToNot(BeNil())
			})
		})
	})
})
