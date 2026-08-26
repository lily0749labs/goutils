package queue

import (
	"context"
	"errors"
	"sync"
)

var (
	// ErrClosed 表示队列已关闭且无法完成当前操作。
	ErrClosed = errors.New("queue: closed")
	// ErrNilContext 表示 Push 或 Pop 收到了 nil context。
	ErrNilContext = errors.New("queue: nil context")
)

// Queue 是先进先出的类型安全队列。
// Queue 必须通过 New 创建；Close 可以安全地重复调用。
type Queue[T any] struct {
	data      chan T
	closed    chan struct{}
	closeOnce sync.Once
}

// New 创建容量为 capacity 的队列。
// capacity 为 0 时，Push 和 Pop 必须直接交付；capacity 为负数时会 panic。
func New[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		data:   make(chan T, capacity),
		closed: make(chan struct{}),
	}
}

// Push 在队列可写、ctx 取消或队列关闭前等待。
func (q *Queue[T]) Push(ctx context.Context, value T) error {
	if ctx == nil {
		return ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	select {
	case <-q.closed:
		return ErrClosed
	default:
	}

	select {
	case <-q.closed:
		return ErrClosed
	case <-ctx.Done():
		return ctx.Err()
	case q.data <- value:
		return nil
	}
}

// Pop 在有元素可读、ctx 取消或队列关闭且已空前等待。
// Close 之前已入队的元素仍可以按先进先出顺序读取。
func (q *Queue[T]) Pop(ctx context.Context) (T, error) {
	var zero T
	if ctx == nil {
		return zero, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return zero, err
	}

	select {
	case value := <-q.data:
		return value, nil
	default:
	}

	select {
	case value := <-q.data:
		return value, nil
	case <-ctx.Done():
		return zero, ctx.Err()
	case <-q.closed:
		select {
		case value := <-q.data:
			return value, nil
		default:
			return zero, ErrClosed
		}
	}
}

// Close 关闭队列并唤醒正在等待的 Push 和 Pop。
func (q *Queue[T]) Close() {
	q.closeOnce.Do(func() {
		close(q.closed)
	})
}

// Closed 报告队列是否已关闭。
func (q *Queue[T]) Closed() bool {
	select {
	case <-q.closed:
		return true
	default:
		return false
	}
}

// Len 返回当前已缓冲的元素数。
func (q *Queue[T]) Len() int { return len(q.data) }

// Cap 返回队列容量。
func (q *Queue[T]) Cap() int { return cap(q.data) }
