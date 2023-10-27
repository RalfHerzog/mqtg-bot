package models

import (
	"gorm.io/gorm"
)

type DbUser struct {
	gorm.Model
	ChatID        int64  `gorm:"type:bigint"`
	UserName      string `gorm:"type:varchar(255)"`
	MqttUrl       string `gorm:"type:varchar(255)"` // (tcp|ssl|ws|wss)://user:password@host:port/path
	Connected     bool
	Subscriptions []*Subscription
	DbMenu        []byte
	Roles         []*DbUserRole
}
