package main

import (
	"context"
	"os/signal"
	"store/integration"
	server "store/rest"
	"syscall"
)

func main() {

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer cancel()

	go func() {
		server.Start()
	}()

	go func() {
		integration.Start(ctx)
	}()

	<-ctx.Done()

	server.Stop()
	integration.Stop()
}
