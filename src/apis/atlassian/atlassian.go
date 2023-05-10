package atlassian

import (
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"net/http"
)

type Atlassian struct {
	clientId string
	apiToken string
	baseUrl  string
	orgId    string
	client   httputil.Client
}

func NewFromEnv() integration.API {
	validateEnvVars()
	return New(clientId, apiToken, orgId, baseUrl, &http.Client{})
}

func New(clientId, apiToken, orgId, baseUrl string, client httputil.Client) integration.API {
	return &Atlassian{
		clientId: clientId,
		apiToken: apiToken,
		orgId:    orgId,
		baseUrl:  baseUrl,
		client:   client,
	}
}

func (a *Atlassian) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	from := timePeriodManager.From()
	to := timePeriodManager.To()
	auditLogs, err := a.getAuditLogs(from, to)
	if err != nil {
		return nil, err
	}

	apiResult := a.collectReports(auditLogs)
	return apiResult, nil
}

func (l *Atlassian) collectReports(auditLogs []*AuditLog) integration.APIResult {
	var ret integration.APIResult
	for _, record := range auditLogs {
		m := make(map[string]interface{})
		m["id"] = record.Id
		m["type"] = record.Type
		m["attributes"] = record.Attributes
		m["links"] = record.Links
		m["context"] = record.Context
		m["container"] = record.Container
		ret = append(ret, m)
	}
	return ret
}
