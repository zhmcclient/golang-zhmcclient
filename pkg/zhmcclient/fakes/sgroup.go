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

	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

type StorageGroupAPI struct {
	CreateStorageGroupsStub        func(string, *zhmcclient.CreateStorageGroupProperties) (*zhmcclient.StorageGroupCreateResponse, int, *zhmcclient.HmcError)
	createStorageGroupsMutex       sync.RWMutex
	createStorageGroupsArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.CreateStorageGroupProperties
	}
	createStorageGroupsReturns struct {
		result1 *zhmcclient.StorageGroupCreateResponse
		result2 int
		result3 *zhmcclient.HmcError
	}
	createStorageGroupsReturnsOnCall map[int]struct {
		result1 *zhmcclient.StorageGroupCreateResponse
		result2 int
		result3 *zhmcclient.HmcError
	}
	DeleteStorageGroupStub        func(string) (int, *zhmcclient.HmcError)
	deleteStorageGroupMutex       sync.RWMutex
	deleteStorageGroupArgsForCall []struct {
		arg1 string
	}
	deleteStorageGroupReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	deleteStorageGroupReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	FulfillStorageGroupStub        func(string, *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError)
	fulfillStorageGroupMutex       sync.RWMutex
	fulfillStorageGroupArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupProperties
	}
	fulfillStorageGroupReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	fulfillStorageGroupReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	GetStorageGroupPartitionsStub        func(string, map[string]string) (*zhmcclient.StorageGroupPartitions, int, *zhmcclient.HmcError)
	getStorageGroupPartitionsMutex       sync.RWMutex
	getStorageGroupPartitionsArgsForCall []struct {
		arg1 string
		arg2 map[string]string
	}
	getStorageGroupPartitionsReturns struct {
		result1 *zhmcclient.StorageGroupPartitions
		result2 int
		result3 *zhmcclient.HmcError
	}
	getStorageGroupPartitionsReturnsOnCall map[int]struct {
		result1 *zhmcclient.StorageGroupPartitions
		result2 int
		result3 *zhmcclient.HmcError
	}
	GetStorageGroupPropertiesStub        func(string) (*zhmcclient.StorageGroupProperties, int, *zhmcclient.HmcError)
	getStorageGroupPropertiesMutex       sync.RWMutex
	getStorageGroupPropertiesArgsForCall []struct {
		arg1 string
	}
	getStorageGroupPropertiesReturns struct {
		result1 *zhmcclient.StorageGroupProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	getStorageGroupPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.StorageGroupProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	GetStorageVolumePropertiesStub        func(string) (*zhmcclient.StorageVolume, int, *zhmcclient.HmcError)
	getStorageVolumePropertiesMutex       sync.RWMutex
	getStorageVolumePropertiesArgsForCall []struct {
		arg1 string
	}
	getStorageVolumePropertiesReturns struct {
		result1 *zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}
	getStorageVolumePropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}
	ListStorageGroupsStub        func(string, string) ([]zhmcclient.StorageGroup, int, *zhmcclient.HmcError)
	listStorageGroupsMutex       sync.RWMutex
	listStorageGroupsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listStorageGroupsReturns struct {
		result1 []zhmcclient.StorageGroup
		result2 int
		result3 *zhmcclient.HmcError
	}
	listStorageGroupsReturnsOnCall map[int]struct {
		result1 []zhmcclient.StorageGroup
		result2 int
		result3 *zhmcclient.HmcError
	}
	ListStorageVolumesStub        func(string) ([]zhmcclient.StorageVolume, int, *zhmcclient.HmcError)
	listStorageVolumesMutex       sync.RWMutex
	listStorageVolumesArgsForCall []struct {
		arg1 string
	}
	listStorageVolumesReturns struct {
		result1 []zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}
	listStorageVolumesReturnsOnCall map[int]struct {
		result1 []zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}
	UpdateStorageGroupPropertiesStub        func(string, *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError)
	updateStorageGroupPropertiesMutex       sync.RWMutex
	updateStorageGroupPropertiesArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupProperties
	}
	updateStorageGroupPropertiesReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	updateStorageGroupPropertiesReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StorageGroupAPI) CreateStorageGroups(arg1 string, arg2 *zhmcclient.CreateStorageGroupProperties) (*zhmcclient.StorageGroupCreateResponse, int, *zhmcclient.HmcError) {
	fake.createStorageGroupsMutex.Lock()
	ret, specificReturn := fake.createStorageGroupsReturnsOnCall[len(fake.createStorageGroupsArgsForCall)]
	fake.createStorageGroupsArgsForCall = append(fake.createStorageGroupsArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.CreateStorageGroupProperties
	}{arg1, arg2})
	stub := fake.CreateStorageGroupsStub
	fakeReturns := fake.createStorageGroupsReturns
	fake.recordInvocation("CreateStorageGroups", []interface{}{arg1, arg2})
	fake.createStorageGroupsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *StorageGroupAPI) CreateStorageGroupsCallCount() int {
	fake.createStorageGroupsMutex.RLock()
	defer fake.createStorageGroupsMutex.RUnlock()
	return len(fake.createStorageGroupsArgsForCall)
}

func (fake *StorageGroupAPI) CreateStorageGroupsCalls(stub func(string, *zhmcclient.CreateStorageGroupProperties) (*zhmcclient.StorageGroupCreateResponse, int, *zhmcclient.HmcError)) {
	fake.createStorageGroupsMutex.Lock()
	defer fake.createStorageGroupsMutex.Unlock()
	fake.CreateStorageGroupsStub = stub
}

func (fake *StorageGroupAPI) CreateStorageGroupsArgsForCall(i int) (string, *zhmcclient.CreateStorageGroupProperties) {
	fake.createStorageGroupsMutex.RLock()
	defer fake.createStorageGroupsMutex.RUnlock()
	argsForCall := fake.createStorageGroupsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *StorageGroupAPI) CreateStorageGroupsReturns(result1 *zhmcclient.StorageGroupCreateResponse, result2 int, result3 *zhmcclient.HmcError) {
	fake.createStorageGroupsMutex.Lock()
	defer fake.createStorageGroupsMutex.Unlock()
	fake.CreateStorageGroupsStub = nil
	fake.createStorageGroupsReturns = struct {
		result1 *zhmcclient.StorageGroupCreateResponse
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) CreateStorageGroupsReturnsOnCall(i int, result1 *zhmcclient.StorageGroupCreateResponse, result2 int, result3 *zhmcclient.HmcError) {
	fake.createStorageGroupsMutex.Lock()
	defer fake.createStorageGroupsMutex.Unlock()
	fake.CreateStorageGroupsStub = nil
	if fake.createStorageGroupsReturnsOnCall == nil {
		fake.createStorageGroupsReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.StorageGroupCreateResponse
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.createStorageGroupsReturnsOnCall[i] = struct {
		result1 *zhmcclient.StorageGroupCreateResponse
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) DeleteStorageGroup(arg1 string) (int, *zhmcclient.HmcError) {
	fake.deleteStorageGroupMutex.Lock()
	ret, specificReturn := fake.deleteStorageGroupReturnsOnCall[len(fake.deleteStorageGroupArgsForCall)]
	fake.deleteStorageGroupArgsForCall = append(fake.deleteStorageGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteStorageGroupStub
	fakeReturns := fake.deleteStorageGroupReturns
	fake.recordInvocation("DeleteStorageGroup", []interface{}{arg1})
	fake.deleteStorageGroupMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StorageGroupAPI) DeleteStorageGroupCallCount() int {
	fake.deleteStorageGroupMutex.RLock()
	defer fake.deleteStorageGroupMutex.RUnlock()
	return len(fake.deleteStorageGroupArgsForCall)
}

func (fake *StorageGroupAPI) DeleteStorageGroupCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.deleteStorageGroupMutex.Lock()
	defer fake.deleteStorageGroupMutex.Unlock()
	fake.DeleteStorageGroupStub = stub
}

func (fake *StorageGroupAPI) DeleteStorageGroupArgsForCall(i int) string {
	fake.deleteStorageGroupMutex.RLock()
	defer fake.deleteStorageGroupMutex.RUnlock()
	argsForCall := fake.deleteStorageGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StorageGroupAPI) DeleteStorageGroupReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteStorageGroupMutex.Lock()
	defer fake.deleteStorageGroupMutex.Unlock()
	fake.DeleteStorageGroupStub = nil
	fake.deleteStorageGroupReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *StorageGroupAPI) DeleteStorageGroupReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteStorageGroupMutex.Lock()
	defer fake.deleteStorageGroupMutex.Unlock()
	fake.DeleteStorageGroupStub = nil
	if fake.deleteStorageGroupReturnsOnCall == nil {
		fake.deleteStorageGroupReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.deleteStorageGroupReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *StorageGroupAPI) FulfillStorageGroup(arg1 string, arg2 *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError) {
	fake.fulfillStorageGroupMutex.Lock()
	ret, specificReturn := fake.fulfillStorageGroupReturnsOnCall[len(fake.fulfillStorageGroupArgsForCall)]
	fake.fulfillStorageGroupArgsForCall = append(fake.fulfillStorageGroupArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupProperties
	}{arg1, arg2})
	stub := fake.FulfillStorageGroupStub
	fakeReturns := fake.fulfillStorageGroupReturns
	fake.recordInvocation("FulfillStorageGroup", []interface{}{arg1, arg2})
	fake.fulfillStorageGroupMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StorageGroupAPI) FulfillStorageGroupCallCount() int {
	fake.fulfillStorageGroupMutex.RLock()
	defer fake.fulfillStorageGroupMutex.RUnlock()
	return len(fake.fulfillStorageGroupArgsForCall)
}

func (fake *StorageGroupAPI) FulfillStorageGroupCalls(stub func(string, *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError)) {
	fake.fulfillStorageGroupMutex.Lock()
	defer fake.fulfillStorageGroupMutex.Unlock()
	fake.FulfillStorageGroupStub = stub
}

func (fake *StorageGroupAPI) FulfillStorageGroupArgsForCall(i int) (string, *zhmcclient.StorageGroupProperties) {
	fake.fulfillStorageGroupMutex.RLock()
	defer fake.fulfillStorageGroupMutex.RUnlock()
	argsForCall := fake.fulfillStorageGroupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *StorageGroupAPI) FulfillStorageGroupReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.fulfillStorageGroupMutex.Lock()
	defer fake.fulfillStorageGroupMutex.Unlock()
	fake.FulfillStorageGroupStub = nil
	fake.fulfillStorageGroupReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *StorageGroupAPI) FulfillStorageGroupReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.fulfillStorageGroupMutex.Lock()
	defer fake.fulfillStorageGroupMutex.Unlock()
	fake.FulfillStorageGroupStub = nil
	if fake.fulfillStorageGroupReturnsOnCall == nil {
		fake.fulfillStorageGroupReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.fulfillStorageGroupReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *StorageGroupAPI) GetStorageGroupPartitions(arg1 string, arg2 map[string]string) (*zhmcclient.StorageGroupPartitions, int, *zhmcclient.HmcError) {
	fake.getStorageGroupPartitionsMutex.Lock()
	ret, specificReturn := fake.getStorageGroupPartitionsReturnsOnCall[len(fake.getStorageGroupPartitionsArgsForCall)]
	fake.getStorageGroupPartitionsArgsForCall = append(fake.getStorageGroupPartitionsArgsForCall, struct {
		arg1 string
		arg2 map[string]string
	}{arg1, arg2})
	stub := fake.GetStorageGroupPartitionsStub
	fakeReturns := fake.getStorageGroupPartitionsReturns
	fake.recordInvocation("GetStorageGroupPartitions", []interface{}{arg1, arg2})
	fake.getStorageGroupPartitionsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *StorageGroupAPI) GetStorageGroupPartitionsCallCount() int {
	fake.getStorageGroupPartitionsMutex.RLock()
	defer fake.getStorageGroupPartitionsMutex.RUnlock()
	return len(fake.getStorageGroupPartitionsArgsForCall)
}

func (fake *StorageGroupAPI) GetStorageGroupPartitionsCalls(stub func(string, map[string]string) (*zhmcclient.StorageGroupPartitions, int, *zhmcclient.HmcError)) {
	fake.getStorageGroupPartitionsMutex.Lock()
	defer fake.getStorageGroupPartitionsMutex.Unlock()
	fake.GetStorageGroupPartitionsStub = stub
}

func (fake *StorageGroupAPI) GetStorageGroupPartitionsArgsForCall(i int) (string, map[string]string) {
	fake.getStorageGroupPartitionsMutex.RLock()
	defer fake.getStorageGroupPartitionsMutex.RUnlock()
	argsForCall := fake.getStorageGroupPartitionsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *StorageGroupAPI) GetStorageGroupPartitionsReturns(result1 *zhmcclient.StorageGroupPartitions, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageGroupPartitionsMutex.Lock()
	defer fake.getStorageGroupPartitionsMutex.Unlock()
	fake.GetStorageGroupPartitionsStub = nil
	fake.getStorageGroupPartitionsReturns = struct {
		result1 *zhmcclient.StorageGroupPartitions
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) GetStorageGroupPartitionsReturnsOnCall(i int, result1 *zhmcclient.StorageGroupPartitions, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageGroupPartitionsMutex.Lock()
	defer fake.getStorageGroupPartitionsMutex.Unlock()
	fake.GetStorageGroupPartitionsStub = nil
	if fake.getStorageGroupPartitionsReturnsOnCall == nil {
		fake.getStorageGroupPartitionsReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.StorageGroupPartitions
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getStorageGroupPartitionsReturnsOnCall[i] = struct {
		result1 *zhmcclient.StorageGroupPartitions
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) GetStorageGroupProperties(arg1 string) (*zhmcclient.StorageGroupProperties, int, *zhmcclient.HmcError) {
	fake.getStorageGroupPropertiesMutex.Lock()
	ret, specificReturn := fake.getStorageGroupPropertiesReturnsOnCall[len(fake.getStorageGroupPropertiesArgsForCall)]
	fake.getStorageGroupPropertiesArgsForCall = append(fake.getStorageGroupPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStorageGroupPropertiesStub
	fakeReturns := fake.getStorageGroupPropertiesReturns
	fake.recordInvocation("GetStorageGroupProperties", []interface{}{arg1})
	fake.getStorageGroupPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *StorageGroupAPI) GetStorageGroupPropertiesCallCount() int {
	fake.getStorageGroupPropertiesMutex.RLock()
	defer fake.getStorageGroupPropertiesMutex.RUnlock()
	return len(fake.getStorageGroupPropertiesArgsForCall)
}

func (fake *StorageGroupAPI) GetStorageGroupPropertiesCalls(stub func(string) (*zhmcclient.StorageGroupProperties, int, *zhmcclient.HmcError)) {
	fake.getStorageGroupPropertiesMutex.Lock()
	defer fake.getStorageGroupPropertiesMutex.Unlock()
	fake.GetStorageGroupPropertiesStub = stub
}

func (fake *StorageGroupAPI) GetStorageGroupPropertiesArgsForCall(i int) string {
	fake.getStorageGroupPropertiesMutex.RLock()
	defer fake.getStorageGroupPropertiesMutex.RUnlock()
	argsForCall := fake.getStorageGroupPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StorageGroupAPI) GetStorageGroupPropertiesReturns(result1 *zhmcclient.StorageGroupProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageGroupPropertiesMutex.Lock()
	defer fake.getStorageGroupPropertiesMutex.Unlock()
	fake.GetStorageGroupPropertiesStub = nil
	fake.getStorageGroupPropertiesReturns = struct {
		result1 *zhmcclient.StorageGroupProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) GetStorageGroupPropertiesReturnsOnCall(i int, result1 *zhmcclient.StorageGroupProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageGroupPropertiesMutex.Lock()
	defer fake.getStorageGroupPropertiesMutex.Unlock()
	fake.GetStorageGroupPropertiesStub = nil
	if fake.getStorageGroupPropertiesReturnsOnCall == nil {
		fake.getStorageGroupPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.StorageGroupProperties
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getStorageGroupPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.StorageGroupProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) GetStorageVolumeProperties(arg1 string) (*zhmcclient.StorageVolume, int, *zhmcclient.HmcError) {
	fake.getStorageVolumePropertiesMutex.Lock()
	ret, specificReturn := fake.getStorageVolumePropertiesReturnsOnCall[len(fake.getStorageVolumePropertiesArgsForCall)]
	fake.getStorageVolumePropertiesArgsForCall = append(fake.getStorageVolumePropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStorageVolumePropertiesStub
	fakeReturns := fake.getStorageVolumePropertiesReturns
	fake.recordInvocation("GetStorageVolumeProperties", []interface{}{arg1})
	fake.getStorageVolumePropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *StorageGroupAPI) GetStorageVolumePropertiesCallCount() int {
	fake.getStorageVolumePropertiesMutex.RLock()
	defer fake.getStorageVolumePropertiesMutex.RUnlock()
	return len(fake.getStorageVolumePropertiesArgsForCall)
}

func (fake *StorageGroupAPI) GetStorageVolumePropertiesCalls(stub func(string) (*zhmcclient.StorageVolume, int, *zhmcclient.HmcError)) {
	fake.getStorageVolumePropertiesMutex.Lock()
	defer fake.getStorageVolumePropertiesMutex.Unlock()
	fake.GetStorageVolumePropertiesStub = stub
}

func (fake *StorageGroupAPI) GetStorageVolumePropertiesArgsForCall(i int) string {
	fake.getStorageVolumePropertiesMutex.RLock()
	defer fake.getStorageVolumePropertiesMutex.RUnlock()
	argsForCall := fake.getStorageVolumePropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StorageGroupAPI) GetStorageVolumePropertiesReturns(result1 *zhmcclient.StorageVolume, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageVolumePropertiesMutex.Lock()
	defer fake.getStorageVolumePropertiesMutex.Unlock()
	fake.GetStorageVolumePropertiesStub = nil
	fake.getStorageVolumePropertiesReturns = struct {
		result1 *zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) GetStorageVolumePropertiesReturnsOnCall(i int, result1 *zhmcclient.StorageVolume, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageVolumePropertiesMutex.Lock()
	defer fake.getStorageVolumePropertiesMutex.Unlock()
	fake.GetStorageVolumePropertiesStub = nil
	if fake.getStorageVolumePropertiesReturnsOnCall == nil {
		fake.getStorageVolumePropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.StorageVolume
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getStorageVolumePropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) ListStorageGroups(arg1 string, arg2 string) ([]zhmcclient.StorageGroup, int, *zhmcclient.HmcError) {
	fake.listStorageGroupsMutex.Lock()
	ret, specificReturn := fake.listStorageGroupsReturnsOnCall[len(fake.listStorageGroupsArgsForCall)]
	fake.listStorageGroupsArgsForCall = append(fake.listStorageGroupsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ListStorageGroupsStub
	fakeReturns := fake.listStorageGroupsReturns
	fake.recordInvocation("ListStorageGroups", []interface{}{arg1, arg2})
	fake.listStorageGroupsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *StorageGroupAPI) ListStorageGroupsCallCount() int {
	fake.listStorageGroupsMutex.RLock()
	defer fake.listStorageGroupsMutex.RUnlock()
	return len(fake.listStorageGroupsArgsForCall)
}

func (fake *StorageGroupAPI) ListStorageGroupsCalls(stub func(string, string) ([]zhmcclient.StorageGroup, int, *zhmcclient.HmcError)) {
	fake.listStorageGroupsMutex.Lock()
	defer fake.listStorageGroupsMutex.Unlock()
	fake.ListStorageGroupsStub = stub
}

func (fake *StorageGroupAPI) ListStorageGroupsArgsForCall(i int) (string, string) {
	fake.listStorageGroupsMutex.RLock()
	defer fake.listStorageGroupsMutex.RUnlock()
	argsForCall := fake.listStorageGroupsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *StorageGroupAPI) ListStorageGroupsReturns(result1 []zhmcclient.StorageGroup, result2 int, result3 *zhmcclient.HmcError) {
	fake.listStorageGroupsMutex.Lock()
	defer fake.listStorageGroupsMutex.Unlock()
	fake.ListStorageGroupsStub = nil
	fake.listStorageGroupsReturns = struct {
		result1 []zhmcclient.StorageGroup
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) ListStorageGroupsReturnsOnCall(i int, result1 []zhmcclient.StorageGroup, result2 int, result3 *zhmcclient.HmcError) {
	fake.listStorageGroupsMutex.Lock()
	defer fake.listStorageGroupsMutex.Unlock()
	fake.ListStorageGroupsStub = nil
	if fake.listStorageGroupsReturnsOnCall == nil {
		fake.listStorageGroupsReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.StorageGroup
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listStorageGroupsReturnsOnCall[i] = struct {
		result1 []zhmcclient.StorageGroup
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) ListStorageVolumes(arg1 string) ([]zhmcclient.StorageVolume, int, *zhmcclient.HmcError) {
	fake.listStorageVolumesMutex.Lock()
	ret, specificReturn := fake.listStorageVolumesReturnsOnCall[len(fake.listStorageVolumesArgsForCall)]
	fake.listStorageVolumesArgsForCall = append(fake.listStorageVolumesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListStorageVolumesStub
	fakeReturns := fake.listStorageVolumesReturns
	fake.recordInvocation("ListStorageVolumes", []interface{}{arg1})
	fake.listStorageVolumesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *StorageGroupAPI) ListStorageVolumesCallCount() int {
	fake.listStorageVolumesMutex.RLock()
	defer fake.listStorageVolumesMutex.RUnlock()
	return len(fake.listStorageVolumesArgsForCall)
}

func (fake *StorageGroupAPI) ListStorageVolumesCalls(stub func(string) ([]zhmcclient.StorageVolume, int, *zhmcclient.HmcError)) {
	fake.listStorageVolumesMutex.Lock()
	defer fake.listStorageVolumesMutex.Unlock()
	fake.ListStorageVolumesStub = stub
}

func (fake *StorageGroupAPI) ListStorageVolumesArgsForCall(i int) string {
	fake.listStorageVolumesMutex.RLock()
	defer fake.listStorageVolumesMutex.RUnlock()
	argsForCall := fake.listStorageVolumesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StorageGroupAPI) ListStorageVolumesReturns(result1 []zhmcclient.StorageVolume, result2 int, result3 *zhmcclient.HmcError) {
	fake.listStorageVolumesMutex.Lock()
	defer fake.listStorageVolumesMutex.Unlock()
	fake.ListStorageVolumesStub = nil
	fake.listStorageVolumesReturns = struct {
		result1 []zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) ListStorageVolumesReturnsOnCall(i int, result1 []zhmcclient.StorageVolume, result2 int, result3 *zhmcclient.HmcError) {
	fake.listStorageVolumesMutex.Lock()
	defer fake.listStorageVolumesMutex.Unlock()
	fake.ListStorageVolumesStub = nil
	if fake.listStorageVolumesReturnsOnCall == nil {
		fake.listStorageVolumesReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.StorageVolume
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listStorageVolumesReturnsOnCall[i] = struct {
		result1 []zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *StorageGroupAPI) UpdateStorageGroupProperties(arg1 string, arg2 *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError) {
	fake.updateStorageGroupPropertiesMutex.Lock()
	ret, specificReturn := fake.updateStorageGroupPropertiesReturnsOnCall[len(fake.updateStorageGroupPropertiesArgsForCall)]
	fake.updateStorageGroupPropertiesArgsForCall = append(fake.updateStorageGroupPropertiesArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupProperties
	}{arg1, arg2})
	stub := fake.UpdateStorageGroupPropertiesStub
	fakeReturns := fake.updateStorageGroupPropertiesReturns
	fake.recordInvocation("UpdateStorageGroupProperties", []interface{}{arg1, arg2})
	fake.updateStorageGroupPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StorageGroupAPI) UpdateStorageGroupPropertiesCallCount() int {
	fake.updateStorageGroupPropertiesMutex.RLock()
	defer fake.updateStorageGroupPropertiesMutex.RUnlock()
	return len(fake.updateStorageGroupPropertiesArgsForCall)
}

func (fake *StorageGroupAPI) UpdateStorageGroupPropertiesCalls(stub func(string, *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError)) {
	fake.updateStorageGroupPropertiesMutex.Lock()
	defer fake.updateStorageGroupPropertiesMutex.Unlock()
	fake.UpdateStorageGroupPropertiesStub = stub
}

func (fake *StorageGroupAPI) UpdateStorageGroupPropertiesArgsForCall(i int) (string, *zhmcclient.StorageGroupProperties) {
	fake.updateStorageGroupPropertiesMutex.RLock()
	defer fake.updateStorageGroupPropertiesMutex.RUnlock()
	argsForCall := fake.updateStorageGroupPropertiesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *StorageGroupAPI) UpdateStorageGroupPropertiesReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.updateStorageGroupPropertiesMutex.Lock()
	defer fake.updateStorageGroupPropertiesMutex.Unlock()
	fake.UpdateStorageGroupPropertiesStub = nil
	fake.updateStorageGroupPropertiesReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *StorageGroupAPI) UpdateStorageGroupPropertiesReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.updateStorageGroupPropertiesMutex.Lock()
	defer fake.updateStorageGroupPropertiesMutex.Unlock()
	fake.UpdateStorageGroupPropertiesStub = nil
	if fake.updateStorageGroupPropertiesReturnsOnCall == nil {
		fake.updateStorageGroupPropertiesReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.updateStorageGroupPropertiesReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *StorageGroupAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createStorageGroupsMutex.RLock()
	defer fake.createStorageGroupsMutex.RUnlock()
	fake.deleteStorageGroupMutex.RLock()
	defer fake.deleteStorageGroupMutex.RUnlock()
	fake.fulfillStorageGroupMutex.RLock()
	defer fake.fulfillStorageGroupMutex.RUnlock()
	fake.getStorageGroupPartitionsMutex.RLock()
	defer fake.getStorageGroupPartitionsMutex.RUnlock()
	fake.getStorageGroupPropertiesMutex.RLock()
	defer fake.getStorageGroupPropertiesMutex.RUnlock()
	fake.getStorageVolumePropertiesMutex.RLock()
	defer fake.getStorageVolumePropertiesMutex.RUnlock()
	fake.listStorageGroupsMutex.RLock()
	defer fake.listStorageGroupsMutex.RUnlock()
	fake.listStorageVolumesMutex.RLock()
	defer fake.listStorageVolumesMutex.RUnlock()
	fake.updateStorageGroupPropertiesMutex.RLock()
	defer fake.updateStorageGroupPropertiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StorageGroupAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.StorageGroupAPI = new(StorageGroupAPI)
