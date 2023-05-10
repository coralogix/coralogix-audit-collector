package lastpass

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/coralogix/c4c-ir-integrations/src/tests"
	"testing"
)

func TestAPIGetReports(t *testing.T) {
	json := map[string][]string{
		"/enterpriseapi.php": []string{
			`{"status":"success","next":null,"data":{"2021-03-01 00:00:00":{"Time":"2021-03-01 00:00:00","Username":"user1","IP_Address":"","Action":"login","Data":"user1"},"2021-03-01 00:00:01":{"Time":"2021-03-01 00:00:01","Username":"user2","IP_Address":"","Action":"login","Data":"user2"}}}`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)
	api := New(
		"http://baseurl",
		"cid",
		"provhash",
		"apiuser",
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
