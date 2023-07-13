package jira

// Issue represents a JIRA issue.
type Issue struct {
	Fields Fields `json:"fields"`
}

// Fields represents a single fields of the Jira Issue.
type Fields struct {
	Summary     string   `json:"summary"`             // The issue Title
	Description string   `json:"description"`         // The issue Description
	Project     KeyField `json:"project"`             // The issue Project ID
	Parent      KeyField `json:"parent,omitempty"`    // Contain the ID or key of the parent issue
	Priority    IDField  `json:"priority,omitempty"`  // The issue priority ID
	IssueType   IDField  `json:"issuetype,omitempty"` // Must be set to a subtask issue type
	Assignee    IDField  `json:"assignee,omitempty"`  // The issue assignee (user account ID)
	Reporter    IDField  `json:"reporter,omitempty"`  // The issue reporter (user account ID)
	Labels      []string `json:"labels,omitempty"`    // The issue label
}

type IDField struct {
	ID string `json:"id,omitempty"`
}

type KeyField struct {
	Key string `json:"key,omitempty"`
}

// IssueResponse represents a successful request response
// in creating a JIRA issue.
type IssueResponse struct {
	ID   string `json:"id"`   // The ID of the created issue or subtask.
	Key  string `json:"key"`  // The key of the created issue or subtask.
	Self string `json:"self"` // The URL of the created issue or subtask.
}
