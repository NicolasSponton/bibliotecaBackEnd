package models

type Categoria struct {
	Id        uint    `json:"id" gorm:"primaryKey"`
	Categoria string  `json:"categoria" gorm:"not null"`
	Libros    []Libro `json:"libros,omitempty"`
}
