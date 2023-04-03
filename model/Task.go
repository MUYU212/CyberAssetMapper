package model

import (
	"fmt"
	"time"
)

func InsertTask(taskName, domain string) (int64, error) {
	sqlStr := "insert into t_task(task_name,created_at,domain,state) values (?,?,?,?)"
	//默认未开始是0，开始是1，完成是2
	ret, err := db.Exec(sqlStr, taskName, time.Now(), domain, 0)
	if err != nil {
		fmt.Println("Failed to insert data:", err)
		return 0, err
	}
	lastInsertId, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}
