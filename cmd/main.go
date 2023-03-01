package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	resp := fmt.Sprintf("Method= %s| Body=%s", request.RequestContext.HTTP.Method, request.Body)
	
	return events.LambdaFunctionURLResponse{
		Body: resp, 
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}