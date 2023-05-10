package importador

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/log"
	"fmt"
	"github.com/tealeg/xlsx/v3"
	"gopkg.in/guregu/null.v3"
)

func ImportarProductos(bytesData []byte, fileName string) (int, int, error) {
	const op errors.Operation = "pkg.importador.importador.ImportarProductos"
	
	fileData, err := xlsx.OpenBinary(bytesData)
	if err != nil {
		log.Error2(op, err)
		return 0, 0, err
	}

	archivoProds := fileData.Sheet[hojaProd]
	if archivoProds == nil {
		log.Error2(op, err)
		return 0, 0, fmt.Errorf("No existe la hoja: %s en el archivo %s", hojaProd, fileName)
	}

	insertadosCorrectamente := leerArchivoEInsertarProds(archivoProds, fileName)
	log.Info("Carga de productos finalizada", log.Int("productosCargados", insertadosCorrectamente), log.Int("productosTotales", archivoProds.MaxRow-1))

	return insertadosCorrectamente, archivoProds.MaxRow - 1, nil
}

func leerArchivoEInsertarProds(datosProds *xlsx.Sheet, fileName string) int {
	const op errors.Operation = "pkg.importador.importador.leerArchivoEInsertarProds"

	insertadosCorrectamente := 0
	
	conn := db.GetDB()
	
	for i := iniciodatosProds; i < datosProds.MaxRow; i++ {

		row, err := datosProds.Row(i)
		if err != nil {
			log.Error2(op, err)
			continue
		}
		prodMetadata := models.ProductoMetadata{NombreArchivo: fileName, LineaArchivo: i + 1}

		producto, err := mapearProducto(row)
		if err != nil {
			log.Error2(op, err)
			continue
		}

		tx := conn.Begin()
		var idProd int	

		 		err = insertarProducto(producto, row, tx)
				if err != nil {
					tx.Rollback()
					log.Error2(op, err)
					completarProductoMetadata(&prodMetadata, true, err)
					continue
				}
			
				idProd = producto.Id
		prodMetadata.IDProducto = null.NewInt(int64(idProd), true)
		completarProductoMetadata(&prodMetadata, false, nil)
		tx.Commit()
		insertadosCorrectamente++	
		
	}

		return insertadosCorrectamente
	}
