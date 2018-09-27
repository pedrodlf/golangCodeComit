package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("eu-west-1"))

func putProduct(account *T4wAccount) error {
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
