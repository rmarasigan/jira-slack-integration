package jira

// IssueType represents the JIRA Projects issue type.
type IssueType struct {
	ID          string `json:"id"`          // The ID of the issue type.
	Name        string `json:"name"`        // The name of the issue type.
	Subtask     bool   `json:"subtask"`     // Whether the issue type is used to create subtasks.
	Description string `json:"description"` // The description of the issue type.
	Scope       Scope  `json:"scope,omitempty"`
}

// Scope is the details of the next-gen projects the issue
// type is available in.
type Scope struct {
	Type    string `json:"type,omitempty"` // The type of scope (valid values: PROJECT, TEMPLATE)
	Project struct {
		ID string `json:"id,omitempty"` // The ID of the project.
	} `json:"project,omitempty"` // The project the item has scope in.
}
