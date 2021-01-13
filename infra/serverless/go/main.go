package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() error {
	fmt.Println("Hello dev!!!")
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
