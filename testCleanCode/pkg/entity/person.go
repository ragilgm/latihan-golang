package entity

type Person struct {
	Id_Person int64 `json:"id_person" gorm:"column:id_person"`
	FirstName string `json:"firstname" gorm:"column:firstname"`
	LastName string `json:"lastname" gorm:"column:lastname"`

}

type Message struct {
	Status int64
}

func (b *Person) TableName() string {
	return "person"
}

