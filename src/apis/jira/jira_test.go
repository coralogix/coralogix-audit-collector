package jira

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/coralogix/c4c-ir-integrations/src/tests"
	"testing"
)

func TestAPIGetReportsWithMapping(t *testing.T) {
	expected := 2
	json := map[string][]string{
		"/rest/api/3/auditing/record": []string{
			`{"offset":0,"limit":1,"total":2,"records":[{"id":15721,"summary":"Custom field created","remoteAddress":"62.90.112.205","authorKey":"ug:16064703-ada5-43df-a2ee-6419522c1506","authorAccountId":"620256b4506317006b07bc02","created":"2022-11-14T15:17:26.697+0000","category":"fields","eventSource":"","objectItem":{"id":"customfield_10140","name":"Who is moving?","typeName":"CUSTOM_FIELD"},"changedValues":[{"fieldName":"Name","changedTo":"Who is moving?"},{"fieldName":"Description","changedTo":""},{"fieldName":"Type","changedTo":"User Picker (multiple users)"}]}]}`,
			`{"offset":1,"limit":1,"total":2,"records":[{"id":15721,"summary":"Custom field created","remoteAddress":"62.90.112.205","authorKey":"ug:16064703-ada5-43df-a2ee-6419522c1506","authorAccountId":"620256b4506317006b07bc02","created":"2022-11-14T15:17:26.697+0000","category":"fields","eventSource":"","objectItem":{"id":"customfield_10140","name":"Who is moving?","typeName":"CUSTOM_FIELD"},"associatedItems":[{"id":"dd1573d1-0099-45db-9833-befb66121124","name":"dd1573d1-0099-45db-9833-befb66121124","typeName":"USER","parentId":"10000","parentName":"com.atlassian.crowd.directory.IdentityPlatformRemoteDirectory"}],"changedValues":[{"fieldName":"Name","changedTo":"Who is moving?"},{"fieldName":"Description","changedTo":""},{"fieldName":"Type","changedTo":"User Picker (multiple users)"}]}]}`,
		},
		"/rest/api/3/user/bulk/migration": []string{
			`[{"username": "dd1573d1-0099-45db-9833-befb66121124", "accountId": "620256b4506317006b07bc02"}]`,
		},
		"/rest/api/3/user": []string{
			`{"accountId": "620256b4506317006b07bc02", "emailAddress": "a@a.com"}`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)

	api := New(
		"username",
		"apiToken",
		"http://baseurl",
		mockClient,
	)
	tpm := integration.NewTimePeriodManagerFromNow(1)
	ret, err := api.GetReports(tpm)

	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if len(ret) != expected {
		t.Errorf("Expected %d, got %d", expected, len(ret))
	}
}

func TestAPIGetReportsWithoutMapping(t *testing.T) {
	expected := 2
	json := map[string][]string{
		"/rest/api/3/auditing/record": []string{
			`{"offset":0,"limit":1,"total":2,"records":[{"id":15721,"summary":"Custom field created","remoteAddress":"62.90.112.205","authorKey":"ug:16064703-ada5-43df-a2ee-6419522c1506","authorAccountId":"620256b4506317006b07bc02","created":"2022-11-14T15:17:26.697+0000","category":"fields","eventSource":"","objectItem":{"id":"customfield_10140","name":"Who is moving?","typeName":"CUSTOM_FIELD"},"changedValues":[{"fieldName":"Name","changedTo":"Who is moving?"},{"fieldName":"Description","changedTo":""},{"fieldName":"Type","changedTo":"User Picker (multiple users)"}]}]}`,
			`{"offset":1,"limit":1,"total":2,"records":[{"id":15721,"summary":"Custom field created","remoteAddress":"62.90.112.205","authorKey":"ug:16064703-ada5-43df-a2ee-6419522c1506","authorAccountId":"620256b4506317006b07bc02","created":"2022-11-14T15:17:26.697+0000","category":"fields","eventSource":"","objectItem":{"id":"customfield_10140","name":"Who is moving?","typeName":"CUSTOM_FIELD"},"associatedItems":[{"id":"dd1573d1-0099-45db-9833-befb66121124","name":"dd1573d1-0099-45db-9833-befb66121124","typeName":"USER","parentId":"10000","parentName":"com.atlassian.crowd.directory.IdentityPlatformRemoteDirectory"}],"changedValues":[{"fieldName":"Name","changedTo":"Who is moving?"},{"fieldName":"Description","changedTo":""},{"fieldName":"Type","changedTo":"User Picker (multiple users)"}]}]}`,
		},
		"/rest/api/3/user/bulk/migration": []string{
			`[]`,
		},
		"/rest/api/3/user": []string{
			`{"accountId": "620256b4506317006b07bc02", "emailAddress": "a@a.com"}`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)

	api := New(
		"username",
		"apiToken",
		"http://baseurl",
		mockClient,
	)
	tpm := integration.NewTimePeriodManagerFromNow(1)
	ret, err := api.GetReports(tpm)

	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if len(ret) != expected {
		t.Errorf("Expected %d, got %d", expected, len(ret))
	}
}
