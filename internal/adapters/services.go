package adapters

import (
	"github.com/ydimitriou/eth-blockchain-parser/internal/adapters/ethereum"
	"github.com/ydimitriou/eth-blockchain-parser/internal/adapters/storage/memory"
	"github.com/ydimitriou/eth-blockchain-parser/internal/app/parser"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
)

type Services struct {
	BlockRepository      block.Repository
	SubscriberRepository subscriber.Repository
	EthereumService      parser.Service
}

func NewServices() Services {
	return Services{
		BlockRepository:      memory.NewBlockRepository(),
		SubscriberRepository: memory.NewSubscriberRepository(),
		EthereumService:      ethereum.NewEthereumService(),
	}
}
