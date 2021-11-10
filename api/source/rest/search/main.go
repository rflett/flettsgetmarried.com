package main

import (
	"flettsgetmarried.com/api/services"
	"flettsgetmarried.com/api/types"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"strings"
)

type responseBody struct {
	PartialMatch bool             `json:"partialMatch"`
	Matches      []types.HTTPBody `json:"matches"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	firstName := request.QueryStringParameters["firstName"]
	lastName := request.QueryStringParameters["lastName"]

	invites, _ := services.FindInvites(strings.ToLower(firstName), strings.ToLower(lastName))

	var matches []types.HTTPBody
	matches = []types.HTTPBody{}
	for _, inviteID := range invites {
		guests, _ := services.GetGuestsOnInvite(inviteID)
		matches = append(matches, types.HTTPBody{
			InviteID: &inviteID,
			Guests:   guests,
		})
	}

	rb := responseBody{
		PartialMatch: len(invites) != 1,
		Matches:      matches,
	}
	return services.ReturnJSON(rb, http.StatusOK)
}

func main() {
	lambda.Start(Handler)
}
