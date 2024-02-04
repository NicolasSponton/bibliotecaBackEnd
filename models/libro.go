package models

import "time"

type Libro struct {
	Id                 int        `json:"id"`
	IdEditorial        int        `json:"ideditorial" gorm:"default:null"`
	IdCategoria        int        `json:"idcategoria" gorm:"default:null"`
	IdAutor            int        `json:"idautor" gorm:"default:null"`
	Titulo             string     `json:"titulo"`
	Edicion            string     `json:"edicion"`
	FechaDePublicacion *time.Time `json:"fechaDePublicacion"`
	Copias             int        `json:"copias"`
	Editorial          Editorial  `json:"editorial,omitempty" gorm:"ForeignKey:IdEditorial"`
	Categoria          Categoria  `json:"categoria,omitempty" gorm:"ForeignKey:IdCategoria"`
	Autor              Autor      `json:"autor,omitempty" gorm:"ForeignKey:IdAutor"`
	CopiasPrestadas    int        `json:"copiasPrestadas,omitempty" gorm:"->;-:migration"`
}
