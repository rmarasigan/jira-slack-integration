package jira

// User represents the JIRA user schema.
type User struct {
	AccountID   string `json:"accountId"`   // The account ID of the user
	AccountType string `json:"accountType"` // The user account type (valid values: atlassian, app, customer, unknown)
	DisplayName string `json:"displayName"` // The display name of the user
	Active      bool   `json:"active"`      // Whether the user is active
}
