package model

import (
	"CyberAssetMapper/model/db"
	"fmt"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model        //GORM自带的基本属性，包括ID，以及更新时间等等
	Domain     string //这里一定要驼峰命令，否则GORM是无法识别出来属性的，并且会自动转换成蛇形命名在数据库中
	IsDeleted  bool
	State      int
}

func Create_taskTable() {
	db.GLOBAL_DB.AutoMigrate(&Task{})
	//创建Task表
}

func Insert_Task(domain string) (*Task, error) {
	task := &Task{
		Domain: domain,
	}
	dbres := db.GLOBAL_DB.Create(task)
	if dbres.Error != nil {
		panic(dbres.Error)
	} else {
		fmt.Println("插入成功")
		return task, nil
	}
}

func Get_Task(id int) Task {
	var task Task
	db.GLOBAL_DB.Select("id,domain").Find(&task, id)
	fmt.Println(task)
	return task
}

// 更新任务状态
func Update_Task(id int, state bool) {
	task := &Task{}

	db := db.GLOBAL_DB.Model(&task).Where("id = ?", id).Update("state", state)
	if db.Error != nil {
		fmt.Println(db.Error)
	}
}
