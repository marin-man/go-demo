package dto

import "helloworld/model"

type UserDto struct {
	Username	string	`json:"name,omitempty"`
	Phone	string	`json:"phone,omitempty"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Username: user.Username,
		Phone: user.Phone,
	}
}
