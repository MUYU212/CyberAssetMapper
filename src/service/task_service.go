package service

import (
	"CyberAssetMapper/src/dao"
	"CyberAssetMapper/src/service/dto"
)

var taskService *TaskService

type TaskService struct {
	BaseService
	Dao *dao.TaskDao
}

func NewTaskService() *TaskService {
	if taskService == nil {
		taskService = &TaskService{
			Dao: dao.NewTaskDao(),
		}
	}
	return taskService
}

func (m *TaskService) AddTask(iTaskAddDTO *dto.TaskAddDTO) error {
	return m.Dao.AddTask(iTaskAddDTO)
}
