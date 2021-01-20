package events

const (
	MessageAttributesEventType = "event"
	MessageAttributesCorrelationID = "correlation-id"
)

type SagaEvent string

const (
	VerifyConsumer SagaEvent = "VERIFY_CONSUMER"
	ConsumerVerified SagaEvent = "CONSUMER_VERIFIED"
	CreateTicket SagaEvent = "CREATE_TICKET"
	TicketCreated SagaEvent = "TICKET_CREATED"
	AuthorizeCard SagaEvent = "AUTHORIZE_CARD"
	CardAuthorized SagaEvent = "CARD_AUTHORIZED"
	ApproveRestaurantOrder SagaEvent = "APPROVE_RESTAURANT_ORDER"
	ApproveOrder SagaEvent = "APPROVE_ORDER"
	CreateOrder SagaEvent = "ORDER_CREATED"
)
