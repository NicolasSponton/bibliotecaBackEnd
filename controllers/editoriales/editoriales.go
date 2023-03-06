package editoriales

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
	Editoriales   []models.Editorial `json:"editoriales,omitempty"`
	Editorial     *models.Editorial  `json:"editorial,omitempty"`
	TotalDataSize int64              `json:"totalDataSize,omitempty"`
}

func GetAll(c echo.Context) error {
	db := database.GetConnection()
	editoriales := []models.Editorial{}

	//Where
	where := "true"

	if c.QueryParam("q") != "" {
		where = where + " AND editorial LIKE '%" + c.QueryParam("q") + "%'"
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

	if err := db.Where(where).Offset(offset).Order(sort).Limit(limit).Find(&editoriales).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	if err := db.Model(&editoriales).Where(where).Count(&totalDataSize).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Editoriales: editoriales, TotalDataSize: totalDataSize}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})

}

func Get(c echo.Context) error {
	db := database.GetConnection()
	editorial := models.Editorial{}

	if err := db.Find(&editorial).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Editorial: &editorial}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Delete(c echo.Context) error {
	db := database.GetConnection()
	id := c.Param("id")

	if err := db.Delete(models.Editorial{}, id).Error; err != nil {
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

	editorial := models.Editorial{}
	if err := c.Bind(&editorial); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&editorial).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Editorial: &editorial}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Update(c echo.Context) error {
	db := database.GetConnection()

	editorial := new(models.Editorial)
	if err := c.Bind(&editorial); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Save(&editorial).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Editorial: editorial}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
