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

package fakes

import (
	"io"
	"net/url"
	"sync"

	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

type ClientAPI struct {
	CloneEndpointURLStub        func() *url.URL
	cloneEndpointURLMutex       sync.RWMutex
	cloneEndpointURLArgsForCall []struct {
	}
	cloneEndpointURLReturns struct {
		result1 *url.URL
	}
	cloneEndpointURLReturnsOnCall map[int]struct {
		result1 *url.URL
	}
	ExecuteRequestStub        func(string, *url.URL, interface{}, string) (int, []byte, *zhmcclient.HmcError)
	executeRequestMutex       sync.RWMutex
	executeRequestArgsForCall []struct {
		arg1 string
		arg2 *url.URL
		arg3 interface{}
		arg4 string
	}
	executeRequestReturns struct {
		result1 int
		result2 []byte
		result3 *zhmcclient.HmcError
	}
	executeRequestReturnsOnCall map[int]struct {
		result1 int
		result2 []byte
		result3 *zhmcclient.HmcError
	}
	IsLogonStub        func(bool) bool
	isLogonMutex       sync.RWMutex
	isLogonArgsForCall []struct {
		arg1 bool
	}
	isLogonReturns struct {
		result1 bool
	}
	isLogonReturnsOnCall map[int]struct {
		result1 bool
	}
	LogoffStub        func() *zhmcclient.HmcError
	logoffMutex       sync.RWMutex
	logoffArgsForCall []struct {
	}
	logoffReturns struct {
		result1 *zhmcclient.HmcError
	}
	logoffReturnsOnCall map[int]struct {
		result1 *zhmcclient.HmcError
	}
	LogoffConsoleStub        func(string) *zhmcclient.HmcError
	logoffConsoleMutex       sync.RWMutex
	logoffConsoleArgsForCall []struct {
		arg1 string
	}
	logoffConsoleReturns struct {
		result1 *zhmcclient.HmcError
	}
	logoffConsoleReturnsOnCall map[int]struct {
		result1 *zhmcclient.HmcError
	}
	LogonStub        func() *zhmcclient.HmcError
	logonMutex       sync.RWMutex
	logonArgsForCall []struct {
	}
	logonReturns struct {
		result1 *zhmcclient.HmcError
	}
	logonReturnsOnCall map[int]struct {
		result1 *zhmcclient.HmcError
	}
	LogonConsoleStub        func() (string, int, *zhmcclient.HmcError)
	logonConsoleMutex       sync.RWMutex
	logonConsoleArgsForCall []struct {
	}
	logonConsoleReturns struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	logonConsoleReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	SetSkipCertVerifyStub        func(bool)
	setSkipCertVerifyMutex       sync.RWMutex
	setSkipCertVerifyArgsForCall []struct {
		arg1 bool
	}
	TraceOffStub        func()
	traceOffMutex       sync.RWMutex
	traceOffArgsForCall []struct {
	}
	TraceOnStub        func(io.Writer)
	traceOnMutex       sync.RWMutex
	traceOnArgsForCall []struct {
		arg1 io.Writer
	}
	UploadRequestStub        func(string, *url.URL, []byte) (int, []byte, *zhmcclient.HmcError)
	uploadRequestMutex       sync.RWMutex
	uploadRequestArgsForCall []struct {
		arg1 string
		arg2 *url.URL
		arg3 []byte
	}
	uploadRequestReturns struct {
		result1 int
		result2 []byte
		result3 *zhmcclient.HmcError
	}
	uploadRequestReturnsOnCall map[int]struct {
		result1 int
		result2 []byte
		result3 *zhmcclient.HmcError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ClientAPI) CloneEndpointURL() *url.URL {
	fake.cloneEndpointURLMutex.Lock()
	ret, specificReturn := fake.cloneEndpointURLReturnsOnCall[len(fake.cloneEndpointURLArgsForCall)]
	fake.cloneEndpointURLArgsForCall = append(fake.cloneEndpointURLArgsForCall, struct {
	}{})
	stub := fake.CloneEndpointURLStub
	fakeReturns := fake.cloneEndpointURLReturns
	fake.recordInvocation("CloneEndpointURL", []interface{}{})
	fake.cloneEndpointURLMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ClientAPI) CloneEndpointURLCallCount() int {
	fake.cloneEndpointURLMutex.RLock()
	defer fake.cloneEndpointURLMutex.RUnlock()
	return len(fake.cloneEndpointURLArgsForCall)
}

func (fake *ClientAPI) CloneEndpointURLCalls(stub func() *url.URL) {
	fake.cloneEndpointURLMutex.Lock()
	defer fake.cloneEndpointURLMutex.Unlock()
	fake.CloneEndpointURLStub = stub
}

func (fake *ClientAPI) CloneEndpointURLReturns(result1 *url.URL) {
	fake.cloneEndpointURLMutex.Lock()
	defer fake.cloneEndpointURLMutex.Unlock()
	fake.CloneEndpointURLStub = nil
	fake.cloneEndpointURLReturns = struct {
		result1 *url.URL
	}{result1}
}

func (fake *ClientAPI) CloneEndpointURLReturnsOnCall(i int, result1 *url.URL) {
	fake.cloneEndpointURLMutex.Lock()
	defer fake.cloneEndpointURLMutex.Unlock()
	fake.CloneEndpointURLStub = nil
	if fake.cloneEndpointURLReturnsOnCall == nil {
		fake.cloneEndpointURLReturnsOnCall = make(map[int]struct {
			result1 *url.URL
		})
	}
	fake.cloneEndpointURLReturnsOnCall[i] = struct {
		result1 *url.URL
	}{result1}
}

func (fake *ClientAPI) ExecuteRequest(arg1 string, arg2 *url.URL, arg3 interface{}, arg4 string) (int, []byte, *zhmcclient.HmcError) {
	fake.executeRequestMutex.Lock()
	ret, specificReturn := fake.executeRequestReturnsOnCall[len(fake.executeRequestArgsForCall)]
	fake.executeRequestArgsForCall = append(fake.executeRequestArgsForCall, struct {
		arg1 string
		arg2 *url.URL
		arg3 interface{}
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.ExecuteRequestStub
	fakeReturns := fake.executeRequestReturns
	fake.recordInvocation("ExecuteRequest", []interface{}{arg1, arg2, arg3, arg4})
	fake.executeRequestMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ClientAPI) ExecuteRequestCallCount() int {
	fake.executeRequestMutex.RLock()
	defer fake.executeRequestMutex.RUnlock()
	return len(fake.executeRequestArgsForCall)
}

func (fake *ClientAPI) ExecuteRequestCalls(stub func(string, *url.URL, interface{}, string) (int, []byte, *zhmcclient.HmcError)) {
	fake.executeRequestMutex.Lock()
	defer fake.executeRequestMutex.Unlock()
	fake.ExecuteRequestStub = stub
}

func (fake *ClientAPI) ExecuteRequestArgsForCall(i int) (string, *url.URL, interface{}, string) {
	fake.executeRequestMutex.RLock()
	defer fake.executeRequestMutex.RUnlock()
	argsForCall := fake.executeRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *ClientAPI) ExecuteRequestReturns(result1 int, result2 []byte, result3 *zhmcclient.HmcError) {
	fake.executeRequestMutex.Lock()
	defer fake.executeRequestMutex.Unlock()
	fake.ExecuteRequestStub = nil
	fake.executeRequestReturns = struct {
		result1 int
		result2 []byte
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ClientAPI) ExecuteRequestReturnsOnCall(i int, result1 int, result2 []byte, result3 *zhmcclient.HmcError) {
	fake.executeRequestMutex.Lock()
	defer fake.executeRequestMutex.Unlock()
	fake.ExecuteRequestStub = nil
	if fake.executeRequestReturnsOnCall == nil {
		fake.executeRequestReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []byte
			result3 *zhmcclient.HmcError
		})
	}
	fake.executeRequestReturnsOnCall[i] = struct {
		result1 int
		result2 []byte
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ClientAPI) IsLogon(arg1 bool) bool {
	fake.isLogonMutex.Lock()
	ret, specificReturn := fake.isLogonReturnsOnCall[len(fake.isLogonArgsForCall)]
	fake.isLogonArgsForCall = append(fake.isLogonArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.IsLogonStub
	fakeReturns := fake.isLogonReturns
	fake.recordInvocation("IsLogon", []interface{}{arg1})
	fake.isLogonMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ClientAPI) IsLogonCallCount() int {
	fake.isLogonMutex.RLock()
	defer fake.isLogonMutex.RUnlock()
	return len(fake.isLogonArgsForCall)
}

func (fake *ClientAPI) IsLogonCalls(stub func(bool) bool) {
	fake.isLogonMutex.Lock()
	defer fake.isLogonMutex.Unlock()
	fake.IsLogonStub = stub
}

func (fake *ClientAPI) IsLogonArgsForCall(i int) bool {
	fake.isLogonMutex.RLock()
	defer fake.isLogonMutex.RUnlock()
	argsForCall := fake.isLogonArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ClientAPI) IsLogonReturns(result1 bool) {
	fake.isLogonMutex.Lock()
	defer fake.isLogonMutex.Unlock()
	fake.IsLogonStub = nil
	fake.isLogonReturns = struct {
		result1 bool
	}{result1}
}

func (fake *ClientAPI) IsLogonReturnsOnCall(i int, result1 bool) {
	fake.isLogonMutex.Lock()
	defer fake.isLogonMutex.Unlock()
	fake.IsLogonStub = nil
	if fake.isLogonReturnsOnCall == nil {
		fake.isLogonReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isLogonReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *ClientAPI) Logoff() *zhmcclient.HmcError {
	fake.logoffMutex.Lock()
	ret, specificReturn := fake.logoffReturnsOnCall[len(fake.logoffArgsForCall)]
	fake.logoffArgsForCall = append(fake.logoffArgsForCall, struct {
	}{})
	stub := fake.LogoffStub
	fakeReturns := fake.logoffReturns
	fake.recordInvocation("Logoff", []interface{}{})
	fake.logoffMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ClientAPI) LogoffCallCount() int {
	fake.logoffMutex.RLock()
	defer fake.logoffMutex.RUnlock()
	return len(fake.logoffArgsForCall)
}

func (fake *ClientAPI) LogoffCalls(stub func() *zhmcclient.HmcError) {
	fake.logoffMutex.Lock()
	defer fake.logoffMutex.Unlock()
	fake.LogoffStub = stub
}

func (fake *ClientAPI) LogoffReturns(result1 *zhmcclient.HmcError) {
	fake.logoffMutex.Lock()
	defer fake.logoffMutex.Unlock()
	fake.LogoffStub = nil
	fake.logoffReturns = struct {
		result1 *zhmcclient.HmcError
	}{result1}
}

func (fake *ClientAPI) LogoffReturnsOnCall(i int, result1 *zhmcclient.HmcError) {
	fake.logoffMutex.Lock()
	defer fake.logoffMutex.Unlock()
	fake.LogoffStub = nil
	if fake.logoffReturnsOnCall == nil {
		fake.logoffReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.HmcError
		})
	}
	fake.logoffReturnsOnCall[i] = struct {
		result1 *zhmcclient.HmcError
	}{result1}
}

func (fake *ClientAPI) LogoffConsole(arg1 string) *zhmcclient.HmcError {
	fake.logoffConsoleMutex.Lock()
	ret, specificReturn := fake.logoffConsoleReturnsOnCall[len(fake.logoffConsoleArgsForCall)]
	fake.logoffConsoleArgsForCall = append(fake.logoffConsoleArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LogoffConsoleStub
	fakeReturns := fake.logoffConsoleReturns
	fake.recordInvocation("LogoffConsole", []interface{}{arg1})
	fake.logoffConsoleMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ClientAPI) LogoffConsoleCallCount() int {
	fake.logoffConsoleMutex.RLock()
	defer fake.logoffConsoleMutex.RUnlock()
	return len(fake.logoffConsoleArgsForCall)
}

func (fake *ClientAPI) LogoffConsoleCalls(stub func(string) *zhmcclient.HmcError) {
	fake.logoffConsoleMutex.Lock()
	defer fake.logoffConsoleMutex.Unlock()
	fake.LogoffConsoleStub = stub
}

func (fake *ClientAPI) LogoffConsoleArgsForCall(i int) string {
	fake.logoffConsoleMutex.RLock()
	defer fake.logoffConsoleMutex.RUnlock()
	argsForCall := fake.logoffConsoleArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ClientAPI) LogoffConsoleReturns(result1 *zhmcclient.HmcError) {
	fake.logoffConsoleMutex.Lock()
	defer fake.logoffConsoleMutex.Unlock()
	fake.LogoffConsoleStub = nil
	fake.logoffConsoleReturns = struct {
		result1 *zhmcclient.HmcError
	}{result1}
}

func (fake *ClientAPI) LogoffConsoleReturnsOnCall(i int, result1 *zhmcclient.HmcError) {
	fake.logoffConsoleMutex.Lock()
	defer fake.logoffConsoleMutex.Unlock()
	fake.LogoffConsoleStub = nil
	if fake.logoffConsoleReturnsOnCall == nil {
		fake.logoffConsoleReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.HmcError
		})
	}
	fake.logoffConsoleReturnsOnCall[i] = struct {
		result1 *zhmcclient.HmcError
	}{result1}
}

func (fake *ClientAPI) Logon() *zhmcclient.HmcError {
	fake.logonMutex.Lock()
	ret, specificReturn := fake.logonReturnsOnCall[len(fake.logonArgsForCall)]
	fake.logonArgsForCall = append(fake.logonArgsForCall, struct {
	}{})
	stub := fake.LogonStub
	fakeReturns := fake.logonReturns
	fake.recordInvocation("Logon", []interface{}{})
	fake.logonMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ClientAPI) LogonCallCount() int {
	fake.logonMutex.RLock()
	defer fake.logonMutex.RUnlock()
	return len(fake.logonArgsForCall)
}

func (fake *ClientAPI) LogonCalls(stub func() *zhmcclient.HmcError) {
	fake.logonMutex.Lock()
	defer fake.logonMutex.Unlock()
	fake.LogonStub = stub
}

func (fake *ClientAPI) LogonReturns(result1 *zhmcclient.HmcError) {
	fake.logonMutex.Lock()
	defer fake.logonMutex.Unlock()
	fake.LogonStub = nil
	fake.logonReturns = struct {
		result1 *zhmcclient.HmcError
	}{result1}
}

func (fake *ClientAPI) LogonReturnsOnCall(i int, result1 *zhmcclient.HmcError) {
	fake.logonMutex.Lock()
	defer fake.logonMutex.Unlock()
	fake.LogonStub = nil
	if fake.logonReturnsOnCall == nil {
		fake.logonReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.HmcError
		})
	}
	fake.logonReturnsOnCall[i] = struct {
		result1 *zhmcclient.HmcError
	}{result1}
}

func (fake *ClientAPI) LogonConsole() (string, int, *zhmcclient.HmcError) {
	fake.logonConsoleMutex.Lock()
	ret, specificReturn := fake.logonConsoleReturnsOnCall[len(fake.logonConsoleArgsForCall)]
	fake.logonConsoleArgsForCall = append(fake.logonConsoleArgsForCall, struct {
	}{})
	stub := fake.LogonConsoleStub
	fakeReturns := fake.logonConsoleReturns
	fake.recordInvocation("LogonConsole", []interface{}{})
	fake.logonConsoleMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ClientAPI) LogonConsoleCallCount() int {
	fake.logonConsoleMutex.RLock()
	defer fake.logonConsoleMutex.RUnlock()
	return len(fake.logonConsoleArgsForCall)
}

func (fake *ClientAPI) LogonConsoleCalls(stub func() (string, int, *zhmcclient.HmcError)) {
	fake.logonConsoleMutex.Lock()
	defer fake.logonConsoleMutex.Unlock()
	fake.LogonConsoleStub = stub
}

func (fake *ClientAPI) LogonConsoleReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.logonConsoleMutex.Lock()
	defer fake.logonConsoleMutex.Unlock()
	fake.LogonConsoleStub = nil
	fake.logonConsoleReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ClientAPI) LogonConsoleReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.logonConsoleMutex.Lock()
	defer fake.logonConsoleMutex.Unlock()
	fake.LogonConsoleStub = nil
	if fake.logonConsoleReturnsOnCall == nil {
		fake.logonConsoleReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.logonConsoleReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ClientAPI) SetSkipCertVerify(arg1 bool) {
	fake.setSkipCertVerifyMutex.Lock()
	fake.setSkipCertVerifyArgsForCall = append(fake.setSkipCertVerifyArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.SetSkipCertVerifyStub
	fake.recordInvocation("SetSkipCertVerify", []interface{}{arg1})
	fake.setSkipCertVerifyMutex.Unlock()
	if stub != nil {
		fake.SetSkipCertVerifyStub(arg1)
	}
}

func (fake *ClientAPI) SetSkipCertVerifyCallCount() int {
	fake.setSkipCertVerifyMutex.RLock()
	defer fake.setSkipCertVerifyMutex.RUnlock()
	return len(fake.setSkipCertVerifyArgsForCall)
}

func (fake *ClientAPI) SetSkipCertVerifyCalls(stub func(bool)) {
	fake.setSkipCertVerifyMutex.Lock()
	defer fake.setSkipCertVerifyMutex.Unlock()
	fake.SetSkipCertVerifyStub = stub
}

func (fake *ClientAPI) SetSkipCertVerifyArgsForCall(i int) bool {
	fake.setSkipCertVerifyMutex.RLock()
	defer fake.setSkipCertVerifyMutex.RUnlock()
	argsForCall := fake.setSkipCertVerifyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ClientAPI) TraceOff() {
	fake.traceOffMutex.Lock()
	fake.traceOffArgsForCall = append(fake.traceOffArgsForCall, struct {
	}{})
	stub := fake.TraceOffStub
	fake.recordInvocation("TraceOff", []interface{}{})
	fake.traceOffMutex.Unlock()
	if stub != nil {
		fake.TraceOffStub()
	}
}

func (fake *ClientAPI) TraceOffCallCount() int {
	fake.traceOffMutex.RLock()
	defer fake.traceOffMutex.RUnlock()
	return len(fake.traceOffArgsForCall)
}

func (fake *ClientAPI) TraceOffCalls(stub func()) {
	fake.traceOffMutex.Lock()
	defer fake.traceOffMutex.Unlock()
	fake.TraceOffStub = stub
}

func (fake *ClientAPI) TraceOn(arg1 io.Writer) {
	fake.traceOnMutex.Lock()
	fake.traceOnArgsForCall = append(fake.traceOnArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	stub := fake.TraceOnStub
	fake.recordInvocation("TraceOn", []interface{}{arg1})
	fake.traceOnMutex.Unlock()
	if stub != nil {
		fake.TraceOnStub(arg1)
	}
}

func (fake *ClientAPI) TraceOnCallCount() int {
	fake.traceOnMutex.RLock()
	defer fake.traceOnMutex.RUnlock()
	return len(fake.traceOnArgsForCall)
}

func (fake *ClientAPI) TraceOnCalls(stub func(io.Writer)) {
	fake.traceOnMutex.Lock()
	defer fake.traceOnMutex.Unlock()
	fake.TraceOnStub = stub
}

func (fake *ClientAPI) TraceOnArgsForCall(i int) io.Writer {
	fake.traceOnMutex.RLock()
	defer fake.traceOnMutex.RUnlock()
	argsForCall := fake.traceOnArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ClientAPI) UploadRequest(arg1 string, arg2 *url.URL, arg3 []byte) (int, []byte, *zhmcclient.HmcError) {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.uploadRequestMutex.Lock()
	ret, specificReturn := fake.uploadRequestReturnsOnCall[len(fake.uploadRequestArgsForCall)]
	fake.uploadRequestArgsForCall = append(fake.uploadRequestArgsForCall, struct {
		arg1 string
		arg2 *url.URL
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	stub := fake.UploadRequestStub
	fakeReturns := fake.uploadRequestReturns
	fake.recordInvocation("UploadRequest", []interface{}{arg1, arg2, arg3Copy})
	fake.uploadRequestMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ClientAPI) UploadRequestCallCount() int {
	fake.uploadRequestMutex.RLock()
	defer fake.uploadRequestMutex.RUnlock()
	return len(fake.uploadRequestArgsForCall)
}

func (fake *ClientAPI) UploadRequestCalls(stub func(string, *url.URL, []byte) (int, []byte, *zhmcclient.HmcError)) {
	fake.uploadRequestMutex.Lock()
	defer fake.uploadRequestMutex.Unlock()
	fake.UploadRequestStub = stub
}

func (fake *ClientAPI) UploadRequestArgsForCall(i int) (string, *url.URL, []byte) {
	fake.uploadRequestMutex.RLock()
	defer fake.uploadRequestMutex.RUnlock()
	argsForCall := fake.uploadRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ClientAPI) UploadRequestReturns(result1 int, result2 []byte, result3 *zhmcclient.HmcError) {
	fake.uploadRequestMutex.Lock()
	defer fake.uploadRequestMutex.Unlock()
	fake.UploadRequestStub = nil
	fake.uploadRequestReturns = struct {
		result1 int
		result2 []byte
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ClientAPI) UploadRequestReturnsOnCall(i int, result1 int, result2 []byte, result3 *zhmcclient.HmcError) {
	fake.uploadRequestMutex.Lock()
	defer fake.uploadRequestMutex.Unlock()
	fake.UploadRequestStub = nil
	if fake.uploadRequestReturnsOnCall == nil {
		fake.uploadRequestReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []byte
			result3 *zhmcclient.HmcError
		})
	}
	fake.uploadRequestReturnsOnCall[i] = struct {
		result1 int
		result2 []byte
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ClientAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloneEndpointURLMutex.RLock()
	defer fake.cloneEndpointURLMutex.RUnlock()
	fake.executeRequestMutex.RLock()
	defer fake.executeRequestMutex.RUnlock()
	fake.isLogonMutex.RLock()
	defer fake.isLogonMutex.RUnlock()
	fake.logoffMutex.RLock()
	defer fake.logoffMutex.RUnlock()
	fake.logoffConsoleMutex.RLock()
	defer fake.logoffConsoleMutex.RUnlock()
	fake.logonMutex.RLock()
	defer fake.logonMutex.RUnlock()
	fake.logonConsoleMutex.RLock()
	defer fake.logonConsoleMutex.RUnlock()
	fake.setSkipCertVerifyMutex.RLock()
	defer fake.setSkipCertVerifyMutex.RUnlock()
	fake.traceOffMutex.RLock()
	defer fake.traceOffMutex.RUnlock()
	fake.traceOnMutex.RLock()
	defer fake.traceOnMutex.RUnlock()
	fake.uploadRequestMutex.RLock()
	defer fake.uploadRequestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ClientAPI) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ zhmcclient.ClientAPI = new(ClientAPI)
