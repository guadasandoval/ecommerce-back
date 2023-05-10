package services

import (
	"ecommerce/pkg/services/db/models"
	"errors"

	"gorm.io/gorm"
)


func InsertarProductoMetadata(conn *gorm.DB, prodMetadata *models.ProductoMetadata) error {

	tablaProdMetadata := models.TABLA_PRODUCTO_METADATA

	result := conn.Table(tablaProdMetadata).Create(prodMetadata)

	return result.Error
}


func InsertarProducto(conn *gorm.DB, prod *models.Producto) (int, error) {
	
	tablaProd := models.TABLA_PRODUCTO

	result := conn.Table(tablaProd).Create(prod)
	return prod.Id, result.Error
}

func ListaDeProductos(conn *gorm.DB) (*[]models.Producto, error) {

	tablaProd := models.TABLA_PRODUCTO
	var prods []models.Producto

	result := conn.Table(tablaProd).Select("*").Find(prods)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &prods, nil
}

func ActualizarProducto(idProd int, prod *models.Producto, conn *gorm.DB) error{
	tablaProd := models.TABLA_PRODUCTO

	result:= conn.Table(tablaProd).Where("Id = ?", idProd).Updates(map[string]interface{}{"Id": idProd, "Nombre":prod.Nombre, "Descripcion": prod.Descripcion, "Stock": prod.Stock, "IdCategoria": prod.IDCategoria, "Precio": prod.Precio})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return  result.Error
	}
	return nil
}
