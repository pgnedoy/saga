package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

const (
	EventTypeFieldName = "eventType"
)

type Event struct {
	Event string `json:"event"`
	Body interface{} `json:"body"`
}

type Req struct {
	Url string
	Event
}

var s = session.Must(session.NewSession(&aws.Config{
	Credentials:      credentials.NewStaticCredentials("foo", "var", ""),
	Region:           aws.String(endpoints.UsEast1RegionID),
	Endpoint:         aws.String("http://localhost:4566"),
}))

var sqsClient = sqs.New(s, &aws.Config{})

func validateMessage(msg *events.SQSMessage) error {
	eventType, ok := msg.MessageAttributes[EventTypeFieldName]
	if !ok {
		return errors.New(fmt.Sprintf("MessageID: %s, Err: %s attribute is required!",
			msg.MessageId, EventTypeFieldName))
	}

	if eventType.StringValue == nil {
		return errors.New(fmt.Sprintf("MessageID: %s, Err: %s value  is required!",
			msg.MessageId, EventTypeFieldName))
	}

	if eventType.DataType != "String" {
		return errors.New(fmt.Sprintf("MessageID: %s, Err: %s type must be String!",
			msg.MessageId, EventTypeFieldName))
	}

	return nil
}

func generateRequest(event *events.SQSMessage) (*http.Request, error) {
	eventType, _ := event.MessageAttributes[EventTypeFieldName]

	switch *eventType.StringValue {
	// TODO: create events
	case "TEST_EVENT":
		// TODO: add event creating logic
		break
	default:
		return nil, errors.New(fmt.Sprintf("MessageID: %s, Err: %s",
			event.MessageId, "Unsupported event type value!!!"))
	}

	return &http.Request{}, nil
}

func createRequestExecutor(ctx context.Context) func(chan *events.SQSMessage, *events.SQSMessage) {
	client := &http.Client{}

	return func(failed chan *events.SQSMessage, record *events.SQSMessage) {
		ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
		defer cancel()

		err := validateMessage(record)
		if err != nil {
			failed <- record
		}

		req, err := generateRequest(record)
		if err != nil {
			failed <- record
		}

		_, err = client.Do(req.WithContext(ctx))
		if err != nil {
			failed <- record
		} else {
			failed <- nil
		}
		return
	}
}

func deleteMessageBatch(ctx context.Context, wg *sync.WaitGroup, entries []*sqs.DeleteMessageBatchRequestEntry) {
	_, err := sqsClient.DeleteMessageBatchWithContext(ctx, &sqs.DeleteMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(""),
	})
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}

func sendMessagesToDeadLetterQueue(ctx context.Context, wg *sync.WaitGroup) {
	// TODO here
	wg.Done()
}

func HandleEvent(ctx context.Context, event events.SQSEvent) error {
	records := event.Records
	recsLen := len(records)
	
	if recsLen == 0 {
		return nil
	}

	failedChannel := make(chan *events.SQSMessage)
	executeRequest := createRequestExecutor(ctx)
	for _, record := range records {
		go executeRequest(failedChannel, &record)
	}

	failedMessages := make([]*events.SQSMessage, 0, recsLen)
	for rec := range failedChannel {
		if rec != nil {
			failedMessages = append(failedMessages, rec)
		}
	}

	batchID := uuid.New().String()
	entries := make([]*sqs.DeleteMessageBatchRequestEntry, 0, recsLen)
	for _, rec := range records {
		entries = append(entries, &sqs.DeleteMessageBatchRequestEntry{
			Id:            aws.String(batchID),
			ReceiptHandle: aws.String(rec.ReceiptHandle),
		})
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go deleteMessageBatch(ctx, &wg, entries)
	go sendMessagesToDeadLetterQueue(ctx, &wg)
	wg.Wait()

	fmt.Println("End! That is consumer-service")
	return nil
}

func main() {
	lambda.Start(HandleEvent)
}
