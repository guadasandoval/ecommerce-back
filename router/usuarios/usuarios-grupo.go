package usuarios

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/permission"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/log"
	"ecommerce/pkg/usuarios"
	"ecommerce/router/common"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetUsuarioGruposSelect(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {

	w.Header().Set("Content-Type", "application/json")

	const error = "Invalid authorization"
	const op errors.Operation = "router.usuarios.usuarios.GetUsuarioGruposSelect"

	// Permission
	if !permission.CanRead(rol.ModUsuarios) {
		common.SendErr(w, http.StatusForbidden, error)
		return
	}

	response := listaDeGruposResult()

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func listaDeGruposResult() map[string]interface{} {
	const op errors.Operation = "router.usuarios.usuarios.listaDeGruposResult"

	listaRoles, err := usuarios.GetGruposSelectOptions()

	var response map[string]interface{}

	if err != nil {
		log.Error2(op, err)
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {

		response = map[string]interface{}{
			"status": "OK",
			"data":   listaRoles,
		}
	}

	return response
}
