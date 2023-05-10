package importador

import (
	"ecommerce/pkg/importador"
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/log"
	"ecommerce/router/common"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const formFileName = "nombre"

func ImportarProductos(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params,
	token string,
	usuario models.Usuario,
	rol models.UsuariosRol,
	grupo models.UsuariosGrupo) {
	const op errors.Operation = "router.importador.importador.ImportarProductos"

	formFile, fileHeader, err := r.FormFile(formFileName)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}

	bytesData, errioutil := ioutil.ReadAll(formFile)
	if errioutil != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, errioutil.Error())
	}

	response := importarProds(bytesData, fileHeader.Filename, usuario)

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error2(op, err)
		common.SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func importarProds(bytesData []byte, fileName string, usuario models.Usuario) map[string]interface{} {
	const op errors.Operation = "router.importador.importador.importarProds"

	cargados, totales, err := importador.ImportarProductos(bytesData, fileName)

	var response map[string]interface{}

	if err != nil {
		log.Error2(op, err)
		response = map[string]interface{}{
			"status": "ERROR",
			"error":  err.Error(),
		}
	} else {

		response = map[string]interface{}{
			"status":                    "OK",
			"dataCargadosCorrectamente": cargados,
			"dataProductosTotales":      totales,
		}
	}

	return response
}