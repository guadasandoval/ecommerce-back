package services

import (
	"ecommerce/pkg/services/db/models"
	"gorm.io/gorm"
)


func CategoriasProductos(conn *gorm.DB) (*[]models.Categorias, error) {
	tablaCat := models.TABLA_CATEGORIAS

	var categoriasModel []models.Categorias

	result := conn.Table(tablaCat).Find(&categoriasModel)

	return &categoriasModel, result.Error
}


