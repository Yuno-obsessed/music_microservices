package interfaces

type TicketInterface interface {
	IsInCatalog(ttype string) (bool, error)
	GetQuantityOfType(ttype string) (int, error)
	Subtruct(ttype string, number int) error
}
