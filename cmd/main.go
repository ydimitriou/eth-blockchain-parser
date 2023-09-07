package main

import (
	"github.com/ydimitriou/eth-blockchain-parser/internal/adapters"
	"github.com/ydimitriou/eth-blockchain-parser/internal/app"
	"github.com/ydimitriou/eth-blockchain-parser/internal/pkg/hex"
	"github.com/ydimitriou/eth-blockchain-parser/internal/ports"
)

func main() {
	hp := hex.NewHexProvider()
	adapters := adapters.NewServices()
	appServices := app.NewServices(adapters.BlockRepository, adapters.SubscriberRepository, adapters.EthereumService, hp)
	ports := ports.NewServices(appServices, hp)
	go ports.Worker.Run()
	ports.HTTPServer.ListenAndServe(":8080")
}
