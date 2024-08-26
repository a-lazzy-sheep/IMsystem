package models

import (
	"ginchat/utils"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      int 
	ClientID      string
	ClientPort    int
	LoginTime     uint64
	HeartbeatTime uint64
	LogoutTime    uint64
	IsLogout      bool
	DeviceInfo    string
}

func (u *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() ([]*UserBasic, error) {
	var userList []*UserBasic
	result := utils.DB.Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	return userList, nil
}
