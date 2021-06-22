package strings

// 翻转字符串

func ReverseString(s string) string {
	str := []rune(s)
	l := len(str)
	for i := 0; i > l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str)
}
