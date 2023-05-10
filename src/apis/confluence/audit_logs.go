package confluence

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ReportingResponse struct {
	Results       []*AuditLog `json:"results"`
	Start         int         `json:"start"`
	Limit         int         `json:"limit"`
	Size          int         `json:"size"`
	Links         []string    `json:"links"`
	ErrorMessages []string    `json:"errorMessages,omitempty"`
}
type AuditLog struct {
	Author            Author           `json:"author"`
	RemoteAddress     string           `json:"remoteAddress"`
	CreationDate      int64            `json:"creationDate"`
	Summary           string           `json:"summary"`
	Description       string           `json:"description"`
	Category          string           `json:"category"`
	SysAdmin          bool             `json:"sysAdmin"`
	SuperAdmin        bool             `json:"superAdmin"`
	AffectedObject    AffectedObject   `json:"affectedObject"`
	ChangedValues     []ChangedValues  `json:"changedValues"`
	AssociatedObjects []AffectedObject `json:"associatedObjects"`
}

type Author struct {
	Type                   string `json:"type"`
	DisplayName            string `json:"displayName"`
	Username               string `json:"username"`
	UserKey                string `json:"userKey,omitempty"`
	AccountId              string `json:"accountId,omitempty"`
	AccountType            string `json:"accountType"`
	ExternalCollaborator   bool   `json:"externalCollaborator,omitempty"`
	PublicName             string `json:"publicName,omitempty"`
	IsExternalCollaborator bool   `json:"isExternalCollaborator,omitempty"`
}

type AffectedObject struct {
	Name       string `json:"name"`
	ObjectType string `json:"objectType"`
}

type ChangedValues struct {
	Name           string `json:"name"`
	OldValue       string `json:"oldValue"`
	HiddenOldValue string `json:"hiddenOldValue"`
	NewValue       string `json:"newValue"`
	HiddenNewValue string `json:"hiddenNewValue"`
}

func (c *Confluence) getAuditLogs(from, to time.Time) ([]*AuditLog, error) {
	auditLogs := make([]*AuditLog, 0)
	nextPage := 0
	for {
		apiEndpoint := c.getEndpoint(
			from.UnixMilli(),
			to.UnixMilli(),
			nextPage,
		)
		request, _ := http.NewRequest("GET", apiEndpoint, nil)
		request = c.setHeaders(request)

		response, err := c.client.Do(request)
		if err != nil {
			return nil, fmt.Errorf("error while requesting reports: %v", err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error response: %d", response.StatusCode)
		}

		contents, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("error while reading reports: %v", err)
		}

		reportingResponse, err := c.parseResponse(contents)
		if err != nil {
			return nil, err
		}

		auditLogs = append(auditLogs, reportingResponse.Results...)

		if c.isLastPage(nextPage, reportingResponse) {
			break
		}
		nextPage += reportingResponse.Limit
	}

	return auditLogs, nil
}

func (c *Confluence) setHeaders(req *http.Request) *http.Request {
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.username, c.apiToken)
	return req
}

func (c *Confluence) parseResponse(response []byte) (*ReportingResponse, error) {
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

func (c *Confluence) getEndpoint(from, to int64, nextPage int) string {
	e := fmt.Sprintf(apiEndpoint, c.baseUrl, from, to)
	if nextPage > 0 {
		e += fmt.Sprintf("&start=%d", nextPage)
	}
	return e
}

func (c *Confluence) isLastPage(currentPage int, reportingResponse *ReportingResponse) bool {
	return currentPage+reportingResponse.Limit >= reportingResponse.Size
}
