package models

import "time"

type Autor struct {
	Id                int        `json:"id" gorm:"primaryKey"`
	Nombre            string     `json:"nombre"`
	Apellido          string     `json:"apellido"`
	FechaDeNacimiento *time.Time `json:"fechaDeNacimiento"`
	FechaDeDefuncion  *time.Time `json:"fechaDeDefuncion"`
	Libros            []Libro    `json:"libros,omitempty" gorm:"foreignKey:IdAutor"`
}

func (Autor) TableName() string {
	return "Autores"
}
