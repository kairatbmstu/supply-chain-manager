package service

import (
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

	//repository.OrganizationRepositoryInstance.FindByIBAN()
	//
	//repository.OrganizationRepositoryInstance.Create()

	return &userDto, nil
}
