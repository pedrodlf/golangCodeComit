package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

//Handler metodo "main" de AWS
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("en el metododo")
	log.Print("request path")
	log.Print(req.Path)
	log.Print("request body")
	log.Print(req.Body)
	log.Print("HTTP Mehtod")
	log.Print(req.HTTPMethod)

	points, err := strconv.ParseInt(req.QueryStringParameters["points"], 10, 32)
	actions, err := strconv.ParseInt(req.QueryStringParameters["action"], 10, 32)
	meetingID, err := strconv.ParseInt(req.QueryStringParameters["meetingID"], 10, 32)
	userID := req.PathParameters["userID"]
	points32 := uint32(points)
	meeting := uint32(meetingID)
	action := uint32(actions)
	if err != nil {

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Number Format Exception ",
		}, nil

	}

	log.Printf("points-----> : %v", points)
	log.Printf("action-----> : %v", action)
	log.Printf("userID-----> : %v", userID)
	log.Printf("meetingID--> : %v", meeting)

	user, err := GetUsserAddress(userID)

	if err != nil {

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Null Address Exception: User common.Address cant be recovered from Database or don't exist err--->" + err.Error(),
		}, nil

	}
	log.Printf("user.Address--> : %v", user)

	result, err := AddPoints(user, meeting, points32, action)
	if err != nil {

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error updating Points to chain err--->" + err.Error(),
		}, nil

	}
	switch result {
	case 900:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "Points Updated",
		}, nil
	case 902:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Points can't be updated ",
		}, nil
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Points can't be updated ",
		}, nil
	}

}
