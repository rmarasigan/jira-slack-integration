package slack

import (
	"fmt"
	"jira-slack-integration/internal/app/jira"
)

type Chat struct {
	ChannelID string  `json:"channel"`          // The public/private channel ID
	Blocks    []Block `json:"blocks,omitempty"` // Array of structured attachments
}

type Block struct {
	Type string `json:"type,omitempty"`
	Text struct {
		Type    string `json:"type"`
		Content string `json:"text"`
	} `json:"text"`
}

// IncidentAlertTemplate formats the Slack Notification message template
// by setting the header and content.
func (chat *Chat) IncidentAlertTemplate(eventmap jira.EventMap) {
	var (
		block   Block
		content string
	)

	// Set the Incident Message Header
	block.Type = "header"
	block.Text.Type = "plain_text"
	block.Text.Content = fmt.Sprintf(":bell: %s", eventmap.Issue.Fields.Summary)
	chat.Blocks = append(chat.Blocks, block)

	// Set the Incident Message Content
	block = Block{}
	block.Type = "section"
	block.Text.Type = "mrkdwn"

	content = fmt.Sprintf("\n*Priority*: %s", eventmap.IssuePriority.Name)
	content += fmt.Sprintf("\n*Project*: %s", eventmap.Project.Name)
	content += fmt.Sprintf("\n*Ticket URL*: %s/browse/%s", eventmap.JiraEndpoint, eventmap.IssueResponse.Key)
	content += fmt.Sprintf("\n\n\n*Description*:\n%s", eventmap.Issue.Fields.Description)
	block.Text.Content = content

	chat.Blocks = append(chat.Blocks, block)
}
