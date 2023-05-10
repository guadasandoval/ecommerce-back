package router

import (
	"ecommerce/pkg/auth"
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/log"
	"ecommerce/router/common"
	"net/http"
	"time"
	"ecommerce/pkg/usuarios"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// Handle type
type Handle func(
	http.ResponseWriter,
	*http.Request,
	httprouter.Params,
)

// HandleWithUsuario type
type HandleWithUsuario func(
	http.ResponseWriter,
	*http.Request,
	httprouter.Params,
	string,
	models.Usuario,
	models.UsuariosRol,
	models.UsuariosGrupo,
)

// MiddlewareLog function
func MiddlewareLog(handler Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		startTime := time.Now()
		handler(w, r, ps)
		now := time.Now()
		timeDur := now.Sub(startTime)

		log.Info("MiddlewareLog", zap.Duration("duration", timeDur), zap.String("method", r.Method),
			zap.String("path", r.URL.Path))
	}
}

// MiddlewareUsuario function
func MiddlewareUsuario(handler HandleWithUsuario) httprouter.Handle {

	return MiddlewareLog(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		token, err := getTokenFromRequest(r)
		if err != nil {
			common.SendErr(w, http.StatusUnauthorized, err.Error())
			return
		}

		usuarioID, err := auth.GetLogin(token)
		if err != nil {
			common.SendErr(w, http.StatusUnauthorized, err.Error())
			return
		}

		usuario, err := usuarios.BuscarUsuarioPorID(usuarioID)
		if err != nil {
			common.SendErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		if usuario.ID == 0 {
			err := errors.Msg("Usuario Id '%d' not found", usuario.ID)
			common.SendErr(w, http.StatusUnauthorized, err.Error())
			return
		}

		rol, err := usuarios.UsuariosRolFind(usuario.RolID)
		if err != nil {
			common.SendErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		grupo, err := usuarios.UsuariosGrupoFind(usuario.GrupoID)
		if err != nil {
			common.SendErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		handler(w, r, ps, token, *usuario, *rol, *grupo)

	})
}

func getTokenFromRequest(r *http.Request) (string, error) {

	authorization := r.Header.Get("Authorization")

	if len(authorization) <= 7 {
		err := errors.Msg("No Authorization header")
		return "", err
	}
	if authorization[:7] != "Bearer " {
		err := errors.Msg("Invalid Authorization header")
		return "", err
	}

	return authorization[7:], nil
}
