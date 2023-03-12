package prestamos

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
	Prestamos     []models.Prestamo `json:"prestamos,omitempty"`
	Prestamo      *models.Prestamo  `json:"prestamo,omitempty"`
	TotalDataSize int64             `json:"totalDataSize,omitempty"`
	Estadisticas  []Estadisticas    `json:"estadisticas,omitempty"`
}

type Estadisticas struct {
	Mes      int
	Cantidad int
}

func GetAll(c echo.Context) error {
	db := database.GetConnection()
	prestamos := []models.Prestamo{}

	//Where
	where := "true"

	if c.QueryParam("pendientes") != "" {
		where += " and fecha_devolucion is null"
	}

	if c.QueryParam("q") != "" {
		where = where + " AND fecha_prestamo LIKE '%" + c.QueryParam("q") + "%'"
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

	if err := db.Where(where).Offset(offset).Order(sort).Limit(limit).Preload("Libro").Preload("Alumno").Find(&prestamos).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	if err := db.Model(&prestamos).Where(where).Count(&totalDataSize).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Prestamos: prestamos, TotalDataSize: totalDataSize}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})

}

func GetAllByMonth(c echo.Context) error {
	db := database.GetConnection()

	estadisticas := []Estadisticas{}

	//===============================================
	result := db.Raw("SELECT month(fecha_prestamo) AS mes, COUNT(*) AS cantidad  FROM prestamos WHERE Year(fecha_prestamo) = YEAR(CURRENT_DATE()) GROUP BY month(fecha_prestamo)").Find(&estadisticas)
	if result.Error != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Estadisticas: estadisticas}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})

}

func Get(c echo.Context) error {
	db := database.GetConnection()
	prestamo := models.Prestamo{}

	if err := db.Find(&prestamo).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Prestamo: &prestamo}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Delete(c echo.Context) error {
	db := database.GetConnection()
	id := c.Param("id")

	if err := db.Delete(models.Prestamo{}, id).Error; err != nil {
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

	prestamo := models.Prestamo{}
	if err := c.Bind(&prestamo); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(&prestamo).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Prestamo: &prestamo}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Update(c echo.Context) error {
	db := database.GetConnection()

	prestamo := new(models.Prestamo)
	if err := c.Bind(&prestamo); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Save(&prestamo).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := Data{Prestamo: prestamo}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}
