package dto

import "CyberAssetMapper/src/model"

type TaskAddDTO struct {
	ID         uint
	TaskName   string `json:"taskName" binding:"required" message:"任务名不能为空"`
	RootDomain string `json:"rootDomain" binding:"required" message:"主域名不能为空"`
}

func (m *TaskAddDTO) ConvertToModel(iTask *model.Task) {
	iTask.TaskName = m.TaskName
}
