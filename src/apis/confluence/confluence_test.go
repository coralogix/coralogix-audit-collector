package confluence

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/coralogix/c4c-ir-integrations/src/tests"
	"testing"
)

func TestAPIGetReports(t *testing.T) {
	expected := 2
	json := map[string][]string{
		"/wiki/rest/api/audit": []string{
			`{"results":[{"author":{"type":"user","displayName":"Shaked Klein","operations":null,"isExternalCollaborator":false,"username":"63bed71b0913ada386aadbe4","userKey":"63bed71b0913ada386aadbe4","accountId":"63bed71b0913ada386aadbe4","accountType":"","publicName":"Unknown user","externalCollaborator":false},"remoteAddress":"","creationDate":1676297029823,"summary":"Space permission added","description":"","category":"Permissions","sysAdmin":true,"superAdmin":false,"affectedObject":{"name":"Chat Notifications","objectType":"User"},"changedValues":[{"name":"Type","oldValue":"","newValue":"SETSPACEPERMISSIONS","hiddenOldValue":"","hiddenNewValue":""},{"name":"Space","oldValue":"","newValue":"~63bed71b0913ada386aadbe4","hiddenOldValue":"","hiddenNewValue":""},{"name":"User","oldValue":"","newValue":"Chat Notifications","hiddenOldValue":"","hiddenNewValue":""}],"associatedObjects":[{"name":"Shaked Klein","objectType":"Space"}]}],"start":0,"limit":25,"size":26,"_links":{"base":"https://confluence.your-atlassian.net/wiki","context":"/wiki","self":"https://confluence.your-atlassian.net/wiki/rest/api/audit?endDate=1676298362104&startDate=1676082362104"}}`,
			`{"results":[{"author":{"type":"user","displayName":"Shaked Klein","operations":null,"isExternalCollaborator":false,"username":"63bed71b0913ada386aadbe4","userKey":"63bed71b0913ada386aadbe4","accountId":"63bed71b0913ada386aadbe4","accountType":"","publicName":"Unknown user","externalCollaborator":false},"remoteAddress":"","creationDate":1676297029823,"summary":"Space permission added","description":"","category":"Permissions","sysAdmin":true,"superAdmin":false,"affectedObject":{"name":"Chat Notifications","objectType":"User"},"changedValues":[{"name":"Type","oldValue":"","newValue":"SETSPACEPERMISSIONS","hiddenOldValue":"","hiddenNewValue":""},{"name":"Space","oldValue":"","newValue":"~63bed71b0913ada386aadbe4","hiddenOldValue":"","hiddenNewValue":""},{"name":"User","oldValue":"","newValue":"Chat Notifications","hiddenOldValue":"","hiddenNewValue":""}],"associatedObjects":[{"name":"Shaked Klein","objectType":"Space"}]}],"start":25,"limit":25,"size":26,"_links":{"base":"https://confluence.your-atlassian.net/wiki","context":"/wiki","self":"https://confluence.your-atlassian.net/wiki/rest/api/audit?endDate=1676298362104&startDate=1676082362104"}}`,
		},
	}
	mockClient := tests.NewMockWithFixtureByUrlPath(json)

	api := New(
		"username",
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
