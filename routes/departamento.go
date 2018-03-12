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
		for i := 0; i < len(departamentos); i++ {
			//fmt.Println(departamentos[i].ToString())
			fmt.Println(departamentos[i].Nombre)
		}
		//fmt.Fprintf(os.Stdout, "%s", departamentos)
		//fmt.Println((json.Marshal(departamentos)))
		//c.String(200, string(departamentos))
		c.JSON(200, gin.H{
			"rpta": departamentos,
		})

	}
}
