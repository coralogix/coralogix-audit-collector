package jfrog

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

type AuditLog struct {
	Date            time.Time
	TraceId         string
	UserIp          string
	User            string
	LoggedPrincipal string
	EntityName      string
	EventType       string
	Event           string
	DataChanged     map[string]interface{}
}

func (j *Jfrog) getAuditLogs(from, to time.Time) ([]*AuditLog, error) {
	todaysLogs, err := j.getTodaysLogs(from)
	if err != nil {
		return nil, err
	}
	auditLogs, err := j.getLogsContents(from, todaysLogs)

	if err != nil {
		return nil, err
	}
	return auditLogs, nil
}

func (j *Jfrog) getLogsContents(from time.Time, todaysLogs []string) ([]*AuditLog, error) {
	logsContents := make([]*AuditLog, 0)
	for _, log := range todaysLogs {
		apiEndpoint := fmt.Sprintf("%s/artifactory/jfrog-logs/artifactory/%s/%s", j.baseUrl, from.Format("2006-01-02"), log)
		logrus.Debugf("Getting logs from %s", apiEndpoint)
		request, _ := http.NewRequest("GET", apiEndpoint, nil)
		request = j.setHeaders(request)

		response, err := j.client.Do(request)
		if err != nil {
			return nil, fmt.Errorf("failed to get audit logs: %w", err)
		}

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status code error: %d", response.StatusCode)
		}

		defer response.Body.Close()
		parsed, err := j.parseResponse(response.Body)
		if err != nil {
			return nil, err
		}

		logsContents = append(logsContents, parsed...)
	}
	return logsContents, nil
}
func (j *Jfrog) getTodaysLogs(from time.Time) ([]string, error) {
	apiEndpoint := fmt.Sprintf("%s/artifactory/api/storage/jfrog-logs/artifactory/%s/", j.baseUrl, from.Format("2006-01-02"))
	request, _ := http.NewRequest("GET", apiEndpoint, nil)
	request = j.setHeaders(request)

	response, err := j.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to get audit logs: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d", response.StatusCode)
	}

	var logs map[string]interface{}
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(contents, &logs)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w, %s", err, string(contents))
	}

	// keep only logs that contain the name access-security-audit
	var todaysLogs []string
	children := logs["children"].([]interface{})
	for _, log := range children {
		l := log.(map[string]interface{})
		uri := l["uri"].(string)
		if strings.Contains(uri, "access-security-audit") {
			todaysLogs = append(todaysLogs, uri)
		}
	}

	return todaysLogs, nil
}

func (i *Jfrog) setHeaders(req *http.Request) *http.Request {
	req.SetBasicAuth(i.username, i.apiToken)
	return req
}

func (i *Jfrog) parseResponse(responseBody io.Reader) ([]*AuditLog, error) {
	gzipReader, err := gzip.NewReader(responseBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	// Read and process the data
	data, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read gzip reader: %w", err)
	}

	auditLogs := make([]*AuditLog, 0)
	responseStr := string(data)
	allLogs := strings.Split(responseStr, "\n")
	for i, logRow := range allLogs {
		auditLog := &AuditLog{}
		if logRow == "" {
			logrus.Debugf("Skipping empty log row: %d", i)
			continue
		}
		r := strings.Split(logRow, "|")
		date, err := time.Parse(time.RFC3339, r[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %w", err)
		}

		var dataChanged map[string]interface{}
		trimmed := strings.Trim(r[8], " ")
		err = json.Unmarshal([]byte(trimmed), &dataChanged)
		if err != nil {
			return nil, fmt.Errorf("failed to parse dataChanged: %w %s vs %s", err, r[8], trimmed)
		}

		auditLog.Date = date
		auditLog.TraceId = r[1]
		auditLog.UserIp = r[2]
		auditLog.User = r[3]
		auditLog.LoggedPrincipal = r[4]
		auditLog.EntityName = r[5]
		auditLog.EventType = r[6]
		auditLog.Event = r[7]
		auditLog.DataChanged = dataChanged

		auditLogs = append(auditLogs, auditLog)
	}

	return auditLogs, nil
}
