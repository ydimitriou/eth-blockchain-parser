package queries

import "github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"

// IsSubscriberRequest Model of IsSubscriberHandler
type IsSubscriberRequest struct {
	Address string
}

// IsSubscriberHandler interface for checking if an address is already a subscriber
type IsSubscriberHandler interface {
	Handle(req IsSubscriberRequest) bool
}

type isSubscriberHandler struct {
	repo subscriber.Repository
}

// NewIsSubscriberHandler constructor
func NewIsSubscriberHandler(repo subscriber.Repository) IsSubscriberHandler {
	return isSubscriberHandler{repo: repo}
}

// Handle handles is subscriber requests
func (h isSubscriberHandler) Handle(req IsSubscriberRequest) bool {
	return h.repo.Exist(req.Address)
}
