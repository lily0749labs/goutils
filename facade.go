// Package goutils 提供各工具子包的统一门面入口。
package goutils

import (
	anyutil "github.com/lily0749labs/goutils/anyto"
	cryptoutil "github.com/lily0749labs/goutils/crypto"
	idutil "github.com/lily0749labs/goutils/id"
	randutil "github.com/lily0749labs/goutils/rand"
	stringutil "github.com/lily0749labs/goutils/stringutil"
	strutil "github.com/lily0749labs/goutils/strto"
	timeutil "github.com/lily0749labs/goutils/time"
	validutil "github.com/lily0749labs/goutils/valid"
)

var (
	// AnyTo 提供任意值转换工具。
	AnyTo = anyutil.AnyTo
	// Crypto 提供哈希、加密和解密工具。
	Crypto = cryptoutil.Crypto
	// ID 提供 UUID 和雪花 ID 生成工具。
	ID = idutil.ID
	// Rand 提供随机数、随机字符串和随机姓名工具。
	Rand = randutil.Rand
	// StrTo 提供字符串类型转换工具。
	StrTo = strutil.StrTo
	// String 提供 Unicode 安全的字符串处理工具。
	String = stringutil.String
	// Time 提供日期和时间处理工具。
	Time = timeutil.Time
	// Valid 提供常见数据格式校验工具。
	Valid = validutil.Valid
)
