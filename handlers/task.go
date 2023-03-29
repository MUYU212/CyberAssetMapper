package handlers

import (
	"CyberAssetMapper/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type TaskResponse struct {
	ID        uint      `json:"id"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"createdAt"`
}

// 根据id查询任务信息
func GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	var task model.Task
	id, err := strconv.Atoi(taskID)
	task = model.Get_Task(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// 插入数据库
func CreateTask(c *gin.Context) {
	var task model.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := model.Insert_Task(task.Domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	taskResponse := TaskResponse{
		ID:        result.ID,
		Domain:    result.Domain,
		CreatedAt: result.CreatedAt,
	}
	c.JSON(http.StatusOK, taskResponse)
}
