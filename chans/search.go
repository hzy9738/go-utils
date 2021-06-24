package chans

import (
	"errors"
	"fmt"
	"reflect"
)

type GetRes func(param interface{}, resChan chan<- interface{})
type GetResSlice func(param []interface{}, resChan chan<- []interface{})

// 分流聚合查询-1
// params参数是切片
// chanLimit是chan的容量和协成查询的条数
// getResFunc回调函数，必须在回调中赋值给resChan

func GetResByParamSlice(params interface{}, chanLimit int, getResFunc GetRes) (interface{}, error) {
	s := reflect.ValueOf(params)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("params given a non-slice type")
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	resChan := make(chan interface{}, chanLimit)   //定义结果channel
	paramChan := make(chan interface{}, chanLimit) //定义参数channel
	closeChan := make(chan bool, chanLimit)        //定义终止channel
	go func() {
		for _, param := range ret {
			paramChan <- param
		}
		close(paramChan)
	}()
	//处理paramChannel中的信息
	for i := 0; i < chanLimit; i++ {
		go func(paramChan <-chan interface{}, resChan chan<- interface{}, closeChan chan<- bool) {
			for param := range paramChan {
				getResFunc(param, resChan)
			}
			closeChan <- true
		}(paramChan, resChan, closeChan)
	}
	//判断是否执行完成，信息聚合
	go func() {
		for i := 0; i < chanLimit; i++ {
			<-closeChan
		}
		close(resChan)
		close(closeChan)
	}()
	//获取结果
	var res []interface{}
	for r := range resChan {
		res = append(res, r)
	}
	return res, nil
}


// 分流聚合查询-2
// params参数是长度巨大的切片
// chanLimit是chan的容量和协成查询的条数
// getResFunc回调函数，必须在回调中赋值给resChan

func GetResByLongParamSlice(params interface{}, chanLimit int, getResFunc GetResSlice) (interface{}, error) {
	s := reflect.ValueOf(params)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("params given a non-slice type")
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	resChan := make(chan []interface{}, chanLimit)   //定义结果channel
	paramChan := make(chan []interface{}, chanLimit) //定义参数channel
	closeChan := make(chan bool, chanLimit)        //定义终止channel
	rangeSize :=  s.Len() / (chanLimit -1 )
	go func() {
		for i := 0; i < chanLimit; i++ {
			start := i * rangeSize
			end := (i + 1) * rangeSize
			if i == (chanLimit - 1) {
				paramChan <- ret[start:]
			} else {
				paramChan <- ret[start:end]
			}
		}
		close(paramChan)
	}()
	//处理paramChannel中的信息
	for i := 0; i < chanLimit; i++ {
		go func(paramChan <-chan []interface{}, resChan chan<- []interface{}, closeChan chan<- bool) {
			for param := range paramChan {
				getResFunc(param, resChan)
			}
			closeChan <- true
		}(paramChan, resChan, closeChan)
	}
	//判断是否执行完成，信息聚合
	go func() {
		for i := 0; i < chanLimit; i++ {
			<-closeChan
		}
		close(resChan)
		close(closeChan)
	}()
	//获取结果
	var res []interface{}
	for r := range resChan {
		res = append(res, r)
	}
	return res, nil
}

func test() {
	res, _ := GetResByParamSlice([]string{"1", "2", "3"}, 10, func(param interface{}, resChan chan<- interface{}) {
		resChan <- param
	})
	fmt.Println(res)
}
