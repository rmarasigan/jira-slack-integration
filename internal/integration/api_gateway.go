package integration

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

var (
	HTTP_METHOD_GET  = jsii.String("GET")
	HTTP_METHOD_POST = jsii.String("POST")
)

// NewRestApi returns a new API Gateway Rest API with pre-defined configuration.
//
// Pre-defined Fields Configuration:
//    - Deploy: True
//    - LoggingLevel: INFO
//    - MetricsEnabled: True
//    - TracingEnabled: True
//    - DataTraceEnabled: True
func NewRestApi(scope constructs.Construct, id, stageName string) awsapigateway.RestApi {
	return awsapigateway.NewRestApi(scope, jsii.String(id), &awsapigateway.RestApiProps{
		Deploy:      ENABLED,
		RestApiName: jsii.String(id),
		DeployOptions: &awsapigateway.StageOptions{
			MetricsEnabled:   ENABLED,
			TracingEnabled:   ENABLED,
			DataTraceEnabled: ENABLED,
			StageName:        jsii.String(stageName),
			LoggingLevel:     awsapigateway.MethodLoggingLevel_INFO,
		},
	})
}

// NewApiResource returns a new child resource where the passed API is the parent.
func NewApiResource(api awsapigateway.RestApi, path string) awsapigateway.Resource {
	return api.Root().AddResource(jsii.String(path), nil)
}

// ApiSubResource returns a sub resource of the child resource.
func ApiSubResource(resource awsapigateway.Resource, path string) awsapigateway.Resource {
	return resource.AddResource(jsii.String(path), nil)
}

// SetupApiKey adds a usage plan to specify who can access the API and
// an API key to identify API clients.
func SetupApiKey(api awsapigateway.RestApi, plan, key string) {
	var usagePlan = api.AddUsagePlan(jsii.String(plan), &awsapigateway.UsagePlanProps{
		Name: jsii.String(plan),
	})

	var apiKey = api.AddApiKey(jsii.String(key), &awsapigateway.ApiKeyOptions{
		ApiKeyName: jsii.String(key),
	})

	usagePlan.AddApiKey(apiKey, nil)
	usagePlan.AddApiStage(&awsapigateway.UsagePlanPerApiStage{
		Api:   api,
		Stage: api.DeploymentStage(),
	})
	usagePlan.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)
}

// SetupGatewayBadRequestBody adds a custom API BadRequest Body Gateway
// response.
func SetupGatewayBadRequestBody(api awsapigateway.RestApi, id string) {
	api.AddGatewayResponse(jsii.String(id), &awsapigateway.GatewayResponseOptions{
		StatusCode: jsii.String("400"),
		Type:       awsapigateway.ResponseType_BAD_REQUEST_BODY(),
		Templates: &map[string]*string{
			"application/json": jsii.String(`{"message": "$context.error.validationErrorString"}`),
		},
	})
}

// ApiRequestBodyValidator returns a configured request validator to validate
// the API request body.
func ApiRequestBodyValidator(api awsapigateway.RestApi) awsapigateway.RequestValidator {
	return api.AddRequestValidator(
		jsii.String("JiraIntegration_RequestBodyValidator"),
		&awsapigateway.RequestValidatorOptions{
			ValidateRequestBody:       ENABLED,
			ValidateRequestParameters: DISABLED,
			RequestValidatorName:      jsii.String("JiraIntegration_RequestBodyValidator"),
		})
}

// ApiRequestParameterValidator returns a configured request validator to validate
// the API request parameters.
func ApiRequestParameterValidator(api awsapigateway.RestApi) awsapigateway.RequestValidator {
	return api.AddRequestValidator(
		jsii.String("JiraIntegration_RequestParameterValidator"),
		&awsapigateway.RequestValidatorOptions{
			ValidateRequestBody:       DISABLED,
			ValidateRequestParameters: ENABLED,
			RequestValidatorName:      jsii.String("JiraIntegration_RequestParameterValidator"),
		})
}
