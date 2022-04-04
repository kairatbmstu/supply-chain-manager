package mapper

import (
	"example.com/m/v2/domain"
	"example.com/m/v2/dto"
)

type UserMapper struct {
}

func (u UserMapper) toDto(user *domain.User) *dto.UserDTO {
	var result = new(dto.UserDTO)
	result.Email = user.Email
	result.Username = user.Username
	result.Password = user.Password
	return result
}

func (u UserMapper) toEntity(user *dto.UserDTO) *domain.User {
	var result = new(domain.User)
	result.Email = user.Email
	result.Username = user.Username
	result.Password = user.Password
	return result
}

type RoleMapper struct {
}

func (u RoleMapper) toDto(role *domain.Role) *dto.RoleDTO {
	var result = new(dto.RoleDTO)
	result.Code = role.Code
	result.NameRu = role.NameRu
	result.NameEn = role.NameEn
	result.NameKz = role.NameKz
	return result
}

func (u RoleMapper) toEntity(role *dto.RoleDTO) *domain.Role {
	var result = new(domain.Role)
	result.Code = role.Code
	result.NameEn = role.NameEn
	result.NameKz = role.NameKz
	result.NameRu = role.NameRu
	return result
}

type OrganizationMapper struct {
}

func (u OrganizationMapper) toDto(entity *domain.Organization) *dto.OrganizationDTO {
	var result = dto.OrganizationDTO{
		ID:         int64(entity.ID),
		ExternalId: int64(entity.ExternalId),
		Bin:        entity.Bin,
		NameRu:     entity.NameRu,
		NameKz:     entity.NameKz,
		NameEn:     entity.NameEn,
		IsResident: entity.IsResident,
		Blocked:    entity.Blocked,
		CorpPhone:  entity.CorpPhone,
		ContactFio: entity.ContactFio,
		Website:    entity.Website,
		Gln:        entity.Gln,
		IsStm:      entity.IsStm,
		IsNds:      entity.IsNds,
		IBAN:       entity.Iban,
		Bic:        entity.Bic,
		KbeCode:    entity.KBE.Code,
		Address:    entity.Address,
		KBE:        dto.KBEDTO(entity.KBE),
	}
	result.FormOfLaw = dto.FormOfLawDTO{
		ID:     entity.FormOfLaw.ID,
		NameKz: entity.FormOfLaw.NameEn,
		NameEn: entity.FormOfLaw.NameRu,
		NameRu: entity.FormOfLaw.NameKz,
	}

	return &result
}

func (u OrganizationMapper) toEntity(dto *dto.OrganizationDTO) *domain.Organization {
	return nil
}

type FormOfLawMapper struct {
}

func (u FormOfLawMapper) toDto(entity *domain.FormOfLaw) *dto.FormOfLawDTO {
	var result = new(dto.FormOfLawDTO)
	result.ID = entity.ID
	result.NameRu = entity.NameRu
	result.NameEn = entity.NameEn
	result.NameKz = entity.NameKz
	return result
}

func (u FormOfLawMapper) toEntity(dto *dto.FormOfLawDTO) *domain.FormOfLaw {
	var result = new(domain.FormOfLaw)
	result.ID = dto.ID
	result.NameEn = dto.NameEn
	result.NameKz = dto.NameKz
	result.NameRu = dto.NameRu
	return result
}

type KBEMapper struct {
}

func (u KBEMapper) toDto(entity *domain.KBE) *dto.KBEDTO {
	var result = new(dto.KBEDTO)
	result.ID = entity.ID
	result.NameRu = entity.NameRu
	result.NameEn = entity.NameEn
	result.NameKz = entity.NameKz
	return result
}

func (u KBEMapper) toEntity(dto *dto.KBEDTO) *domain.KBE {
	var result = new(domain.KBE)
	result.ID = dto.ID
	result.NameEn = dto.NameEn
	result.NameKz = dto.NameKz
	result.NameRu = dto.NameRu
	return result
}

type MeasurementMapper struct {
}

func (u MeasurementMapper) toDto(entity *domain.Measurement) *dto.MeasurementDTO {
	var result = new(dto.MeasurementDTO)
	result.ID = entity.ID
	result.NameRu = entity.NameRu
	result.NameEn = entity.NameEn
	result.NameKz = entity.NameKz
	return result
}

func (u MeasurementMapper) toEntity(dto *dto.MeasurementDTO) *domain.Measurement {
	var result = new(domain.Measurement)
	result.ID = dto.ID
	result.NameEn = dto.NameEn
	result.NameKz = dto.NameKz
	result.NameRu = dto.NameRu
	return result
}

type DepartmentMapper struct {
}

func (u DepartmentMapper) toDto(entity *domain.Department) *dto.DepartmentDTO {
	var result = new(dto.DepartmentDTO)
	result.ID = entity.ID
	result.NameRu = entity.NameRu
	result.NameEn = entity.NameEn
	result.NameKz = entity.NameKz
	return result
}

func (u DepartmentMapper) toEntity(dto *dto.DepartmentDTO) *domain.Department {
	var result = new(domain.Department)
	result.ID = dto.ID
	result.NameEn = dto.NameEn
	result.NameKz = dto.NameKz
	result.NameRu = dto.NameRu
	return result
}

type DeliveryConditionsMapper struct {
}

func (u DeliveryConditionsMapper) toDto(entity *domain.DeliveryConditions) *dto.DeliveryConditionsDTO {
	var result = new(dto.DeliveryConditionsDTO)
	result.ID = entity.ID
	result.NameRu = entity.NameRu
	result.NameEn = entity.NameEn
	result.NameKz = entity.NameKz
	return result
}

func (u DeliveryConditionsMapper) toEntity(dto *dto.DeliveryConditionsDTO) *domain.DeliveryConditions {
	var result = new(domain.DeliveryConditions)
	result.ID = dto.ID
	result.NameEn = dto.NameEn
	result.NameKz = dto.NameKz
	result.NameRu = dto.NameRu
	return result
}

type PackageTypeMapper struct {
}

func (u PackageTypeMapper) toDto(entity *domain.PackageType) *dto.PackageTypeDTO {
	var result = new(dto.PackageTypeDTO)
	result.ID = entity.ID
	result.NameRu = entity.NameRu
	result.NameEn = entity.NameEn
	result.NameKz = entity.NameKz
	return result
}

func (u PackageTypeMapper) toEntity(dto *dto.PackageTypeDTO) *domain.PackageType {
	var result = new(domain.PackageType)
	result.ID = dto.ID
	result.NameEn = dto.NameEn
	result.NameKz = dto.NameKz
	result.NameRu = dto.NameRu
	return result
}
