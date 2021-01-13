aws-prefix=aws --endpoint-url=http://localhost:4566 --region us-east-1

##### AWS QUEUE #####

QUEUE_NAME?=test
QUEUE_URL?=test

services := consumer-service \
		kitchen-service \
		accounting-service \
		order-service \
		create-order-saga

queue_prefix := channel
handler_prefix := handler

queues := $(foreach srv, $(services), $(srv)-$(queue_prefix))
handlers := $(foreach srv, $(services), $(srv)-$(handler_prefix))
sqs_mapping := $(services)

up:
	$(MAKE) create-queues 2>&1 | tee logs/create-queues.logs
	$(MAKE) create-handlers 2>&1 | tee logs/create-handlers.logs
	$(MAKE) create-mapping 2>&1 | tee logs/create-mapping.logs

###############################

create-queues: $(queues)

$(queues):
	$(MAKE) create-queue QUEUE_NAME=$@

create-queue:
	$(aws-prefix) sqs create-queue \
		--queue-name=$(QUEUE_NAME)

list-queues:
	$(aws-prefix) sqs list-queues \
		--output=json 2>&1 | tee -a logs/list-queues.logs

send-message:

send-message-batch:

delete-queue:
	$(aws-prefix) sqs delete-queue \
		--queue-url=$(QUEUE_NAME)


##### AWS LAMBDA #####

FUNCTION_NAME?=my-function
HANDLER?=main
QUEUE_NAME?=test-queue
QUEUE_UUID?=

create-handlers: $(handlers)

$(handlers):
	$(MAKE) -C ./handlers/$@ build
	$(MAKE) create FUNCTION_NAME=$@

create:
	$(aws-prefix) lambda create-function \
		--function-name=$(FUNCTION_NAME) \
		--runtime="go1.x" \
		--role=fakerole \
		--handler=main \
		--zip-file=fileb://handlers/$(FUNCTION_NAME)/bin/main.zip

delete:
	$(aws-prefix) lambda delete-function \
		--function-name=$(FUNCTION_NAME)

update-function-code:
	$(aws-prefix) lambda update-function-code \
		--function-name=$(FUNCTION_NAME) \
		--zip-file fileb://bin/main.zip

upload:
	make build
	make create

update:
	make build
	make update-function-code

invoke:
	$(aws-prefix) lambda invoke \
		--function-name=$(FUNCTION_NAME) \
		bin/response.json

list-handlers:
	$(aws-prefix) lambda list-functions 2>&1 | tee -a logs/list-handlers.logs

update-function-configuration:
	$(aws-prefix) lambda update-function-configuration \
		--function-name=$(FUNCTION_NAME) \
		--handler=$(HANDLER)

############## sqs-event-sourcing

create-mapping: $(sqs_mapping)

$(sqs_mapping):
	$(MAKE) create-event-source-mapping FUNCTION_NAME=$@-$(handler_prefix) QUEUE_NAME=$@-$(queue_prefix)

create-event-source-mapping:
	$(aws-prefix) lambda create-event-source-mapping \
		--function-name=$(FUNCTION_NAME) \
		--batch-size=5 \
		--maximum-batching-window-in-seconds=60 \
		--event-source-arn=arn:aws:sqs:us-east-1:000000000000:$(QUEUE_NAME)

delete-event-source-mapping:
	$(aws-prefix) lambda delete-event-source-mapping \
		--uuid=$(QUEUE_UUID)

list-event-source-mappings:
	$(aws-prefix) lambda list-event-source-mappings 2>&1 | tee -a logs/list-event-source-mappings.logs





