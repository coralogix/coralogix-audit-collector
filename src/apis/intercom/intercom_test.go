package intercom

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/coralogix/c4c-ir-integrations/src/tests"
	"testing"
)

func TestAPIGetReports(t *testing.T) {
	json := map[string][]string{
		"/admins/activity_logs": []string{
			`{"type":"activity_log.list","activity_logs":[{"activity_type":"admin.deleted","activity_description":"Admin deleted","metadata":{"admin_id":"2","admin_name":"John Doe","t":{"test123":123}},"created_at":1677437424,"performed_by":{"type":"admin","id":"2","name":"John Doe","email":"a@a.com","job_title":"CEO","away_mode_enabled":false,"away_mode_reassign":false,"has_inbox_seat":true,"team_ids":["123"],"avatar":"https://example.com/avatar.png","team_priority_level":"team"},"id":"1"}],"pages":{"next":"https://api.helpscout.net/admins/activity_logs?page=2","perPage":50,"page":1,"totalPage":2}}`,
			`{"type":"activity_log.list","activity_logs":[{"activity_type":"admin.created","activity_description":"Admin created","metadata":{"admin_id":"2","admin_name":"John Doe","t":{"test123":123}},"created_at":1677091824,"performed_by":{"type":"admin","id":"2","name":"John Doe","email":"a@a.com","job_title":"CEO","away_mode_enabled":false,"away_mode_reassign":false,"has_inbox_seat":true,"team_ids":["123"],"avatar":"https://example.com/avatar.png","team_priority_level":"team"},"id":"1"}],"pages":{"next":"","perPage":50,"page":2,"totalPage":2}}`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)

	api := New(
		"123",
		"https://api.intercom.io",
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

	activityType := ret[0]["activity_type"]
	if activityType != "admin.deleted" {
		t.Errorf("Expected admin.deleted, got %s", activityType)
	}
}
