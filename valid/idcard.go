package valid

import (
	"regexp"
	"strconv"
)

// IDCard 验证身份证号（18 或 15 位）。
func (f valid) IDCard(str string) bool {
	if len(str) != 15 && len(str) != 18 {
		return false
	}
	if len(str) == 18 {
		return f.IDCard18(str)
	}
	return f.IDCard15(str)
}

// IDCard18 验证 18 位身份证号。
func (valid) IDCard18(id string) bool {
	regExp := "^[1-9]\\d{5}(19|20)\\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$"
	if match, _ := regexp.MatchString(regExp, id); !match {
		return false
	}

	// 对身份证号码的最后一位进行校验
	// 根据身份证号码的规则，最后一位可能是数字0-9，也可能是字符X（表示10）
	// 将字符X转换成数字10进行校验
	lastChar := id[len(id)-1]
	var lastNum int
	if lastChar == 'X' || lastChar == 'x' {
		lastNum = 10
	} else {
		lastNum, _ = strconv.Atoi(string(lastChar))
	}
	// 对身份证号码的前17位进行加权和校验
	// 加权系数，根据规则固定
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 计算加权和
	sum := 0
	for i := 0; i < len(weights); i++ {
		num, _ := strconv.Atoi(string(id[i]))
		sum += num * weights[i]
	}

	// 计算校验码
	checkCode := sum % 11
	checkCodeMap := map[int]string{
		0:  "1",
		1:  "0",
		2:  "10", // 身份证最后一位是X，加权求和是10
		3:  "9",
		4:  "8",
		5:  "7",
		6:  "6",
		7:  "5",
		8:  "4",
		9:  "3",
		10: "2",
	}
	// 校验身份证号码的最后一位
	return checkCodeMap[checkCode] == strconv.Itoa(lastNum)
}

// IDCard15 验证 15 位身份证号。
func (valid) IDCard15(idCard string) bool {
	// 验证是否为15位数字
	if match, _ := regexp.MatchString(`^\d{15}$`, idCard); !match {
		return false
	}

	// 将身份证号前两位转换成省份代码
	provinceCode, err := strconv.Atoi(idCard[:2])
	if err != nil || provinceCode < 11 || provinceCode > 91 {
		return false
	}

	// 验证生日是否正确
	year := strconv.Itoa(1900 + int(idCard[6]-'0')*10 + int(idCard[7]-'0'))
	month := string(idCard[8:10])
	day := string(idCard[10:12])
	if match, _ := regexp.MatchString(`^(19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])$`, year+month+day); !match {
		return false
	}

	return true
}

// Deprecated: 使用 Valid.IDCard。
func IsIDCard(str string) bool { return Valid.IDCard(str) }

// Deprecated: 使用 Valid.IDCard18。
func IsIDCard18(id string) bool { return Valid.IDCard18(id) }

// Deprecated: 使用 Valid.IDCard15。
func IsIDCard15(idCard string) bool { return Valid.IDCard15(idCard) }
