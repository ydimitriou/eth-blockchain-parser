package app

import (
	blockCommands "github.com/ydimitriou/eth-blockchain-parser/internal/app/block/commands"
	blockQueries "github.com/ydimitriou/eth-blockchain-parser/internal/app/block/queries"
	"github.com/ydimitriou/eth-blockchain-parser/internal/app/parser"
	subCommands "github.com/ydimitriou/eth-blockchain-parser/internal/app/subscriber/commands"
	subQueries "github.com/ydimitriou/eth-blockchain-parser/internal/app/subscriber/queries"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/block"
	"github.com/ydimitriou/eth-blockchain-parser/internal/domain/subscriber"
	"github.com/ydimitriou/eth-blockchain-parser/internal/pkg/hex"
)

type SubscriberCommands struct {
	AddSubscriberHandler    subCommands.AddSubscriberHandler
	UpdateSubscriberHandler subCommands.UpdateSubscriberHandler
}

type SubscriberQueries struct {
	GetSubscriberHandler subQueries.GetSubscriberHandler
	IsSubscriberHandler  subQueries.IsSubscriberHandler
}

type SubscriberServices struct {
	Commands SubscriberCommands
	Queries  SubscriberQueries
}

type BlockCommands struct {
	AddBlockHandler blockCommands.AddBlockHandler
}

type BlockQueries struct {
	GetBlockHandler blockQueries.GetBlockHandler
}

type BlockServices struct {
	Commands BlockCommands
	Queries  BlockQueries
}

type Services struct {
	SubscriberServices SubscriberServices
	BlockServices      BlockServices
	ParserServices     parser.Service
}

func NewServices(blockRepo block.Repository, subRepo subscriber.Repository, parser parser.Service, hp hex.Provider) Services {
	return Services{
		SubscriberServices: SubscriberServices{
			Commands: SubscriberCommands{
				AddSubscriberHandler:    subCommands.NewAddSubscriberHandler(subRepo),
				UpdateSubscriberHandler: subCommands.NewUpdateSubscriberHandler(subRepo),
			},
			Queries: SubscriberQueries{
				GetSubscriberHandler: subQueries.NewGetSubscriberHandler(subRepo),
				IsSubscriberHandler:  subQueries.NewIsSubscriberHandler(subRepo),
			},
		},
		BlockServices: BlockServices{
			Commands: BlockCommands{
				AddBlockHandler: blockCommands.NewAddBlockHandler(blockRepo),
			},
			Queries: BlockQueries{
				GetBlockHandler: blockQueries.NewGetBlockHandler(blockRepo, hp),
			},
		},
		ParserServices: parser,
	}
}
