package valid

import "net"

// IPv4 验证是否为 IPv4 地址。
func (valid) IPv4(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil && ip.To4() != nil
}

// IPv6 验证是否为 IPv6 地址。
func (valid) IPv6(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil && ip.To4() == nil
}

// Deprecated: 使用 Valid.IPv4。
func IsIPv4(input string) bool { return Valid.IPv4(input) }

// Deprecated: 使用 Valid.IPv6。
func IsIPv6(input string) bool { return Valid.IPv6(input) }
