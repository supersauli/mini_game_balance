package model

import (
	"log"
	"mini_game_balance/global"
)

func Init() {
	err := global.MySql.AutoMigrate(&RoleDB{})
	if err != nil {
		log.Panic("init mysql fail ,err:", err)
	}
}
