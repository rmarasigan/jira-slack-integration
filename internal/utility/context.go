package utility

import (
	"context"

	"github.com/aws/aws-lambda-go/lambdacontext"
)

var (
	aws_request_id       string
	lambda_function_name string
	invoked_function_arn string
)

// LambdaContext contains the metadata of a lambda function.
type LambdaContext struct {
	AWSRequestID       string `json:"aws_request_id,omitempty"`
	FunctionName       string `json:"lambda_function_name,omitempty"`
	InvokedFunctionArn string `json:"invoked_function_arn,omitempty"`
}

// InitContext initializes the lambda context information.
func InitContext(ctx context.Context) {
	lambdactx, _ := lambdacontext.FromContext(ctx)

	aws_request_id = lambdactx.AwsRequestID
	lambda_function_name = lambdacontext.FunctionName
	invoked_function_arn = lambdactx.InvokedFunctionArn
}
