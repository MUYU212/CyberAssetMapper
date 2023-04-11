package api

import (
	"CyberAssetMapper/src/global"
	"CyberAssetMapper/src/service"
	"CyberAssetMapper/src/service/dto"
	"CyberAssetMapper/src/utils"
	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_ADD_USER      = 10011
	ERR_CODE_GET_USER_LIST = 10013
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
	iUser, err := m.Service.Login(iUserLoginDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}
	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)

	m.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
		Msg: "Login Success",
	})
}

func (m UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx: c,
		DTO: &iUserAddDTO,
	}).GetError(); err != nil {
		return
	}
	global.Logger.Info(iUserAddDTO.Name)
	err := m.Service.AddUser(&iUserAddDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data: iUserAddDTO,
	})
}

func (m *UserApi) GetUserList(c *gin.Context) {
	var iUserListDTO dto.UserListDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx: c,
		DTO: &iUserListDTO,
	}).GetError(); err != nil {
		return
	}
	iUserList, total, err := m.Service.GetUserList(&iUserListDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data:  iUserList,
		Total: total,
	})
}
