package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ajay-308/student-api/internal/config"
)

func main() {

	fmt.Println("welcome to student api")
	cfg := config.MustLoad()


	router := http.NewServeMux()
	router.Handle("/students", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I AM rich"))
	}))

	server := http.Server{
		Addr:    cfg.HTTPServer.Address, 
		Handler: router,
	}
	slog.Info("starting servr",slog.String("adress",cfg.HTTPServer.Address))
	fmt.Printf("Starting server on %s...\n", cfg.HTTPServer.Address)
	//var wg sync.WaitGroup
	//wg.Add(1)

	done := make(chan os.Signal, 1) // create a channel to listen for os signals
	signal.Notify(done, os.Interrupt, syscall.SIGINT , syscall.SIGALRM)// listen for the interrupt signal
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("failed to start server:", err)
		}
	}()

	<-done // block until we receive a signal

	slog.Info("shutting down server...")
	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil{
		log.Fatal("failed to shutdown server:",err)
	}

	slog.Info("server shutdown successfully")

	// Print that the server has started
	fmt.Println("server started")

	// Wait for the server goroutine to complete
	//wg.Wait()
}
