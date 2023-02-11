package methods

import (
	"context"
	"sso/config/cache"
	"sso/config/errors"
	"sso/config/errors/status"
	"time"
)

// 기본 콘텍스트
var ctx = context.Background()

// GetValue 캐시 조회
func GetValue(key string) (string, error) {
	client := cache.GetRedis()
	return client.Get(ctx, key).Result()
}

// SetValue 캐시 수정
func SetValue(key string, value string, expiration int) error {
	client := cache.GetRedis()
	if err := client.Set(ctx, key, value, time.Duration(expiration)*time.Second).Err(); err != nil {
		return errors.New(status.CacheDataConfigError, err)
	}
	return nil
}

// PushValue 캐시 값 추가
func PushValue(key string, value string) error {
	client := cache.GetRedis()
	if err := client.LPush(ctx, key, value).Err(); err != nil {
		errors.New(status.CacheDataInsertError, err)
	}
	return nil
}

// PopValue 캐시 값 제거
func PopValue(key string) (string, error) {
	client := cache.GetRedis()
	value, err := client.RPop(ctx, key).Result()
	if err != nil {
		return "", errors.New(status.CacheDataExtractionError, err)
	}
	return value, nil
}

// GetRange 값 리스트 조회
func GetRange(key string) ([]string, error) {
	client := cache.GetRedis()
	values, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, errors.New(status.CacheListRetrievalError, err)
	}
	return values, nil
}

// GetLength 전체 카운트 조회
func GetLength(key string) (int, error) {
	client := cache.GetRedis()
	count, err := client.LLen(ctx, key).Result()
	if err != nil {
		return 0, errors.New(status.CacheDataCountRetrievalError, err)
	}
	return int(count), nil
}

// PushValueSetLimit 최대 제한을 두고 추가
func PushValueSetLimit(key string, value string, limit int) error {
	var err error
	client := cache.GetRedis()

	// 데이터 수 조회
	count, err := client.LLen(ctx, key).Result()
	if err != nil {
		return errors.New(status.CacheDataCountRetrievalError, err)
	}

	// 최대 제한 수를 초과하는 경우 초과되는 값들은 제거
	if int(count) >= limit {
		for i := 0; i <= int(count)-limit; i++ {
			client.RPop(ctx, key)
		}
	}
	return client.LPush(ctx, key, value).Err()
}

// ExistsValue 리스트 내 해당 데이터가 있는지 확인
func ExistsValue(key string, value string) (bool, error) {
	client := cache.GetRedis()

	// 리스트 조회
	values, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return false, errors.New(status.CacheListRetrievalError, err)
	}

	// 리스트내 존재 여부 확인
	for _, v := range values {
		if v == value {
			return true, nil
		}
	}
	return false, nil
}
