package handlers

import (
	"net/http"

	"github.com/shnartho/go-serverless-api-gateway-dynamodb/pkg/user"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbifac"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"` // *string -> can be filed w/ nil value if no value assigned to it 
}

// this functions will call functions inside the user.go which talks to the db
func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbifac.DynamoDBAPI)(*events.APIGatewayProxyResponse, error){
	
	email := req.QueryStringParameters{"email"}
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}

	result, err := user.FetchUsers(tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, result)
}


func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbifac.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error
){
	
}


func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbifac.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error
){
	
}


func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbifac.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error
){
	
}


func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}

