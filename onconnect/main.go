package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request *events.APIGatewayWebsocketProxyRequest) (response *events.APIGatewayProxyResponse, err error) {
	defer func() {
	}()
	return nil, nil
}
