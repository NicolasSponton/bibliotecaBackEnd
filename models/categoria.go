package models

type Categoria struct {
	Id        int     `json:"id" gorm:"primaryKey"`
	Categoria string  `json:"categoria"`
	Libros    []Libro `json:"libros,omitempty" gorm:"foreignKey:IdCategoria"`
}
