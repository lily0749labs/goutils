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

	"github.com/lily0749labs/goutils/anyto"
)

func main() {
	value, err := anyto.AnyToInt("123")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
```

## 发布前检查

```bash
go build ./...
go vet ./...
```
