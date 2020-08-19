package domain

type Person struct {
	Id_User int64
	FirstName string
	LastName string
}

type PersonInputPort interface {
	PrintPerson()(Person, error)
	AddPerson(interface{}) (int64, error)
	DeletePerson(interface{}) (int64, error)
	EditPersonNById(interface{}) (interface{}, error)
	EditPersonByName(interface{}) (interface{}, error)
	PrintSinglePerson(interface{})(interface{}, error)
}

type PersonOutputPort interface {
	GetPrintPerson(interface{})(Person, error)
	GetEditPersonNById(interface{}) (interface{}, error)
	GetEditPersonByName(interface{}) (interface{}, error)
}