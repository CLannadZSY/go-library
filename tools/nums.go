package nums

// Comma inserts commas in a non-negative decimal integer string.
// 在非负十进制整数字符串中插入逗号
// 123456 -> 123,456
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return Comma(s[:n-3]) + "," + s[n-3:]
}