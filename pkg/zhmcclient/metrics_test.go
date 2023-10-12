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
	"net/http"
	"net/url"
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
	"github.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient/fakes"
)

var _ = Describe("Metrics", func() {
	var (
		manager              *MetricsManager
		fakeClient           *fakes.ClientAPI
		lparid               string
		url                  *url.URL
		hmcErr               *HmcError

	)

	BeforeEach(func() {
		fakeClient = &fakes.ClientAPI{}
		manager = NewMetricsManager(fakeClient)

		url, _ = url.Parse("https://127.0.0.1:443")
		lparid = "lparid"

		hmcErr = &HmcError{
			Reason:  int(ERR_CODE_HMC_BAD_REQUEST),
			Message: "error message",
		}
	})


	Describe("GetLiveEnergyDetailsforLPAR", func() {
		var (
			bytes      []byte
			metricsContext *MetricsContextDef
		)

		BeforeEach(func() {
			jsonString := `{
				"metrics-context-uri": "/example/metrics",
				"metric-group-infos": [
				  {
					"name": "power-consumption-watts",
					"metrics": ["metric1", "metric2", "metric3"]
				  },
				  {
					"name": "group2",
					"metrics": ["metric4", "metric5"]
				  }
				],
				"metric-group-infos-by-name": {
				  "group1": {
					"name": "group1",
					"metrics": ["metric1", "metric2", "metric3"]
				  },
				  "group2": {
					"name": "group2",
					"metrics": ["metric4", "metric5"]
				  }
				}
			  }
			`		
			bytes = []byte(jsonString)
			json.Unmarshal(bytes, &metricsContext)
		})

		Context("When GetLiveEnergyDetailsforLPAR returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)				
				fakeClient.GetMetricsContextReturns(metricsContext)

				rets, _, err := manager.GetLiveEnergyDetailsforLPAR(lparid)

				Expect(err).To(BeNil())
				Expect(rets).To(Equal(uint64(0)))
			})
		})

		Context("When GetEnergyDetailsforLPAR returns error due to hmcErr", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, hmcErr)
				fakeClient.GetMetricsContextReturns(metricsContext)

				rets, _, err := manager.GetLiveEnergyDetailsforLPAR(lparid)

				Expect(*err).To(Equal(*hmcErr))
				Expect(rets).To(Equal(uint64(0)))

			})
		})

	})

})
