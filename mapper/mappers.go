package mapper

import (
	"example.com/m/v2/domain"
	"example.com/m/v2/dto"
)

type UserMapper struct {
}

func (u UserMapper) toDto(user *domain.User) *dto.UserDTO {
	return nil
}

func (u UserMapper) toEntity(user *dto.UserDTO) *domain.User {
	return nil
}

type RoleMapper struct {
}

func (u RoleMapper) toDto(role *domain.Role) *dto.RoleDTO {
	return nil
}

func (u RoleMapper) toEntity(role *dto.RoleDTO) *domain.Role {
	return nil
}

type OrganizationMapper struct {
}

func (u OrganizationMapper) toDto(entity *domain.Organization) *dto.OrganizationDTO {
	return nil
}

func (u OrganizationMapper) toEntity(dto *dto.OrganizationDTO) *domain.Organization {
	return nil
}

type FormOfLawMapper struct {
}

func (u FormOfLawMapper) toDto(entity *domain.FormOfLaw) *dto.FormOfLawDTO {
	return nil
}

func (u FormOfLawMapper) toEntity(dto *dto.FormOfLawDTO) *domain.FormOfLaw {
	return nil
}

type KBEMapper struct {
}

func (u KBEMapper) toDto(entity *domain.KBE) *dto.KBEDTO {
	return nil
}

func (u KBEMapper) toEntity(dto *dto.KBEDTO) *domain.KBE {
	return nil
}

type MeasurementMapper struct {
}

func (u MeasurementMapper) toDto(entity *domain.Measurement) *dto.MeasurementDTO {
	return nil
}

func (u MeasurementMapper) toEntity(dto *dto.MeasurementDTO) *domain.Measurement {
	return nil
}

type DepartmentMapper struct {
}

func (u DepartmentMapper) toDto(entity *domain.Department) *dto.DepartmentDTO {
	return nil
}

func (u DepartmentMapper) toEntity(dto *dto.DepartmentDTO) *domain.Department {
	return nil
}

type DeliveryConditionsMapper struct {
}

func (u DeliveryConditionsMapper) toDto(entity *domain.DeliveryConditions) *dto.DeliveryConditionsDTO {
	return nil
}

func (u DeliveryConditionsMapper) toEntity(dto *dto.DeliveryConditionsDTO) *domain.DeliveryConditions {
	return nil
}

type PackageTypeMapper struct {
}

func (u PackageTypeMapper) toDto(entity *domain.PackageType) *dto.PackageTypeDTO {
	return nil
}

func (u PackageTypeMapper) toEntity(dto *dto.PackageTypeDTO) *domain.PackageType {
	return nil
}
