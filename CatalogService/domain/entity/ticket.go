package entity

type Ticket struct {
	TicketId int `json:"ticket_id"`
	EventId  int `json:"event_id"`
	Quantity int `json:"amount"`
}
