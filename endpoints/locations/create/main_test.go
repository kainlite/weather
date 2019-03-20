package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Location", func() {
	var (
		response events.APIGatewayProxyResponse
		request  events.APIGatewayProxyRequest
		err      error
	)

	JustBeforeEach(func() {
		response, err = Handler(request)
		Expect(err).To(BeNil())
	})

	It(`Returns a location object and an OK status`, func() {
		Expect(response.Body).To(Equal("{\"LocationId\":\"\",\"UserId\":\"\"}"))
		Expect(response.StatusCode).To(Equal(http.StatusOK))
	})
})
