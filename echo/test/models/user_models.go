package models

type User struct {
	Id int `json:"id"`
	NamaDepan string `json:"namadepan"`
	NamaBelakang string `json:"namabelakang"`
	Email string `json:"email"`
}

type Users []User