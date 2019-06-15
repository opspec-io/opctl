// Code generated by counterfeiter. DO NOT EDIT.
package dirs

import (
	"sync"

	"github.com/opctl/sdk-golang/model"
)

type FakeInterpreter struct {
	InterpretStub        func(opHandle model.DataHandle, scope map[string]*model.Value, scgContainerCallFiles map[string]string, scratchDirPath string) (map[string]string, error)
	interpretMutex       sync.RWMutex
	interpretArgsForCall []struct {
		opHandle              model.DataHandle
		scope                 map[string]*model.Value
		scgContainerCallFiles map[string]string
		scratchDirPath        string
	}
	interpretReturns struct {
		result1 map[string]string
		result2 error
	}
	interpretReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterpreter) Interpret(opHandle model.DataHandle, scope map[string]*model.Value, scgContainerCallFiles map[string]string, scratchDirPath string) (map[string]string, error) {
	fake.interpretMutex.Lock()
	ret, specificReturn := fake.interpretReturnsOnCall[len(fake.interpretArgsForCall)]
	fake.interpretArgsForCall = append(fake.interpretArgsForCall, struct {
		opHandle              model.DataHandle
		scope                 map[string]*model.Value
		scgContainerCallFiles map[string]string
		scratchDirPath        string
	}{opHandle, scope, scgContainerCallFiles, scratchDirPath})
	fake.recordInvocation("Interpret", []interface{}{opHandle, scope, scgContainerCallFiles, scratchDirPath})
	fake.interpretMutex.Unlock()
	if fake.InterpretStub != nil {
		return fake.InterpretStub(opHandle, scope, scgContainerCallFiles, scratchDirPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interpretReturns.result1, fake.interpretReturns.result2
}

func (fake *FakeInterpreter) InterpretCallCount() int {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return len(fake.interpretArgsForCall)
}

func (fake *FakeInterpreter) InterpretArgsForCall(i int) (model.DataHandle, map[string]*model.Value, map[string]string, string) {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return fake.interpretArgsForCall[i].opHandle, fake.interpretArgsForCall[i].scope, fake.interpretArgsForCall[i].scgContainerCallFiles, fake.interpretArgsForCall[i].scratchDirPath
}

func (fake *FakeInterpreter) InterpretReturns(result1 map[string]string, result2 error) {
	fake.InterpretStub = nil
	fake.interpretReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpreter) InterpretReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.InterpretStub = nil
	if fake.interpretReturnsOnCall == nil {
		fake.interpretReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.interpretReturnsOnCall[i] = struct {
		result1 map[string]string
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

var _ Interpreter = new(FakeInterpreter)
