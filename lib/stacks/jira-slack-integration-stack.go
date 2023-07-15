package stacks

import (
	"jira-slack-integration/internal/app/config"
	"jira-slack-integration/internal/integration"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func JiraSlackIntegrationStack(scope constructs.Construct, props *JiraSlackIntegrationStackProps) {
	var (
		stack       = newStack(scope, props)
		queue       = integration.NewQueueWithDLQ(stack, "jira-integration-queue", true, jsii.Number(60000))
		jirasecret  = awssecretsmanager.Secret_FromSecretCompleteArn(stack, jsii.String("JiraIntegrationSecret"), config.JIRA_SECRET)
		slacksecret = awssecretsmanager.Secret_FromSecretCompleteArn(stack, jsii.String("SlackIntegrationSecret"), config.SLACK_SECRET)
	)

	// *************** Lambda Functions *************** //
	getJiraUsers := integration.NewLambdaFunction(stack, "get-jira-users", "cmd/get-jira-users")
	getJiraUsers.AddEnvironment(jsii.String("JIRA_SECRET"), jirasecret.SecretArn(), nil)
	getJiraUsers.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)
	jirasecret.GrantRead(getJiraUsers, nil)

	getJiraProject := integration.NewLambdaFunction(stack, "get-jira-project", "cmd/get-jira-project")
	getJiraProject.AddEnvironment(jsii.String("JIRA_SECRET"), jirasecret.SecretArn(), nil)
	getJiraProject.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)
	jirasecret.GrantRead(getJiraProject, nil)

	getJiraPriorities := integration.NewLambdaFunction(stack, "get-jira-priorities", "cmd/get-jira-priorities")
	getJiraPriorities.AddEnvironment(jsii.String("JIRA_SECRET"), jirasecret.SecretArn(), nil)
	getJiraPriorities.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)
	jirasecret.GrantRead(getJiraPriorities, nil)

	receiveTicket := integration.NewLambdaFunction(stack, "receive-ticket", "cmd/receive-ticket")
	receiveTicket.AddEnvironment(jsii.String("JIRA_SECRET"), jirasecret.SecretArn(), nil)
	receiveTicket.AddEnvironment(jsii.String("JIRA_QUEUE"), queue.QueueUrl(), nil)
	receiveTicket.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)
	jirasecret.GrantRead(receiveTicket, nil)
	queue.GrantSendMessages(receiveTicket)

	createTicket := integration.NewLambdaFunction(stack, "create-ticket", "cmd/create-ticket")
	createTicket.AddEnvironment(jsii.String("SLACK_SECRET"), slacksecret.SecretArn(), nil)
	createTicket.AddEnvironment(jsii.String("JIRA_SECRET"), jirasecret.SecretArn(), nil)
	createTicket.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)
	slacksecret.GrantRead(createTicket, nil)
	jirasecret.GrantRead(createTicket, nil)

	createTicket.AddEventSource(
		awslambdaeventsources.NewSqsEventSource(queue,
			&awslambdaeventsources.SqsEventSourceProps{
				BatchSize:               jsii.Number(1),
				ReportBatchItemFailures: integration.ENABLED,
			}))

	// *************** API Gateway *************** //
	api := integration.NewRestApi(stack, "jira-integration", "production")
	integration.SetupGatewayBadRequestBody(api, "JiraIntegration_BadRequestGatewayResponse")
	integration.SetupApiKey(api, "jira-usage-plan", "jira-api-key")
	api.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)

	JiraApiModel := integration.ApiJIRAModel(api)
	JiraApiModel.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)

	RequestBodyValidator := integration.ApiRequestBodyValidator(api)
	RequestBodyValidator.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)

	RequestParameterValidator := integration.ApiRequestParameterValidator(api)
	RequestParameterValidator.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)

	jira := integration.NewApiResource(api, "jira")
	jira.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)

	users := integration.ApiSubResource(jira, "users")
	users.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)
	users.AddMethod(
		integration.HTTP_METHOD_GET,
		awsapigateway.NewLambdaIntegration(getJiraUsers, nil),
		&awsapigateway.MethodOptions{
			ApiKeyRequired: integration.ENABLED,
		})

	project := integration.ApiSubResource(jira, "project")
	project.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)
	project.AddMethod(
		integration.HTTP_METHOD_GET,
		awsapigateway.NewLambdaIntegration(getJiraProject, nil),
		&awsapigateway.MethodOptions{
			ApiKeyRequired: integration.ENABLED,
			RequestParameters: &map[string]*bool{
				"method.request.querystring.key": jsii.Bool(true),
			},
			RequestValidator: RequestParameterValidator,
		},
	)

	priorities := integration.ApiSubResource(jira, "priorities")
	priorities.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)
	priorities.AddMethod(
		integration.HTTP_METHOD_GET,
		awsapigateway.NewLambdaIntegration(getJiraPriorities, nil),
		&awsapigateway.MethodOptions{
			ApiKeyRequired: integration.ENABLED,
		})

	ticket := integration.ApiSubResource(jira, "ticket")
	ticket.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)

	ticket.AddMethod(
		integration.HTTP_METHOD_POST,
		awsapigateway.NewLambdaIntegration(receiveTicket, nil),
		&awsapigateway.MethodOptions{
			ApiKeyRequired: integration.ENABLED,
			RequestModels: &map[string]awsapigateway.IModel{
				"application/json": JiraApiModel,
			},
			RequestValidator: RequestBodyValidator,
		})
}
