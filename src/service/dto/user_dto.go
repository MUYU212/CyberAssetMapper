package dto

import "CyberAssetMapper/src/model"

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,first_is_a" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" binding:"required" message:"用户名填写错误"`
	Password string `json:"password,omitempty" binding:"required" message:"密码不能为空"`
}

func (m *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Name = m.Name
	iUser.Password = m.Password
}
