package models

type Carrera struct {
	Id      uint     `json:"id" gorm:"primaryKey"`
	Carrera string   `json:"carrera" gorm:"not null"`
	Alumnos []Alumno `json:"alumnos,omitempty"`
}
