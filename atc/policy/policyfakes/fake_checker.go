// Code generated by counterfeiter. DO NOT EDIT.
package policyfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/policy"
)

type FakeChecker struct {
	CheckStub        func(policy.PolicyCheckInput) (policy.PolicyCheckOutput, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 policy.PolicyCheckInput
	}
	checkReturns struct {
		result1 policy.PolicyCheckOutput
		result2 error
	}
	checkReturnsOnCall map[int]struct {
		result1 policy.PolicyCheckOutput
		result2 error
	}
	ShouldCheckActionStub        func(string) bool
	shouldCheckActionMutex       sync.RWMutex
	shouldCheckActionArgsForCall []struct {
		arg1 string
	}
	shouldCheckActionReturns struct {
		result1 bool
	}
	shouldCheckActionReturnsOnCall map[int]struct {
		result1 bool
	}
	ShouldCheckHttpMethodStub        func(string) bool
	shouldCheckHttpMethodMutex       sync.RWMutex
	shouldCheckHttpMethodArgsForCall []struct {
		arg1 string
	}
	shouldCheckHttpMethodReturns struct {
		result1 bool
	}
	shouldCheckHttpMethodReturnsOnCall map[int]struct {
		result1 bool
	}
	ShouldSkipActionStub        func(string) bool
	shouldSkipActionMutex       sync.RWMutex
	shouldSkipActionArgsForCall []struct {
		arg1 string
	}
	shouldSkipActionReturns struct {
		result1 bool
	}
	shouldSkipActionReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeChecker) Check(arg1 policy.PolicyCheckInput) (policy.PolicyCheckOutput, error) {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 policy.PolicyCheckInput
	}{arg1})
	stub := fake.CheckStub
	fakeReturns := fake.checkReturns
	fake.recordInvocation("Check", []interface{}{arg1})
	fake.checkMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeChecker) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeChecker) CheckCalls(stub func(policy.PolicyCheckInput) (policy.PolicyCheckOutput, error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeChecker) CheckArgsForCall(i int) policy.PolicyCheckInput {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeChecker) CheckReturns(result1 policy.PolicyCheckOutput, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 policy.PolicyCheckOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeChecker) CheckReturnsOnCall(i int, result1 policy.PolicyCheckOutput, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 policy.PolicyCheckOutput
			result2 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 policy.PolicyCheckOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeChecker) ShouldCheckAction(arg1 string) bool {
	fake.shouldCheckActionMutex.Lock()
	ret, specificReturn := fake.shouldCheckActionReturnsOnCall[len(fake.shouldCheckActionArgsForCall)]
	fake.shouldCheckActionArgsForCall = append(fake.shouldCheckActionArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ShouldCheckActionStub
	fakeReturns := fake.shouldCheckActionReturns
	fake.recordInvocation("ShouldCheckAction", []interface{}{arg1})
	fake.shouldCheckActionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeChecker) ShouldCheckActionCallCount() int {
	fake.shouldCheckActionMutex.RLock()
	defer fake.shouldCheckActionMutex.RUnlock()
	return len(fake.shouldCheckActionArgsForCall)
}

func (fake *FakeChecker) ShouldCheckActionCalls(stub func(string) bool) {
	fake.shouldCheckActionMutex.Lock()
	defer fake.shouldCheckActionMutex.Unlock()
	fake.ShouldCheckActionStub = stub
}

func (fake *FakeChecker) ShouldCheckActionArgsForCall(i int) string {
	fake.shouldCheckActionMutex.RLock()
	defer fake.shouldCheckActionMutex.RUnlock()
	argsForCall := fake.shouldCheckActionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeChecker) ShouldCheckActionReturns(result1 bool) {
	fake.shouldCheckActionMutex.Lock()
	defer fake.shouldCheckActionMutex.Unlock()
	fake.ShouldCheckActionStub = nil
	fake.shouldCheckActionReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeChecker) ShouldCheckActionReturnsOnCall(i int, result1 bool) {
	fake.shouldCheckActionMutex.Lock()
	defer fake.shouldCheckActionMutex.Unlock()
	fake.ShouldCheckActionStub = nil
	if fake.shouldCheckActionReturnsOnCall == nil {
		fake.shouldCheckActionReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.shouldCheckActionReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeChecker) ShouldCheckHttpMethod(arg1 string) bool {
	fake.shouldCheckHttpMethodMutex.Lock()
	ret, specificReturn := fake.shouldCheckHttpMethodReturnsOnCall[len(fake.shouldCheckHttpMethodArgsForCall)]
	fake.shouldCheckHttpMethodArgsForCall = append(fake.shouldCheckHttpMethodArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ShouldCheckHttpMethodStub
	fakeReturns := fake.shouldCheckHttpMethodReturns
	fake.recordInvocation("ShouldCheckHttpMethod", []interface{}{arg1})
	fake.shouldCheckHttpMethodMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeChecker) ShouldCheckHttpMethodCallCount() int {
	fake.shouldCheckHttpMethodMutex.RLock()
	defer fake.shouldCheckHttpMethodMutex.RUnlock()
	return len(fake.shouldCheckHttpMethodArgsForCall)
}

func (fake *FakeChecker) ShouldCheckHttpMethodCalls(stub func(string) bool) {
	fake.shouldCheckHttpMethodMutex.Lock()
	defer fake.shouldCheckHttpMethodMutex.Unlock()
	fake.ShouldCheckHttpMethodStub = stub
}

func (fake *FakeChecker) ShouldCheckHttpMethodArgsForCall(i int) string {
	fake.shouldCheckHttpMethodMutex.RLock()
	defer fake.shouldCheckHttpMethodMutex.RUnlock()
	argsForCall := fake.shouldCheckHttpMethodArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeChecker) ShouldCheckHttpMethodReturns(result1 bool) {
	fake.shouldCheckHttpMethodMutex.Lock()
	defer fake.shouldCheckHttpMethodMutex.Unlock()
	fake.ShouldCheckHttpMethodStub = nil
	fake.shouldCheckHttpMethodReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeChecker) ShouldCheckHttpMethodReturnsOnCall(i int, result1 bool) {
	fake.shouldCheckHttpMethodMutex.Lock()
	defer fake.shouldCheckHttpMethodMutex.Unlock()
	fake.ShouldCheckHttpMethodStub = nil
	if fake.shouldCheckHttpMethodReturnsOnCall == nil {
		fake.shouldCheckHttpMethodReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.shouldCheckHttpMethodReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeChecker) ShouldSkipAction(arg1 string) bool {
	fake.shouldSkipActionMutex.Lock()
	ret, specificReturn := fake.shouldSkipActionReturnsOnCall[len(fake.shouldSkipActionArgsForCall)]
	fake.shouldSkipActionArgsForCall = append(fake.shouldSkipActionArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ShouldSkipActionStub
	fakeReturns := fake.shouldSkipActionReturns
	fake.recordInvocation("ShouldSkipAction", []interface{}{arg1})
	fake.shouldSkipActionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeChecker) ShouldSkipActionCallCount() int {
	fake.shouldSkipActionMutex.RLock()
	defer fake.shouldSkipActionMutex.RUnlock()
	return len(fake.shouldSkipActionArgsForCall)
}

func (fake *FakeChecker) ShouldSkipActionCalls(stub func(string) bool) {
	fake.shouldSkipActionMutex.Lock()
	defer fake.shouldSkipActionMutex.Unlock()
	fake.ShouldSkipActionStub = stub
}

func (fake *FakeChecker) ShouldSkipActionArgsForCall(i int) string {
	fake.shouldSkipActionMutex.RLock()
	defer fake.shouldSkipActionMutex.RUnlock()
	argsForCall := fake.shouldSkipActionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeChecker) ShouldSkipActionReturns(result1 bool) {
	fake.shouldSkipActionMutex.Lock()
	defer fake.shouldSkipActionMutex.Unlock()
	fake.ShouldSkipActionStub = nil
	fake.shouldSkipActionReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeChecker) ShouldSkipActionReturnsOnCall(i int, result1 bool) {
	fake.shouldSkipActionMutex.Lock()
	defer fake.shouldSkipActionMutex.Unlock()
	fake.ShouldSkipActionStub = nil
	if fake.shouldSkipActionReturnsOnCall == nil {
		fake.shouldSkipActionReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.shouldSkipActionReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.shouldCheckActionMutex.RLock()
	defer fake.shouldCheckActionMutex.RUnlock()
	fake.shouldCheckHttpMethodMutex.RLock()
	defer fake.shouldCheckHttpMethodMutex.RUnlock()
	fake.shouldSkipActionMutex.RLock()
	defer fake.shouldSkipActionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeChecker) recordInvocation(key string, args []interface{}) {
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

var _ policy.Checker = new(FakeChecker)
