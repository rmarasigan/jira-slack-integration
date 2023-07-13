package schema

import (
	"jira-slack-integration/internal/app/jira"
)

type JiraTicket struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ProjectKey  string   `json:"project_key"`
	PriorityID  string   `json:"priority_id"`
	ReporterID  string   `json:"reporter_id"`
	IssueTypeID string   `json:"issue_type_id"`
	ParentKey   string   `json:"parent_key,omitempty"`
	AssigneeID  string   `json:"assignee_id,omitempty"`
	Labels      []string `json:"labels,omitempty"`
}

func (ticket JiraTicket) JiraIssueMapping() jira.Issue {
	var issue jira.Issue

	issue.Fields.Labels = ticket.Labels
	issue.Fields.Summary = ticket.Title
	issue.Fields.Parent.Key = ticket.ParentKey
	issue.Fields.Priority.ID = ticket.PriorityID
	issue.Fields.Assignee.ID = ticket.AssigneeID
	issue.Fields.Reporter.ID = ticket.ReporterID
	issue.Fields.Project.Key = ticket.ProjectKey
	issue.Fields.Description = ticket.Description
	issue.Fields.IssueType.ID = ticket.IssueTypeID

	return issue
}
