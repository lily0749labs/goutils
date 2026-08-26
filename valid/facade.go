package valid

// Valid provides object-style access to the validation functions in this package.
var Valid = facade{}

type facade struct{}

func (facade) IsAllChinese(value string) bool     { return IsAllChinese(value) }
func (facade) IsBankCardNo(value string) bool     { return IsBankCardNo(value) }
func (facade) IsChineseName(value string) bool    { return IsChineseName(value) }
func (facade) IsContainChinese(value string) bool { return IsContainChinese(value) }
func (facade) IsDate(value string) bool           { return IsDate(value) }
func (facade) IsDateTime(value string) bool       { return IsDateTime(value) }
func (facade) IsDecimal(value string) bool        { return IsDecimal(value) }
func (facade) IsEmail(value string) bool          { return IsEmail(value) }
func (facade) IsEnglishName(value string) bool    { return IsEnglishName(value) }
func (facade) IsIDCard(value string) bool         { return IsIDCard(value) }
func (facade) IsIDCard15(value string) bool       { return IsIDCard15(value) }
func (facade) IsIDCard18(value string) bool       { return IsIDCard18(value) }
func (facade) IsIPv4(value string) bool           { return IsIPv4(value) }
func (facade) IsIPv6(value string) bool           { return IsIPv6(value) }
func (facade) IsJSON(value string) bool           { return IsJSON(value) }
func (facade) IsMobile(value string) bool         { return IsMobile(value) }
func (facade) IsNumber(value string) bool         { return IsNumber(value) }
func (facade) IsPassword(value string) bool       { return IsPassword(value) }
func (facade) IsPostalCode(value string) bool     { return IsPostalCode(value) }
func (facade) IsQQ(value string) bool             { return IsQQ(value) }
func (facade) IsTelephone(value string) bool      { return IsTelephone(value) }
func (facade) IsTime(value string) bool           { return IsTime(value) }
func (facade) IsURL(value string) bool            { return IsURL(value) }
func (facade) IsWeChat(value string) bool         { return IsWeChat(value) }
func (facade) IsWeibo(value string) bool          { return IsWeibo(value) }
