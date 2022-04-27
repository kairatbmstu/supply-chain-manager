package service

import (
	"example.com/m/v2/database"
	"example.com/m/v2/dto"
	"log"
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

func (OrganizationService) nextID() (int, error) {
	var id int
	query := "select nextval('s_organization')"
	rows, err := database.DB.Query(query, nil)

	if err != nil {
		log.Println("error  : ", err)
		return 0, err
	}
	if !rows.Next() {
		return 0, nil
	}

	err = rows.Scan(&id)
	if err != nil {
		log.Println("error  : ", err)
		return 0, err
	}

	return id, nil
}
