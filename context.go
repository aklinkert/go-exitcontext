package exitcontext

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// New uses a default background context to create the exit handler
func New() context.Context {
	return NewWithContext(context.Background())

}

// NewWithContext returns a context instance that gets cancelled if the process receives a SIGTERM or a SIGINT.
func NewWithContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
