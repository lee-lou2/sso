package util

// Contains 슬라이스 포함 여부 확인
func Contains[T int | string](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// GetKeys 맵 키 조회
func GetKeys[T int | string](elems map[T][]interface{}) []T {
	keys := make([]T, len(elems))
	i := 0
	for k := range elems {
		keys[i] = k
		i++
	}
	return keys
}
