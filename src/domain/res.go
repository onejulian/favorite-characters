package domain

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type SuccessBody struct {
	SuccessMsg *string `json:"successMsg"`
}

func ApiResponse(statusCode int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = statusCode

	stringBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp.Body = string(stringBody)
	return &resp, nil
}
