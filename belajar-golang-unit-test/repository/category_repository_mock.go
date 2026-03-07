package repository

import (
	"belajar-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock // composition adalah memiliki objek dari struct lain dimana struct tersebut dapat mengakses semua field dan method dari struct yang dimilikinya disini memiliki objek mock dari testify/mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id) //memanggil method Called dari objek mock dengan parameter id, mengembalikan objek arguments yang berisi data hasil pemanggilan mock
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category //mengembalikan pointer ke entity Category karena method FindById mengembalikan pointer
	}
}
