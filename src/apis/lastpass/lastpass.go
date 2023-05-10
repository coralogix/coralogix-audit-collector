package lastpass

import (
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"net/http"
)

type LastPass struct {
	baseUrl  string
	cid      string
	provhash string
	apiuser  string
	client   httputil.Client
}

func NewFromEnv() integration.API {
	validateEnvVars()
	return New(baseUrl, cid, provhash, apiuser, &http.Client{})
}

func New(baseUrl, cid, provhash, apiuser string, client httputil.Client) integration.API {
	return &LastPass{
		baseUrl:  baseUrl,
		cid:      cid,
		provhash: provhash,
		apiuser:  apiuser,
		client:   client,
	}
}

func (l *LastPass) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	startFrom := timePeriodManager.From()
	toEnd := timePeriodManager.To()

	auditLogs, err := l.getAuditLogs(startFrom, toEnd)
	if err != nil {
		return nil, err
	}

	apiResult := l.collectReports(auditLogs)
	return apiResult, nil
}
