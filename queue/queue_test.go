package queue

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestQueueFIFO(t *testing.T) {
	t.Parallel()

	q := New[int](2)
	ctx := context.Background()
	if err := q.Push(ctx, 1); err != nil {
		t.Fatal(err)
	}
	if err := q.Push(ctx, 2); err != nil {
		t.Fatal(err)
	}
	if q.Len() != 2 || q.Cap() != 2 {
		t.Fatalf("Len/Cap = %d/%d, want 2/2", q.Len(), q.Cap())
	}

	for _, want := range []int{1, 2} {
		got, err := q.Pop(ctx)
		if err != nil || got != want {
			t.Fatalf("Pop() = (%d, %v), want (%d, nil)", got, err, want)
		}
	}
}

func TestQueueContextCancellation(t *testing.T) {
	t.Parallel()

	q := New[int](1)
	if err := q.Push(context.Background(), 1); err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := q.Push(ctx, 2); !errors.Is(err, context.Canceled) {
		t.Fatalf("Push(canceled) error = %v, want context.Canceled", err)
	}

	empty := New[int](1)
	if _, err := empty.Pop(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("Pop(canceled) error = %v, want context.Canceled", err)
	}
	if err := empty.Push(nil, 1); !errors.Is(err, ErrNilContext) {
		t.Fatalf("Push(nil) error = %v, want ErrNilContext", err)
	}
	if _, err := empty.Pop(nil); !errors.Is(err, ErrNilContext) {
		t.Fatalf("Pop(nil) error = %v, want ErrNilContext", err)
	}
}

func TestQueueCloseDrainsBufferedValues(t *testing.T) {
	t.Parallel()

	q := New[string](2)
	ctx := context.Background()
	if err := q.Push(ctx, "first"); err != nil {
		t.Fatal(err)
	}
	if err := q.Push(ctx, "second"); err != nil {
		t.Fatal(err)
	}
	q.Close()
	q.Close()

	if !q.Closed() {
		t.Fatal("Closed() = false, want true")
	}
	if err := q.Push(ctx, "third"); !errors.Is(err, ErrClosed) {
		t.Fatalf("Push() after Close error = %v, want ErrClosed", err)
	}
	for _, want := range []string{"first", "second"} {
		got, err := q.Pop(ctx)
		if err != nil || got != want {
			t.Fatalf("Pop() after Close = (%q, %v), want (%q, nil)", got, err, want)
		}
	}
	if _, err := q.Pop(ctx); !errors.Is(err, ErrClosed) {
		t.Fatalf("empty Pop() after Close error = %v, want ErrClosed", err)
	}
}

func TestQueueCloseUnblocksWaiters(t *testing.T) {
	t.Parallel()

	full := New[int](1)
	if err := full.Push(context.Background(), 1); err != nil {
		t.Fatal(err)
	}
	pushResult := make(chan error, 1)
	go func() {
		pushResult <- full.Push(context.Background(), 2)
	}()
	full.Close()
	if err := <-pushResult; !errors.Is(err, ErrClosed) {
		t.Fatalf("blocked Push() error = %v, want ErrClosed", err)
	}

	empty := New[int](1)
	popResult := make(chan error, 1)
	go func() {
		_, err := empty.Pop(context.Background())
		popResult <- err
	}()
	empty.Close()
	if err := <-popResult; !errors.Is(err, ErrClosed) {
		t.Fatalf("blocked Pop() error = %v, want ErrClosed", err)
	}
}

func TestUnbufferedQueue(t *testing.T) {
	t.Parallel()

	q := New[int](0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		errCh <- q.Push(ctx, 42)
	}()
	got, err := q.Pop(ctx)
	if err != nil || got != 42 {
		t.Fatalf("Pop() = (%d, %v), want (42, nil)", got, err)
	}
	if err := <-errCh; err != nil {
		t.Fatalf("Push() error = %v", err)
	}
}

func TestConcurrentPushAndClose(t *testing.T) {
	t.Parallel()

	for range 100 {
		q := New[int](1)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		var wait sync.WaitGroup
		wait.Add(2)
		go func() {
			defer wait.Done()
			err := q.Push(ctx, 1)
			if err != nil && !errors.Is(err, ErrClosed) {
				t.Errorf("Push() error = %v", err)
			}
		}()
		go func() {
			defer wait.Done()
			q.Close()
		}()
		wait.Wait()
		cancel()
	}
}
