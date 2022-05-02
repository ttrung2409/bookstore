package kafka

import "context"

type Consumer interface {
	FetchMessage(ctx context.Context) (Message, error)
	CommitMessage(ctx context.Context, msg Message) error
	Dispose() error
}
