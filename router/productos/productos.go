package productos

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/productos"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/log"
	"ecommerce/router/common"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ListarProductos(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {
	w.Header().Set("Content-Type", "application/json")
	const op errors.Operation = "router.productos.productos.ListarProductos"

	response := listaDeProductos()

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func listaDeProductos() map[string]interface{} {
	const op errors.Operation = "router.productos.productos.listaDeProductos"

	listaProductos, err := productos.ListarProductos()

	var response map[string]interface{}

	if err != nil {
		log.Error2(op, err)
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {
		response = map[string]interface{}{
			"status":        "OK",
			"data":          listaProductos,
		}
	}

	return response
}