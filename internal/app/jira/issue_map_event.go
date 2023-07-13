package jira

type EventMap struct {
	Issue         Issue         `json:"issue"`
	Project       Project       `json:"project"`
	IssuePriority Priority      `json:"issue_priority"`
	IssueTypeName string        `json:"issue_type_name"`
	JiraEndpoint  string        `json:"jira_endpoint,omitempty"`
	IssueResponse IssueResponse `json:"issue_response,omitempty"`
}
