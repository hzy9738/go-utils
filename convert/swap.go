package convert

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"time"
)

func SwapTo(request, obj interface{}) error {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, obj)
}

//类型转换
func TypeSwapTo(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}

	//else if .......增加其他一些类型的转换

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}

//将传入的值按照需要的类型转换
func ValueSwapTo(value interface{}, toType string) (val interface{}, err error) {
	valueType := reflect.TypeOf(value).Name()
	if valueType == toType {
		return value, nil
	}
	switch toType {
	case "string":
		switch valueType {
		case "float64":
			val = strconv.FormatFloat(value.(float64), 'E', -1, 64)
			break
		case "float32":
			val = strconv.FormatFloat(float64(value.(float32)), 'E', -1, 32)
			break
		case "int":
			val = strconv.Itoa(value.(int))
			break
		case "int8":
			val = strconv.Itoa(int(value.(int8)))
			break
		case "int16":
			val = strconv.Itoa(int(value.(int16)))
			break
		case "int32":
			val = strconv.Itoa(int(value.(int32)))
			break
		case "int64":
			val = strconv.Itoa(int(value.(int64)))
			break
		}
		break
	case "float32":
		switch valueType {
		case "float64":
			val = float32(value.(float64))
			break
		}
		break
	case "float64":
		switch valueType {
		case "float32":
			val = float64(value.(float32))
			break
		}
		break
	case "int":
		switch valueType {
		case "string":
			val, err = strconv.Atoi(value.(string))
			break
		case "float32":
			val = int(value.(float32))
			break
		case "float64":
			val = int(value.(float64))
			break
		case "int8":
			val = int(value.(int8))
			break
		case "int16":
			val = int(value.(int16))
			break
		case "int32":
			val = int(value.(int32))
			break
		case "int64":
			val = int(value.(int64))
			break
		}
		break
	case "int8":
		switch valueType {
		case "string":
			var val1 int
			val1, err = strconv.Atoi(value.(string))
			if err != nil {
				val = int8(val1)
			}
			break
		case "float32":
			val = int8(value.(float32))
			break
		case "float64":
			val = int8(value.(float64))
			break
		case "int":
			val = int8(value.(int))
			break
		case "int16":
			val = int8(value.(int16))
			break
		case "int32":
			val = int8(value.(int32))
			break
		case "int64":
			val = int8(value.(int64))
			break
		}
		break
	case "int16":
		switch valueType {
		case "string":
			var val1 int
			val1, err = strconv.Atoi(value.(string))
			if err != nil {
				val = int16(val1)
			}
			break
		case "float32":
			val = int16(value.(float32))
			break
		case "float64":
			val = int16(value.(float64))
			break
		case "int":
			val = int16(value.(int))
			break
		case "int8":
			val = int16(value.(int8))
			break
		case "int32":
			val = int16(value.(int32))
			break
		case "int64":
			val = int16(value.(int64))
			break
		}
		break
	case "int32":
		switch valueType {
		case "string":
			var val1 int
			val1, err = strconv.Atoi(value.(string))
			if err != nil {
				val = int32(val1)
			}
			break
		case "float32":
			val = int32(value.(float32))
			break
		case "float64":
			val = int32(value.(float64))
			break
		case "int":
			val = int32(value.(int))
			break
		case "int8":
			val = int32(value.(int8))
			break
		case "int16":
			val = int32(value.(int16))
			break
		case "int64":
			val = int32(value.(int64))
			break
		}
		break
	case "int64":
		switch valueType {
		case "string":
			var val1 int
			val1, err = strconv.Atoi(value.(string))
			if err != nil {
				val = int64(val1)
			}
			break
		case "float32":
			val = int64(value.(float32))
			break
		case "float64":
			val = int64(value.(float64))
			break
		case "int":
			val = int64(value.(int))
			break
		case "int8":
			val = int64(value.(int8))
			break
		case "int16":
			val = int64(value.(int16))
			break
		case "int32":
			val = int64(value.(int32))
			break
		}
		break
	default:
		err = errors.New("未知类型" + toType)
		break
	}

	//else if .......增加其他一些类型的转换
	return val, err
}
