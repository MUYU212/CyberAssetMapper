package model

import (
	"CyberAssetMapper/model/db"
	"fmt"
)

func QueryData() {
	// 编写SQL语句
	query := "SELECT * FROM t_task"

	// 执行查询
	rows, err := db.GLOBAL_DB.Query(query)
	if err != nil {
		fmt.Println("Failed to query data:", err)
		return
	}
	defer rows.Close()

	// 遍历结果集
	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println("Failed to scan row:", err)
			return
		}

		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// 检查错误
	err = rows.Err()
	if err != nil {
		fmt.Println("Error during row iteration:", err)
		return
	}
}

//type Task struct {
//	gorm.Model        //GORM自带的基本属性，包括ID，以及更新时间等等
//	Domain     string //这里一定要驼峰命令，否则GORM是无法识别出来属性的，并且会自动转换成蛇形命名在数据库中
//	IsDeleted  bool
//	State      int
//}
//
//func Create_taskTable() {
//	db.GLOBAL_DB.AutoMigrate(&Task{})
//	//创建Task表
//}
//
//func Insert_Task(domain string) (*Task, error) {
//	task := &Task{
//		Domain: domain,
//	}
//	dbres := db.GLOBAL_DB.Create(task)
//	if dbres.Error != nil {
//		panic(dbres.Error)
//	} else {
//		fmt.Println("插入成功")
//		return task, nil
//	}
//}
//
//func Get_Task(id int) Task {
//	var task Task
//	db.GLOBAL_DB.Select("id,domain").Find(&task, id)
//	fmt.Println(task)
//	return task
//}
//
//// 更新任务状态
//func Update_Task(id int, state int) {
//	task := &Task{}
//
//	db := db.GLOBAL_DB.Model(&task).Where("id = ?", id).Update("state", state)
//	if db.Error != nil {
//		fmt.Println(db.Error)
//	}
//}
