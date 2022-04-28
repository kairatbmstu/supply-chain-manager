package service

import (
	"example.com/m/v2/dto"
)

var OrganizationServiceInstance OrganizationService

type OrganizationService struct {
}

func (OrganizationService) Create(organizationDTO *dto.OrganizationDTO) (*dto.OrganizationDTO, error) {
	return nil, nil
}

func (OrganizationService) PartialUpdate(organizationDTO *dto.OrganizationDTO) (*dto.OrganizationDTO, error) {
	return nil, nil
}

func (OrganizationService) FindByGLN() (*dto.OrganizationDTO, error) {
	return nil, nil
}

func (OrganizationService) FindByIBAN() (*dto.OrganizationDTO, error) {
	return nil, nil
}
