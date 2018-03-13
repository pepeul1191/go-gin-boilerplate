package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-pp/config"
	"github.com/gin-pp/routes"
)

func GetPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.Use(config.BeforeAll())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/pong", GetPong)
	r.GET("/departamento/listar", routes.DepartamentoListar)
	r.GET("/provincia/listar/:departamento_id", routes.ProvinciaListar)
	r.GET("/distrito/buscar", routes.DistritoBuscar)
	r.GET("/distrito/listar/:provincia_id", routes.DistritoListar)
	r.GET("/distrito/nombre/:distrito_id", routes.DistritoNombre)
	r.Run() // listen and serve on 0.0.0.0:8080
}
