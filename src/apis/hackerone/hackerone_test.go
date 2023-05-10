package hackerone

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/coralogix/c4c-ir-integrations/src/tests"
	"testing"
	"time"
)

func TestHackerOne(t *testing.T) {
	programId := "1"
	json := map[string][]string{
		"/v1/programs/1/audit_log": []string{
			`{"data":[{"id":"1","type":"audit-log-item","attributes":{"log":"\"@page1\" changed groups for \"na\" by adding \"H1 Triage\"","event":"teams.groups.update","source":"User#32123","subject":"TeamMember#2123","user_agent":null,"country":null,"parameters":"{\"group_ids\":[\"1234\"]}","created_at":"2021-12-13T21:20:42.516Z"}},{"id":"581427","type":"audit-log-item","attributes":{"log":"\"@name2\" changed groups for \"na\" by adding \"H1 Triage\"","event":"teams.groups.update","source":"User#32123","subject":"TeamMember#2123","user_agent":null,"country":null,"parameters":"{\"group_ids\":[\"1234\"]}","created_at":"2021-12-10T21:20:42.516Z"}}],"links":{"self":"https://api.hackerone.com/v1/programs/1/audit_log?page%5Bnumber%5D=1","next":"https://api.hackerone.com/v1/programs/1/audit_log?page%5Bnumber%5D=2","last":"https://api.hackerone.com/v1/programs/1/audit_log?page%5Bnumber%5D=2"}}`,
			`{"data":[{"id":"2","type":"audit-log-item","attributes":{"log":"\"@page2\" changed groups for \"na\" by adding \"H1 Triage\"","event":"teams.groups.update","source":"User#32123","subject":"TeamMember#2123","user_agent":null,"country":null,"parameters":"{\"group_ids\":[\"1234\"]}","created_at":"2021-12-14T21:20:42.516Z"}}],"links":{"self":"https://api.hackerone.com/v1/programs/1/audit_log?page%5Bnumber%5D=2","prev":"https://api.hackerone.com/v1/programs/1/audit_log?page%5Bnumber%5D=1"}}`,
			`{"data":[{"id":"1","type":"audit-log-item","attributes":{"log":"\"@page1\" changed groups for \"na\" by adding \"H1 Triage\"","event":"teams.groups.update","source":"User#32123","subject":"TeamMember#2123","user_agent":null,"country":null,"parameters":"{\"group_ids\":[\"1234\"]}","created_at":"2021-12-13T21:20:42.516Z"}},{"id":"581427","type":"audit-log-item","attributes":{"log":"\"@name2\" changed groups for \"na\" by adding \"H1 Triage\"","event":"teams.groups.update","source":"User#32123","subject":"TeamMember#2123","user_agent":null,"country":null,"parameters":"{\"group_ids\":[\"1234\"]}","created_at":"2021-12-10T21:20:42.516Z"}}],"links":{"self":"https://api.hackerone.com/v1/programs/1/audit_log?page%5Bnumber%5D=1","next":"https://api.hackerone.com/v1/programs/1/audit_log?page%5Bnumber%5D=2","last":"https://api.hackerone.com/v1/programs/1/audit_log?page%5Bnumber%5D=2"}}`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)

	startFrom, _ := time.Parse(time.RFC3339, "2021-12-14T21:20:42.516Z")
	api := New(
		"client_id",
		"client_secret",
		"https://api.hackerone.com",
		programId,
		mockClient,
	)
	i := integration.NewTimePeriodManager(startFrom, 1440)
	result, err := api.GetReports(i)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if result == nil {
		t.Errorf("Expected not nil, got nil")
	}

	if result[0]["id"] != "2" {
		t.Errorf("Expected 2, got %s", result[0]["id"])
	}

	if result[1]["id"] != "1" {
		t.Errorf("Expected 1, got %s", result[1]["id"])
	}

	if len(result) != 2 {
		t.Errorf("Expected 2, got %d", len(result))
	}
}
