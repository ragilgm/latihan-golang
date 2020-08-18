package repository

type PersonRepository interface {
	PrintPerson(data interface{})(interface{}, error)
}
