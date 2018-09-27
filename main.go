package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//T4wAccount cuenta de usuario
type T4wAccount struct {
	AUserID      uint32 `json:"userID"`
	InitialOints uint32 `json:"points"`
	Key          string `json:"key"`
	Address      string `json:"address"`
	AURL         string `json: "url"`
	MeetingID    uint32 `json:"meetingID"`
}

// User datos de Usuario
type User struct {
	UserID       uint32 `json:"userID"`
	InitialOints uint32 `json:"points"`
	MeetingID    uint32 `json:"meetingID"`
}

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

	user := new(User)
	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		log.Print(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Smart Contract : Unsupported Media Type ",
		}, nil
	}

	tfwAccount := new(T4wAccount)
	tfwAccount.AUserID = user.UserID
	tfwAccount.InitialOints = user.InitialOints
	tfwAccount.MeetingID = user.MeetingID
	log.Printf("tfwAccount : %v", tfwAccount)
	log.Printf("AUser : %v", user)
	log.Printf("tfwAccount.AUserID : %v", user)

	tfwAccount.Key, tfwAccount.Address, tfwAccount.AURL, err = CreateAccount(user.UserID)
	log.Printf("tfwAccount.Key : %v", tfwAccount.Key)
	if err != nil {

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Account error :  cant get new account",
		}, nil

	}

	err = putProduct(tfwAccount)

	if err != nil {

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusConflict,
			Body:       "Account error :  cant save the account",
		}, nil

	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Account: Created ",
	}, nil
}
