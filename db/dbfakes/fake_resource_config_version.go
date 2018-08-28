// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakeResourceConfigVersion struct {
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	VersionStub        func() db.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 db.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 db.Version
	}
	MetadataStub        func() db.ResourceConfigMetadataFields
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct{}
	metadataReturns     struct {
		result1 db.ResourceConfigMetadataFields
	}
	metadataReturnsOnCall map[int]struct {
		result1 db.ResourceConfigMetadataFields
	}
	CheckOrderStub        func() int
	checkOrderMutex       sync.RWMutex
	checkOrderArgsForCall []struct{}
	checkOrderReturns     struct {
		result1 int
	}
	checkOrderReturnsOnCall map[int]struct {
		result1 int
	}
	ResourceConfigStub        func() db.ResourceConfig
	resourceConfigMutex       sync.RWMutex
	resourceConfigArgsForCall []struct{}
	resourceConfigReturns     struct {
		result1 db.ResourceConfig
	}
	resourceConfigReturnsOnCall map[int]struct {
		result1 db.ResourceConfig
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct{}
	reloadReturns     struct {
		result1 bool
		result2 error
	}
	reloadReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	SaveMetadataStub        func(db.ResourceConfigMetadataFields) error
	saveMetadataMutex       sync.RWMutex
	saveMetadataArgsForCall []struct {
		arg1 db.ResourceConfigMetadataFields
	}
	saveMetadataReturns struct {
		result1 error
	}
	saveMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceConfigVersion) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.iDReturns.result1
}

func (fake *FakeResourceConfigVersion) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeResourceConfigVersion) IDReturns(result1 int) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceConfigVersion) IDReturnsOnCall(i int, result1 int) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceConfigVersion) Version() db.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.versionReturns.result1
}

func (fake *FakeResourceConfigVersion) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeResourceConfigVersion) VersionReturns(result1 db.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 db.Version
	}{result1}
}

func (fake *FakeResourceConfigVersion) VersionReturnsOnCall(i int, result1 db.Version) {
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 db.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 db.Version
	}{result1}
}

func (fake *FakeResourceConfigVersion) Metadata() db.ResourceConfigMetadataFields {
	fake.metadataMutex.Lock()
	ret, specificReturn := fake.metadataReturnsOnCall[len(fake.metadataArgsForCall)]
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct{}{})
	fake.recordInvocation("Metadata", []interface{}{})
	fake.metadataMutex.Unlock()
	if fake.MetadataStub != nil {
		return fake.MetadataStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.metadataReturns.result1
}

func (fake *FakeResourceConfigVersion) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *FakeResourceConfigVersion) MetadataReturns(result1 db.ResourceConfigMetadataFields) {
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 db.ResourceConfigMetadataFields
	}{result1}
}

func (fake *FakeResourceConfigVersion) MetadataReturnsOnCall(i int, result1 db.ResourceConfigMetadataFields) {
	fake.MetadataStub = nil
	if fake.metadataReturnsOnCall == nil {
		fake.metadataReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfigMetadataFields
		})
	}
	fake.metadataReturnsOnCall[i] = struct {
		result1 db.ResourceConfigMetadataFields
	}{result1}
}

func (fake *FakeResourceConfigVersion) CheckOrder() int {
	fake.checkOrderMutex.Lock()
	ret, specificReturn := fake.checkOrderReturnsOnCall[len(fake.checkOrderArgsForCall)]
	fake.checkOrderArgsForCall = append(fake.checkOrderArgsForCall, struct{}{})
	fake.recordInvocation("CheckOrder", []interface{}{})
	fake.checkOrderMutex.Unlock()
	if fake.CheckOrderStub != nil {
		return fake.CheckOrderStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.checkOrderReturns.result1
}

func (fake *FakeResourceConfigVersion) CheckOrderCallCount() int {
	fake.checkOrderMutex.RLock()
	defer fake.checkOrderMutex.RUnlock()
	return len(fake.checkOrderArgsForCall)
}

func (fake *FakeResourceConfigVersion) CheckOrderReturns(result1 int) {
	fake.CheckOrderStub = nil
	fake.checkOrderReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceConfigVersion) CheckOrderReturnsOnCall(i int, result1 int) {
	fake.CheckOrderStub = nil
	if fake.checkOrderReturnsOnCall == nil {
		fake.checkOrderReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.checkOrderReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceConfigVersion) ResourceConfig() db.ResourceConfig {
	fake.resourceConfigMutex.Lock()
	ret, specificReturn := fake.resourceConfigReturnsOnCall[len(fake.resourceConfigArgsForCall)]
	fake.resourceConfigArgsForCall = append(fake.resourceConfigArgsForCall, struct{}{})
	fake.recordInvocation("ResourceConfig", []interface{}{})
	fake.resourceConfigMutex.Unlock()
	if fake.ResourceConfigStub != nil {
		return fake.ResourceConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.resourceConfigReturns.result1
}

func (fake *FakeResourceConfigVersion) ResourceConfigCallCount() int {
	fake.resourceConfigMutex.RLock()
	defer fake.resourceConfigMutex.RUnlock()
	return len(fake.resourceConfigArgsForCall)
}

func (fake *FakeResourceConfigVersion) ResourceConfigReturns(result1 db.ResourceConfig) {
	fake.ResourceConfigStub = nil
	fake.resourceConfigReturns = struct {
		result1 db.ResourceConfig
	}{result1}
}

func (fake *FakeResourceConfigVersion) ResourceConfigReturnsOnCall(i int, result1 db.ResourceConfig) {
	fake.ResourceConfigStub = nil
	if fake.resourceConfigReturnsOnCall == nil {
		fake.resourceConfigReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfig
		})
	}
	fake.resourceConfigReturnsOnCall[i] = struct {
		result1 db.ResourceConfig
	}{result1}
}

func (fake *FakeResourceConfigVersion) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	ret, specificReturn := fake.reloadReturnsOnCall[len(fake.reloadArgsForCall)]
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct{}{})
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if fake.ReloadStub != nil {
		return fake.ReloadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.reloadReturns.result1, fake.reloadReturns.result2
}

func (fake *FakeResourceConfigVersion) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeResourceConfigVersion) ReloadReturns(result1 bool, result2 error) {
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigVersion) ReloadReturnsOnCall(i int, result1 bool, result2 error) {
	fake.ReloadStub = nil
	if fake.reloadReturnsOnCall == nil {
		fake.reloadReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.reloadReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigVersion) SaveMetadata(arg1 db.ResourceConfigMetadataFields) error {
	fake.saveMetadataMutex.Lock()
	ret, specificReturn := fake.saveMetadataReturnsOnCall[len(fake.saveMetadataArgsForCall)]
	fake.saveMetadataArgsForCall = append(fake.saveMetadataArgsForCall, struct {
		arg1 db.ResourceConfigMetadataFields
	}{arg1})
	fake.recordInvocation("SaveMetadata", []interface{}{arg1})
	fake.saveMetadataMutex.Unlock()
	if fake.SaveMetadataStub != nil {
		return fake.SaveMetadataStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.saveMetadataReturns.result1
}

func (fake *FakeResourceConfigVersion) SaveMetadataCallCount() int {
	fake.saveMetadataMutex.RLock()
	defer fake.saveMetadataMutex.RUnlock()
	return len(fake.saveMetadataArgsForCall)
}

func (fake *FakeResourceConfigVersion) SaveMetadataArgsForCall(i int) db.ResourceConfigMetadataFields {
	fake.saveMetadataMutex.RLock()
	defer fake.saveMetadataMutex.RUnlock()
	return fake.saveMetadataArgsForCall[i].arg1
}

func (fake *FakeResourceConfigVersion) SaveMetadataReturns(result1 error) {
	fake.SaveMetadataStub = nil
	fake.saveMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigVersion) SaveMetadataReturnsOnCall(i int, result1 error) {
	fake.SaveMetadataStub = nil
	if fake.saveMetadataReturnsOnCall == nil {
		fake.saveMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigVersion) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	fake.checkOrderMutex.RLock()
	defer fake.checkOrderMutex.RUnlock()
	fake.resourceConfigMutex.RLock()
	defer fake.resourceConfigMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	fake.saveMetadataMutex.RLock()
	defer fake.saveMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceConfigVersion) recordInvocation(key string, args []interface{}) {
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

var _ db.ResourceConfigVersion = new(FakeResourceConfigVersion)
