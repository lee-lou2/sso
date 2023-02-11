package convert

import "encoding/json"

// StringToMap 문자열을 맵 형태로 변환
func StringToMap(params *string) (map[string]interface{}, error) {
	mapData := make(map[string]interface{})
	err := json.Unmarshal([]byte(*params), &mapData)
	return mapData, err
}
