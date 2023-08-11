package cwlogs

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	awswrapper "jira-slack-integration/internal/aws_wrapper"
	"jira-slack-integration/internal/trail"
	"jira-slack-integration/internal/utility"
	"net/http"
	"strings"
)

type Configuration struct {
	PriorityID  string `json:"priority_id"`          // The issue priority ID
	ProjectKey  string `json:"project_key"`          // The issue Project key
	IssueTypeID string `json:"issue_type_id"`        // The issue type ID
	Title       string `json:"title"`                // The issue title
	Description string `json:"description"`          // The issue description
	ReporterID  string `json:"reporter_id"`          // The issue reporter ID (user account ID)
	ParentKey   string `json:"parent_key,omitempty"` // Contain the ID or key of the parent issue
	ApiKey      string `json:"api_key"`              // The API Gateway key
}

// CloudWatchTicketMapping sets the ticket title and description based on the
// received CloudWatch data.
func (cfg *Configuration) CloudWatchTicketMapping(data awswrapper.CloudWatchData) {
	var (
		description string
		logGroup    = strings.Split(data.LogGroup, "/")
	)

	description = fmt.Sprintf("*Owner*: %s\n", data.Owner)
	description += fmt.Sprintf("*Log Group*: %s\n", data.LogGroup)
	description += fmt.Sprintf("*Log Stream*: %s\n\n", data.LogStream)

	for index, log := range data.LogEvents {
		description += fmt.Sprintf("*Error #%v*\n", index+1)
		description += fmt.Sprintf("%s\n\n", log.Message)
	}

	cfg.Title = fmt.Sprintf("[%s] %s", logGroup[len(logGroup)-1], data.LogStream)
	cfg.Description = description
}

// CreateTicket sends an API POST request method to the API Gateway configured
// in the stack to ensure its reliability and FIFO sequence when creating a ticket.
func (cfg Configuration) CreateTicket(ctx context.Context, endpoint string, ticket []byte) error {
	var client http.Client

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(ticket))
	if err != nil {
		utility.Error(err, "CWError", "failed to create a new request with context",
			utility.KVP{Key: "ticket", Value: string(ticket)})

		return err
	}

	// Supply the reqest headers
	request.Header.Set("X-Api-Key", cfg.ApiKey)
	request.Header.Set("Content-Type", "application/json")

	// Send an HTTP request and return a response
	trail.Info("Start sending a request to the API Gateway....")
	response, err := client.Do(request)
	if err != nil {
		utility.Error(err, "CWError", "failed to send an HTTP request",
			utility.KVP{Key: "ticket", Value: string(ticket)})

		return err
	}
	defer response.Body.Close()

	// Read response from the request
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		utility.Error(err, "CWError", "failed to read the response body",
			utility.KVP{Key: "ticket", Value: ticket},
			utility.KVP{Key: "status", Value: response.Status})

		return err
	}

	switch response.StatusCode {
	case http.StatusOK:
		utility.Info("CWAlertTicket", "successfully created a JIRA ticket",
			utility.KVP{Key: "response", Value: string(result)})

	default:
		utility.Warning("CWAlertTicket", "failed to create a JIRA ticket",
			utility.KVP{Key: "ticket", Value: string(ticket)},
			utility.KVP{Key: "response", Value: string(result)},
			utility.KVP{Key: "status", Value: response.Status})
	}

	return nil
}
