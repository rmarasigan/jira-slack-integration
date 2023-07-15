package jira

// IssueType represents the JIRA Projects issue type.
type IssueType struct {
	ID          string `json:"id"`          // The ID of the issue type.
	Name        string `json:"name"`        // The name of the issue type.
	Subtask     bool   `json:"subtask"`     // Whether the issue type is used to create subtasks.
	Description string `json:"description"` // The description of the issue type.
}
