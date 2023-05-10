package hackerone

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type ReportingResponse struct {
	Data  []AuditLogs       `json:"data"`
	Links map[string]string `json:"links"`
}

type AuditLogs map[string]interface{}

func (h *HackerOne) getAuditLogs(from, to time.Time) ([]AuditLogs, error) {
	auditLogs := make([]AuditLogs, 0)
	nextPageToFetch := ""
	isFetchingFirstPage := true
	for {
		u := h.buildUrl(nextPageToFetch, from, to)
		request, _ := http.NewRequest(
			"GET",
			u,
			nil,
		)
		request = h.setHeaders(request)
		response, err := h.client.Do(request)
		if err != nil {
			return nil, fmt.Errorf("error while sending request: %s", err.Error())
		}

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error while fetching audit logs: %s", response.Status)
		}

		contents, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("error while reading response body: %s", err.Error())
		}

		reportingResponse, err := h.parseResponse(contents)
		if err != nil {
			return nil, err
		}

		if isFetchingFirstPage {
			nextPageToFetch = reportingResponse.Links["last"]
			isFetchingFirstPage = false
			continue
		}

		filteredAuditLogs, areRestOfTheLogsBeforeTheStartingDate, err := h.filterAuditLogs(reportingResponse.Data, from, to)
		if err != nil {
			return nil, fmt.Errorf("error while filtering audit logs: %s", err.Error())
		}

		auditLogs = append(auditLogs, filteredAuditLogs...)

		if areRestOfTheLogsBeforeTheStartingDate {
			break
		}

		prevPage, ok := reportingResponse.Links["prev"]
		if !ok {
			break
		}
		nextPageToFetch = prevPage
	}

	return auditLogs, nil
}

func (h *HackerOne) filterAuditLogs(auditLogs []AuditLogs, from, to time.Time) ([]AuditLogs, bool, error) {
	areRestOfTheLogsBeforeTheStartingDate := false
	data := make([]AuditLogs, 0)
	for _, report := range auditLogs {
		if _, ok := report["attributes"]; !ok {
			logrus.Warnf("attributes not found in report: %v", report)
			continue
		}

		attributes := report["attributes"].(map[string]interface{})
		if _, ok := attributes["created_at"]; !ok {
			logrus.Warnf("created_at not found in report: %v", report)
			continue
		}

		createdAt, err := time.Parse(time.RFC3339, attributes["created_at"].(string))
		if err != nil {
			return nil, false, err
		}

		if from.After(createdAt) {
			areRestOfTheLogsBeforeTheStartingDate = true
			break
		}
		data = append(data, report)
	}

	return data, areRestOfTheLogsBeforeTheStartingDate, nil
}

func (h *HackerOne) buildUrl(nextPage string, from, to time.Time) string {
	if nextPage != "" {
		return nextPage
	}
	return fmt.Sprintf(auditLogAPIEndpoint, h.baseUrl, h.programId)
}

func (h *HackerOne) setHeaders(req *http.Request) *http.Request {
	accessToken := fmt.Sprintf("%s:%s", h.clientId, h.clientSecret)
	accessToken = base64.StdEncoding.EncodeToString([]byte(accessToken))
	req.Header.Set("Authorization", "Basic "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return req
}

func (h *HackerOne) parseResponse(response []byte) (*ReportingResponse, error) {
	var r ReportingResponse
	err := json.Unmarshal(response, &r)
	if err != nil {
		return nil, fmt.Errorf("error while parsing response: %s", err.Error())
	}
	return &r, nil
}
