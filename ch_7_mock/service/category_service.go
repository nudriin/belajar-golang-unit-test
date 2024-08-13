package service

import (
	"errors"

	"github.com/nudriin/belajar-golang-unit-test/ch_7_mock/entities"
	"github.com/nudriin/belajar-golang-unit-test/ch_7_mock/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entities.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
