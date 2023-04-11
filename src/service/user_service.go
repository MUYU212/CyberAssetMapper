package service

import (
	"CyberAssetMapper/src/dao"
	"CyberAssetMapper/src/model"
	"CyberAssetMapper/src/service/dto"
	"errors"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (m *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, error) {
	var errResult error
	iUser := m.Dao.GetUserByNameAndPassword(iUserDTO.Name, iUserDTO.Password)
	if iUser.ID == 0 {
		errResult = errors.New("Invalid User Or Password")
	}
	return iUser, errResult
}

func (m *UserService) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	if m.Dao.CheckUserExist(iUserAddDTO.Name) {
		return errors.New("User Already Exist")
	}
	return m.Dao.AddUser(iUserAddDTO)
}
func (m *UserService) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return m.Dao.GetUserList(iUserListDTO)
}
