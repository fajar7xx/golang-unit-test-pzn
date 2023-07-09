package repository

import (
	"golang-unit-test-pzn/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

// kode ini tidak memanggil database melainkan mock object
func (repository *CategoryRepositoryMock) FindById(Id string) *entity.Category {
	arguments := repository.Mock.Called(Id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		// karena kembalianny interface maka di konversi menjadi kategory
		category := arguments.Get(0).(entity.Category)
		// karena category pointer maka tambahkan &
		return &category
	}
}
