package intercom

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type ReportingResponse struct {
	Type         string      `json:"type"`
	ActivityLogs []*AuditLog `json:"activity_logs"`
	Pages        Pages       `json:"pages"`
}

type Pages struct {
	Next      string `json:"next,omitempty"`
	Page      int    `json:"page"`
	PerPage   int    `json:"per_page"`
	TotalPage int    `json:"total_pages"`
}

type AuditLog struct {
	ActivityType        string                 `json:"activity_type"`
	ActivityDescription string                 `json:"activity_description"`
	Metadata            map[string]interface{} `json:"metadata"`
	CreatedAt           int64                  `json:"created_at"`
	PerformedBy         Performer              `json:"performed_by"`
	Id                  string                 `json:"id"`
}

type Performer struct {
	Type              string   `json:"type"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	JobTitle          string   `json:"job_title"`
	AwayModeEnabled   bool     `json:"away_mode_enabled"`
	AwayModeReassign  bool     `json:"away_mode_reassign"`
	HasInboxSeat      bool     `json:"has_inbox_seat"`
	TeamIds           []string `json:"team_ids"`
	Avatar            string   `json:"avatar"`
	TeamPriorityLevel string   `json:"team_priority_level"`
}

func (i *Intercom) getAuditLogs(from, to time.Time) ([]*AuditLog, error) {
	auditLogs := make([]*AuditLog, 0)
	nextPage := ""
	for {
		apiEndpoint := i.getEndpoint(
			from.Unix(),
			to.Unix(),
			nextPage,
		)
		request, _ := http.NewRequest("GET", apiEndpoint, nil)
		request = i.setHeaders(request)

		response, err := i.client.Do(request)
		if err != nil {
			return nil, fmt.Errorf("failed to get audit logs: %w", err)
		}

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status code error: %d", response.StatusCode)
		}

		defer response.Body.Close()
		contents, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		reportingResponse, err := i.parseResponse(contents)
		if err != nil {
			return nil, err
		}

		auditLogs = append(auditLogs, reportingResponse.ActivityLogs...)

		nextPage = reportingResponse.Pages.Next

		if nextPage == "" {
			break
		}
	}

	return auditLogs, nil
}

func (i *Intercom) setHeaders(req *http.Request) *http.Request {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+i.accessToken)
	return req
}

func (i *Intercom) parseResponse(response []byte) (*ReportingResponse, error) {
	var reportingResponse ReportingResponse
	err := json.Unmarshal(response, &reportingResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &reportingResponse, nil
}

func (i *Intercom) getEndpoint(from, to int64, nextPage string) string {
	if nextPage != "" {
		return nextPage
	}
	return fmt.Sprintf(apiEndpoint, i.baseUrl, strconv.Itoa(int(from)), strconv.Itoa(int(to)))
}
