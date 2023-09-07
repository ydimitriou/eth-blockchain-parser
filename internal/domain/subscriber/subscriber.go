package subscriber

import "github.com/ydimitriou/eth-blockchain-parser/internal/domain/transaction"

// Subscriber Moder that represent the subscriber
type Subscriber struct {
	Address              string
	InboundTransactions  []transaction.Transaction
	OutboundTransactions []transaction.Transaction
}
