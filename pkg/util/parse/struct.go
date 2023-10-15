package parse

import (
	"encoding/json"
	"roofix/pkg/util/log"
)

func StructToMap(obj interface{}) map[string]interface{} {
	// Convert to a json string
	data, err := json.Marshal(obj)
	if err != nil {
		log.Error(err)
		return nil
	}

	var res map[string]interface{}
	if err = json.Unmarshal(data, &res); err != nil {
		log.Error(err)
		return nil
	}

	return res
}
