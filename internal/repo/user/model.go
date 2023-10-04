package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Ksuid      string
	Email      string
	Password   string
	InsertedAt time.Time
	UpdatedAt  time.Time
	DeletedAt  *gorm.DeletedAt
}
