package models

import "time"

type Libro struct {
	Id                 int        `json:"id"`
	IdEditorial        int        `json:"ideditorial"`
	IdCategoria        int        `json:"idcategoria"`
	IdAutor            int        `json:"idautor"`
	Titulo             string     `json:"titulo" gorm:"not null"`
	Edicion            string     `json:"edicion"`
	FechaDePublicacion *time.Time `json:"fechaDePublicacion"`
	Copias             int        `json:"copias"`
	Editorial          Editorial  `json:"editorial,omitempty" gorm:"ForeignKey:IdEditorial"`
	Categoria          Categoria  `json:"categoria,omitempty" gorm:"ForeignKey:IdCategoria"`
	Autor              Autor      `json:"autor,omitempty" gorm:"ForeignKey:IdAutor"`
	CopiasPrestadas    int        `json:"copiasPrestadas,omitempty" gorm:"->;-:migration"`
}
