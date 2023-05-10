package jira

import (
	"encoding/json"
	"fmt"
	"github.com/coralogix/c4c-ir-integrations/src/integration"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type AuditLog struct {
	Records       []Record `json:"records"`
	Offset        int      `json:"offset"`
	Limit         int      `json:"limit"`
	Total         int      `json:"total"`
	ErrorMessages []string `json:"errorMessages,omitempty"`
}

type Record struct {
	Id              int64            `json:"id"`
	Summary         string           `json:"summary"`
	RemoteAddress   string           `json:"remoteAddress"`
	AuthorKey       string           `json:"authorKey"`
	AuthorAccountId string           `json:"authorAccountId"`
	Created         string           `json:"created"`
	Category        string           `json:"category"`
	EventSource     string           `json:"eventSource"`
	Description     string           `json:"description"`
	ObjectItem      AssociatedItem   `json:"objectItem"`
	ChangedValues   []ChangedValues  `json:"changedValues"`
	AssociatedItems []AssociatedItem `json:"associatedItems"`
}

type AssociatedItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ParentId   string `json:"parentId"`
	ParentName string `json:"parentName"`
	TypeName   string `json:"typeName"`
}

type ChangedValues struct {
	ChangedFrom string `json:"changedFrom"`
	ChangedTo   string `json:"changedTo"`
	FieldName   string `json:"fieldName"`
}

func (j *Jira) getAuditLogs(fromDate, toDate time.Time) (*AuditLog, error) {
	auditLogs := &AuditLog{}
	currentPageNumber := 0
	for {
		apiEndpoint := j.getEndpoint(
			fromDate.Format(timeFormat),
			toDate.Format(timeFormat),
			currentPageNumber,
		)
		request, _ := http.NewRequest("GET", apiEndpoint, nil)
		request = j.setHeaders(request)
		response, err := j.client.Do(request)
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

		logrus.Debugf("response: %v", string(contents))

		parsed, err := j.parseResponse(contents)
		if err != nil {
			logrus.Debugf("Error parsing response: %v", err)
			break
		}

		auditLogs.Records = append(auditLogs.Records, parsed.Records...)

		if j.isLastPage(currentPageNumber, parsed) {
			break
		}

		currentPageNumber += parsed.Limit
	}

	return auditLogs, nil
}

func (j *Jira) parseResponse(response []byte) (*AuditLog, error) {
	var reportingResponse AuditLog
	err := json.Unmarshal(response, &reportingResponse)
	if err != nil {
		return nil, err
	}

	if len(reportingResponse.ErrorMessages) > 0 {
		return nil, fmt.Errorf("error while getting reports: %v", reportingResponse.ErrorMessages)
	}

	return &reportingResponse, nil
}
func (j *Jira) getEndpoint(fromDate, toDate string, currentPageNumber int) string {
	e := fmt.Sprintf(apiEndpoint, j.baseUrl, fromDate, toDate)
	if currentPageNumber > 0 {
		e += fmt.Sprintf("&offset=%d", currentPageNumber)
	}
	return e
}

func (j *Jira) isLastPage(currentPageNumber int, reportingResponse *AuditLog) bool {
	return currentPageNumber+reportingResponse.Limit >= reportingResponse.Total
}

func (j *Jira) collectReports(auditLog *AuditLog, usernameToUserMap, accountIdsToMap map[string]*JiraUser) integration.APIResult {
	ret := make(integration.APIResult, 0)
	for _, record := range auditLog.Records {
		m := make(map[string]interface{})
		m["id"] = fmt.Sprintf("%d", record.Id)
		m["summary"] = record.Summary
		m["remoteAddress"] = record.RemoteAddress
		m["authorKey"] = record.AuthorKey
		m["authorAccountId"] = record.AuthorAccountId
		if record.AuthorAccountId != "" {
			if user, ok := accountIdsToMap[record.AuthorAccountId]; ok {
				m["authorUser"] = user
			}
		}
		m["created"] = record.Created
		m["category"] = record.Category
		m["eventSource"] = record.EventSource
		m["description"] = record.Description
		m["changedValues"] = record.ChangedValues
		associatedItems := map[string]interface{}{}
		for _, item := range record.AssociatedItems {
			associatedItems["Id"] = item.Id
			associatedItems["Name"] = item.Name
			associatedItems["ParentId"] = item.ParentId
			associatedItems["ParentName"] = item.ParentName
			associatedItems["TypeName"] = item.TypeName
			if item.Id != "" {
				if user, ok := usernameToUserMap[item.Id]; ok {
					associatedItems["User"] = user
				}
			}
		}
		m["associatedItems"] = associatedItems
		m["ObjectItemId"] = record.ObjectItem.Id
		m["ObjectItemName"] = record.ObjectItem.Name
		m["ObjectItemParentId"] = record.ObjectItem.ParentId
		ret = append(ret, m)
	}
	return ret
}
