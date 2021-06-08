package common_utils

import "encoding/json"

func SwapTo(request, obj interface{}) error {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, obj)
}

