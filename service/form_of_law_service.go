package service

import (
	"example.com/m/v2/dto"
	"example.com/m/v2/repository"
)

var FormOfLawServiceInstance FormOfLawService

type FormOfLawService struct {
}

func (c FormOfLawService) FindResident() ([]dto.FormOfLawDTO, error) {
	var result = make([]dto.FormOfLawDTO, 0)

	formoflaws, err := repository.FormOfLawRepositoryInstance.FindResident()
	if err != nil {
		return nil, err
	}

	for _, fol := range formoflaws {
		result = append(result, dto.FormOfLawDTO{
			ID:         fol.ID,
			Code:       fol.Code,
			IsResident: fol.IsResident,
			NameRu:     fol.NameRu,
			NameKz:     fol.NameKz,
			NameEn:     fol.NameEn,
		})
	}

	return result, nil
}

func (c FormOfLawService) FindNonResident() ([]dto.FormOfLawDTO, error) {
	var result = make([]dto.FormOfLawDTO, 0)

	formoflaws, err := repository.FormOfLawRepositoryInstance.FindNonResident()
	if err != nil {
		return nil, err
	}

	for _, fol := range formoflaws {
		result = append(result, dto.FormOfLawDTO{
			ID:         fol.ID,
			Code:       fol.Code,
			IsResident: fol.IsResident,
			NameRu:     fol.NameRu,
			NameKz:     fol.NameKz,
			NameEn:     fol.NameEn,
		})
	}

	return result, nil
}

func (c FormOfLawService) FindAll() ([]dto.FormOfLawDTO, error) {
	var result = make([]dto.FormOfLawDTO, 0)

	formoflaws, err := repository.FormOfLawRepositoryInstance.FindAll()
	if err != nil {
		return nil, err
	}

	for _, fol := range formoflaws {
		result = append(result, dto.FormOfLawDTO{
			ID:         fol.ID,
			Code:       fol.Code,
			IsResident: fol.IsResident,
			NameRu:     fol.NameRu,
			NameKz:     fol.NameKz,
			NameEn:     fol.NameEn,
		})
	}

	return result, nil
}
