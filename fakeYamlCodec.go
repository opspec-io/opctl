// This file was generated by counterfeiter
package opspec

import "sync"

type fakeYamlCodec struct {
	ToYamlStub        func(in interface{}) (opFileBytes []byte, err error)
	toYamlMutex       sync.RWMutex
	toYamlArgsForCall []struct {
		in interface{}
	}
	toYamlReturns struct {
		result1 []byte
		result2 error
	}
	FromYamlStub        func(in []byte, out interface{}) (err error)
	fromYamlMutex       sync.RWMutex
	fromYamlArgsForCall []struct {
		in  []byte
		out interface{}
	}
	fromYamlReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeYamlCodec) ToYaml(in interface{}) (opFileBytes []byte, err error) {
	fake.toYamlMutex.Lock()
	fake.toYamlArgsForCall = append(fake.toYamlArgsForCall, struct {
		in interface{}
	}{in})
	fake.recordInvocation("ToYaml", []interface{}{in})
	fake.toYamlMutex.Unlock()
	if fake.ToYamlStub != nil {
		return fake.ToYamlStub(in)
	} else {
		return fake.toYamlReturns.result1, fake.toYamlReturns.result2
	}
}

func (fake *fakeYamlCodec) ToYamlCallCount() int {
	fake.toYamlMutex.RLock()
	defer fake.toYamlMutex.RUnlock()
	return len(fake.toYamlArgsForCall)
}

func (fake *fakeYamlCodec) ToYamlArgsForCall(i int) interface{} {
	fake.toYamlMutex.RLock()
	defer fake.toYamlMutex.RUnlock()
	return fake.toYamlArgsForCall[i].in
}

func (fake *fakeYamlCodec) ToYamlReturns(result1 []byte, result2 error) {
	fake.ToYamlStub = nil
	fake.toYamlReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *fakeYamlCodec) FromYaml(in []byte, out interface{}) (err error) {
	var inCopy []byte
	if in != nil {
		inCopy = make([]byte, len(in))
		copy(inCopy, in)
	}
	fake.fromYamlMutex.Lock()
	fake.fromYamlArgsForCall = append(fake.fromYamlArgsForCall, struct {
		in  []byte
		out interface{}
	}{inCopy, out})
	fake.recordInvocation("FromYaml", []interface{}{inCopy, out})
	fake.fromYamlMutex.Unlock()
	if fake.FromYamlStub != nil {
		return fake.FromYamlStub(in, out)
	} else {
		return fake.fromYamlReturns.result1
	}
}

func (fake *fakeYamlCodec) FromYamlCallCount() int {
	fake.fromYamlMutex.RLock()
	defer fake.fromYamlMutex.RUnlock()
	return len(fake.fromYamlArgsForCall)
}

func (fake *fakeYamlCodec) FromYamlArgsForCall(i int) ([]byte, interface{}) {
	fake.fromYamlMutex.RLock()
	defer fake.fromYamlMutex.RUnlock()
	return fake.fromYamlArgsForCall[i].in, fake.fromYamlArgsForCall[i].out
}

func (fake *fakeYamlCodec) FromYamlReturns(result1 error) {
	fake.FromYamlStub = nil
	fake.fromYamlReturns = struct {
		result1 error
	}{result1}
}

func (fake *fakeYamlCodec) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.toYamlMutex.RLock()
	defer fake.toYamlMutex.RUnlock()
	fake.fromYamlMutex.RLock()
	defer fake.fromYamlMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeYamlCodec) recordInvocation(key string, args []interface{}) {
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
