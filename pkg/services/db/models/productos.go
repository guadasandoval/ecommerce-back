package models

import "gopkg.in/guregu/null.v3"

const TABLA_PRODUCTO = "productos"
const TABLA_PRODUCTO_METADATA = "productos_metadata"

type Producto struct {
	Id          int    `gorm:"column:Id"`
	Nombre      string `gorm:"column:Nombre"`
	Descripcion string `gorm:"column:Descripcion"`
	Stock       bool   `gorm:"column:Stock"`
	Precio      int    `gorm:"column:Precio"`
	IDCategoria int    `gorm:"column:IdCategoria"`
}

type ProductoMetadata struct {
	ID             int         `gorm:"column:Id"`
	IDProducto     null.Int    `gorm:"column:IdProducto"`
	NombreArchivo  string      `gorm:"column:NombreArchivo"`
	LineaArchivo   int         `gorm:"column:LineaArchivo"`
	VersionArchivo string      `gorm:"column:VersionArchivo"`
	OcurrioError   bool        `gorm:"column:OcurrioError"`
	Error          null.String `gorm:"column:Error"`
}
