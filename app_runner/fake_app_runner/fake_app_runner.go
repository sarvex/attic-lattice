// This file was generated by counterfeiter
package fake_app_runner

import (
	"sync"

	"github.com/pivotal-cf-experimental/lattice-cli/app_runner"
)

type FakeAppRunner struct {
	StartDockerAppStub        func(name, startCommand, dockerImagePath string, appArgs []string, environmentVariables map[string]string, privileged bool, instances, memoryMB, diskMB, port int) error
	startDockerAppMutex       sync.RWMutex
	startDockerAppArgsForCall []struct {
		name                 string
		startCommand         string
		dockerImagePath      string
		appArgs              []string
		environmentVariables map[string]string
		privileged           bool
		instances            int
		memoryMB             int
		diskMB               int
		port                 int
	}
	startDockerAppReturns struct {
		result1 error
	}
	ScaleAppStub        func(name string, instances int) error
	scaleAppMutex       sync.RWMutex
	scaleAppArgsForCall []struct {
		name      string
		instances int
	}
	scaleAppReturns struct {
		result1 error
	}
	RemoveAppStub        func(name string) error
	removeAppMutex       sync.RWMutex
	removeAppArgsForCall []struct {
		name string
	}
	removeAppReturns struct {
		result1 error
	}
	AppExistsStub        func(name string) (bool, error)
	appExistsMutex       sync.RWMutex
	appExistsArgsForCall []struct {
		name string
	}
	appExistsReturns struct {
		result1 bool
		result2 error
	}
	NumOfRunningAppInstancesStub        func(name string) (int, error)
	numOfRunningAppInstancesMutex       sync.RWMutex
	numOfRunningAppInstancesArgsForCall []struct {
		name string
	}
	numOfRunningAppInstancesReturns struct {
		result1 int
		result2 error
	}
}

func (fake *FakeAppRunner) StartDockerApp(name string, startCommand string, dockerImagePath string, appArgs []string, environmentVariables map[string]string, privileged bool, instances int, memoryMB int, diskMB int, port int) error {
	fake.startDockerAppMutex.Lock()
	fake.startDockerAppArgsForCall = append(fake.startDockerAppArgsForCall, struct {
		name                 string
		startCommand         string
		dockerImagePath      string
		appArgs              []string
		environmentVariables map[string]string
		privileged           bool
		instances            int
		memoryMB             int
		diskMB               int
		port                 int
	}{name, startCommand, dockerImagePath, appArgs, environmentVariables, privileged, instances, memoryMB, diskMB, port})
	fake.startDockerAppMutex.Unlock()
	if fake.StartDockerAppStub != nil {
		return fake.StartDockerAppStub(name, startCommand, dockerImagePath, appArgs, environmentVariables, privileged, instances, memoryMB, diskMB, port)
	} else {
		return fake.startDockerAppReturns.result1
	}
}

func (fake *FakeAppRunner) StartDockerAppCallCount() int {
	fake.startDockerAppMutex.RLock()
	defer fake.startDockerAppMutex.RUnlock()
	return len(fake.startDockerAppArgsForCall)
}

func (fake *FakeAppRunner) StartDockerAppArgsForCall(i int) (string, string, string, []string, map[string]string, bool, int, int, int, int) {
	fake.startDockerAppMutex.RLock()
	defer fake.startDockerAppMutex.RUnlock()
	return fake.startDockerAppArgsForCall[i].name, fake.startDockerAppArgsForCall[i].startCommand, fake.startDockerAppArgsForCall[i].dockerImagePath, fake.startDockerAppArgsForCall[i].appArgs, fake.startDockerAppArgsForCall[i].environmentVariables, fake.startDockerAppArgsForCall[i].privileged, fake.startDockerAppArgsForCall[i].instances, fake.startDockerAppArgsForCall[i].memoryMB, fake.startDockerAppArgsForCall[i].diskMB, fake.startDockerAppArgsForCall[i].port
}

func (fake *FakeAppRunner) StartDockerAppReturns(result1 error) {
	fake.StartDockerAppStub = nil
	fake.startDockerAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppRunner) ScaleApp(name string, instances int) error {
	fake.scaleAppMutex.Lock()
	fake.scaleAppArgsForCall = append(fake.scaleAppArgsForCall, struct {
		name      string
		instances int
	}{name, instances})
	fake.scaleAppMutex.Unlock()
	if fake.ScaleAppStub != nil {
		return fake.ScaleAppStub(name, instances)
	} else {
		return fake.scaleAppReturns.result1
	}
}

func (fake *FakeAppRunner) ScaleAppCallCount() int {
	fake.scaleAppMutex.RLock()
	defer fake.scaleAppMutex.RUnlock()
	return len(fake.scaleAppArgsForCall)
}

func (fake *FakeAppRunner) ScaleAppArgsForCall(i int) (string, int) {
	fake.scaleAppMutex.RLock()
	defer fake.scaleAppMutex.RUnlock()
	return fake.scaleAppArgsForCall[i].name, fake.scaleAppArgsForCall[i].instances
}

func (fake *FakeAppRunner) ScaleAppReturns(result1 error) {
	fake.ScaleAppStub = nil
	fake.scaleAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppRunner) RemoveApp(name string) error {
	fake.removeAppMutex.Lock()
	fake.removeAppArgsForCall = append(fake.removeAppArgsForCall, struct {
		name string
	}{name})
	fake.removeAppMutex.Unlock()
	if fake.RemoveAppStub != nil {
		return fake.RemoveAppStub(name)
	} else {
		return fake.removeAppReturns.result1
	}
}

func (fake *FakeAppRunner) RemoveAppCallCount() int {
	fake.removeAppMutex.RLock()
	defer fake.removeAppMutex.RUnlock()
	return len(fake.removeAppArgsForCall)
}

func (fake *FakeAppRunner) RemoveAppArgsForCall(i int) string {
	fake.removeAppMutex.RLock()
	defer fake.removeAppMutex.RUnlock()
	return fake.removeAppArgsForCall[i].name
}

func (fake *FakeAppRunner) RemoveAppReturns(result1 error) {
	fake.RemoveAppStub = nil
	fake.removeAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppRunner) AppExists(name string) (bool, error) {
	fake.appExistsMutex.Lock()
	fake.appExistsArgsForCall = append(fake.appExistsArgsForCall, struct {
		name string
	}{name})
	fake.appExistsMutex.Unlock()
	if fake.AppExistsStub != nil {
		return fake.AppExistsStub(name)
	} else {
		return fake.appExistsReturns.result1, fake.appExistsReturns.result2
	}
}

func (fake *FakeAppRunner) AppExistsCallCount() int {
	fake.appExistsMutex.RLock()
	defer fake.appExistsMutex.RUnlock()
	return len(fake.appExistsArgsForCall)
}

func (fake *FakeAppRunner) AppExistsArgsForCall(i int) string {
	fake.appExistsMutex.RLock()
	defer fake.appExistsMutex.RUnlock()
	return fake.appExistsArgsForCall[i].name
}

func (fake *FakeAppRunner) AppExistsReturns(result1 bool, result2 error) {
	fake.AppExistsStub = nil
	fake.appExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeAppRunner) NumOfRunningAppInstances(name string) (int, error) {
	fake.numOfRunningAppInstancesMutex.Lock()
	fake.numOfRunningAppInstancesArgsForCall = append(fake.numOfRunningAppInstancesArgsForCall, struct {
		name string
	}{name})
	fake.numOfRunningAppInstancesMutex.Unlock()
	if fake.NumOfRunningAppInstancesStub != nil {
		return fake.NumOfRunningAppInstancesStub(name)
	} else {
		return fake.numOfRunningAppInstancesReturns.result1, fake.numOfRunningAppInstancesReturns.result2
	}
}

func (fake *FakeAppRunner) NumOfRunningAppInstancesCallCount() int {
	fake.numOfRunningAppInstancesMutex.RLock()
	defer fake.numOfRunningAppInstancesMutex.RUnlock()
	return len(fake.numOfRunningAppInstancesArgsForCall)
}

func (fake *FakeAppRunner) NumOfRunningAppInstancesArgsForCall(i int) string {
	fake.numOfRunningAppInstancesMutex.RLock()
	defer fake.numOfRunningAppInstancesMutex.RUnlock()
	return fake.numOfRunningAppInstancesArgsForCall[i].name
}

func (fake *FakeAppRunner) NumOfRunningAppInstancesReturns(result1 int, result2 error) {
	fake.NumOfRunningAppInstancesStub = nil
	fake.numOfRunningAppInstancesReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

var _ app_runner.AppRunner = new(FakeAppRunner)
