package ticktick

import "fmt"

// APIError represents an error from the TickTick API
type APIError struct {
	StatusCode int
	Message    string
}

// Error implements the error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("TickTick API error (status %d): %s", e.StatusCode, e.Message)
}

// IsNotFound checks if the error is a 404 Not Found error
func (e *APIError) IsNotFound() bool {
	return e.StatusCode == 404
}

// IsUnauthorized checks if the error is a 401 Unauthorized error
func (e *APIError) IsUnauthorized() bool {
	return e.StatusCode == 401
}

// IsForbidden checks if the error is a 403 Forbidden error
func (e *APIError) IsForbidden() bool {
	return e.StatusCode == 403
}

// IsRateLimited checks if the error is a 429 Too Many Requests error
func (e *APIError) IsRateLimited() bool {
	return e.StatusCode == 429
}

// IsServerError checks if the error is a 5xx server error
func (e *APIError) IsServerError() bool {
	return e.StatusCode >= 500 && e.StatusCode < 600
}
