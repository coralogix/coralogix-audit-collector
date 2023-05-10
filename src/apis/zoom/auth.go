package zoom

import (
	"fmt"
	"io"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpireIn    int    `json:"expire_in"`
	Scope       string `json:"scope"`
	Reason      string `json:"reason,omitempty"`
}

func (z *Zoom) getAccessToken() (*TokenResponse, error) {
	preAuthRequest := z.createPreAuthRequest()
	preAuthResponse, err := z.client.Do(preAuthRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to get pre-auth response: %w", err)
	}
	defer preAuthResponse.Body.Close()
	contents, err := io.ReadAll(preAuthResponse.Body)
	if err != nil {
		return nil, err
	}

	tokenResponse, err := z.parseAuthResponse(contents)
	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}
