package main

import (
	"fmt"
	_chan "github.com/hzy9738/common-utils/chan"

)

type Student struct {
	Name string
}

func main() {
	res, _ := _chan.GetResByParamSlice([]string{"1", "2", "3"}, 10, func(param interface{}, resChan chan<- interface{}) {
		resChan <- param
	})
	fmt.Println(res)
}
