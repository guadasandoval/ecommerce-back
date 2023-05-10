package services

import (
	"ecommerce/pkg/services/db/models"
	"errors"

	"gorm.io/gorm"
)


func UsuarioCreate(conn *gorm.DB, usuario *models.Usuario) error {

	tablaUsuarios := models.TABLE_USUARIOS

	result := conn.Table(tablaUsuarios).Create(usuario)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRegistered) {
			return nil
		}

		return result.Error
	}
	return nil
}

func UsuarioUpdate(conn *gorm.DB, usuario *models.Usuario) error {
	tablaUsuarios := models.TABLE_USUARIOS
	result := conn.Table(tablaUsuarios).Where("usr_id = ?", usuario.ID).Updates(map[string]interface{}{"usrrol_id": usuario.RolID, "usrgrp_id": usuario.GrupoID, "usr_email": usuario.Email, "usr_nombres": usuario.Nombres, "usr_apellidos": usuario.Apellidos, "usr_habilitado": usuario.Habilitado})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRegistered) {
			return nil
		}
		return result.Error
	}
	return nil
}

func UsuarioUpdatePassword(conn *gorm.DB, idUsr int, passUsr string) error {
	tablaUsuarios := models.TABLE_USUARIOS
	result := conn.Table(tablaUsuarios).Where("usr_id = ?", idUsr).Update("usr_password", passUsr)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRegistered) {
			return nil
		}
		return result.Error
	}
	return nil
}

// UsuarioFindByEMail function
func UsuarioFindByEMail(conn *gorm.DB, email string) (*models.Usuario, error) {
	tablaUsuarios := models.TABLE_USUARIOS
	var usuario models.Usuario

	result := conn.Table(tablaUsuarios).Where("usr_email = ?", email).Find(&usuario)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &usuario, nil

}

//BuscarPersonaPorID Servicio que obtiene una persona con los datos completos
func BuscarUsuarioPorID(conn *gorm.DB, ID int) (*models.Usuario, error) {

	tablaUsuarios := models.TABLE_USUARIOS

	var usuario models.Usuario
	result := conn.Table(tablaUsuarios).Where("usr_id = ?", ID).Find(&usuario)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &usuario, result.Error
}

func UsuariosTable(conn *gorm.DB) (*[]models.Usuario, error) {

	tablaUsuarios := models.TABLE_USUARIOS
	var usuarios []models.Usuario

	result := conn.Table(tablaUsuarios).Find(&usuarios)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &usuarios, nil
}
