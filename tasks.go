package ticktick

import (
	"encoding/json"
	"fmt"
)

// TasksService handles task-related operations
type TasksService struct {
	client *Client
}

// List retrieves all tasks for a specific project
func (s *TasksService) List(projectID string) ([]*Task, error) {
	endpoint := fmt.Sprintf("/project/%s/data", projectID)

	respBody, err := s.client.doRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	if err := json.Unmarshal(respBody, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse tasks response: %w", err)
	}

	return tasks, nil
}

// Get retrieves a specific task by ID
func (s *TasksService) Get(projectID, taskID string) (*Task, error) {
	endpoint := fmt.Sprintf("/project/%s/task/%s", projectID, taskID)

	respBody, err := s.client.doRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var task Task
	if err := json.Unmarshal(respBody, &task); err != nil {
		return nil, fmt.Errorf("failed to parse task response: %w", err)
	}

	return &task, nil
}

// Create creates a new task
func (s *TasksService) Create(req *CreateTaskRequest) (*Task, error) {
	endpoint := "/task"

	respBody, err := s.client.doRequest("POST", endpoint, req)
	if err != nil {
		return nil, err
	}

	var task Task
	if err := json.Unmarshal(respBody, &task); err != nil {
		return nil, fmt.Errorf("failed to parse task response: %w", err)
	}

	return &task, nil
}

// Update updates an existing task
func (s *TasksService) Update(projectID string, req *UpdateTaskRequest) (*Task, error) {
	endpoint := fmt.Sprintf("/project/%s/task/%s", projectID, req.ID)

	respBody, err := s.client.doRequest("POST", endpoint, req)
	if err != nil {
		return nil, err
	}

	var task Task
	if err := json.Unmarshal(respBody, &task); err != nil {
		return nil, fmt.Errorf("failed to parse task response: %w", err)
	}

	return &task, nil
}

// Complete marks a task as completed
func (s *TasksService) Complete(projectID, taskID string) (*Task, error) {
	endpoint := fmt.Sprintf("/project/%s/task/%s/complete", projectID, taskID)

	respBody, err := s.client.doRequest("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var task Task
	if err := json.Unmarshal(respBody, &task); err != nil {
		return nil, fmt.Errorf("failed to parse task response: %w", err)
	}

	return &task, nil
}

// Delete deletes a task
func (s *TasksService) Delete(projectID, taskID string) error {
	endpoint := fmt.Sprintf("/project/%s/task/%s", projectID, taskID)

	_, err := s.client.doRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	return nil
}
