package models

type Usuario struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
	Usuario  string `json:"usuario"`
	Clave    string `json:"clave"`
}
