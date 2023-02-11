package util

import "math/rand"

const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const IntBytes = "0123456789"

// RandStringBytes 랜덤 텍스트 생성
func RandStringBytes(n int, letter string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
