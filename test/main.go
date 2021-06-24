package main

import (
	"fmt"
	_chan "github.com/hzy9738/go-utils/chans"
	"github.com/hzy9738/go-utils/strings"
)

type Student struct {
	Name string
}

func main() {
	res, _ := _chan.GetResByParamSlice([]string{"1", "2", "3"}, 10, func(param interface{}, resChan chan<- interface{}) {
		resChan <- param
	})
	fmt.Println(res)
	strings.ReverseString("abc")
}
