package zoom

import (
	"fmt"
	"github.com/coralogix/c4c-ir-integrations/src/httputil"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"net/http"
)

type Zoom struct {
	accountId      string
	clientId       string
	clientSecret   string
	baseUrl        string
	accessTokenUrl string
	client         httputil.Client
}

func NewFromEnv() integration.API {
	validateEnvVars()
	return New(accountId, clientId, clientSecret, baseUrl, accessTokenUrl, &http.Client{})
}

func New(accountId, clientId, clientSecret, baseUrl, accessTokenUrl string, client httputil.Client) integration.API {
	return &Zoom{
		accountId:      accountId,
		clientId:       clientId,
		clientSecret:   clientSecret,
		baseUrl:        baseUrl,
		accessTokenUrl: accessTokenUrl,
		client:         client,
	}
}

func (z *Zoom) GetReports(timePeriodManager *integration.TimePeriodManager) (integration.APIResult, error) {
	fromDate := timePeriodManager.From()
	toDate := timePeriodManager.To()

	tokenResponse, err := z.getAccessToken()
	if err != nil {
		return nil, fmt.Errorf("could not get access token: %w", err)
	}

	auditLogs, err := z.getAuditLogs(tokenResponse, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	apiResult := z.collectReports(auditLogs)
	return apiResult, nil
}
