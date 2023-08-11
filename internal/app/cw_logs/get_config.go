package cwlogs

import (
	"context"
	"errors"
	"fmt"
	"jira-slack-integration/api"
	awswrapper "jira-slack-integration/internal/aws_wrapper"
	"jira-slack-integration/internal/trail"
	"jira-slack-integration/internal/utility"
	"os"
	"strings"
)

// GetConfiguration fetches the CloudWatch Alert secret from the Secrets Manager
// and returns an object of the CloudWatch configuration.
func GetConfiguration(ctx context.Context) (Configuration, error) {
	var (
		cfg           Configuration
		cwalertsecret = os.Getenv("CLOUDWATCH_ALERT_SECRET")
	)

	// Check if the CloudWatch Alert SecretsManager is configured
	if cwalertsecret == "" {
		err := errors.New("secretsmanager CLOUDWATCH_ALERT_SECRET environment is not set")
		utility.Error(err, "SMError", "secretsmanager CLOUDWATCH_ALERT_SECRET is not configured on the environment")

		return cfg, err
	}

	// Get the CloudWatch Alert secret values
	trail.Info("Start getting CloudWatch Alert Configuration...")
	result, err := awswrapper.SecretGetValue(ctx, cwalertsecret)
	if err != nil {
		trail.Error("failed to fetch the CloudWatch Alert secret on the Secrets Manager")
		return cfg, err
	}

	// Unmarshal the CloudWatch Alert secret valuse
	err = api.ParseJSON([]byte(*result.SecretString), &cfg)
	if err != nil {
		trail.Error("failed to unmarshal the JSON-encoded secret")
		return cfg, err
	}

	err = cfg.isRequiredFieldsEmpty()
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

// isRequiredFieldsEmpty checks if the required fields are not
// empty.
func (cfg Configuration) isRequiredFieldsEmpty() error {
	var (
		fields  []string
		isEmpty = func(field string) bool {
			return field == ""
		}
	)

	if isEmpty(cfg.PriorityID) {
		fields = append(fields, "priority_id")
	}

	if isEmpty(cfg.ProjectKey) {
		fields = append(fields, "project_key")
	}

	if isEmpty(cfg.IssueTypeID) {
		fields = append(fields, "issue_type_id")
	}

	if isEmpty(cfg.ReporterID) {
		fields = append(fields, "reporter_id")
	}

	if isEmpty(cfg.ApiKey) {
		fields = append(fields, "api_key")
	}

	if len(fields) > 0 {
		return fmt.Errorf("missing field(s): %s", strings.Join(fields, ", "))
	}

	return nil
}
