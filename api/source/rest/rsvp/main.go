package main

import (
	"encoding/json"
	"flettsgetmarried.com/api/services"
	"flettsgetmarried.com/api/types"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// unmarshall request body to HTTPBody struct
	requestBody := types.HTTPBody{}
	jsonErr := json.Unmarshal([]byte(request.Body), &requestBody)
	if jsonErr != nil {
		return events.APIGatewayProxyResponse{Body: jsonErr.Error(), StatusCode: http.StatusBadRequest}, nil
	}

	for _, guest := range requestBody.Guests {
		if err := guest.Update(); err != nil {
			return services.ReturnError(err, http.StatusInternalServerError)
		}
	}

	return events.APIGatewayProxyResponse{Body: "", StatusCode: http.StatusNoContent}, nil
}

func main() {
	lambda.Start(Handler)
}
