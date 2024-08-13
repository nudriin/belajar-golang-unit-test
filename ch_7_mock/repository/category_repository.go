package repository

import "github.com/nudriin/belajar-golang-unit-test/ch_7_mock/entities"

type CategoryRepository interface {
	FindById(id string) *entities.Category
}
