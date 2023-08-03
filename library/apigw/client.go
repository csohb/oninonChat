package apigw

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"log"
)

func getConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panicf("Configuration error, " + err.Error())
	}
	return cfg
}

func NewAPIGatewayManagementClient(domain, stage string) *apigatewaymanagementapi.Client {
	cfg := getConfig()
	return apigatewaymanagementapi.NewFromConfig(cfg)
}
