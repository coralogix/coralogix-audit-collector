package main

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"testing"
	"time"
)

type MemoryAPI struct {
	reports integration.APIResult
}
type MemoryCollector struct {
	collected integration.APIResult
}

func (m *MemoryAPI) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	reporting := integration.APIResult{}
	for _, v := range m.reports {
		t, _ := time.Parse(m.GetDateTimeFormat(), v["Time"].(string))
		if t.After(timePeriodManager.From()) {
			reporting = append(reporting, v)
		}
	}
	return reporting, nil
}

func (m *MemoryAPI) GetDateTimeFormat() string {
	return "2006-01-02 15:04:05"
}

func (m *MemoryCollector) Collect(v map[string]interface{}) {
	m.collected = append(m.collected, v)
}

func (m *MemoryCollector) Send() {

}

func TestNewPublishReportsUseCaseWithResults(t *testing.T) {
	runEveryIntervalMinutes := 1
	reports := integration.APIResult{
		{
			"Data": "in",
			"Time": time.Now().Format("2006-01-02 15:04:05"),
		}, {
			"Data": "out",
			"Time": time.Now().Add(-time.Duration(24) * time.Hour).Format("2006-01-02 15:04:05"),
		},
	}
	expected := len(reports) - 1
	api := &MemoryAPI{
		reports: reports,
	}
	collector := &MemoryCollector{
		collected: integration.APIResult{},
	}
	i := integration.NewTimePeriodManagerFromNow(runEveryIntervalMinutes)
	usecase := NewPublishReportsUseCase(api, collector)
	usecase.execute(i)
	if len(collector.collected) != expected {
		t.Errorf("expected %d reports, got %d", len(reports), len(collector.collected))
	}
}
