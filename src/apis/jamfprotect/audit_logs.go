package jamfprotect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ReportingResponse struct {
	Data map[string]*ReportingResponseData `json:"data"`
}

type ReportingResponseData struct {
	Items    []*AuditLog `json:"items"`
	PageInfo PageInfo    `json:"pageInfo"`
	TypeName string      `json:"__typename"`
}

type PageInfo struct {
	Next     string `json:"next"`
	TypeName string `json:"__typename"`
}

type AuditLog struct {
	ResourceId string `json:"resourceId"`
	Date       string `json:"date"`
	Args       string `json:"args"`
	Error      string `json:"error"`
	Ips        string `json:"ips"`
	Op         string `json:"op"`
	User       string `json:"user"`
	TypeName   string `json:"__typename"`
}

type ReportDuration struct {
	From string
	To   string
}

func (j *JamfProtect) getAuditLogs(accessToken string, from, to time.Time) ([]*AuditLog, error) {
	auditLogs := make([]*AuditLog, 0)
	nextPage := ""
	for {
		u := j.getEndpoint()
		postData := j.buildQueryData(
			from.Format(time.RFC3339),
			to.Format(time.RFC3339),
			nextPage,
		)
		postDataJson, _ := json.Marshal(postData)

		request, _ := http.NewRequest("POST", u, bytes.NewBuffer(postDataJson))
		request = j.setHeaders(accessToken, request)

		response, err := j.client.Do(request)
		if err != nil {
			return nil, fmt.Errorf("error while getting reports: %s", err)
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status code error: %d", response.StatusCode)
		}

		content, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("error while reading response body: %s", err)
		}

		reportingResponseData, err := j.parseResponse(content)
		if err != nil {
			return nil, err
		}

		auditLogs = append(auditLogs, reportingResponseData.Items...)

		nextPage = reportingResponseData.PageInfo.Next
		if nextPage == "" {
			break
		}
	}

	return auditLogs, nil
}

func (j *JamfProtect) setHeaders(accessToken string, req *http.Request) *http.Request {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	return req
}
func (j *JamfProtect) parseResponse(response []byte) (*ReportingResponseData, error) {
	var reportingResponse ReportingResponse
	err := json.Unmarshal(response, &reportingResponse)
	if err != nil {
		return nil, fmt.Errorf("error while parsing response: %s", err)
	}

	data, ok := reportingResponse.Data["listAuditLogsByDate"]
	if !ok {
		return nil, fmt.Errorf("error while getting reports: %s", response)
	}

	return data, nil
}

func (j *JamfProtect) getEndpoint() string {
	e := fmt.Sprintf("%s%s", j.baseUrl, apiUri)
	return e
}

func (j *JamfProtect) buildQueryData(from, to, nextPage string) map[string]interface{} {
	condition := map[string]map[string]string{
		"dateRange": {
			"startDate": from,
			"endDate":   to,
		},
	}
	variables := map[string]interface{}{
		"condition": condition,
	}
	if nextPage != "" {
		variables["next"] = nextPage
	}

	postData := map[string]interface{}{
		"query":         query,
		"operationName": "listAuditLogsByDate",
		"variables":     variables,
	}
	return postData
}

func (j *JamfProtect) isLastPage(data ReportingResponseData) bool {
	return data.PageInfo.Next == ""
}
