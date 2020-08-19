package usecase

import (
	"fmt"
	"github.com/ragilmaulana/Latihan/clean/pkg/builder"
	"github.com/ragilmaulana/Latihan/clean/pkg/domain"

	er "github.com/ragilmaulana/Latihan/clean/error"
	//"github.com/ragilmaulana/Latihan/clean/pkg/entity"
	"github.com/ragilmaulana/Latihan/clean/pkg/model"
)

type PersonUsecase struct {
	personRepo model.PersonHandler
	out        builder.PersonBuilder
}

func NewPersonUsecase(a model.PersonHandler, out builder.PersonBuilder) *PersonUsecase {
	return &PersonUsecase{
		personRepo: a,
		out:        out,
	}
}

func (p *PersonUsecase) PrintPerson() (interface{}, error) {
	fmt.Println("called")
	response, err := p.personRepo.PrintPerson()
	if err != nil {
		return domain.Person{}, err
	}
	fmt.Println(response)
	return p.out.GetPrintPerson(response), nil
}

func (p *PersonUsecase) AddPerson(in interface{}) (int64, error) {
	if in == "" {
		return 0, fmt.Errorf(er.INVALID_ARGUMENT)
	}
	data, ok := in.(*domain.Person)
	if !ok {
		fmt.Println("data tidak valid")
		return 0, fmt.Errorf(er.INVALID_ARGUMENT)
	}
	response, err := p.personRepo.AddPerson(data)
	if err != nil {
		return 0, err
	}

	if response != 1 {
		return response, err
	} else {
		return response, nil
	}
}
func (p *PersonUsecase) DeletePerson(person interface{}) (interface{}, error) {
	data, ok := person.(*domain.Person)
	if !ok {
		return nil, fmt.Errorf(er.INVALID_ARGUMENT)
	}
	response, err := p.personRepo.DeletePerson(data)
	if err != nil {
		panic(err)
	}
	return response, nil
}
//
//func (p *PersonUsecase)EditPersonNById(interface{}) (interface{}, error){
//
//}

//func (p *PersonUsecase)EditPersonByName(interface{}) (interface{}, error){
//
//}
