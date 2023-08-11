package config

import (
	"github.com/aws/jsii-runtime-go"
)

var (
	JIRA_SQS_GROUP_ID  = "jira.ticket"
	AWS_ACCOUNT_REGION = jsii.String("us-east-1")
	AWS_ACCOUNT_ID     = jsii.String("123456789101")
	STACK_ID           = jsii.String("JiraSlackIntegrationStack")
	JIRA_SECRET        = jsii.String("JIRA-SECRET-ARN")
	SLACK_SECRET       = jsii.String("SLACK-SECRET-ARN")
	CLOUDWATCH_SECRET  = jsii.String("CLOUDWATCH-SECRET-ARN")
)
