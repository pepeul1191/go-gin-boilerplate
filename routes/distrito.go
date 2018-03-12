package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-pp/config"
	"github.com/gin-pp/models"
)

func DistritoListar(c *gin.Context) {
	provincia_id, err := strconv.ParseInt(c.Param("provincia_id"), 10, 64)
	if err != nil {
		c.String(500, "Error: El parámetro enviado no se pudo parsear a entero")
	} else {
		var distritos []models.Distrito
		if err := config.Database().Where("provincia_id = ?", provincia_id).Find(&distritos).Error; err != nil {
			fmt.Println(err)
		} else {
			c.JSON(200, gin.H{
				"rpta": distritos,
			})
		}
	}
}
