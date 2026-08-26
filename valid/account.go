package valid

import (
	"encoding/json"
	"regexp"
)

// Email 验证是否为电子邮箱地址。
func (facade) Email(input string) bool {
	// 定义邮箱地址的正则表达式
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, input)
	return match && err == nil
}

// JSON 验证是否为 JSON。
func (facade) JSON(input string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(input), &js) == nil
}

// QQ 验证是否为 QQ 号。
func (facade) QQ(qq string) bool {
	match, _ := regexp.MatchString(`^[1-9][0-9]{4,12}$`, qq)
	return match
}

// WeChat 验证是否为微信号。
func (facade) WeChat(wechat string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z][-_a-zA-Z0-9]{6,20}$`, wechat)
	return match
}

// Weibo 验证是否为微博 ID。
func (facade) Weibo(weibo string) bool {
	if len(weibo) < 6 || len(weibo) > 20 {
		return false
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z][\w-]*$`, weibo); !matched {
		return false
	}

	return true
}

// Password 验证密码是否合法。
// 密码长度在6-20个字符之间，必须包含数字、字母和特殊符号
func (facade) Password(password string) bool {
	if len(password) < 6 || len(password) > 20 {
		return false
	}

	if matched, _ := regexp.MatchString(`[a-zA-Z]`, password); !matched {
		return false
	}

	if matched, _ := regexp.MatchString(`\d`, password); !matched {
		return false
	}

	if matched, _ := regexp.MatchString(`[^a-zA-Z\d]`, password); !matched {
		return false
	}

	return true
}

// Deprecated: 使用 Valid.Email。
func IsEmail(input string) bool { return Valid.Email(input) }

// Deprecated: 使用 Valid.JSON。
func IsJSON(input string) bool { return Valid.JSON(input) }

// Deprecated: 使用 Valid.QQ。
func IsQQ(qq string) bool { return Valid.QQ(qq) }

// Deprecated: 使用 Valid.WeChat。
func IsWeChat(wechat string) bool { return Valid.WeChat(wechat) }

// Deprecated: 使用 Valid.Weibo。
func IsWeibo(weibo string) bool { return Valid.Weibo(weibo) }

// Deprecated: 使用 Valid.Password。
func IsPassword(password string) bool { return Valid.Password(password) }
