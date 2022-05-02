package kafka

import "context"

type Producer interface {
	Send(ctx context.Context, msgs ...Message) error
	Dispose() error
}
