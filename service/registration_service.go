package service

import (
	"database/sql"
	"example.com/m/v2/database"
	"example.com/m/v2/domain"
	"example.com/m/v2/dto"
	"example.com/m/v2/errors"
	"example.com/m/v2/repository"
	"log"
)

var RegistrationServiceInstance RegistrationService

type RegistrationService struct {
}

func (r RegistrationService) RegisterUser(regCommand dto.RegistrationDTO) (*dto.UserDTO, error) {

	userExist, err := repository.UserRepositoryInstance.FindByUsername(regCommand.Email)

	if err != nil {
		log.Println("err : ", err)
		return nil, err
	}

	if userExist != nil {
		err = errors.NewError(errors.UserAlreadyExists)
		log.Println("err : ", err)
		return nil, err
	}

	userExist, err = repository.UserRepositoryInstance.FindByEmail(regCommand.Email)

	if err != nil {
		log.Println("err : ", err)
		return nil, err
	}

	if userExist != nil {
		log.Println("err : ", err)
		err = errors.NewError(errors.UserAlreadyExists)
		return nil, err
	}

	var userDto = dto.UserDTO{
		Email:    regCommand.Email,
		Username: regCommand.Email,
		Password: regCommand.Password,
	}

	var organizationDto = dto.OrganizationDTO{
		Bin:        regCommand.IINBIN,
		NameRu:     regCommand.OrganizationName,
		NameKz:     regCommand.OrganizationName,
		NameEn:     regCommand.OrganizationName,
		IsResident: regCommand.IsResident,
		Blocked:    false,
		CorpPhone:  regCommand.OrganizationPhone,
		ContactFio: regCommand.ContactPersonFullname,
		Website:    regCommand.OrganizationSite,
		Gln:        regCommand.GLN,
		IsStm:      regCommand.IsSTM,
		IsNds:      regCommand.IsVAT,
		IBAN:       "",
		Bic:        regCommand.BIC,
		KbeCode:    regCommand.KBE,
		Address:    regCommand.OrgAddress,
		FormOfLaw:  dto.FormOfLawDTO{},
		KBE:        dto.KBEDTO{},
	}

	var organizationExist *domain.Organization

	organizationExist, err = repository.OrganizationRepositoryInstance.FindByBIN(organizationDto.Bin)
	if err != nil {
		log.Println("err : ", err)
		err = errors.NewError(errors.InternalServerError)
		return nil, err
	}
	organizationExist, err = repository.OrganizationRepositoryInstance.FindByIBAN(organizationDto.IBAN)
	if err != nil {
		log.Println("err : ", err)
		err = errors.NewError(errors.InternalServerError)
		return nil, err
	}
	organizationExist, err = repository.OrganizationRepositoryInstance.FindByGLN(organizationDto.Gln)
	if err != nil {
		log.Println("err : ", err)
		err = errors.NewError(errors.InternalServerError)
		return nil, err
	}

	if organizationExist != nil {
		log.Println("err : ", err)
		err = errors.NewError(errors.OrganizationAlreadyExists)
		return nil, err
	}

	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("err : ", err)
		err = errors.NewError(errors.InternalServerError)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("err : ", err)
		err = errors.NewError(errors.InternalServerError)
		return nil, err
	}
	return &userDto, nil
}

func (r RegistrationService) CreateOrganization(organization dto.OrganizationDTO, tx *sql.Tx) (*dto.OrganizationDTO, error) {
	return nil, nil
}

func (r RegistrationService) CreateUser(user dto.UserDTO, tx *sql.Tx) (*dto.UserDTO, error) {
	return nil, nil
}
