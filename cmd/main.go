package main

import (
	"context"
	"diagram-server/internal/bootstrap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := bootstrap.NewApplication()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		log.Fatalf("[Error] Application error: %v", err)
	}
}
