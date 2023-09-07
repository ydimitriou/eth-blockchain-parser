package memory

import (
	"fmt"

	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
)

type SubscriberRepository struct {
	subscribers map[string]subscriber.Subscriber
}

func NewSubscriberRepository() SubscriberRepository {
	return SubscriberRepository{
		subscribers: make(map[string]subscriber.Subscriber),
	}
}

func (mr SubscriberRepository) Add(sub subscriber.Subscriber) error {
	mr.subscribers[sub.Address] = sub

	return nil
}

func (mr SubscriberRepository) Update(sub subscriber.Subscriber) error {
	_, exists := mr.subscribers[sub.Address]
	if !exists {
		return fmt.Errorf("subscriber with address %v does not exist", sub.Address)
	}
	mr.subscribers[sub.Address] = sub

	return nil
}

func (mr SubscriberRepository) Exist(address string) bool {
	_, ok := mr.subscribers[address]

	return ok
}

func (mr SubscriberRepository) GetByAddress(address string) (*subscriber.Subscriber, error) {
	val, exists := mr.subscribers[address]
	if !exists {
		return nil, nil
	}

	return &val, nil
}
