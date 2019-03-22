package weatherapi

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/satori/go.uuid"
)

var db *dynamodb.DynamoDB
var tableName string

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

func CreateLocation(inputLocation Location) (error, Location) {
	id := uuid.Must(uuid.NewV4(), nil).String()

	// Initialize location
	location := &Location{
		Id:         id,
		UserId:     inputLocation.UserId,
		LocationId: inputLocation.LocationId,
		CreatedAt:  time.Now().String(),
	}

	// Write to DynamoDB
	item, _ := dynamodbattribute.MarshalMap(location)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: &tableName,
	}

	_, err := db.PutItem(input)

	return err, *location
}

func DeleteLocation(locationkey LocationKey) error {
	fmt.Println("DeleteLocation")

	key, err := dynamodbattribute.MarshalMap(LocationKey{
		UserId:     locationkey.UserId,
		LocationId: locationkey.LocationId,
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	input := &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: aws.String("Locations"),
	}

	_, err = db.DeleteItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return err
}

func ListLocations(userId string) (error, []Location) {
	fmt.Println("ListLocations")

	var err error

	// Read from DynamoDB
	input := &dynamodb.ScanInput{
		TableName: &tableName,
	}
	result, _ := db.Scan(input)

	// Construct locations from response
	var locations []Location
	for _, i := range result.Items {
		location := Location{}
		if err := dynamodbattribute.UnmarshalMap(i, &location); err != nil {
			fmt.Println("Failed to unmarshal")
			fmt.Println(err)
		}
		fmt.Printf("%+v", location)
		locations = append(locations, location)
	}

	return err, locations
}

func SearchLocation(name string) (error, []Location) {
	// TODO: Query external endpoints for ids
	var locations []Location
	readJSONFile("example_search.json", &locations)

	return nil, locations
}
