package main

import (
	"jira-slack-integration/internal/app/config"
	"jira-slack-integration/lib/stacks"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

// env determines the AWS environment (account+region) in which our stack is to be deployed.
// For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: config.AWS_ACCOUNT_ID,
		Region:  config.AWS_ACCOUNT_REGION,
	}
}

func main() {
	// To make sure that the CDK app cleans up after itself.
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	stacks.JiraSlackIntegrationStack(app, &stacks.JiraSlackIntegrationStackProps{
		StackProps: awscdk.StackProps{
			Env:         env(),
			StackName:   config.STACK_ID,
			Description: jsii.String("JIRA + Slack Integration"),
		},
	})

	app.Synth(nil)
}
