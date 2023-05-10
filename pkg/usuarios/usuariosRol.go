package usuarios

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/db/services"
	"ecommerce/pkg/services/log"
	"strconv"

	"go.uber.org/zap"
)

// UsuariosRolFind function
func UsuariosRolFind(ID int) (*models.UsuariosRol, error) {
	const op errors.Operation = "pkg.usuarios.usuariossRol.UsuariosRolFind"
	conn := db.GetDB()

	usuarioRol, err := services.BuscarUsuarixRolPorID(conn, ID)
	if err != nil {
		log.Error2(op, err, zap.String("usuarioID", strconv.Itoa(ID)))
		return nil, errors.NewError(op, errors.KindNotFound, err)
	} else {
		log.Debug("Busqueda de rol usuario finalizada correctamente", log.Int("usuarioRolID", ID))
	}
	return usuarioRol, nil
}

// GetRolSelectOptions function
func GetRolSelectOptions() (*[]models.UsuariosRol, error) {
	const op errors.Operation = "pkg.usuarios.usuariosRol.GetRolSelectOptions"
	conn := db.GetDB()

	roles, err := services.UsuariosRolesAll(conn)
	if err != nil {
		log.Error2(op, err)
		return nil, errors.NewError(op, errors.KindNotFound, err)
	} else {
		log.Debug("Busqueda de rol usuario finalizada correctamente")
	}

	return roles, nil
}
