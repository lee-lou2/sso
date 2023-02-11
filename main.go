package main

import (
	"sso/api/gin"
	"sso/api/grpc"
	"sso/config/env"
	"sso/config/orm"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// 환경 설정
	env.Load()

	// 데이터 마이그레이션
	orm.AutoMigrate()

	// API 리스너 실행
	wg.Add(2)
	go func() {
		gin.Run()
		wg.Done()
	}()
	go func() {
		grpc.Run()
		wg.Done()
	}()
	wg.Wait()
}
