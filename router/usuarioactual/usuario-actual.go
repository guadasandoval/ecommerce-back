package usuarioactual

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/log"
	"ecommerce/router/common"
	"encoding/json"
	"net/http"
	"ecommerce/pkg/usuarioactual"
	"github.com/julienschmidt/httprouter"
)

// GetMisDatos function
func GetMisDatos(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {

	w.Header().Set("Content-Type", "application/json")

	const op errors.Operation = "router.usuarioactual.usuarioactual.GetMisDatos"

	misDatos, err := usuarioactual.GetFormInputMisDatos(&usuario)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	// Response
	response := map[string]interface{}{
		"status": "OK",
		"data":   misDatos,
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	return
}

// CambiarContrasena function
func CambiarContrasena(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo,
) {
	const op errors.Operation = "router.usuarioactual.usuarioactual.CambiarContrasena"

	// Input
	var input usuarioactual.FormInputCambiarContrasena

	eerr := json.NewDecoder(r.Body).Decode(&input)
	if eerr != nil {
		log.Error2(op, eerr)
		common.SendErr(w, http.StatusBadRequest, eerr.Error())
		return
	}
	// Logic
	response := updatePassUsrActual(input, usuario)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	return
}

func updatePassUsrActual(input usuarioactual.FormInputCambiarContrasena, usuario models.Usuario) map[string]interface{} {
	const op errors.Operation = "router.usuarioactual.usuarioactual.updatePassUsrActual"

	err := input.Update(usuario)

	var response map[string]interface{}

	if err != nil {
		if errors.Is(err, usuarioactual.ErrInvalidPassword) {
			response = map[string]interface{}{
				"status": "ERROR",
				"error":  "La contraseña actual es inválida",
			}
		} else {
			log.Error2(op, err)
			response = map[string]interface{}{
				"status": "ERROR",
				"error":  err.Error(),
			}
		}
	} else {
		response = map[string]interface{}{
			"status": "OK",
		}
	}

	return response
}
