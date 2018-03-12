package routes

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-pp/config"
	"github.com/gin-pp/models"
)

func ProvinciaListar(c *gin.Context) {
	departamento_id, err := strconv.ParseInt(c.Param("departamento_id"), 10, 64)
	if err != nil {
		c.String(500, "Error: El parámetro enviado no se pudo parsear a entero")
	} else {
		var provincias []models.Provincia
		if err := config.Database().Where("departamento_id = ?", departamento_id).Find(&provincias).Error; err != nil {
			fmt.Println(err)
		} else {
			c.JSON(200, gin.H{
				"rpta": provincias,
			})
		}
	}
}
