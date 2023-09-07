package adapters

import (
	"github.com/ydimitriou/eth-blockchain-parser/internal/adapters/ethereum"
	"github.com/ydimitriou/eth-blockchain-parser/internal/adapters/storage/memory"
	"github.com/ydimitriou/eth-blockchain-parser/internal/app/parser"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
)

// Services contains the available adapters services
type Services struct {
	BlockRepository      block.Repository
	SubscriberRepository subscriber.Repository
	EthereumService      parser.Service
}

// NewServices instantiates the adapter services
func NewServices() Services {
	return Services{
		BlockRepository:      memory.NewBlockRepository(),
		SubscriberRepository: memory.NewSubscriberRepository(),
		EthereumService:      ethereum.NewEthereumService(),
	}
}
