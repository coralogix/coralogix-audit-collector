package atlassian

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ReportingResponse struct {
	Data          []*AuditLog `json:"data"`
	Meta          Meta        `json:"meta"`
	Links         Links       `json:"links"`
	ErrorMessages []string    `json:"errorMessages,omitempty"`
}
type Meta struct {
	Next     string `json:"next"`
	PageSize int    `json:"pageSize"`
}
type Links struct {
	Next string `json:"next,omitempty"`
	Self string `json:"self,omitempty"`
	Prev string `json:"prev,omitempty"`
}
type AuditLog struct {
	Id         string      `json:"id"`
	Type       string      `json:"type"`
	Attributes Attributes  `json:"attributes"`
	Links      Links       `json:"links"`
	Context    []Context   `json:"context"`
	Container  []Container `json:"container"`
}
type Attributes struct {
	Time   string `json:"time"`
	Action string `json:"action"`
	Actor  Actor  `json:"actor"`
}
type Actor struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Links Links  `json:"links"`
}

type Context struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
	Links      Links      `json:"links"`
}

type Container struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
	Links      Links      `json:"links"`
}

type Location struct {
	Ip          string `json:"ip"`
	Geo         string `json:"geo"`
	CountryName string `json:"countryName"`
	RegionName  string `json:"regionName"`
	City        string `json:"city"`
}

func (a *Atlassian) getAuditLogs(from, to time.Time) ([]*AuditLog, error) {
	auditLogs := make([]*AuditLog, 0)
	nextPage := ""
	for {
		apiEndpoint := a.getEndpoint(
			from.UnixMilli(),
			to.UnixMilli(),
			nextPage,
		)
		request, _ := http.NewRequest("GET", apiEndpoint, nil)
		request = a.setHeaders(request)

		response, err := a.client.Do(request)
		if err != nil {
			return nil, fmt.Errorf("error while requesting reports: %v", err)
		}

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error status code: %d", response.StatusCode)
		}

		contents, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("error while reading response body: %v", err)
		}

		reportingResponse, err := a.parseResponse(contents)
		if err != nil {
			return nil, err
		}

		auditLogs = append(auditLogs, reportingResponse.Data...)

		if a.isLastPage(reportingResponse) {
			break
		}
		nextPage = reportingResponse.Meta.Next
	}

	return auditLogs, nil
}

func (a *Atlassian) setHeaders(req *http.Request) *http.Request {
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(a.clientId, a.apiToken)
	return req
}

func (a *Atlassian) parseResponse(response []byte) (*ReportingResponse, error) {
	var reportingResponse ReportingResponse
	err := json.Unmarshal(response, &reportingResponse)
	if err != nil {
		return nil, err
	}

	if len(reportingResponse.ErrorMessages) > 0 {
		return nil, fmt.Errorf("error while getting reports: %v", reportingResponse.ErrorMessages)
	}

	return &reportingResponse, nil
}

func (a *Atlassian) getEndpoint(from, to int64, nextPage string) string {
	if nextPage != "" {
		return nextPage
	}
	return fmt.Sprintf(apiEndpoint, a.baseUrl, a.orgId, from, to)
}

func (l *Atlassian) isLastPage(reportingResponse *ReportingResponse) bool {
	return reportingResponse.Meta.Next == ""
}
