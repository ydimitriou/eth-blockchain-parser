package infra

import (
	"context"
	"log"
	"time"

	"github.com/ydimitriou/eth-blockchain-parser/internal/app"
	blockCommands "github.com/ydimitriou/eth-blockchain-parser/internal/app/block/commands"
	blockQueries "github.com/ydimitriou/eth-blockchain-parser/internal/app/block/queries"
	"github.com/ydimitriou/eth-blockchain-parser/internal/app/parser"
	subCommands "github.com/ydimitriou/eth-blockchain-parser/internal/app/subscriber/commands"
	subQueries "github.com/ydimitriou/eth-blockchain-parser/internal/app/subscriber/queries"
	"github.com/ydimitriou/eth-blockchain-parser/internal/pkg/hex"
)

// Worker responsible for polling eth blockcain to get new blocks,
// extract transactions and update subscribers for inbound or outbound transactions in memory
type Worker struct {
	blockServices  app.BlockServices
	subServices    app.SubscriberServices
	parserServices parser.Service
	interval       time.Duration
	hexProvider    hex.Provider
}

// NewWorker constructor
func NewWorker(as app.Services, hp hex.Provider) Worker {
	return Worker{
		blockServices:  as.BlockServices,
		subServices:    as.SubscriberServices,
		parserServices: as.ParserServices,
		interval:       5 * time.Second,
		hexProvider:    hp,
	}
}

func (w *Worker) Run() {
	ctx := context.Background()
	err := w.parseFirstBlock(ctx)
	if err != nil {
		log.Printf("parseFirstBlock error: %v", err.Error())
	}

	ticker := time.NewTicker(w.interval)
	for range ticker.C {
		ethBlockNum, err := w.getBlockNumberFromBlockchain(ctx)
		if err != nil {
			log.Printf("getBlockNumberFromBlockchain error: %v", err.Error())
		}
		latestParsedBlock, err := w.getCurrentBlockFromStorage()
		if err != nil {
			log.Printf("getCurrentStorageBlock error: %v", err.Error())
		}
		nextBlockNum := latestParsedBlock.Number + 1
		if nextBlockNum <= w.hexProvider.HexToInt(*ethBlockNum).Int64() {
			blockNum := w.hexProvider.IntToHex(nextBlockNum)
			err := w.processBlock(ctx, blockNum)
			if err != nil {
				log.Printf("processBlock error: %v", err.Error())
			}
		}
	}
}

// parseFirstBlock runs only the first time that the worker starts to set the first parsed block in memory
func (w *Worker) parseFirstBlock(ctx context.Context) error {
	ethBlockNum, err := w.getBlockNumberFromBlockchain(ctx)
	if err != nil {
		return err
	}
	return w.processBlock(ctx, *ethBlockNum)
}

// processBlock requests the transactions from eth blockchain for the given block number and updates subscribers in memory if they have transactions in this block
func (w *Worker) processBlock(ctx context.Context, blockNumber string) error {
	block, err := w.parserServices.GetBlockByNumber(ctx, blockNumber)
	if err != nil {
		return err
	}
	blockTransactions := block.Result.Transactions
	for _, tx := range blockTransactions {
		if w.isSubscriber(tx.From) {
			err := w.updateSubscriber(tx.From, tx)
			if err != nil {
				return err
			}
		}

		if w.isSubscriber(tx.To) {
			err := w.updateSubscriber(tx.To, tx)
			if err != nil {
				return err
			}
		}
	}
	return w.addBlockInStorage(blockNumber)
}

// updateSubscriber updates subscriber transactions for the given address using the app SubscriberServices
func (w *Worker) updateSubscriber(subAddress string, tx parser.Transaction) error {
	fromAddress := subQueries.IsSubscriberRequest{Address: subAddress}
	subscriber, err := w.subServices.Queries.GetSubscriberHandler.Handle(subQueries.GetSubscriberRequest(fromAddress))
	if err != nil {
		return err
	}
	transactionRequest := subCommands.TransactionsRequest{From: tx.From, To: tx.To}
	updateSubscriberReq := subCommands.UpdateSubscriberRequest{Address: subscriber.Address, Transaction: transactionRequest}
	return w.subServices.Commands.UpdateSubscriberHandler.Handle(updateSubscriberReq)
}

// isSubscriber checks if there is a subscriber in memory using the app SubscriberServices
func (w *Worker) isSubscriber(address string) bool {
	isSubReq := subQueries.IsSubscriberRequest{Address: address}
	return w.subServices.Queries.IsSubscriberHandler.Handle(isSubReq)
}

// getBlockNumberFromBlockchain return the latest block number on ethereum block chain using the app ParserServices
func (w *Worker) getBlockNumberFromBlockchain(ctx context.Context) (*string, error) {
	block, err := w.parserServices.GetCurrentBlock(ctx)
	if err != nil {
		return nil, err
	}

	return block, err
}

// getCurrentBlockFromStorage returns the latest parsed block from memory using the app BlockServices
func (w *Worker) getCurrentBlockFromStorage() (*blockQueries.GetBlockResult, error) {
	res, err := w.blockServices.Queries.GetBlockHandler.Handle()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// addBlockInStorage add block in memory using the app BlockServices
func (w *Worker) addBlockInStorage(blocknumber string) error {
	req := blockCommands.AddBlockRequest{Number: blocknumber}
	return w.blockServices.Commands.AddBlockHandler.Handle(req)
}
