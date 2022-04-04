package service

import "example.com/m/v2/dto"

type CommercialConditionsService struct {
}

func (c CommercialConditionsService) GetOne() *dto.CommercialConditionsDTO {
	return nil
}

func (c CommercialConditionsService) Save(commercialConditionsDTO *dto.CommercialConditionsDTO) {

}

func (c CommercialConditionsService) Update(commercialConditionsDTO *dto.CommercialConditionsDTO) {

}

func (c CommercialConditionsService) Delete(commercialConditionsDTO *dto.CommercialConditionsDTO) {

}

func (c CommercialConditionsService) GetAll() []*dto.CommercialConditionsDTO {
	return make([]*dto.CommercialConditionsDTO, 0)
}
