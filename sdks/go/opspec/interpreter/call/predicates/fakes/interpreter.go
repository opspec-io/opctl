// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/predicates"
)

type FakeInterpreter struct {
	InterpretStub        func([]*model.PredicateSpec, map[string]*model.Value) (bool, error)
	interpretMutex       sync.RWMutex
	interpretArgsForCall []struct {
		arg1 []*model.PredicateSpec
		arg2 map[string]*model.Value
	}
	interpretReturns struct {
		result1 bool
		result2 error
	}
	interpretReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterpreter) Interpret(arg1 []*model.PredicateSpec, arg2 map[string]*model.Value) (bool, error) {
	var arg1Copy []*model.PredicateSpec
	if arg1 != nil {
		arg1Copy = make([]*model.PredicateSpec, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.interpretMutex.Lock()
	ret, specificReturn := fake.interpretReturnsOnCall[len(fake.interpretArgsForCall)]
	fake.interpretArgsForCall = append(fake.interpretArgsForCall, struct {
		arg1 []*model.PredicateSpec
		arg2 map[string]*model.Value
	}{arg1Copy, arg2})
	fake.recordInvocation("Interpret", []interface{}{arg1Copy, arg2})
	fake.interpretMutex.Unlock()
	if fake.InterpretStub != nil {
		return fake.InterpretStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.interpretReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInterpreter) InterpretCallCount() int {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return len(fake.interpretArgsForCall)
}

func (fake *FakeInterpreter) InterpretCalls(stub func([]*model.PredicateSpec, map[string]*model.Value) (bool, error)) {
	fake.interpretMutex.Lock()
	defer fake.interpretMutex.Unlock()
	fake.InterpretStub = stub
}

func (fake *FakeInterpreter) InterpretArgsForCall(i int) ([]*model.PredicateSpec, map[string]*model.Value) {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	argsForCall := fake.interpretArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInterpreter) InterpretReturns(result1 bool, result2 error) {
	fake.interpretMutex.Lock()
	defer fake.interpretMutex.Unlock()
	fake.InterpretStub = nil
	fake.interpretReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) InterpretReturnsOnCall(i int, result1 bool, result2 error) {
	fake.interpretMutex.Lock()
	defer fake.interpretMutex.Unlock()
	fake.InterpretStub = nil
	if fake.interpretReturnsOnCall == nil {
		fake.interpretReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.interpretReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterpreter) recordInvocation(key string, args []interface{}) {
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

var _ predicates.Interpreter = new(FakeInterpreter)
