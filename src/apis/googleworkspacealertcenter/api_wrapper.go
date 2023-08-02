package googleworkspacealertcenter

import (
	alertcenter "google.golang.org/api/alertcenter/v1beta1"
	"log"
)

type service interface {
	List() *alertcenter.AlertsListCall
}

type serviceWrapper struct {
	svc *alertcenter.Service
}

func (g *serviceWrapper) List() *alertcenter.AlertsListCall {
	return g.svc.Alerts.List()
}

func newGoogleWorkspaceServiceWrapper(targetPrincipal, rawJsonKey, impersonateUserEmail string) service {
	svc, err := createService(targetPrincipal, rawJsonKey, impersonateUserEmail)
	if err != nil {
		log.Fatalf("Unable to create service: %v", err)
	}
	return &serviceWrapper{svc: svc}
}
