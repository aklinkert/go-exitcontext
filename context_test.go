package exitcontext

import (
	"context"
	"errors"
	"runtime"
	"testing"
	"time"
)

func waitDone(t *testing.T, ctx context.Context) {
	t.Helper()

	select {
	case <-ctx.Done():
	case <-time.After(5 * time.Second):
		t.Fatal("context was not cancelled")
	}
}

func TestNewWithContextFollowsParent(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	ctx := NewWithContext(parent)

	if err := ctx.Err(); err != nil {
		t.Fatalf("expected a live context, got %v", err)
	}

	cancel()
	waitDone(t, ctx)

	if err := ctx.Err(); !errors.Is(err, context.Canceled) {
		t.Errorf("Err() = %v, want %v", err, context.Canceled)
	}
}

func TestNewWithContextKeepsValues(t *testing.T) {
	type key struct{}

	ctx := NewWithContext(context.WithValue(context.Background(), key{}, "value"))

	if got := ctx.Value(key{}); got != "value" {
		t.Errorf("Value() = %v, want %q", got, "value")
	}
}

func TestNewWithContextReleasesGoroutine(t *testing.T) {
	newCancelled := func() {
		parent, cancel := context.WithCancel(context.Background())
		ctx := NewWithContext(parent)
		cancel()
		waitDone(t, ctx)
	}

	// The first signal.Notify starts a long-lived runtime goroutine, keep it
	// out of the baseline.
	newCancelled()
	time.Sleep(50 * time.Millisecond)

	before := runtime.NumGoroutine()
	newCancelled()

	deadline := time.Now().Add(5 * time.Second)
	for runtime.NumGoroutine() > before {
		if time.Now().After(deadline) {
			t.Fatalf("goroutines = %d, want at most %d", runtime.NumGoroutine(), before)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
