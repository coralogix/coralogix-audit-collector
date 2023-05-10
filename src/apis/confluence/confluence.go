package confluence

import "C"
import (
	"fmt"
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"net/http"
	"strconv"
)

type Confluence struct {
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
	return &Confluence{
		username: username,
		apiToken: apiToken,
		baseUrl:  baseUrl,
		client:   client,
	}
}

func (c *Confluence) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	from := timePeriodManager.From()
	to := timePeriodManager.To()
	auditLogs, err := c.getAuditLogs(from, to)
	if err != nil {
		return nil, err
	}
	apiResult := c.collectReports(auditLogs)
	return apiResult, nil
}

func (l *Confluence) collectReports(auditLogs []*AuditLog) integration.APIResult {
	ret := make(integration.APIResult, 0)
	for _, record := range auditLogs {
		m := make(map[string]interface{})
		m["author"] = record.Author.DisplayName
		m["remoteAddress"] = record.RemoteAddress
		m["creationDate"] = strconv.Itoa(int(record.CreationDate))
		m["summary"] = record.Summary
		m["description"] = record.Description
		m["category"] = record.Category
		m["sysAdmin"] = strconv.FormatBool(record.SysAdmin)
		m["superAdmin"] = strconv.FormatBool(record.SuperAdmin)
		m["affectedObjectName"] = record.AffectedObject.Name
		m["affectedObjectType"] = record.AffectedObject.ObjectType
		m["changedValues"] = fmt.Sprintf("%v", record.ChangedValues)
		m["associatedObjects"] = fmt.Sprintf("%v", record.AssociatedObjects)
		ret = append(ret, m)
	}
	return ret
}
