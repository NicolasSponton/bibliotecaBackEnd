package models

type Alumno struct {
	Id        int     `json:"id" gorm:"primaryKey"`
	IdCarrera int     `json:"idcarrera" gorm:"default:null"`
	Nombre    string  `json:"nombre"`
	Apellido  string  `json:"apellido"`
	DNI       int     `json:"dni"`
	Celular   string  `json:"celular"`
	Email     string  `json:"email"`
	Carrera   Carrera `json:"carrera,omitempty" gorm:"foreignkey:IdCarrera"`
}
