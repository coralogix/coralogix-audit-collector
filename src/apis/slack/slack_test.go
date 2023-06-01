package slack

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/coralogix/c4c-ir-integrations/src/tests"
	"testing"
)

func TestAPIGetReports(t *testing.T) {
	json := map[string][]string{
		"/audit/v1/logs": []string{
			`{"entries":[{"id":"1123a45b-6c7d-8900-e12f-3456789gh0i1","date_create":1521214343,"action":"user_login","actor":{"type":"user","user":{"id":"W123AB456","name":"Charlie Parker","email":"bird@slack.com"}},"entity":{"type":"user","user":{"id":"W123AB456","name":"Charlie Parker","email":"bird@slack.com"}},"context":{"location":{"type":"enterprise","id":"E1701NCCA","name":"Birdland","domain":"birdland"},"ua":"Mozilla\/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/64.0.3282.186 Safari\/537.36","session_id":"847288190092","ip_address":"1.23.45.678"}}],"response_metadata":{"next_cursor":"dXNlcjpVMEc5V0ZYTlo="}}`,
			`{"entries":[{"id":"2123a45b-6c7d-8900-e12f-3456789gh0i1","date_create":1521214343,"action":"user_login","actor":{"type":"user","user":{"id":"W123AB456","name":"Charlie Parker","email":"bird@slack.com"}},"entity":{"type":"user","user":{"id":"W123AB456","name":"Charlie Parker","email":"bird@slack.com"}},"context":{"location":{"type":"enterprise","id":"E1701NCCA","name":"Birdland","domain":"birdland"},"ua":"Mozilla\/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/64.0.3282.186 Safari\/537.36","session_id":"847288190092","ip_address":"1.23.45.678"}}],"response_metadata":{"next_cursor":""}}`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)

	api := New(
		"123",
		"https://api.slack.com",
		mockClient,
	)
	i := integration.NewTimePeriodManagerFromNow(1)

	ret, err := api.GetReports(i)

	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if len(ret) != 2 {
		t.Errorf("Expected 0, got %d", len(ret))
	}

	if ret[0]["id"] != "1123a45b-6c7d-8900-e12f-3456789gh0i1" {
		t.Errorf("Expected admin.deleted, got %s", ret[0]["id"])
	}

	if ret[1]["id"] != "2123a45b-6c7d-8900-e12f-3456789gh0i1" {
		t.Errorf("Expected admin.deleted, got %s", ret[1]["id"])
	}
}
