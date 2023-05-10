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
	"strconv"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// RegistrarNuevoUsuarix método POST crea usuarie en la bd
func RegistrarNuevoUsuario(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {

	w.Header().Set("Content-Type", "application/json")
	const op errors.Operation = "router.usuarios.usuarios.RegistrarNuevoUsuario"
	var usuarioForm usuarios.FormInputUsuario

	json.NewDecoder(r.Body).Decode(&usuario)

	response := registrarUsuario(&usuarioForm)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateUsuario(w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {

	const op errors.Operation = "router.usuarios.usuarios.UpdateUsuario"

	// Permission
	if !permission.CanWrite(rol.ModUsuarios) {
		common.SendErr(w, http.StatusForbidden, "Invalid authorization")
		return
	}

	// Input
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Error2(op, err, zap.String("id", ps.ByName("id")))
		common.SendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	var input usuarios.FormInputUsuario
	eerr := json.NewDecoder(r.Body).Decode(&input)
	if eerr != nil {
		log.Error2(op, eerr)
		common.SendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	if id == 0 && len(input.Password) < 6 {
		common.SendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	response := update(input, id)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	
	return

}

// GetUsuario function
func GetUsuario(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo,
) {
	const op errors.Operation = "router.usuarios.usuarios.GetUsuario"

	// Permission
	if !permission.CanRead(rol.ModUsuarios) {
		common.SendErr(w, http.StatusForbidden, "Invalid authorization")
		return
	}

	// Input
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	// Logic
	usuarioInfo, eerr := usuarios.GetFormInputUsuario(id)
	if eerr != nil {
		log.Error2(op, eerr)
		common.SendErr(w, http.StatusInternalServerError, eerr.Error())
		return
	}

	// Response
	response := map[string]interface{}{
		"status": "OK",
		"data":   usuarioInfo,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	return
}

// UpdateUsuarioPassword function
func UpdateUsuarioPassword(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo,
) {
	const op errors.Operation = "router.usuarios.usuarios.UpdateUsuarioPassword"

	// Permission
	if !permission.CanWrite(rol.ModUsuarios) {
		common.SendErr(w, http.StatusForbidden, "Invalid authorization")
		return
	}

	// Input
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Error2(op, err, zap.String("id", ps.ByName("id")))
		common.SendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	var input usuarios.FormInputPassword
	eerr := json.NewDecoder(r.Body).Decode(&input)
	if eerr != nil {
		log.Error2(op, eerr)
		common.SendErr(w, http.StatusBadRequest, eerr.Error())
		return
	}

	// Logic
	response := updatePass(input, id)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	return
}

func ListarUsuariosTable(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {
	const op errors.Operation = "router.usuarios.usuarios.ListarUsuariosTable"

	response := listaDeUsuarios()
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	return
}

func listaDeUsuarios() map[string]interface{} {
	const op errors.Operation = "router.usuarios.usuarios.listaDeUsuarios"

	usuarios, err := usuarios.ListarUsuariosTable()

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
			"data":   usuarios,
		}
	}

	return response
}

func update(input usuarios.FormInputUsuario, id int) map[string]interface{} {
	const op errors.Operation = "router.usuarios.usuarios.update"

	err := input.Update(id)

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
			"data":   id,
		}
	}

	return response
}

func updatePass(input usuarios.FormInputPassword, id int) map[string]interface{} {
	const op errors.Operation = "router.usuarios.usuarios.updatePass"

	err := input.UpdatePassword(id)

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
			"data":   id,
		}
	}

	return response
}

func IsUsuarioEMailAvailable(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {

	const error = "Invalid authorization"
	const op errors.Operation = "router.usuarios.usuarios.IsUsuarioEMailAvailable"

	// Permission
	if !permission.CanRead(rol.ModUsuarios) {
		common.SendErr(w, http.StatusForbidden, error)
		return
	}

	// Input
	var input usuarios.FormInputEMailAvailable
	json.NewDecoder(r.Body).Decode(&input)


	// Logic
	mailAvailable, err := input.Check()
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	// Response
	response := map[string]interface{}{
		"status": "OK",
		"data":   mailAvailable,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	return

}

func registrarUsuario(usuarioForm *usuarios.FormInputUsuario) map[string]interface{} {
	const op errors.Operation = "router.usuarios.usuarios.registrarUsuario"
	usuarixID, err := usuarioForm.CrearUsuario()

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
			"data":   usuarixID,
		}
	}

	return response
}
