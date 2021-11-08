package types

import (
	"context"
	"errors"
	"flettsgetmarried.com/api/logger"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dbTypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"time"
)

// Guest represents a single Guest and their requirements
type Guest struct {
	PK         string  `json:"-" dynamodbav:"PK"`
	SK         string  `json:"-" dynamodbav:"SK"`
	ID         string  `json:"id"`
	InviteID   string  `json:"inviteID"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	RSVP       bool    `json:"rsvp"`
	Vaccinated bool    `json:"vaccinated"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  *string `json:"updatedAt"`
	Mobile     *string `json:"mobile"`
	Email      *string `json:"email"`
	Diet       *string `json:"diet"`
	Music      *string `json:"music"`
}

func (g *Guest) Create(inviteID string) error {
	// Guest
	g.ID = uuid.NewString()
	g.PK = fmt.Sprintf("GUEST#%s", g.ID)
	g.SK = fmt.Sprintf("#PROFILE#%s", g.ID)
	g.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	guestAv, _ := attributevalue.MarshalMap(g)
	guestInput := &dynamodb.PutItemInput{
		TableName:    &DynamoTable,
		Item:         guestAv,
		ReturnValues: dbTypes.ReturnValueNone,
	}

	if _, err := DynamoClient.PutItem(context.TODO(), guestInput); err != nil {
		logger.Log.Error().Err(err).Str("guestID", g.ID).Msg("Error adding Guest to DynamoTable")
		return err
	}

	// Invite
	inviteAv, _ := attributevalue.MarshalMap(Invite{
		PK:      fmt.Sprintf("INVITE#%s", inviteID),
		SK:      fmt.Sprintf("GUEST#%s", g.ID),
		ID:      inviteID,
		GuestID: g.ID,
	})
	inviteInput := &dynamodb.PutItemInput{
		TableName:    &DynamoTable,
		Item:         inviteAv,
		ReturnValues: dbTypes.ReturnValueNone,
	}

	if _, err := DynamoClient.PutItem(context.TODO(), inviteInput); err != nil {
		logger.Log.Error().Err(err).Str("inviteID", inviteID).Str("guestID", g.ID).Msg("Error adding Invite to DynamoTable")
		return err
	}

	logger.Log.Info().Msg(fmt.Sprintf("Created guest %s %s with ID %s", g.FirstName, g.LastName, g.ID))
	return nil
}

func (g *Guest) Update() error {
	updatedAt := time.Now().UTC().Format(time.RFC3339)
	g.UpdatedAt = &updatedAt
	g.PK = fmt.Sprintf("GUEST#%s", g.ID)
	g.SK = fmt.Sprintf("#PROFILE#%s", g.ID)

	av, _ := attributevalue.MarshalMap(g)

	// update query
	input := &dynamodb.PutItemInput{
		TableName:    &DynamoTable,
		Item:         av,
		ReturnValues: dbTypes.ReturnValueNone,
	}

	if _, err := DynamoClient.PutItem(context.TODO(), input); err != nil {
		logger.Log.Error().Err(err).Str("guestID", g.ID).Msg("Unable to update the guest")
		return err
	}

	logger.Log.Info().Msg(fmt.Sprintf("Updated guest %s %s with ID %s", g.FirstName, g.LastName, g.ID))
	return nil
}

func (g *Guest) Get() error {
	input := &dynamodb.GetItemInput{
		Key: map[string]dbTypes.AttributeValue{
			"PK": &dbTypes.AttributeValueMemberS{Value: fmt.Sprintf("GUEST#%s", g.ID)},
			"SK": &dbTypes.AttributeValueMemberS{Value: fmt.Sprintf("#PROFILE#%s", g.ID)},
		},
		TableName: &DynamoTable,
	}

	// getItem
	result, err := DynamoClient.GetItem(context.TODO(), input)

	// handle errors
	if err != nil {
		logger.Log.Error().Err(err).Str("guestID", g.ID).Msg("error getting guest from database")
		return err
	}

	if len(result.Item) == 0 {
		return errors.New("user not found")
	}

	// unmarshal item into struct
	if err = attributevalue.UnmarshalMap(result.Item, &g); err != nil {
		logger.Log.Error().Err(err).Str("guestID", g.ID).Msg("failed to unmarshal dynamo item map to guest")
	}

	return nil
}
