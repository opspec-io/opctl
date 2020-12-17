// Code generated by counterfeiter. DO NOT EDIT.
package dataresolver

import (
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
)

type Fake struct {
	ResolveStub        func(string, *model.Creds) model.DataHandle
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		arg1 string
		arg2 *model.Creds
	}
	resolveReturns struct {
		result1 model.DataHandle
	}
	resolveReturnsOnCall map[int]struct {
		result1 model.DataHandle
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Resolve(arg1 string, arg2 *model.Creds) model.DataHandle {
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		arg1 string
		arg2 *model.Creds
	}{arg1, arg2})
	fake.recordInvocation("Resolve", []interface{}{arg1, arg2})
	fake.resolveMutex.Unlock()
	if fake.ResolveStub != nil {
		return fake.ResolveStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.resolveReturns
	return fakeReturns.result1
}

func (fake *Fake) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *Fake) ResolveCalls(stub func(string, *model.Creds) model.DataHandle) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = stub
}

func (fake *Fake) ResolveArgsForCall(i int) (string, *model.Creds) {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	argsForCall := fake.resolveArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Fake) ResolveReturns(result1 model.DataHandle) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 model.DataHandle
	}{result1}
}

func (fake *Fake) ResolveReturnsOnCall(i int, result1 model.DataHandle) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 model.DataHandle
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 model.DataHandle
	}{result1}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ DataResolver = new(Fake)