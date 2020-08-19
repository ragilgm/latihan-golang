package builder

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/ragilmaulana/Latihan/clean/pkg/domain"
)

type PersonBuilder struct{}

func (*PersonBuilder) GetPrintPerson(in interface{}) interface{} {
	req := in.(*[]domain.Person)

	var out *[]domain.Person

	err1 := mapstructure.Decode(req, &out)

	if err1 != nil {
		fmt.Println("data tidak valid atau tidak sama (builder)")
	}
	return out
}
