// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
)

type FakeResourceConfigFactory struct {
	FindOrCreateResourceConfigForBuildStub        func(build *dbng.Build, resourceType string, source atc.Source, pipeline *dbng.Pipeline, resourceTypes atc.ResourceTypes) (*dbng.UsedResourceConfig, error)
	findOrCreateResourceConfigForBuildMutex       sync.RWMutex
	findOrCreateResourceConfigForBuildArgsForCall []struct {
		build         *dbng.Build
		resourceType  string
		source        atc.Source
		pipeline      *dbng.Pipeline
		resourceTypes atc.ResourceTypes
	}
	findOrCreateResourceConfigForBuildReturns struct {
		result1 *dbng.UsedResourceConfig
		result2 error
	}
	FindOrCreateResourceConfigForResourceStub        func(resource *dbng.Resource, resourceType string, source atc.Source, pipeline *dbng.Pipeline, resourceTypes atc.ResourceTypes) (*dbng.UsedResourceConfig, error)
	findOrCreateResourceConfigForResourceMutex       sync.RWMutex
	findOrCreateResourceConfigForResourceArgsForCall []struct {
		resource      *dbng.Resource
		resourceType  string
		source        atc.Source
		pipeline      *dbng.Pipeline
		resourceTypes atc.ResourceTypes
	}
	findOrCreateResourceConfigForResourceReturns struct {
		result1 *dbng.UsedResourceConfig
		result2 error
	}
	FindOrCreateResourceConfigForResourceTypeStub        func(resourceTypeName string, source atc.Source, pipeline *dbng.Pipeline, resourceTypes atc.ResourceTypes) (*dbng.UsedResourceConfig, error)
	findOrCreateResourceConfigForResourceTypeMutex       sync.RWMutex
	findOrCreateResourceConfigForResourceTypeArgsForCall []struct {
		resourceTypeName string
		source           atc.Source
		pipeline         *dbng.Pipeline
		resourceTypes    atc.ResourceTypes
	}
	findOrCreateResourceConfigForResourceTypeReturns struct {
		result1 *dbng.UsedResourceConfig
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForBuild(build *dbng.Build, resourceType string, source atc.Source, pipeline *dbng.Pipeline, resourceTypes atc.ResourceTypes) (*dbng.UsedResourceConfig, error) {
	fake.findOrCreateResourceConfigForBuildMutex.Lock()
	fake.findOrCreateResourceConfigForBuildArgsForCall = append(fake.findOrCreateResourceConfigForBuildArgsForCall, struct {
		build         *dbng.Build
		resourceType  string
		source        atc.Source
		pipeline      *dbng.Pipeline
		resourceTypes atc.ResourceTypes
	}{build, resourceType, source, pipeline, resourceTypes})
	fake.recordInvocation("FindOrCreateResourceConfigForBuild", []interface{}{build, resourceType, source, pipeline, resourceTypes})
	fake.findOrCreateResourceConfigForBuildMutex.Unlock()
	if fake.FindOrCreateResourceConfigForBuildStub != nil {
		return fake.FindOrCreateResourceConfigForBuildStub(build, resourceType, source, pipeline, resourceTypes)
	} else {
		return fake.findOrCreateResourceConfigForBuildReturns.result1, fake.findOrCreateResourceConfigForBuildReturns.result2
	}
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForBuildCallCount() int {
	fake.findOrCreateResourceConfigForBuildMutex.RLock()
	defer fake.findOrCreateResourceConfigForBuildMutex.RUnlock()
	return len(fake.findOrCreateResourceConfigForBuildArgsForCall)
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForBuildArgsForCall(i int) (*dbng.Build, string, atc.Source, *dbng.Pipeline, atc.ResourceTypes) {
	fake.findOrCreateResourceConfigForBuildMutex.RLock()
	defer fake.findOrCreateResourceConfigForBuildMutex.RUnlock()
	return fake.findOrCreateResourceConfigForBuildArgsForCall[i].build, fake.findOrCreateResourceConfigForBuildArgsForCall[i].resourceType, fake.findOrCreateResourceConfigForBuildArgsForCall[i].source, fake.findOrCreateResourceConfigForBuildArgsForCall[i].pipeline, fake.findOrCreateResourceConfigForBuildArgsForCall[i].resourceTypes
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForBuildReturns(result1 *dbng.UsedResourceConfig, result2 error) {
	fake.FindOrCreateResourceConfigForBuildStub = nil
	fake.findOrCreateResourceConfigForBuildReturns = struct {
		result1 *dbng.UsedResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForResource(resource *dbng.Resource, resourceType string, source atc.Source, pipeline *dbng.Pipeline, resourceTypes atc.ResourceTypes) (*dbng.UsedResourceConfig, error) {
	fake.findOrCreateResourceConfigForResourceMutex.Lock()
	fake.findOrCreateResourceConfigForResourceArgsForCall = append(fake.findOrCreateResourceConfigForResourceArgsForCall, struct {
		resource      *dbng.Resource
		resourceType  string
		source        atc.Source
		pipeline      *dbng.Pipeline
		resourceTypes atc.ResourceTypes
	}{resource, resourceType, source, pipeline, resourceTypes})
	fake.recordInvocation("FindOrCreateResourceConfigForResource", []interface{}{resource, resourceType, source, pipeline, resourceTypes})
	fake.findOrCreateResourceConfigForResourceMutex.Unlock()
	if fake.FindOrCreateResourceConfigForResourceStub != nil {
		return fake.FindOrCreateResourceConfigForResourceStub(resource, resourceType, source, pipeline, resourceTypes)
	} else {
		return fake.findOrCreateResourceConfigForResourceReturns.result1, fake.findOrCreateResourceConfigForResourceReturns.result2
	}
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForResourceCallCount() int {
	fake.findOrCreateResourceConfigForResourceMutex.RLock()
	defer fake.findOrCreateResourceConfigForResourceMutex.RUnlock()
	return len(fake.findOrCreateResourceConfigForResourceArgsForCall)
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForResourceArgsForCall(i int) (*dbng.Resource, string, atc.Source, *dbng.Pipeline, atc.ResourceTypes) {
	fake.findOrCreateResourceConfigForResourceMutex.RLock()
	defer fake.findOrCreateResourceConfigForResourceMutex.RUnlock()
	return fake.findOrCreateResourceConfigForResourceArgsForCall[i].resource, fake.findOrCreateResourceConfigForResourceArgsForCall[i].resourceType, fake.findOrCreateResourceConfigForResourceArgsForCall[i].source, fake.findOrCreateResourceConfigForResourceArgsForCall[i].pipeline, fake.findOrCreateResourceConfigForResourceArgsForCall[i].resourceTypes
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForResourceReturns(result1 *dbng.UsedResourceConfig, result2 error) {
	fake.FindOrCreateResourceConfigForResourceStub = nil
	fake.findOrCreateResourceConfigForResourceReturns = struct {
		result1 *dbng.UsedResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForResourceType(resourceTypeName string, source atc.Source, pipeline *dbng.Pipeline, resourceTypes atc.ResourceTypes) (*dbng.UsedResourceConfig, error) {
	fake.findOrCreateResourceConfigForResourceTypeMutex.Lock()
	fake.findOrCreateResourceConfigForResourceTypeArgsForCall = append(fake.findOrCreateResourceConfigForResourceTypeArgsForCall, struct {
		resourceTypeName string
		source           atc.Source
		pipeline         *dbng.Pipeline
		resourceTypes    atc.ResourceTypes
	}{resourceTypeName, source, pipeline, resourceTypes})
	fake.recordInvocation("FindOrCreateResourceConfigForResourceType", []interface{}{resourceTypeName, source, pipeline, resourceTypes})
	fake.findOrCreateResourceConfigForResourceTypeMutex.Unlock()
	if fake.FindOrCreateResourceConfigForResourceTypeStub != nil {
		return fake.FindOrCreateResourceConfigForResourceTypeStub(resourceTypeName, source, pipeline, resourceTypes)
	} else {
		return fake.findOrCreateResourceConfigForResourceTypeReturns.result1, fake.findOrCreateResourceConfigForResourceTypeReturns.result2
	}
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForResourceTypeCallCount() int {
	fake.findOrCreateResourceConfigForResourceTypeMutex.RLock()
	defer fake.findOrCreateResourceConfigForResourceTypeMutex.RUnlock()
	return len(fake.findOrCreateResourceConfigForResourceTypeArgsForCall)
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForResourceTypeArgsForCall(i int) (string, atc.Source, *dbng.Pipeline, atc.ResourceTypes) {
	fake.findOrCreateResourceConfigForResourceTypeMutex.RLock()
	defer fake.findOrCreateResourceConfigForResourceTypeMutex.RUnlock()
	return fake.findOrCreateResourceConfigForResourceTypeArgsForCall[i].resourceTypeName, fake.findOrCreateResourceConfigForResourceTypeArgsForCall[i].source, fake.findOrCreateResourceConfigForResourceTypeArgsForCall[i].pipeline, fake.findOrCreateResourceConfigForResourceTypeArgsForCall[i].resourceTypes
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigForResourceTypeReturns(result1 *dbng.UsedResourceConfig, result2 error) {
	fake.FindOrCreateResourceConfigForResourceTypeStub = nil
	fake.findOrCreateResourceConfigForResourceTypeReturns = struct {
		result1 *dbng.UsedResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateResourceConfigForBuildMutex.RLock()
	defer fake.findOrCreateResourceConfigForBuildMutex.RUnlock()
	fake.findOrCreateResourceConfigForResourceMutex.RLock()
	defer fake.findOrCreateResourceConfigForResourceMutex.RUnlock()
	fake.findOrCreateResourceConfigForResourceTypeMutex.RLock()
	defer fake.findOrCreateResourceConfigForResourceTypeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResourceConfigFactory) recordInvocation(key string, args []interface{}) {
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

var _ dbng.ResourceConfigFactory = new(FakeResourceConfigFactory)
