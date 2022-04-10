package mapper

import (
	"example.com/m/v2/domain"
	"example.com/m/v2/dto"
)

type UserMapper struct {
}

func (u UserMapper) ToDto(user *domain.User) *dto.UserDTO {
	var result = new(dto.UserDTO)
	result.Email = user.Email
	result.Username = user.Username
	result.Password = user.Password
	return result
}

func (u UserMapper) ToEntity(user *dto.UserDTO) *domain.User {
	var result = new(domain.User)
	result.Email = user.Email
	result.Username = user.Username
	result.Password = user.Password
	return result
}
