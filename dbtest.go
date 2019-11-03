package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"os"
	//	"strconv"
)

type Contact struct {
	Phone   string
	Dir     string
	Station string
	Line    string
}

var tableName string = "db_test"
var PhoneN string = "15551234567"

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	fmt.Println("=addItem")
	addItem(svc)
	fmt.Println("=getItem")
	result := getItem(svc)
	fmt.Println(result)
}

func getItem(svc *dynamodb.DynamoDB) *dynamodb.GetItemOutput {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Phone": {
				S: aws.String(PhoneN),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	return result
}

func addItem(svc *dynamodb.DynamoDB) {
	contact := Contact{
		Phone:   PhoneN,
		Dir:     "s",
		Station: "MONT",
		Line:    "YELLOW",
	}

	av, err := dynamodbattribute.MarshalMap(contact)

	fmt.Println(av)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	put, err := svc.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(put)

}
