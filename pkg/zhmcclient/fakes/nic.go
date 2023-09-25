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

type NicAPI struct {
	CreateNicStub        func(string, *zhmcclient.NIC) (string, int, *zhmcclient.HmcError)
	createNicMutex       sync.RWMutex
	createNicArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.NIC
	}
	createNicReturns struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	createNicReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	DeleteNicStub        func(string) (int, *zhmcclient.HmcError)
	deleteNicMutex       sync.RWMutex
	deleteNicArgsForCall []struct {
		arg1 string
	}
	deleteNicReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	deleteNicReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	GetNicPropertiesStub        func(string) (*zhmcclient.NIC, int, *zhmcclient.HmcError)
	getNicPropertiesMutex       sync.RWMutex
	getNicPropertiesArgsForCall []struct {
		arg1 string
	}
	getNicPropertiesReturns struct {
		result1 *zhmcclient.NIC
		result2 int
		result3 *zhmcclient.HmcError
	}
	getNicPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.NIC
		result2 int
		result3 *zhmcclient.HmcError
	}
	UpdateNicPropertiesStub        func(string, *zhmcclient.NIC) (int, *zhmcclient.HmcError)
	updateNicPropertiesMutex       sync.RWMutex
	updateNicPropertiesArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.NIC
	}
	updateNicPropertiesReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	updateNicPropertiesReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NicAPI) CreateNic(arg1 string, arg2 *zhmcclient.NIC) (string, int, *zhmcclient.HmcError) {
	fake.createNicMutex.Lock()
	ret, specificReturn := fake.createNicReturnsOnCall[len(fake.createNicArgsForCall)]
	fake.createNicArgsForCall = append(fake.createNicArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.NIC
	}{arg1, arg2})
	stub := fake.CreateNicStub
	fakeReturns := fake.createNicReturns
	fake.recordInvocation("CreateNic", []interface{}{arg1, arg2})
	fake.createNicMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *NicAPI) CreateNicCallCount() int {
	fake.createNicMutex.RLock()
	defer fake.createNicMutex.RUnlock()
	return len(fake.createNicArgsForCall)
}

func (fake *NicAPI) CreateNicCalls(stub func(string, *zhmcclient.NIC) (string, int, *zhmcclient.HmcError)) {
	fake.createNicMutex.Lock()
	defer fake.createNicMutex.Unlock()
	fake.CreateNicStub = stub
}

func (fake *NicAPI) CreateNicArgsForCall(i int) (string, *zhmcclient.NIC) {
	fake.createNicMutex.RLock()
	defer fake.createNicMutex.RUnlock()
	argsForCall := fake.createNicArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *NicAPI) CreateNicReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createNicMutex.Lock()
	defer fake.createNicMutex.Unlock()
	fake.CreateNicStub = nil
	fake.createNicReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *NicAPI) CreateNicReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createNicMutex.Lock()
	defer fake.createNicMutex.Unlock()
	fake.CreateNicStub = nil
	if fake.createNicReturnsOnCall == nil {
		fake.createNicReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.createNicReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *NicAPI) DeleteNic(arg1 string) (int, *zhmcclient.HmcError) {
	fake.deleteNicMutex.Lock()
	ret, specificReturn := fake.deleteNicReturnsOnCall[len(fake.deleteNicArgsForCall)]
	fake.deleteNicArgsForCall = append(fake.deleteNicArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteNicStub
	fakeReturns := fake.deleteNicReturns
	fake.recordInvocation("DeleteNic", []interface{}{arg1})
	fake.deleteNicMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *NicAPI) DeleteNicCallCount() int {
	fake.deleteNicMutex.RLock()
	defer fake.deleteNicMutex.RUnlock()
	return len(fake.deleteNicArgsForCall)
}

func (fake *NicAPI) DeleteNicCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.deleteNicMutex.Lock()
	defer fake.deleteNicMutex.Unlock()
	fake.DeleteNicStub = stub
}

func (fake *NicAPI) DeleteNicArgsForCall(i int) string {
	fake.deleteNicMutex.RLock()
	defer fake.deleteNicMutex.RUnlock()
	argsForCall := fake.deleteNicArgsForCall[i]
	return argsForCall.arg1
}

func (fake *NicAPI) DeleteNicReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteNicMutex.Lock()
	defer fake.deleteNicMutex.Unlock()
	fake.DeleteNicStub = nil
	fake.deleteNicReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *NicAPI) DeleteNicReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteNicMutex.Lock()
	defer fake.deleteNicMutex.Unlock()
	fake.DeleteNicStub = nil
	if fake.deleteNicReturnsOnCall == nil {
		fake.deleteNicReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.deleteNicReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *NicAPI) GetNicProperties(arg1 string) (*zhmcclient.NIC, int, *zhmcclient.HmcError) {
	fake.getNicPropertiesMutex.Lock()
	ret, specificReturn := fake.getNicPropertiesReturnsOnCall[len(fake.getNicPropertiesArgsForCall)]
	fake.getNicPropertiesArgsForCall = append(fake.getNicPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetNicPropertiesStub
	fakeReturns := fake.getNicPropertiesReturns
	fake.recordInvocation("GetNicProperties", []interface{}{arg1})
	fake.getNicPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *NicAPI) GetNicPropertiesCallCount() int {
	fake.getNicPropertiesMutex.RLock()
	defer fake.getNicPropertiesMutex.RUnlock()
	return len(fake.getNicPropertiesArgsForCall)
}

func (fake *NicAPI) GetNicPropertiesCalls(stub func(string) (*zhmcclient.NIC, int, *zhmcclient.HmcError)) {
	fake.getNicPropertiesMutex.Lock()
	defer fake.getNicPropertiesMutex.Unlock()
	fake.GetNicPropertiesStub = stub
}

func (fake *NicAPI) GetNicPropertiesArgsForCall(i int) string {
	fake.getNicPropertiesMutex.RLock()
	defer fake.getNicPropertiesMutex.RUnlock()
	argsForCall := fake.getNicPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *NicAPI) GetNicPropertiesReturns(result1 *zhmcclient.NIC, result2 int, result3 *zhmcclient.HmcError) {
	fake.getNicPropertiesMutex.Lock()
	defer fake.getNicPropertiesMutex.Unlock()
	fake.GetNicPropertiesStub = nil
	fake.getNicPropertiesReturns = struct {
		result1 *zhmcclient.NIC
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *NicAPI) GetNicPropertiesReturnsOnCall(i int, result1 *zhmcclient.NIC, result2 int, result3 *zhmcclient.HmcError) {
	fake.getNicPropertiesMutex.Lock()
	defer fake.getNicPropertiesMutex.Unlock()
	fake.GetNicPropertiesStub = nil
	if fake.getNicPropertiesReturnsOnCall == nil {
		fake.getNicPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.NIC
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getNicPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.NIC
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *NicAPI) UpdateNicProperties(arg1 string, arg2 *zhmcclient.NIC) (int, *zhmcclient.HmcError) {
	fake.updateNicPropertiesMutex.Lock()
	ret, specificReturn := fake.updateNicPropertiesReturnsOnCall[len(fake.updateNicPropertiesArgsForCall)]
	fake.updateNicPropertiesArgsForCall = append(fake.updateNicPropertiesArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.NIC
	}{arg1, arg2})
	stub := fake.UpdateNicPropertiesStub
	fakeReturns := fake.updateNicPropertiesReturns
	fake.recordInvocation("UpdateNicProperties", []interface{}{arg1, arg2})
	fake.updateNicPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *NicAPI) UpdateNicPropertiesCallCount() int {
	fake.updateNicPropertiesMutex.RLock()
	defer fake.updateNicPropertiesMutex.RUnlock()
	return len(fake.updateNicPropertiesArgsForCall)
}

func (fake *NicAPI) UpdateNicPropertiesCalls(stub func(string, *zhmcclient.NIC) (int, *zhmcclient.HmcError)) {
	fake.updateNicPropertiesMutex.Lock()
	defer fake.updateNicPropertiesMutex.Unlock()
	fake.UpdateNicPropertiesStub = stub
}

func (fake *NicAPI) UpdateNicPropertiesArgsForCall(i int) (string, *zhmcclient.NIC) {
	fake.updateNicPropertiesMutex.RLock()
	defer fake.updateNicPropertiesMutex.RUnlock()
	argsForCall := fake.updateNicPropertiesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *NicAPI) UpdateNicPropertiesReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.updateNicPropertiesMutex.Lock()
	defer fake.updateNicPropertiesMutex.Unlock()
	fake.UpdateNicPropertiesStub = nil
	fake.updateNicPropertiesReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *NicAPI) UpdateNicPropertiesReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.updateNicPropertiesMutex.Lock()
	defer fake.updateNicPropertiesMutex.Unlock()
	fake.UpdateNicPropertiesStub = nil
	if fake.updateNicPropertiesReturnsOnCall == nil {
		fake.updateNicPropertiesReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.updateNicPropertiesReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *NicAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createNicMutex.RLock()
	defer fake.createNicMutex.RUnlock()
	fake.deleteNicMutex.RLock()
	defer fake.deleteNicMutex.RUnlock()
	fake.getNicPropertiesMutex.RLock()
	defer fake.getNicPropertiesMutex.RUnlock()
	fake.updateNicPropertiesMutex.RLock()
	defer fake.updateNicPropertiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *NicAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.NicAPI = new(NicAPI)
