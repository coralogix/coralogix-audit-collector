package atlassian

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/coralogix/c4c-ir-integrations/src/tests"
	"testing"
)

func TestAPIGetReports(t *testing.T) {
	expected := 2
	orgId := "fixture_org_id"
	json := map[string][]string{
		"/admin/v1/orgs/fixture_org_id/events": []string{
			`{"data":[{"id":"<string>","type":"events","attributes":{"time":"<string>","action":"<string>","actor":{"id":"<string>","name":"<string>","email":"<string>","links":{"self":"<string>"}},"context":[{"id":"<string>","type":"<string>","attributes":{},"links":{"self":"<string>","alt":"<string>"}}],"container":[{"id":"<string>","type":"<string>","attributes":{},"links":{"self":"<string>","alt":"<string>"}}],"location":{"ip":"<string>","geo":"<string>","countryName":"<string>","regionName":"<string>","city":"<string>"}},"links":{"self":"<string>"}}],"meta":{"next":"/admin/v1/orgs/fixture_org_id/events","page_size":1},"links":{"self":"<string>","prev":"<string>","next":"<string>"}}`,
			`{"data":[{"id":"<string>","type":"events","attributes":{"time":"<string>","action":"<string>","actor":{"id":"<string>","name":"<string>","email":"<string>","links":{"self":"<string>"}},"context":[{"id":"<string>","type":"<string>","attributes":{},"links":{"self":"<string>","alt":"<string>"}}],"container":[{"id":"<string>","type":"<string>","attributes":{},"links":{"self":"<string>","alt":"<string>"}}],"location":{"ip":"<string>","geo":"<string>","countryName":"<string>","regionName":"<string>","city":"<string>"}},"links":{"self":"<string>"}}],"meta":{"next":"","page_size":1},"links":{"self":"<string>","prev":"<string>","next":"<string>"}  }`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)

	api := New(
		"clientId",
		"apiToken",
		orgId,
		"https://api.atlassian.com",
		mockClient,
	)
	i := integration.NewTimePeriodManagerFromNow(1)

	ret, err := api.GetReports(i)

	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if len(ret) != expected {
		t.Errorf("Expected %d, got %d", expected, len(ret))
	}
}
