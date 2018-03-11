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
	r.GET("/provincia/listar", routes.ProvinciaListar)
	r.GET("/distrito/listar", routes.DistritoListar)
	r.Run() // listen and serve on 0.0.0.0:8080
}
