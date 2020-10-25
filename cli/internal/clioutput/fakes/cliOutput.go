// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/opctl/opctl/cli/internal/clioutput"
	"github.com/opctl/opctl/sdks/go/model"
)

type FakeCliOutput struct {
	AttentionStub        func(string)
	attentionMutex       sync.RWMutex
	attentionArgsForCall []struct {
		arg1 string
	}
	ErrorStub        func(string)
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
		arg1 string
	}
	EventStub        func(*model.Event)
	eventMutex       sync.RWMutex
	eventArgsForCall []struct {
		arg1 *model.Event
	}
	InfoStub        func(string)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		arg1 string
	}
	SuccessStub        func(string)
	successMutex       sync.RWMutex
	successArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCliOutput) Attention(arg1 string) {
	fake.attentionMutex.Lock()
	fake.attentionArgsForCall = append(fake.attentionArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Attention", []interface{}{arg1})
	fake.attentionMutex.Unlock()
	if fake.AttentionStub != nil {
		fake.AttentionStub(arg1)
	}
}

func (fake *FakeCliOutput) AttentionCallCount() int {
	fake.attentionMutex.RLock()
	defer fake.attentionMutex.RUnlock()
	return len(fake.attentionArgsForCall)
}

func (fake *FakeCliOutput) AttentionCalls(stub func(string)) {
	fake.attentionMutex.Lock()
	defer fake.attentionMutex.Unlock()
	fake.AttentionStub = stub
}

func (fake *FakeCliOutput) AttentionArgsForCall(i int) string {
	fake.attentionMutex.RLock()
	defer fake.attentionMutex.RUnlock()
	argsForCall := fake.attentionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCliOutput) Error(arg1 string) {
	fake.errorMutex.Lock()
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Error", []interface{}{arg1})
	fake.errorMutex.Unlock()
	if fake.ErrorStub != nil {
		fake.ErrorStub(arg1)
	}
}

func (fake *FakeCliOutput) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeCliOutput) ErrorCalls(stub func(string)) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = stub
}

func (fake *FakeCliOutput) ErrorArgsForCall(i int) string {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	argsForCall := fake.errorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCliOutput) Event(arg1 *model.Event) {
	fake.eventMutex.Lock()
	fake.eventArgsForCall = append(fake.eventArgsForCall, struct {
		arg1 *model.Event
	}{arg1})
	fake.recordInvocation("Event", []interface{}{arg1})
	fake.eventMutex.Unlock()
	if fake.EventStub != nil {
		fake.EventStub(arg1)
	}
}

func (fake *FakeCliOutput) EventCallCount() int {
	fake.eventMutex.RLock()
	defer fake.eventMutex.RUnlock()
	return len(fake.eventArgsForCall)
}

func (fake *FakeCliOutput) EventCalls(stub func(*model.Event)) {
	fake.eventMutex.Lock()
	defer fake.eventMutex.Unlock()
	fake.EventStub = stub
}

func (fake *FakeCliOutput) EventArgsForCall(i int) *model.Event {
	fake.eventMutex.RLock()
	defer fake.eventMutex.RUnlock()
	argsForCall := fake.eventArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCliOutput) Info(arg1 string) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Info", []interface{}{arg1})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		fake.InfoStub(arg1)
	}
}

func (fake *FakeCliOutput) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeCliOutput) InfoCalls(stub func(string)) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = stub
}

func (fake *FakeCliOutput) InfoArgsForCall(i int) string {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	argsForCall := fake.infoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCliOutput) Success(arg1 string) {
	fake.successMutex.Lock()
	fake.successArgsForCall = append(fake.successArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Success", []interface{}{arg1})
	fake.successMutex.Unlock()
	if fake.SuccessStub != nil {
		fake.SuccessStub(arg1)
	}
}

func (fake *FakeCliOutput) SuccessCallCount() int {
	fake.successMutex.RLock()
	defer fake.successMutex.RUnlock()
	return len(fake.successArgsForCall)
}

func (fake *FakeCliOutput) SuccessCalls(stub func(string)) {
	fake.successMutex.Lock()
	defer fake.successMutex.Unlock()
	fake.SuccessStub = stub
}

func (fake *FakeCliOutput) SuccessArgsForCall(i int) string {
	fake.successMutex.RLock()
	defer fake.successMutex.RUnlock()
	argsForCall := fake.successArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCliOutput) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attentionMutex.RLock()
	defer fake.attentionMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	fake.eventMutex.RLock()
	defer fake.eventMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.successMutex.RLock()
	defer fake.successMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCliOutput) recordInvocation(key string, args []interface{}) {
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

var _ clioutput.CliOutput = new(FakeCliOutput)