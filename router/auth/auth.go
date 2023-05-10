package auth

import (
	"ecommerce/pkg/auth"
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/log"
	"ecommerce/router/common"
	"encoding/json"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Login function
func LoginNuevo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	const op errors.Operation = "router.auth.auth.LoginNuevo"

	// Input
	var input auth.FormInputLogin

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	// Validation
	// validationInput := validation.CreateValidation(input)
	// if !validationInput.Valid() {
	// 	log.Error(op, errors.Msg("Invalid validation"), zap.String("Validation", validationInput.ErrorsDebug()))
	// 	sendErr(w, http.StatusUnprocessableEntity)
	// 	return
	// }

	// Logic
	loginResponse, err := input.Login()

	if err != nil {
		if errors.Is(err, auth.ErrUserEMailNotFound) || errors.Is(err, auth.ErrUserInvalidPassword) {
			log.Error2(op, err)
			response := map[string]interface{}{
				"status": "ERROR",
				"error":  "Correo electrónico o contraseña inválida",
			}

			w.WriteHeader(http.StatusOK)

			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				log.Error2(op, err)
				common.SendErr(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		if errors.Is(err, auth.ErrUserDisabled) {
			log.Error2(op, err)
			response := map[string]interface{}{
				"status": "ERROR",
				"error":  "El usuario se encuentra deshabilitado",
			}
			w.WriteHeader(http.StatusOK)

			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				log.Error2(op, err)
				common.SendErr(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		log.Error2(op, err)
		response := map[string]interface{}{
			"status": "ERROR",
			"error":  err,
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Error2(op, err)
			common.SendErr(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Response
	response := map[string]interface{}{
		"status": "OK",
		"data":   loginResponse,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
	return
}

// CheckLogin function
func CheckLogin(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {
	const op errors.Operation = "router.auth.auth.CheckLogin"

	// Response
	response := map[string]interface{}{
		"status": "OK",
	}

	//w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	return
}

// Logout function
func Logout(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {
	const op errors.Operation = "router.auth.auth.Logout"

	var usuarix auth.UsuarioToken

	err := json.NewDecoder(r.Body).Decode(&usuarix)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	// Logic
	err = auth.Logout(usuarix.Token)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Response
	response := map[string]interface{}{
		"status": "OK",
	}

	//w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
	return
}
