package integration

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// NewLambdaFunction returns an AWS Lambda Function with pre-defined configuration.
//
// Pre-defined Fields Configuration:
//    - Runtime: Go
//    - Tracing: Active
//    - RetryAttempts: 2
//    - MemorySize: 1024
//    - Timeout: 1 minute
//    - Architecture: x86_64
func NewLambdaFunction(scope constructs.Construct, id, code string) awslambda.Function {
	return awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
		RetryAttempts: jsii.Number(2),
		Handler:       jsii.String(id),
		FunctionName:  jsii.String(id),
		MemorySize:    jsii.Number(1024),
		Tracing:       awslambda.Tracing_ACTIVE,
		Runtime:       awslambda.Runtime_GO_1_X(),
		Architecture:  awslambda.Architecture_X86_64(),
		Timeout:       awscdk.Duration_Millis(jsii.Number(60000)),
		Code:          awslambda.Code_FromAsset(jsii.String(code), nil),
	})
}
