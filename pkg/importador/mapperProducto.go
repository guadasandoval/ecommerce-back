package importador

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/db/services"
	"ecommerce/pkg/services/log"
	"github.com/tealeg/xlsx/v3"
	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

func mapearProducto(datoProd *xlsx.Row) (*models.Producto, error) {

	var producto models.Producto

	precio, _ := leerComoInt(datoProd.GetCell(colPrecio))
	producto.Precio = precio
	nombre := leerComoString(datoProd.GetCell(colNombre))
	producto.Nombre = nombre
	descripcion := leerComoString(datoProd.GetCell(colDescripcion))
	producto.Descripcion = descripcion
	stock := leerComoBool(datoProd.GetCell(colStock))
	producto.Stock = stock
	idCat, _ := leerComoInt(datoProd.GetCell(colCategoria))
	producto.IDCategoria = idCat

	return &producto, nil
}

func completarProductoMetadata(productoMetadata *models.ProductoMetadata, ocurrioError bool, err error) error {
	const op errors.Operation = "pkg.importador.mapperproducto.completarproductoMetadata"
	productoMetadata.OcurrioError = ocurrioError

	if ocurrioError {
		productoMetadata.Error = null.NewString(err.Error(), true)
	} else {
		productoMetadata.Error = null.NewString("", false)
	}

	err = services.InsertarProductoMetadata(db.GetDB(), productoMetadata)
	if err != nil {
		log.Error2(op, err)
		return errors.NewError(op, errors.KindUnexpected, err)
	}
	return nil
}

func insertarProducto(personaF *models.Producto, row *xlsx.Row, tx *gorm.DB) error {
	const op errors.Operation = "pkg.importador.mapperPersona.insertarproducto"

	_, err := services.InsertarProducto(tx, personaF)
	if err != nil {
		log.Error2(op, err)
		return errors.NewError(op, errors.KindUnexpected, err)
	}

	return nil
}


