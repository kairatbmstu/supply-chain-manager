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
	user, err := repository.UserRepositoryInstance.Create(user)
	if err != nil {
		return nil, err
	}
	result := c.userMapper.ToDto(user)
	return result, nil
}

func (c UserService) GrantRoleToUser(user *dto.UserDTO, role dto.RoleDTO) (*dto.UserDTO, error) {
	return nil, nil
}

func (c UserService) Update(userdto *dto.UserDTO) (*dto.UserDTO, error) {
	user := c.userMapper.ToEntity(userdto)
	user, err := repository.UserRepositoryInstance.Update(user)
	if err != nil {
		return nil, err
	}
	result := c.userMapper.ToDto(user)
	return result, nil
}

func (c UserService) GetAll() ([]*dto.UserDTO, error) {
	return nil, nil
}

func (c UserService) FindByUsername(username string) (*dto.UserDTO, error) {
	user, err := repository.UserRepositoryInstance.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	result := c.userMapper.ToDto(user)
	return result, nil
}

func (c UserService) FindByEmail(email string) (*dto.UserDTO, error) {
	user, err := repository.UserRepositoryInstance.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	result := c.userMapper.ToDto(user)
	return result, nil
}
