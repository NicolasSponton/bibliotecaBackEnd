package routes

import (
	"biblioteca/controllers/alumnos"
	"biblioteca/controllers/auth"
	"biblioteca/controllers/autores"
	"biblioteca/controllers/carreras"
	"biblioteca/controllers/categorias"
	"biblioteca/controllers/editoriales"
	"biblioteca/controllers/libros"
	"biblioteca/controllers/prestamos"
	"biblioteca/controllers/usuarios"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func HandleRoutes(e *echo.Echo) {

	r := e.Group("/biblioteca")

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))
	e.POST("/login", auth.Login)

	r.GET("/usuarios", usuarios.GetAll)
	r.GET("/usuarios/:id", usuarios.Get)
	r.POST("/usuarios", usuarios.Create)
	r.PUT("/usuarios", usuarios.Update)
	r.DELETE("/usuarios/:id", usuarios.Delete)

	r.GET("/carreras", carreras.GetAll)
	r.GET("/carreras/:id", carreras.Get)
	r.POST("/carreras", carreras.Create)
	r.PUT("/carreras", carreras.Update)
	r.DELETE("/carreras/:id", carreras.Delete)

	r.GET("/alumnos", alumnos.GetAll)
	r.GET("/alumnos/:id", alumnos.Get)
	r.POST("/alumnos", alumnos.Create)
	r.PUT("/alumnos", alumnos.Update)
	r.DELETE("/alumnos/:id", alumnos.Delete)

	r.GET("/prestamos", prestamos.GetAll)
	r.GET("/prestamos/getAllByMonth", prestamos.GetAllByMonth)
	r.GET("/prestamos/:id", prestamos.Get)
	r.POST("/prestamos", prestamos.Create)
	r.PUT("/prestamos", prestamos.Update)
	r.DELETE("/prestamos/:id", prestamos.Delete)

	r.GET("/libros", libros.GetAll)
	r.GET("/libros/:id", libros.Get)
	r.POST("/libros", libros.Create)
	r.PUT("/libros", libros.Update)
	r.DELETE("/libros/:id", libros.Delete)

	r.GET("/editoriales", editoriales.GetAll)
	r.GET("/editoriales/:id", editoriales.Get)
	r.POST("/editoriales", editoriales.Create)
	r.PUT("/editoriales", editoriales.Update)
	r.DELETE("/editoriales/:id", editoriales.Delete)

	r.GET("/categorias", categorias.GetAll)
	r.GET("/categorias/:id", categorias.Get)
	r.POST("/categorias", categorias.Create)
	r.PUT("/categorias", categorias.Update)
	r.DELETE("/categorias/:id", categorias.Delete)

	r.GET("/autores", autores.GetAll)
	r.GET("/autores/:id", autores.Get)
	r.POST("/autores", autores.Create)
	r.PUT("/autores", autores.Update)
	r.DELETE("/autores/:id", autores.Delete)

}
