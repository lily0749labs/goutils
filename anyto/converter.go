package anyto

// AnyTo 是任意值转换方法的统一入口。
// 该值不保存状态，可以安全地在多个 goroutine 中复用。
var AnyTo = anyto{}

type anyto struct{}
