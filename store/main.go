package main

import (
	"context"
	"os/signal"
	"store/container"
	"store/kafka"
	"store/messaging"
	repository "store/repository"
	server "store/rest"
	"syscall"
)

func main() {
	builder := container.ContainerBuilder()

	repository.Install(builder)
	kafka.Install(builder)

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
		messaging.Start(ctx)
	}()

	<-ctx.Done()

	server.Stop()
	messaging.Stop()
}
