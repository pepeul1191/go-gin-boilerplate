package routes

import "github.com/gin-gonic/gin"

func ProvinciaListar(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "listar provincias",
	})
}
