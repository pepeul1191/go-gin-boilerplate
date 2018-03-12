package main

import "github.com/gin-gonic/gin"
import "github.com/gin-pp/routes"

func GetPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
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
