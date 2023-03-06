package models

type Usuario struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
	Usuario  string `json:"usuario" gorm:"not null"`
	Clave    string `json:"clave" gorm:"not null"`
}
