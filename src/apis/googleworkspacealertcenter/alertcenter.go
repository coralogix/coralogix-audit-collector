package googleworkspacealertcenter

import (
	"context"
	"fmt"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	alertcenter "google.golang.org/api/alertcenter/v1beta1"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
	"strconv"
)

type GoogleWorkspaceAlertCenter struct {
	svc         service
	maxPageSize int64
}

func createService(targetPrincipal, rawJsonKey, impersonateUserEmail string) (*alertcenter.Service, error) {
	var err error
	ctx := context.Background()

	var keyFileBytes []byte
	var tokenSource oauth2.TokenSource

	if rawJsonKey != "" {
		keyFileBytes = []byte(rawJsonKey)
	}
	if keyFileBytes != nil {
		logrus.Debugf("Using raw json key")
		config, err := google.JWTConfigFromJSON(keyFileBytes, alertcenter.AppsAlertsScope)

		if err != nil {
			return nil, fmt.Errorf("unable to impersonate service account: %v", err)
		}
		config.Subject = impersonateUserEmail
		tokenSource = config.TokenSource(ctx)
	} else {
		ts, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
			TargetPrincipal: targetPrincipal,
			Scopes:          []string{alertcenter.AppsAlertsScope},
			Subject:         impersonateUserEmail,
		})

		if err != nil {
			return nil, fmt.Errorf("unable to impersonate service account: %v", err)
		}
		logrus.Debugf("Using default credentials")
		tokenSource = ts
	}

	svc, err := alertcenter.NewService(ctx,
		option.WithTokenSource(tokenSource),
	)

	if err != nil {
		return nil, fmt.Errorf("unable to retrieve reports Client %v", err)
	}

	return svc, nil
}

func NewFromEnv() integration.API {
	validateEnvVars()
	svc := newGoogleWorkspaceServiceWrapper(targetPrincipal, rawJsonKey, impersonateUserEmail)
	var err error
	maxPageSizeInt := int64(100)
	if maxPageSize != "" {
		maxPageSizeInt, err = strconv.ParseInt(maxPageSize, 10, 64)
		if err != nil {
			logrus.Fatalf("maxPageSize is not a number: %s", maxPageSize)
		}
	}
	return New(svc, maxPageSizeInt)
}

func New(svc service, maxPageSize int64) integration.API {
	return &GoogleWorkspaceAlertCenter{
		svc:         svc,
		maxPageSize: maxPageSize,
	}
}

func (g *GoogleWorkspaceAlertCenter) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	from := timePeriodManager.From()
	to := timePeriodManager.To()
	auditLogs, err := g.getAuditLogs(from, to)
	if err != nil {
		return nil, err
	}

	apiResult := g.collectReports(auditLogs)
	logrus.Debugf("Google Workspace Alerts API result: %d records", len(apiResult))
	return apiResult, nil
}

func (g *GoogleWorkspaceAlertCenter) collectReports(auditLogs []AuditLog) integration.APIResult {
	ret := make(integration.APIResult, 0)
	for _, auditLog := range auditLogs {
		ret = append(ret, auditLog)
	}
	return ret
}
