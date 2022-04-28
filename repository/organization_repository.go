package repository

import (
	"example.com/m/v2/database"
	"example.com/m/v2/domain"
	"log"
)

var OrganizationRepositoryInstance OrganizationRepository

type OrganizationRepository struct {
}

func (OrganizationRepository) FindByGLN(gln string) (*domain.Organization, error) {
	return nil, nil
}

func (OrganizationRepository) FindByIBAN(iban string) (*domain.Organization, error) {
	return nil, nil
}

func (OrganizationRepository) Create(organization *domain.Organization) (*domain.Organization, error) {
	return nil, nil
}

func (OrganizationRepository) Update(organization *domain.Organization) (*domain.Organization, error) {
	return nil, nil
}

func (OrganizationRepository) Delete(organization *domain.Organization) error {
	return nil
}

func (c OrganizationRepository) nextID() (int, error) {
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
