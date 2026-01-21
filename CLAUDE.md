# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go client library for the TickTick API, providing a structured interface to manage tasks and projects in the TickTick task management application. The library handles OAuth2 authentication and provides typed request/response structures for all API operations.

## Development Commands

### Build and Run
```bash
# Build the library
go build

# Run the example
go run examples/main.go

# Get dependencies
go mod tidy
go mod download
```

### Testing
```bash
# Run tests (when available)
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for a specific package
go test -v ./
```

### Code Quality
```bash
# Format code
go fmt ./...

# Run linter (if golangci-lint is installed)
golangci-lint run

# Vet code
go vet ./...
```

## Architecture

### Core Components

**Client Structure (client.go)**
- `Client` is the main entry point that manages HTTP communication with the TickTick API
- Uses service pattern: `Client.Tasks` and `Client.Projects` are specialized services
- All HTTP requests flow through `Client.doRequest()` which handles authentication, error responses, and JSON marshaling
- Base URL: `https://api.ticktick.com/open/v1`
- Authenticates using Bearer tokens in the Authorization header

**Service Pattern**
- `TasksService` (tasks.go) - handles all task CRUD operations
- `ProjectsService` (projects.go) - handles all project CRUD operations
- Services hold a reference to the parent `Client` and use its `doRequest()` method
- This pattern separates concerns and makes the API intuitive: `client.Tasks.Create()`, `client.Projects.List()`

**Authentication Flow (auth.go)**
- OAuth2 implementation with authorization code flow
- `OAuthConfig` handles authorization URL generation and token exchange
- Supports both initial authorization (`ExchangeCode`) and token refresh (`RefreshToken`)
- Token endpoint: `https://ticktick.com/oauth/token`
- Auth endpoint: `https://ticktick.com/oauth/authorize`

**Type System (types.go)**
- `Task` - comprehensive struct with fields for title, content, dates, reminders, priority, status, tags
- `Project` - represents TickTick lists with name, color, sort order
- `TaskItem` - checklist items within tasks
- `Reminder` - task reminder configuration
- Separate request structs for Create/Update operations (e.g., `CreateTaskRequest`, `UpdateTaskRequest`)

**Error Handling (errors.go)**
- `APIError` wraps HTTP status codes and response messages
- Helper methods: `IsNotFound()`, `IsUnauthorized()`, `IsForbidden()`, `IsRateLimited()`, `IsServerError()`
- Returned by `Client.doRequest()` when HTTP status is outside 200-299 range

### API Request Flow

1. User calls service method (e.g., `client.Tasks.Create(&CreateTaskRequest{...})`)
2. Service constructs endpoint path (e.g., `/task` or `/project/{id}/task/{taskId}`)
3. Service calls `client.doRequest(method, endpoint, body)`
4. `doRequest()` marshals request body to JSON, adds auth headers, executes HTTP request
5. Response is read and checked for errors (non-2xx status codes return `APIError`)
6. Successful response body is returned to service method
7. Service unmarshals JSON response into appropriate struct and returns to caller

### Key Design Patterns

**Request/Response Separation**: Create and Update operations use separate request structs rather than reusing the main model structs, providing clearer API boundaries and avoiding confusion about which fields are settable.

**Error Type Assertions**: Callers can type-assert errors to `*APIError` to access status-specific helpers, enabling clean error handling patterns.

**Service Composition**: The Client doesn't implement task/project methods directly; instead it composes specialized services, making the codebase more maintainable as new resource types are added.

## Important Implementation Notes

### API Endpoint Patterns

Tasks require projectID in most operations:
- List tasks: `GET /project/{projectId}/data`
- Get task: `GET /project/{projectId}/task/{taskId}`
- Update task: `POST /project/{projectId}/task/{taskId}`
- Complete task: `POST /project/{projectId}/task/{taskId}/complete`
- Delete task: `DELETE /project/{projectId}/task/{taskId}`

Create task is an exception: `POST /task` (projectID in body)

Projects use simpler paths:
- List: `GET /project`
- Get: `GET /project/{projectId}`
- Create: `POST /project`
- Update: `POST /project/{projectId}`
- Delete: `DELETE /project/{projectId}`

### HTTP Methods

Note that TickTick API uses POST for updates (not PUT/PATCH). The Update methods in both services use POST requests.

### OAuth Token Management

The library handles OAuth2 code exchange and token refresh, but does not automatically refresh expired tokens. Callers are responsible for implementing token refresh logic based on `ExpiresIn` from `TokenResponse`.

### Time Handling

Date/time fields use `*time.Time` (pointers) to distinguish between "not set" (nil) and "set to zero value". The TickTick API expects ISO 8601 format dates.

### Required vs Optional Fields

Task creation requires only `Title`; all other fields including `ProjectID` are optional (defaults to inbox if not specified). Project creation requires only `Name`.
