package main

import (
	"errors"
	"flettsgetmarried.com/api/services"
	"flettsgetmarried.com/api/types"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type responseBody struct {
	Guests []types.Guest `json:"guests"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.QueryStringParameters["password"] != types.AdminPassword {
		return services.ReturnError(errors.New("forbidden"), http.StatusForbidden)
	}

	guests, err := services.GetAllGuests()
	if err != nil {
		return services.ReturnError(err, http.StatusInternalServerError)
	}

	rb := responseBody{
		Guests: guests,
	}
	return services.ReturnJSON(rb, http.StatusOK)
}

func main() {
	lambda.Start(Handler)
}
