# goutils

Go 常用工具库。

## 安装

```bash
go get github.com/lily0749labs/goutils@latest
```

## 模块

- `AnyTo`：任意值转换
- `StrTo`：字符串转换
- `Valid`：常见格式校验
- `Crypto`：哈希、AES、RSA 和 bcrypt
- `Rand`：随机数、随机字符串和英文名
- `ID`：UUID 和雪花 ID
- `Money`：整数最小货币单位格式化
- `Time`：日期、时间戳和时间范围
- `String`：Unicode 字符串、脱敏和命名转换
- `slice`：类型安全的泛型切片函数
- `maps`：类型安全的泛型 Map 函数
- `queue`：支持 Context 取消和并发关闭的泛型队列

## 使用门面

```go
package main

import (
	"fmt"

	"github.com/lily0749labs/goutils"
)

func main() {
	value, err := goutils.AnyTo.Int("123")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
	number, err := goutils.StrTo.IntE("456")
	if err != nil {
		panic(err)
	}

	ciphertext, err := goutils.Crypto.EncryptGCM("1234567890abcdef", "hello")
	if err != nil {
		panic(err)
	}
	plaintext, err := goutils.Crypto.DecryptGCM("1234567890abcdef", ciphertext)
	if err != nil {
		panic(err)
	}

	fmt.Println(number)
	fmt.Println(goutils.Valid.Email("user@example.com"))
	fmt.Println(goutils.Crypto.MD5("hello"))
	fmt.Println(string(plaintext))
	secureText, err := goutils.Rand.SecureString(16)
	if err != nil {
		panic(err)
	}
	fmt.Println(secureText)
	fmt.Println(goutils.ID.Token())
	fmt.Println(goutils.Money.FormatCents(12345))
	fmt.Println(goutils.Time.NowTime())
	fmt.Println(goutils.String.Mask("13800138000", 3, 4))
}
```

结构体方法承载实际实现。原有包级函数仍保留为兼容入口，例如
`anyto.AnyToInt("123")`，但新代码推荐使用 `goutils.AnyTo.Int("123")`。

字符串转整数或布尔值时，如果需要识别非法输入，推荐使用带 `E` 后缀的严格方法，
例如 `StrTo.IntE`、`StrTo.Uint64E` 和 `StrTo.BoolE`。不带 `E` 的兼容方法在失败时返回零值。

新代码中的对称加密推荐使用 `Crypto.EncryptGCM` 和 `Crypto.DecryptGCM`。
GCM 会为每次加密生成随机 Nonce，并校验密文完整性。

RSA 加密推荐使用 `Crypto.EncryptRSAOAEP` / `DecryptRSAOAEP`，签名验签使用
`Crypto.SignRSAPSS` / `VerifyRSAPSS`。安全令牌或密码材料推荐使用
`Rand.SecureString`、`Rand.SecureInt`，这些方法会返回随机源错误。

严格切片转换使用 `StrTo.IntsE`、`StrTo.UintsE`、`StrTo.BoolsE` 等方法，
失败时 `ElementError` 会提供元素下标和原始值。时间解析推荐使用 `Time.ParseE`
或 `Time.ParseLayoutE`；测试中可通过 `Time.WithNow` 注入固定时钟。

## 开发与验证

```bash
go test ./...
go test -race ./...
go vet ./...
```
