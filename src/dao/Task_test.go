package dao

import (
	"CyberAssetMapper/src/service/dto"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("hello world")
	var taskDao = NewTaskDao()
	var iTaskAddDTO *dto.TaskAddDTO
	iTaskAddDTO.TaskName = "顺丰SRC"
	iTaskAddDTO.RootDomain = "sf-express.com"
	err := taskDao.AddTask(iTaskAddDTO)
	if err != nil {
		fmt.Println(err)
	}
}
