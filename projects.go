package ticktick

import (
	"encoding/json"
	"fmt"
)

// ProjectsService handles project-related operations
type ProjectsService struct {
	client *Client
}

// List retrieves all projects
func (s *ProjectsService) List() ([]*Project, error) {
	endpoint := "/project"

	respBody, err := s.client.doRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var projects []*Project
	if err := json.Unmarshal(respBody, &projects); err != nil {
		return nil, fmt.Errorf("failed to parse projects response: %w", err)
	}

	return projects, nil
}

// Get retrieves a specific project by ID
func (s *ProjectsService) Get(projectID string) (*Project, error) {
	endpoint := fmt.Sprintf("/project/%s", projectID)

	respBody, err := s.client.doRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var project Project
	if err := json.Unmarshal(respBody, &project); err != nil {
		return nil, fmt.Errorf("failed to parse project response: %w", err)
	}

	return &project, nil
}

// Create creates a new project
func (s *ProjectsService) Create(req *CreateProjectRequest) (*Project, error) {
	endpoint := "/project"

	respBody, err := s.client.doRequest("POST", endpoint, req)
	if err != nil {
		return nil, err
	}

	var project Project
	if err := json.Unmarshal(respBody, &project); err != nil {
		return nil, fmt.Errorf("failed to parse project response: %w", err)
	}

	return &project, nil
}

// Update updates an existing project
func (s *ProjectsService) Update(req *UpdateProjectRequest) (*Project, error) {
	endpoint := fmt.Sprintf("/project/%s", req.ID)

	respBody, err := s.client.doRequest("POST", endpoint, req)
	if err != nil {
		return nil, err
	}

	var project Project
	if err := json.Unmarshal(respBody, &project); err != nil {
		return nil, fmt.Errorf("failed to parse project response: %w", err)
	}

	return &project, nil
}

// Delete deletes a project
func (s *ProjectsService) Delete(projectID string) error {
	endpoint := fmt.Sprintf("/project/%s", projectID)

	_, err := s.client.doRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	return nil
}
