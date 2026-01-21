package ticktick

import "time"

// Task represents a TickTick task
type Task struct {
	ID            string     `json:"id,omitempty"`
	ProjectID     string     `json:"projectId"`
	Title         string     `json:"title"`
	Content       string     `json:"content,omitempty"`
	Desc          string     `json:"desc,omitempty"`
	AllDay        bool       `json:"allDay,omitempty"`
	StartDate     *time.Time `json:"startDate,omitempty"`
	DueDate       *time.Time `json:"dueDate,omitempty"`
	TimeZone      string     `json:"timeZone,omitempty"`
	IsFloating    bool       `json:"isFloating,omitempty"`
	Reminder      string     `json:"reminder,omitempty"`
	Reminders     []Reminder `json:"reminders,omitempty"`
	Priority      int        `json:"priority,omitempty"`
	Status        int        `json:"status,omitempty"`
	CompletedTime *time.Time `json:"completedTime,omitempty"`
	SortOrder     int64      `json:"sortOrder,omitempty"`
	Items         []TaskItem `json:"items,omitempty"`
	Progress      int        `json:"progress,omitempty"`
	ModifiedTime  *time.Time `json:"modifiedTime,omitempty"`
	Etag          string     `json:"etag,omitempty"`
	Deleted       int        `json:"deleted,omitempty"`
	CreatedTime   *time.Time `json:"createdTime,omitempty"`
	Creator       int64      `json:"creator,omitempty"`
	Tags          []string   `json:"tags,omitempty"`
	Kind          string     `json:"kind,omitempty"`
}

// TaskItem represents a checklist item within a task
type TaskItem struct {
	ID          string     `json:"id,omitempty"`
	Title       string     `json:"title"`
	Status      int        `json:"status,omitempty"`
	CompletedAt *time.Time `json:"completedTime,omitempty"`
	SortOrder   int64      `json:"sortOrder,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	IsAllDay    bool       `json:"isAllDay,omitempty"`
	TimeZone    string     `json:"timeZone,omitempty"`
}

// Reminder represents a task reminder
type Reminder struct {
	ID      string `json:"id,omitempty"`
	Trigger string `json:"trigger"`
}

// Project represents a TickTick project (list)
type Project struct {
	ID           string     `json:"id,omitempty"`
	Name         string     `json:"name"`
	Color        string     `json:"color,omitempty"`
	InAll        bool       `json:"inAll,omitempty"`
	SortOrder    int64      `json:"sortOrder,omitempty"`
	SortType     string     `json:"sortType,omitempty"`
	UserCount    int        `json:"userCount,omitempty"`
	Etag         string     `json:"etag,omitempty"`
	ModifiedTime *time.Time `json:"modifiedTime,omitempty"`
	Closed       bool       `json:"closed,omitempty"`
	Muted        bool       `json:"muted,omitempty"`
	Kind         string     `json:"kind,omitempty"`
}

// CreateTaskRequest represents a request to create a new task
type CreateTaskRequest struct {
	Title     string     `json:"title"`
	ProjectID string     `json:"projectId,omitempty"`
	Content   string     `json:"content,omitempty"`
	Desc      string     `json:"desc,omitempty"`
	AllDay    bool       `json:"allDay,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	DueDate   *time.Time `json:"dueDate,omitempty"`
	TimeZone  string     `json:"timeZone,omitempty"`
	Reminders []Reminder `json:"reminders,omitempty"`
	Priority  int        `json:"priority,omitempty"`
	Tags      []string   `json:"tags,omitempty"`
}

// UpdateTaskRequest represents a request to update a task
type UpdateTaskRequest struct {
	ID        string     `json:"id"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content,omitempty"`
	Desc      string     `json:"desc,omitempty"`
	AllDay    bool       `json:"allDay,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	DueDate   *time.Time `json:"dueDate,omitempty"`
	TimeZone  string     `json:"timeZone,omitempty"`
	Reminders []Reminder `json:"reminders,omitempty"`
	Priority  int        `json:"priority,omitempty"`
	Status    int        `json:"status,omitempty"`
	Tags      []string   `json:"tags,omitempty"`
}

// CreateProjectRequest represents a request to create a new project
type CreateProjectRequest struct {
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

// UpdateProjectRequest represents a request to update a project
type UpdateProjectRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}
