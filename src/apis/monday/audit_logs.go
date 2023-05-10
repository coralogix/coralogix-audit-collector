package monday

import (
	"encoding/json"
	"fmt"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type AuditLogRequest struct {
	UserId    string `json:"user_id"`
	Event     string `json:"event"`
	IpAddress string `json:"ip_address"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type ReportingResponse struct {
	Data     []*AuditLog `json:"data,omitempty"`
	Page     int         `json:"page,omitempty"`
	PerPage  int         `json:"per_page,omitempty"`
	NextPage int         `json:"next_page,omitempty"`
	Message  string      `json:"message,omitempty"`
}

type AuditLog struct {
	Timestamp        string      `json:"timestamp"`
	AccountID        int         `json:"account_id"`
	UserId           int         `json:"user_id"`
	Event            string      `json:"event"`
	Slug             string      `json:"slug"`
	IPAddress        string      `json:"ip_address"`
	UserAgent        string      `json:"user_agent"`
	ClientName       string      `json:"client_name"`
	ClientVersion    string      `json:"client_version"`
	OsName           string      `json:"os_name"`
	OsVersion        string      `json:"os_version"`
	DeviceName       string      `json:"device_name"`
	DeviceType       string      `json:"device_type"`
	ActivityMetadata interface{} `json:"activity_metadata,omitempty"`
}

func (m *Monday) getAuditLogs(from, to time.Time) ([]*AuditLog, error) {
	auditLogs := make([]*AuditLog, 0)
	nextPage := 0
	for {
		u := m.getEndpoint(
			nextPage,
			from.Format(timeFormat),
		)

		request, _ := http.NewRequest("GET", u, nil)
		request = m.setHeaders(request)

		response, err := m.client.Do(request)
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

		reportingResponse, err := m.arseResponse(contents)
		if err != nil {
			return nil, err
		}

		auditLogs = append(auditLogs, reportingResponse.Data...)

		if reportingResponse.NextPage == 0 {
			break
		}

		nextPage = reportingResponse.NextPage
	}

	return auditLogs, nil
}

func (m *Monday) setHeaders(req *http.Request) *http.Request {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+m.apiToken)
	return req
}

func (m *Monday) arseResponse(response []byte) (*ReportingResponse, error) {
	var reportingResponse ReportingResponse
	err := json.Unmarshal(response, &reportingResponse)
	if err != nil {
		return nil, err
	}

	if reportingResponse.Message != "" {
		return nil, fmt.Errorf(reportingResponse.Message)
	}

	return &reportingResponse, nil
}

func (m *Monday) getEndpoint(pageNumber int, from string) string {
	filters := map[string]string{
		"start_time": from,
	}
	filterJson, _ := json.Marshal(filters)
	endpoint := fmt.Sprintf(apiEndpoint, m.baseUrl)
	u, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	q := u.Query()

	q.Set("page", strconv.Itoa(pageNumber))
	q.Set("per_page", strconv.Itoa(resultsPerPage))
	q.Set("filters", string(filterJson))
	u.RawQuery = q.Encode()

	return u.String()
}

func (l *Monday) collectReports(auditLogs []*AuditLog) integration.APIResult {
	ret := make(integration.APIResult, 0)
	for _, v := range auditLogs {
		m := make(map[string]interface{})
		m["timestamp"] = v.Timestamp
		m["account_id"] = strconv.Itoa(v.AccountID)
		m["user_id"] = strconv.Itoa(v.UserId)
		m["user_page"] = fmt.Sprintf(userPage, l.baseUrl, strconv.Itoa(v.UserId))
		m["event"] = v.Event
		m["slug"] = v.Slug
		m["ip_address"] = v.IPAddress
		m["user_agent"] = v.UserAgent
		m["client_name"] = v.ClientName
		m["client_version"] = v.ClientVersion
		m["os_name"] = v.OsName
		m["os_version"] = v.OsVersion
		m["device_name"] = v.DeviceName
		m["device_type"] = v.DeviceType
		m["activity_metadata"] = v.ActivityMetadata
		ret = append(ret, m)
	}
	return ret
}
