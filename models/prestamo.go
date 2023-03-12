package models

import "time"

type Prestamo struct {
	Id              int        `json:"id"`
	IdAlumno        int        `json:"idalumno"`
	IdLibro         int        `json:"idlibro"`
	FechaPrestamo   time.Time  `json:"fechaPrestamo"`
	FechaLimite     time.Time  `json:"fechaLimite" gorm:"not null"`
	FechaDevolucion *time.Time `json:"fechaDevolucion" gorm:"default: null"`
	Alumno          Alumno     `json:"alumno,omitempty" gorm:"ForeignKey:IdAlumno"`
	Libro           Libro      `json:"libro,omitempty" gorm:"ForeignKey:IdLibro"`
	Apellido        string     `json:"apellido,omitempty"`
}
