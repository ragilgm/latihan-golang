package builder

import "github.com/ragilmaulana/Latihan/clean/pkg/usecase"

type PersonBuilder struct {}

func GetPrintPerson(in interface{}) (interface{}, error) {
	data := in.(*usecase.PersonRequest)
	out := &usecase.PersonRespond{
		Id_user:   data.Id_user,
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}
	return out, nil
}
