package queries

import "github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"

type IsSubscriberRequest struct {
	Address string
}

type IsSubscriberHandler interface {
	Handle(req IsSubscriberRequest) bool
}

type isSubscriberHandler struct {
	repo subscriber.Repository
}

func NewIsSubscriberHandler(repo subscriber.Repository) IsSubscriberHandler {
	return isSubscriberHandler{repo: repo}
}

func (h isSubscriberHandler) Handle(req IsSubscriberRequest) bool {
	return h.repo.Exist(req.Address)
}
