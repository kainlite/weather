package weatherapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var db *dynamodb.DynamoDB
var tableName *aws.String

func init() {
	region := "us-east-1"
	tableName = "Locations"

	if session, err := session.NewSession(&aws.Config{
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		db = dynamodb.New(session)
	}
}

func CreateLocation(location weatherapi.Location) (error, weatherapi.Location) {
	id := uuid.Must(uuid.NewV4(), nil).String()

	// Initialize location
	location := &weatherapi.Location{
		ID:         id,
		UserId:     "todo",
		LocationId: "ARBA0009:1:AR",
		CreatedAt:  time.Now().String(),
	}

	// Parse request body
	json.Unmarshal([]byte(request.Body), location)

	// Write to DynamoDB
	item, _ := dynamodbattribute.MarshalMap(location)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: tableName,
	}

	_, err := db.PutItem(input)

	return err, location
}

func DeleteLocation(id string) (error, weatherapi.Location) {
	fmt.Println("DeleteLocation")

	// Delete location
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: tableName,
	}
	_, err := db.DeleteItem(input)

	return err, location
}

func ListLocations() []weatherapi.Location {
	fmt.Println("ListLocations")

	// Read from DynamoDB
	input := &dynamodb.ScanInput{
		TableName: tableName,
	}
	result, _ := db.Scan(input)

	// Construct locations from response
	var locations []weatherapi.Location
	for _, i := range result.Items {
		location := weatherapi.Location{}
		if err := dynamodbattribute.UnmarshalMap(i, &location); err != nil {
			fmt.Println("Failed to unmarshal")
			fmt.Println(err)
		}
		locations = append(locations, location)
	}

	return locations
}
