package commands

import (
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
)

// AddSubscriberRequest Model of AddSubscriberHandler
type AddSubscriberRequest struct {
	Address string
}

// AddSubscriberHandler interface for adding an address as subscriber
type AddSubscriberHandler interface {
	Handle(req AddSubscriberRequest) error
}

type addSubscriberHandler struct {
	repo subscriber.Repository
}

// NewAddSubscriberHandler constructor
func NewAddSubscriberHandler(repo subscriber.Repository) AddSubscriberHandler {
	return addSubscriberHandler{repo: repo}
}

// Handle handles subscribe requests
func (h addSubscriberHandler) Handle(req AddSubscriberRequest) error {
	s := subscriber.Subscriber{
		Address: req.Address,
	}
	err := h.repo.Add(s)
	if err != nil {
		return err
	}

	return nil
}
