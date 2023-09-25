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
	"sync"

	"github.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

type VirtualSwitchAPI struct {
	GetVirtualSwitchPropertiesStub        func(string) (*zhmcclient.VirtualSwitchProperties, int, *zhmcclient.HmcError)
	getVirtualSwitchPropertiesMutex       sync.RWMutex
	getVirtualSwitchPropertiesArgsForCall []struct {
		arg1 string
	}
	getVirtualSwitchPropertiesReturns struct {
		result1 *zhmcclient.VirtualSwitchProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	getVirtualSwitchPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.VirtualSwitchProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	ListVirtualSwitchesStub        func(string, map[string]string) ([]zhmcclient.VirtualSwitch, int, *zhmcclient.HmcError)
	listVirtualSwitchesMutex       sync.RWMutex
	listVirtualSwitchesArgsForCall []struct {
		arg1 string
		arg2 map[string]string
	}
	listVirtualSwitchesReturns struct {
		result1 []zhmcclient.VirtualSwitch
		result2 int
		result3 *zhmcclient.HmcError
	}
	listVirtualSwitchesReturnsOnCall map[int]struct {
		result1 []zhmcclient.VirtualSwitch
		result2 int
		result3 *zhmcclient.HmcError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *VirtualSwitchAPI) GetVirtualSwitchProperties(arg1 string) (*zhmcclient.VirtualSwitchProperties, int, *zhmcclient.HmcError) {
	fake.getVirtualSwitchPropertiesMutex.Lock()
	ret, specificReturn := fake.getVirtualSwitchPropertiesReturnsOnCall[len(fake.getVirtualSwitchPropertiesArgsForCall)]
	fake.getVirtualSwitchPropertiesArgsForCall = append(fake.getVirtualSwitchPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetVirtualSwitchPropertiesStub
	fakeReturns := fake.getVirtualSwitchPropertiesReturns
	fake.recordInvocation("GetVirtualSwitchProperties", []interface{}{arg1})
	fake.getVirtualSwitchPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *VirtualSwitchAPI) GetVirtualSwitchPropertiesCallCount() int {
	fake.getVirtualSwitchPropertiesMutex.RLock()
	defer fake.getVirtualSwitchPropertiesMutex.RUnlock()
	return len(fake.getVirtualSwitchPropertiesArgsForCall)
}

func (fake *VirtualSwitchAPI) GetVirtualSwitchPropertiesCalls(stub func(string) (*zhmcclient.VirtualSwitchProperties, int, *zhmcclient.HmcError)) {
	fake.getVirtualSwitchPropertiesMutex.Lock()
	defer fake.getVirtualSwitchPropertiesMutex.Unlock()
	fake.GetVirtualSwitchPropertiesStub = stub
}

func (fake *VirtualSwitchAPI) GetVirtualSwitchPropertiesArgsForCall(i int) string {
	fake.getVirtualSwitchPropertiesMutex.RLock()
	defer fake.getVirtualSwitchPropertiesMutex.RUnlock()
	argsForCall := fake.getVirtualSwitchPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *VirtualSwitchAPI) GetVirtualSwitchPropertiesReturns(result1 *zhmcclient.VirtualSwitchProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getVirtualSwitchPropertiesMutex.Lock()
	defer fake.getVirtualSwitchPropertiesMutex.Unlock()
	fake.GetVirtualSwitchPropertiesStub = nil
	fake.getVirtualSwitchPropertiesReturns = struct {
		result1 *zhmcclient.VirtualSwitchProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *VirtualSwitchAPI) GetVirtualSwitchPropertiesReturnsOnCall(i int, result1 *zhmcclient.VirtualSwitchProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getVirtualSwitchPropertiesMutex.Lock()
	defer fake.getVirtualSwitchPropertiesMutex.Unlock()
	fake.GetVirtualSwitchPropertiesStub = nil
	if fake.getVirtualSwitchPropertiesReturnsOnCall == nil {
		fake.getVirtualSwitchPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.VirtualSwitchProperties
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getVirtualSwitchPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.VirtualSwitchProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *VirtualSwitchAPI) ListVirtualSwitches(arg1 string, arg2 map[string]string) ([]zhmcclient.VirtualSwitch, int, *zhmcclient.HmcError) {
	fake.listVirtualSwitchesMutex.Lock()
	ret, specificReturn := fake.listVirtualSwitchesReturnsOnCall[len(fake.listVirtualSwitchesArgsForCall)]
	fake.listVirtualSwitchesArgsForCall = append(fake.listVirtualSwitchesArgsForCall, struct {
		arg1 string
		arg2 map[string]string
	}{arg1, arg2})
	stub := fake.ListVirtualSwitchesStub
	fakeReturns := fake.listVirtualSwitchesReturns
	fake.recordInvocation("ListVirtualSwitches", []interface{}{arg1, arg2})
	fake.listVirtualSwitchesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *VirtualSwitchAPI) ListVirtualSwitchesCallCount() int {
	fake.listVirtualSwitchesMutex.RLock()
	defer fake.listVirtualSwitchesMutex.RUnlock()
	return len(fake.listVirtualSwitchesArgsForCall)
}

func (fake *VirtualSwitchAPI) ListVirtualSwitchesCalls(stub func(string, map[string]string) ([]zhmcclient.VirtualSwitch, int, *zhmcclient.HmcError)) {
	fake.listVirtualSwitchesMutex.Lock()
	defer fake.listVirtualSwitchesMutex.Unlock()
	fake.ListVirtualSwitchesStub = stub
}

func (fake *VirtualSwitchAPI) ListVirtualSwitchesArgsForCall(i int) (string, map[string]string) {
	fake.listVirtualSwitchesMutex.RLock()
	defer fake.listVirtualSwitchesMutex.RUnlock()
	argsForCall := fake.listVirtualSwitchesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *VirtualSwitchAPI) ListVirtualSwitchesReturns(result1 []zhmcclient.VirtualSwitch, result2 int, result3 *zhmcclient.HmcError) {
	fake.listVirtualSwitchesMutex.Lock()
	defer fake.listVirtualSwitchesMutex.Unlock()
	fake.ListVirtualSwitchesStub = nil
	fake.listVirtualSwitchesReturns = struct {
		result1 []zhmcclient.VirtualSwitch
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *VirtualSwitchAPI) ListVirtualSwitchesReturnsOnCall(i int, result1 []zhmcclient.VirtualSwitch, result2 int, result3 *zhmcclient.HmcError) {
	fake.listVirtualSwitchesMutex.Lock()
	defer fake.listVirtualSwitchesMutex.Unlock()
	fake.ListVirtualSwitchesStub = nil
	if fake.listVirtualSwitchesReturnsOnCall == nil {
		fake.listVirtualSwitchesReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.VirtualSwitch
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listVirtualSwitchesReturnsOnCall[i] = struct {
		result1 []zhmcclient.VirtualSwitch
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *VirtualSwitchAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getVirtualSwitchPropertiesMutex.RLock()
	defer fake.getVirtualSwitchPropertiesMutex.RUnlock()
	fake.listVirtualSwitchesMutex.RLock()
	defer fake.listVirtualSwitchesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *VirtualSwitchAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.VirtualSwitchAPI = new(VirtualSwitchAPI)
