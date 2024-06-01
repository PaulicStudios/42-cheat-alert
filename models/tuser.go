package models

import "gorm.io/gorm"

const (
	ROLE_USER  = 0
	ROLE_ADMIN = 1
)

type TUser struct {
	gorm.Model
	TUserID int64
	Role    int
	Notify  bool
}
