package main

import (
	"context"
	"os/signal"
	"store/container"
	"store/integration"
	repository "store/repository"
	server "store/rest"
	"syscall"
)

func main() {
	builder := container.ContainerBuilder()

	repository.RegisterDependencies(builder)

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

	repository.GetEventDispatcher().Dispose()
}
