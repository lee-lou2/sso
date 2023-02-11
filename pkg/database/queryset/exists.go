package queryset

import (
	"gorm.io/gorm"
)

// Exists 포함 여부 확인
func Exists(tx *gorm.DB) bool {
	var exists bool

	err := tx.Select("count(*) > 0").Find(&exists).Error
	if err != nil {
		return false
	}
	return exists
}
