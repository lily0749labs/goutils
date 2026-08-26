# queue

提供类型安全、先进先出的并发队列：

```go
q := queue.New[string](16)
defer q.Close()

if err := q.Push(ctx, "job"); err != nil {
    return err
}
value, err := q.Pop(ctx)
```

`Push` 和 `Pop` 使用 `context.Context` 控制取消和超时。`Close` 可以安全地重复调用，
会唤醒所有等待中的操作。关闭前已缓冲的元素仍可读取；队列关闭且已空时，
`Pop` 返回 `ErrClosed`。

容量为 0 的队列不缓冲元素，每次 `Push` 都需要与 `Pop` 直接交付。
