package analyzer

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i mydb/internal/database/compute.Analyzer -o ./analyzer/analyzer_mock.go -n AnalyzerMock

import (
	"mydb/internal/domain/operations"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// AnalyzerMock implements compute.Analyzer
type AnalyzerMock struct {
	t minimock.Tester

	funcAnalyze          func(args []string) (o1 operations.Operation, err error)
	inspectFuncAnalyze   func(args []string)
	afterAnalyzeCounter  uint64
	beforeAnalyzeCounter uint64
	AnalyzeMock          mAnalyzerMockAnalyze
}

// NewAnalyzerMock returns a mock for compute.Analyzer
func NewAnalyzerMock(t minimock.Tester) *AnalyzerMock {
	m := &AnalyzerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AnalyzeMock = mAnalyzerMockAnalyze{mock: m}
	m.AnalyzeMock.callArgs = []*AnalyzerMockAnalyzeParams{}

	return m
}

type mAnalyzerMockAnalyze struct {
	mock               *AnalyzerMock
	defaultExpectation *AnalyzerMockAnalyzeExpectation
	expectations       []*AnalyzerMockAnalyzeExpectation

	callArgs []*AnalyzerMockAnalyzeParams
	mutex    sync.RWMutex
}

// AnalyzerMockAnalyzeExpectation specifies expectation struct of the Analyzer.Analyze
type AnalyzerMockAnalyzeExpectation struct {
	mock    *AnalyzerMock
	params  *AnalyzerMockAnalyzeParams
	results *AnalyzerMockAnalyzeResults
	Counter uint64
}

// AnalyzerMockAnalyzeParams contains parameters of the Analyzer.Analyze
type AnalyzerMockAnalyzeParams struct {
	args []string
}

// AnalyzerMockAnalyzeResults contains results of the Analyzer.Analyze
type AnalyzerMockAnalyzeResults struct {
	o1  operations.Operation
	err error
}

// Expect sets up expected params for Analyzer.Analyze
func (mmAnalyze *mAnalyzerMockAnalyze) Expect(args []string) *mAnalyzerMockAnalyze {
	if mmAnalyze.mock.funcAnalyze != nil {
		mmAnalyze.mock.t.Fatalf("AnalyzerMock.Analyze mock is already set by Set")
	}

	if mmAnalyze.defaultExpectation == nil {
		mmAnalyze.defaultExpectation = &AnalyzerMockAnalyzeExpectation{}
	}

	mmAnalyze.defaultExpectation.params = &AnalyzerMockAnalyzeParams{args}
	for _, e := range mmAnalyze.expectations {
		if minimock.Equal(e.params, mmAnalyze.defaultExpectation.params) {
			mmAnalyze.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAnalyze.defaultExpectation.params)
		}
	}

	return mmAnalyze
}

// Inspect accepts an inspector function that has same arguments as the Analyzer.Analyze
func (mmAnalyze *mAnalyzerMockAnalyze) Inspect(f func(args []string)) *mAnalyzerMockAnalyze {
	if mmAnalyze.mock.inspectFuncAnalyze != nil {
		mmAnalyze.mock.t.Fatalf("Inspect function is already set for AnalyzerMock.Analyze")
	}

	mmAnalyze.mock.inspectFuncAnalyze = f

	return mmAnalyze
}

// Return sets up results that will be returned by Analyzer.Analyze
func (mmAnalyze *mAnalyzerMockAnalyze) Return(o1 operations.Operation, err error) *AnalyzerMock {
	if mmAnalyze.mock.funcAnalyze != nil {
		mmAnalyze.mock.t.Fatalf("AnalyzerMock.Analyze mock is already set by Set")
	}

	if mmAnalyze.defaultExpectation == nil {
		mmAnalyze.defaultExpectation = &AnalyzerMockAnalyzeExpectation{mock: mmAnalyze.mock}
	}
	mmAnalyze.defaultExpectation.results = &AnalyzerMockAnalyzeResults{o1, err}
	return mmAnalyze.mock
}

// Set uses given function f to mock the Analyzer.Analyze method
func (mmAnalyze *mAnalyzerMockAnalyze) Set(f func(args []string) (o1 operations.Operation, err error)) *AnalyzerMock {
	if mmAnalyze.defaultExpectation != nil {
		mmAnalyze.mock.t.Fatalf("Default expectation is already set for the Analyzer.Analyze method")
	}

	if len(mmAnalyze.expectations) > 0 {
		mmAnalyze.mock.t.Fatalf("Some expectations are already set for the Analyzer.Analyze method")
	}

	mmAnalyze.mock.funcAnalyze = f
	return mmAnalyze.mock
}

// When sets expectation for the Analyzer.Analyze which will trigger the result defined by the following
// Then helper
func (mmAnalyze *mAnalyzerMockAnalyze) When(args []string) *AnalyzerMockAnalyzeExpectation {
	if mmAnalyze.mock.funcAnalyze != nil {
		mmAnalyze.mock.t.Fatalf("AnalyzerMock.Analyze mock is already set by Set")
	}

	expectation := &AnalyzerMockAnalyzeExpectation{
		mock:   mmAnalyze.mock,
		params: &AnalyzerMockAnalyzeParams{args},
	}
	mmAnalyze.expectations = append(mmAnalyze.expectations, expectation)
	return expectation
}

// Then sets up Analyzer.Analyze return parameters for the expectation previously defined by the When method
func (e *AnalyzerMockAnalyzeExpectation) Then(o1 operations.Operation, err error) *AnalyzerMock {
	e.results = &AnalyzerMockAnalyzeResults{o1, err}
	return e.mock
}

// Analyze implements compute.Analyzer
func (mmAnalyze *AnalyzerMock) Analyze(args []string) (o1 operations.Operation, err error) {
	mm_atomic.AddUint64(&mmAnalyze.beforeAnalyzeCounter, 1)
	defer mm_atomic.AddUint64(&mmAnalyze.afterAnalyzeCounter, 1)

	if mmAnalyze.inspectFuncAnalyze != nil {
		mmAnalyze.inspectFuncAnalyze(args)
	}

	mm_params := &AnalyzerMockAnalyzeParams{args}

	// Record call args
	mmAnalyze.AnalyzeMock.mutex.Lock()
	mmAnalyze.AnalyzeMock.callArgs = append(mmAnalyze.AnalyzeMock.callArgs, mm_params)
	mmAnalyze.AnalyzeMock.mutex.Unlock()

	for _, e := range mmAnalyze.AnalyzeMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.o1, e.results.err
		}
	}

	if mmAnalyze.AnalyzeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAnalyze.AnalyzeMock.defaultExpectation.Counter, 1)
		mm_want := mmAnalyze.AnalyzeMock.defaultExpectation.params
		mm_got := AnalyzerMockAnalyzeParams{args}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAnalyze.t.Errorf("AnalyzerMock.Analyze got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAnalyze.AnalyzeMock.defaultExpectation.results
		if mm_results == nil {
			mmAnalyze.t.Fatal("No results are set for the AnalyzerMock.Analyze")
		}
		return (*mm_results).o1, (*mm_results).err
	}
	if mmAnalyze.funcAnalyze != nil {
		return mmAnalyze.funcAnalyze(args)
	}
	mmAnalyze.t.Fatalf("Unexpected call to AnalyzerMock.Analyze. %v", args)
	return
}

// AnalyzeAfterCounter returns a count of finished AnalyzerMock.Analyze invocations
func (mmAnalyze *AnalyzerMock) AnalyzeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAnalyze.afterAnalyzeCounter)
}

// AnalyzeBeforeCounter returns a count of AnalyzerMock.Analyze invocations
func (mmAnalyze *AnalyzerMock) AnalyzeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAnalyze.beforeAnalyzeCounter)
}

// Calls returns a list of arguments used in each call to AnalyzerMock.Analyze.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAnalyze *mAnalyzerMockAnalyze) Calls() []*AnalyzerMockAnalyzeParams {
	mmAnalyze.mutex.RLock()

	argCopy := make([]*AnalyzerMockAnalyzeParams, len(mmAnalyze.callArgs))
	copy(argCopy, mmAnalyze.callArgs)

	mmAnalyze.mutex.RUnlock()

	return argCopy
}

// MinimockAnalyzeDone returns true if the count of the Analyze invocations corresponds
// the number of defined expectations
func (m *AnalyzerMock) MinimockAnalyzeDone() bool {
	for _, e := range m.AnalyzeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AnalyzeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAnalyzeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAnalyze != nil && mm_atomic.LoadUint64(&m.afterAnalyzeCounter) < 1 {
		return false
	}
	return true
}

// MinimockAnalyzeInspect logs each unmet expectation
func (m *AnalyzerMock) MinimockAnalyzeInspect() {
	for _, e := range m.AnalyzeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AnalyzerMock.Analyze with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AnalyzeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAnalyzeCounter) < 1 {
		if m.AnalyzeMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AnalyzerMock.Analyze")
		} else {
			m.t.Errorf("Expected call to AnalyzerMock.Analyze with params: %#v", *m.AnalyzeMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAnalyze != nil && mm_atomic.LoadUint64(&m.afterAnalyzeCounter) < 1 {
		m.t.Error("Expected call to AnalyzerMock.Analyze")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *AnalyzerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAnalyzeInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *AnalyzerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *AnalyzerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAnalyzeDone()
}
