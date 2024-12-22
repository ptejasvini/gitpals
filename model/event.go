package model

import "time"
// GitHubEvent represents a GitHub event (activity).
type GitHubEvent struct {
	ID string `json:"id"`
	Type  EventType  `json:"type"` // Event type (e.g., "PushEvent")
	Actor GitHubUser `json:"actor"`
	Repo GitHubRepo `json:"repo"` // Repository involved in the event
	Payload   Payload    `json:"payload"` // Event-specific details (e.g., commits, issue data)
	Public    bool       `json:"public"`
	CreatedAt time.Time  `json:"created_at"` // Time the event occurred
}

// EventType is a custom type to restrict event types to known GitHub event types.
type EventType string

const (
	PushEvent   EventType = "PushEvent"
	CreateEvent EventType = "CreateEvent"
	WatchEvent  EventType = "WatchEvent"
	IssueEvent  EventType = "IssuesEvent"
	// Add other event types as needed
)

// GitHubRepo represents the repository involved in the event.
type GitHubRepo struct {
	Name string `json:"name"` // Name of the repository
	URL  string `json:"url"`  // URL of the repository
}

// Payload represents specific event data (e.g., commits for PushEvent, issue data for IssuesEvent).
type Payload struct {
	Ref     string   `json:"ref,omitempty"`     // For PushEvent (e.g., branch name)
	Commits []Commit `json:"commits,omitempty"` // List of commits for PushEvent
	Action  string   `json:"action,omitempty"`  // Action (opened/closed) for IssuesEvent
}

// Commit represents a single commit in a PushEvent.
type Commit struct {
	Message string `json:"message"` // Commit message
	SHA     string `json:"sha"`     // Commit SHA (unique identifier)
	URL     string `json:"url"`
}
