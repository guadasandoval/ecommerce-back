package productos

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/db/services"
	"ecommerce/pkg/services/log"
)

func ListarProductos() (*[]models.Producto, error) {
	const op errors.Operation = "pkg.productos.productos.ListarProductos"
	conn := db.GetDB()

	productosList, err := services.ListaDeProductos(conn)
	if err != nil {
		err = errors.Msg("Ocurrió un error en ListaDeProductos")
		log.Error2(op, err)
		return nil, errors.NewError(op, errors.KindUnexpected, err)
	}
	
	log.Debug("Se listaron los prods", log.Int("cantidadDeproductosList", len(*productosList)))

	return productosList, nil
}