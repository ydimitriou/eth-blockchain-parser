package ports

import (
	"github.com/ydimitriou/eth-blockchain-parser/internal/app"
	"github.com/ydimitriou/eth-blockchain-parser/internal/pkg/hex"
	"github.com/ydimitriou/eth-blockchain-parser/internal/ports/http"
	"github.com/ydimitriou/eth-blockchain-parser/internal/ports/infra"
)

// Services contains ports services
type Services struct {
	HTTPServer http.Server
	Worker     infra.Worker
}

// NewServices instantiates ports services
func NewServices(as app.Services, hp hex.Provider) Services {
	return Services{
		HTTPServer: http.NewServer(as),
		Worker:     infra.NewWorker(as, hp),
	}
}
