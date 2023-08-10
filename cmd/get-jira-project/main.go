package main

import (
	"context"
	"errors"
	"jira-slack-integration/api"
	"jira-slack-integration/internal/app/jira"
	"jira-slack-integration/internal/utility"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	utility.InitContext(ctx)

	var projectKey = request.QueryStringParameters["key"]

	// 1. Get the JIRA Configuration.
	jiracfg, err := jira.GetConfiguration(ctx)
	if err != nil {
		utility.Error(err, "JIRAError", "failed to fetch configuration")
		return api.StatusInternalServer(), nil
	}

	// 2. Get the JIRA Project details.
	project, err := jiracfg.GetProject(ctx, projectKey)
	if err != nil {
		utility.Error(err, "JIRA", "failed to fetch project details")

		// 2.1 Check if the error is "not found" and return an HTTP
		// BadRequest Status response.
		if strings.HasSuffix(err.Error(), "not found") {
			response, err := api.EncodeJSONString(api.SetErrorMessage(errors.New("'key' not found")))
			if err != nil {
				utility.Error(err, "JSONError", "failed to unmarshal the JSON-encoded data")
				return api.StatusInternalServer(), nil
			}

			return api.StatusBadRequest(response), nil
		}

		return api.StatusInternalServer(), nil
	}

	response, err := api.EncodeJSONString(project)
	if err != nil {
		utility.Error(err, "JSONError", "failed to unmarshal the JSON-encoded data")
		return api.StatusInternalServer(), nil
	}

	return api.StatusOK(response)
}
