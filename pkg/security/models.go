package security

import "time"

// RSAKeyPair 공개/비밀키
type RSAKeyPair struct {
	ID         uint   `gorm:"primarykey"`
	PrivateKey string `gorm:"not null" json:"private_key"`
	PublicKey  string `gorm:"not null" json:"public_key"`
	CreatedAt  time.Time
}
