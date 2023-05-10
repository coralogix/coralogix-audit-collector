package jamfprotect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AuthResponse struct {
	AccessToken      string `json:"access_token,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

func (j *JamfProtect) getAccessToken() (string, error) {
	u := fmt.Sprintf("%s%s", j.baseUrl, accessTokenUri)
	postData := map[string]string{
		"client_id": j.clientId,
		"password":  j.apiToken,
	}
	postDataJson, _ := json.Marshal(postData)

	request, _ := http.NewRequest("POST", u, bytes.NewBuffer(postDataJson))

	response, err := j.client.Do(request)
	if err != nil {
		return "", fmt.Errorf("error while getting access token: %s", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d", response.StatusCode)
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error while reading response body: %s", err)
	}

	var authResponse AuthResponse
	err = json.Unmarshal(contents, &authResponse)
	if err != nil {
		return "", fmt.Errorf("error while parsing response: %s", err)
	}

	if authResponse.Error != "" {
		return "", fmt.Errorf("error while getting access token: %s (%s)", authResponse.Error, authResponse.ErrorDescription)
	}

	return authResponse.AccessToken, nil
}
