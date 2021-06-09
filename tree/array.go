package tree

import common_utils "github.com/hzy9738/common-utils"


type Tree struct {
	Name      string `json:"name"`
	Pid       int    `json:"pid"`
	ChildList []Tree `json:"childlist"`
}

/**
 * 得到子级切片
 * @param int64
 * @return []interface{}
 */
func (t OriginData) GetChild(myID int64) []interface{} {
	var data []interface{}
	for _, v := range t.dataMap {
		if _, ok := v[t.MyID]; !ok {
			continue
		}
		if _, ok := v[t.PName]; ok {
			p := int64(v[t.PName].(float64))
			if p == myID {
				data = append(data, v)
			}
		}
	}
	return data
}

/**
 *
 * 获取树状数组
 * @param string $myid 要查询的ID
 * @param string $itemprefix 前缀
 * @return map[int64]interface
 */
func (t OriginData) GetTreeArray(myID int64) []interface{} {
	childs := t.GetChild(myID)
	n := 0
	var data []interface{}
	total := len(childs)
	if total > 0 {
		for _, v := range childs {
			newData := map[string]interface{}{}
			_ = common_utils.SwapTo(v, &newData)
			id := int64(newData[t.MyID].(float64))
			newData[t.CName] = t.GetChild(id)
			data = append(data, newData)
			n++
		}
	}
	return data
}
