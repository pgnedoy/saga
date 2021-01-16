package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//func getEndpoint() {}

func HandleRequest(ctx context.Context, event events.SQSEvent) error {
	for _, record := range event.Records {
		fmt.Println(record)
	}
	fmt.Println("End!")
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
