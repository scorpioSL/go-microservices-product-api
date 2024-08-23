package main

import (
	"PRODUCT-API/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "products-api ", log.LstdFlags)

	// create the handler
	ph := handlers.NewProducts(logger)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting the server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	logger.Println("Got signal:", sig)

	// gracefully shutdown
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
