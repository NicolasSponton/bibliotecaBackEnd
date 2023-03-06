package usuarios

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
	Usuarios      []models.Usuario `json:"usuarios,omitempty"`
	Usuario       *models.Usuario  `json:"usuario,omitempty"`
	TotalDataSize int64            `json:"totalDataSize,omitempty"`
}

func GetAll(c echo.Context) error {
	db := database.GetConnection()
	usuarios := []models.Usuario{}

	//Where
	where := "true"

	if c.QueryParam("q") != "" {
		where = where + " AND apellido LIKE '%" + c.QueryParam("q") + "%'"
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

	if err := db.Where(where).Offset(offset).Order(sort).Limit(limit).Find(&usuarios).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	if err := db.Model(&usuarios).Where(where).Count(&totalDataSize).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Usuarios: usuarios, TotalDataSize: totalDataSize}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})

}

func Get(c echo.Context) error {
	db := database.GetConnection()
	usuario := models.Usuario{}

	if err := db.Find(&usuario).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Usuario: &usuario}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Delete(c echo.Context) error {
	db := database.GetConnection()
	id := c.Param("id")

	if err := db.Delete(models.Usuario{}, id).Error; err != nil {
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

	usuario := models.Usuario{}
	if err := c.Bind(&usuario); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&usuario).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Usuario: &usuario}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Update(c echo.Context) error {
	db := database.GetConnection()

	usuario := new(models.Usuario)
	if err := c.Bind(&usuario); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Save(&usuario).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Usuario: usuario}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
