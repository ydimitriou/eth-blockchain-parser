package subscriber

// Repository interface for subscribers
type Repository interface {
	Add(subscriber Subscriber) error
	Update(subscriber Subscriber) error
	Exist(address string) bool
	GetByAddress(address string) (*Subscriber, error)
}
