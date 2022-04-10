package service

import (
	"example.com/m/v2/dto"
	"example.com/m/v2/mapper"
	"example.com/m/v2/repository"
)

var UserServiceInstance UserService

type UserService struct {
	userMapper mapper.UserMapper
}

func (c UserService) GetOne(id int) (*dto.UserDTO, error) {
	user, err := repository.UserRepositoryInstance.GetOne(id)
	if err != nil {
		return nil, err
	}
	userDto := c.userMapper.ToDto(user)
	return userDto, nil
}

func (c UserService) Create(userDto *dto.UserDTO) (*dto.UserDTO, error) {
	user := c.userMapper.ToEntity(userDto)
	return repository.UserRepositoryInstance.Create(user)
}

func (c UserService) GrantRoleToUser(user *dto.UserDTO, role dto.RoleDTO) (*dto.UserDTO, error) {
	return repository.UserRepositoryInstance.GrantRoleToUser(user, role)
}

func (c UserService) Update(user *dto.UserDTO) (*dto.UserDTO, error) {
	return repository.UserRepositoryInstance.Update(user)
}

func (c UserService) GetAll() ([]*dto.UserDTO, error) {
	return repository.UserRepositoryInstance.GetAll()
}

func (c UserService) FindByUsername(username string) (*dto.UserDTO, error) {
	return repository.UserRepositoryInstance.FindByUsername(username)
}

func (c UserService) FindByEmail(email string) (*dto.UserDTO, error) {
	return repository.UserRepositoryInstance.FindByEmail(email)
}
