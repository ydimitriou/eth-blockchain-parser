package commands

import (
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
)

type AddSubscriberRequest struct {
	Address string
}

type AddSubscriberHandler interface {
	Handle(req AddSubscriberRequest) error
}

type addSubscriberHandler struct {
	repo subscriber.Repository
}

func NewAddSubscriberHandler(repo subscriber.Repository) AddSubscriberHandler {
	return addSubscriberHandler{repo: repo}
}

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
