package api

import (
	"CyberAssetMapper/src/global"
	"CyberAssetMapper/src/service"
	"CyberAssetMapper/src/service/dto"
	"github.com/gin-gonic/gin"
)

type TaskApi struct {
	BaseApi
	Service *service.TaskService
}

func NewTaskApi() TaskApi {
	return TaskApi{
		BaseApi: NewBaseApi(),
		Service: service.NewTaskService(),
	}
}

func (m TaskApi) AddTask(c *gin.Context) {
	var iTaskAddDTO dto.TaskAddDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx: c,
		DTO: &iTaskAddDTO,
	}).GetError(); err != nil {
		return
	}
	global.Logger.Info(iTaskAddDTO.TaskName)
	err := m.Service.AddTask(&iTaskAddDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10011,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data: iTaskAddDTO,
	})
}
