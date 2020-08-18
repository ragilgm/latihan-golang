package usecase

import (
	"fmt"
	"github.com/ragilmaulana/Latihan/clean/pkg/entity"
	"github.com/ragilmaulana/Latihan/clean/pkg/model"
	"github.com/ragilmaulana/Latihan/clean/pkg/repository"
)

type PersonUsecase struct {
	repo repository.PersonRepository
	model model.PersonModel
}

func (p *PersonUsecase) PrintPerson(data interface{}) (interface{}, error) {
	fmt.Println("called")
	res := data.(entity.Person)
	respons, err := p.model.PrintPerson(res)
	fmt.Println("responds",respons)
	if err != nil {
		println(err)
	}
	return respons, nil
}
