package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	s := session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("foo", "var", ""),
		Region:           aws.String(endpoints.UsEast1RegionID),
		Endpoint:         aws.String("http://localhost:4566"),
	}))

	snsClient := sns.New(s, &aws.Config{})
	input := &sns.PublishInput{
		Message:  aws.String("Hello world!"),
		TopicArn: aws.String("arn:aws:sns:us-east-1:000000000000:test-topic"),
	}

	result, err := snsClient.Publish(input)
	if err != nil {
		fmt.Println("Publish error:", err)
		return
	}

	fmt.Println(result)
}
