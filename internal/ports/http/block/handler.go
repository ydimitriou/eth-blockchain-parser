package block

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ydimitriou/eth-blockchain-parser/internal/app"
)

// Handler is block http request handler
type Handler struct {
	blockServices app.BlockServices
}

// NewHandler return an http request handler
func NewHandler(bs app.BlockServices) Handler {
	return Handler{blockServices: bs}
}

// GetLast returns the last parsed block
func (h Handler) GetLast(w http.ResponseWriter, _ *http.Request) {
	block, err := h.blockServices.Queries.GetBlockHandler.Handle()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(block)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
}
