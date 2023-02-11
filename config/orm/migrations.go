package orm

import (
	"sso/cmd/client"
	"sso/cmd/user"
	"sso/config/database"
	"sso/pkg/security"
)

// AutoMigrate 마이그레이션
func AutoMigrate() {
	db, err := database.GetDatabase()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&user.User{},
		&user.Provider{},
		&user.Group{},
		&user.GroupUser{},
		&user.Role{},
		&user.TFA{},
		&client.Client{},
		&security.RSAKeyPair{},
	)
}
