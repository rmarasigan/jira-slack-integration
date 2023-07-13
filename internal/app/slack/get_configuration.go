package slack

import (
	"context"
	"errors"
	"jira-slack-integration/api"
	awswrapper "jira-slack-integration/internal/aws_wrapper"
	"jira-slack-integration/internal/trail"
	"jira-slack-integration/internal/utility"
	"os"
)

// GetConfiguration fetches the Slack secret from the Secrets Manager and
// returns an object of the Slack configuration.
func GetConfiguration(ctx context.Context) (Configuration, error) {
	var (
		cfg         Configuration
		slacksecret = os.Getenv("SLACK_SECRET")
	)

	// Check if the Slack SecretsManager is configured
	if slacksecret == "" {
		err := errors.New("secretsmanager SLACK_SECRET environment is not set")
		utility.Error(err, "SMError", "secretsmanager SLACK_SECRET is not configured on the environment")

		return cfg, err
	}

	// Get the Slack secret values
	trail.Info("Start getting Slack Configuration...")
	result, err := awswrapper.SecretGetValue(ctx, slacksecret)
	if err != nil {
		trail.Error("failed to fetch the Slack secret on the Secrets Manager")
		return cfg, err
	}

	// Unmarshal the Slack secret values
	err = api.ParseJSON([]byte(*result.SecretString), &cfg)
	if err != nil {
		trail.Error("failed to unmarshal the JSON-encoded secret")
		return cfg, err
	}

	return cfg, nil
}
