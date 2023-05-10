package jira

import (
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const (
	apiEndpoint = "%s/rest/api/3/auditing/record?from=%s&to=%s"
	timeFormat  = "2006-01-02T15:04:05"
)

var (
	username = os.Getenv("JIRA_USERNAME")
	apiToken = os.Getenv("JIRA_API_TOKEN")
	baseUrl  = os.Getenv("BASE_URL")
)

func validateEnvVars() {
	if username == "" {
		logrus.Fatalf("JIRA_USERNAME env var is not set")
	}
	if apiToken == "" {
		logrus.Fatalf("JIRA_API_TOKEN env var is not set")
	}
	if baseUrl == "" {
		logrus.Fatalf("BASE_URL env var is not set")
	}
}

type Jira struct {
	username string
	apiToken string
	baseUrl  string
	client   httputil.Client
}

func NewFromEnv() integration.API {
	validateEnvVars()
	return New(username, apiToken, baseUrl, &http.Client{})
}

func New(username, apiToken, baseUrl string, client httputil.Client) integration.API {
	return &Jira{
		username: username,
		apiToken: apiToken,
		baseUrl:  baseUrl,
		client:   client,
	}
}

func (j *Jira) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	fromDate := timePeriodManager.From()
	toDate := timePeriodManager.To()
	auditLogs, err := j.getAuditLogs(fromDate, toDate)
	if err != nil {
		return nil, err
	}
	usernameToUserMap, accountIdsToMap, err := j.mapUsernamesToUserObjects(auditLogs)
	if err != nil {
		return nil, err
	}
	apiResult := j.collectReports(auditLogs, usernameToUserMap, accountIdsToMap)
	return apiResult, nil
}

func (j *Jira) setHeaders(req *http.Request) *http.Request {
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(j.username, j.apiToken)
	return req
}

func (j *Jira) mapUsernamesToUserObjects(auditLog *AuditLog) (map[string]*JiraUser, map[string]*JiraUser, error) {
	usernames := map[string]string{}
	accountIds := map[string]string{}
	for _, result := range auditLog.Records {
		for _, item := range result.AssociatedItems {
			usernames[item.Id] = item.Id
		}
		accountIds[result.AuthorAccountId] = result.AuthorAccountId
	}

	var err error
	usernameAccountIds := []map[string]string{}
	if len(usernames) > 0 {
		usernameAccountIds, err = j.GetAccountIdsByUsername(usernames)
		if err != nil {
			return nil, nil, err
		}
	}

	usernameToUserMap := map[string]*JiraUser{}
	accountIdsToMap := map[string]*JiraUser{}
	for _, accountId := range usernameAccountIds {
		user, err := j.GetUserByAccountId(accountId["accountId"])
		if err != nil {
			logrus.Warnf("Error getting user by account id: %v", err)
			continue
		}
		accountIdsToMap[accountId["accountId"]] = user
		usernameToUserMap[accountId["username"]] = user
	}

	for _, accountId := range accountIds {
		if _, ok := accountIdsToMap[accountId]; ok {
			continue
		}

		user, err := j.GetUserByAccountId(accountId)
		if err != nil {
			logrus.Warnf("Error getting user by account id: %v", err)
			continue
		}
		accountIdsToMap[accountId] = user
	}

	return usernameToUserMap, accountIdsToMap, nil
}
