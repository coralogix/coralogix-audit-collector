package jira

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BulkMigrationResponse struct {
	Username      string   `json:"username"`
	AccountId     string   `json:"accountId"`
	ErrorMessages []string `json:"errorMessages,omitempty"`
	Errors        []string `json:"errors,omitempty"`
}

func (j *Jira) GetAccountIdsByUsername(usernames map[string]string) ([]map[string]string, error) {
	usernameQueryParams := j.buildUsernamesQueryParams(usernames)
	u := fmt.Sprintf("%s/rest/api/3/user/bulk/migration?%s", baseUrl, usernameQueryParams)
	req, _ := http.NewRequest("GET", u, nil)
	req = j.setHeaders(req)
	resp, err := j.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result []BulkMigrationResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		var resultErr BulkMigrationResponse
		err = json.Unmarshal(body, &resultErr)
		if err != nil {

		}
		return nil, fmt.Errorf("error while getting account ids: %v", resultErr.ErrorMessages)
	}

	var accountIds []map[string]string
	for _, item := range result {
		accountIds = append(accountIds, map[string]string{
			"username":  item.Username,
			"accountId": item.AccountId,
		})
	}

	return accountIds, nil
}

func (j *Jira) buildUsernamesQueryParams(usernames map[string]string) string {
	queryParams := ""
	for _, username := range usernames {
		queryParams = fmt.Sprintf("%s&username=%s", queryParams, username)
	}
	return queryParams
}
