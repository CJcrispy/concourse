// Code generated by counterfeiter. DO NOT EDIT.
package lockfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/db/lock"
)

type FakeLockFactory struct {
	AcquireStub        func(lager.Logger, lock.LockID) (lock.Lock, bool, error)
	acquireMutex       sync.RWMutex
	acquireArgsForCall []struct {
		arg1 lager.Logger
		arg2 lock.LockID
	}
	acquireReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	acquireReturnsOnCall map[int]struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLockFactory) Acquire(arg1 lager.Logger, arg2 lock.LockID) (lock.Lock, bool, error) {
	fake.acquireMutex.Lock()
	ret, specificReturn := fake.acquireReturnsOnCall[len(fake.acquireArgsForCall)]
	fake.acquireArgsForCall = append(fake.acquireArgsForCall, struct {
		arg1 lager.Logger
		arg2 lock.LockID
	}{arg1, arg2})
	stub := fake.AcquireStub
	fakeReturns := fake.acquireReturns
	fake.recordInvocation("Acquire", []interface{}{arg1, arg2})
	fake.acquireMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeLockFactory) AcquireCallCount() int {
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	return len(fake.acquireArgsForCall)
}

func (fake *FakeLockFactory) AcquireCalls(stub func(lager.Logger, lock.LockID) (lock.Lock, bool, error)) {
	fake.acquireMutex.Lock()
	defer fake.acquireMutex.Unlock()
	fake.AcquireStub = stub
}

func (fake *FakeLockFactory) AcquireArgsForCall(i int) (lager.Logger, lock.LockID) {
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	argsForCall := fake.acquireArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLockFactory) AcquireReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.acquireMutex.Lock()
	defer fake.acquireMutex.Unlock()
	fake.AcquireStub = nil
	fake.acquireReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLockFactory) AcquireReturnsOnCall(i int, result1 lock.Lock, result2 bool, result3 error) {
	fake.acquireMutex.Lock()
	defer fake.acquireMutex.Unlock()
	fake.AcquireStub = nil
	if fake.acquireReturnsOnCall == nil {
		fake.acquireReturnsOnCall = make(map[int]struct {
			result1 lock.Lock
			result2 bool
			result3 error
		})
	}
	fake.acquireReturnsOnCall[i] = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLockFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLockFactory) recordInvocation(key string, args []interface{}) {
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

var _ lock.LockFactory = new(FakeLockFactory)
