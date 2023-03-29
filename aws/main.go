package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func helloGateway() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              "hello from Golang",
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	lambda.Start(helloGateway)
}
