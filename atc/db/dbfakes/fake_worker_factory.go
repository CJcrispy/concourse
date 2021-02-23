// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"
	"time"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
)

type FakeWorkerFactory struct {
	BuildContainersCountPerWorkerStub        func() (map[string]int, error)
	buildContainersCountPerWorkerMutex       sync.RWMutex
	buildContainersCountPerWorkerArgsForCall []struct {
	}
	buildContainersCountPerWorkerReturns struct {
		result1 map[string]int
		result2 error
	}
	buildContainersCountPerWorkerReturnsOnCall map[int]struct {
		result1 map[string]int
		result2 error
	}
	FindWorkersForContainerByOwnerStub        func(db.ContainerOwner) ([]db.Worker, error)
	findWorkersForContainerByOwnerMutex       sync.RWMutex
	findWorkersForContainerByOwnerArgsForCall []struct {
		arg1 db.ContainerOwner
	}
	findWorkersForContainerByOwnerReturns struct {
		result1 []db.Worker
		result2 error
	}
	findWorkersForContainerByOwnerReturnsOnCall map[int]struct {
		result1 []db.Worker
		result2 error
	}
	GetWorkerStub        func(string) (db.Worker, bool, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		arg1 string
	}
	getWorkerReturns struct {
		result1 db.Worker
		result2 bool
		result3 error
	}
	getWorkerReturnsOnCall map[int]struct {
		result1 db.Worker
		result2 bool
		result3 error
	}
	HeartbeatWorkerStub        func(atc.Worker, time.Duration) (db.Worker, error)
	heartbeatWorkerMutex       sync.RWMutex
	heartbeatWorkerArgsForCall []struct {
		arg1 atc.Worker
		arg2 time.Duration
	}
	heartbeatWorkerReturns struct {
		result1 db.Worker
		result2 error
	}
	heartbeatWorkerReturnsOnCall map[int]struct {
		result1 db.Worker
		result2 error
	}
	SaveWorkerStub        func(atc.Worker, time.Duration) (db.Worker, error)
	saveWorkerMutex       sync.RWMutex
	saveWorkerArgsForCall []struct {
		arg1 atc.Worker
		arg2 time.Duration
	}
	saveWorkerReturns struct {
		result1 db.Worker
		result2 error
	}
	saveWorkerReturnsOnCall map[int]struct {
		result1 db.Worker
		result2 error
	}
	VisibleWorkersStub        func([]string) ([]db.Worker, error)
	visibleWorkersMutex       sync.RWMutex
	visibleWorkersArgsForCall []struct {
		arg1 []string
	}
	visibleWorkersReturns struct {
		result1 []db.Worker
		result2 error
	}
	visibleWorkersReturnsOnCall map[int]struct {
		result1 []db.Worker
		result2 error
	}
	WorkersStub        func() ([]db.Worker, error)
	workersMutex       sync.RWMutex
	workersArgsForCall []struct {
	}
	workersReturns struct {
		result1 []db.Worker
		result2 error
	}
	workersReturnsOnCall map[int]struct {
		result1 []db.Worker
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkerFactory) BuildContainersCountPerWorker() (map[string]int, error) {
	fake.buildContainersCountPerWorkerMutex.Lock()
	ret, specificReturn := fake.buildContainersCountPerWorkerReturnsOnCall[len(fake.buildContainersCountPerWorkerArgsForCall)]
	fake.buildContainersCountPerWorkerArgsForCall = append(fake.buildContainersCountPerWorkerArgsForCall, struct {
	}{})
	stub := fake.BuildContainersCountPerWorkerStub
	fakeReturns := fake.buildContainersCountPerWorkerReturns
	fake.recordInvocation("BuildContainersCountPerWorker", []interface{}{})
	fake.buildContainersCountPerWorkerMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerFactory) BuildContainersCountPerWorkerCallCount() int {
	fake.buildContainersCountPerWorkerMutex.RLock()
	defer fake.buildContainersCountPerWorkerMutex.RUnlock()
	return len(fake.buildContainersCountPerWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) BuildContainersCountPerWorkerCalls(stub func() (map[string]int, error)) {
	fake.buildContainersCountPerWorkerMutex.Lock()
	defer fake.buildContainersCountPerWorkerMutex.Unlock()
	fake.BuildContainersCountPerWorkerStub = stub
}

func (fake *FakeWorkerFactory) BuildContainersCountPerWorkerReturns(result1 map[string]int, result2 error) {
	fake.buildContainersCountPerWorkerMutex.Lock()
	defer fake.buildContainersCountPerWorkerMutex.Unlock()
	fake.BuildContainersCountPerWorkerStub = nil
	fake.buildContainersCountPerWorkerReturns = struct {
		result1 map[string]int
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) BuildContainersCountPerWorkerReturnsOnCall(i int, result1 map[string]int, result2 error) {
	fake.buildContainersCountPerWorkerMutex.Lock()
	defer fake.buildContainersCountPerWorkerMutex.Unlock()
	fake.BuildContainersCountPerWorkerStub = nil
	if fake.buildContainersCountPerWorkerReturnsOnCall == nil {
		fake.buildContainersCountPerWorkerReturnsOnCall = make(map[int]struct {
			result1 map[string]int
			result2 error
		})
	}
	fake.buildContainersCountPerWorkerReturnsOnCall[i] = struct {
		result1 map[string]int
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) FindWorkersForContainerByOwner(arg1 db.ContainerOwner) ([]db.Worker, error) {
	fake.findWorkersForContainerByOwnerMutex.Lock()
	ret, specificReturn := fake.findWorkersForContainerByOwnerReturnsOnCall[len(fake.findWorkersForContainerByOwnerArgsForCall)]
	fake.findWorkersForContainerByOwnerArgsForCall = append(fake.findWorkersForContainerByOwnerArgsForCall, struct {
		arg1 db.ContainerOwner
	}{arg1})
	stub := fake.FindWorkersForContainerByOwnerStub
	fakeReturns := fake.findWorkersForContainerByOwnerReturns
	fake.recordInvocation("FindWorkersForContainerByOwner", []interface{}{arg1})
	fake.findWorkersForContainerByOwnerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerFactory) FindWorkersForContainerByOwnerCallCount() int {
	fake.findWorkersForContainerByOwnerMutex.RLock()
	defer fake.findWorkersForContainerByOwnerMutex.RUnlock()
	return len(fake.findWorkersForContainerByOwnerArgsForCall)
}

func (fake *FakeWorkerFactory) FindWorkersForContainerByOwnerCalls(stub func(db.ContainerOwner) ([]db.Worker, error)) {
	fake.findWorkersForContainerByOwnerMutex.Lock()
	defer fake.findWorkersForContainerByOwnerMutex.Unlock()
	fake.FindWorkersForContainerByOwnerStub = stub
}

func (fake *FakeWorkerFactory) FindWorkersForContainerByOwnerArgsForCall(i int) db.ContainerOwner {
	fake.findWorkersForContainerByOwnerMutex.RLock()
	defer fake.findWorkersForContainerByOwnerMutex.RUnlock()
	argsForCall := fake.findWorkersForContainerByOwnerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWorkerFactory) FindWorkersForContainerByOwnerReturns(result1 []db.Worker, result2 error) {
	fake.findWorkersForContainerByOwnerMutex.Lock()
	defer fake.findWorkersForContainerByOwnerMutex.Unlock()
	fake.FindWorkersForContainerByOwnerStub = nil
	fake.findWorkersForContainerByOwnerReturns = struct {
		result1 []db.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) FindWorkersForContainerByOwnerReturnsOnCall(i int, result1 []db.Worker, result2 error) {
	fake.findWorkersForContainerByOwnerMutex.Lock()
	defer fake.findWorkersForContainerByOwnerMutex.Unlock()
	fake.FindWorkersForContainerByOwnerStub = nil
	if fake.findWorkersForContainerByOwnerReturnsOnCall == nil {
		fake.findWorkersForContainerByOwnerReturnsOnCall = make(map[int]struct {
			result1 []db.Worker
			result2 error
		})
	}
	fake.findWorkersForContainerByOwnerReturnsOnCall[i] = struct {
		result1 []db.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) GetWorker(arg1 string) (db.Worker, bool, error) {
	fake.getWorkerMutex.Lock()
	ret, specificReturn := fake.getWorkerReturnsOnCall[len(fake.getWorkerArgsForCall)]
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetWorkerStub
	fakeReturns := fake.getWorkerReturns
	fake.recordInvocation("GetWorker", []interface{}{arg1})
	fake.getWorkerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeWorkerFactory) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) GetWorkerCalls(stub func(string) (db.Worker, bool, error)) {
	fake.getWorkerMutex.Lock()
	defer fake.getWorkerMutex.Unlock()
	fake.GetWorkerStub = stub
}

func (fake *FakeWorkerFactory) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	argsForCall := fake.getWorkerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWorkerFactory) GetWorkerReturns(result1 db.Worker, result2 bool, result3 error) {
	fake.getWorkerMutex.Lock()
	defer fake.getWorkerMutex.Unlock()
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 db.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerFactory) GetWorkerReturnsOnCall(i int, result1 db.Worker, result2 bool, result3 error) {
	fake.getWorkerMutex.Lock()
	defer fake.getWorkerMutex.Unlock()
	fake.GetWorkerStub = nil
	if fake.getWorkerReturnsOnCall == nil {
		fake.getWorkerReturnsOnCall = make(map[int]struct {
			result1 db.Worker
			result2 bool
			result3 error
		})
	}
	fake.getWorkerReturnsOnCall[i] = struct {
		result1 db.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerFactory) HeartbeatWorker(arg1 atc.Worker, arg2 time.Duration) (db.Worker, error) {
	fake.heartbeatWorkerMutex.Lock()
	ret, specificReturn := fake.heartbeatWorkerReturnsOnCall[len(fake.heartbeatWorkerArgsForCall)]
	fake.heartbeatWorkerArgsForCall = append(fake.heartbeatWorkerArgsForCall, struct {
		arg1 atc.Worker
		arg2 time.Duration
	}{arg1, arg2})
	stub := fake.HeartbeatWorkerStub
	fakeReturns := fake.heartbeatWorkerReturns
	fake.recordInvocation("HeartbeatWorker", []interface{}{arg1, arg2})
	fake.heartbeatWorkerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerFactory) HeartbeatWorkerCallCount() int {
	fake.heartbeatWorkerMutex.RLock()
	defer fake.heartbeatWorkerMutex.RUnlock()
	return len(fake.heartbeatWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) HeartbeatWorkerCalls(stub func(atc.Worker, time.Duration) (db.Worker, error)) {
	fake.heartbeatWorkerMutex.Lock()
	defer fake.heartbeatWorkerMutex.Unlock()
	fake.HeartbeatWorkerStub = stub
}

func (fake *FakeWorkerFactory) HeartbeatWorkerArgsForCall(i int) (atc.Worker, time.Duration) {
	fake.heartbeatWorkerMutex.RLock()
	defer fake.heartbeatWorkerMutex.RUnlock()
	argsForCall := fake.heartbeatWorkerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeWorkerFactory) HeartbeatWorkerReturns(result1 db.Worker, result2 error) {
	fake.heartbeatWorkerMutex.Lock()
	defer fake.heartbeatWorkerMutex.Unlock()
	fake.HeartbeatWorkerStub = nil
	fake.heartbeatWorkerReturns = struct {
		result1 db.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) HeartbeatWorkerReturnsOnCall(i int, result1 db.Worker, result2 error) {
	fake.heartbeatWorkerMutex.Lock()
	defer fake.heartbeatWorkerMutex.Unlock()
	fake.HeartbeatWorkerStub = nil
	if fake.heartbeatWorkerReturnsOnCall == nil {
		fake.heartbeatWorkerReturnsOnCall = make(map[int]struct {
			result1 db.Worker
			result2 error
		})
	}
	fake.heartbeatWorkerReturnsOnCall[i] = struct {
		result1 db.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) SaveWorker(arg1 atc.Worker, arg2 time.Duration) (db.Worker, error) {
	fake.saveWorkerMutex.Lock()
	ret, specificReturn := fake.saveWorkerReturnsOnCall[len(fake.saveWorkerArgsForCall)]
	fake.saveWorkerArgsForCall = append(fake.saveWorkerArgsForCall, struct {
		arg1 atc.Worker
		arg2 time.Duration
	}{arg1, arg2})
	stub := fake.SaveWorkerStub
	fakeReturns := fake.saveWorkerReturns
	fake.recordInvocation("SaveWorker", []interface{}{arg1, arg2})
	fake.saveWorkerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerFactory) SaveWorkerCallCount() int {
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	return len(fake.saveWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) SaveWorkerCalls(stub func(atc.Worker, time.Duration) (db.Worker, error)) {
	fake.saveWorkerMutex.Lock()
	defer fake.saveWorkerMutex.Unlock()
	fake.SaveWorkerStub = stub
}

func (fake *FakeWorkerFactory) SaveWorkerArgsForCall(i int) (atc.Worker, time.Duration) {
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	argsForCall := fake.saveWorkerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeWorkerFactory) SaveWorkerReturns(result1 db.Worker, result2 error) {
	fake.saveWorkerMutex.Lock()
	defer fake.saveWorkerMutex.Unlock()
	fake.SaveWorkerStub = nil
	fake.saveWorkerReturns = struct {
		result1 db.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) SaveWorkerReturnsOnCall(i int, result1 db.Worker, result2 error) {
	fake.saveWorkerMutex.Lock()
	defer fake.saveWorkerMutex.Unlock()
	fake.SaveWorkerStub = nil
	if fake.saveWorkerReturnsOnCall == nil {
		fake.saveWorkerReturnsOnCall = make(map[int]struct {
			result1 db.Worker
			result2 error
		})
	}
	fake.saveWorkerReturnsOnCall[i] = struct {
		result1 db.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) VisibleWorkers(arg1 []string) ([]db.Worker, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.visibleWorkersMutex.Lock()
	ret, specificReturn := fake.visibleWorkersReturnsOnCall[len(fake.visibleWorkersArgsForCall)]
	fake.visibleWorkersArgsForCall = append(fake.visibleWorkersArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.VisibleWorkersStub
	fakeReturns := fake.visibleWorkersReturns
	fake.recordInvocation("VisibleWorkers", []interface{}{arg1Copy})
	fake.visibleWorkersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerFactory) VisibleWorkersCallCount() int {
	fake.visibleWorkersMutex.RLock()
	defer fake.visibleWorkersMutex.RUnlock()
	return len(fake.visibleWorkersArgsForCall)
}

func (fake *FakeWorkerFactory) VisibleWorkersCalls(stub func([]string) ([]db.Worker, error)) {
	fake.visibleWorkersMutex.Lock()
	defer fake.visibleWorkersMutex.Unlock()
	fake.VisibleWorkersStub = stub
}

func (fake *FakeWorkerFactory) VisibleWorkersArgsForCall(i int) []string {
	fake.visibleWorkersMutex.RLock()
	defer fake.visibleWorkersMutex.RUnlock()
	argsForCall := fake.visibleWorkersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWorkerFactory) VisibleWorkersReturns(result1 []db.Worker, result2 error) {
	fake.visibleWorkersMutex.Lock()
	defer fake.visibleWorkersMutex.Unlock()
	fake.VisibleWorkersStub = nil
	fake.visibleWorkersReturns = struct {
		result1 []db.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) VisibleWorkersReturnsOnCall(i int, result1 []db.Worker, result2 error) {
	fake.visibleWorkersMutex.Lock()
	defer fake.visibleWorkersMutex.Unlock()
	fake.VisibleWorkersStub = nil
	if fake.visibleWorkersReturnsOnCall == nil {
		fake.visibleWorkersReturnsOnCall = make(map[int]struct {
			result1 []db.Worker
			result2 error
		})
	}
	fake.visibleWorkersReturnsOnCall[i] = struct {
		result1 []db.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) Workers() ([]db.Worker, error) {
	fake.workersMutex.Lock()
	ret, specificReturn := fake.workersReturnsOnCall[len(fake.workersArgsForCall)]
	fake.workersArgsForCall = append(fake.workersArgsForCall, struct {
	}{})
	stub := fake.WorkersStub
	fakeReturns := fake.workersReturns
	fake.recordInvocation("Workers", []interface{}{})
	fake.workersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorkerFactory) WorkersCallCount() int {
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	return len(fake.workersArgsForCall)
}

func (fake *FakeWorkerFactory) WorkersCalls(stub func() ([]db.Worker, error)) {
	fake.workersMutex.Lock()
	defer fake.workersMutex.Unlock()
	fake.WorkersStub = stub
}

func (fake *FakeWorkerFactory) WorkersReturns(result1 []db.Worker, result2 error) {
	fake.workersMutex.Lock()
	defer fake.workersMutex.Unlock()
	fake.WorkersStub = nil
	fake.workersReturns = struct {
		result1 []db.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) WorkersReturnsOnCall(i int, result1 []db.Worker, result2 error) {
	fake.workersMutex.Lock()
	defer fake.workersMutex.Unlock()
	fake.WorkersStub = nil
	if fake.workersReturnsOnCall == nil {
		fake.workersReturnsOnCall = make(map[int]struct {
			result1 []db.Worker
			result2 error
		})
	}
	fake.workersReturnsOnCall[i] = struct {
		result1 []db.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildContainersCountPerWorkerMutex.RLock()
	defer fake.buildContainersCountPerWorkerMutex.RUnlock()
	fake.findWorkersForContainerByOwnerMutex.RLock()
	defer fake.findWorkersForContainerByOwnerMutex.RUnlock()
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	fake.heartbeatWorkerMutex.RLock()
	defer fake.heartbeatWorkerMutex.RUnlock()
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	fake.visibleWorkersMutex.RLock()
	defer fake.visibleWorkersMutex.RUnlock()
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorkerFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.WorkerFactory = new(FakeWorkerFactory)
