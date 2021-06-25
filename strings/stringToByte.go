package strings

import (
	"reflect"
	"unsafe"
)

//  unsafe.Pointer(&a)方法可以得到变量a的地址。
// (*reflect.StringHeader)(unsafe.Pointer(&a)) 可以把字符串a转成底层结构的形式。
// (*[]byte)(unsafe.Pointer(&ssh)) 可以把ssh底层结构体转成byte的切片的指针。
// 再通过 *转为指针指向的实际内容。

func To2Byte(s string) []byte {
	ss := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&ss))
}

func From2Byte(b []byte) string {
	bb := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	return *(*string)(unsafe.Pointer(&bb))
}
