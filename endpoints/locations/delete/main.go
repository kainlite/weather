package main

import (
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
		ginEngine := weatherapi.MountAuthorizedRoute("/locations/:id", "delete", processRequest)
		ginLambda = ginadapter.New(ginEngine)
		initialized = true
	}
	return ginLambda.Proxy(req)
}

func processRequest(c *gin.Context) {
	fmt.Println("Delete")

	location := weatherapi.DeleteLocation(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else {
		return c.JSON(http.StatusNoContent, nil)
	}
}

func main() {
	lambda.Start(Handler)
}
