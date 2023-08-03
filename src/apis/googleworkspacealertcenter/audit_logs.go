package googleworkspacealertcenter

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	alertcenter "google.golang.org/api/alertcenter/v1beta1"
	"strings"
	"time"
)

type AuditLog map[string]interface{}

func (g *GoogleWorkspaceAlertCenter) getAuditLogs(from, to time.Time) ([]AuditLog, error) {
	startTime := from.Format(time.RFC3339)
	endTime := to.Format(time.RFC3339)

	alertsRequest := g.svc.List()
	alertsRequest.PageSize(g.maxPageSize)
	filter := fmt.Sprintf("startTime >= \"%s\" AND startTime < \"%s\"", startTime, endTime)
	alertsRequest.Filter(filter)
	alerts, err := alertsRequest.Do()
	if err != nil {
		return nil, err
	}
	//t, _ := alerts.MarshalJSON()
	//logrus.Fatalf("alerts: %s", string(t))
	logrus.Debugf("Retrieved %d audit logs for filter: %s", len(alerts.Alerts), filter)
	var auditLogs []AuditLog
	for {
		for _, alert := range alerts.Alerts {
			auditLog, err := g.extractAlert(alert)
			if err != nil {
				return nil, err
			}
			data, ok := auditLog["data"]
			if !ok {
				return nil, fmt.Errorf("alert data is nil")
			}
			typeName := data.(map[string]interface{})["@type"].(string)
			auditLog["AlertType"] = strings.Replace(typeName, "type.googleapis.com/google.apps.alertcenter.type.", "", -1)
			auditLogs = append(auditLogs, auditLog)
		}

		if alerts.NextPageToken == "" {
			break
		}

		alertsRequest.PageToken(alerts.NextPageToken)
		alerts, err = alertsRequest.Do()
		if err != nil {
			return nil, err
		}
	}
	logrus.Debugf("Retrieved %d audit logs", len(auditLogs))

	return auditLogs, nil
}

func (g *GoogleWorkspaceAlertCenter) extractAlert(alert *alertcenter.Alert) (AuditLog, error) {
	var alertMap map[string]interface{}
	jsonBytes, err := alert.MarshalJSON()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBytes, &alertMap)
	if err != nil {
		return nil, err
	}
	return alertMap, nil
}
