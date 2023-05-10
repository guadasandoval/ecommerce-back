package categorias

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/db/services"
	"ecommerce/pkg/services/log"
)

func Categorias() (*[]models.Categorias, error) {
	const op errors.Operation = "pkg.categorias.categorias.Categorias"

	conn := db.GetDB()
	categorias, err := services.CategoriasProductos(conn)
	if err != nil {
		log.Error2(op, err)
		return nil, errors.NewError(op, errors.KindNotFound, err)
	} else {
		log.Debug("Se listaron los categorias", log.Int("categorias", len(*categorias)))
	}

	return categorias, nil
}



