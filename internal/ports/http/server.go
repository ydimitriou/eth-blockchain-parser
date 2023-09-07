package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ydimitriou/eth-blockchain-parser/internal/app"
	"github.com/ydimitriou/eth-blockchain-parser/internal/ports/http/block"
	"github.com/ydimitriou/eth-blockchain-parser/internal/ports/http/subscriber"
)

// Server is the http server
type Server struct {
	appServices app.Services
	router      *mux.Router
}

// NewServer HTTP server constructor
func NewServer(as app.Services) Server {
	router := mux.NewRouter()
	httpServer := Server{
		appServices: as,
		router:      router,
	}
	httpServer.createHTTPRoutes()
	http.Handle("/", httpServer.router)

	return httpServer
}

// createHTTPRoutes generates routes
func (httpServer *Server) createHTTPRoutes() {
	httpServer.router.HandleFunc("/subscriber", subscriber.NewHandler(httpServer.appServices.SubscriberServices).Create).Methods("POST")
	httpServer.router.HandleFunc("/subscriber"+"/{"+subscriber.GetTransactionsURLParam+"}", subscriber.NewHandler(httpServer.appServices.SubscriberServices).GetTransactions).Methods("GET")
	httpServer.router.HandleFunc("/last-block", block.NewHandler(httpServer.appServices.BlockServices).GetLast).Methods("GET")
}

// ListenAndServer wraps HTTP listenAndServe (initiate listening for request)
func (httpServer *Server) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
