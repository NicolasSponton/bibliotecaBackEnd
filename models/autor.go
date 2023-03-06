package models

import "time"

type Autor struct {
	Id                uint       `json:"id" gorm:"primaryKey"`
	Nombre            string     `json:"nombre" gorm:"not null"`
	Apellido          string     `json:"apellido"`
	FechaDeNacimiento *time.Time `json:"fechaDeNacimiento"`
	FechaDeDefuncion  *time.Time `json:"fechaDeDefuncion"`
	Libros            []Libro    `json:"libros,omitempty"`
}

func (Autor) TableName() string {
	return "Autores"
}
