package service

import (
	"golang-unit-test-pzn/entity"
	"golang-unit-test-pzn/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// program mock
	// wajib
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Dell XPS 16",
	}

	categoryRepository.Mock.On("FindById", "100").Return(category)

	categoryResult, err := categoryService.Get("100")
	// fmt.Println(category.Id, categoryResult.Id)
	assert.Nil(t, err)
	assert.NotNil(t, categoryResult)
	assert.Equal(t, category.Id, categoryResult.Id)
	assert.Equal(t, category.Name, categoryResult.Name)

}
