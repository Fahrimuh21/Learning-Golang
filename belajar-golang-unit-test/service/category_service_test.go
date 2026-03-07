package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}} //membuat objek categoryRepository dari struct CategoryRepositoryMock dengan inisialisasi field Mock dengan objek mock.Mock{}
var categoryService = CategoryService{Repository: categoryRepository}          //membuat objek categoryService dari struct CategoryService dengan inisialisasi field Repository dengan objek categoryRepository

func TestCategoryService_GetNotFound(t *testing.T) { //test untuk kasus data tidak ditemukan

	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil) //mengatur mock untuk method FindById dengan parameter "1" agar mengembalikan nil

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
