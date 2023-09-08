package http

import (
	"log"
	"net/http"
	"strings"

	"github.com/ydimitriou/eth-blockchain-parser/internal/app"
	"github.com/ydimitriou/eth-blockchain-parser/internal/ports/http/block"
	"github.com/ydimitriou/eth-blockchain-parser/internal/ports/http/subscriber"
)

// Server is the http server
type Server struct {
	appServices app.Services
	mux         *http.ServeMux
}

// NewServer HTTP server constructor
func NewServer(as app.Services) Server {
	mux := http.NewServeMux()
	httpServer := Server{
		appServices: as,
		mux:         mux,
	}

	httpServer.createHTTPRoutes()
	http.Handle("/", httpServer.mux)

	return httpServer
}

// createHTTPRoutes generates routes
func (httpServer *Server) createHTTPRoutes() {
	subscribeHandler := subscriber.NewHandler(httpServer.appServices.SubscriberServices)
	blockHandler := block.NewHandler(httpServer.appServices.BlockServices)

	httpServer.mux.HandleFunc(subscriber.BasePath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			subscribeHandler.Create(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	httpServer.mux.HandleFunc(subscriber.BasePath+"/", func(w http.ResponseWriter, r *http.Request) {
		pathSegments := strings.Split(r.URL.Path, "/")
		if len(pathSegments) > 3 {
			subAddress := pathSegments[3]
			if r.Method == http.MethodGet {
				subscribeHandler.GetTransactions(w, r, subAddress)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		} else {
			http.NotFound(w, r)
		}
	})

	httpServer.mux.HandleFunc(block.BasePath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			blockHandler.GetLast(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

// ListenAndServer wraps HTTP listenAndServe (initiate listening for request)
func (httpServer *Server) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
