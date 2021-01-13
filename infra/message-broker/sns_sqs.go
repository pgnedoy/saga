package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

const (
	SqsEndpoint = "http://localhost:4566"
)

const (
	TestTopic = "test-topic"
)

const (
	TestQueue = "test-queue"
)

func createSqsQueue(sess *session.Session) {
	sqsClient := sqs.New(sess, &aws.Config{})
	queueAtr, err := sqsClient.CreateQueue(&sqs.CreateQueueInput{
		QueueName:  aws.String(TestQueue),
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(queueAtr)

	q, err := sqsClient.ListQueues(&sqs.ListQueuesInput{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}

func createSnsTopics(sess *session.Session) {
	snsClient := sns.New(sess, &aws.Config{})

	result, err := snsClient.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(TestTopic),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.TopicArn)
}

func subscribeToTopics(sess *session.Session) {
	snsClient := sns.New(sess, &aws.Config{})

	snsClient.Subscribe(&sns.SubscribeInput{
		Attributes:            nil,
		Endpoint:              nil,
		Protocol:              aws.String("sqs"),
		ReturnSubscriptionArn: nil,
		TopicArn:              nil,
	})
}

func main() {
	s := session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("foo", "var", ""),
		Region:           aws.String(endpoints.UsEast1RegionID),
		Endpoint:         aws.String(SqsEndpoint),
	}))

	snsClient := sns.New(s, &aws.Config{})
	sqsClient := sqs.New(s, &aws.Config{})

	queueAtr, err := sqsClient.CreateQueue(&sqs.CreateQueueInput{
		QueueName:  aws.String(TestQueue),
	})

	if err != nil {
		fmt.Println(err)
	}

	topicAtr, err := snsClient.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(TestTopic),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	subscriptionAtr, err := snsClient.Subscribe(&sns.SubscribeInput{
		Attributes:            nil,
		Endpoint:              queueAtr.QueueUrl,
		Protocol:              aws.String("sqs"),
		ReturnSubscriptionArn: aws.Bool(true),
		TopicArn:              topicAtr.TopicArn,
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*queueAtr.QueueUrl)
	fmt.Println(*topicAtr.TopicArn)
	fmt.Println(subscriptionAtr)
}
