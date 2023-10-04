package user

import "gorm.io/gorm"

type UserModule struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserModule {
	return UserModule{
		db: db,
	}
}
