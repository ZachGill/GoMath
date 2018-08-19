package handlers

import (
	"context"
	"log"
	"net/http"
	"sync"

	middleware "github.com/ZachGill/transaction-mw"
	"github.com/gorilla/mux"
)

// Server starts the applications HTTP Server
type Server struct {
	ServerMutex *sync.Mutex
	WaitGroup   *sync.WaitGroup

	HTTPListenAddr string
	HTTPLogger     *log.Logger

	// Addition, subtraction, multiplication, and division
	Add      http.Handler
	Subtract http.Handler
	Multiply http.Handler
	Divide   http.Handler
	// Keeps track of answered problems
	Problems http.Handler
	Problem  http.Handler

	httpServer *http.Server
}

// Start starts the http.Server
func (server *Server) Start() {

	router := server.Router()

	server.ServerMutex.Lock()
	server.httpServer = &http.Server{
		Addr:     server.HTTPListenAddr,
		Handler:  router,
		ErrorLog: server.HTTPLogger,
	}
	server.ServerMutex.Unlock()

	log.Println("I'm starting the server")
	if err := server.httpServer.ListenAndServe(); err != nil {
		server.HTTPLogger.Println("Unable to listen and serve", err.Error())
	}
}

// Stop tells the httpServer to shutdown
func (server *Server) Stop(ctx context.Context) {
	server.ServerMutex.Lock()
	defer server.ServerMutex.Unlock()

	err := server.httpServer.Shutdown(ctx)

	if err != nil {
		server.HTTPLogger.Print("unable to shutdown. error:", err.Error())
	}
	log.Println("I'm stopping the server")
	server.WaitGroup.Done()
}

// Router parses the request URI and supplies the needed handler
func (server *Server) Router() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.NewHandler)

	r.Handle("/", http.FileServer(http.Dir("./cmd/math/static")))
	r.Handle("/add", server.Add).Methods("GET")
	r.Handle("/subtract", server.Subtract).Methods("GET")
	r.Handle("/multiply", server.Multiply).Methods("GET")
	r.Handle("/divide", server.Divide).Methods("GET")
	r.Handle("/problems", server.Problems).Methods("GET")
	r.Handle("/problems/{problem_id}", server.Problem).Methods("GET")

	return r
}
