package services

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flettsgetmarried.com/api/logger"
	"flettsgetmarried.com/api/types"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"os"
	"strings"
)

func removeDuplicates(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func scanForGuests(filter expression.ConditionBuilder) (invites []string, err error) {
	proj := expression.NamesList(expression.Name("InviteID"))

	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(proj).Build()

	if err != nil {
		logger.Log.Error().Err(err).Msg("error building expression for FindInvites")
		return nil, err
	}

	input := &dynamodb.ScanInput{
		TableName:                 &types.DynamoTable,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
	}
	response, err := types.DynamoClient.Scan(context.TODO(), input)

	if err != nil {
		logger.Log.Error().Err(err).Msg("error Scanning for invites")
		return nil, err
	}

	for _, item := range response.Items {
		guest := types.Guest{}
		if err = attributevalue.UnmarshalMap(item, &guest); err != nil {
			logger.Log.Error().Err(err).Msg("error unmarshalling item to guest")
		}
		invites = append(invites, guest.InviteID)
	}

	return removeDuplicates(invites), nil
}

func FindInvites(firstName, lastName string) (invites []string, err error) {
	var filter expression.ConditionBuilder

	firstNameFilter := expression.Name("FirstName").Equal(expression.Value(firstName))
	lastNameFilter := expression.Name("LastName").Equal(expression.Value(lastName))

	filter = expression.And(firstNameFilter, lastNameFilter)
	if invites, err = scanForGuests(filter); invites != nil {
		return invites, nil
	}

	filter = expression.Or(firstNameFilter, lastNameFilter)
	if invites, err = scanForGuests(filter); invites != nil {
		return invites, nil
	}

	return nil, nil
}

func LoadCSV() {
	for {
		records, err := readData("../../invites.csv")

		if err != nil {
			logger.Log.Fatal().Err(err).Msg("Error calling readData")
		}
		for _, record := range records {
			inviteID := record[0]
			var mobile *string
			var email *string

			if len(record[3]) == 0 {
				mobile = nil
			} else {
				mobile = &record[3]
			}

			if len(record[4]) == 0 {
				email = nil
			} else {
				email = &record[4]
			}

			guest := types.Guest{
				FirstName: strings.ToLower(record[1]),
				LastName:  strings.ToLower(record[2]),
				Mobile:    mobile,
				Email:     email,
				InviteID:  inviteID,
			}
			if err = guest.Create(inviteID); err != nil {
				logger.Log.Error().Err(err).Str("Name", record[1]).Msg("Error creating guest")
			}
		}
		break
	}
}

func ReturnJSON(body interface{}, status int) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{"Content-Type": "application/json"}
	marshalledBody, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{Body: string(marshalledBody), StatusCode: status, Headers: headers}, nil
}

// ReturnError returns an error from APIGW in a standard format
func ReturnError(err error, status int) (events.APIGatewayProxyResponse, error) {
	logger.Log.Info().Str("status", string(rune(status))).Str("err", err.Error()).Msg("Returning error from APIGW")
	body := map[string]interface{}{
		"success": false,
		"error":   err.Error(),
	}
	return ReturnJSON(body, status)
}

func GetGuestsOnInvite(inviteID string) ([]types.Guest, error) {
	pkCondition := expression.Key("PK").Equal(expression.Value(fmt.Sprintf("INVITE#%s", inviteID)))
	skCondition := expression.Key("SK").BeginsWith("GUEST#")
	keyCondition := expression.KeyAnd(pkCondition, skCondition)

	projExpr := expression.NamesList(expression.Name("GuestID"))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).WithProjection(projExpr).Build()

	if err != nil {
		logger.Log.Error().Err(err).Msg("error building expression for GetGuestsOnInvite")
	}

	// input
	input := &dynamodb.QueryInput{
		TableName:                 &types.DynamoTable,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
	}

	invites, err := types.DynamoClient.Query(context.TODO(), input)
	if err != nil {
		logger.Log.Error().Err(err).Str("inviteID", inviteID).Msg("error getting guests on invite")
		return []types.Guest{}, err
	}

	var guests []types.Guest = nil
	for _, invitee := range invites.Items {
		guest := types.Guest{}
		invite := types.Invite{}
		if err = attributevalue.UnmarshalMap(invitee, &invite); err != nil {
			logger.Log.Error().Err(err).Msg("Unable to unmarshal invitee to guest")
			continue
		}
		guest.ID = invite.GuestID
		if err = guest.Get(); err != nil {
			logger.Log.Error().Err(err).Msg("Unable to get guest")
			continue
		}
		guests = append(guests, guest)
	}
	return guests, nil
}

func GetAllGuests() ([]types.Guest, error) {
	pkFilter := expression.Name("PK").BeginsWith("GUEST#")
	expr, err := expression.NewBuilder().WithFilter(pkFilter).Build()

	if err != nil {
		logger.Log.Error().Err(err).Msg("error building expression for GetGroupFromCode func")
	}

	// input
	input := &dynamodb.ScanInput{
		TableName:                 &types.DynamoTable,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	results, err := types.DynamoClient.Scan(context.TODO(), input)

	if err != nil {
		logger.Log.Error().Err(err).Msg("error getting all guests")
		return []types.Guest{}, err
	}

	var guests []types.Guest = nil
	for _, result := range results.Items {
		guest := types.Guest{}
		if err = attributevalue.UnmarshalMap(result, &guest); err != nil {
			logger.Log.Error().Err(err).Msg("Unable to unmarshal result to guest")
			continue
		}
		guests = append(guests, guest)
	}

	return guests, nil
}