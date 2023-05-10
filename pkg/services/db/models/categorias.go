package models

const TABLA_CATEGORIAS = "categorias"

type Categorias struct {
	Id          int    `gorm:"column:Id"`
	Nombre      string `gorm:"column:Nombre"`
}