package main

import (
	"context"
	"errors"
	"fmt"
	"jira-slack-integration/api"
	cwlogs "jira-slack-integration/internal/app/cw_logs"
	awswrapper "jira-slack-integration/internal/aws_wrapper"
	"jira-slack-integration/internal/utility"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event awswrapper.CloudWatchEvent) error {
	utility.InitContext(ctx)

	var (
		apiroot   = os.Getenv("API_ROOT_URL")
		apiticket = os.Getenv("API_TICKET_PATH")
	)

	// 1. Check if the apiroot is configured
	if apiroot == "" {
		err := errors.New("api gateway API_ROOT_URL environment variable is not set")
		utility.Error(err, "APIGateway", "api gateway API_ROOT_URL is not configured on the environment")

		return err
	}

	// 2. Check if the apiticket is configured
	if apiticket == "" {
		err := errors.New("api gateway API_TICKET_PATH environment variable is not set")
		utility.Error(err, "APIGateway", "api gateway API_TICKET_PATH is not configured on the environment")

		return err
	}

	// 3. Decode and decompress the received CloudWatch Event
	data, err := event.DecodeData()
	if err != nil {
		utility.Error(err, "CWError", "failed to decode and decompress the received event from CloudWatch",
			utility.KVP{Key: "event", Value: event})

		return err
	}

	// ******************** CloudWatch ******************** //
	// 4. Get the CloudWatch Alert Configuration.
	cwcfg, err := cwlogs.GetConfiguration(ctx)
	if err != nil {
		utility.Error(err, "CWError", "failed to fetch configuration",
			utility.KVP{Key: "cloudwatch_data", Value: data})

		return err
	}

	// 5. Map the decoded CloudWatch Event.
	cwcfg.CloudWatchTicketMapping(*data)

	// 6. Create a JIRA Ticket by sending a POST request to
	// the API Gateway Endpoint.
	ticket, err := api.EncodeJSON(cwcfg)
	if err != nil {
		utility.Error(err, "JSONError", "failed to marshal the cloudwatch ticket data",
			utility.KVP{Key: "data", Value: data})

		return err
	}

	endpoint := fmt.Sprintf("%s%s", apiroot, strings.TrimPrefix(apiticket, "/"))
	err = cwcfg.CreateTicket(ctx, endpoint, ticket)
	if err != nil {
		return err
	}

	return nil
}
