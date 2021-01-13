prefix

build:
	GOARCH=amd64 GOOS=linux go build -ldflags="-d -s -w" -o bin/main main.go
	chmod +x bin/main
	zip -j bin/main.zip bin/main

create-event-source-mapping:
	aws --endpoint-url=http://localhost:4566 --region us-east-1 lambda create-event-source-mapping --function-name=my-function --batch-size 5 --maximum-batching-window-in-seconds 60 --event-source-arn arn:aws:sqs:us-east-1:000000000000:test-queue

delete-event-source-mapping:
	aws --endpoint-url=http://localhost:4566 --region us-east-1 lambda delete-event-source-mapping --uuid="81cc2c11-09ab-463a-bf42-8ce3e2acedfa"

create-lambda:
	aws --endpoint-url=http://localhost:4566 --region us-east-1 lambda create-function --function-name=my-function --runtime="go1.x" --role=fakerole --handler=main --zip-file fileb://bin/main.zip

list-event-source-mappings:
	aws --endpoint-url=http://localhost:4566 --region us-east-1 lambda list-event-source-mappings

list-functions:
	aws --endpoint-url=http://localhost:4566 --region us-east-1 lambda list-functions

update:
	aws --endpoint-url=http://localhost:4566 --region us-east-1 lambda update-function-code --function-name=my-function --zip-file fileb://bin/main.zip

update-function-configuration:
	aws --endpoint-url=http://localhost:4566 --region us-east-1 lambda update-function-configuration --function-name=my-function --handler=index.handler

upload:
	make build && make update

invoke:
	aws --endpoint-url=http://localhost:4566 --region us-east-1 lambda invoke --function-name=my-function bin/response.json

remove:
	aws --endpoint-url=http://localhost:4566 --region us-east-1 lambda delete-function --function-name=my-function