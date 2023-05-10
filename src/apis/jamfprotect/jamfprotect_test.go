package jamfprotect

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/coralogix/c4c-ir-integrations/src/tests"
	"testing"
)

func TestAPIGetReports(t *testing.T) {
	expected := 6
	json := map[string][]string{
		"/token": []string{
			`{"access_token": "ACCESS_TOKEN", "expires_in": 86366, "token_type": "Bearer"}`,
		},
		"/graphql": []string{
			`{"data":{"listAuditLogsByDate":{"items":[{"resourceId":"1","date":"2023-02-13T17:36:15.053Z","args":"{\"input\":{\"roleIds\":[\"2\"],\"name\":\"c4c-ir\"}}","error":null,"ips":"41.121.112.118","op":"createApiClient","user":"test","__typename":"AuditLog"},{"resourceId":"199","date":"2023-02-07T08:21:29.117Z","args":"{\"input\":\"test\"}","error":null,"ips":"","op":"createUser","user":"test","__typename":"AuditLog"},{"resourceId":"166","date":"2023-02-06T15:44:13.385Z","args":"{\"input\":\"test\"}","error":null,"ips":"","op":"createUser","user":"test","__typename":"AuditLog"}],"pageInfo":{"next":"test","__typename":"PageInfo"},"__typename":"AuditLogConnection"}}}`,
			`{"data":{"listAuditLogsByDate":{"items":[{"resourceId":"1","date":"2023-02-11T17:36:15.053Z","args":"{\"input\":{\"roleIds\":[\"2\"],\"name\":\"c4c-ir\"}}","error":null,"ips":"41.121.112.118","op":"createApiClient","user":"test","__typename":"AuditLog"},{"resourceId":"199","date":"2023-02-07T08:21:29.117Z","args":"{\"input\":\"test\"}","error":null,"ips":"","op":"createUser","user":"test","__typename":"AuditLog"},{"resourceId":"166","date":"2023-02-06T15:44:13.385Z","args":"{\"input\":\"test\"}","error":null,"ips":"","op":"createUser","user":"test","__typename":"AuditLog"}],"pageInfo":{"next":null,"__typename":"PageInfo"},"__typename":"AuditLogConnection"}}}`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)

	api := New(
		"clientId",
		"apiToken",
		"http://baseurl",
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

func TestAPIGetReportsNoAccess(t *testing.T) {
	json := map[string][]string{
		"/token": []string{
			`{"error": "access_denied", "error_description": "Unauthorized"}`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)

	api := New(
		"clientId",
		"apiToken",
		"http://baseurl",
		mockClient,
	)

	i := integration.NewTimePeriodManagerFromNow(1)

	_, err := api.GetReports(i)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
