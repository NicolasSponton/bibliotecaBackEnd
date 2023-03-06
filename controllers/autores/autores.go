package autores

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
	Autores       []models.Autor `json:"autores,omitempty"`
	Autor         *models.Autor  `json:"autor,omitempty"`
	TotalDataSize int64          `json:"totalDataSize,omitempty"`
}

func GetAll(c echo.Context) error {
	db := database.GetConnection()
	autores := []models.Autor{}

	//Where
	where := "true"

	if c.QueryParam("q") != "" {
		where = where + " AND apellido LIKE '%" + c.QueryParam("q") + "%'"
	}

	sort := "id ASC"
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

	if err := db.Where(where).Offset(offset).Order(sort).Limit(limit).Find(&autores).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	if err := db.Model(&autores).Where(where).Count(&totalDataSize).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Autores: autores, TotalDataSize: totalDataSize}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})

}

func Get(c echo.Context) error {
	db := database.GetConnection()
	autor := models.Autor{}

	if err := db.Find(&autor).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Autor: &autor}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Delete(c echo.Context) error {
	db := database.GetConnection()
	id := c.Param("id")

	if err := db.Delete(models.Autor{}, id).Error; err != nil {
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

	autor := models.Autor{}
	if err := c.Bind(&autor); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&autor).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Autor: &autor}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Update(c echo.Context) error {
	db := database.GetConnection()

	autor := new(models.Autor)
	if err := c.Bind(&autor); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Save(&autor).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Autor: autor}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
