package dao

import (
	"CyberAssetMapper/src/model"
	"CyberAssetMapper/src/service/dto"
	"fmt"
)

var taskDao *TaskDao

type TaskDao struct {
	BaseDao
}

func NewTaskDao() *TaskDao {
	if taskDao == nil {
		taskDao = &TaskDao{NewBaseDao()}
	}
	return taskDao
}

// 添加任务，但是添加了任务之后还得添加任务对应的域名
func (m *TaskDao) AddTask(iTaskAddDTO *dto.TaskAddDTO) error {
	var iTask model.Task
	//这里的逻辑得是先添加了任务之后，获取到任务的id，然后再添加任务对应的域名
	iTaskAddDTO.ConvertToModel(&iTask)
	fmt.Println(iTask)
	err := m.Orm.Save(&iTask).Error

	if err == nil {
		fmt.Printf("添加任务成功，任务id为：%d\n", iTask.ID)
	}
	return err
}
