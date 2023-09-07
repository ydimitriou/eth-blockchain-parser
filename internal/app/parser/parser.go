package parser

import "context"

type GetBlockByNumberResponse struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
}

type Result struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Service interface {
	GetCurrentBlock(ctx context.Context) (*string, error)
	GetBlockByNumber(ctx context.Context, blockNumber string) (*GetBlockByNumberResponse, error)
}
