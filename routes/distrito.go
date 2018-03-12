package routes

import "github.com/gin-gonic/gin"

func DistritoListar(c *gin.Context) {
	c.Param("provincia_id")
	c.JSON(200, gin.H{
		"message": "listar distritos",
	})
}
