package monday

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/coralogix/c4c-ir-integrations/src/tests"
	"testing"
)

func TestAPIGetReports(t *testing.T) {
	json := map[string][]string{
		"/audit-api/get-logs": []string{
			`{"data":[{"timestamp":"2020-10-01T12:00:00Z","account_id":123456,"user_id":123456,"event":"login","slug":"login","ip_address":"123","user_agent":"test123","client_name":"test123","client_version":"test123","os_name":"test123","os_version":"test123","device_name":"test123","device_type":"test123","activity_metadata":{"user_id":123456,"user_name":"test123"}}],"page":1,"per_page":1,"next_page":2,"message":""}`,
			`{"data":[{"timestamp":"2020-10-01T12:00:00Z","account_id":123456,"user_id":123456,"event":"login","slug":"login","ip_address":"123","user_agent":"test123","client_name":"test123","client_version":"test123","os_name":"test123","os_version":"test123","device_name":"test123","device_type":"test123","activity_metadata":{"user_id":123456,"user_name":"test123"}}],"page":2,"per_page":1,"next_page":null,"message":""}`,
		},
	}

	mockClient := tests.NewMockWithFixtureByUrlPath(json)
	api := New(
		"http://baseurl",
		"apiToken",
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

	if ret[0]["user_id"] != "123456" {
		t.Errorf("Expected 123456, got %s", ret[0]["user_id"])
	}
}
