package monday

import (
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"net/http"
)

type Monday struct {
	baseUrl  string
	apiToken string
	client   httputil.Client
}

func NewFromEnv() integration.API {
	validateEnvVars()
	return New(baseUrl, apiToken, &http.Client{})
}

func New(baseUrl, apiToken string, client httputil.Client) integration.API {
	return &Monday{
		baseUrl:  baseUrl,
		apiToken: apiToken,
		client:   client,
	}
}

func (m *Monday) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	from := timePeriodManager.From()
	to := timePeriodManager.To()

	auditLogs, err := m.getAuditLogs(from, to)
	if err != nil {
		return nil, err
	}

	apiResult := m.collectReports(auditLogs)
	return apiResult, nil
}
