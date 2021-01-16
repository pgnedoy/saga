package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

func deleteMessageBatch(ctx context.Context, wg *sync.WaitGroup, entries []*sqs.DeleteMessageBatchRequestEntry) {
	_, err := GetSqsConnection().DeleteMessageBatchWithContext(ctx, &sqs.DeleteMessageBatchInput{
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
	executeRequest := NewRequestExecutor(ctx)
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
