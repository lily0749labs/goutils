# money

使用整数最小货币单位格式化金额，避免浮点精度问题：

```go
positive := money.FormatCents(12345) // "123.45"
negative := money.FormatCents(-1)    // "-0.01"
```

也可通过根包门面调用：

```go
value := goutils.Money.FormatCents(12345)
```

`FormatCents` 只进行确定性的数值格式化，不添加货币符号、千位分隔符或本地化规则。
