# slice

提供类型安全的泛型切片函数。Go 不支持在普通门面方法上声明新的类型参数，因此本包使用包级泛型函数：

```go
numbers := []int{1, 2, 2, 3}
unique := slice.Unique(numbers)
even := slice.Filter(numbers, func(value int) bool { return value%2 == 0 })
texts := slice.Map(numbers, strconv.Itoa)
total := slice.Reduce(numbers, 0, func(total, value int) int { return total + value })
remaining := slice.DeleteAt(numbers, 1)

result := slice.New(numbers).
    Filter(func(value int) bool { return value > 1 }).
    Unique().
    Values()
```

所有返回切片的操作都不会修改输入。`DeleteAt` 在下标越界时返回输入副本。
`Intersection` 和 `Difference` 返回去重结果，
并保留第一个切片中元素的首次出现顺序。
