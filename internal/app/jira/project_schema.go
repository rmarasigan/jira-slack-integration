package jira

// Project represents the JIRA projects.
type Project struct {
	ID          string      `json:"id"`             // The ID of the project.
	Key         string      `json:"key"`            // The key of the project.
	Name        string      `json:"name"`           // The name of the project.
	IssueTypes  []IssueType `json:"issueTypes"`     // List of the issue types available in the project.
	ProjectType string      `json:"projectTypeKey"` // The project type of the project (valid values: software, service_desk, business)
}

// IsValidIssueType returns if it is a valid issue type, a subtask or not, and its name.
func (project Project) IsValidIssueType(id string) (validIssueType bool, isSubtask bool, typeName string) {
	for _, issue := range project.IssueTypes {
		if issue.ID == id {
			if issue.Subtask {
				isSubtask = true
				validIssueType = true
				typeName = issue.Name

				return

			} else {
				isSubtask = false
				validIssueType = true
				typeName = issue.Name

				return
			}
		}
	}

	return
}
