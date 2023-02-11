package errors

import (
	"log"
)

// Alert 오류 알람
func Alert(err error) {
	log.Println(err.Error())
}
