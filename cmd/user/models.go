package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sso/cmd/client"
)

// User 사용자
type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Password string `gorm:"null" json:"password"`
	Verified bool   `gorm:"default:false;not null"`
}

func (u *User) SetPassword(password string) {
	hashPassword, _ := hashPassword(password)
	u.Password = hashPassword
}

// Provider 공급사 - 사용자 연동
type Provider struct {
	gorm.Model
	UserID     int
	IsActive   bool   `gorm:"default:true"`
	User       User   `gorm:"constraint:OnDelete:CASCADE;"`
	Provider   int    `gorm:"not null;default:0" json:"provider"`
	ProviderID string `gorm:"null" json:"provider_id"`
}

// Group 그룹
type Group struct {
	gorm.Model
	UUID     string `gorm:"type:uuid"`
	Name     string `gorm:"not null;type:varchar(50)" json:"name"`
	ClientId int
	Client   client.Client `gorm:"constraint:OnDelete:CASCADE;" json:"client"`
}

// BeforeCreate UUID 생성
func (g *Group) BeforeCreate(tx *gorm.DB) error {
	// UUID 생성
	_uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	g.UUID = _uuid.String()
	return nil
}

// GroupUser 그룹 지정자
type GroupUser struct {
	gorm.Model
	UserID  int
	User    User `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
	GroupId int
	Group   Group `gorm:"constraint:OnDelete:CASCADE;" json:"group"`
}

// Role 그룹별 권한
type Role struct {
	gorm.Model
	Name        string `gorm:"unique;not null;type:varchar(50)" json:"name"`
	Description string `gorm:"null;type:varchar(255)" json:"description"`
	GroupID     int
	Group       Group `gorm:"constraint:OnDelete:CASCADE;"`
}

// TFA 그룹별 2단계 인증
type TFA struct {
	gorm.Model
	Name        string `gorm:"unique;not null;type:varchar(50)" json:"name"`
	Description string `gorm:"null;type:varchar(255)" json:"description"`
	Interval    int    `gorm:"default:0" json:"interval"`
	GroupID     int
	Group       Group `gorm:"constraint:OnDelete:CASCADE;"`
}
