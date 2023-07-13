package api

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

const (
	CONTENT_TYPE = "application/json"
)

type Message struct {
	Error  string `json:"error,omitempty"`
	Custom string `json:"message,omitempty"`
}

// SetErrorMessage returns the formatted error message.
func SetErrorMessage(err error) Message {
	return Message{Error: err.Error()}
}

// StatusOK returns a response of an HTTP StatusOK with body.
func StatusOK(body string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": CONTENT_TYPE,
		},
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

// StatusOKWithoutBody returns a response of an HTTP StatusOK without body.
func StatusOKWithoutBody() (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": CONTENT_TYPE,
		},
		StatusCode: http.StatusOK,
	}, nil
}

// StatusBadRequest returns a response of an HTTP StatusBadRequest.
func StatusBadRequest(body string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Contet-Type": CONTENT_TYPE,
		},
		StatusCode: http.StatusBadRequest,
		Body:       body,
	}
}

// StatusInternalServer returns a response of an HTTP StatusInternalServer without body.
func StatusInternalServer() *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": CONTENT_TYPE,
		},
		StatusCode: http.StatusInternalServerError,
	}
}
