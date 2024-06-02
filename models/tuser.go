package models

import "gorm.io/gorm"

const (
	ROLE_USER  = 0
	ROLE_ADMIN = 1
)

type TUser struct {
	gorm.Model
	TUserID  int64
	UserName string
	Role     int  `gorm:"default:0"`
	Notify   bool `gorm:"default:false"`
	CmdCount uint `gorm:"default:0"`
}
