package types

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	awsConfig, _ = config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-2"))
	DynamoClient = dynamodb.NewFromConfig(awsConfig)
	DynamoTable  = "flettsgetmarried"
)

// HTTPBody is the expected request for managing an Invite and its guests
type HTTPBody struct {
	InviteID *string `json:"inviteId"`
	Guests   []Guest `json:"guests"`
}
