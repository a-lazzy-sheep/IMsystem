package models

import (
	"ginchat/utils"
	// "github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

// All fields are required to at least have one validation defined
// using valid:"-" or valid:"email,optional"
// func init() {
//   valid.SetFieldsRequiredByDefault(true)
// }

type UserBasic struct {
	gorm.Model
	Name          string `valid:"-"`
	Password      string `valid:"-"`
	Phone         string `valid:"numeric"`
	Email         string `valid:"required,email"`
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

func CreateUser(user *UserBasic) error {
	result := utils.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(user *UserBasic) error {
	result := utils.DB.Where("name = ?",user.Name).Delete(&UserBasic{}) 
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateUser(user *UserBasic) error {
	result := utils.DB.Model(&UserBasic{})
	result = result.Where("id = ?",user.ID)
	result = result.Updates(UserBasic{
		Name:          user.Name,
		Password:      user.Password,
		Phone:         user.Phone,
		Email:         user.Email,
		// Identity:      user.Identity,
		// ClientID:      user.ClientID,
		// ClientPort:    user.ClientPort,
		// LoginTime:     user.LoginTime,
		// HeartbeatTime: user.HeartbeatTime,
		// LogoutTime:    user.LogoutTime,
		// IsLogout:      user.IsLogout,
		// DeviceInfo:    user.DeviceInfo,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindUserByName(name string) (*UserBasic, error) {
	var user UserBasic
	result := utils.DB.Where("name = ?",name).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func FindUserByEmail(email string) (*UserBasic, error) {
	var user UserBasic
	result := utils.DB.Where("email = ?",email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

