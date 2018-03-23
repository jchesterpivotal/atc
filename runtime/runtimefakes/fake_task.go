// Code generated by counterfeiter. DO NOT EDIT.
package runtimefakes

import (
	"context"
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/runtime"
	"github.com/concourse/atc/worker"
)

type FakeTask struct {
	RunStub        func(context.Context, runtime.IOConfig, atc.TaskConfig) (chan runtime.TaskResult, []worker.VolumeMount, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 context.Context
		arg2 runtime.IOConfig
		arg3 atc.TaskConfig
	}
	runReturns struct {
		result1 chan runtime.TaskResult
		result2 []worker.VolumeMount
		result3 error
	}
	runReturnsOnCall map[int]struct {
		result1 chan runtime.TaskResult
		result2 []worker.VolumeMount
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTask) Run(arg1 context.Context, arg2 runtime.IOConfig, arg3 atc.TaskConfig) (chan runtime.TaskResult, []worker.VolumeMount, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 context.Context
		arg2 runtime.IOConfig
		arg3 atc.TaskConfig
	}{arg1, arg2, arg3})
	fake.recordInvocation("Run", []interface{}{arg1, arg2, arg3})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.runReturns.result1, fake.runReturns.result2, fake.runReturns.result3
}

func (fake *FakeTask) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeTask) RunArgsForCall(i int) (context.Context, runtime.IOConfig, atc.TaskConfig) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1, fake.runArgsForCall[i].arg2, fake.runArgsForCall[i].arg3
}

func (fake *FakeTask) RunReturns(result1 chan runtime.TaskResult, result2 []worker.VolumeMount, result3 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 chan runtime.TaskResult
		result2 []worker.VolumeMount
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTask) RunReturnsOnCall(i int, result1 chan runtime.TaskResult, result2 []worker.VolumeMount, result3 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 chan runtime.TaskResult
			result2 []worker.VolumeMount
			result3 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 chan runtime.TaskResult
		result2 []worker.VolumeMount
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTask) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTask) recordInvocation(key string, args []interface{}) {
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

var _ runtime.Task = new(FakeTask)