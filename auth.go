package ticktick

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	// OAuth endpoints
	AuthURL  = "https://ticktick.com/oauth/authorize"
	TokenURL = "https://ticktick.com/oauth/token"
)

// OAuthConfig represents OAuth2 configuration
type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scope        string
}

// TokenResponse represents the OAuth token response
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

// GetAuthorizationURL generates the OAuth authorization URL
func (c *OAuthConfig) GetAuthorizationURL(state string) string {
	params := url.Values{}
	params.Add("client_id", c.ClientID)
	params.Add("redirect_uri", c.RedirectURI)
	params.Add("response_type", "code")
	params.Add("scope", c.Scope)
	params.Add("state", state)

	return AuthURL + "?" + params.Encode()
}

// ExchangeCode exchanges an authorization code for an access token
func (c *OAuthConfig) ExchangeCode(code string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", c.ClientID)
	data.Set("client_secret", c.ClientSecret)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", c.RedirectURI)
	data.Set("scope", c.Scope)

	return c.requestToken(data)
}

// RefreshToken refreshes an access token using a refresh token
func (c *OAuthConfig) RefreshToken(refreshToken string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", c.ClientID)
	data.Set("client_secret", c.ClientSecret)
	data.Set("refresh_token", refreshToken)
	data.Set("grant_type", "refresh_token")

	return c.requestToken(data)
}

// requestToken performs the token request
func (c *OAuthConfig) requestToken(data url.Values) (*TokenResponse, error) {
	req, err := http.NewRequest("POST", TokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}

	return &tokenResp, nil
}
