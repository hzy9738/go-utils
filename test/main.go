package main

import (
	"fmt"
	"github.com/hzy9738/go-utils/strings"
)

type Student struct {
	Name string
}

func main() {
	//res, _ := _chan.GetResByParamSlice([]string{"1", "2", "3"}, 10, func(param interface{}, resChan chan<- interface{}) {
	//	resChan <- param
	//})
	//fmt.Println(res)
	//strings.ReverseString("abc")
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(  rand.Intn(999)  )
	c := "hzy"
	cd := strings.To2Byte(c)
	cc := strings.From2Byte(cd)
	bb := []byte(c)
	fmt.Println(&c)
	fmt.Println(&cd == &bb)
	fmt.Println(&cc)


}
// 1 6 11 16 21
// 2 7 12 17 22
