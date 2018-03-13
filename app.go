package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-pp/config"
	"github.com/gin-pp/routes"
)

func GetPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":    "pong",
		"BASE_URL":   config.BaseUrl,
		"STATIC_URL": config.StaticUrl,
	})
}

func main() {
	r := gin.Default()
	// configuraciones de vistas y archivos estáticos
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	// middleware personalizado
	r.Use(config.BeforeAll())
	// requests simple
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.GET("/pong", GetPong)
	// renderizado de vista
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	// rutas a otros arhcivos
	r.GET("/distrito/buscar", routes.DistritoBuscar)
	r.GET("/distrito/listar/:provincia_id", routes.DistritoListar)
	r.GET("/distrito/nombre/:distrito_id", routes.DistritoNombre)
	r.GET("/departamento/listar", routes.DepartamentoListar)
	r.GET("/provincia/listar/:departamento_id", routes.ProvinciaListar)
	// inicio de aplicación
	r.Run() // listen and serve on 0.0.0.0:8080
}
