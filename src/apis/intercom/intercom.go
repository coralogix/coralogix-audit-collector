package intercom

import (
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"net/http"
	"strconv"
)

type Intercom struct {
	accessToken string
	baseUrl     string
	client      httputil.Client
}

func NewFromEnv() integration.API {
	validateEnvVars()
	return New(accessToken, baseUrl, &http.Client{})
}

func New(accessToken string, baseUrl string, client httputil.Client) integration.API {
	return &Intercom{
		accessToken: accessToken,
		baseUrl:     baseUrl,
		client:      client,
	}
}

func (i *Intercom) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	from := timePeriodManager.From()
	to := timePeriodManager.To()

	auditLogs, err := i.getAuditLogs(from, to)
	if err != nil {
		return nil, err
	}

	apiResult := i.collectReports(auditLogs)
	return apiResult, nil
}

func (l *Intercom) collectReports(auditLogs []*AuditLog) integration.APIResult {
	var ret []map[string]interface{}
	for _, v := range auditLogs {
		m := make(map[string]interface{})
		m["activity_type"] = v.ActivityType
		m["activity_description"] = v.ActivityDescription
		m["created_at"] = strconv.FormatInt(v.CreatedAt, 10)
		m["id"] = v.Id
		m["metadata"] = v.Metadata
		m["performed_by"] = v.PerformedBy
		ret = append(ret, m)
	}
	return ret
}
