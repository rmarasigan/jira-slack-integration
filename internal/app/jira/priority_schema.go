package jira

// Priority represents the JIRA priority.
type Priority struct {
	ID          string `json:"id"`          // The ID of the issue priority.
	Name        string `json:"name"`        // The name of the issue priority.
	Description string `json:"description"` // The description of the issue priority.
	Self        string `json:"self"`        // The URL of the issue priority.
}
