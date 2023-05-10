package googleworkspace

import (
	reports "google.golang.org/api/admin/reports/v1"
	"log"
	"time"
)

type reportsService interface {
	GetActivities(applicationName string, from, to time.Time) (*reports.Activities, error)
}

type reportsServiceWrapper struct {
	svc *reports.Service
}

func (g *reportsServiceWrapper) GetActivities(applicationName string, from, to time.Time) (*reports.Activities, error) {
	opt := g.svc.Activities.List("all", applicationName)
	opt.StartTime(from.Format(time.RFC3339))
	opt.EndTime(to.Format(time.RFC3339))

	r, err := opt.Do()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func newGoogleWorkspaceReportsServiceWrapper(targetPrincipal, rawJsonKey, googleApplicationCredentials, impersonateUserEmail string) reportsService {
	svc, err := createService(targetPrincipal, rawJsonKey, googleApplicationCredentials, impersonateUserEmail)
	if err != nil {
		log.Fatalf("Unable to create service: %v", err)
	}
	return &reportsServiceWrapper{svc: svc}
}
