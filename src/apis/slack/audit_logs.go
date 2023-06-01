package slack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type auditLog struct {
	Entries          []auditLogEntry   `json:"entries"`
	ResponseMetadata map[string]string `json:"response_metadata"`
}

type auditLogEntry map[string]interface{}

func (s *Slack) getAuditLogs(from, to time.Time) ([]auditLogEntry, error) {
	auditLogs := make([]auditLogEntry, 0)
	nextPage := ""
	for {
		apiEndpoint := s.getEndpoint(
			from.Unix(),
			to.Unix(),
			nextPage,
		)
		request, _ := http.NewRequest("GET", apiEndpoint, nil)
		request = s.setHeaders(request)

		response, err := s.client.Do(request)
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

		auditLog, err := s.parseResponse(contents)
		if err != nil {
			return nil, err
		}

		auditLogs = append(auditLogs, auditLog.Entries...)

		nextPage = auditLog.ResponseMetadata["next_cursor"]

		if nextPage == "" {
			break
		}
	}

	return auditLogs, nil
}

func (s *Slack) setHeaders(req *http.Request) *http.Request {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.accessToken)
	return req
}

func (s *Slack) parseResponse(response []byte) (*auditLog, error) {
	var auditLog *auditLog
	err := json.Unmarshal(response, &auditLog)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return auditLog, nil
}

func (s *Slack) getEndpoint(from, to int64, nextPage string) string {
	if nextPage != "" {
		nextPage = fmt.Sprintf("&cursor=%s", nextPage)
	}
	return fmt.Sprintf("%s/audit/v1/logs?oldest=%s&latest=%s%s",
		s.baseUrl,
		strconv.Itoa(int(from)),
		strconv.Itoa(int(to)),
		nextPage,
	)
}
