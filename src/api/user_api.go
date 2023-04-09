package api

import (
	"CyberAssetMapper/src/service"
	"CyberAssetMapper/src/service/dto"
	"github.com/gin-gonic/gin"
)

// 定义了一个UserApi的结构体
type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// @Tags 用户管理
// @Summary 用户登录
// @Description 用户详细描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(c *gin.Context) {

	var iUserLoginDTO dto.UserLoginDTO

	if err := m.BuildRequest(BuildRequestOption{
		Ctx: c,
		DTO: &iUserLoginDTO,
	}).GetError(); err != nil {
		return
	}

	//结合service了
	iUser, error := m.Service.Login()

	m.OK(ResponseJson{
		Data: iUserLoginDTO,
		Msg:  "Login Success",
	})

}
