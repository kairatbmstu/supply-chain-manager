package service

import "example.com/m/v2/dto"

var RegistrationServiceInstance RegistrationService

type RegistrationService struct {
}

func (r RegistrationService) RegisterUser(registrationCommand dto.RegistrationCommand) (*dto.UserDTO, error) {
	return nil, nil
}
