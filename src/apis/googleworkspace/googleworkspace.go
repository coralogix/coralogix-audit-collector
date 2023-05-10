package googleworkspace

import (
	"context"
	"fmt"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	reports "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

type GoogleWorkspace struct {
	results                []map[string]interface{}
	logTypes               []string
	ignoredAuditParameters []string
	svc                    reportsService
}

func createService(targetPrincipal, rawJsonKey, googleApplicationCredentials, impersonateUserEmail string) (*reports.Service, error) {
	var err error
	ctx := context.Background()

	var keyFileBytes []byte
	var tokenSource oauth2.TokenSource

	if rawJsonKey != "" {
		keyFileBytes = []byte(rawJsonKey)
	}
	if keyFileBytes != nil {
		logrus.Debugf("Using raw json key")
		config, err := google.JWTConfigFromJSON(keyFileBytes, reports.AdminReportsAuditReadonlyScope)

		if err != nil {
			return nil, fmt.Errorf("unable to impersonate service account: %v", err)
		}
		config.Subject = impersonateUserEmail
		tokenSource = config.TokenSource(ctx)
	} else {
		ts, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
			TargetPrincipal: targetPrincipal,
			Scopes:          []string{"https://www.googleapis.com/auth/admin.reports.audit.readonly"},
			Subject:         impersonateUserEmail,
		})

		if err != nil {
			return nil, fmt.Errorf("unable to impersonate service account: %v", err)
		}
		logrus.Debugf("Using default credentials")
		tokenSource = ts
	}

	svc, err := reports.NewService(ctx,
		option.WithTokenSource(tokenSource),
	)

	if err != nil {
		return nil, fmt.Errorf("unable to retrieve reports Client %v", err)
	}

	return svc, nil
}

func NewFromEnv() integration.API {
	validateEnvVars()
	svc := newGoogleWorkspaceReportsServiceWrapper(targetPrincipal, rawJsonKey, googleApplicationCredentials, impersonateUserEmail)
	return New(logTypes, ignoredAuditParameters, svc)
}

func New(logTypes, ignoredAuditParameters []string, svc reportsService) integration.API {
	return &GoogleWorkspace{
		logTypes:               logTypes,
		ignoredAuditParameters: ignoredAuditParameters,
		svc:                    svc,
	}
}

func (g *GoogleWorkspace) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	from := timePeriodManager.From()
	to := timePeriodManager.To()

	auditLogs, err := g.getAuditLogs(from, to)
	if err != nil {
		return nil, err
	}

	apiResult := g.collectReports(auditLogs)
	logrus.Debugf("Google Workspace API result: %d records", len(apiResult))
	return apiResult, nil
}

func (g *GoogleWorkspace) collectReports(auditLogs []AuditLog) integration.APIResult {
	ret := make(integration.APIResult, 0)
	for _, auditLog := range auditLogs {
		ret = append(ret, auditLog)
	}
	return ret
}
