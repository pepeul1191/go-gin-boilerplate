package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-pp/config"
	"github.com/gin-pp/models"
)

func DepartamentoListar(c *gin.Context) {
	var departamentos []models.Departamento
	if err := config.Database().Find(&departamentos).Error; err != nil {
		fmt.Println(err)
	} else {
		c.JSON(200, departamentos)
	}
}
