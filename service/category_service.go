package service

import (
	"errors"
	"golang-unit-test-pzn/entity"
	"golang-unit-test-pzn/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}