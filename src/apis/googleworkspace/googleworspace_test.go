package googleworkspace

import (
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	reports "google.golang.org/api/admin/reports/v1"
	"testing"
	"time"
)

type reportsServiceMemory struct {
	ret map[string]*reports.Activities
}

func (g *reportsServiceMemory) GetActivities(applicationName string, from, to time.Time) (*reports.Activities, error) {
	return g.ret[applicationName], nil
}

func TestGoogleWorkspace_GetReports(t *testing.T) {
	loginEmail := "login@test.com"
	samlEmail := "saml@test.com"
	logTypes := []string{"saml", "login"}
	timePeriodManager := integration.NewTimePeriodManagerFromNow(1)
	activities := map[string]*reports.Activities{}
	activities["saml"] = &reports.Activities{
		Items: []*reports.Activity{
			{
				Actor: &reports.ActivityActor{
					Email: samlEmail,
				},
			},
		},
	}
	activities["login"] = &reports.Activities{
		Items: []*reports.Activity{
			{
				Actor: &reports.ActivityActor{
					Email: loginEmail,
				},
				Events: []*reports.ActivityEvents{
					{
						Type: "test_type",
						Name: "test_name",
						Parameters: []*reports.ActivityEventsParameters{
							{
								Name:  "test",
								Value: "test-value",
							},
							{
								Name:  "blah",
								Value: "blah_value",
							},
						},
					},
				},
			},
		},
	}
	svc := &reportsServiceMemory{
		ret: activities,
	}
	ignoredAuditParameters = []string{"test"}
	gw := New(logTypes, ignoredAuditParameters, svc)
	auditLogs, err := gw.GetReports(timePeriodManager)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(auditLogs) != 2 {
		t.Errorf("Expected 2 audit log, got %d", len(auditLogs))
	}

	if auditLogs[0]["ApplicationName"] != "saml" {
		t.Errorf("Expected ApplicationName to be 'saml', got %v", auditLogs[0]["ApplicationName"])
	}

	if auditLogs[0]["ActorEmail"] != samlEmail {
		t.Errorf("Expected ActorEmail to be '%s', got %s", samlEmail, auditLogs[0]["ActorEmail"])
	}

	if auditLogs[1]["ApplicationName"] != "login" {
		t.Errorf("Expected ApplicationName to be 'login', got %v", auditLogs[1]["ApplicationName"])
	}

	events := auditLogs[1]["Events"].([]map[string]interface{})
	if events[0]["Type"] != "test_type" {
		t.Errorf("Expected Type to be 'test_type', got %v", events[0]["Type"])
	}

	var ok bool
	_, ok = events[0]["Parameters"].(map[string]interface{})["blah"]
	if !ok {
		t.Errorf("Expected blah to be in the parameters, got %v", events[0]["Parameters"])
	}

	_, ok = events[0]["Parameters"].(map[string]interface{})["test"]
	if ok {
		t.Errorf("Expected test to be ignored, got %v", events[0])
	}
}
