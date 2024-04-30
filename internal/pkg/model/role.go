package model

import "gorm.io/gorm"

type RoleDB struct {
	gorm.Model
	AccountID int64  `json:"account_id" gorm:"column:account_id;type:bigint(20);not null;comment:账号id"`
	UID       int64  `json:"uid" gorm:"column:uid;type:bigint(20);not null;comment:角色id"`
	RoleData  string `json:"role_data" gorm:"column:role_data;type:text;comment:账号数据"`
	ServerID  int32  `json:"server_id" gorm:"column:server_id;type:int;not null;comment:服务器id"`
}

func (*RoleDB) TableName() string {
	return "role_db"
}
