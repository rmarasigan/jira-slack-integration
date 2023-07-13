package main

import (
	"context"
	"jira-slack-integration/api"
	"jira-slack-integration/internal/app/jira"
	"jira-slack-integration/internal/app/slack"
	"jira-slack-integration/internal/utility"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.SQSEvent) error {
	utility.InitContext(ctx)

	var (
		chat     slack.Chat
		eventmap jira.EventMap
		records  = event.Records
	)

	if len(records) == 0 {
		utility.Info("SQSEvent", "no records found")
		return nil
	}

	// ******************** Fetch the configurations ******************** //
	// 1. Get the JIRA Configuration.
	jiracfg, err := jira.GetConfiguration(ctx)
	if err != nil {
		utility.Error(err, "JIRAError", "failed to fetch configuration")
		return err
	}

	// 2. Get the Slack configuration.
	slackcfg, err := slack.GetConfiguration(ctx)
	if err != nil {
		utility.Error(err, "SlackError", "failed to fetch slack configuration")
		return err
	}

	for _, record := range records {
		var body = record.Body

		// Unmarshal the event message
		err := api.ParseJSON([]byte(body), &eventmap)
		if err != nil {
			utility.Error(err, "JSONError", "failed to unmarshal the JSON-encoded data",
				utility.KVP{Key: "event", Value: body})

			return err
		}

		// ******************** Process the ticket ******************** //
		// 1. Marshal the event issue.
		ticket, err := api.EncodeJSON(eventmap.Issue)
		if err != nil {
			utility.Error(err, "JSONError", "failed to marshal JIRA issue", utility.KVP{Key: "event_map", Value: eventmap})
			return err
		}

		// 2. Create a JIRA ticket.
		response, err := jiracfg.CreateTicket(ctx, ticket)
		if err != nil {
			return err
		}

		// 3. Set the event mapping details.
		eventmap.IssueResponse = response
		eventmap.JiraEndpoint = jiracfg.Endpoint

		slackEnabled, err := slackcfg.IsEnabled()
		if err != nil {
			utility.Error(err, "SlackError", "failed to parse a string into a boolean value",
				utility.KVP{Key: "event_map", Value: eventmap})

			return err
		}

		// 3. If Slack Integration is enabled, send a
		// notification to the channel configured.
		if slackEnabled {
			chat.ChannelID = slackcfg.Channel
			chat.IncidentAlertTemplate(eventmap)

			data, err := api.EncodeJSON(chat)
			if err != nil {
				utility.Error(err, "JSONError", "failed to marshal the slack notification data", utility.KVP{Key: "event_map", Value: eventmap})
				return err
			}

			// 4. Send notification to the Slack channel.
			err = slackcfg.ChatPostMessage(ctx, data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
