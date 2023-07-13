package jira

import (
	"context"
	"errors"
	"jira-slack-integration/api"
	awswrapper "jira-slack-integration/internal/aws_wrapper"
	"jira-slack-integration/internal/trail"
	"jira-slack-integration/internal/utility"
	"os"
)

// GetConfiguration fetches the JIRA secret from the Secrets Manager and
// returns an object of the JIRA configuration.
func GetConfiguration(ctx context.Context) (Configuration, error) {
	var (
		cfg        Configuration
		jirasecret = os.Getenv("JIRA_SECRET")
	)

	// Check if the JIRA SecretsManager is configured
	if jirasecret == "" {
		err := errors.New("secretsmanager JIRA_SECRET environment is not set")
		utility.Error(err, "SMError", "secretsmanager JIRA_SECRET is not configured on the environment")

		return cfg, err
	}

	// Get the JIRA secret values
	trail.Info("Start getting JIRA Configuration...")
	result, err := awswrapper.SecretGetValue(ctx, jirasecret)
	if err != nil {
		trail.Error("failed to fetch the JIRA secret on the Secrets Manager")
		return cfg, err
	}

	// Unmarshal the JIRA secret values
	err = api.ParseJSON([]byte(*result.SecretString), &cfg)
	if err != nil {
		trail.Error("failed to unmarshal the JSON-encoded secret")
		return cfg, err
	}

	return cfg, nil
}
