package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/kainlite/weather/weatherapi"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !initialized {
		ginEngine := weatherapi.MountAuthorizedRoute("/locations", "post", processRequest)
		ginLambda = ginadapter.New(ginEngine)
		initialized = true
	}
	return ginLambda.Proxy(req)
}

func processRequest(c *gin.Context) {
	fmt.Println("Create")

	var input weatherapi.Location
	c.BindJSON(&input)
	err, location := weatherapi.CreateLocation(input)

	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, nil)
	} else {
		body, _ := json.Marshal(location)

		c.JSON(http.StatusCreated, body)
	}
}

func main() {
	lambda.Start(Handler)
}
