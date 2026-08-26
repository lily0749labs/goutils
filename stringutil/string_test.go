package stringutil

import "testing"

// TestUnicodeStringOperations 验证截断、反转和脱敏不会拆开 Unicode 字符。
func TestUnicodeStringOperations(t *testing.T) {
	t.Parallel()
	if got := String.Truncate("你好世界", 3); got != "你好世" {
		t.Fatalf("Truncate() = %q", got)
	}
	if got := String.TruncateWithSuffix("你好世界", 3, "…"); got != "你好…" {
		t.Fatalf("TruncateWithSuffix() = %q", got)
	}
	if got := String.TruncateWithSuffix("abcdef", 2, "..."); got != ".." {
		t.Fatalf("过长 suffix = %q", got)
	}
	if got := String.Reverse("A你🙂"); got != "🙂你A" {
		t.Fatalf("Reverse() = %q", got)
	}
	if got := String.Mask("张三丰一", 1, 1); got != "张**一" {
		t.Fatalf("Mask() = %q", got)
	}
	if got := String.MaskWith("13800138000", 3, 4, '•'); got != "138••••8000" {
		t.Fatalf("MaskWith() = %q", got)
	}
}

// TestNamingConversions 验证常见命名格式和连续大写缩写的转换。
func TestNamingConversions(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		input, snake, kebab, camel, pascal string
	}{
		{input: "helloWorld", snake: "hello_world", kebab: "hello-world", camel: "helloWorld", pascal: "HelloWorld"},
		{input: "HTTPServer2Port", snake: "http_server_2_port", kebab: "http-server-2-port", camel: "httpServer2Port", pascal: "HttpServer2Port"},
		{input: " hello-world_test ", snake: "hello_world_test", kebab: "hello-world-test", camel: "helloWorldTest", pascal: "HelloWorldTest"},
	}
	for _, tc := range testCases {
		if got := String.Snake(tc.input); got != tc.snake {
			t.Errorf("Snake(%q) = %q，期望 %q", tc.input, got, tc.snake)
		}
		if got := String.Kebab(tc.input); got != tc.kebab {
			t.Errorf("Kebab(%q) = %q，期望 %q", tc.input, got, tc.kebab)
		}
		if got := String.Camel(tc.input); got != tc.camel {
			t.Errorf("Camel(%q) = %q，期望 %q", tc.input, got, tc.camel)
		}
		if got := String.Pascal(tc.input); got != tc.pascal {
			t.Errorf("Pascal(%q) = %q，期望 %q", tc.input, got, tc.pascal)
		}
	}
}

// TestWhitespaceAndBoundaries 验证空白归一化与边界参数。
func TestWhitespaceAndBoundaries(t *testing.T) {
	t.Parallel()
	if !String.IsBlank(" \t\n") || String.IsBlank(" a ") {
		t.Fatal("IsBlank() 结果不正确")
	}
	if got := String.NormalizeSpace("  hello\t世界\n test "); got != "hello 世界 test" {
		t.Fatalf("NormalizeSpace() = %q", got)
	}
	if String.Truncate("abc", 0) != "" || String.Mask("abc", 2, 2) != "abc" {
		t.Fatal("边界参数结果不正确")
	}
}
