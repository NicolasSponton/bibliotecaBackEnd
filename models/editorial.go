package models

type Editorial struct {
	Id        int     `json:"id" gorm:"primaryKey"`
	Editorial string  `json:"editorial"`
	Libros    []Libro `json:"libros,omitempty" gorm:"foreignKey:IdEditorial"`
}

func (Editorial) TableName() string {
	return "Editoriales"
}
