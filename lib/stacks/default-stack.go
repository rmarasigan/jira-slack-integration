package stacks

import (
	"jira-slack-integration/internal/app/config"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type JiraSlackIntegrationStackProps struct{ awscdk.StackProps }

// newStack creates a new stack and will serve as the parent construct or scope.
func newStack(scope constructs.Construct, props *JiraSlackIntegrationStackProps) awscdk.Stack {
	var stackProps awscdk.StackProps

	if props != nil {
		stackProps = props.StackProps
	}

	return awscdk.NewStack(scope, config.STACK_ID, &stackProps)
}
