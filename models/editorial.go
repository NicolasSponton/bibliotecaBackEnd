package models

type Editorial struct {
	Id        uint    `json:"id" gorm:"primaryKey"`
	Editorial string  `json:"editorial" gorm:"not null"`
	Libros    []Libro `json:"libros,omitempty"`
}

func (Editorial) TableName() string {
	return "Editoriales"
}
