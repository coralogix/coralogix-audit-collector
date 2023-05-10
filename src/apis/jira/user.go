package jira

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JiraUser struct {
	Username     string `json:"username,omitempty"`
	AccountId    string `json:"accountId"`
	AccountType  string `json:"accountType"`
	EmailAddress string `json:"emailAddress"`
	Self         string `json:"self"`
	Active       bool   `json:"active"`
}

func (j *Jira) GetUserByAccountId(accountId string) (*JiraUser, error) {
	u := fmt.Sprintf("%s/rest/api/3/user?accountId=%s", baseUrl, accountId)
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
	var result *JiraUser
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
