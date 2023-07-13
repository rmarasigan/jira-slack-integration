package integration

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/jsii-runtime-go"
)

func ApiJIRAModel(api awsapigateway.RestApi) awsapigateway.Model {
	return api.AddModel(jsii.String("JIRATicketApiModel"), &awsapigateway.ModelOptions{
		ModelName:   jsii.String("JIRATicketApiModel"),
		ContentType: jsii.String("application/json"),
		Schema: &awsapigateway.JsonSchema{
			Type: awsapigateway.JsonSchemaType_OBJECT,
			Properties: &map[string]*awsapigateway.JsonSchema{
				"title": {
					Pattern: jsii.String("^.+"),
					Type:    awsapigateway.JsonSchemaType_STRING,
				},
				"description": {
					Pattern: jsii.String("^.+"),
					Type:    awsapigateway.JsonSchemaType_STRING,
				},
				"project_key": {
					Pattern: jsii.String("^.+"),
					Type:    awsapigateway.JsonSchemaType_STRING,
				},
				"issue_type_id": {
					Pattern: jsii.String("^.+"),
					Type:    awsapigateway.JsonSchemaType_STRING,
				},
				"priority_id": {
					Pattern: jsii.String("^.+"),
					Type:    awsapigateway.JsonSchemaType_STRING,
				},
				"parent_key": {
					Pattern: jsii.String("^.+"),
					Type:    awsapigateway.JsonSchemaType_STRING,
				},
				"assignee_id": {
					Pattern: jsii.String("^.+"),
					Type:    awsapigateway.JsonSchemaType_STRING,
				},
				"reporter_id": {
					Pattern: jsii.String("^.+"),
					Type:    awsapigateway.JsonSchemaType_STRING,
				},
				"labels": {
					Type: awsapigateway.JsonSchemaType_ARRAY,
					Items: awsapigateway.JsonSchema{
						Type: awsapigateway.JsonSchemaType_STRING,
					},
				},
			},
			Required: jsii.Strings("title", "description", "project_key", "issue_type_id", "priority_id", "reporter_id"),
		},
	})
}
