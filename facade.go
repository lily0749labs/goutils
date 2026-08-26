// Package goutils 提供各工具子包的统一门面入口。
package goutils

import (
	anyutil "github.com/lily0749labs/goutils/anyto"
	cryptoutil "github.com/lily0749labs/goutils/crypto"
	floatutil "github.com/lily0749labs/goutils/floatutil"
	idutil "github.com/lily0749labs/goutils/id"
	jsonutil "github.com/lily0749labs/goutils/jsonutil"
	moneyutil "github.com/lily0749labs/goutils/money"
	randutil "github.com/lily0749labs/goutils/rand"
	stringutil "github.com/lily0749labs/goutils/stringutil"
	strutil "github.com/lily0749labs/goutils/strto"
	syncutil "github.com/lily0749labs/goutils/syncutil"
	timeutil "github.com/lily0749labs/goutils/time"
	validutil "github.com/lily0749labs/goutils/valid"
)

var (
	// AnyTo 提供任意值转换工具。
	AnyTo = anyutil.AnyTo
	// Crypto 提供哈希、加密和解密工具。
	Crypto = cryptoutil.Crypto
	// Float 提供语义明确的浮点数辅助函数。
	Float = floatutil.Float
	// ID 提供 UUID 和雪花 ID 生成工具。
	ID = idutil.ID
	// JSON 提供 JSON 编解码辅助函数。
	JSON = jsonutil.JSON
	// Money 提供基于整数最小货币单位的金额工具。
	Money = moneyutil.Money
	// Rand 提供随机数、随机字符串和随机姓名工具。
	Rand = randutil.Rand
	// StrTo 提供字符串类型转换工具。
	StrTo = strutil.StrTo
	// String 提供 Unicode 安全的字符串处理工具。
	String = stringutil.String
	// Sync 提供标准同步容器辅助函数。
	Sync = syncutil.Sync
	// Time 提供日期和时间处理工具。
	Time = timeutil.Time
	// Valid 提供常见数据格式校验工具。
	Valid = validutil.Valid
)
