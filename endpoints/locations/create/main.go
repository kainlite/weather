package main

import (
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

type Input struct {
	LocationId string `form:"location_id" json:"location_id" binding:"required"`
	UserId     string `form:"user_id" json:"user_id" binding:"required"`
}

func processRequest(c *gin.Context) {
	var input Input
	c.BindJSON(&input)
	location := weatherapi.CreateLocation(input.LocationId, input.UserId)
	c.JSON(http.StatusCreated, location)
}

func main() {
	lambda.Start(Handler)
}
