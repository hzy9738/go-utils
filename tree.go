package common_utils

import (
	"fmt"
)

type Tree struct {
	Name      string `json:"name"`
	Pid       int    `json:"pid"`
	ChildList []Tree `json:"childlist"`
}

type TreeOrigin struct {
	Data  interface{}
	PName string
}

type TreeOption func(*TreeOrigin)

func SetTreeOriginData(data interface{}) TreeOption {
	return func(client *TreeOrigin) {
		client.Data = data
	}
}

func SetTreePidName(pName string) TreeOption {
	return func(client *TreeOrigin) {
		client.PName = pName
	}
}

func NewTree(opts ...TreeOption) TreeOrigin {
	defaultClient := TreeOrigin{
		Data:  nil,
		PName: "pid",
	}
	for _, o := range opts {
		o(&defaultClient)
	}
	return defaultClient
}

func (t *TreeOrigin) GetChild() {
	for  v := t.Data {
		fmt.Println(v)
	}

}
