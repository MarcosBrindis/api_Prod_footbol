package ports

import "context"

type EventPublisher interface {
	PublishEvent(ctx context.Context, message string) error
}
