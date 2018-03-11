package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Departamento struct {
	ID     int    `gorm:"primary_key"`
	Nombre string `gorm:"column:nombre"`
}

func (Departamento) TableName() string {
	return "departamentos"
}

func main() {
	db, err := gorm.Open("sqlite3", "db/ubicaciones.db")
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	//db.Create(&Rol{Nombre: "Rol ORM", Sistema_id: 8})
	var departamentos []Departamento
	if err := db.Find(&departamentos).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(departamentos)
	}

	db.Close()
}
