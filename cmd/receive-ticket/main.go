package main

import (
	"context"
	"errors"
	"jira-slack-integration/api"
	"jira-slack-integration/api/schema"
	"jira-slack-integration/internal/app/config"
	"jira-slack-integration/internal/app/jira"
	awswrapper "jira-slack-integration/internal/aws_wrapper"
	"jira-slack-integration/internal/utility"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	utility.InitContext(ctx)

	var (
		eventmap jira.EventMap
		body     = request.Body
		ticket   schema.JiraTicket
		queue    = os.Getenv("JIRA_QUEUE")
	)

	// Check if the queue is configured
	if queue == "" {
		err := errors.New("sqs JIRA_QUEUE environment variable is not set")
		utility.Error(err, "SQSError", "sqs JIRA_QUEUE is not configured on the environment",
			utility.KVP{Key: "payload", Value: body})

		return api.StatusInternalServer(), nil
	}

	// Unmarshal the received JSON-encoded data
	err := api.ParseJSON([]byte(body), &ticket)
	if err != nil {
		utility.Error(err, "JSONError", "failed to unmarshal the JSON-encoded data",
			utility.KVP{Key: "payload", Value: body})

		return api.StatusInternalServer(), nil
	}

	// ******************** JIRA ******************** //
	// 1. Get the JIRA Configuration.
	jiracfg, err := jira.GetConfiguration(ctx)
	if err != nil {
		utility.Error(err, "JIRAError", "failed to fetch configuration",
			utility.KVP{Key: "payload", Value: ticket})

		return api.StatusInternalServer(), nil
	}

	// 2. Get the JIRA Project details to validate if the project exists.
	project, err := jiracfg.GetProject(ctx, ticket.ProjectKey)
	if err != nil {
		utility.Error(err, "JIRAError", "failed to fetch project details",
			utility.KVP{Key: "payload", Value: ticket})

		// 2.1 Check if the error is "not found" and return an HTTP
		// BadRequest Status response.
		if strings.HasSuffix(err.Error(), "not found") {
			response, err := api.EncodeJSONString(api.SetErrorMessage(errors.New("'project_key' not found")))
			if err != nil {
				utility.Error(err, "JSONError", "failed to unmarshal the JSON-encoded data",
					utility.KVP{Key: "payload", Value: ticket})

				return api.StatusInternalServer(), nil
			}

			return api.StatusBadRequest(response), nil
		}

		return api.StatusInternalServer(), nil
	}

	// 3. Check if it is a valid Issue Type and return an HTTP
	// BadRequest Status response if it is invalid.
	isValidIssueType, isSubtask, typeName := project.IsValidIssueType(ticket.IssueTypeID)
	if !isValidIssueType {
		err := errors.New("'issue_type_id' not found")
		utility.Error(err, "APIError", "invalid issue_type_id", utility.KVP{Key: "payload", Value: ticket})

		response, err := api.EncodeJSONString(api.SetErrorMessage(err))
		if err != nil {
			utility.Error(err, "JSONError", "failed to unmarshal the JSON-encoded data",
				utility.KVP{Key: "payload", Value: ticket})

			return api.StatusInternalServer(), nil
		}

		return api.StatusBadRequest(response), nil
	}

	// 4. Check if it is a Subtask issue and requires the 'parent_key'.
	if isSubtask && ticket.ParentKey == "" {
		err := errors.New("'parent_key' is required")
		utility.Error(err, "APIError", "parent_key is not set", utility.KVP{Key: "payload", Value: ticket})

		response, err := api.EncodeJSONString(api.SetErrorMessage(err))
		if err != nil {
			utility.Error(err, "JSONError", "failed to unmarshal the JSON-encoded data",
				utility.KVP{Key: "payload", Value: ticket})

			return api.StatusInternalServer(), nil
		}

		return api.StatusBadRequest(response), nil
	}

	// 5. Get the JIRA Issue Priority to validate if the priority exists.
	priority, err := jiracfg.GetPriority(ctx, ticket.PriorityID)
	if err != nil {
		utility.Error(err, "JIRAError", "failed to fetch issue priority",
			utility.KVP{Key: "payload", Value: ticket})

		// 5.1 Check if the error is "not found" and return an HTTP
		// BadRequest Status response.
		if strings.HasSuffix(err.Error(), "not found") {
			response, err := api.EncodeJSONString(api.SetErrorMessage(errors.New("'priority_id' not found")))
			if err != nil {
				utility.Error(err, "JSONError", "failed to unmarshal the JSON-encoded data",
					utility.KVP{Key: "payload", Value: ticket})

				return api.StatusInternalServer(), nil
			}

			return api.StatusBadRequest(response), nil
		}

		return api.StatusInternalServer(), nil
	}

	// ******************** Event Mapping ******************** //
	// 1. Map the data and marshal the mapped JIRA issue.
	eventmap.Project = project
	eventmap.IssuePriority = priority
	eventmap.IssueTypeName = typeName
	eventmap.Issue = ticket.JiraIssueMapping()

	data, err := api.EncodeJSONString(eventmap)
	if err != nil {
		utility.Error(err, "JSONError", "failed to marshal mapped event",
			utility.KVP{Key: "payload", Value: ticket},
			utility.KVP{Key: "event_mapping", Value: eventmap})

		return api.StatusInternalServer(), nil
	}

	// 2. Send the message to the queue.
	err = awswrapper.SQSSendMessage(ctx, queue, data, config.JIRA_SQS_GROUP_ID)
	if err != nil {
		utility.Error(err, "SQSError", "failed to send message", utility.KVP{Key: "payload", Value: ticket})
		return api.StatusInternalServer(), nil
	}

	return api.StatusOKWithoutBody()
}
