package services

import (
	"ecommerce/pkg/services/db/models"
	"errors"

	"gorm.io/gorm"
)

//BuscarPersonaPorID Servicio que obtiene una persona con los datos completos
func BuscarUsuarixRolPorID(conn *gorm.DB, rolID int) (*models.UsuariosRol, error) {

	tablaUsuarixsRol := models.TABLE_USUARIOSROLES

	var usuarixRol models.UsuariosRol
	result := conn.Table(tablaUsuarixsRol).Where("rol_id = ?", rolID).Find(&usuarixRol)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &usuarixRol, result.Error
}

// UsuariosRolesAll lista los roles
func UsuariosRolesAll(conn *gorm.DB) (*[]models.UsuariosRol, error) {
	tablaUsuariosRoles := models.TABLE_USUARIOSROLES

	var usuariosRoles []models.UsuariosRol

	result := conn.Table(tablaUsuariosRoles).Scan(&usuariosRoles)

	return &usuariosRoles, result.Error
}
