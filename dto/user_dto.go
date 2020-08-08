package dto

import "ginEsseential/model"

type UserDto struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User_info) UserDto {
	return UserDto{
		Name: 	user.Name,
		Telephone: user.Telephone,
	}
}