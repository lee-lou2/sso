package client

import (
	"gorm.io/gorm"
	"sso/pkg/util"
	"time"
)

// Client 클라이언트 정보
type Client struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string `gorm:"unique;not null;type:varchar(50)" json:"name"`
	GlobalService bool   `gorm:"default:false;" json:"global_service"`
	Homepage      string `gorm:"null;type:varchar(255)" json:"homepage"`
	CallbackUri   string `gorm:"null;type:varchar(255)" json:"callback_uri"`
	SecretKey     string `gorm:"not null;type:varchar(16)" json:"secret_key"`
}

// BeforeCreate 암호키 생성
func (c *Client) BeforeCreate(tx *gorm.DB) error {
	// 시크릿 키 생성
	secretKey, err := util.GenerateKey(10)
	if err != nil {
		return err
	}
	c.SecretKey = secretKey
	return nil
}
