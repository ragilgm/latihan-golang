package model

import (
	"github.com/ragilmaulana/Latihan/clean/pkg/entity"
)

type PersonModel struct {
	data []entity.Person
}

func (p PersonModel) PrintPerson(data interface{}) (interface{}, error) {
	request := data.(entity.Person)
	p.data = append(p.data, request)
	return data, nil
}
