package model

import (
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"github.com/ragilmaulana/Latihan/clean/config"
	"github.com/ragilmaulana/Latihan/clean/pkg/entity"
	//er "github.com/ragilmaulana/Latihan/clean/error"
	"github.com/ragilmaulana/Latihan/clean/pkg/domain"
)

type PersonHandler struct {
	Db     *gorm.DB
	config config.DbConfig
}

func (p *PersonHandler) PrintPerson() (interface{}, error) {
	var person []entity.Person

	// connect to databse
	connect, err := p.config.Connect()
	if err != nil {
		return nil, err
	}

	// query to databse
	err2 := connect.Find(&person).Error
	if err2 != nil {
		return nil, err2
	}

	// close konneksi database
	connect.Close()

	var respon *[]domain.Person

	// decode object respon dari database
	err = mapstructure.Decode(person, &respon)
	if err != nil {
		return nil, nil
	}
	return respon, nil
}

func (p *PersonHandler) AddPerson(request interface{}) (int64, error) {
	data := request.(*domain.Person)

	// insert ke database, id tidak dimasukan karna sudah auto increment
	insert := entity.Person{
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	// connect to databse
	connect, err := p.config.Connect()
	if err != nil {
		return 0, err
	}

	// query to databse
	rows := connect.Create(&insert).RowsAffected

	return rows, nil
}
func (p *PersonHandler) DeletePerson(person interface{}) (int64, error) {
	data := person.(*domain.Person)
	id := entity.Person{
		Id_Person: data.Id_User,
	}
	connect, err := p.config.Connect()
	if err != nil {
		return 0, err
	}
	rows := connect.Where("id_person = ?", id.Id_Person).Delete(&id).RowsAffected
	return rows, nil
}

//func (p *PersonHandler) EditPersonNById(interface{}) (interface{}, error)
//func (p *PersonHandler) EditPersonByName(interface{}) (interface{}, error)
