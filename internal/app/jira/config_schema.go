package jira

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
)

type Configuration struct {
	Token        string `json:"token"`         // JIRA API Token for authentication.
	Username     string `json:"username"`      // Username to be used for Basic Auth.
	Endpoint     string `json:"endpoint"`      // The Atlassian Domain of your JIRA.
	IssuePath    string `json:"issue_path"`    // The REST API resource path for the JIRA issue.
	ProjectPath  string `json:"project_path"`  // The REST API resource path for the JIRA project.
	PriorityPath string `json:"priority_path"` // The REST API resource path for the JIRA priority.
}

// GetProject returns the representation of the project by passing the JIRA Project 'key' or 'id'.
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-get
func (cfg Configuration) GetProject(ctx context.Context, key string) (Project, error) {
	var (
		project  Project
		client   http.Client
		endpoint = fmt.Sprintf("%s%s%s", cfg.Endpoint, cfg.ProjectPath, key)
	)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		trail.Error("failed to create a new request with context")
		return project, err
	}

	// Supply the request headers
	request.SetBasicAuth(cfg.Username, cfg.Token)
	request.Header.Set("Content-Type", "application/json")

	// Send an HTTP request and return a response
	trail.Info("Start fetching a JIRA Project...")
	response, err := client.Do(request)
	if err != nil {
		trail.Error("failed to send an HTTP request")
		return project, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		// Read response from the request
		result, err := ioutil.ReadAll(response.Body)
		if err != nil {
			trail.Error("failed to read the response body")
			return project, err
		}

		err = api.ParseJSON(result, &project)
		if err != nil {
			utility.Error(err, "JSONError", "failed to unmarshal JSON-encoded data",
				utility.KVP{Key: "response", Value: string(result)})

			return project, err
		}
		utility.Info("JIRAProject", "successfully fetched project details",
			utility.KVP{Key: "project", Value: project})

		return project, nil

	default:
		// Read response from the request
		result, err := ioutil.ReadAll(response.Body)
		if err != nil {
			trail.Error("failed to read the response body")
			return project, err
		}

		err = errors.New("project not found")
		utility.Error(err, "JIRAError", "unable to fetch project details",
			utility.KVP{Key: "project_key", Value: key},
			utility.KVP{Key: "response", Value: string(result)},
			utility.KVP{Key: "status", Value: response.Status})

		return project, err
	}
}

// GetPriority returns the representation of the issue priority by passing the JIRA Issue Priority
// 'id'.
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-priorities/#api-rest-api-2-priority-id-get
func (cfg Configuration) GetPriority(ctx context.Context, id string) (Priority, error) {
	var (
		client   http.Client
		priority Priority
		endpoint = fmt.Sprintf("%s%s%s", cfg.Endpoint, cfg.PriorityPath, id)
	)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		trail.Error("failed to create a new request with context")
		return priority, err
	}

	// Supply the request headers
	request.SetBasicAuth(cfg.Username, cfg.Token)
	request.Header.Set("Content-Type", "application/json")

	// Send an HTTP request and return a response
	trail.Info("Start fetching a JIRA Priority...")
	response, err := client.Do(request)
	if err != nil {
		trail.Error("failed to send an HTTP request")
		return priority, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		// Read response from the request
		result, err := ioutil.ReadAll(response.Body)
		if err != nil {
			trail.Error("failed to read the response body")
			return priority, err
		}

		err = api.ParseJSON(result, &priority)
		if err != nil {
			utility.Error(err, "JSONError", "failed to unmarshal JSON-encoded data",
				utility.KVP{Key: "response", Value: string(result)})

			return priority, err
		}
		utility.Info("JIRAProject", "successfully fetched issue priority",
			utility.KVP{Key: "priority", Value: priority})

		return priority, nil

	default:
		// Read response from the request
		result, err := ioutil.ReadAll(response.Body)
		if err != nil {
			trail.Error("failed to read the response body")
			return priority, err
		}

		err = errors.New("priority not found")
		utility.Error(err, "JIRAError", "unable to fetch issue priorities",
			utility.KVP{Key: "priority_id", Value: id},
			utility.KVP{Key: "response", Value: string(result)},
			utility.KVP{Key: "status", Value: response.Status})

		return priority, err
	}
}

// CreateTicket creates an issue and returns the issue response if it is created successfully.
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-post
func (cfg Configuration) CreateTicket(ctx context.Context, ticket []byte) (IssueResponse, error) {
	var (
		client        http.Client
		issueresponse IssueResponse
		endpoint      = fmt.Sprintf("%s%s", cfg.Endpoint, cfg.IssuePath)
	)

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(ticket))
	if err != nil {
		utility.Error(err, "JIRAError", "failed to create a new request with context",
			utility.KVP{Key: "ticket", Value: string(ticket)})

		return issueresponse, err
	}

	// Supply the request headers
	request.SetBasicAuth(cfg.Username, cfg.Token)
	request.Header.Set("Content-Type", "application/json")

	// Send an HTTP request and return a response
	trail.Info("Start creating a JIRA ticket...")
	response, err := client.Do(request)
	if err != nil {
		utility.Error(err, "JIRAError", "failed to send an HTTP request",
			utility.KVP{Key: "ticket", Value: string(ticket)})

		return issueresponse, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusCreated:
		// Read response from the request
		result, err := ioutil.ReadAll(response.Body)
		if err != nil {
			utility.Error(err, "JIRAError", "failed to read the response body",
				utility.KVP{Key: "ticket", Value: string(ticket)},
				utility.KVP{Key: "status", Value: response.Status})

			return issueresponse, err
		}

		err = api.ParseJSON(result, &issueresponse)
		if err != nil {
			utility.Error(err, "JSONError", "failed to unmarshal JSON-encoded response",
				utility.KVP{Key: "ticket", Value: string(ticket)},
				utility.KVP{Key: "response", Value: string(result)},
				utility.KVP{Key: "status", Value: response.Status})

			return issueresponse, err
		}
		utility.Info("JIRATicket", "successfully created a JIRA issue",
			utility.KVP{Key: "ticket", Value: string(ticket)},
			utility.KVP{Key: "response", Value: issueresponse},
			utility.KVP{Key: "status", Value: response.Status})

		return issueresponse, nil

	default:
		// Read response from the request
		result, err := ioutil.ReadAll(response.Body)
		if err != nil {
			utility.Error(err, "JIRAError", "failed to read the response body",
				utility.KVP{Key: "ticket", Value: string(ticket)},
				utility.KVP{Key: "status", Value: response.Status})

			return issueresponse, err
		}

		err = errors.New("failed to create JIRA issue")
		utility.Error(err, "JIRAError", "unable to create ticket",
			utility.KVP{Key: "ticket", Value: string(ticket)},
			utility.KVP{Key: "response", Value: string(result)},
			utility.KVP{Key: "status", Value: response.Status})

		return issueresponse, err
	}
}
