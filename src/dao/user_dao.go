package dao

import (
	"CyberAssetMapper/src/global"
	"CyberAssetMapper/src/model"
	"CyberAssetMapper/src/service/dto"
	"fmt"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{NewBaseDao()}
	}
	return userDao
}

func (m *UserDao) GetUserByNameAndPassword(stUserName, stPassword string) model.User {
	var iUser model.User
	m.Orm.Model(&iUser).Where("username = ? AND password = ?", stUserName, stPassword).Find(&iUser)
	return iUser
}

func (m *UserDao) CheckUserExist(stUserName string) bool {
	var nTotal int64
	m.Orm.Model(&model.User{}).Where("name = ?", stUserName).Count(&nTotal)
	global.Logger.Info(fmt.Sprintf("CheckUserExist: %s, %d", stUserName, nTotal))
	return nTotal > 0
}

func (m *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	iUserAddDTO.ConvertToModel(&iUser)
	err := m.Orm.Save(&iUser).Error
	if err == nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}
	return err
}
