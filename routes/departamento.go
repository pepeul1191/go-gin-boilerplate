package routes

import "fmt"
import "github.com/gin-gonic/gin"

func DepartamentoListar(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "listar deptartamento",
	})
}

func Listar() {
	fmt.Println("departamento listar")
}
