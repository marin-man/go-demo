package dao

import (
	"helloworld/db"
	"helloworld/model"
)

func GetUserByPhone(phone string) (*model.User, error) {
	user := new(model.User)
	if err := db.DB.Debug().Where("phone=?", phone).First(&user).Error; nil != err {
		return nil, err
	}
	return user, nil
}

func IsPhoneExist(phone string) bool {
	var user model.User
	db.DB.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
