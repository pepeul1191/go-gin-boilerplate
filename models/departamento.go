package models

import "strconv"

type Departamento struct {
	ID     int    `gorm:"primary_key"`
	Nombre string `gorm:"column:nombre"`
}

func (Departamento) TableName() string {
	return "departamentos"
}

func (departamento *Departamento) GetId() int {
	return departamento.ID
}

func (departamento *Departamento) GetNombre() string {
	return departamento.Nombre
}

func (departamento *Departamento) ToString() string {
	return "Departamento : { id : " + strconv.Itoa(departamento.ID) +
		", nombre : " + departamento.Nombre + " }"
}
