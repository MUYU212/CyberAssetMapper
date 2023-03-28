package model

import (
	"CyberAssetMapper/model/db"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model        //GORM自带的基本属性，包括ID，以及更新时间等等
	Domain     string //这里一定要驼峰命令，否则GORM是无法识别出来属性的，并且会自动转换成蛇形命名在数据库中
	IsDeleted  bool
}

func Create_taskTable() {
	db.GLOBAL_DB.AutoMigrate(&Task{})
	//创建Task表
}
