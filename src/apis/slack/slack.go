package slack

import (
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"net/http"
)

type Slack struct {
	accessToken string
	baseUrl     string
	client      httputil.Client
}

func NewFromEnv() integration.API {
	validateEnvVars()
	return New(accessToken, baseUrl, &http.Client{})
}
func New(accessToken, baseUrl string, client httputil.Client) integration.API {
	return &Slack{
		accessToken: accessToken,
		baseUrl:     baseUrl,
		client:      client,
	}
}
func (s *Slack) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	from := timePeriodManager.From()
	to := timePeriodManager.To()

	auditLogs, err := s.getAuditLogs(from, to)
	if err != nil {
		return nil, err
	}

	apiResult := s.collectReports(auditLogs)
	return apiResult, nil
}
func (s *Slack) collectReports(auditLogs []auditLogEntry) integration.APIResult {
	ret := make(integration.APIResult, 0)
	for _, auditLog := range auditLogs {
		ret = append(ret, auditLog)
	}
	return ret
}
