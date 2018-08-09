package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ZachGill/math/cmd/math/handlers"
)

func main() {
	var (
		problemsAccessor = &handlers.Accessor{
			Problems: map[string]handlers.MathResponse{},
			Mutex:    &sync.RWMutex{},
		}

		waitGroup = &sync.WaitGroup{}

		httpServer = &handlers.Server{
			ServerMutex:    &sync.Mutex{},
			WaitGroup:      waitGroup,
			HTTPListenAddr: ":8080",
			HTTPLogger:     log.New(os.Stderr, "HTTPLogger: ", log.Lshortfile),
			Add: &handlers.Add{
				ProblemsAccessor: problemsAccessor,
			},
			Subtract: &handlers.Subtract{
				ProblemsAccessor: problemsAccessor,
			},
			Multiply: &handlers.Multiply{
				ProblemsAccessor: problemsAccessor,
			},
			Divide: &handlers.Divide{
				ProblemsAccessor: problemsAccessor,
			},
			Problems: &handlers.Problems{
				ProblemsAccessor: problemsAccessor,
			},
			Problem: &handlers.Problem{
				ProblemsAccessor: problemsAccessor,
			},
		}
	)

	waitGroup.Add(1)
	go httpServer.Start()
	go waitForSignal(make(chan os.Signal, 1), httpServer)
	waitGroup.Wait()
}

func waitForSignal(c <-chan os.Signal, server *handlers.Server) {
	<-c
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	server.Stop(ctx)
	cancelFunc()
}
