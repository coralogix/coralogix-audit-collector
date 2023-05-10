package zoom

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ReportingResponse struct {
	NextPageToken string           `json:"next_page_token"`
	PageSize      int              `json:"page_size"`
	OperationLogs []*OperationLogs `json:"operation_logs"`
}

type OperationLogs struct {
	Action          string `json:"action"`
	CategoryType    string `json:"category_type"`
	OperationDetail string `json:"operation_detail"`
	Operator        string `json:"operator"`
	Time            string `json:"time"`
}

func (z *Zoom) getAuditLogs(tokenResponse *TokenResponse, fromDate, toDate time.Time) ([]*OperationLogs, error) {
	from := fromDate.Format(time.RFC3339)
	to := toDate.Format(time.RFC3339)
	nextPageToken := ""

	auditLogs := make([]*OperationLogs, 0)
	for {
		u := z.getUrl(from, to, nextPageToken)
		request, _ := http.NewRequest("GET", u, nil)
		request = z.setHeaders(request, tokenResponse)

		response, err := z.client.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get reports: %s", response.Status)
		}

		contents, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		reportingResponse, err := z.parseResponse(contents)
		if err != nil {
			return nil, err
		}

		auditLogs = append(auditLogs, reportingResponse.OperationLogs...)

		if reportingResponse.NextPageToken == "" {
			break
		}

		nextPageToken = reportingResponse.NextPageToken
	}

	return auditLogs, nil
}

func (z *Zoom) setHeaders(req *http.Request, tokenResponse *TokenResponse) *http.Request {
	authHeader := fmt.Sprintf("%s %s", tokenResponse.TokenType, tokenResponse.AccessToken)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Accept", "application/json")
	return req
}

func (z *Zoom) parseResponse(response []byte) (*ReportingResponse, error) {
	var reportingResponse ReportingResponse
	err := json.Unmarshal(response, &reportingResponse)
	if err != nil {
		return nil, err
	}

	return &reportingResponse, nil
}

func (z *Zoom) createPreAuthRequest() *http.Request {
	accessTokenUrl := fmt.Sprintf("%s/oauth/token", z.baseUrl)
	req, _ := http.NewRequest("POST", accessTokenUrl, nil)
	q := req.URL.Query()
	q.Add("grant_type", "account_credentials")
	q.Add("account_id", z.accountId)
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(z.clientId, z.clientSecret)
	return req
}
func (z *Zoom) parseAuthResponse(response []byte) (*TokenResponse, error) {
	var tokenResponse TokenResponse
	err := json.Unmarshal(response, &tokenResponse)
	if err != nil {
		return nil, err
	}

	if tokenResponse.Reason != "" {
		return nil, fmt.Errorf("reason: %s", tokenResponse.Reason)
	}
	if tokenResponse.AccessToken == "" {
		return nil, fmt.Errorf("access_token is empty")
	}

	return &tokenResponse, nil
}

func (z *Zoom) getUrl(from, to, nextPageToken string) string {
	u := fmt.Sprintf(apiEndpoint, z.baseUrl, from, to)
	if nextPageToken != "" {
		u = fmt.Sprintf("%s&next_page_token=%s", u, nextPageToken)
	}
	return u
}

func (z *Zoom) collectReports(operationLogs []*OperationLogs) []map[string]interface{} {
	ret := make([]map[string]interface{}, 0)
	for _, v := range operationLogs {
		m := map[string]interface{}{}
		m["action"] = v.Action
		m["category_type"] = v.CategoryType
		m["operation_detail"] = v.OperationDetail
		m["operator"] = v.Operator
		m["time"] = v.Time
		ret = append(ret, m)
	}
	return ret
}
