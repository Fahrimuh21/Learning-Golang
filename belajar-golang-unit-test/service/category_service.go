package service

import (
	"belajar-golang-unit-test/entity"     //ambil entity Category
	"belajar-golang-unit-test/repository" //ambil interface dari repository package
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository //ambil interface dari repository package
}

func (service CategoryService) Get(id string) (*entity.Category, error) { //mengembalikan pointer ke entity Category dan error, entity Category diambil dari package entity
	category := service.Repository.FindById(id)
	if category == nil { // jika category nil
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
