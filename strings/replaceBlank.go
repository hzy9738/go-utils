package strings

import (
	"errors"
	"strings"
	"unicode"
)

// 字符串替换
// 将字符串中的空格全部替换为“%20”,同时保证字符串由【⼤⼩写的英⽂字⺟组成】

func ReplaceBlank(s string) (string, error) {
	for _, v := range s {
		if string(v) != "" && unicode.IsLetter(v) == false {
			return s, errors.New("包含非英文字符")
		}
	}
	return strings.Replace(s, " ", "%20", -1), nil
}

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case student:
		msg.Name = "111"
	}
}
