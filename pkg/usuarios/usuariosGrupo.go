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
func UsuariosGrupoFind(ID int) (*models.UsuariosGrupo, error) {
	const op errors.Operation = "pkg.usuarios.usuariosGrupo.UsuariosGrupoFind"
	conn := db.GetDB()

	usuarioGrupo, err := services.BuscarUsuarioGrupoPorID(conn, ID)
	if err != nil {
		log.Error2(op, err, zap.String("usuarioGrupoID", strconv.Itoa(ID)))
		return nil, errors.NewError(op, errors.KindNotFound, err)
	} else {
		log.Debug("Busqueda de grupo usuario finalizada correctamente", log.Int("usuarioGrupoID", ID))
	}
	return usuarioGrupo, nil
}

// GetRolSelectOptions function
func GetGruposSelectOptions() (*[]models.UsuariosGrupo, error) {
	const op errors.Operation = "pkg.usuarios.usuariosGrupo.GetGruposSelectOptions"
	conn := db.GetDB()

	grupos, err := services.UsuariosGruposAll(conn)
	if err != nil {
		err = errors.Msg("Ocurrió un error en UsuariosGruposAll")
		log.Error2(op, err)
		return nil, errors.NewError(op, errors.KindNotFound, err)
	} else {
		log.Debug("Busqueda de rol usuario finalizada correctamente")
	}

	return grupos, nil
}
