package tree

import (
	"github.com/hzy9738/common-utils/convert"
	"reflect"
)

/**
 * 得到子级切片
 * @param interface{}
 * @return []interface{}
 */
func (t OriginData) GetChild(myID interface{}) []interface{} {
	myIDType := reflect.TypeOf(myID).Name()
	var data []interface{}
	for _, v := range t.dataMap {
		if _, ok := v[t.MyID]; !ok {
			continue
		}
		if _, ok := v[t.PName]; ok {
			p, err := convert.ValueSwapTo(v[t.PName], myIDType)
			if err == nil && p == myID {
				data = append(data, v)
			}
		}
	}
	return data
}

/**
 *
 * 获取树状切片
 * @param interface{}
 * @return []interface{}
 */
func (t OriginData) GetTreeArray(myID interface{}) []interface{} {
	myIDType := reflect.TypeOf(myID).Name()
	childs := t.GetChild(myID)
	n := 0
	var data []interface{}
	total := len(childs)
	if total > 0 {
		for _, v := range childs {
			newData := map[string]interface{}{}
			_ = convert.SwapTo(v, &newData)
			id, err := convert.ValueSwapTo(newData[t.MyID], myIDType)
			if err == nil {
				newData[t.CName] = t.GetTreeArray(id)
				data = append(data, newData)
				n++
			}
		}
	}
	return data
}
