package utils

import (
	"biblioteca/database"
	"biblioteca/models"
)

func MigrateSchemas() {

	db := database.GetConnection()
	db.AutoMigrate(&models.Carrera{})
	db.AutoMigrate(&models.Alumno{})
	db.AutoMigrate(&models.Prestamo{})
	db.AutoMigrate(&models.Usuario{})
	db.AutoMigrate(&models.Libro{})
	db.AutoMigrate(&models.Categoria{})
	db.AutoMigrate(&models.Autor{})
	db.AutoMigrate(&models.Editorial{})

}
