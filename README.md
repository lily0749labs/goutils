# goutils

Go 语言常用工具库。

## 安装

```bash
go get github.com/lily0749labs/goutils@latest
```

## 工具包

- `anyto`：任意类型转换
- `strto`：字符串类型转换
- `valid`：常用数据格式校验
- `crypto`：哈希与加解密
- `rand`：随机数与随机字符串
- `id`：UUID 与雪花 ID
- `time`：日期与时间处理

## 使用

```go
package main

import (
	"fmt"

	"github.com/lily0749labs/goutils"
)

func main() {
	fmt.Println(goutils.Time.NowTime())
	fmt.Println(goutils.AnyTo.Int("123"))
	fmt.Println(goutils.StrTo.Int("123"))
	fmt.Println(goutils.Valid.Email("test@example.com"))
	fmt.Println(goutils.Crypto.MD5("hello"))
	fmt.Println(goutils.Rand.Rand6())
	fmt.Println(goutils.ID.Token())
}
```

根包提供以下门面入口：

- `goutils.AnyTo`
- `goutils.StrTo`
- `goutils.Valid`
- `goutils.Crypto`
- `goutils.Rand`
- `goutils.ID`
- `goutils.Time`

各子包的结构体方法承载实际实现；原有包级函数作为兼容入口保留。例如，
`anyto.AnyToInt("123")` 仍然可用，但新代码推荐使用 `goutils.AnyTo.Int("123")`。

## 发布前检查

```bash
go build ./...
go vet ./...
```
