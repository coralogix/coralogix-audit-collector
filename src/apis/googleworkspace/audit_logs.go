package googleworkspace

import (
	"fmt"
	"github.com/sirupsen/logrus"
	reports "google.golang.org/api/admin/reports/v1"
	"strconv"
	"time"
)

type AuditLog map[string]interface{}

func (g *GoogleWorkspace) getAuditLogs(from, to time.Time) ([]AuditLog, error) {
	var auditLogs []AuditLog
	for _, logType := range g.logTypes {
		ret, err := g.getAuditLogsForApp(logType, from, to)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve audit logs for app: %s, err: %s", logType, err)
		}

		logrus.Debugf("Retrieved %d audit logs for app: %s", len(ret), logType)
		auditLogs = append(auditLogs, ret...)
	}

	return auditLogs, nil
}

func (g *GoogleWorkspace) getAuditLogsForApp(applicationName string, from, to time.Time) ([]AuditLog, error) {
	r, err := g.svc.GetActivities(applicationName, from, to)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve activities. %v", err)
	}

	var auditLogs []AuditLog

	for _, activity := range r.Items {
		auditLog := AuditLog{}
		auditLog["ApplicationName"] = applicationName
		auditLog["ActorEmail"] = activity.Actor.Email
		auditLog["ActorProfileId"] = activity.Actor.ProfileId

		var events []map[string]interface{}

		for _, event := range activity.Events {
			collected := g.collectEventParameters(event)
			events = append(events, collected)
		}
		auditLog["Events"] = events
		auditLogs = append(auditLogs, auditLog)
	}

	return auditLogs, nil
}

func (g *GoogleWorkspace) collectEventParameters(event *reports.ActivityEvents) map[string]interface{} {
	parameters := event.Parameters

	normalizedParameters := map[string]interface{}{}
	for _, parameter := range parameters {
		if g.isParameterIgnored(parameter.Name) {
			continue
		}

		var value interface{}
		if parameter.Value != "" {
			value = parameter.Value
		} else if parameter.IntValue != 0 {
			value = strconv.Itoa(int(parameter.IntValue))
		} else if parameter.BoolValue {
			value = strconv.FormatBool(parameter.BoolValue)
		} else if len(parameter.MultiIntValue) > 0 {
			t := []string{}
			for _, v := range parameter.MultiIntValue {
				t = append(t, strconv.Itoa(int(v)))
			}
			value = t
		} else if len(parameter.MultiValue) > 0 {
			t := []string{}
			for _, v := range parameter.MultiValue {
				t = append(t, v)
			}
			value = t
		} else {
			continue
		}

		normalizedParameters[parameter.Name] = value
	}

	collected := map[string]interface{}{
		"Name":       event.Name,
		"Type":       event.Type,
		"Parameters": normalizedParameters,
	}
	return collected
}

func (g *GoogleWorkspace) isParameterIgnored(name string) bool {
	for _, ignoredParameter := range g.ignoredAuditParameters {
		if name == ignoredParameter {
			return true
		}
	}
	return false
}
