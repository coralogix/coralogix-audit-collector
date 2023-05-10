package zoom

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/coralogix/c4c-ir-integrations/src/tests"
	"testing"
)

func TestAPIGetReports(t *testing.T) {
	json := map[string][]string{
		"/v2/report/operationlogs": []string{
			`{"next_page_token": "page_token","page_size":0,"operation_logs":[{"action":"string","category_type":"string","created_at":"string","email":"string","ip_address":"string","login_type":"string","operation_time":"string"}]}`,
			`{"next_page_token": null,"page_size":0,"operation_logs":[{"action":"string","category_type":"string","created_at":"string","email":"string","ip_address":"string","login_type":"string","operation_time":"string"}]}`,
		},
		"/oauth/token": []string{
			`{"access_token":"string","token_type":"string","expire_in":0,"scope":"string"}`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)

	api := New(
		"accountId",
		"clientId",
		"clientSecret",
		"http://baseUrl",
		"http://accessTokenUrl",
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
}
