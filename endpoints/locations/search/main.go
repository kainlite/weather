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

type SearchParams struct {
	Name string `json:"name"`
}

var initialized = false
var ginLambda *ginadapter.GinLambda

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !initialized {
		ginEngine := weatherapi.MountAuthorizedRoute("/locations/search", "post", processRequest)
		ginLambda = ginadapter.New(ginEngine)
		initialized = true
	}
	return ginLambda.Proxy(req)
}

func processRequest(c *gin.Context) {
	fmt.Println("Search")

	var params SearchParams
	c.BindJSON(&params)

	err, locations := weatherapi.SearchLocation(params.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, locations)
	}
}

func main() {
	lambda.Start(Handler)
}
