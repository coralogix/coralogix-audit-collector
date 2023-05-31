package jfrog

import (
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"net/http"
	"time"
)

type Jfrog struct {
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
	return &Jfrog{
		username: username,
		apiToken: apiToken,
		baseUrl:  baseUrl,
		client:   client,
	}
}

func (j *Jfrog) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	from := timePeriodManager.From()
	to := timePeriodManager.To()

	auditLogs, err := j.getAuditLogs(from, to)
	if err != nil {
		return nil, err
	}

	apiResult := j.collectReports(auditLogs)
	return apiResult, nil
}

func (j *Jfrog) collectReports(auditLogs []*AuditLog) integration.APIResult {
	ret := make(integration.APIResult, 0)
	for _, record := range auditLogs {
		m := make(map[string]interface{})
		m["date"] = record.Date.Format(time.RFC3339)
		m["traceId"] = record.TraceId
		m["userIp"] = record.UserIp
		m["user"] = record.User
		m["loggedPrincipal"] = record.LoggedPrincipal
		m["entityName"] = record.EntityName
		m["eventType"] = record.EventType
		m["event"] = record.Event
		m["dataChanged"] = record.DataChanged
		ret = append(ret, m)
	}
	return ret
}
