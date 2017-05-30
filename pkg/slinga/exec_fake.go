package slinga

import "time"

// FakeCodeExecutor is a mock executor that does nothing
type FakeCodeExecutor struct {
	Code  *Code
	Delay time.Duration
}

// NewFakeCodeExecutor constructs FakeCodeExecutor from given *Code
func NewFakeCodeExecutor(code *Code, key string, codeMetadata map[string]string, codeParams interface{}, clusters map[string]*Cluster) FakeCodeExecutor {
	return FakeCodeExecutor{Code: code}
}

// NewFakeCodeExecutorWithDelay constructs FakeCodeExecutor from given *Code with specified delay
func NewFakeCodeExecutorWithDelay(code *Code, key string, codeMetadata map[string]string, codeParams interface{}, clusters map[string]*Cluster, delay time.Duration) FakeCodeExecutor {
	return FakeCodeExecutor{Code: code, Delay: delay}
}

// Install for FakeCodeExecutor does nothing
func (executor FakeCodeExecutor) Install() error {
	time.Sleep(executor.Delay)
	return nil
}

// Update for FakeCodeExecutor does nothing
func (executor FakeCodeExecutor) Update() error {
	time.Sleep(executor.Delay)
	return nil
}

// Destroy for FakeCodeExecutor does nothing
func (executor FakeCodeExecutor) Destroy() error {
	time.Sleep(executor.Delay)
	return nil
}

func (executor FakeCodeExecutor) Endpoints() (map[string]string, error) {
	return nil, nil
}