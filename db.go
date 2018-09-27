package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ethereum/go-ethereum/common"
)

type T4wAccount struct {
	AUserID      uint32 `json:"userID"`
	InitialOints uint32 `json:"points"`
	Key          string `json:"key"`
	Address      string `json:"address"`
	AURL         string `json:"url"`
	MeetingID    uint32 `json:"meetingID"`
}

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("eu-west-1"))

func saveT4wAccount(account *T4wAccount) error {
	av, err := dynamodbattribute.MarshalMap(account)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("likedat-accounts"),
	}

	_, err = db.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

//GetUsserAddress by userID
func GetUsserAddress(userID string) (common.Address, error) {

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("likedat-accounts"),
		Key: map[string]*dynamodb.AttributeValue{
			"userID": {
				N: aws.String(userID),
			},
		},
	})

	if err != nil {
		log.Println(err.Error())
		log.Printf("User not found-->: %v", userID)
	}

	item := T4wAccount{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	addres := common.HexToAddress(item.Address)
	return addres, nil
}
