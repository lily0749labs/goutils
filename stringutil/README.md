# stringutil

提供 Unicode 安全的字符串门面，推荐通过根包调用：

```go
short := goutils.String.TruncateWithSuffix("很长的文本", 5, "…")
masked := goutils.String.Mask("13800138000", 3, 4)
snake := goutils.String.Snake("HTTPServerPort")
camel := goutils.String.Camel("user_name")
clean := goutils.String.NormalizeSpace("  hello\tworld  ")
```

`Truncate`、`Mask` 和 `Reverse` 均按 Unicode 字符处理，不会把中文或 Emoji 拆成无效字节。
