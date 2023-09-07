package infra

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ydimitriou/eth-blockchain-parser/internal/app"
	blockCommands "github.com/ydimitriou/eth-blockchain-parser/internal/app/block/commands"
	blockQueries "github.com/ydimitriou/eth-blockchain-parser/internal/app/block/queries"
	subCommands "github.com/ydimitriou/eth-blockchain-parser/internal/app/subscriber/commands"
	subQueries "github.com/ydimitriou/eth-blockchain-parser/internal/app/subscriber/queries"
	"github.com/ydimitriou/eth-blockchain-parser/internal/pkg/hex"
)

type Worker struct {
	appServices app.Services
	interval    time.Duration
	isFirstRun  bool
	hexProvider hex.Provider
}

func NewWorker(as app.Services, hp hex.Provider) Worker {
	return Worker{
		appServices: as,
		interval:    5 * time.Second,
		isFirstRun:  true,
		hexProvider: hp,
	}
}

func (w *Worker) Run() {
	ctx := context.Background()
	ticker := time.NewTicker(w.interval)

	for range ticker.C {
		if w.isFirstRun {
			parsedBlock, err := w.parseBlock(ctx)
			if err != nil {
				log.Printf(err.Error())
			}
			w.processBlock(ctx, *parsedBlock)
			w.isFirstRun = false
		} else {
			parsedBlock, err := w.parseBlock(ctx)
			if err != nil {
				log.Printf(err.Error())
			}
			currentBlockInStorage, err := w.getCurrentStorageBlock()
			if err != nil {
				log.Printf(err.Error())
			}
			if currentBlockInStorage.Number+1 <= w.hexProvider.HexToInt(*parsedBlock).Int64() {
				blockNum := w.hexProvider.IntToHex(currentBlockInStorage.Number + 1)
				w.processBlock(ctx, blockNum)
			}
		}
	}
}

func (w *Worker) processBlock(ctx context.Context, blockNumber string) error {
	block, err := w.appServices.ParserServices.GetBlockByNumber(ctx, blockNumber)
	fmt.Println("Get transaction details from eth blockhain for block: ", w.hexProvider.HexToInt(blockNumber))
	if err != nil {
		log.Printf(err.Error())
	}
	blockTransactions := block.Result.Transactions
	fmt.Println("Block transactions are: ", blockTransactions)
	for _, tx := range blockTransactions {
		fromAddress := subQueries.IsSubscriberRequest{Address: tx.From}
		if w.appServices.SubscriberServices.Queries.IsSubscriberHandler.Handle(fromAddress) {
			subscriber, err := w.appServices.SubscriberServices.Queries.GetSubscriberHandler.Handle(subQueries.GetSubscriberRequest(fromAddress))
			if err != nil {
				return err
			}
			transactionRequest := subCommands.TransactionsRequest{From: tx.From, To: tx.To}
			updateSubscriberReq := subCommands.UpdateSubscriberRequest{Address: subscriber.Address, Transaction: transactionRequest}
			err = w.appServices.SubscriberServices.Commands.UpdateSubscriberHandler.Handle(updateSubscriberReq)
			if err != nil {
				return err
			}
		}

		toAddress := subQueries.IsSubscriberRequest{Address: tx.To}
		if w.appServices.SubscriberServices.Queries.IsSubscriberHandler.Handle(toAddress) {
			subscriber, err := w.appServices.SubscriberServices.Queries.GetSubscriberHandler.Handle(subQueries.GetSubscriberRequest(toAddress))
			if err != nil {
				return err
			}
			transactionRequest := subCommands.TransactionsRequest{From: tx.From, To: tx.To}
			updateSubscriberReq := subCommands.UpdateSubscriberRequest{Address: subscriber.Address, Transaction: transactionRequest}
			err = w.appServices.SubscriberServices.Commands.UpdateSubscriberHandler.Handle(updateSubscriberReq)
			if err != nil {
				return err
			}
		}
	}
	err = w.addBlockInStorage(blockNumber)
	if err != nil {
		return err
	}

	return nil
}

func (w *Worker) parseBlock(ctx context.Context) (*string, error) {
	block, err := w.appServices.ParserServices.GetCurrentBlock(ctx)
	fmt.Println("Parsing from eth blockchain block: ", w.hexProvider.HexToInt(*block))
	if err != nil {
		return nil, err
	}

	return block, err
}

func (w *Worker) getCurrentStorageBlock() (*blockQueries.GetBlockResult, error) {
	res, err := w.appServices.BlockServices.Queries.GetBlockHandler.Handle()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (w *Worker) addBlockInStorage(blocknumber string) error {
	req := blockCommands.AddBlockRequest{Number: blocknumber}
	err := w.appServices.BlockServices.Commands.AddBlockHandler.Handle(req)
	if err != nil {
		return err
	}

	return nil
}
