# strto 字符串解析与转换

`strto` 用于把字符串解析为常见 Go 基础类型，适合配置项、命令行参数、表单字段和文本数据。
推荐通过根包门面 `goutils.StrTo` 调用；只有在处理 `ElementError` 等包内类型时，才需要额外导入
`strto` 子包。

## 快速选型

| 需求 | 选择 | 失败行为 |
| --- | --- | --- |
| 必须识别非法输入 | `ParseXxx` | 返回目标值和 `error` |
| 允许使用零值回退 | `XxxOrZero` / `BoolOrFalse` | 数字、时长返回 `0`，布尔值返回 `false` |
| 解析 `[]string` | 对应的复数方法，如 `ParseInts` | 任一元素失败时返回 `ElementError` |
| 逐项转换 `[]string`，非法元素写零值 | `IntsOrZero`、`BoolsOrFalse` 等 | 保留输入顺序和长度 |
| 解析逗号分隔的整数 | `ParseCommaInts` | 字段非法时返回 `ElementError` |
| 从逗号列表中提取合法整数 | `ParseCommaIntsLoose` | 忽略空字段和非法字段 |
| 获取字符串的字节副本 | `Bytes` | 不会失败 |

业务输入通常应选择 `ParseXxx`。只有调用方明确接受零值，而且无需区分合法零值与非法输入时，
才适合使用 `OrZero` / `OrFalse` 方法。

## 调用入口

```go
import "github.com/lily-study-utils/goutils"

port, err := goutils.StrTo.ParseInt("8080")
timeout := goutils.StrTo.DurationOrZero("30s")
```

需要检查切片元素错误时导入子包：

```go
import strutil "github.com/lily-study-utils/goutils/strto"
```

## 功能列表

### 1. 有符号整数

所有整数方法都按十进制处理。

| 目标类型 | 返回错误的解析方法 | 零值回退方法 |
| --- | --- | --- |
| `int` | `ParseInt` | `IntOrZero` |
| `int8` | `ParseInt8` | `Int8OrZero` |
| `int16` | `ParseInt16` | `Int16OrZero` |
| `int32` | `ParseInt32` | `Int32OrZero` |
| `int64` | `ParseInt64` | `Int64OrZero` |

### 2. 无符号整数

| 目标类型 | 返回错误的解析方法 | 零值回退方法 |
| --- | --- | --- |
| `uint` | `ParseUint` | `UintOrZero` |
| `uint8` | `ParseUint8` | `Uint8OrZero` |
| `uint16` | `ParseUint16` | `Uint16OrZero` |
| `uint32` | `ParseUint32` | `Uint32OrZero` |
| `uint64` | `ParseUint64` | `Uint64OrZero` |

负数不能解析为无符号整数。`int` 和 `uint` 的范围取决于当前运行平台的字长。

### 3. 布尔值、有限浮点数、时长和字节

| 目标类型 | 返回错误的解析方法 | 回退或直接转换方法 | 说明 |
| --- | --- | --- | --- |
| `bool` | `ParseBool` | `BoolOrFalse` | 接受不区分大小写的 `true`、`false`，以及数值 `0`、`1` |
| `float32` | `ParseFiniteFloat32` | `Float32OrZero` | 拒绝 NaN、无穷大和超出范围的值 |
| `float64` | `ParseFiniteFloat64` | `Float64OrZero` | 拒绝 NaN、无穷大和超出范围的值 |
| `time.Duration` | `ParseDuration` | `DurationOrZero` | 使用 Go `time.Duration` 语法 |
| `[]byte` | — | `Bytes` | 返回字符串内容的可写字节副本 |

### 4. 字符串切片

切片方法与标量方法使用相同的解析规则。

| 分类 | 返回错误的解析方法 | 逐项零值回退方法 |
| --- | --- | --- |
| 布尔值 | `ParseBools` | `BoolsOrFalse` |
| 有符号整数 | `ParseInts`、`ParseInt8s`、`ParseInt16s`、`ParseInt32s`、`ParseInt64s` | `IntsOrZero`、`Int8sOrZero`、`Int16sOrZero`、`Int32sOrZero`、`Int64sOrZero` |
| 无符号整数 | `ParseUints`、`ParseUint8s`、`ParseUint16s`、`ParseUint32s`、`ParseUint64s` | `UintsOrZero`、`Uint8sOrZero`、`Uint16sOrZero`、`Uint32sOrZero`、`Uint64sOrZero` |
| 有限浮点数 | `ParseFiniteFloat32s`、`ParseFiniteFloat64s` | `Float32sOrZero`、`Float64sOrZero` |
| Go 时长 | `ParseDurations` | `DurationsOrZero` |

### 5. 逗号整数列表

| 方法 | 行为 |
| --- | --- |
| `ParseCommaInts` | 去除每个字段两侧的空白；字段为空、语法错误或溢出时返回 `ElementError` |
| `ParseCommaIntsLoose` | 去除字段两侧空白，并忽略空字段、非法字段和溢出字段 |

`ParseCommaInts` 允许列表首尾各出现一个空字段，便于处理 `",1,2,"` 这类输入；中间空字段仍然报错。
空字符串在两种模式下都返回 `nil`；Loose 模式没有提取到有效整数时也返回 `nil`。
返回错误的解析模式不返回部分结果。

## 行为约定

### Parse 与零值回退

`ParseXxx` 方法保留解析失败信息，适合需要校验输入的场景：

```go
port, err := goutils.StrTo.ParseInt("8080")
if err != nil {
    return fmt.Errorf("端口格式错误: %w", err)
}
```

`OrZero` / `OrFalse` 方法忽略解析错误：

```go
port := goutils.StrTo.IntOrZero("无效")          // 0
enabled := goutils.StrTo.BoolOrFalse("无效")    // false
timeout := goutils.StrTo.DurationOrZero("无效") // 0
```

零值回退方法不能区分 `"0"` 与非法输入。若这个区别会影响业务逻辑，应使用 `ParseXxx`。

### 输入格式

- 整数只按十进制解析，不接受小数，也不会自动去除首尾空白。
- `ParseBool` 不区分 `true`、`false` 的大小写；为兼容既有行为，`00`、`01` 等带前导零的数值形式也有效。
- 有限浮点数方法保留 `strconv.ParseFloat` 的语法，但拒绝 NaN、正负无穷大以及目标类型范围外的值。
- `ParseDuration` 使用 Go `time.Duration` 语法，例如 `250ms`、`1h30m`；裸数字不表示秒。
- 只有逗号列表方法会主动去除每个字段两侧的空白。

### 切片结果

返回错误的切片解析采用“全部成功或全部失败”语义：任一元素失败时返回 `nil` 和
`*strto.ElementError`，不返回部分结果。零值回退方法则保留输入顺序和长度，失败元素写入零值。

| 输入 | `ParseXxxs` | `XxxsOrZero` / `BoolsOrFalse` |
| --- | --- | --- |
| `nil` | `(nil, nil)` | `nil` |
| 非 nil 空切片 | 非 nil 空切片和 `nil` 错误 | `nil` |
| 包含非法元素 | `nil` 和 `ElementError` | 等长切片，非法位置为零值 |

## 错误分类

解析错误支持 `errors.Is`；切片和列表的元素错误支持 `errors.As`。

| 错误 | 含义 |
| --- | --- |
| `ErrSyntax` | 输入不符合目标类型语法；对应 `strconv.ErrSyntax` |
| `ErrOutOfRange` | 数值超出目标类型范围；对应 `strconv.ErrRange` |
| `ErrInvalidBool` | 输入不是受支持的布尔值，同时属于 `ErrSyntax` |
| `ErrNonFiniteFloat` | 输入为 NaN 或无穷大，同时属于 `ErrOutOfRange` |
| `ErrInvalidDuration` | 输入不是有效的 Go 时长，同时属于 `ErrSyntax` |
| `*ElementError` | 切片或列表中某个元素解析失败，记录从 0 开始的 `Index`、`Value` 和底层 `Err` |

整数和普通浮点解析保留标准库的 `*strconv.NumError`，可以用 `errors.As` 读取原始输入与解析函数。

```go
_, err := goutils.StrTo.ParseInts([]string{"1", "错误", "3"})
if err != nil {
    var elementErr *strutil.ElementError
    if errors.As(err, &elementErr) {
        log.Printf("第 %d 个元素 %q 解析失败: %v",
            elementErr.Index, elementErr.Value, elementErr.Err)
    }
}
```

## 常用场景示例

### 解析配置

```go
port, err := goutils.StrTo.ParseUint16(config.Port)
if err != nil {
    return fmt.Errorf("无效端口 %q: %w", config.Port, err)
}

timeout, err := goutils.StrTo.ParseDuration(config.Timeout)
if err != nil {
    return fmt.Errorf("无效超时时间 %q: %w", config.Timeout, err)
}
```

### 使用明确的默认值

需要业务默认值时，建议先解析，再由调用方设置默认值，避免把库的零值误当成业务默认值：

```go
retry, err := goutils.StrTo.ParseInt(config.Retry)
if err != nil {
    retry = 3
}
```

## 兼容名称迁移

旧名称集中在 `compat.go` 和 `compat_methods.go`，均已标记为 `Deprecated`。已有代码仍可编译，
新代码应使用下表中的规范名称。

| 旧名称 | 推荐名称 |
| --- | --- |
| `StrictInt` 至 `StrictInt64`、`IntE` 至 `Int64E` | `ParseInt` 至 `ParseInt64` |
| `StrictUint` 至 `StrictUint64`、`UintE` 至 `Uint64E` | `ParseUint` 至 `ParseUint64` |
| `StrictBool`、`BoolE` | `ParseBool` |
| `StrictFloat32`、`Float32E`、`StrictFloat64`、`Float64E` | `ParseFiniteFloat32`、`ParseFiniteFloat64` |
| `StrictDuration`、`DurationE` | `ParseDuration` |
| `StrictInts`、`IntsE` 等旧切片解析名称 | 对应的 `ParseXxxs` |
| `Int`、`Uint64`、`Float32`、`Duration` 等旧短名称 | 对应的 `OrZero` 名称 |
| `Bool`、`Bools` | `BoolOrFalse`、`BoolsOrFalse` |
| `CommaInts`、`StrictCommaInts` | `ParseCommaIntsLoose`、`ParseCommaInts` |
| `StrToInt`、`StrToBool`、`StrToBytes` 等包级函数 | `StrTo.IntOrZero`、`StrTo.BoolOrFalse`、`StrTo.Bytes` 等方法 |
| `ArrStrToInt`、`ArrStrToUint` 等包级函数 | `StrTo.IntsOrZero`、`StrTo.UintsOrZero` 等方法 |
