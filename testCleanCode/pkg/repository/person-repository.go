package repository

type Person struct {
	Id_User   int64
	FirstName string
	LastName  string
}

type PersonRepository interface {
	PrintPerson() (interface{}, error)
	AddPerson(interface{}) (interface{}, error)
	DeletePerson(interface{}) (interface{}, error)
	EditPersonNById(interface{}) (interface{}, error)
	EditPersonByName(interface{}) (interface{}, error)
	PrintSinglePerson(interface{})(interface{},error)
}
