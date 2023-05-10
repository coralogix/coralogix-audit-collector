package lastpass

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"io"
	"net/http"
	"time"
)

type Request struct {
	Cid      string `json:"cid"`
	Provhash string `json:"provhash"`
	Apiuser  string `json:"apiuser"`
	Cmd      string `json:"cmd"`
}

type ReportingRequest struct {
	Request
	Data *DateTimeRange `json:"data"`
}

type DateTimeRange struct {
	From string `json:"from"`
	To   string `json:"to"`
	Page string `json:"page,omitempty"`
}

type ReportingResponse struct {
	Status string               `json:"status"`
	Next   string               `json:"next"`
	Data   map[string]*AuditLog `json:"data,omitempty"`
}

type AuditLog struct {
	Time      string `json:"Time"`
	Username  string `json:"Username"`
	IPAddress string `json:"IP_Address"`
	Action    string `json:"Action"`
	Data      string `json:"Data"`
}

func (l *LastPass) getAuditLogs(from, to time.Time) ([]*AuditLog, error) {
	auditLogs := make([]*AuditLog, 0)
	page := ""
	for {
		body := l.getPostData(
			page,
			from.Format(timeFormat),
			to.Format(timeFormat),
		)
		endpoint := fmt.Sprintf(apiEndpoint, l.baseUrl)
		request, _ := http.NewRequest("POST", endpoint, body)
		request = l.setHeaders(request)

		response, err := l.client.Do(request)
		if err != nil {
			return nil, err
		}

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get audit logs, status code: %d", response.StatusCode)
		}

		contents, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		reportingResponse, err := l.parseResponse(contents)
		if err != nil {
			return nil, err
		}

		for _, v := range reportingResponse.Data {
			auditLogs = append(auditLogs, v)
		}

		if reportingResponse.Next == "" {
			break
		}

		page = reportingResponse.Next
	}

	return auditLogs, nil
}

func (l *LastPass) setHeaders(req *http.Request) *http.Request {
	req.Header.Set("Accept", "application/json")
	return req
}

func (l *LastPass) getPostData(page, from, to string) io.Reader {
	d := &DateTimeRange{
		From: from,
		To:   to,
	}

	if page != "" {
		d.Page = page
	}

	reportingRequest := &ReportingRequest{
		Request: Request{
			Cid:      l.cid,
			Apiuser:  l.apiuser,
			Cmd:      "reporting",
			Provhash: l.provhash,
		},
		Data: d,
	}

	jsonm, _ := json.Marshal(reportingRequest)
	return bytes.NewBuffer(jsonm)
}

func (l *LastPass) parseResponse(response []byte) (*ReportingResponse, error) {
	var reportingResponse ReportingResponse
	err := json.Unmarshal(response, &reportingResponse)
	if err != nil {
		return nil, err
	}

	return &reportingResponse, nil
}

func (l *LastPass) collectReports(auditLogs []*AuditLog) integration.APIResult {
	ret := make(integration.APIResult, 0)
	for _, v := range auditLogs {
		ret = append(ret, map[string]interface{}{
			"Time":      v.Time,
			"Username":  v.Username,
			"IPAddress": v.IPAddress,
			"Action":    v.Action,
			"Data":      v.Data,
		})
	}
	return ret
}
