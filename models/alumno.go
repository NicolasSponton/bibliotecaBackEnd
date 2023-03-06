package models

type Alumno struct {
	Id        int     `json:"id"`
	IdCarrera int     `json:"idcarrera" gorm:"not null"`
	Nombre    string  `json:"nombre"`
	Apellido  string  `json:"apellido"`
	DNI       int     `json:"dni" gorm:"not null"`
	Celular   string  `json:"celular"`
	Email     string  `json:"email"`
	Carrera   Carrera `json:"carrera,omitempty" gorm:"ForeignKey:IdCarrera"`
}
