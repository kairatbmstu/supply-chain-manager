package service

import (
	"example.com/m/v2/dto"
	"example.com/m/v2/errors"
	"example.com/m/v2/repository"
)

var RegistrationServiceInstance RegistrationService

type RegistrationService struct {
}

func (r RegistrationService) RegisterUser(regCommand dto.RegistrationCommand) (*dto.UserDTO, error) {

	userExist, err := repository.UserRepositoryInstance.FindByUsername(regCommand.Email)

	if err != nil {
		return nil, err
	}

	if userExist != nil {
		err = errors.NewError(errors.UserAlreadyExists)
		return nil, err
	}

	userExist, err = repository.UserRepositoryInstance.FindByEmail(regCommand.Email)

	if err != nil {
		return nil, err
	}

	if userExist != nil {
		err = errors.NewError(errors.UserAlreadyExists)
		return nil, err
	}

	var userDto = new(dto.UserDTO)

	return userDto, nil
}
