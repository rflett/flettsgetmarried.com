package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const rsvpTable = "wedding-rsvps"

var (
	awsSession, _ = session.NewSession(&aws.Config{Region: aws.String("ap-southeast-2")})
	db            = dynamodb.New(awsSession)
)

// RequestBody is the expected HTTP request body
type RequestBody struct {
	Contact   string   `json:"contact"`
	Diet      string   `json:"diet"`
	Music     string   `json:"music"`
	Attendees []string `json:"attendees" dynamodbav:"Attendees,stringset"`
	Absentees []string `json:"absentees" dynamodbav:"Absentees,stringset"`
}

// RSVP is the guest invite response
type RSVP struct {
	Contact   string
	Diet      string
	Music     string
	Attendees []string
	Absentees []string
}

func (r RSVP) Create() error {
	// create attribute value
	av, _ := dynamodbattribute.MarshalMap(r)

	// define input
	input := &dynamodb.PutItemInput{
		TableName: aws.String(rsvpTable),
		Item: av,
		ReturnValues: aws.String("NONE"),
	}

	// create item
	_, err := db.PutItem(input)

	// handle errors
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return err
	}
	return nil
}

// Handler is the entrypoint to the lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// unmarshall request body to RequestBody struct
	requestBody := RequestBody{}
	jsonErr := json.Unmarshal([]byte(request.Body), &requestBody)
	if jsonErr != nil {
		return events.APIGatewayProxyResponse{Body: jsonErr.Error(), StatusCode: 400}, nil
	}

	// create in database
	rsvp := RSVP{
		Contact:   requestBody.Contact,
		Diet:      requestBody.Diet,
		Music:     requestBody.Music,
		Attendees: requestBody.Attendees,
		Absentees: requestBody.Absentees,
	}
	createErr := rsvp.Create()
	if createErr != nil {
		return events.APIGatewayProxyResponse{Body: createErr.Error(), StatusCode: 500}, nil
	}

	// create and send the response
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 201}, nil
}

func main() {
	lambda.Start(Handler)
}
