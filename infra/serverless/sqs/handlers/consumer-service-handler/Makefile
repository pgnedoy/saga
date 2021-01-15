aws-prefix=aws --endpoint-url=http://localhost:4566 --region us-east-1

FUNCTION_NAME?=consumer-service-handler

build:
	GOARCH=amd64 GOOS=linux go build -ldflags="-d -s -w" -o bin/main main.go
	chmod +x bin/main
	zip -j bin/main.zip bin/main

clean:
	rm -r bin

update:
	$(MAKE) build
	$(MAKE) update-function-code

update-function-code:
	$(aws-prefix) lambda update-function-code \
		--function-name=$(FUNCTION_NAME) \
		--zip-file fileb://bin/main.zip


