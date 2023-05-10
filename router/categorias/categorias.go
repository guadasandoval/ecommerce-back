package categorias

import (
	"ecommerce/pkg/categorias"
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/log"
	"ecommerce/router/common"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CategoriasProductos(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {

	w.Header().Set("Content-Type", "application/json")
	const op errors.Operation = "router.categorias.categorias.CategoriasProductos"

	response := obtenerCategorias()

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}


func obtenerCategorias() map[string]interface{} {
	const op errors.Operation = "router.categorias.categorias.obtenerCategorias"

	categorias, err := categorias.Categorias()

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
			"data":   categorias,
		}
	}

	return response
}

