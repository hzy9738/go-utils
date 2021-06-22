package strings

import "strings"

// 判断字符串中字符是否全都不同
// 预设s为ASCII字符，ASCII字符 字符⼀共有256个，其中128个是常⽤字符，可以在键盘上输⼊。128之后的是键盘上⽆法找到的

func IsUniqueString(s string) bool {
	//if strings.Count(s, "") > 3000 {
	//	return false
	//}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

func IsUniqueString2(s string) bool {
	//if strings.Count(s, "") > 3000 {
	//	return false
	//}
	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}
