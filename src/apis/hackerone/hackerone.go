package hackerone

import (
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"net/http"
)

type HackerOne struct {
	results      []map[string]interface{}
	clientId     string
	clientSecret string
	programId    string
	baseUrl      string
	client       httputil.Client
}

func NewFromEnv() integration.API {
	validateEnvVars()
	return New(clientId, clientSecret, baseUrl, programId, &http.Client{})
}

func New(clientId, clientSecret, baseUrl, programId string, client httputil.Client) integration.API {
	return &HackerOne{
		clientId:     clientId,
		clientSecret: clientSecret,
		baseUrl:      baseUrl,
		programId:    programId,
		client:       client,
	}
}

func (h *HackerOne) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	from := timePeriodManager.From()
	to := timePeriodManager.To()

	auditLogs, err := h.getAuditLogs(from, to)
	if err != nil {
		return nil, err
	}

	apiResult := h.collectReports(auditLogs)
	return apiResult, nil
}

func (h *HackerOne) collectReports(auditLogs []AuditLogs) integration.APIResult {
	ret := make(integration.APIResult, 0)
	for _, auditLog := range auditLogs {
		ret = append(ret, auditLog)
	}
	return ret
}
