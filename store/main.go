package main

import (
	"context"
	"os/signal"
	"store/app/messaging"
	"store/container"
	"store/integration"
	"store/kafka"
	repository "store/repository"
	server "store/rest"
	"store/utils"
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
		integration.Start(ctx)
	}()

	<-ctx.Done()

	server.Stop()
	integration.Stop()

	eventDispatcher := container.Instance().Get(utils.Nameof((*messaging.EventDispatcher)(nil))).(messaging.EventDispatcher)
	eventDispatcher.Dispose()
}
