package repository

import "example.com/m/v2/domain"

type ProductFieldStageRepository struct {
}

func (c ProductFieldStageRepository) GetOne() (*domain.ProductFieldStage, error) {
	return nil, nil
}

func (c ProductFieldStageRepository) Save(*domain.ProductFieldStage) (*domain.ProductFieldStage, error) {
	return nil, nil
}

func (c ProductFieldStageRepository) Update(*domain.ProductFieldStage) (*domain.ProductFieldStage, error) {
	return nil, nil
}

func (c ProductFieldStageRepository) GetAll() ([]*domain.ProductFieldStage, error) {
	return nil, nil
}
