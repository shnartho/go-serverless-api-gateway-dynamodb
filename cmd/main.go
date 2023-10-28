package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.Newsession(&aws.Config{
		Region: aws.String(region)})
}
