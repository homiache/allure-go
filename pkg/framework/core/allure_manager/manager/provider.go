package manager

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/adapter"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type allureManager struct {
	testMeta  provider.TestMeta
	suiteMeta provider.SuiteMeta

	executionContext provider.ExecutionContext
}

func NewProvider(cfg ProviderConfig) provider.Provider {
	suiteMeta := adapter.NewSuiteMeta(cfg.PackageName(), cfg.Runner(), cfg.FullName(), cfg.SuiteName())
	return &allureManager{suiteMeta: suiteMeta, testMeta: &adapter.TestAdapter{}}
}

func (a *allureManager) safely(f func(result *allure.Result)) {
	if result := a.GetResult(); result != nil {
		f(result)
	}
}

func (a *allureManager) UpdateResultStatus(msg string, trace string) {
	a.GetResult().SetStatusMessage(msg)
	a.GetResult().SetStatusTrace(trace)
}

func (a *allureManager) StopResult(status allure.Status) {
	a.safely(func(result *allure.Result) {
		result.Status = status
		result.Stop = allure.GetNow()
	})
}

func (a *allureManager) GetResult() *allure.Result {
	return a.testMeta.GetResult()
}

func (a *allureManager) GetTestMeta() provider.TestMeta {
	return a.testMeta
}

func (a *allureManager) GetSuiteMeta() provider.SuiteMeta {
	return a.suiteMeta
}

func (a *allureManager) ExecutionContext() provider.ExecutionContext {
	return a.executionContext
}

func (a *allureManager) NewTest(testName, packageName string, tags ...string) {
	a.testMeta = adapter.NewTestMeta(
		a.suiteMeta.GetSuiteFullName(),
		a.suiteMeta.GetSuiteName(),
		testName,
		packageName,
		tags...,
	)
	a.suiteMeta.GetContainer().AddChild(a.testMeta.GetResult().UUID)
}

func (a *allureManager) FinishTest() {
	a.testMeta.GetResult().Done()
}
