//go:build unix

package exitcontext

import (
	"context"
	"errors"
	"syscall"
	"testing"
)

func TestCancelledBySignal(t *testing.T) {
	tests := []struct {
		name   string
		signal syscall.Signal
		new    func() context.Context
	}{
		{name: "New on SIGINT", signal: syscall.SIGINT, new: New},
		{name: "New on SIGTERM", signal: syscall.SIGTERM, new: New},
		{
			name:   "NewWithContext on SIGTERM",
			signal: syscall.SIGTERM,
			new:    func() context.Context { return NewWithContext(context.Background()) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.new()

			if err := ctx.Err(); err != nil {
				t.Fatalf("expected a live context, got %v", err)
			}

			if err := syscall.Kill(syscall.Getpid(), tt.signal); err != nil {
				t.Fatalf("sending %v: %v", tt.signal, err)
			}

			waitDone(t, ctx)

			if err := ctx.Err(); !errors.Is(err, context.Canceled) {
				t.Errorf("Err() = %v, want %v", err, context.Canceled)
			}
		})
	}
}
