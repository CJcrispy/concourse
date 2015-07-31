// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/concourse/atc/exec"
)

type FakeGetDelegate struct {
	CompletedStub        func(exec.ExitStatus, *exec.VersionInfo)
	completedMutex       sync.RWMutex
	completedArgsForCall []struct {
		arg1 exec.ExitStatus
		arg2 *exec.VersionInfo
	}
	FailedStub        func(error)
	failedMutex       sync.RWMutex
	failedArgsForCall []struct {
		arg1 error
	}
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct{}
	stdoutReturns     struct {
		result1 io.Writer
	}
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct{}
	stderrReturns     struct {
		result1 io.Writer
	}
}

func (fake *FakeGetDelegate) Completed(arg1 exec.ExitStatus, arg2 *exec.VersionInfo) {
	fake.completedMutex.Lock()
	fake.completedArgsForCall = append(fake.completedArgsForCall, struct {
		arg1 exec.ExitStatus
		arg2 *exec.VersionInfo
	}{arg1, arg2})
	fake.completedMutex.Unlock()
	if fake.CompletedStub != nil {
		fake.CompletedStub(arg1, arg2)
	}
}

func (fake *FakeGetDelegate) CompletedCallCount() int {
	fake.completedMutex.RLock()
	defer fake.completedMutex.RUnlock()
	return len(fake.completedArgsForCall)
}

func (fake *FakeGetDelegate) CompletedArgsForCall(i int) (exec.ExitStatus, *exec.VersionInfo) {
	fake.completedMutex.RLock()
	defer fake.completedMutex.RUnlock()
	return fake.completedArgsForCall[i].arg1, fake.completedArgsForCall[i].arg2
}

func (fake *FakeGetDelegate) Failed(arg1 error) {
	fake.failedMutex.Lock()
	fake.failedArgsForCall = append(fake.failedArgsForCall, struct {
		arg1 error
	}{arg1})
	fake.failedMutex.Unlock()
	if fake.FailedStub != nil {
		fake.FailedStub(arg1)
	}
}

func (fake *FakeGetDelegate) FailedCallCount() int {
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	return len(fake.failedArgsForCall)
}

func (fake *FakeGetDelegate) FailedArgsForCall(i int) error {
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	return fake.failedArgsForCall[i].arg1
}

func (fake *FakeGetDelegate) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct{}{})
	fake.stdoutMutex.Unlock()
	if fake.StdoutStub != nil {
		return fake.StdoutStub()
	} else {
		return fake.stdoutReturns.result1
	}
}

func (fake *FakeGetDelegate) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakeGetDelegate) StdoutReturns(result1 io.Writer) {
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeGetDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct{}{})
	fake.stderrMutex.Unlock()
	if fake.StderrStub != nil {
		return fake.StderrStub()
	} else {
		return fake.stderrReturns.result1
	}
}

func (fake *FakeGetDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeGetDelegate) StderrReturns(result1 io.Writer) {
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

var _ exec.GetDelegate = new(FakeGetDelegate)
