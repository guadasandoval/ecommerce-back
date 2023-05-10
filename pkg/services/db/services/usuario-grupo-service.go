package services

import (
	"ecommerce/pkg/services/db/models"
	"errors"

	"gorm.io/gorm"
)

//BuscarPersonaPorID Servicio que obtiene una persona con los datos completos
func BuscarUsuarioGrupoPorID(conn *gorm.DB, grupoID int) (*models.UsuariosGrupo, error) {

	tablaUsuariosGrupo := models.TABLE_USUARIOSGRUPOS

	var usuarioGrupo models.UsuariosGrupo
	result := conn.Table(tablaUsuariosGrupo).Where("grp_id = ?", grupoID).Find(&usuarioGrupo)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &usuarioGrupo, result.Error
}

// UsuariosGruposAll lista de grupos
func UsuariosGruposAll(conn *gorm.DB) (*[]models.UsuariosGrupo, error) {
	tablaGruposUsuarios := models.TABLE_USUARIOSGRUPOS

	var usuariosGrupos []models.UsuariosGrupo

	result := conn.Table(tablaGruposUsuarios).Scan(&usuariosGrupos)

	return &usuariosGrupos, result.Error
}
