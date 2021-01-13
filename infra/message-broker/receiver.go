package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	s := session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("foo", "var", ""),
		Region:           aws.String(endpoints.UsEast1RegionID),
		Endpoint:         aws.String("http://localhost:4566"),
	}))

	sqsClient := sqs.New(s, &aws.Config{})

	getQueueUrlOutput, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("test-queue"),
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	receiveMessageOutput, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: getQueueUrlOutput.QueueUrl,
		AttributeNames: aws.StringSlice([]string{
			"SentTimestamp",
		}),
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: aws.Int64(20),
	})

	fmt.Println(receiveMessageOutput)

}