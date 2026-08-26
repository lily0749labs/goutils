package rand

import (
	cryptorand "crypto/rand"
	"errors"
	"io"
	"math/big"
)

var (
	// ErrInvalidLength 表示随机字符串长度为负数。
	ErrInvalidLength = errors.New("random string length must not be negative")
	// ErrInvalidCharset 表示字符集不是由 2 至 128 个不重复 ASCII 字符组成。
	ErrInvalidCharset = errors.New("random charset must contain between 2 and 128 unique ASCII characters")
	// ErrInvalidRange 表示随机整数区间为空或上下界颠倒。
	ErrInvalidRange = errors.New("random range must satisfy min < max")
	// ErrNilReader 表示调用方传入了空随机源。
	ErrNilReader = errors.New("random reader must not be nil")
)

// SecureString 使用密码学安全随机源生成由默认字符集组成的字符串。
func (r rand) SecureString(length int) (string, error) {
	return r.SecureStringWithCharset(length, string(Chars))
}

// SecureStringWithCharset 使用密码学安全随机源和指定的单字节字符集生成字符串。
func (r rand) SecureStringWithCharset(length int, charset string) (string, error) {
	return r.SecureStringFrom(cryptorand.Reader, length, charset)
}

// SecureStringFrom 使用调用方提供的随机源生成字符串，便于测试和受控环境注入随机源。
func (rand) SecureStringFrom(reader io.Reader, length int, charset string) (string, error) {
	if reader == nil {
		return "", ErrNilReader
	}
	if length < 0 {
		return "", ErrInvalidLength
	}
	if length == 0 {
		return "", nil
	}
	charsetLength := len(charset)
	if charsetLength < 2 || charsetLength > 128 {
		return "", ErrInvalidCharset
	}
	var seen [128]bool
	for _, value := range []byte(charset) {
		if value >= 128 || seen[value] {
			return "", ErrInvalidCharset
		}
		seen[value] = true
	}

	// 仅接收可被均匀映射的字节，避免直接取模带来的偏差。
	limit := 256 - (256 % charsetLength)
	result := make([]byte, length)
	buffer := make([]byte, length+(length/4)+1)
	for index := 0; index < length; {
		if _, err := io.ReadFull(reader, buffer); err != nil {
			return "", err
		}
		for _, value := range buffer {
			if int(value) >= limit {
				continue
			}
			result[index] = charset[int(value)%charsetLength]
			index++
			if index == length {
				break
			}
		}
	}
	return string(result), nil
}

// SecureInt 返回区间 [min, max) 内的密码学安全随机整数。
func (r rand) SecureInt(min, max int) (int, error) {
	return r.SecureIntFrom(cryptorand.Reader, min, max)
}

// SecureIntFrom 使用调用方提供的随机源返回区间 [min, max) 内的整数。
func (rand) SecureIntFrom(reader io.Reader, min, max int) (int, error) {
	if reader == nil {
		return 0, ErrNilReader
	}
	if min >= max {
		return 0, ErrInvalidRange
	}

	minimum := big.NewInt(int64(min))
	span := new(big.Int).Sub(big.NewInt(int64(max)), minimum)
	value, err := cryptorand.Int(reader, span)
	if err != nil {
		return 0, err
	}
	value.Add(value, minimum)
	return int(value.Int64()), nil
}
