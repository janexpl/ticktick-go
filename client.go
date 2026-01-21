package ticktick

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// BaseURL is the base URL for TickTick API
	BaseURL = "https://api.ticktick.com/open/v1"
)

// Client represents a TickTick API client
type Client struct {
	httpClient *http.Client
	baseURL    string
	accessToken string

	// Services
	Tasks    *TasksService
	Projects *ProjectsService
}

// NewClient creates a new TickTick API client
func NewClient(accessToken string) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL:     BaseURL,
		accessToken: accessToken,
	}

	client.Tasks = &TasksService{client: client}
	client.Projects = &ProjectsService{client: client}

	return client
}

// SetHTTPClient sets a custom HTTP client
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

// doRequest performs an HTTP request
func (c *Client) doRequest(method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	url := c.baseURL + endpoint
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(respBody),
		}
	}

	return respBody, nil
}
