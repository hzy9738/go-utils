package tree

import (
	"errors"
	"github.com/hzy9738/common-utils/convert"
	"reflect"
)

type OriginData struct {
	Data    interface{}
	dataMap []map[string]interface{}
	MyID    string
	Name    string
	PName   string
	CName   string
}

type Option func(*OriginData)

func SetTreeOriginData(data interface{}) Option {
	return func(client *OriginData) {
		client.Data = data
	}
}

func SetTreeChildName(cName string) Option {
	return func(client *OriginData) {
		client.CName = cName
	}
}

func SetTreePidName(pName string) Option {
	return func(client *OriginData) {
		client.PName = pName
	}
}

func SetTreeIdName(myIDName string) Option {
	return func(client *OriginData) {
		client.MyID = myIDName
	}
}

func NewInit(opts ...Option) (OriginData, error) {
	defaultClient := OriginData{
		Data:    nil,
		dataMap: nil,
		MyID:    "id",
		Name:    "name",
		PName:   "pid",
		CName:   "childlist",
	}
	for _, o := range opts {
		o(&defaultClient)
	}
	if defaultClient.Data == nil {
		return defaultClient, errors.New("Data不能为空")
	} else {
		typeOfData := reflect.TypeOf(defaultClient.Data).Kind()
		if typeOfData != reflect.Slice {
			return defaultClient, errors.New("Data必须是切片类型")
		}
	}
	err := convert.SwapTo(defaultClient.Data, &defaultClient.dataMap)
	if err != nil {
		return defaultClient, nil
	}
	return defaultClient, nil
}
