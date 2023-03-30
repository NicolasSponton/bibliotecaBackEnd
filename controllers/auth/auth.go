package auth

import (
	"biblioteca/database"
	"biblioteca/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type ResponseMessage struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Login(c echo.Context) error {
	db := database.GetConnection()

	username := c.FormValue("username")
	password := c.FormValue("password")

	usuario := models.Usuario{}
	println(username)
	println(password)
	result := db.Where("usuario = ?", username).Where("clave = ?", password).Find(&usuario)
	if result.Error != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	if result.RowsAffected == 0 {
		response := ResponseMessage{
			Status:  "error",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusUnauthorized, response)
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	usuario.Clave = ""

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data: echo.Map{
			"token": t,
			"user":  usuario,
		},
	})
}
