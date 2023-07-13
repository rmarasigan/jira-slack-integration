package slack

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"jira-slack-integration/api"
	"jira-slack-integration/internal/trail"
	"jira-slack-integration/internal/utility"
	"net/http"
	"strconv"
)

type Configuration struct {
	Enabled      string `json:"enabled"`       // If this integration is enabled or not
	Token        string `json:"token"`         // Authentication token bearing required scopes
	Channel      string `json:"channel"`       // The public/private channel ID
	ChatEndpoint string `json:"chat_endpoint"` // Endpoint of chat.postMessage
}

// IsEnabled returns if the Slack integration is enabled.
func (cfg Configuration) IsEnabled() (bool, error) {
	enabled, err := strconv.ParseBool(cfg.Enabled)
	if err != nil {
		return false, err
	}

	return enabled, nil
}

func (cfg Configuration) authorization() string {
	return fmt.Sprintf("Bearer %s", cfg.Token)
}

// ChatPostMessage sends a message to the configured channel.
func (cfg Configuration) ChatPostMessage(ctx context.Context, data []byte) error {
	var (
		slackresponse Response
		client        http.Client
		endpoint      = cfg.ChatEndpoint
	)

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(data))
	if err != nil {
		utility.Error(err, "SlackError", "failed to create a new request with context",
			utility.KVP{Key: "data", Value: string(data)})

		return err
	}

	// Supply the request headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", cfg.authorization())

	// Send an HTTP request and return a response
	trail.Info("Start sending request to Slack Integration...")
	response, err := client.Do(request)
	if err != nil {
		utility.Error(err, "SlackError", "failed to send an HTTP request",
			utility.KVP{Key: "data", Value: string(data)})

		return err
	}
	defer response.Body.Close()

	// Read response from the request
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		utility.Error(err, "SlackError", "failed to read the response body",
			utility.KVP{Key: "data", Value: string(data)},
			utility.KVP{Key: "status", Value: response.Status})

		return err
	}

	err = api.ParseJSON(result, &slackresponse)
	if err != nil {
		utility.Error(err, "JSONError", "failed to unmarshal JSON-encoded response",
			utility.KVP{Key: "data", Value: string(data)},
			utility.KVP{Key: "response", Value: string(result)},
			utility.KVP{Key: "status", Value: response.Status})

		return err
	}

	if !slackresponse.OK {
		err = errors.New("failed to notify slack channel")
		utility.Error(err, "SlackError", "unable to create notification",
			utility.KVP{Key: "data", Value: string(data)},
			utility.KVP{Key: "response", Value: slackresponse},
			utility.KVP{Key: "status", Value: response.Status})

		return err
	}

	utility.Info("Slack", "successfully created slack notification",
		utility.KVP{Key: "data", Value: string(data)},
		utility.KVP{Key: "response", Value: slackresponse},
		utility.KVP{Key: "status", Value: response.Status})

	return nil
}
