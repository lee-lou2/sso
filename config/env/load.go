package env

import (
	"github.com/joho/godotenv"
	"log"
)

// Load 환경 변수 불러오기
func Load() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("환경 변수 파일이 존재하지 않아 서버 환경 변수를 참조합니다")
	}
}
