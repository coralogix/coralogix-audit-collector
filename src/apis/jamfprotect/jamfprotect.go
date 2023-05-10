package jamfprotect

import (
	"encoding/json"
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/sirupsen/logrus"
	"net/http"
)

type JamfProtect struct {
	clientId string
	apiToken string
	baseUrl  string
	client   httputil.Client
}

func NewFromEnv() integration.API {
	validateEnvVars()
	return New(clientId, apiToken, baseUrl, &http.Client{})
}

func New(clientId, apiToken, baseUrl string, client httputil.Client) integration.API {
	return &JamfProtect{
		clientId: clientId,
		apiToken: apiToken,
		baseUrl:  baseUrl,
		client:   client,
	}
}

func (j *JamfProtect) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	from := timePeriodManager.From()
	to := timePeriodManager.To()

	accessToken, err := j.getAccessToken()
	if err != nil {
		return nil, err
	}

	auditLogs, err := j.getAuditLogs(accessToken, from, to)
	if err != nil {
		return nil, err
	}

	apiResult := j.collectReports(auditLogs)
	return apiResult, nil
}

func (j *JamfProtect) collectReports(auditLogs []*AuditLog) integration.APIResult {
	ret := make(integration.APIResult, 0)
	for _, record := range auditLogs {
		m := make(map[string]interface{})
		m["resourceId"] = record.ResourceId
		m["date"] = record.Date
		var args map[string]interface{}
		err := json.Unmarshal([]byte(record.Args), &args)
		if err != nil {
			logrus.Debugf("Error while unmarshalling args: %s", err.Error())
		} else {
			m["args"] = args
		}
		m["error"] = record.Error
		m["ips"] = record.Ips
		m["op"] = record.Op
		m["user"] = record.User
		m["__typename"] = record.TypeName
		ret = append(ret, m)
	}
	return ret
}
