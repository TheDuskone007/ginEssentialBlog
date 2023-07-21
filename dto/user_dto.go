package dto

import "ginEssential2/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

//ToUserDto 将user转换为UserDto
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
