package libros

import (
	"biblioteca/database"
	"biblioteca/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ResponseMessage struct {
	Status  string `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	Libros        []models.Libro `json:"libros,omitempty"`
	Libro         *models.Libro  `json:"libro,omitempty"`
	TotalDataSize int64          `json:"totalDataSize,omitempty"`
}

func GetAll(c echo.Context) error {
	db := database.GetConnection()
	libros := []models.Libro{}

	//Where
	where := "true"

	if c.QueryParam("q") != "" {
		where = where + " AND titulo LIKE '%" + c.QueryParam("q") + "%'"
	}

	sort := "id DESC"
	if c.QueryParam("sort") != "" {
		sort = c.QueryParam("sort")
	}

	var page int = 1
	var limit int = 1000
	var offset int = 0
	var totalDataSize int64 = 0

	if c.QueryParam("limit") != "" {
		limit, _ = strconv.Atoi(c.QueryParam("limit"))
	}

	if c.QueryParam("page") != "" {
		page, _ = strconv.Atoi(c.QueryParam("page"))
	}

	offset = limit * (page - 1)

	//===============================================

	if err := db.Select("libros.*, COUNT(prestamos.id) as CopiasPrestadas").Joins("left JOIN prestamos ON (prestamos.id_libro = libros.id AND prestamos.fecha_devolucion IS null)").Where(where).Group("libros.id").Offset(offset).Order(sort).Limit(limit).Preload("Autor").Preload("Editorial").Preload("Categoria").Find(&libros).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	if err := db.Model(&libros).Where(where).Count(&totalDataSize).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Libros: libros, TotalDataSize: totalDataSize}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})

}

func Get(c echo.Context) error {
	db := database.GetConnection()
	libro := models.Libro{}

	if err := db.Find(&libro).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Libro: &libro}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Delete(c echo.Context) error {
	db := database.GetConnection()
	id := c.Param("id")

	if err := db.Delete(models.Libro{}, id).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
	})
}

func Create(c echo.Context) error {
	db := database.GetConnection()

	libro := models.Libro{}
	if err := c.Bind(&libro); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&libro).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Libro: &libro}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Update(c echo.Context) error {
	db := database.GetConnection()

	libro := new(models.Libro)
	if err := c.Bind(&libro); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Save(&libro).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Libro: libro}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
