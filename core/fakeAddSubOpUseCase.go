// This file was generated by counterfeiter
package core

import (
	"sync"

	"github.com/opctl/engine/core/models"
)

type fakeAddSubOpUseCase struct {
	ExecuteStub        func(req models.AddSubOpReq) (err error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		req models.AddSubOpReq
	}
	executeReturns struct {
		result1 error
	}
}

func (fake *fakeAddSubOpUseCase) Execute(req models.AddSubOpReq) (err error) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		req models.AddSubOpReq
	}{req})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(req)
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *fakeAddSubOpUseCase) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *fakeAddSubOpUseCase) ExecuteArgsForCall(i int) models.AddSubOpReq {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].req
}

func (fake *fakeAddSubOpUseCase) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}