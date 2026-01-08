package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dealance.co/backend/internal/config"
	"dealance.co/backend/internal/httpx"
)

func main() {
	cfg := config.Load()

	server := httpx.NewServer(httpx.Config{
		Addr: cfg.HTTP.Addr, // explicit mapping
	})

	go func() {
		log.Printf("auth-service listening on %s\n", cfg.HTTP.Addr)
		if err := server.Start(); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("shutting down auth-service...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}
}
