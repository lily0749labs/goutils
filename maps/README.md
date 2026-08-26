# maps

提供类型安全的泛型 Map 工具：

```go
keys := maps.Keys(values)
items := maps.Values(values)
copy := maps.Clone(values)
merged := maps.Merge(defaults, overrides)
filtered := maps.Filter(values, func(key string, value int) bool { return value > 0 })
texts := maps.MapValues(values, func(_ string, value int) string { return strconv.Itoa(value) })
complete := maps.ContainsKeys(values, "name", "age")
missing := maps.MissingKeys(values, "name", "age", "email")
```

`Clone`、`Merge`、`Filter`、`MapValues`、`Pick` 和 `Omit` 都返回新 Map，
不会修改输入。`Merge` 中后传入的 Map 会覆盖前面相同键的值。

`ContainsKeys` 和 `MissingKeys` 只检查键是否存在，不判断对应值是否为零值或 nil。
`MissingKeys` 按请求顺序返回缺失键，重复键只会出现一次。
